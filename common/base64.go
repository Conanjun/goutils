package common

import (
	"encoding/base64"
	"log"
)

func base64encodebytes2string(input []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func base64decodestring(encodeString string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	return string(decodeBytes)
}

func base64urlencodebytes2string(input []byte) string {
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	return uEnc
}

func base64decodeurlstring(encodeString string) string {
	uDec, err := base64.URLEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	return string(uDec)
}
