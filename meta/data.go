package meta

import (
	"github.com/carno-php/protoc-gen/carno"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/twitchtv/protogen"
)

type Description struct {
	G      *carno.Generator
	File   *descriptor.FileDescriptorProto
	Source string
}

func Package(g *carno.Generator, d *protogen.FileDescriptor) *Description {
	return &Description{
		G:      g,
		File:   d.FileDescriptorProto,
		Source: d.GetName(),
	}
}
