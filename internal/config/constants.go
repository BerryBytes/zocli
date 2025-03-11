package config

import (
	"os"
)

// each organization will be saved
type Context struct {
	OrganizationID   int    `yaml:"organization_id"`
	OrganizationName string `yaml:"organization_name,omitempty"`

	DefaultProject     int  `yaml:"project,omitempty"`
	DefaultApplication int  `yaml:"application,omitempty"`
	DefaultEnvironment int  `yaml:"environment,omitempty"`
	Active             bool `yaml:"active"`
}

type Config struct {
	// ConfigFolder is the path to the configuration folder.
	ConfigFolder string
	// AuthFile is the path to the authentication file.
	AuthFile string

	Populated     ConfigPopulated
	ConfPopulated bool

	// this file sets up the default context which consists
	// of environment, project and app information
	ContextFile   string
	Contexts      []*Context `yaml:"contexts"`
	ActiveContext *Context
}

// New
//
// returns a new Config instance.
func New() *Config {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return &Config{}
	}
	config := &Config{
		ConfigFolder: userHome + "/.01cloud",
		AuthFile:     "/token.yaml",
		ContextFile:  "/contexts.yaml",
	}
	config.loadFromEnv()
	return config
}

type ConfigPopulated struct {
	Email     string
	Password  string
	Debug     bool
	Dev       bool
	AuthToken string
	WebToken  string
}

var pickingList = []string{
	"EMAIL",
	"PASSWORD",
	"DEBUG",
	"DEV",
	"WEBTOKEN", // this is the token created from settings on web
}

func (c *Config) loadFromEnv() {
	for _, pick := range pickingList {
		c.Populated.setVal(pick, os.Getenv("CLI_"+pick))
	}
	c.ConfPopulated = true
}

func (c *ConfigPopulated) setVal(picked string, val string) {
	switch picked {
	case "EMAIL":
		c.Email = val
	case "PASSWORD":
		c.Password = val
	case "WEBTOKEN":
		c.WebToken = val
	case "DEBUG":
		if val == "true" {
			c.Debug = true
		}
	case "DEV":
		if val == "true" {
			c.Dev = true
		}
	}
}
