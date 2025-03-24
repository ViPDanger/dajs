package logger

import "fmt"


func Printf(format string,v ...any) (int, error){
	return fmt.Printf(format,v...)
}

func Println(v ...any) (int, error){
	return fmt.Println(v...)
}

func Print(v ...any) (int, error){
	return fmt.Print(v...)
}
