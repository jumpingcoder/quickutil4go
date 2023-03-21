package cryptoutil

import (
	"testing"
)

func TestEncryptConfigHandler(t *testing.T) {
	t.Log(EncryptConfigHandler("12345678", "0000000011111111"))
}

func TestDecryptConfigHandler(t *testing.T) {
	t.Log(DecryptConfigHandler("ENC(ozPh7o7XOvEup69IjSAbOg==)", "0000000011111111"))
}

func TestMD5Encrypt(t *testing.T) {
	t.Log(MD5Encrypt([]byte("123456")))
}

func TestHmacMD5Encrypt(t *testing.T) {
	t.Log(HmacMD5("test0",[]byte("zxcv1234")))
}

func TestBase64Encrypt(t *testing.T) {
	t.Log(Base64Encrypt([]byte("shopee_de_csdata@Inhouse_2021")))
}
