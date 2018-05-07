package main

import (
	"./system/app"
	//"flag"
)

//var port string

//func init(){
//	flag.StringVar(&port, "port","9000","Assign port")
//	flag.Parse()
//}

func main() {
	server := app.NewServer()
	//server.Start(port)
	server.Start()
}
