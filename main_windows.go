//+build go1.12

package main

import (
	"flag"
	"fmt"
	"github.com/Bios-Marcel/cordless/app"
	"github.com/Bios-Marcel/cordless/config"
	"github.com/Bios-Marcel/cordless/version"
)

func main() {
	showVersion := flag.Bool("version", false, "Show the version instead of starting cordless")
	flag.BoolVar(&config.DisableUTF8, "disable-UTF8", true, "Replaces certain characters with question marks in order to avoid broken rendering")
	flag.Parse()

	if showVersion != nil && *showVersion {
		fmt.Printf("You are running cordless version %s\nKeep in mind that this version might not be correct for manually built versions, as those can contain additional commits.\n", version.Version)
	} else {
		app.Run()
	}
}