package validation

import (
	"strconv"

	"github.com/berrybytes/zocli/pkg/utils/factory"
)

func CheckValidID(f *factory.Factory, id string) {
	if id == "" {
		return
	}
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		f.Printer.Fatal(1, "invalid id supplied")
		return
	}
}

func GetIDOrName(data string) (id string, name string) {
	_, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return "", data
	}
	return data, ""
}
