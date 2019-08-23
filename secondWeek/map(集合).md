## map（集合） ##

map是一种无序的键值对集合，map最重要的一点是通过key来快速检索数据，key类似于所以，定向数据的值。

map是一种集合，所以我们可以像迭代数组和切片那样来迭代它。不过，map是无序的，我们无法决定它的返回顺序，这是因为map 使用hash表来实现的。

## 定义map ##

可以使用内建函数make 也可以使用map关键字来定义map。

    //声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type

	//使用 make 函数
	map_variable := make(make[key_data_type]value_data_type)

如果不初始化map，那么就会创建一个nil map。nil map 不能用来存放简直对。

例：

    package main
	import "fmt"
	
	func main(){
		//创建集合
		var countryCapitalMap map[string]string  
		countryCapitalMap = make(map[string]string)

		//map插入 key - value对，各个国家对应的首都
		countryCapitalMap['France'] = "巴黎"
		countryCapitalMap['Italy'] = "罗马"
		countryCapitalMap['Japan'] = "东京"
		countryCapitalMap['India'] = "新德里"

		//使用键输出地图值
		for country := range countryCapitalMap {
			fmt.Println(country,"首都是",countryCapitalMap [country])
		}

		//查看元素在集合中是否存在(确定是真实的则存在，反之 不存在)
		copital,ok := countryCapitalMap['American']
		
		if(ok) {
			fmt.Println("American 的首都是",capital)
		}else{
			fmt.Println("American 的首都不存在")
		}
	}

输出结果：

    France 首都是 巴黎
	Italy  首都是 罗马
	Japan  首都是 东京
	India  首都是 新德里
	American  的首都不存在

## delete()函数 ##

delete（）函数用于删除集合的元素，