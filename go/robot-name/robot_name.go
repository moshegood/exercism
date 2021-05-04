package robotname

import (
	"errors"
	"fmt"
)

type Robot string

var id int = 0
var letters []rune = []rune{'A', 'A'}
var ErrNoMore = errors.New("No more robots allowed")

func (r *Robot) Name() (string, error) {
	if *r == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return string(*r), nil
}

func (r *Robot) Reset() error {
	if letters[0] > 'Z' {
		return ErrNoMore
	}
	*r = Robot(fmt.Sprintf(string(letters)+"%03d", id))
	id++
	if id > 999 {
		id = 0
		letters[1]++
		if letters[1] > 'Z' {
			letters[1] = 'A'
			letters[0]++
		}
	}
	return nil
}
