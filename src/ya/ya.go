package ya

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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


	return y
}
func (y *Ya) Run(){
	c := cmdy.New()
	args := os.Args
	// projectName := ""

	if len(args) == 1{ fmt.Println("ya help page")} else {
		switch args[1] {
			case "cmd":{
				if len(args) == 2 { fmt.Println("ya cmd help page") } else {
					u.Logg("in cmd")
					c.Run([]string{
						args[2],
					})
				}
			}
			case "install-gh": {
				if len(args) == 2 {

						fmt.Println("pre-install start...")
						// get ls -a string and identify contains .dockerenv or not
						execStr := "ls / -a"
						result, err:=exec.Command("/bin/sh", "-c", execStr).Output()
						u.Checke(err, "exec command failed")
						c.Run([]string{
							"git config --global user.name 'cbot918'",
							"git config --global user.email 'cbot918@gmail.com'",
						})
						if (strings.Contains(string(result), ".dockerenv")) {
							// docker env
							c.Run([]string{
								"apt update && apt install curl -y",
								fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",y.GhUrl, y.GhName),
								fmt.Sprintf("mkdir -p /usr/local/bin && cp %s/bin/gh /usr/local/bin ",y.GhName),
								fmt.Sprintf("rm -r %s && rm %s.tar.gz",y.GhName,y.GhName ),
							})
						} else {
							//VM env
							c.Run([]string{
								"sudo apt update && sudo apt install curl",
								fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",y.GhUrl, y.GhName),
								fmt.Sprintf("mkdir -p /usr/local/bin && sudo cp %s/bin/gh /usr/local/bin ",y.GhName),
								fmt.Sprintf("rm -rf %s && rm %s.tar.gz",y.GhName,y.GhName),
							})
						}
				} 		
			}
			}
	}
}