package provider

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
