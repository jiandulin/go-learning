package main
import "fmt"
func main(){
    var a bool = true
    var b bool = false
    if(a && b){
        fmt.Printf("first line is true\n")
    }
    if(a || b){
        fmt.Printf("second line is true\n")
    }

    a = false
    b = true
    if(a && b){
        fmt.Printf("third line is true\n")
    }else{
        fmt.Printf("third line is false\n")
    }

    if(!(a && b)){
        fmt.Printf("fourth line is true\n")
    }
}

