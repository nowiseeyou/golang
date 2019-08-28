## fallthrough ##

使用fallthrough 会强制执行后面的case语句，fallthrough不会判断下一条 case 的表达式是否为true。

实例：

    package main

	import "fmt"

	func main(){
		switch {
			case false:
				fmt.Println("1.case 条件语句为 false")
				fallthrough
			case true:
				fmt.Println("2.case 条件语句为 true")
				fallthrough
			case false:
				fmt.Println("3.case 条件语句为 false")
				fallthrough
			case true:
				fmt.Println("4.case 条件语句为 true")
			case false:
				fmt.Println("5.case 条件语句为 false")
				fallthrough
			defalut:
				fmt.Println("6.默认 case")
		}
	}

输出结果：

		2.case 条件语句为 true
    	3.case 条件语句为 false
    	4.case 条件语句为 true  

结论：switch从第一个判断表达式为true 的case开始执行，如果case带有 fallthrough,程序会继续执行下一条case,且它不会去判断下一个case的表达式是否为true。