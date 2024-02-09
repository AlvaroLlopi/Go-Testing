package clase3

import (
	"reflect"
	"testing"
)

func TestPerro(t *testing.T) {
	want := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Criollo",
		},
	}

	got := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Criollo",
		},
	}
	t.Logf("Want %p,got %p", want, got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Se esperaba: %v, se obtuvo %v", want, got)
	}
}
