package main

import "api_ticket/router"

func main() {
	StartApp()
}

// StartApp start app
func StartApp() {
	
	router.InitRouter().Run(":8001")
}
