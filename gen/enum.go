package gen

import (
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Enum struct {
	Meta   *meta.Description
	Name   string
	Class  php.ClassName
	Values []EValue
}

type EValue struct {
	Anno string
	Key  string
	Val  int32
}

func Enums(md *meta.Description, dss ...*descriptor.EnumDescriptorProto) {
	for _, ds := range dss {
		enum := Enum{
			Meta:  md,
			Name:  ds.GetName(),
			Class: php.Namespace(php.Package(md.File), php.Class(ds.GetName())),
		}

		for _, v := range ds.Value {
			enum.Values = append(enum.Values, EValue{
				Key: *v.Name,
				Val: *v.Number,
			})
		}

		template.Rendering(md.G, "enum.php", enum.Class, enum)
	}
}
