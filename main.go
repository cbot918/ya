package main

import (
	"fmt"
	"os"

	"github.com/cbot918/ya/src/cli"
)

func main(){
	c := cli.NewYacli()
	args := os.Args
	
	// 

	// fmt.Println(args)
	// fmt.Printf("%T\n", args)
	// fmt.Println(len(args))

	if (len(args)==1){
		fmt.Println(cli.NewHelp().Get())
	}

	if (len(args)>1){
		arg := args[1]
		switch arg {
			case "list": {
				c.List()
			}
			case "install": {
				c.Install()
			}
			case "delete": {
				c.Delete()
			}
			case "installed": {
				c.Installed()
			}
			default: {
				fmt.Println("unknown command")
			}  
	 }
	}



	// service.NewYa().Run(":3333")
		
}
