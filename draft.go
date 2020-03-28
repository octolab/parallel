package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	platform "go.octolab.org/toolkit/cli/cobra"
)

const unknown = "unknown"

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

var tool cli = func(executor commander, output io.Writer, shutdown func(code int)) {
	defer func() {
		if r := recover(); r != nil {
			shutdown(failure)
		}
	}()
	executor.AddCommand(platform.NewCompletionCommand(), platform.NewVersionCommand(commit, date, version))
	if err := executor.Execute(); err != nil {
		shutdown(failure)
	}
	shutdown(success)
}

type cli func(executor commander, output io.Writer, shutdown func(code int))

type commander interface {
	AddCommand(...*cobra.Command)
	Execute() error
}

func New() *cobra.Command {
	return &cobra.Command{Use: "semaphore",
		Short: "CLI tool based on https://github.com/kamilsk/semaphore package to execute commands in parallel."}
}

func draft() { tool(New(), os.Stderr, os.Exit) }
