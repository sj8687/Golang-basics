package main

import (
	"fmt"
	"errors"
)


func divide(a,b int)(int,error){
	if b == 0 {
		return 0, errors.New("canot");
	}

	return a / b , nil
}


func main() {
	result, err := divide(10,0);
	 if err != nil{
		fmt.Println("error",err);
	 }

	 fmt.Println(result);
}