// DES、3DES加解密示例，用于生产环境请修改代码使之健壮
package helpers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func testDes() {
	key := []byte("sfe023f_")
	result, err := DesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

//func test3Des() {
//	key := []byte("sfe023f_sefiel!fi32lf3e?")
//	result, err := TripleDesEncrypt([]byte("polaris@studygol"), key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(base64.StdEncoding.EncodeToString(result))
//	origData, err := TripleDesDecrypt(result, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return string(crypted), nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}

// 3DES加密 返回string
func TripleDesEncryptString(origDataStr, keyStr string) (string, error) {
	origData := []byte(origDataStr)

	//若keyStr不够24个字符  用0补齐到24个
	fillKey := FillStringToLength(keyStr, 24)
	key := []byte(fillKey)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	cryptedStr := base64.StdEncoding.EncodeToString(crypted)
	urlCrypt := url.QueryEscape(cryptedStr)
	return urlCrypt, nil
}

// 3DES解密 返回string
func TripleDesDecryptString(strCrypted, strkey string) (string, error) {
	subkey := FillStringToLength(strkey, 24)
	key := []byte(subkey)
	//urlUnescape, _ := url.QueryUnescape(strCrypted)

	crypted, _ := base64.StdEncoding.DecodeString(strCrypted)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}

func FillStringToLength(key string, length int) string {
	keyByte := []byte(key)
	var subByte []byte
	var subString string
	if len(key) >= 24 {
		subByte = keyByte[0:24]
		subString = string(subByte)
	} else {
		num := 24 - len(key)
		subString = key + strings.Repeat("0", num)
	}
	return subString
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
