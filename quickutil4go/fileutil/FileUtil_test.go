package fileutil

import (
	"github.com/jumpingcoder/quickutil4go/quickutil4go/environmentutil"
	"testing"
)

func TestFile2String(t *testing.T) {
	path, err := environmentutil.HomePath()
	if err != nil {
		t.Error(err)
	}
	err = String2File(path+"/a.txt", "test")
	if err != nil {
		t.Error(err)
	}
	bytes, err := File2Byte(path + "/a.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(bytes))
}

func TestFileExists(t *testing.T) {
	exists, err := FileExists("/notexists")
	if err != nil {
		t.Error(err)
	}
	t.Log(exists)
}
