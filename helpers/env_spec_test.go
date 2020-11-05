package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvWhenIsTestingIsTrue(t *testing.T) {
	assert := assert.New(t)
	expectedEnv := ".env.test"
	isTesting = true
	actualEnv := GetEnv(isTesting)

	assert.Equal(expectedEnv, actualEnv)
}

func TestGetEnvWhenIsTestingIsFalse(t *testing.T) {
	assert := assert.New(t)
	isTesting = false
	expectedEnv := ".env"
	actualEnv := GetEnv(isTesting)

	assert.Equal(expectedEnv, actualEnv)
}
