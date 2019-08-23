package reader

type GeneralReader interface {
	Load(path string) int
	Current() string
	Next() string
	Prev() string
	CurrentPos() int
	Goto(pos int) string
	GetProgress() string
}
