package stringsext

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		alphabet string
		length   int
	}{
		{"12", 10},
		{"1a", 10},
		{"ab", 10},
		{"11", 10},
		{"aa", 10},
	}

	for i := range tests {
		generatedString := Generate(tests[i].alphabet, tests[i].length)
		assert.Len(t, generatedString, tests[i].length)
		assert.Subset(t, []byte(tests[i].alphabet), []byte(generatedString))
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		str    string
		slice  []string
		exists bool
	}{
		{"one", []string{}, false},
		{"one", []string{"one"}, true},
		{"one", []string{"twe", "three"}, false},
		{"one", []string{"one", "twe", "three"}, true},
		{"", []string{"one", "twe", "three"}, false},
	}

	for i := range tests {
		assert.Equal(t, In(tests[i].str, tests[i].slice...), tests[i].exists)
	}
}

func TestTrimSlice(t *testing.T) {
	tests := []struct {
		slice  []string
		result []string
	}{
		{[]string{"one", "twe", "three"}, []string{"one", "twe", "three"}},
		{[]string{"one", "twe ", " three"}, []string{"one", "twe ", " three"}},
		{[]string{"one", "twe", " "}, []string{"one", "twe"}},
		{[]string{"one", "twe", ""}, []string{"one", "twe"}},
	}

	for i := range tests {
		assert.Equal(t, tests[i].result, TrimSlice(tests[i].slice))
	}
}

func TestPointerToString(t *testing.T) {
	str := "test"
	assert.Equal(t, ToString(&str), "test")
	assert.Equal(t, ToString(nil), "")
}

func TestStringToPointer(t *testing.T) {
	assert.Equal(t, *ToPointer("test"), "test")
	assert.Nil(t, ToPointer(""))
}
