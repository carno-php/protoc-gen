package main

import (
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/gen"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/utils"
	"github.com/gobuffalo/packr"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twitchtv/protogen"
)

type PluginG struct {
	gen *carno.Generator
}

func NewPluginG(box packr.Box) *PluginG {
	return &PluginG{
		gen: carno.New(box),
	}
}

func (p *PluginG) Generate(in *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	gFiles, _, aNamedFiles, err := protogen.WrapTypes(in)
	if err != nil {
		utils.Fatal(err)
	}

	p.gen.Init(in)
	p.gen.NamedFiles(aNamedFiles)

	for _, fd := range gFiles {
		gen.Metadata(meta.Package(p.gen, fd), fd)
		gen.Enums(meta.Package(p.gen, fd), fd.GetEnumType()...)
		gen.Messages(meta.Package(p.gen, fd), fd.GetMessageType()...)
		gen.Services(meta.Package(p.gen, fd), fd.GetService()...)
	}

	return p.gen.Response(), nil
}
