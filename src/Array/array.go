package main
import "fmt"

func main()  {

	var a[5] int
	{
	fmt.Print("array",a)
	fmt.Print("\n")
	}
	a[4] = 100
    fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // Length of an array

	b := [5]int{1, 2, 3, 4, 5} //Array initialization single line
	fmt.Println("dcl:", b)
	

	var twoD [2][5]int   		//Multidimentional Array
    for i := 0; i < 2; i++ {
        for j := 0; j <5; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d array: ", twoD)

}