//Go supports constants of character, string, boolean and numeric.
package main
import "fmt"
import "math"

const s string ="constant" // const decalres a constant value
func main()  {

	fmt.Println(s)   
	const a= 500   		 //constant as variable

	const b = 3e20 / a   //Arithmatic with binary precision
	fmt.Println(b)
	
	fmt.Println(int64(b))  //Numeric constant

	fmt.Println(math.Sin(a))	//variable asgn or function call (math.Sin) expect float64
	
}