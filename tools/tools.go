package essential

import (
	"io/ioutil"
	"crypto/md5"
	"encoding/base64"
)


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
