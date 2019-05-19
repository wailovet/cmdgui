package helper

import (
	"bytes"
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
)

type Pty struct {
	cmd     *exec.Cmd
	ptyFile *os.File
	Buffer  *bytes.Buffer
}

type ptyWriter struct {
	callback func(p []byte)
}

func (that *ptyWriter) Write(p []byte) (int, error) {
	that.callback(p)
	return len(p), nil
}

func (that *Pty) Write(p []byte) (int, error) {
	n, err := that.ptyFile.Write(p)
	return n, err
}

func (that *Pty) Start(callback func(data []byte)) error {
	var err error
	that.ptyFile, err = pty.Start(that.cmd)
	if err != nil {
		return err
	}

	go func() {
		_, _ = io.Copy(&ptyWriter{
			callback: func(p []byte) {
				that.Buffer.Write(p)
				callback(p)
			},
		}, that.ptyFile)
	}()

	return nil
}

func NewPty(name string, arg ...string) *Pty {
	cmd := exec.Command(name, arg...)
	ptytmp := Pty{
		cmd:    cmd,
		Buffer: bytes.NewBuffer(nil),
	}
	return &ptytmp
}

var instancePty *Pty

func GetInstancePtyBySh() *Pty {
	if instancePty == nil {
		instancePty = NewPty("sh")
	}
	return instancePty
}
