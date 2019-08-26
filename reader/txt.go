package reader

type TxtReader struct {
}

func (txt *TxtReader) Load(path string) int {
	return 0
}

func (txt *TxtReader) Current() string {
	return ""
}

func (txt *TxtReader) Next() string {
	return ""
}

func (txt *TxtReader) Prev() string {
	return ""
}

func (txt *TxtReader) CurrentPos() int {
	return 0
}

func (txt *TxtReader) Goto(pos int) string {
	return ""
}

func (txt *TxtReader) GetProgress() string {
	return ""
}
