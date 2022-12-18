package goini

import (
	"testing"
	"time"
)

type custom string

func (c custom) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *custom) UnmarshalText(b []byte) error {
	v := custom(string((b)))
	*c = v

	return nil
}

func (c custom) String() string {
	return string(c)
}

type customTestStruct struct {
	Foo  custom    `ini:"foo"`
	Date time.Time `ini:"time"`
	Int  int       `ini:"int"`
}

var customTestIni = []byte(`foo = custom-type
time = 2222-12-12T09:49:42.783607764-08:00
int = 42
`)

func TestFoo(t *testing.T) {
	s := &customTestStruct{}

	err := Unmarshal(customTestIni, &s)
	assertNil(t, err, "failed to unmarshal")

	b, err := Marshal(s)
	assertNil(t, err, "failed to test marshal")

	assertEqualBytes(t, b, customTestIni)
}
