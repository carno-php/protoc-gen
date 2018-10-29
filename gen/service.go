package gen

import (
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Service struct {
	CTX        *Context
	Meta       *meta.Description
	Name       string
	Class      php.ClassName
	Contracted string
	Methods    []Method
}

type Method struct {
	Package string
	Service string
	Anno    string
	Name    string
	Input   string
	Output  string
}

func Services(g *carno.Generator, md *meta.Description, dss ...*descriptor.ServiceDescriptorProto) {
	for _, ds := range dss {
		svc := Service{
			CTX:  NewContext(g, md),
			Meta: md,
			Name: ds.GetName(),
		}

		for _, m := range ds.GetMethod() {
			m := Method{
				Package: md.File.GetPackage(),
				Service: svc.Name,
				Name:    m.GetName(),
				Input:   svc.CTX.Using(php.MessageName(g, m.GetInputType())),
				Output:  svc.CTX.Using(php.MessageName(g, m.GetOutputType())),
			}
			svc.Methods = append(svc.Methods, m)
		}

		svc.Class = php.Namespace(php.Package(md.File), php.Class("Contracts"), php.Class(ds.GetName()))
		template.Rendering(g, "interface.php", svc.Class, svc)

		svc.Class = php.Namespace(php.Package(md.File), php.Class("Clients"), php.Class(ds.GetName()))
		template.Rendering(g, "client.php", svc.Class, svc)
	}
}
