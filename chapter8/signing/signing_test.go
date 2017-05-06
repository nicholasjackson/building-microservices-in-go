package signing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var message = `"It will be dark soon," he said. "Then I should see the glow of Havana. If I am too far to the eastward I will see the lights of one of the new beaches." I cannot be too far out now, he thought. I hope no one has been too worried. There is only the boy to worry, of course. But I am sure he would have confidence. Many of the older fishermen will worry. Many others too, he thought. I live in a good town.`

var signature = "I9OxVapU2k+cWB3ixfu9249oJt7cbZqTCToROiSWG5YNOz5QsS4joHS8QiLsvyBl5zyRA6vAdalS2ymwVyKnqNL9K4bEnXU6fU38GoYHBAwi93C5o/HItdmQGn1xk1xUgVSLfyvc9Sjqr6Mrlvw//xSduWg2ho4/qo8b5KtuCenb9UbosKU2cUx/8XjNg0uzzCJwYMZitfni0LRn5bkNn2f3FbGij2GY5X+SM0adUh/fPfRRj/W6heWQC3Rx+VA7EqMNAuBScZr76b1I2z1Q0oyjChgqiWd2xHEbvVH1DZ2ebh/H7Q/IxECYzlitnR7I8UYLQKEWOTigDgd66RijGw=="

func TestCorrectlySignsMessage(t *testing.T) {
	s, err := SignMessageWithPrivateKey(message)

	assert.Nil(t, err)
	assert.Equal(t, s, signature)
}

func TestValidateMessageSignature(t *testing.T) {
	err := ValidateMessageWithPublicKey(message, signature)

	assert.Nil(t, err)
}
