package fileutil

import (
	"testing"
)

func TestFile2String(t *testing.T) {
	bytes, err := File2Byte("/Users/dji/Downloads/a.txt")
	if err != nil {
		t.Error(err)
	}
	t.Error(len(bytes), err)
}

func TestFileExists(t *testing.T) {
	_, err := FileExists("/notexists")
	if err != nil {
		t.Error(err)
	}
}
