package matcher

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// RuleTest describe a pattern matcher
type RuleTest func(item *ScrapItem) bool

// Rule describe a whole rule object
type Rule struct {
	Name  string
	Match string
	Type  ItemType
	With  []string
	Store string
	Tests []RuleTest
}

// Reader is an interface capable of reading rules
type Reader interface {
}

// ReadRules loads rules from the given file
func ReadRules(filename string) ([]Rule, error) {
	extension := filepath.Ext(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Rule{}, err
	}

	switch extension {
	case ".yml":
		fallthrough
	case ".yaml":
		return loadFromYamlFile(content)
	default:
		return nil, fmt.Errorf("%s is not a known format", filename)
	}
}

func loadFromYamlFile(contents []byte) ([]Rule, error) {
	rules := []Rule{}

	err := yaml.Unmarshal(contents, &rules)
	if err != nil {
		return nil, fmt.Errorf("error while reading rules: %s", err)
	}

	for i := range rules {
		rules[i].ParseMatchers()
		log.Info(fmt.Sprintf("%s\n  match: %s\n  type: %s\n  store: %s",
			rules[i].Name, rules[i].Match, rules[i].Type, rules[i].Store))
	}

	return rules, nil
}
