// Package main ...
package main

import (
	"fmt"

	"github.com/GoCodeAlone/rod/lib/launcher"
	"github.com/GoCodeAlone/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
