package context

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"gopkg.in/yaml.v3"
)

// Loader
//
// this function is responsible for calling the necessary methods
// which will populate the contexts on the config
func Loader(f *factory.Factory) {
	loadFromConfig(f)
	checkEnv(f)
}

// checkEnv
//
// checks if any temporary context has been set on the environment
func checkEnv(f *factory.Factory) {
	active := activeContext(f)
	if active == -1 {
		return
	}

	var pickingList = []string{
		"PROJECT",
		"APPLICATION",
		"ORGANIZATION",
	}

	for _, pick := range pickingList {
		setVal(f, pick, os.Getenv("CLI_"+pick), active)
	}
}

func setVal(f *factory.Factory, picked, value string, activeContext int) {
	if value == "" {
		return
	}

	actualVal, err := strconv.Atoi(value)
	if err != nil {
		return
	}

	switch picked {
	case "PROJECT":
		f.Config.Contexts[activeContext].DefaultProject = actualVal
	case "APPLICATION":
		f.Config.Contexts[activeContext].DefaultApplication = actualVal
	case "ORGANIZATION":
		f.Config.Contexts[activeContext].OrganizationID = actualVal
	}
}

// loadFromConfig
//
// this function will load the contexts from the configuration file
// which is present on the system's config folder for 01cloud
// If none of the configuration found, or if any error is detected then,
// this function returns nil and the checking is aborted
func loadFromConfig(f *factory.Factory) {
	_, err := os.Stat(f.Config.ConfigFolder + f.Config.ContextFile)
	if err != nil {
		return
	}

	var contexts []*config.Context
	file, err := os.ReadFile(f.Config.ConfigFolder + f.Config.ContextFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &contexts)
	if err != nil {
		f.Printer.Errorf("possible tampering of contexts file.\nErr: ", err)
		return
	}

	f.Config.Contexts = contexts
	active := activeContext(f)
	if active != -1 && active != -2 {
		f.Config.ActiveContext = f.Config.Contexts[active]
	}
}

// SaveContextChanges
//
// this function will be responsible to save any changes on the context
// to the contexts configuration file
func SaveContextChanges(f *factory.Factory, changesID int, changesOn, changesName string) {
	active := activeContext(f)
	if active == -1 {
		createNewContextFile(f)
		active = 0
	}

	switch changesOn {
	case "orgranization":
		f.Config.Contexts[active].OrganizationID = changesID
		f.Config.Contexts[active].OrganizationName = changesName
	case "project":
		f.Config.Contexts[active].DefaultProject = changesID
	case "application":
		f.Config.Contexts[active].DefaultApplication = changesID
	case "environment":
		f.Config.Contexts[active].DefaultEnvironment = changesID
	default:
		f.Printer.Fatal(10, "no such context configuration can be saved")
		return
	}
	val := true
	f.Config.Contexts[active].Active = val

	// now save the changes
	file, err := os.Create(f.Config.ConfigFolder + f.Config.ContextFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		f.Printer.Fatal(10, "cannot create the config file")
		return
	}

	data, err := yaml.Marshal(f.Config.Contexts)
	if err != nil {
		f.Printer.Fatalf(9, "cannor Marshal")
		return
	}

	if _, err := file.Write(data); err != nil {
		f.Printer.Fatal(10, "cannot write to config file")
		return
	}
}

// createNewContextFile
//
// if no any context file is detected, and the user wants to add a context,
// then this function is responsible to create a new one
func createNewContextFile(f *factory.Factory) {
	file, err := os.Create(f.Config.ConfigFolder + f.Config.ContextFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		f.Printer.Fatal(10, "cannot create the config file")
		return
	}

	if len(f.Config.Contexts) == 0 {
		f.Config.Contexts = make([]*config.Context, 1)
	}

	val := true
	f.Config.Contexts[0].Active = val

	data, err := yaml.Marshal(f.Config.Contexts)
	if err != nil {
		f.Printer.Fatalf(9, "cannot Marshal")
		return
	}

	if _, err := file.Write(data); err != nil {
		fmt.Println(err)
		f.Printer.Fatal(1, "cannot write to config file")
	}
}

// activeContext
//
// finds the active context which is being used from the contexts
// that had been previously loaded
func activeContext(f *factory.Factory) int {
	if len(f.Config.Contexts) == 0 {
		return -1
	}

	for index, context := range f.Config.Contexts {
		if context.Active {
			return index
		}
	}

	return -2
}

// SetActiveFalse
//
// this function sets all of the context available on
// the factory instance active state as false
func SetActiveFalse(f *factory.Factory) *factory.Factory {
	for _, context := range f.Config.Contexts {
		context.Active = false
	}
	return f
}

func SetFirstActiveContext(f *factory.Factory) {
	if len(f.Config.Contexts) != 0 {
		return
	}
	f.Config.Contexts = append(f.Config.Contexts, &config.Context{
		OrganizationID:   0,
		OrganizationName: "Default",
		Active:           true,
	})
	createNewContextFile(f)
}

func CheckIfOrgInContext(target int, f *factory.Factory) bool {
	for _, context := range f.Config.Contexts {
		if context.OrganizationID == target {
			return true
		}
	}
	return false
}

func SetActive(targetOrgID int, f *factory.Factory) (error, *factory.Factory) {
	for _, context := range f.Config.Contexts {
		if context.OrganizationID == targetOrgID {
			context.Active = true
			return nil, f
		}
	}
	return errors.New("no such org found"), f
}
