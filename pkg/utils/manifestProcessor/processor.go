package manifestprocessor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/fs"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"gopkg.in/yaml.v3"
)

type Interface interface {
	// process the standard input and convert the input to the
	// respective struct along with the struct name for mapping
	ParseFromStdIn() (interface{}, string)
	// process the file input to the respective struct along
	// with the struct name for mapping
	GetFileData(string) (interface{}, string)
	// helps for the separation of the struct which is to be used
	SegregateYaml([]byte) (interface{}, string)
	// helps for the separation of the struct which is to be used
	SegregateJson([]byte) (interface{}, string)
	// generate the manifest file from the data which is received
	// from the request response, supplied as the argument
	MakeManifest(string, interface{}) interface{}
}

type Opts struct {
	F *factory.Factory
}

// New
//
// responsible for returning the interface which contains functions
// necessary for using and processing the manifest file
func New(f *factory.Factory) Interface {
	o := Opts{f}
	return &o
}

// ParseFromStdIn
//
// is responsible for parsing the yaml or the json file from the
// standard input i.e. using pipe symbol on cli
func (o *Opts) ParseFromStdIn() (interface{}, string) {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var inputLines string

	for scanner.Scan() {
		inputLines += "\n" + scanner.Text()
	}

	if scanner.Err() != nil {
		o.F.Printer.Fatal(1, scanner.Err())
	}

	// check if the first character in the input is equal to '{'
	for _, char := range inputLines {
		if fmt.Sprint(char) == "{" {
			return o.SegregateJson([]byte(inputLines))
		} else {
			break
		}
	}

	return o.SegregateYaml([]byte(inputLines))
}

// GetFileData
//
// is responsible for opening the file which name is supplied from
// the console and calling the Segregate function to map to their own struct
func (o *Opts) GetFileData(file string) (interface{}, string) {
	data, err := fs.LoadFile(file)
	if err != nil {
		o.F.Printer.Fatal(10, err.Error())
	}
	words := strings.Split(file, ".")
	ext := words[len(words)-1]
	if ext == "yaml" || ext == "yml" {
		return o.SegregateYaml(data)
	} else if ext == "json" {
		return o.SegregateJson(data)
	}
	return nil, ""
}

// SegregateYaml
//
// helps for the separation of the struct which is to be used
func (o *Opts) SegregateYaml(data []byte) (interface{}, string) {
	var manifest map[string]string
	_ = yaml.Unmarshal(data, &manifest)
	switch strings.ToLower(manifest["kind"]) {
	case "project":
		var project models.Project
		err := yaml.Unmarshal(data, &project)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return project, "project"
	case "organization":
		var organization models.Organization
		err := yaml.Unmarshal(data, &organization)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return organization, "organization"
	case "application":
		var app models.App
		err := yaml.Unmarshal(data, &app)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return app, "application"
	case "member":
		var member models.OrganizationMember
		err := yaml.Unmarshal(data, &member)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return member, "member"
	case "environment":
		var env models.Env
		err := yaml.Unmarshal(data, &env)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return env, "environment"
	default:
		o.F.Printer.Fatal(1, "no such kind of manifest found")
	}
	return nil, ""
}

// SegregateJson
//
// is responsible for switching the data and mapping to respective struct by
// checking the kind value on the yaml
func (o *Opts) SegregateJson(data []byte) (interface{}, string) {
	var manifest map[string]string
	_ = json.Unmarshal(data, &manifest)
	switch strings.ToLower(manifest["kind"]) {
	case "member":
		var member models.OrganizationMember
		err := json.Unmarshal(data, &member)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return member, "member"
	case "project":
		var project models.Project
		err := json.Unmarshal(data, &project)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return project, "project"
	case "application":
		var app models.App
		err := json.Unmarshal(data, &app)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
		return app, "application"
	case "environment":
		return nil, "environment"
	default:
		o.F.Printer.Fatal(1, "no such kind of manifest found")
	}
	return nil, ""
}
