package main

import (
	"github.com/jumpingcoder/quickutil4go/utils/cryptoutil"
	"testing"
)

func TestDecryptConfig(t *testing.T) {
	t.Log(decryptConfig("https://hello:ENC(rnz6vt4CTah2XIhtwLoDkw==)@quickutil.com"))
	t.Log(decryptConfig("https://hello:world@quickutil.com"))
	t.Log(string(cryptoutil.Base64Decrypt("NDQxNTEy")))
}
