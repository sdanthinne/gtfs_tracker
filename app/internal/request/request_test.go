package request

import (
    "testing"
)

func TestInit(t *testing.T) {
    exp_out := "Something!"
    out := DoSomething()
    if exp_out != out {
        t.Fatalf("Failure! exp: \"%s\" act: \"%s\"", exp_out, out)
    }
}
