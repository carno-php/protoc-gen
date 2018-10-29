package gen

import (
	"fmt"
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"hash/crc32"
)

type Context struct {
	Generator *carno.Generator
	Metadata  *meta.Description
	Imported  map[string]php.ClassName
}

func NewContext(g *carno.Generator, md *meta.Description) *Context {
	return &Context{
		Generator: g,
		Metadata:  md,
		Imported:  make(map[string]php.ClassName),
	}
}

func (ctx *Context) Using(class php.ClassName) string {
	named := string(class.Named())
	used := ""
	if _, exists := ctx.Imported[named]; exists {
		used = CRC32(string(class))
	} else {
		used = named
	}
	ctx.Imported[used] = class
	return used
}

func (ctx *Context) Namespaces() []string {
	uses := make([]string, 0)
	for alias, class := range ctx.Imported {
		if string(class.Named()) == alias {
			uses = append(uses, string(class))
		} else {
			uses = append(uses, string(class)+" as "+alias)
		}
	}
	return uses
}

func CRC32(input string) (output string) {
	return fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(input)))
}
