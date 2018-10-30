package gen

import (
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/carno-php/protoc-gen/utils"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Message struct {
	CTX     *php.Context
	Meta    *meta.Description
	Name    string
	Class   php.ClassName
	Fields  []MField
	GPBUtil string
}

type MField struct {
	Anno     string
	Name     string
	Type     string
	Default  string
	Repeated bool
	Mapped   *MMapped
}

type MMapped struct {
	Key string
	Val string
}

func Messages(md *meta.Description, dss ...*descriptor.DescriptorProto) {
	for _, ds := range dss {
		msg := Message{
			CTX:   php.NewContext(md),
			Meta:  md,
			Name:  ds.GetName(),
			Class: php.Namespace(php.Package(md.File), php.Class(ds.GetName())),
		}

		msg.GPBUtil = php.GPBUtil(msg.CTX)

		for _, f := range ds.GetField() {
			typed, defaults, _ := TypeExplains(msg.CTX, f)
			mf := MField{
				Name:     f.GetName(),
				Type:     typed,
				Default:  defaults,
				Repeated: TypeIsRepeated(f),
			}
			if typed == "array" {
				TypeMapped(msg.CTX, &mf, f)
			}
			msg.Fields = append(msg.Fields, mf)
		}

		template.Rendering(md.G, "message.php", msg.Class, msg)
	}
}

func TypeExplains(ctx *php.Context, fd *descriptor.FieldDescriptorProto) (typed, defaults string, comments []string) {
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
		msg := ctx.Meta.G.Message(fd.GetTypeName())
		if msg.Descriptor.GetOptions().GetMapEntry() {
			typed, defaults = "array", "[]"
		} else {
			typed, defaults = MessageImported(ctx, fd), "null"
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		comments = append(comments, "@see "+MessageImported(ctx, fd))
		typed, defaults = "int", "0"
	default:
		utils.Error("unknown type for", fd.GetName())
	}
	return
}

func TypeMapped(ctx *php.Context, mf *MField, f *descriptor.FieldDescriptorProto) {
	mf.Repeated = false
	msg := ctx.Meta.G.Message(f.GetTypeName())

	key := ""
	val := ""

	for _, ff := range msg.Descriptor.GetField() {
		if ff.GetName() == "key" {
			key = php.GPBType(ctx, ff.GetType())
		} else if ff.GetName() == "value" {
			val = php.GPBType(ctx, ff.GetType())
		}
	}

	mf.Mapped = &MMapped{
		Key: key,
		Val: val,
	}
}

func TypeIsRepeated(fd *descriptor.FieldDescriptorProto) bool {
	return fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

func MessageImported(ctx *php.Context, fd *descriptor.FieldDescriptorProto) string {
	return ctx.Using(php.Protoc(fd.GetTypeName()))
}
