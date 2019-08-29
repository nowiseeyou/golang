## select 语句 ##

select 是Go 中的一个控制结构，类似于用于通信的switch 语句，每个case 必须是一个通信操作，要么是发送，要么是接收。

select 随机执行一个可运行的 case。 如果没有case 可运行，她将阻塞，直到有case可运行，一个默认的字句应该总是可运行的。

## 语法 ##

Go编程语句中 select 语句的语法如下：

    select {
		case communication clause :
			statement(s);
		case communication clause :
			statement(s)
		default :
			statement(s)
	}

以下描述了select 语句的语法：

- 每个case都必须是一个通信
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，她就执行，其他被忽略。
- 如果多个case都可以运行，select 会随机公平的选出一个执行。其他不会执行。
	
否则：
1. 如果有default 子句，则执行该语句。
2. 如果没有default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对channel或值进行求值。

例：

    package main
	import "fmt"

	func main(){
		var c1,c2,c3 chan int
		var i1,i2 int
		
		select {
			case i1 = <-c1:
				fmt.Printlf("received ",i1," from c1\n")
			case c2 <- i2:
				fmt.Printlf("sent ",i2,"from c2\n")
			case i3,ok := (<-c3): //same as :i3,ok := <-c3
				if ok {
					fmt.Printf("received",i3," from c3\n")
				}else{
					fmt.Printf("c3 is closed\n")
				}
			default:
				fmt.Printf("no communication\n")
		}
	}

输出结果：

    no communication