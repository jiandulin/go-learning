package main
import "fmt"
func main(){
	/*定于局部变量*/
	var a int = 100
	var b int = 200

	/*判断条件*/
	if a == 100{
		/*if条件语句为true执行*/
		if b == 200{
			/*if条件语句为true执行*/
			fmt.Printf("a 的值为100，b的值为200\n");
		}
	}
	fmt.Printf("a 值为：%d\n",a)
	fmt.Printf("b 值为：%d\n",b)
}