package main

import (
	"github.com/cbot918/ya/src/ya"
	// "github.com/cbot918/liby/cmdy"
)



func main(){
	// c := cmdy.New()
	// script1 := []string{"git clone -b go https://github.com/cbot918/template"}
	// script2 := []string{"rm -rf template"}
	// c.Run(script1)
	// c.Run(script2)

	y := ya.New()
	y.Run()

	// branch 
	// name
	// rm -rf .git && git init
	// git add .
	// git commit -m "init project"
	// gh repo create $name --public --add-readme --push --source .

}