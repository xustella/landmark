package landmark

import (
	"errors"
	"unicode/utf8"

	"github.com/xustella/coordinates"
)

type Places struct {
	name string
	coordinates.Coordinates
}

func (p *Places) SetLandmark(l string) error {
	if utf8.RuneCountInString(l) > 25 {
		return errors.New("Landmark name too long")
	}
	p.name = l
	return nil
}

func (p *Places) Landmark() string {
	return p.name
}
