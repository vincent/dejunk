package matcher

import (
	"regexp"
	"strings"
)

var reExpressionList = regexp.MustCompile(`(\w*)*\(([^)]*)\)`)

type pattern struct {
	Function string
	Params   []string
}

// ParseMatchers parse the given rule
func (rule *Rule) ParseMatchers() error {
	res := reExpressionList.FindAllStringSubmatch(rule.Match, -1)

	for i := range res {
		p := pattern{
			Function: res[i][1],
			Params:   strings.Split(res[i][2], ","),
		}
		rule.Tests = append(rule.Tests, testFromExpression(&p))
	}

	return nil
}

// MatchItem returns true if all rule's tests against the given item are truthy
func (rule *Rule) MatchItem(item *ScrapItem) bool {
	for _, test := range rule.Tests {
		if !test(item) {
			return false
		}
	}
	return true
}

func testFromExpression(p *pattern) RuleTest {
	switch p.Function {
	case "ext":
		params := replaceREPlaceholders(p.Params, true)
		re := regexp.MustCompile("(" + strings.Join(params, "|") + ")$")
		return func(item *ScrapItem) bool {
			return re.Match([]byte(item.SourcePath))
		}
	case "is":
		params := replaceREPlaceholders(p.Params, true)
		re := regexp.MustCompile("(" + strings.Join(params, "|") + ")")
		return func(item *ScrapItem) bool {
			return re.Match([]byte(item.SourcePath))
		}
	case "not":
		params := replaceREPlaceholders(p.Params, true)
		re := regexp.MustCompile("(" + strings.Join(params, "|") + ")")
		return func(item *ScrapItem) bool {
			return !re.Match([]byte(item.SourcePath))
		}
	default:
		return func(item *ScrapItem) bool {
			return true
		}
	}
}

func replaceREPlaceholders(terms []string, removeOthers bool) []string {
	newTerms := []string{}
	for _, term := range terms {
		switch term {
		case ":episode":
			newTerms = append(newTerms, `([Ss]?)(\d{1,2})([xXeE\.\-])(\d{1,2})`)
		case ":audio":
			newTerms = append(newTerms, "("+strings.Join(AudioFileExts, "|")+")")
		case ":video":
			newTerms = append(newTerms, "("+strings.Join(VideoFileExts, "|")+")")
		default:
			if !removeOthers {
				newTerms = append(newTerms, term)
			}
		}
	}
	return newTerms
}
