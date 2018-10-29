package php

import (
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/utils"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var reserved = []string{
	"abstract", "and", "array", "as", "break",
	"callable", "case", "catch", "class", "clone",
	"const", "continue", "declare", "default", "die",
	"do", "echo", "else", "elseif", "empty",
	"enddeclare", "endfor", "endforeach", "endif", "endswitch",
	"endwhile", "eval", "exit", "extends", "final",
	"for", "foreach", "function", "global", "goto",
	"if", "implements", "include", "include_once", "instanceof",
	"insteadof", "interface", "isset", "list", "namespace",
	"new", "or", "print", "private", "protected",
	"public", "require", "require_once", "return", "static",
	"switch", "throw", "trait", "try", "unset",
	"use", "var", "while", "xor", "int",
	"float", "bool", "string", "true", "false",
	"null", "void", "iterable",
}

type ClassName string

type PathName string

func (cn ClassName) Path(suffix string) PathName {
	return PathName(strings.Replace(string(cn), "\\", "/", -1) + "." + suffix)
}

func (cn ClassName) Namespaced() ClassName {
	sps := strings.Split(string(cn), "\\")
	return ClassName(strings.Join(sps[:len(sps)-1], "\\"))
}

func (cn ClassName) Named() ClassName {
	sps := strings.Split(string(cn), "\\")
	return ClassName(sps[len(sps)-1])
}

func (pn PathName) Class() ClassName {
	return ClassName(strings.Replace(string(pn), "/", "\\", -1))
}

func Protoc(name string) ClassName {
	return Namespace(strings.TrimSuffix(name, ".proto"))
}

func Class(name string) ClassName {
	isReserved := false
	for _, match := range reserved {
		if match == strings.ToLower(name) {
			isReserved = true
			break
		}
	}

	if isReserved {
		utils.Error("class name is reserved:", name)
	}

	return ClassName(name)
}

func Package(file *descriptor.FileDescriptorProto) string {
	options := file.GetOptions()

	if options != nil && options.PhpNamespace != nil {
		return options.GetPhpNamespace()
	}

	return file.GetPackage()
}

func Namespace(pkg string, more ...ClassName) ClassName {
	parts := make([]string, 0)

	fn := func(r rune) bool {
		return r == '.' || r == '/'
	}

	for _, value := range strings.FieldsFunc(pkg, fn) {
		parts = append(parts, strings.Title(value))
	}

	for _, class := range more {
		parts = append(parts, strings.Title(string(class)))
	}

	return ClassName(strings.Join(parts, "\\"))
}

func MessageName(g *carno.Generator, t string) ClassName {
	msg := g.CTX().Registry().MessageDefinition(t)

	if msg == nil {
		utils.Error("message definition not found", t)
	}

	className := ""

	if lineage := msg.Lineage(); len(lineage) > 0 {
		for _, parent := range lineage {
			className += parent.Descriptor.GetName() + "_"
		}
	}

	className += msg.Descriptor.GetName()

	return Namespace(Package(msg.File), ClassName(className))
}
