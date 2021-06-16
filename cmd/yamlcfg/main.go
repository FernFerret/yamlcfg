package main

import (
	"fmt"
	"os"

	"github.com/creasty/defaults"
	"github.com/sanity-io/litter"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Columns []Column   `yaml:"columns"`
	Params  Parameters `yaml:"params"`
}

// Without the parent item
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := defaults.Set(c)
	if err != nil {
		return err
	}
	type alias Config
	if err := unmarshal((*alias)(c)); err != nil {
		return err
	}
	return nil
}

// Since Parameters is a direct descendent of Config, it will be drilled down
// into via the initial defaults.Set(c) call above, so we don't need to specify
// a default function for it unless it's going to be created on its own.
type Parameters struct {
	Foo    string  `yaml:"foo"`
	Bar    int     `yaml:"bar" default:"30"`
	Baz    bool    `yaml:"baz" default:"true"`
	Fooooo float64 `yaml:"fooooo" default:"43.6"`
}

// However, I've included it anyway, it doesn't hurt to have it and this way
// we'll always get the defaults no matter how we unmarshal it.
func (p *Parameters) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := defaults.Set(p)
	if err != nil {
		return err
	}
	type alias Parameters
	if err := unmarshal((*alias)(p)); err != nil {
		return err
	}
	return nil
}

// Column is created dynamically. IE: just by looking at the Config struct
// we don't know how many Columns or what their content is going to be
type Column struct {
	Key     string `yaml:"key"`
	Width   int    `yaml:"width" default:"30"`
	Visible bool   `yaml:"visible" default:"true"`
	Fmt     string `yaml:"fmt" default:"%s"`
}

func (c *Column) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := defaults.Set(c)
	if err != nil {
		return err
	}
	type alias Column
	if err := unmarshal((*alias)(c)); err != nil {
		return err
	}
	return nil
}

func main() {
	config := &Config{}
	data, err := os.ReadFile("base.yaml")
	if err != nil {
		fmt.Printf("ERROR: failed to read file: %s\n", err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Printf("ERROR: failed to read file: %s\n", err)
	}
	litter.Dump(config)
}
