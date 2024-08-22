package main

import "testing"

func TestSayHello(t *testing.T){
	msg := sayHello("Golang")

	want:= "Hello Golang"

	if want != msg{
		t.Error()
	}
}