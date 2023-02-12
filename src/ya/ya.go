package ya

import (
	"fmt"
	"os"

	"github.com/cbot918/liby/cmdy"
	u "github.com/cbot918/liby/util"
)


type Ya struct {
	GhName string
	GhUrl string
}

func New() *Ya{
	y := new(Ya)

	y.GhName = "gh_2.23.0_linux_amd64"
	y.GhUrl = "https://github.com/cli/cli/releases/download/v2.23.0/gh_2.23.0_linux_amd64.tar.gz"

	fmt.Println("hello this is ya ( yale-agent )")

	return y
}
func (y *Ya) Run(){
	c := cmdy.New()
	args := os.Args
	projectName := ""
	switch args[1] {
	case "cmd":{
		u.Logg("in cmd")
		c.Run([]string{
			args[2],
		})
	}
	case "install": {
		if args[2] == "docker"{
			c.Run([]string{
				"apt update && apt install curl -y",
				fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",y.GhUrl, y.GhName),
				fmt.Sprintf("mkdir -p /usr/local/bin && cp %s/bin/gh /usr/local/bin ",y.GhName),
				fmt.Sprintf("rm -r %s && rm %s.tar.gz",y.GhName,y.GhName ),
				"git config --global user.name 'cbot918'",
				"git config --global user.email 'cbot918@gmail.com'",
			})
		}
		if args[2] == "vm" {
			c.Run([]string{
				"sudo apt update && sudo apt install curl",
				fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",y.GhUrl, y.GhName),
				fmt.Sprintf("mkdir -p /usr/local/bin && sudo cp %s/bin/gh /usr/local/bin ",y.GhName),
				fmt.Sprintf("rm -rf %s && rm %s.tar.gz",y.GhName,y.GhName),
				"git config --global user.name 'cbot918'",
				"git config --global user.email 'cbot918@gmail.com'",
			})
		}

	}

	case "init":{
		if args[2] == "."{
			u.Logg("in .")
		}else {
			projectName = args[2]
			c.Run([]string{"git clone -b go https://github.com/cbot918/template"})
			c.Run([]string{fmt.Sprintf("mv template %s", projectName)})
		}
	}
	case "gitc":{
		u.Logg("in gitc")
		if args[2] == "."{
			u.Logg("in gitc .")
			c.Run([]string{
				"rm -rf .git && git init",
				"git checkout -b main",
				// fmt.Sprintf("echo \"# %s\" > README.md", projectName ),
				"git add .",
				"git commit -m 'init project'",
				"echo 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 | gh auth login --with-token",
				"gh repo create --public --push --source .",
			})
		}else{
			fmt.Println("後面+個點")
			return
		}
	
	}
	}
}