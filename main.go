package main

import (
	"flag"
	"fmt"

	"spicychicken.top/auto-mysql/installer"
)

func main() {
	fmt.Println("")
	mode := *flag.String("m", "install", "install/remove")

	if mode == "install" {
		if err := installer.Install(); err != nil {
			fmt.Println(err)
		}
	} else {
		installer.Remove()
	}

}
