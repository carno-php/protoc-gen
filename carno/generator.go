package carno

import (
	"github.com/carno-php/protoc-gen/utils"
	"github.com/gobuffalo/packr"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twitchtv/protogen"
	"github.com/twitchtv/protogen/typemap"
)

type Generator struct {
	box     packr.Box
	ctx     *Context
	files   map[string]*protogen.FileDescriptor
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

func (g *Generator) NamedFiles(list map[string]*protogen.FileDescriptor) {
	g.files = list
}

func (g *Generator) GetFileD(name string) (found *protogen.FileDescriptor, exists bool) {
	found, exists = g.files[name]
	return
}

func (g *Generator) Box() packr.Box {
	return g.box
}

func (g *Generator) CTX() *Context {
	return g.ctx
}

func (g *Generator) Message(typed string) *typemap.MessageDefinition {
	msg := g.CTX().Registry().MessageDefinition(typed)

	if msg == nil {
		utils.Error("message definition not found", typed)
	}

	return msg
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
