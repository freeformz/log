package plain

import "bytes"
import "fmt"

type Formatter struct{}

func (f *Formatter) Format(v ...interface{}) ([]byte, error) {
	var b bytes.Buffer
	for i := range v {
		switch k := v[i].(type) {
		case string:
			if _, err := b.WriteString(k); err != nil {
				return b.Bytes(), err
			}
		case fmt.Stringer:
			if _, err := b.WriteString(k.String()); err != nil {
				return b.Bytes(), err
			}

		default:
			b.WriteString(fmt.Sprint(k))
		}
		if i < len(v)-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString(`\n`)
	return b.Bytes(), nil
}
