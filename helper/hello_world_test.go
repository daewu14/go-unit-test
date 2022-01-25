package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Daewu")
	}
}

func BenchmarkHelloWorldBintara(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bintara")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test") //  bisa connect db atau init variable

	m.Run()

	fmt.Println("After Unit Test")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Daewu")
	assert.Equal(t, "Hello world Daewu", result, "result harusnya => Hello world Daewu")
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Daewu")
	if result != "Hello world Daewu" {
		t.Error("result harusnya => Hello world Daewu")
		//t.Fail()
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldDaewu(t *testing.T) {
	result := HelloWorld("Daewu")
	if result != "Hello world Daewu" {
		t.Fatal("result harusnya => Hello world Daewu")
		//t.FailNow()
	}
	fmt.Println("TestHelloWorldDaewu done")
}