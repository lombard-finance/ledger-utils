package common

import (
	"bytes"
	"errors"
	"testing"
)

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func AssertError(t *testing.T, err error, errorTypes ...error) {
	if err == nil {
		t.FailNow()
	}
	if len(errorTypes) > 0 {
		for _, target := range errorTypes {
			if !errors.Is(err, target) {
				t.Errorf("error '%s' is supposed to have '%s' as well", err.Error(), target.Error())
			}
		}
	}
}

func AssertTrue(t *testing.T, value bool) {
	if !value {
		t.FailNow()
	}
}

func AssertFalse(t *testing.T, value bool) {
	if value {
		t.FailNow()
	}
}

func EqualStrings(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func EqualBytes(t *testing.T, expected []byte, actual []byte) {
	if !bytes.Equal(expected, actual) {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}
