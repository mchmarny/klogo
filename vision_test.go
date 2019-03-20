package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAccessToken(t *testing.T) {

	url := "https://storage.googleapis.com/kdemo-logos/0.png"

	desc, err := getLogoFromUrl(url)

	assert.Nil(t, err)
	assert.NotEmpty(t, desc)
}
