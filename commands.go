package main

import (
	"github.com/dcaba/autoscaling-optimizer/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"fetch": func() (cli.Command, error) {
			return &command.FetchCommand{
				Meta: *meta,
			}, nil
		},
		"simulate": func() (cli.Command, error) {
			return &command.SimulateCommand{
				Meta: *meta,
			}, nil
		},
		"combine": func() (cli.Command, error) {
			return &command.CombineCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
