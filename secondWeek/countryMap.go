package main
import "fmt"

func main(){
    //创建map
    countryCapitalMap := map[string]string{"France":"Paris",
                                            "Italy":"Rome",
                                            "Japan":"Tokyo",
                                            "India":"New delhi"}

    fmt.Println("原始地图")

    for country := range countryCapitalMap {
        fmt.Println(country,"首都  ",countryCapitalMap[country])
    }

    //删除 元素
    delete(countryCapitalMap,"France")

    fmt.Println("法国条目被删除")

    fmt.Println("删除后的地图")

    for country := range countryCapitalMap {
        fmt.Println(country,"首都 ",countryCapitalMap[country])
    }
}