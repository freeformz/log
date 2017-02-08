package logfmt

import (
	"bytes"
	"sync"

	"github.com/go-logfmt/logfmt"
)

type Formatter struct{}

type logfmtEncoder struct {
	*logfmt.Encoder
	buf bytes.Buffer
}

func (l *logfmtEncoder) Reset() {
	l.Encoder.Reset()
	l.buf.Reset()
}

var logfmtEncoderPool = sync.Pool{
	New: func() interface{} {
		var enc logfmtEncoder
		enc.Encoder = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

func (f *Formatter) Format(v ...interface{}) ([]byte, error) {
	enc := logfmtEncoderPool.Get().(*logfmtEncoder)
	enc.Reset()
	defer logfmtEncoderPool.Put(enc)

	if err := enc.EncodeKeyvals(v...); err != nil {
		return nil, err
	}

	// Add newline to the end of the buffer
	if err := enc.EndRecord(); err != nil {
		return nil, err
	}

	return enc.buf.Bytes(), nil
}
