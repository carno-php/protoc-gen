package gen

import (
	"encoding/hex"
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/carno-php/protoc-gen/utils"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/twitchtv/protogen"
)

type GPBMeta struct {
	CTX     *Context
	Meta    *meta.Description
	Class   php.ClassName
	Imports []string
	Lines   []string
}

func Metadata(g *carno.Generator, md *meta.Description, fd *protogen.FileDescriptor) {
	gpb := &GPBMeta{
		CTX:   NewContext(g, md),
		Meta:  md,
		Class: php.Namespace(php.Package(md.File), php.Class("Metadata"), php.Protoc(fd.GetName()).Named()),
	}

	for _, file := range fd.GetDependency() {
		gpb.Imports = append(gpb.Imports, gpb.CTX.Using(php.Protoc(file)))
	}

	simplify := &descriptor.FileDescriptorProto{
		Name:        fd.Name,
		Package:     fd.Package,
		EnumType:    fd.EnumType,
		MessageType: fd.MessageType,
	}

	sets := &descriptor.FileDescriptorSet{
		File: []*descriptor.FileDescriptorProto{simplify},
	}

	bytes, err := proto.Marshal(sets)
	if err != nil {
		utils.Fatal(err)
	}

	hexed := hex.EncodeToString(bytes)

	for i := 0; i < len(hexed); i += 60 {
		if i+60 > len(hexed) {
			gpb.Lines = append(gpb.Lines, hexed[i:])
		} else {
			gpb.Lines = append(gpb.Lines, hexed[i:i+60])
		}
	}

	template.Rendering(g, "metadata.php", gpb.Class, gpb)
}
