package maps

type Dictionary map[string] string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exit")
)

type DictionaryErr string 

func (e DictionaryErr) Error() string {
	return string(e)
}

func(d Dictionary) Search(word string) (string, error) {
	// map look up can return two values, 
	// the second value is a boolean that indicates whether
	// the key was found succesfully. 
	// allows us to differentiate between keys that just don't 
	// have values vs. keys that don't exist.
	definition, ok :=  d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func(d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
			return err
	}
	return nil
}

func(d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word) 
	
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
		
	}
	return nil
}

func(d Dictionary) Delete(word string) {
	delete(d, word)
}
