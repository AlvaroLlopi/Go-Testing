package clase1

import "testing"

func TestAdd(t *testing.T) {
	want := 8
	t.Logf("want vale: %d\n", want)
	got := Add(2, 3)
	t.Logf("got vale: %d\n", got)
	if got != want {
		t.Errorf("Error: Se esperaba %d, se obtuvo %d", want, got)
	}
}

func TestAddAcum(t *testing.T) {
	want := 7
	got := AddAcum(1, 2, 3)
	if got != want {
		t.Errorf("Error: Se esperaba %d, se obtuvo %d", want, got)
	}
}
