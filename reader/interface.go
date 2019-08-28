package reader

type GeneralReader interface {
	Load(path string) error
	Current() string
	Next() string
	Prev() string
	CurrentPos() int
	Goto(pos int) string
	GetProgress() string
}
