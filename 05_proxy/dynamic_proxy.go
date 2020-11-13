package proxy

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"text/template"
)

func generate(filename string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	data := proxyData{
		Package:    f.Name.Name,
		StructName: "User",
	}
	m := &proxyMethod{
		Name:       "Login",
		Params:     "username string,password string",
		ParamNames: "username,password",
		Results:    "error",
	}
	data.Methods = append(data.Methods, m)
	tpl, err := template.New("").Parse(proxyTpl)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, data)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	return string(src), nil
}

const proxyTpl = `package {{.Package}}

type {{.StructName}}Proxy struct {
	user *{{.StructName}}
}

func New{{.StructName}}Proxy(u *{{.StructName}}) *{{.StructName}}Proxy {
	return &{{.StructName}}Proxy{user: u}
}

{{ range.Methods}}
func (up *{{$.StructName}}Proxy) {{.Name}}({{.Params}}) ({{.Results}}) {
	fmt.Println("check username and password")

	if err := up.user.{{.Name}}({{.ParamNames}}); err != nil {
		fmt.Println("login failed")
		return nil
	}

	fmt.Printf("login info: %s", username)
	return nil
}
{{end}}`

type proxyData struct {
	Package    string
	StructName string
	Methods    []*proxyMethod
}

type proxyMethod struct {
	Name       string
	Params     string
	ParamNames string
	Results    string
}
