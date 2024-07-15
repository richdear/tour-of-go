package main

import ("fmt"; "runtime")

func main(){
	fmt.Println("Go runs on ")
	switch os:=runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Printf("Linus fan %s))",os)
	default:
		fmt.Printf("%s. \n",os)
	}
}