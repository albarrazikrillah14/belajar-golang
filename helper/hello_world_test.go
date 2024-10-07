package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T) {
	t.Run("Test for Albarra", func(t *testing.T) {
		result := HelloWorld("Albarra")

		require.Equal(t, result, "Hi, Albarra")
	})

	t.Run("Test for Zikrillah", func(t *testing.T) {
		result := HelloWorld("Zikrillah")

		require.Equal(t, result, "Hi, Zikrillah")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before test")

	m.Run()

	//after
	fmt.Println("After test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on macOS")
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zikri")

	assert.Equal(t, result, "Hi, Zikri", "Result must be 'Hello, Zikri'")
}

func TestHelloWorld(t *testing.T) {
	name := "Albarra Zikrillah"
	result := HelloWorld(name)

	if result != "Hi, "+name {
		t.Errorf("result must be 'Hello, %s', got '%s'", name, result)
	}
}
