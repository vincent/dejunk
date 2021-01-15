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

func (rule Rule) isValid() bool {
	return rule.Match != "" && rule.Name != "" && rule.Store != "" && rule.Type != ""
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
	rules, returned := []Rule{}, []Rule{}

	err := yaml.Unmarshal(contents, &rules)
	if err != nil {
		return nil, fmt.Errorf("error while reading rules: %s", err)
	}

	for i, rule := range rules {
		if rule.isValid() {
			err = rules[i].ParseMatchers()
			if err == nil {
				returned = append(returned, rules[i])
				log.Info(fmt.Sprintf("%s\n  match: %s\n  type: %s\n  store: %s",
					rules[i].Name, rules[i].Match, rules[i].Type, rules[i].Store))
			} else {
				log.Warn(fmt.Sprintf("failed to parse matchers of rule %d:%s", i, rule.Name))
			}
		} else {
			log.Warn(fmt.Sprintf("rule %d is invalid", i))
		}
	}

	if len(returned) == 0 {
		return nil, fmt.Errorf("no rules found")
	}

	return returned, nil
}
