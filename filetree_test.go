package filetree_test


import (
    "github.com/hahnicity/filetree"
    "path"
    "os"
    "testing"
)

func TestGetDirSuccess(t *testing.T) {
    dirname, _ := os.Getwd()
    d, err := filetree.GetDir(dirname)
    CheckDirError(t, dirname, err)
    f, err := os.OpenFile(dirname, os.O_RDONLY, os.ModeDir)
    defer f.Close()
    expInfo, _ := f.Stat()
    if expInfo.Name() != d.Info.Name() {
        t.Errorf("Expected dir name: " + expInfo.Name() + " Actual dir Name: " + d.Info.Name())    
    }
    if path.Dir(dirname) != d.Path {
        t.Errorf("Expected dir path: " + path.Dir(dirname) + " Actual dir path " + d.Path)    
    }
}

func TestGetDirFailure(t *testing.T) {
    invalidDir := "FHALKFOOLAJFLKAHJ"
    _, err := filetree.GetDir(invalidDir)
    if err == nil {
        t.Errorf("No error was raised using an invalid directory name: " + invalidDir)
    }
}

func TestGetFilePaths(t *testing.T) {
    dirname, _ := os.Getwd()
    d, err := filetree.GetDir(dirname)
    CheckDirError(t, dirname, err)
    files, err := d.GetFilePaths()
    if err != nil {
        t.Errorf("You were unable to get files from " + dirname)   
    }
    if len(files) == 0 {
        t.Errorf("No files were found in your current working directory")
    }
}

func CheckDirError(t *testing.T, dirname string, err error) {
    if err != nil {
        t.Errorf("The provided directory " + dirname + " was invalid")    
    }
}
