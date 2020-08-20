package singleton_test

import (
	singleton "github.com/mohuishou/go-design-pattern/01_singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, singleton.GetInstance(), singleton.GetInstance())
		}
	})
}

