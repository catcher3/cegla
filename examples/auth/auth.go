package main

import (
	"fmt"

	"github.com/catcher3/cegla"
	"github.com/catcher3/cegla/atr"
	"github.com/catcher3/cegla/tags"
)

func main() {

	authPage := cegla.HTML{
		cegla.Body{
			tags.Form{
				tags.Input{
					atr.TypeAttr("text"),
					atr.NameAttr("username"),
				},
			},
		},
	}

	fmt.Println(authPage)
}
