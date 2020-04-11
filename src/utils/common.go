package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func GenerateUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return strings.ToUpper(uuid)
}

func GenerateRandom(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x", b)
	fmt.Println(uuid)
	return strings.ToUpper(uuid)
}

func EncodeBase64(plainText string) string {
	encodedText := base64.StdEncoding.EncodeToString([]byte(plainText))
	return encodedText
}

func DecodeBase64(cipherText string) (string, error) {
	decodedText, err := base64.StdEncoding.DecodeString(string(cipherText))
	return string(decodedText), err
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func DoesFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func GenerateSHA256Hash(plainText []byte) string {
	hashVal := sha256.New()
	hashVal.Write(plainText)
	return base64.URLEncoding.EncodeToString(hashVal.Sum(nil))
}

func VerifySHA256Hash(cipherText string, plainText []byte) bool {
	hashVal := sha256.New()
	hashVal.Write(plainText)
	cipherPlain := base64.URLEncoding.EncodeToString(hashVal.Sum(nil))
	return strings.EqualFold(cipherPlain, cipherText)
}
