package main
import "fmt"

func main(){
	var x interface{}

	switch i := x.(type){
		case nil:
			fmt.Printf("x 的类型：%T",i)
		case int:
			fmt.Printf("x 是int型")
		case float64:
			fmt.Printf("x 是float64型")
		case func(int) float64:
			fmt.Printf("x是func(int)型")
		case bool,string:
			fmt.Printf("x是bool或者string型")
		default:
			fmt.Printf("未知型")
	}
}