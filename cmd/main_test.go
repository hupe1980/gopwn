package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmdHelp(t *testing.T) {
	var b bytes.Buffer
	cmd := newRootCmd("")
	cmd.SetOut(&b)
	cmd.SetArgs([]string{"-h"})
	assert.NoError(t, cmd.Execute())
}

func TestRootCmdVersion(t *testing.T) {
	var b bytes.Buffer
	cmd := newRootCmd("1.2.3")
	cmd.SetOut(&b)
	cmd.SetArgs([]string{"-v"})
	assert.NoError(t, cmd.Execute())
	assert.Equal(t, "gopwn version 1.2.3\n", b.String())
}
