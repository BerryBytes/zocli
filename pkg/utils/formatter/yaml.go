package formatter

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"gopkg.in/yaml.v3"
)

func PrintYaml(f *factory.Factory, data interface{}) {
	var b bytes.Buffer

	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	err := yamlEncoder.Encode(&data)
	if err != nil {
		f.Printer.Fatal(9, "YAML Marshal error: ", err)
		return
	}

	f.Printer.Print(b.String())
}
