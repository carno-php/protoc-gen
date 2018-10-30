package php

import (
	"fmt"
	"github.com/carno-php/protoc-gen/meta"
	"hash/crc32"
)

type Context struct {
	Meta     *meta.Description
	Imported map[string]ClassName
}

func NewContext(md *meta.Description) *Context {
	return &Context{
		Meta:     md,
		Imported: make(map[string]ClassName),
	}
}

func (ctx *Context) Using(class ClassName) string {
	named := string(class.Named())
	used := ""
	if imported, exists := ctx.Imported[named]; exists && imported != class {
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
	return fmt.Sprintf("C_%08x", crc32.ChecksumIEEE([]byte(input)))
}
