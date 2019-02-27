package tools

import (
	"io/ioutil"
	"crypto/md5"
	"encoding/base64"
)

func GetPasswordMd5(password string) (passwordMd5 string, err error) {
	m := md5.New()
	m.Write([]byte(password))
	value := m.Sum(nil)
	passwordMd5 = base64.StdEncoding.EncodeToString(value)
	return passwordMd5, nil
}

func GetFileContentMd5(file string) (contentMd5 string, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	m := md5.New()
	m.Write(data)
	value := m.Sum(nil)
	contentMd5 = base64.StdEncoding.EncodeToString(value)
	return contentMd5, nil
}
