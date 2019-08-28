package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type TxtReader struct {
	content []string
	pos     int
}

func NewTxtReader() *TxtReader {
	return &TxtReader{}
}

func (txt *TxtReader) Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	cmd := exec.Command("fold", "-w", "80", "-s", path)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return err
	}

	r := bufio.NewScanner(stdout)
	r.Split(bufio.ScanRunes)
	buffer := bytes.NewBuffer(make([]byte, 0))

	for r.Scan() {
		line := r.Text()
		if line == "\r" {
			continue
		}

		if line == "\n" {
			if buffer.Len() > 0 {
				txt.content = append(txt.content, buffer.String())
				buffer.Reset()
			}
		} else {
			buffer.Write(r.Bytes())
		}
	}

	txt.pos = 0
	return nil
}

func (txt *TxtReader) Current() string {
	return txt.content[txt.pos]
}

func (txt *TxtReader) Next() string {
	txt.pos++

	if txt.pos < len(txt.content)-1 {
		return txt.content[txt.pos]
	}

	return "END"
}

func (txt *TxtReader) Prev() string {
	txt.pos--

	if txt.pos < 0 {
		txt.pos = 0
	}

	return txt.content[txt.pos]
}

func (txt *TxtReader) CurrentPos() int {
	return txt.pos
}

func (txt *TxtReader) Goto(pos int) string {
	if pos-1 < 0 || pos > len(txt.content) {
		return txt.content[txt.pos]
	}

	txt.pos = pos - 1
	return txt.content[txt.pos]
}

func (txt *TxtReader) GetProgress() string {
	return fmt.Sprintf("(%d / %d)", txt.pos+1, len(txt.content))
}
