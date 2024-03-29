package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/jumpingcoder/quickutil4go/quickutil4go/logutil"
	"strings"
)

func Base64Encrypt(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

func Base64Decrypt(content string) []byte {
	output, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		logutil.Error("Base64 deserialization failed", err)
	}
	return output
}

func MD5Encrypt(content []byte) string {
	sum := md5.Sum(content)
	return fmt.Sprintf("%x", sum)
}

func SHA1Encrypt(content []byte) string {
	sum := sha1.Sum(content)
	return fmt.Sprintf("%x", sum)
}

func HmacMD5(key string, data []byte) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write(data)
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func HmacSHA1(key string, data []byte) string {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write(data)
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func AESCBCEncrypt(content []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	originData := pad(content, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(originData))
	blockMode.CryptBlocks(encrypted, originData)
	return encrypted
}

func AESCBCDecrypt(content []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(content))
	blockMode.CryptBlocks(decrypted, content)
	return unpad(decrypted)
}

func EncryptConfigHandler(content string, key string) string {
	encrypted := Base64Encrypt(AESCBCEncrypt([]byte(content), []byte(key), make([]byte, 16)))
	return "ENC(" + encrypted + ")"
}

func DecryptConfigHandler(content string, key string) string {
	start := strings.Index(content, "ENC(")
	if start < 0 {
		return content
	}
	end := strings.Index(content[start:len(content)], ")")
	password := content[start+4 : start+end]
	decrypted := string(AESCBCDecrypt(Base64Decrypt(password), []byte(key), make([]byte, 16)))
	newContent := content[0:start] + decrypted + content[start+end+1:len(content)]
	return newContent
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	unpadText := int(ciphertext[length-1])
	return ciphertext[:(length - unpadText)]
}
