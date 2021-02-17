package echo2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Service(t *testing.T) {
	result := Service()
	assert := assert.New(t)

	assert.Equal(result, "Howdy: 20")
}
