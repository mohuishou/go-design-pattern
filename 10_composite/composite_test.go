package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	assert.Equal(t, 20, got)
}
