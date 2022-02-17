package main

import (
	"go-form/routers"
)


func main() {
	r:=routers.SetupRouter()
	r.Run()
}