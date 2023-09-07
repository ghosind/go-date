package date

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

func TestMarshalAndUnmarshalText(t *testing.T) {
	a := assert.New(t)
	time := Now()
	wantTime := "2000-01-01T00:00:00Z"

	err := time.UnmarshalText([]byte(wantTime))
	a.NilNow(err)

	txt, err := time.MarshalText()
	a.NilNow(err)
	a.Equal(string(txt), wantTime)
}

func TestMarshalAndUnmarshalJSON(t *testing.T) {
	a := assert.New(t)
	time := Now()
	wantTime := `"2000-01-01T00:00:00Z"`

	err := time.UnmarshalJSON([]byte(wantTime))
	a.NilNow(err)

	txt, err := time.MarshalJSON()
	a.NilNow(err)
	a.Equal(string(txt), wantTime)
}

func TestMarshalAndUnmarshalBinary(t *testing.T) {
	a := assert.New(t)
	testTime := Now()
	wantTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	bin, _ := wantTime.MarshalBinary()

	err := testTime.UnmarshalBinary(bin)
	a.NilNow(err)

	testBin, err := testTime.MarshalBinary()
	a.NilNow(err)
	a.Equal(bin, testBin)
}

func TestGeoEncodeAndDecode(t *testing.T) {
	a := assert.New(t)
	testTime := Now()
	wantTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	gob, _ := wantTime.GobEncode()

	err := testTime.GobDecode(gob)
	a.NilNow(err)

	testGob, err := testTime.GobEncode()
	a.NilNow(err)
	a.Equal(gob, testGob)
}
