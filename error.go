package date

import "errors"

var (
	ErrNotTime error = errors.New("not a Time")
)

// ParseError is the error that happens when parsing the time string by the layout.
type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
}

func (pe *ParseError) Error() string {
	return `parsing time "` +
		pe.Value + `" as "` +
		pe.Layout + `": cannot parse "` +
		pe.LayoutElem + `" as "` +
		pe.ValueElem + `"`
}

func newParseError(layout, value, layoutElem, valueElem string) error {
	return &ParseError{
		Layout:     layout,
		Value:      value,
		LayoutElem: layoutElem,
		ValueElem:  valueElem,
	}
}
