package env_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gg/env"
)

func testEnvAquire[T any](t *testing.T, key string, expected error, opts ...env.AquireOption) *T {
	if key == "" {
		key = uuid.NewString()
	}

	result, err := env.Aquire[T](key, opts...)

	if expected != nil {
		assert.NotNil(t, err)
		assert.IsType(t, expected, err)
	} else {
		assert.Nil(t, err)
		assert.NotNil(t, result)
	}

	return result
}

func TestEnvAquireEmpty(t *testing.T) {
	testEnvAquire[any](t, "", env.AquireUnmarshalError, env.IgnoreEmpty())
}

func TestEnvAquireEmptyError(t *testing.T) {
	testEnvAquire[any](t, "", env.AquireEmptyError)
}

func TestEnvAquireEmptyString(t *testing.T) {
	key := uuid.NewString()
	os.Setenv(key, "\"\"")

	testEnvAquire[any](t, key, nil)
}
