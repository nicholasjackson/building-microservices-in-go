package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var message = `"It will be dark soon," he said. "Then I should see the glow of Havana. If I am too far to the eastward I will see the lights of one of the new beaches." I cannot be too far out now, he thought. I hope no one has been too worried. There is only the boy to worry, of course. But I am sure he would have confidence. Many of the older fishermen will worry. Many others too, he thought. I live in a good town.`

var keyPassphrase = "mysecretkey"

func TestCorrectlySignsMessage(t *testing.T) {
	signature, err := SignMessage(message, []byte{})

	assert.NotNil(t, err)
	assert.Equal(t, "abcde", signature)
}

func TestValidateMessageSignature(t *testing.T) {
	valid, err := ValidateSignature(message, "", []byte{})

	assert.NotNil(t, err)
	assert.True(t, valid)
}
