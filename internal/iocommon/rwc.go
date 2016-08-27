package iocommon

import (
	"io"
)

type MergedRWC struct {
	io.Reader
	io.Writer
	io.Closer
}

type multiCloser struct {
	closers []io.Closer
}

func (mc multiCloser) Close() error {
	var err error
	for _, closer := range mc.closers {
		newErr := closer.Close()
		if err == nil {
			err = newErr
		}
	}
	return err
}

func MultiCloser(closers ...io.Closer) io.Closer {
	return multiCloser{closers}
}
