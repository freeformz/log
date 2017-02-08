package level

type levels int

const (
	Debug levels = iota
	Info
	Warn
	Err
)

func (l levels) String() string {
	switch l {
	case Err:
		return "error"
	case Warn:
		return "warn"
	case Info:
		return "info"
	case Debug:
		return "debug"
	default:
		return "unknown"
	}
}

func (l levels) KVExpand() (interface{}, interface{}) {
	return "level", l
}

type Filter struct {
	At levels
}

func (f *Filter) Filter(v ...interface{}) []interface{} {
	for i, _ := range v {
		if l, ok := v[i].(levels); ok && l >= f.At {
			return v
		}
	}
	return nil
}
