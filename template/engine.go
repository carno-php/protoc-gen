package template

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/carno-php/protoc-gen/carno"
	"github.com/carno-php/protoc-gen/php"
	"github.com/carno-php/protoc-gen/utils"
	"strings"
	"text/template"
)

func Rendering(g *carno.Generator, template string, class php.ClassName, data interface{}) {
	if rendered, err := Execute(g, template, data); err != nil {
		utils.Fatal(err)
	} else {
		g.Output(string(class.Path("php")), rendered)
	}
}

func Execute(g *carno.Generator, file string, data interface{}) (string, error) {
	tData, err := g.Box().FindString(file)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	tObj, err := template.New("").Funcs(funcMaps()).Parse(tData)
	if err != nil {
		return "", err
	}

	err = tObj.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func funcMaps() template.FuncMap {
	funcMap := sprig.TxtFuncMap()
	funcMap["Titled"] = strings.Title
	funcMap["MDInit"] = php.MDLoading
	return funcMap
}
