package main
import "fmt"

func main()  {

	i:=1                  //Single Condition For Loop
	for i<=10{
		fmt.Println(i)
		i=i+1
	}

	for a:=11;a<=100;a++{  //ICI For Loop(Normal)
		fmt.Println(a)
	}

	for{

		fmt.Println("HIIIII")
		break
	}
	
	for b:=2;b<=20;b++{
		if b%2==0{
			fmt.Println(b)
			
		}
		continue
	}
}