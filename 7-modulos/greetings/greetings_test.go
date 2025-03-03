package greetings

import (
	"regexp"
	"testing"
)

//Deben comenzar con el nombre Test
//(t *tetsing.T) para reportar el estado de la prueba
func TestHelloName(t *testing.T){
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}