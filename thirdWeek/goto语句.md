## goto语句 ##

Go语言的goto语句可以无条件的转移到过程中指定的行。

goto语句通常与条件语句配合使用。可以来实现条件转移，构成循环，跳出循环体等功能。

但是，在结构化程序中一般不主张使用goto语句，以免造成程序流程混乱，使理解和调试都产生困难。

## 语法 ##

goto语法格式如下：

    goto label;
	..
	.
	label:statement;

## 实例 ##

	package main

	import "fmt"
	
	func main(){
		//定义局部变量
		var a int = 10

		//循环
		LOOP:for a< 20 {
			if a ==15 {
				//跳过迭代
				a = a + 1
				goto LOOP
			}
			fmt.Printf("a的值为： %d\n",a)
			a++
		}
	}

输出结果：

    a的值为： 10
	a的值为： 11
	a的值为： 12
	a的值为： 13
	a的值为： 14
	a的值为： 15
	a的值为： 16
	a的值为： 17
	a的值为： 18
	a的值为： 19

