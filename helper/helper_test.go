package helper

import "testing"

func TestPathHasFile(t *testing.T) {
	t.Log(GetConfigFile("go.mod"))
}
