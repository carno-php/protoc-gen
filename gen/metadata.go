package gen

import (
	"encoding/hex"
	"github.com/carno-php/protoc-gen/meta"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/template"
	"github.com/carno-php/protoc-gen/utils"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/twitchtv/protogen"
)

type GPBMeta struct {
	CTX     *php.Context
	Meta    *meta.Description
	Class   php.ClassName
	Imports []Imported
	Lines   []string
}

type Imported struct {
	Class string
	WKT   bool
}

func Metadata(md *meta.Description, fd *protogen.FileDescriptor) {
	gpb := &GPBMeta{
		CTX:   php.NewContext(md),
		Meta:  md,
		Class: php.MDClass(md.File),
	}

	for _, file := range fd.GetDependency() {
		if imported, exists := md.G.GetFileD(file); exists {
			gpb.Imports = append(gpb.Imports, Imported{
				Class: gpb.CTX.Using(php.MDClass(imported.FileDescriptorProto)),
				WKT:   imported.GetPackage() == php.PackageWKT,
			})
		}
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

	template.Rendering(md.G, "metadata.php", gpb.Class, gpb)
}
