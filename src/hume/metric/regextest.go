package metric

import (
	"hume/record"
	"regexp"
)

type RegexTest struct {
	BaseMetric
	Counter
	Field string `json:"field"`
	Regex string `json:"regex"`
	r     *regexp.Regexp
}

func (rt *RegexTest) Init() error {
	r, err := regexp.Compile(rt.Regex)
	rt.r = r
	if err != nil {
		return err
	}
	return rt.Counter.Initialize("true", "false")
}

func (rt *RegexTest) Process(rec *record.Record) {
	v, ok := rec.Map[rt.Field]
	l := "false"
	if ok {
		if rt.r.MatchString(v) {
			l = "true"
		}
	}
	rt.Count(l)
}
