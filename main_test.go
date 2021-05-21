package main

import (
	"github.com/jumpingcoder/quickutil4go/utils/cryptoutil"
	"testing"
)

func TestDecryptConfig(t *testing.T) {
	t.Log(string(cryptoutil.Base64Decrypt("77u/MTIz")))
}
