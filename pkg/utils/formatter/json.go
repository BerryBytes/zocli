package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/berrybytes/zocli/pkg/utils/factory"
)

func PrintJson(f *factory.Factory, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		f.Printer.Fatal(9, "JSON Marshal error: ", err)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		f.Printer.Fatal(9, "JSON parse error: ", err)
		return
	}

	f.Printer.Print(prettyJSON.String())
}
