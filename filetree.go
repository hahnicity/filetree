package filetree

import (
    "errors"
    "path"
    "os"
)

type Dir struct {
    Path  string
    Info  os.FileInfo
}

func GetDir(dirname string) (*Dir, error) {
    /*Get a directory object*/
    f, err := os.OpenFile(dirname, os.O_RDONLY, os.ModeDir)
    if errorCheck(err) {
        return nil, err    
    }
    defer f.Close()
    Info, err := f.Stat()
    if errorCheck(err) {
        return nil, err    
    }
    if !Info.IsDir() {
        return nil, IsNotDirError(dirname)
    }
    d := new(Dir)
    d.Path = path.Dir(dirname)
    d.Info, _ = f.Stat()
    return d, nil
}

func (d *Dir) GetFilePaths() ([]string, error) {
    /* Return a slice of all file paths in the chosen directory*/    
    allPaths := make([]string, 0)
    f, err := os.OpenFile(path.Join(d.Path, d.Info.Name()), os.O_RDONLY, os.ModePerm)
    files, err := f.Readdir(0)
    if errorCheck(err) {
        return nil, err    
    }
    defer f.Close()
    for _, i := range files {
        if !i.IsDir() {
            allPaths = append(allPaths, path.Join(d.Path, i.Name()))    
        }
    }
    return allPaths, nil
}

// Error Object //
func IsNotDirError(dirname string) error {
    return errors.New(dirname + " is not a valid directory")
}

// This can just as easily be placed into an if block, but I like the readability
// and I'm not that concerned about memory allocation atm
func errorCheck(err error) bool { 
    return err != nil 
}
