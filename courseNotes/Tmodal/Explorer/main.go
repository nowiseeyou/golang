package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Fibonacci() intGen {
	a,b := 0,1
	return func() int {
		a,b = b, a+b
		return a
	}
}

type intGen func() int

func tryDefer()  {
	for i := 0; i<100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func tryDefer2(){
	for i := 0; i<100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string){
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := Fibonacci()

	for i := 0; i < 20 ; i++ {
		fmt.Fprintln(writer,f())
	}
}


// 错误处理概念
func writeFile2(filename string){
	file,err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	err = errors.New("this is a custom error")

	if err != nil {
		if pathError,ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s ,%s, %s \n",
					pathError.Op,
					pathError.Path,
					pathError.Err,
				)
			return
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := Fibonacci()

		for i := 0; i < 20; i++ {
			fmt.Fprintln(writer,f())
		}
	}
}



func main()  {
	writeFile("text.txt")

	return
	defer fmt.Print(1)
	defer fmt.Print(2)

	fmt.Print(3)

	return
	panic("error")
	fmt.Print(4)
}
