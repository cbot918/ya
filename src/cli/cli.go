/*
list
search
install
delete
installed
*/

package cli

import (
	u "github.com/cbot918/liby/util"
)

type Yacli struct{}

func NewYacli() *Yacli {
	y := new(Yacli)
	return y
}

func (y *Yacli ) List (){
	u.Logg("list package")
}

func (y *Yacli ) Install (){
	u.Logg("install package")
}

func (y *Yacli ) Delete (){
	u.Logg("delete package")
}

func (y *Yacli ) Installed (){
	u.Logg("list installed")
}
