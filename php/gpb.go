package php

import (
	"fmt"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const PackageWKT = "google.protobuf"

func MDClass(f *descriptor.FileDescriptorProto) ClassName {
	pkg := Package(f)
	if f.GetPackage() == PackageWKT {
		class := ClassName(GPBFilter(pkg, string(Protoc(f.GetName()).Named())))
		return Namespace(fmt.Sprintf("GPBMetadata.%s", pkg), class)
	}
	return Namespace(pkg, Class("Metadata"), Protoc(f.GetName()).Named())
}

func MDLoading(md *meta.Description) string {
	return fmt.Sprintf("\\%s::init();", MDClass(md.File))
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

func GPBFilter(pkg, class string) string {
	if pkg == PackageWKT {
		if class == "Empty" {
			class = "GPBEmpty"
		}
	}
	return class
}
