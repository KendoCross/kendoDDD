package utility

import (
	"encoding/xml"
	"errors"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	XmlGbkHeader  = `<?xml version="1.0" encoding="GBK"?>`
	XmlUtf8Header = `<?xml version="1.0" encoding="UTF-8"?>`
)

// GbkToUtf8
func GbkToUtf8(gbkByte []byte) ([]byte, error) {
	b, err := simplifiedchinese.GBK.NewDecoder().Bytes(gbkByte)
	if err != nil {
		err = errors.New("simplifiedchinese.GBK.NewDecoder().Bytes() error : " + err.Error())
	}
	return b, err
}

// Utf8ToGbk
func Utf8ToGbk(utf8Byte []byte) ([]byte, error) {
	b, err := simplifiedchinese.GBK.NewEncoder().Bytes(utf8Byte)
	if err != nil {
		err = errors.New("simplifiedchinese.GBK.NewEncoder().Bytes() error : " + err.Error())
	}
	return b, err
}

//序列化为Gbk的XML
func ToGbkXmlByte(data interface{}) ([]byte, error) {
	// req to xml
	b, err := xml.Marshal(data)
	if err != nil {
		err = errors.New("xml.Marshal error : " + err.Error())
		return b, err
	}
	// xml header
	xmlString := XmlGbkHeader + string(b)
	// utf8 conversion gbk
	return Utf8ToGbk([]byte(xmlString))
}
