package main

import (
	"os"
	"testing"
)

func TestSomeGarbage(t *testing.T) {
	var err error

	file, err := os.Open("test-data/CTR2013.gpx")
	if err != nil {
		t.Error(err)
	}
	g, err := FromXml(file)
	if err != nil {
		t.Error(err)
	}

	l := len(g.TrackPoints())
	if l != 10000 {
		t.Errorf("Did not unmarshal the test data correctly. %d != 10000", l)
	}

	t.Log(g.TrackPoints()[0])
	t.Fail()
}
