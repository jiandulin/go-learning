package main
import "fmt"

func main(){
    var a int = 21
    var b int = 10
    var c int

    c = a + b
    fmt.Printf("first line - c value is %d\n",c)
    c = a - b
    fmt.Printf("second line - c value is %d\n",c)
    c = a * b
    fmt.Printf("third line - c value is %d\n",c)
    c = a / b
    fmt.Printf("fourth line - c value is %d\n",c)
    c = a % b
    fmt.Printf("fifth line - c value is %d\n",c)
    a++
    fmt.Printf("sixth line - a value is %d\n",a)
    a = 21
    a--
    fmt.Printf("seventh line - a value is %d\n",a)
}
