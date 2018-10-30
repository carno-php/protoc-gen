package gen

import (
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Service struct {
	CTX        *php.Context
	Meta       *meta.Description
	Name       string
	Class      php.ClassName
	Package    string
	Contracted string
	Methods    []Method
}

type Method struct {
	Anno   string
	Name   string
	Input  string
	Output string
}

func Services(md *meta.Description, dss ...*descriptor.ServiceDescriptorProto) {
	for _, ds := range dss {
		svc := Service{
			CTX:     php.NewContext(md),
			Meta:    md,
			Name:    ds.GetName(),
			Package: md.File.GetPackage(),
		}

		for _, m := range ds.GetMethod() {
			m := Method{
				Name:   m.GetName(),
				Input:  svc.CTX.Using(php.MessageName(md.G, m.GetInputType())),
				Output: svc.CTX.Using(php.MessageName(md.G, m.GetOutputType())),
			}
			svc.Methods = append(svc.Methods, m)
		}

		svc.Class = php.Namespace(php.Package(md.File), php.Class("Contracts"), php.Class(ds.GetName()))
		template.Rendering(md.G, "interface.php", svc.Class, svc)

		svc.Class = php.Namespace(php.Package(md.File), php.Class("Clients"), php.Class(ds.GetName()))
		svc.Contracted = svc.CTX.Using(php.Namespace(php.Package(md.File), php.Class("Contracts"), php.Class(ds.GetName())))
		template.Rendering(md.G, "client.php", svc.Class, svc)
	}
}
