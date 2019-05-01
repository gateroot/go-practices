package __bowling_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBowling(t *testing.T) {
	assert.Equal(t, 0, Score("--------------------"))
	assert.Equal(t, 10, Score("X------------------"))
	assert.Equal(t, 10, Score("-/------------------"))
	assert.Equal(t, 20, Score("11111111111111111111"))
	assert.Equal(t, 24, Score("X34----------------"))
	assert.Equal(t, 20, Score("5/34----------------"))
	assert.Equal(t, 60, Score("XXX--------------"))
	assert.Equal(t, 40, Score("5/5/5/--------------"))
	assert.Equal(t, 300, Score("XXXXXXXXXXXX"))
	assert.Equal(t, 150, Score("5/5/5/5/5/5/5/5/5/5/5"))
	assert.Equal(t, 90, Score("9-9-9-9-9-9-9-9-9-9-"))
}
