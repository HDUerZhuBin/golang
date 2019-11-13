package main
import "fmt"

func main(){
	doDBOperations()
}

func connectToDB(){
	fmt.Println("ok,connected to db")
}

func disconnectFromDB(){
	fmt.Println("ok,disconnect from db")
}

func doDBOperations(){
	connectToDB()
	fmt.Println("defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("doing some DB oeration...")
	fmt.Println("oops,some crash or network error...")
	fmt.Println("returning from function here")
	return 
}