package main
import "testing"


func TestHello(t*testing.T){
	got := hello("feroz")
	want := "Hello, feroz"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

