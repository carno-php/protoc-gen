package gen

import (
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/carno-php/protoc-gen/utils"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Message struct {
	CTX    *Context
	Meta   *meta.Description
	Name   string
	Class  php.ClassName
	Fields []MField
}

type MField struct {
	Anno     string
	Name     string
	Type     string
	Default  string
	Repeated bool
}

func Messages(g *carno.Generator, md *meta.Description, dss ...*descriptor.DescriptorProto) {
	for _, ds := range dss {
		msg := Message{
			CTX:   NewContext(g, md),
			Meta:  md,
			Name:  ds.GetName(),
			Class: php.Namespace(php.Package(md.File), php.Class(ds.GetName())),
		}

		for _, f := range ds.GetField() {
			typed, defaults, _ := TypeExplains(msg.CTX, f)
			mf := MField{
				Name:     f.GetName(),
				Type:     typed,
				Default:  defaults,
				Repeated: TypeIsRepeated(f),
			}
			msg.Fields = append(msg.Fields, mf)
		}

		template.Rendering(g, "message.php", msg.Class, msg)
	}
}

func TypeExplains(ctx *Context, fd *descriptor.FieldDescriptorProto) (typed, defaults string, comments []string) {
	switch fd.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		typed, defaults = "float", "0.0"
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		typed, defaults = "int", "0"
	case descriptor.FieldDescriptorProto_TYPE_STRING,
		descriptor.FieldDescriptorProto_TYPE_BYTES:
		typed, defaults = "string", "\"\""
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		typed, defaults = "bool", "false"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		typed, defaults = MessageImported(ctx, fd), "null"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		comments = append(comments, "@see "+MessageImported(ctx, fd))
		typed, defaults = "int", "0"
	default:
		utils.Error("unknown type for", fd.GetName())
	}
	return
}

func TypeIsRepeated(fd *descriptor.FieldDescriptorProto) bool {
	return fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

func MessageImported(ctx *Context, fd *descriptor.FieldDescriptorProto) string {
	return ctx.Using(php.Namespace(php.Package(ctx.Metadata.File), php.ClassName(fd.GetName())))
}
