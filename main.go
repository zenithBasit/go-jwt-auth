package main

import (
	"fmt"

	"github.com/zenithBasit/jwt-authentication/intializers"
)

func init() {
	intializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello 2")
}
