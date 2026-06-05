package reader

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type TxtReader struct {
	content  []string
	pos      int
	filePath string
}

func NewTxtReader() *TxtReader {
	return &TxtReader{}
}

func (txt *TxtReader) Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	txt.filePath = path

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
				if strings.TrimSpace(buffer.String()) != "" {
					txt.content = append(txt.content, buffer.String())
					buffer.Reset()
				} else {
					buffer.Reset()
				}
			}
		} else {
			buffer.Write(r.Bytes())
		}
	}

	txt.pos = 0
	txt.loadProgress()
	return nil
}

func (txt *TxtReader) getProgressFile() string {
	hash := md5.Sum([]byte(txt.filePath))
	hashStr := hex.EncodeToString(hash[:])
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".glance", hashStr+".progress")
}

func (txt *TxtReader) SaveProgress() error {
	progressFile := txt.getProgressFile()
	dir := filepath.Dir(progressFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(progressFile, []byte(strconv.Itoa(txt.pos)), 0644)
}

func (txt *TxtReader) loadProgress() {
	progressFile := txt.getProgressFile()
	data, err := os.ReadFile(progressFile)
	if err != nil {
		return
	}
	pos, err := strconv.Atoi(string(data))
	if err != nil {
		return
	}
	if pos >= 0 && pos < len(txt.content) {
		txt.pos = pos
	}
}

func (txt *TxtReader) Current() string {
	return txt.content[txt.pos]
}

func (txt *TxtReader) Next() string {
	txt.pos++

	if txt.pos <= len(txt.content)-1 {
		return txt.content[txt.pos]
	}
	txt.pos = len(txt.content) - 1

	return "END"
}

func (txt *TxtReader) Prev() string {
	txt.pos--

	if txt.pos < 0 {
		txt.pos = 0
	}

	return txt.content[txt.pos]
}

func (txt *TxtReader) First() string {
	txt.pos = 0
	return txt.content[0]
}

func (txt *TxtReader) Last() string {
	txt.pos = len(txt.content) - 1
	return txt.content[len(txt.content)-1]
}

func (txt *TxtReader) CurrentPos() int {
	return txt.pos
}

func (txt *TxtReader) Goto(pos int) string {
	if pos < 0 {
		pos = 0
	}

	if pos > len(txt.content)-1 {
		pos = len(txt.content) - 1
	}

	txt.pos = pos
	return txt.content[txt.pos]
}

func (txt *TxtReader) GetProgress() string {
	return fmt.Sprintf("(%d / %d)", txt.pos+1, len(txt.content))
}
