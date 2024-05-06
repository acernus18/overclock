package overclock

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

func generateIV(key string) []byte {
	iv := md5.Sum([]byte(key))
	return []byte(hex.EncodeToString(iv[:])[12:28])
}

func pkcs7padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 1 || blockSize >= 256 {
		return nil, fmt.Errorf("PKCS7: invalid block size %d", blockSize)
	}
	length := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(length)}, length)...), nil
}

func pkcs7strip(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("PKCS7: data is empty")
	}
	if len(data)%blockSize != 0 {
		return nil, errors.New("PKCS7: data is not block-aligned")
	}
	length := int(data[len(data)-1])
	if length > blockSize || length == 0 || !bytes.HasSuffix(data, bytes.Repeat([]byte{byte(length)}, length)) {
		return nil, errors.New("PKCS7: invalid padding")
	}
	return data[:len(data)-length], nil
}

func encryptAES(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	content, err := pkcs7padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(content))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(result, content)
	return result, nil
}

func decryptAES(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(result, data)
	content, err := pkcs7strip(result, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return content, nil
}

func EncryptContent(key, data string) (string, error) {
	encryptedContent, err := encryptAES([]byte(key), generateIV(key), []byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encryptedContent), nil
	//result := hex.EncodeToString(encryptedContent)
	//rows := make([]string, len(result)/(len(key)*2))
	//for i := range rows {
	//	rows[i] = result[i*32 : i*32+32]
	//}
	//return strings.Join(rows, "\n"), nil
}

func DecryptContent(key, data string) (string, error) {
	contentBytes, err := hex.DecodeString(data)
	//contentBytes, err := hex.DecodeString(strings.Join(strings.Split(data, "\n"), ""))
	if err != nil {
		return "", err
	}
	//iv := md5.Sum([]byte(key))
	originData, err := decryptAES([]byte(key), generateIV(key), contentBytes)
	if err != nil {
		return "", err
	}
	return string(originData), nil
}
