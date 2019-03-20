package processor

import (
	"golang.org/x/net/context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAccessToken(t *testing.T) {

	url := "https://www.logomaker.com/wp-content/uploads/2015/06/Logo-Samples2-07-min.jpg"
	ctx := context.Background()

	err := Process(ctx, url)

	assert.Nil(t, err)
}
