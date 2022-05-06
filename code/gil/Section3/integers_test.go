package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewIntegers(t *testing.T) {
	i := newIntegers()

	expectedLen := 27

	assert.Equal(t, expectedLen, len(i), "Length should be equal")

	expectedLastElement := 9 * 1000

	assert.Equal(t, expectedLastElement, i.getLastElement(), "Unexpected last element")
}

func TestSaveToFileAndLoad(t *testing.T) {
	i1 := newIntegers()
	const fileName = "testFile.dat"

	os.Remove(fileName)

	i1.saveTofile(fileName)

	i2 := newIntegersFromFile(fileName)

	assert.Equal(t, len(i1), len(i2), "Lengths should be the same")

	for i, _ := range i1 {
		assert.Equal(t, i1[i], i2[i], fmt.Sprintf("Elements if index %v should be equal", i))
	}

	os.Remove(fileName)
}
