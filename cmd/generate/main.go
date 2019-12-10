package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type config struct {
	Handlers []configHandler `hcl:"handler,block"`
}

type configHandler struct {
	Struct string   `hcl:"struct"`
	Output []string `hcl:"output,optional"`
}

func main() {
	var c config
	if err := hclsimple.DecodeFile("generate.hcl", nil, &c); err != nil {
		fmt.Printf("failed to decode the configuration file 'generate.hcl': %s\n", err.Error())
		os.Exit(1)
	}

	tmpl, err := parseTemplate()
	if err != nil {
		fmt.Printf("failed to generate the template: %s\n", err.Error())
		os.Exit(1)
	}

	buf := bytes.NewBuffer([]byte{})
	for _, handler := range c.Handlers {
		buf.Reset()
		if err := tmpl.Execute(buf, handler); err != nil {
			fmt.Printf("failed to execute the template: %s\n", err.Error())
			os.Exit(1)
		}

		payload, err := ioutil.ReadAll(buf)
		if err != nil {
			fmt.Printf("failed to read the generated handler: %s\n", err.Error())
			os.Exit(1)
		}

		filename := fmt.Sprintf("../../lambda/handler_%s.go", toSnakeCase(handler.Struct))
		if err := ioutil.WriteFile(filename, payload, 0644); err != nil {
			fmt.Printf("failed to write the generated handler: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

func parseTemplate() (*template.Template, error) {
	tmpl, err := template.New("base.tmpl").Funcs(template.FuncMap{
		"generateParameters": generateParameters,
	}).ParseFiles(
		"./template/base.tmpl",
		"./template/function_no_return.tmpl",
		"./template/function_single_return.tmpl",
		"./template/function_multiple_return.tmpl",
	)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
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
	return strings.TrimSpace(result)
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
