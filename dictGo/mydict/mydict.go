package mydict

import "errors"

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("that word already exists")
	errCantUpdate = errors.New("can't update non-existing word")
)

// Dictionary type
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		// nil == null
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
