package ttt 

import (
	"testing"
	"github.com/stretchrcom/testify/assert"
)

func TestValidateInputForTrue(t *testing.T) {
	t.Log("Returns True for valid input")
  assert.True(t, ValidateInput("7"))
}

func TestValidateInputForFalse(t *testing.T) {
	t.Log("Returns False for an invalid input")
	assert.False(t, ValidateInput("b"))
}

func TestValidateInputForFalseAlso(t *testing.T) {
	t.Log("Returns False for invlaid number")
	assert.False(t, ValidateInput("10"))
}