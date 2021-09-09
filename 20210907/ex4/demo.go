package main

import(
    "fmt"
)

func main(){
    //%d,%s
    var stockcode = 123
    var enddate = "2021-09-07"
    var url = "Code = %d & endDate = %s"
    var target_url = fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)
}
