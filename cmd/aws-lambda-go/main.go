package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type config struct {
	Entries []configEntry `hcl:"event,block"`
}

type configEntry struct {
	Struct string   `hcl:"struct"`
	Output []string `hcl:"output,optional"`
}

func main() {
	var c config
	if err := hclsimple.DecodeFile("generate.hcl", nil, &c); err != nil {
		panic(err)
	}

	tmpl := parseTemplate()
	buf := bytes.NewBuffer([]byte{})
	for _, ce := range c.Entries {
		buf.Reset()
		generate(tmpl, buf, ce)
		payload, err := ioutil.ReadAll(buf)
		if err != nil {
			panic(err)
		}
		filename := fmt.Sprintf("../../lambda/handler_%s.go", toSnakeCase(ce.Struct))
		if err := ioutil.WriteFile(filename, payload, 0644); err != nil {
			log.Fatal(err)
		}
	}

}

func parseTemplate() *template.Template {
	tmpl, err := template.New("base.tmpl").Funcs(template.FuncMap{
		"generateParameters":  generateParameters,
		"generateHandlerCall": generateHandlerCall,
	}).ParseFiles(
		"./template/base.tmpl",
		"./template/function_no_return.tmpl",
		"./template/function_single_return.tmpl",
		"./template/function_multiple_return.tmpl",
	)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func generate(tmpl *template.Template, buf *bytes.Buffer, ce configEntry) {
	if err := tmpl.Execute(buf, ce); err != nil {
		panic(err)
	}
}

func generateParameters(structName string, output []string) string {
	var outputFragment string
	switch x := len(output); {
	case x == 1:
		outputFragment = output[0]
	case x > 1:
		outputFragment = fmt.Sprintf("(%s)", strings.Join(output, ", "))
	}

	input := []string{"context.Context", "events." + structName}
	result := fmt.Sprintf("handler func(%s) %s", strings.Join(input, ", "), outputFragment)
	result = strings.TrimSpace(result)
	return result
}

// ao invés de fazer isso, usar templates, muito mais fácil.
func generateHandlerCall(output []string) string {
	if len(output) == 1 && output[0] == "error" {
		return "return handler(ctx, e), nil"
	}

	var sb strings.Builder
	sb.WriteString("response, err := handler(ctx, e)\n")
	sb.WriteString("\n")
	sb.WriteString("if trace.ResponseEvent != nil {\n")
	sb.WriteString("	trace.ResponseEvent(ctx, response)\n")
	sb.WriteString("}\n")
	sb.WriteString("\n")
	sb.WriteString("return response, err\n")
	return sb.String()
}

func toSnakeCase(str string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
