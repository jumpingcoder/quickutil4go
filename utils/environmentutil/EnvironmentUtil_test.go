package environmentutil

import "testing"

func TestFile2String(t *testing.T) {
	path, _ := HomePath()
	t.Log(path)
}
