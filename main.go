package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

func LogInfo(info string) {
	
}

func LogError(err string) {
	panic(err)
}

func Encode(source, target string, password []byte) {
	if len(password) != 32 {
		LogError("密码应该为32位")
	}
	
	sourceF, err := os.OpenFile(source, os.O_RDONLY, 0777)
	if err != nil {
		LogError("错误：" + err.Error())
	}
	defer sourceF.Close()
	stat, err := sourceF.Stat()
	if err != nil {
		LogError("错误：" + err.Error())
	}
	targetF, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		LogError("错误：" + err.Error())
	}
	defer targetF.Close()
	
	LogInfo("AES - 导出中")
	var iv = RandomIV()
	sourceBytes := make([]byte, stat.Size())
	sourceF.Read(sourceBytes)
	b := AesCBCEncrypt(sourceBytes, password, iv)
	targetF.Write(b)
	LogInfo("AES - 导出完成")
}

// AesCBCEncrypt AES/CBC/PKCS7Padding 加密
func AesCBCEncrypt(content []byte, key []byte, iv []byte) (dst []byte) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		LogError(err.Error())
	}
	
	// PKCS7 填充
	src := PaddingPKCS7(content, aes.BlockSize)
	dst = make([]byte, len(src))
	
	// CBC 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dst, src)
	return
}

// PaddingPKCS7 PKCS7 填充
func PaddingPKCS7(content []byte, blockSize int) []byte {
	paddingSize := blockSize - len(content)%blockSize
	paddingData := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(content, paddingData...)
}
func RandomIV() []byte {
	b := make([]byte, aes.BlockSize)
	rand.Read(b)
	return b
}

// func ToPassword(s string) []byte {
// 	b := make([]byte, 32)
// 	_ = append(b[:0], s...)
// 	return b
// }

func Password() (b []byte) {
	b = make([]byte, 32)
	f, _ := os.OpenFile("password.txt", os.O_RDONLY, 0777)
	f.Read(b)
	return
}

func Copy(source, target string) {
	
}

func main() {
	Encode("./春语.html", "./docs/chunyu.bin", Password())
}
