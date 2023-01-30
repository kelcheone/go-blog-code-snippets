package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func willError() error {
	return errors.New("Error!")
}

func willNotError() error {
	return nil
}

func TestError(t *testing.T) {
	assert.Error(t, willError())
	assert.NoError(t, willNotError())
}
