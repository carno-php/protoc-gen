package main

import (
	"github.com/gobuffalo/packr"
	"github.com/twitchtv/protogen"
)

func main() {
	protogen.RunProtocPlugin(NewPluginG(packr.NewBox("./templates")))
}
