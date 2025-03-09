package request

import (
	"sdanthinne/gtfs_tracker/internal/request"
	"testing"
)

func TestInit(t *testing.T) {
	exp_out := "Something!"
	out := request.DoSomething()
	if exp_out != out {
		t.Fatalf("Failure! exp: \"%s\" act: \"%s\"", exp_out, out)
	}
}
