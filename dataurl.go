package dataurl

import (
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

type DataURL struct {
	io.Reader

	Type   string
	Params map[string]string
}

func Parse(rawurl string) (*DataURL, error) {
	u := &DataURL{
		Params: map[string]string{},
	}

	r := rawurl
	if !strings.HasPrefix(r, "data:") {
		return nil, &Error{"parse", rawurl, errors.New("missing protocol scheme")}
	}
	r = r[5:]

	if r[0] != ';' {
		for i := 0; i < len(r); i++ {
			if r[i] == ';' || r[i] == ',' {
				r = r[i:]
				break
			}
			u.Type += string(r[i])
		}
	}

L:
	for r[0] != ',' {
		if r[0] == ';' {
			k := ""
			v := ""
			for i := 0; i < len(r); i++ {
				if r[i] == '=' {
					r = r[i:]
					break
				}
				if r[i] == ',' {
					r = r[i:]
					break L
				}
				k += string(r[i])
			}
			for i := 0; i < len(r); i++ {
				if r[i] == ';' {
					r = r[i:]
					continue L
				}
				if r[i] == ',' {
					r = r[i:]
					break L
				}
				v += string(r[i])
			}
			u.Params[k] = v
		}
	}

	if r[0] != ',' {
		return nil, &Error{"parse", rawurl, errors.New("expected comma, reached end-of-line")}
	}
	u.Reader = base64.NewDecoder(base64.StdEncoding, strings.NewReader(r[1:]))

	return u, nil
}

type Error struct {
	Op  string
	URL string
	Err error
}

func (e *Error) Error() string {
	return e.Op + " " + e.URL + ": " + e.Err.Error()
}
