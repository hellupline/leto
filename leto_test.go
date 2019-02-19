package leto

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func TestOpen(t *testing.T) {
	name, blob := "hello.txt", []byte("hello world")
	Register(name, blob)

	f, err := Open(name)
	require.NoError(t, err)

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	assert.Equal(t, blob, data)
}
