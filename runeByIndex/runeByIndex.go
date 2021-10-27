package runeByIndex

import "errors"

func RuneByIndex(s *string, i *int) (r rune, nilPointerErr error) {

	defer func(){
		if err := recover(); err != nil {
			nilPointerErr = errors.New("nil pointer dereference")
			r = 0
		}
	}()
	r = []rune(*s)[*i]
	return r, nilPointerErr
}
