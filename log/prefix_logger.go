package log

import (
	"fmt"
	"io"
)

type prefixLogger struct {
	io.Writer
}

// NewPrefixLogger returns a basic logger that encodes keyvals as simple "k=v"
// pairs to the Writer.
func NewPrefixLogger(w io.Writer) Logger {
	return &prefixLogger{w}
}

func (l prefixLogger) Log(keyvals ...interface{}) error {
	if len(keyvals)%2 == 1 {
		panic("odd number of keyvals")
	}
	for i := 0; i < len(keyvals); i += 2 {
		if i != 0 {
			if _, err := fmt.Fprint(l.Writer, " "); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(l.Writer, "%s=%v", keyvals[i], keyvals[i+1]); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(l.Writer); err != nil {
		return err
	}
	return nil
}
