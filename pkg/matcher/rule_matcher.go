package matcher

import (
	log "github.com/sirupsen/logrus"
)

// Matcher can find a rule for a given file
type Matcher struct {
	Rules []Rule
}

// NewMatcher return a matcher from config
func NewMatcher(rulesfile string) Matcher {
	rules, err := ReadRules(rulesfile)
	if err != nil {
		log.Warn(err)
		return Matcher{Rules: []Rule{}}
	}
	return Matcher{Rules: rules}
}

// FindRuleFor return the first matching rull for a given file
func (m *Matcher) FindRuleFor(item *ScrapItem) (*Rule, error) {
	for _, rule := range m.Rules {
		if rule.MatchItem(item) {
			return &rule, nil
		}
	}
	log.Info(item.SourcePath, "does not satisfy any rules")
	return nil, nil
}
