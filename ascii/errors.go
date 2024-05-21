package ascii

import "errors"

var (
	ErrInvalideTemplate = errors.New("invalide template name")
	ErrNoASCII          = errors.New("only ASCII characters are allowed")
)
