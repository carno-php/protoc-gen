package carno

import (
	"github.com/gobuffalo/packr"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twitchtv/protogen/typemap"
)

type Generator struct {
	box     packr.Box
	ctx     *Context
	outputs []*plugin.CodeGeneratorResponse_File
}

type Context struct {
	registry *typemap.Registry
}

func New(box packr.Box) *Generator {
	return &Generator{
		box: box,
	}
}

func (g *Generator) Init(in *plugin.CodeGeneratorRequest) {
	g.ctx = &Context{
		registry: typemap.New(in.ProtoFile),
	}
}

func (g *Generator) Box() packr.Box {
	return g.box
}

func (g *Generator) CTX() *Context {
	return g.ctx
}

func (g *Generator) Output(file string, content string) {
	g.outputs = append(g.outputs, &plugin.CodeGeneratorResponse_File{
		Name:    &file,
		Content: &content,
	})
}

func (g *Generator) Response() *plugin.CodeGeneratorResponse {
	return &plugin.CodeGeneratorResponse{
		File: g.outputs,
	}
}

func (c *Context) Registry() *typemap.Registry {
	return c.registry
}
