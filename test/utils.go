package test

import (
	"bytes"
	"context"
	"testing"

	"blobdev.com/pandaroll/cmd"
	"github.com/stretchr/testify/require"
)

func executeCommand(t *testing.T, command string, args ...string) (string, error) {
	rootCmd := cmd.NewRootCommand()

	require.NotNil(t, rootCmd, "Command not defined in executeCommand")

	args = append([]string{command}, args...)

	b := bytes.NewBufferString("")

	for _, c := range rootCmd.Commands() {
		c.SetOut(b)
		c.SetErr(b)
	}

	rootCmd.SetOut(b)
	rootCmd.SetErr(b)
	rootCmd.SetArgs(args)
	err := rootCmd.ExecuteContext(context.Background())

	return b.String(), err
}

// TODO: Figure out how to set env before the root init is called
/*
type Env map[string]string

func executeCommandWithEnv(t *testing.T, command string, env Env, args ...string) (string, error) {
	for key, value := range env {
		os.Setenv(key, value)
	}

	rootCmd := cmd.NewRootCommand()

	require.NotNil(t, rootCmd, "Command not defined in executeCommand")

	args = append([]string{command}, args...)

	b := bytes.NewBufferString("")

	for _, c := range rootCmd.Commands() {
		c.SetOut(b)
		c.SetErr(b)
	}

	rootCmd.SetOut(b)
	rootCmd.SetErr(b)
	rootCmd.SetArgs(args)

	err := rootCmd.ExecuteContext(context.Background())

	return b.String(), err
}
*/
