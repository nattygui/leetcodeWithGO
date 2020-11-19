package reformat

import "testing"

func TestReformat(t *testing.T) {
	str := "covid2019"
	want := "c2o0v1i9d"
	result := reformat(str)
	if result != want {
		t.Error("error")
	}
}
