package main

import (
	"cegla/core"
	"cegla/core/atr"
	"cegla/core/tags"
	"fmt"
)

func main() {

	authPage := core.HTML{
		core.Body{
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
