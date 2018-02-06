package main

import (
	"log"
	"io/ioutil"
	"os"
)

func main() {
	var path = "testFile.md"
	var content string = "hello world!"
	writeFile(content, path)
	log.Println(readFile(path))

	osOpenAndWrite(path,content)

	osOpenAndRead(path)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
	write content
 */
func writeFile(content string, path string) {
	log.Println("wirte content:", content, ",file path :", path)
	err := ioutil.WriteFile(path, []byte(content), 0644)
	check(err)
}

/**
	read file content
 */
func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	check(err)
	return string(data)
}

func osOpenAndRead(path string) {
	file, err := os.Open(path)
	defer file.Close()
	check(err)
	//设置读取字节的缓存
	b := make([]byte, 20)
	//返回值n表示实际读取的字节数
	n,err:=file.Read(b)
	check(err)
	log.Println(string(b),n)
}


func osOpenAndWrite(path string,content string){
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	n1,err:=file.Write([]byte(content))
	check(err)
	log.Println(n1)

	n2,err:=file.WriteString("\n")
	check(err)
	log.Println(n2)

	n3,err:=file.WriteString(content)
	check(err)
	log.Println(n3)
}
