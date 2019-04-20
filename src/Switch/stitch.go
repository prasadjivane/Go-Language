package main
import "fmt"
import "time"

func main()  {

	i:=2
	fmt.Print("Value Of ",i,"As")
	switch i {
	case 1:
		{
			fmt.Print("One")
		}
	case 2:
		{
			fmt.Print("Two")
		}
	case 3:
		{
			fmt.Print("Three")
		}	
		
	}
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	t:= time.Now()
	switch{
	case t.Hour() <12 :
		fmt.Print("Before Noon")
	default:
		fmt.Print("After Noon")
	}
	test := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("bool")
        case int:
            fmt.Println("int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    test(true)
    test(1)
	test("hey")
}