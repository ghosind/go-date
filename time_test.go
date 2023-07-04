package date

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

func TestMarshalAndUnmarshalText(t *testing.T) {
	time := Now()
	wantTime := "2000-01-01T00:00:00Z"

	if err := time.UnmarshalText([]byte(wantTime)); err != nil {
		t.Errorf("T.UnmarshalText() error = %v, want nil", err)
	}

	if txt, err := time.MarshalText(); err != nil {
		t.Errorf("Time.MarshalText() error = %v, want nil", err)
	} else {
		assert.DeepEqual(t, string(txt), wantTime)
	}
}

func TestMarshalAndUnmarshalJSON(t *testing.T) {
	time := Now()
	wantTime := `"2000-01-01T00:00:00Z"`

	if err := time.UnmarshalJSON([]byte(wantTime)); err != nil {
		t.Errorf("T.UnmarshalJSON() error = %v, want nil", err)
	}

	if txt, err := time.MarshalJSON(); err != nil {
		t.Errorf("Time.MarshalJSON() error = %v, want nil", err)
	} else {
		assert.DeepEqual(t, string(txt), wantTime)
	}
}

func TestMarshalAndUnmarshalBinary(t *testing.T) {
	testTime := Now()
	wantTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	bin, _ := wantTime.MarshalBinary()

	if err := testTime.UnmarshalBinary(bin); err != nil {
		t.Errorf("Time.UnmarshalBinary() error = %v, want nil", err)
	}

	if testBin, err := testTime.MarshalBinary(); err != nil {
		t.Errorf("Time.MarshalBinary() error = %v, want nil", err)
	} else {
		assert.DeepEqual(t, bin, testBin)
	}
}

func TestGeoEncodeAndDecode(t *testing.T) {
	testTime := Now()
	wantTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	gob, _ := wantTime.GobEncode()

	if err := testTime.GobDecode(gob); err != nil {
		t.Errorf("Time.GobDecode() error = %v, want nil", err)
	}

	if testGob, err := testTime.GobEncode(); err != nil {
		t.Errorf("Time.GobEncode() error = %v, want nil", err)
	} else {
		assert.DeepEqual(t, gob, testGob)
	}
}
