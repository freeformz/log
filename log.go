package log

import (
	"io"
	"io/ioutil"

	"github.com/freeformz/log/plain"
)

type KVExpander interface {
	KVExpand() (interface{}, interface{})
}

type Logger interface {
	Log(v ...interface{}) error
}

type Filter interface {
	Filter(v ...interface{}) []interface{}
}

type Formatter interface {
	Format(v ...interface{}) ([]byte, error)
}

type logger struct {
	data      []interface{}
	filters   []Filter
	formatter Formatter
	writer    io.Writer
}

func New() *logger {
	return &logger{formatter: &plain.Formatter{}, writer: ioutil.Discard}
}

func (l *logger) WithFilter(f Filter) *logger {
	l.filters = append(l.filters, f)
	return l
}

func (l *logger) WithFormatter(f Formatter) *logger {
	l.formatter = f
	return l
}

func (l *logger) WithWriter(w io.Writer) *logger {
	l.writer = w
	return l
}

func (l *logger) Log(v ...interface{}) error {
	// Filter out any messages that would be filtered out.
	for i := range l.filters {
		v = l.filters[i].Filter(v...)
		if len(v) == 0 {
			return nil
		}
	}

	// Expand any type that need to be expanded.
	for i := 0; i < len(v); i++ {
		if e, ok := v[i].(KVExpander); ok {
			key, val := e.KVExpand()
			v[i] = key
			i++
			v = append(v[:i], append([]interface{}{val}, v[i:]...)...)
		}
	}

	// Format the values to a []byte.
	d, err := l.formatter.Format(v...)
	if err != nil {
		return err
	}

	// Write the values out to the writer.
	_, err = l.writer.Write(d)
	return err
}
