package main
import "fmt"

func main()  {
	var no int
	fmt.Print("Enter The No To Check\n")
	fmt.Scanln(&no)
	if(no%2==0){
		fmt.Print("Even Number\n")
	}else {
		fmt.Print("Odd number\n")
	}
}