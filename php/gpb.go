package php

import (
	"fmt"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func MDClass(f *descriptor.FileDescriptorProto) ClassName {
	return Namespace(Package(f), Class("Metadata"), Protoc(f.GetName()).Named())
}

func MDLoading(md *meta.Description) string {
	return string(MDClass(md.File)) + "::init();"
}

func GPBUtil(ctx *Context) string {
	return ctx.Using(Class("Google\\Protobuf\\Internal\\GPBUtil"))
}

func GPBType(ctx *Context, typ descriptor.FieldDescriptorProto_Type) string {
	base := ctx.Using(Class("Google\\Protobuf\\Internal\\GPBType"))
	expr := ""

	switch typ {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		expr = "DOUBLE"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		expr = "FLOAT"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		expr = "INT64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		expr = "UINT64"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		expr = "INT32"
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		expr = "FIXED64"
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		expr = "FIXED32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		expr = "BOOL"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		expr = "STRING"
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		expr = "GROUP"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		expr = "MESSAGE"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		expr = "BYTES"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		expr = "UINT32"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		expr = "ENUM"
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		expr = "SFIXED32"
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		expr = "SFIXED64"
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		expr = "SINT32"
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		expr = "SINT64"
	}

	return fmt.Sprintf("%s::%s", base, expr)
}
