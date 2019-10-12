package main

import (
	"fmt"
	"io/ioutil"
)
// go里有相对路径是执行命令时的目录，不是相对可执行文件的
func main(){
	const filename  = "./courseNotes/abc.txt"
	if contents,err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
}
