package sample01

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "hello world, hoge"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

type Cases []Case
type Case struct {
	in, want string
}

func TestTableSample(t *testing.T) {
	c1 := Case{"hoge", "egoh"}
	c2 := Case{"fuga", "aguf"}

	cs := Cases{}

	for i := 0; i < 100; i++ {
		cs = append(cs, c1)
		cs = append(cs, c2)

		fmt.Printf("loop %d\n", i+1)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("%+v", cs)
}
