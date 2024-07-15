package main

import ("fmt"
"math"
)

func Sqrt(x float64) float64{
	var z float64=1.0
	prev:=0.0
	for  {
		dif :=math.Abs(z-prev); 
		if dif<=0.1{
			// fmt.Printf("PREV:%g Z is %g DIFF: %f\n",prev,z,dif)
			break
		}
		prev=z
		z -= (z*z - x) / (2*z)
		// fmt.Printf("PREV:%g Z is %g DIFF: %f\n",prev,z,dif)
	}
	return z
}

func main(){
	for number:=2.0; number<=10.0;number++{
		fmt.Printf("Number: %f Sqrt: %f\n",number,Sqrt(number))
	}
	
}