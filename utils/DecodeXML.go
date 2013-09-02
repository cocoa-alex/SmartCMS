package utils

import (
	"encoding/xml"
	"os"
)

const (
	Header = `<?xml version="1.0" encoding="UTF-8"?>`
)

type WebSetting struct {
	XMLName   xml.Name `xml:websetting`
	WebName   string   `xml:name`
	Url       string   `xml:url`
	KeyWords  string   `xml:keywords`
	Person    string   `xml:person`
	Email     string   `xml:Email`
	Telephone string   `xml:telephone`
	Address   string   `xml:address`
}

func CreateXML(path string, v WebSetting) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	output, err1 := xml.MarshalIndent(v, " ", "  ")
	if err != nil && os.IsExist(err) {
		file.WriteString(Header + string(output))
		defer file.Close()
	} else {
		file1, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file1.Close()
		file1.WriteString(Header + string(output))
	}
	if err1 != nil {
		panic(err1)
	}
}

func DecodeXML(path string) WebSetting {
	var v WebSetting
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := make([]byte, 1024)
	count, err := file.Read(data)
	err1 := xml.Unmarshal(data[38:count], &v)
	if err1 != nil {
		panic(err1)
	}
	return v
}
