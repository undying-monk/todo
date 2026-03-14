package main

import "fmt"

func main(){
	defer func ()  {
		fmt.Println("1")	
	}()

	defer func ()  {
		fmt.Println("2")	
	}()
}