package main

import (
	"LinkCutter/internal/profile"
	"fmt"
)

func main() {
	user := profile.User{"hello"}
	fmt.Print(user.Name)
}
