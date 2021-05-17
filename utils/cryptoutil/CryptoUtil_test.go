package cryptoutil

import (
	"testing"
)

func TestAESCBCEncrypt(t *testing.T) {
	content := "zxcvbnm"
	key := "1234567812345678"
	encrypted := AESCBCEncrypt([]byte(content), []byte(key), make([]byte, 16))
	t.Log(Base64Encrypt(encrypted))
	decrypted := AESCBCDecrypt(encrypted, []byte(key), make([]byte, 16))
	t.Log(string(decrypted))
}
