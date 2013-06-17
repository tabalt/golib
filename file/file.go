package file

import (
	"log"
	"os"
)

//读取文件内容
func ReadFile(fileName string) (content string) {
	content = string(GetFileByte(fileName))
	return
}

//读取文件内容
func GetFileByte(fileName string) []byte {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
	}
	return buf
}

//写文件
func WriteFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}

//删除文件 文件目录操作 etc.
