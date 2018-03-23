package sample01

import (
	"fmt"
	"testing"
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

func tableSample() *Cases {
	c1 := Case{"hoge", "egoh"}
	c2 := Case{"fuga", "aguf"}

	cs := Cases{}

	for i := 0; i < 100000; i++ {
		cs = append(cs, c1)
		cs = append(cs, c2)
	}

	return &cs
}

func TestTableSample1(t *testing.T) {
	cases := tableSample()

	i := 0
	for _, c := range *cases {
		i = i + 1
		if i%100 == 0 {
			fmt.Println(c)
			t.Errorf("fail test, i=%d", i)
		}
	}
}

func TestTableSample2(t *testing.T) {
	cases := tableSample()

	i := 0
	for _, c := range *cases {
		i = i + 1
		fmt.Println(fmt.Sprintf("%+v", c))
		if i%100 == 0 {
			fmt.Println(c)
			t.Errorf("fail test, i=%d", i)
		}
	}
}
