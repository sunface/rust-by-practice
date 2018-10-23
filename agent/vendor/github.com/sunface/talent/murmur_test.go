package talent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//  16.1 ns/op             0 B/op          0 allocs/op
func BenchmarkGetHash(b *testing.B) {
	v := []byte("a/b/c/d/e/f/g/h/this/is/emitter")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MurMurHash(v)
	}
}

func TestMeHash(t *testing.T) {
	h := MurMurHash([]byte("me"))
	assert.Equal(t, uint32(2539734036), h)
}

func TestGetHash(t *testing.T) {
	h := MurMurHash([]byte("+"))
	if h != 1815237614 {
		t.Errorf("Hash %d is not equal to %d", h, 1815237614)
	}
}

func TestGetHash2(t *testing.T) {
	h := MurMurHash([]byte("hello world"))
	if h != 4008393376 {
		t.Errorf("Hash %d is not equal to %d", h, 1815237614)
	}
}
