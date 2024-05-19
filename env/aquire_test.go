package env_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gg/env"
)

func TestEnvAquireEmpty(t *testing.T) {
	key := uuid.NewString()
	result, err := env.Aquire[any](key, env.IgnoreEmpty())

	assert.NotNil(t, err)
	assert.IsType(t, env.AquireUnmarshalError, err)
	assert.Nil(t, result)
}

func TestEnvAquireEmptyError(t *testing.T) {
	key := uuid.NewString()
	result, err := env.Aquire[any](key)

	assert.NotNil(t, err)
	assert.IsType(t, env.AquireEmptyError, err)
	assert.Nil(t, result)
}

func TestEnvAquireEmptyString(t *testing.T) {
	key := uuid.NewString()
	os.Setenv(key, "\"\"")

	result, err := env.Aquire[string](key)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestEnvAquireString(t *testing.T) {
}
