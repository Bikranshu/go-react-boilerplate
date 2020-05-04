package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
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

func GenerateSHA256Hash(plainText string) string {
	hashVal := sha256.New()
	hashVal.Write([]byte(plainText))
	return base64.URLEncoding.EncodeToString(hashVal.Sum(nil))
}

func VerifySHA256Hash(cipherText string, plainText string) bool {
	hashVal := sha256.New()
	hashVal.Write([]byte(plainText))
	cipherPlain := base64.URLEncoding.EncodeToString(hashVal.Sum(nil))
	return strings.EqualFold(cipherPlain, cipherText)
}

func IndexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func ContainsString(slice []string, value string) bool {
	return IndexOf(slice, value) != -1
}

// All fields marked with the tag `hidden:"true"` will be set to their zero value.
// For example, this allows to send user models to the client without their password field.
func OmitHiddenFields(model interface{}) {
	if model == nil {
		return
	}
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Struct {
			value := reflect.ValueOf(model).Elem()
			for i := 0; i < t.NumField(); i++ {
				field := value.Field(i)
				fieldType := t.Field(i)

				if !field.CanSet() {
					continue
				}

				if field.Kind() == reflect.Struct && fieldType.Anonymous {
					OmitHiddenFields(field.Addr().Interface())
					continue
				}

				tagName := strings.Split(fieldType.Tag.Get("hidden"), ";")
				if ContainsString(tagName, "true") {
					field.Set(reflect.Zero(fieldType.Type))
				}
			}
		}
	}
}
