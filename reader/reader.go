package reader

type Reader interface {
	Load(path string) error
	Current() string
	Next() string
	Prev() string
	First() string
	Last() string
	CurrentPos() int
	Goto(pos int) string
	GetProgress() string
}
