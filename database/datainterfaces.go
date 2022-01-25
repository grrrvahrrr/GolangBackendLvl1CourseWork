package database

type Writer interface {
	WriteURL(fullURL string, shortURL string) error
	WriteData(shortURL string, data string) error
}

type Reader interface {
	ReadURL(ShortURL string) (FullURL string, err error)
	ReadData(ShortURL string) (Data string, err error)
}

type Sorter interface {
	Sort(filename string) error
}
