package main

import (
	"fmt"
	"io"
	"os"
)

// Contents 将文件的内容作为字符串返回
func Contents (filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		 return "", err
	}
	
	defer f.Close() // f.Close 会在我们结束之后运行
	
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF{
				break
			}
			return "", err  // 这里返回后，f 将会关闭
		}
	}
	
	return  string(result),nil // 这里返回后，f 将会关闭
}

func main()  {
	
	content,err := Contents("note.txt")
	fmt.Printf("content: %s ",content)
	panic(err)
	world := "世界"
	fmt.Printf("hello %s" ,world)
}

