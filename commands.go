package main

import (
	"github.com/mitchellh/cli"
	"github.com/munjeli/supermodel/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"checkout": func() (cli.Command, error) {
			return &command.CheckoutCommand{
				Meta: *meta,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Meta: *meta,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				Meta: *meta,
			}, nil
		},
		"merge": func() (cli.Command, error) {
			return &command.MergeCommand{
				Meta: *meta,
			}, nil
		},
		"pull": func() (cli.Command, error) {
			return &command.PullCommand{
				Meta: *meta,
			}, nil
		},
		"push": func() (cli.Command, error) {
			return &command.PushCommand{
				Meta: *meta,
			}, nil
		},
		"show": func() (cli.Command, error) {
			return &command.ShowCommand{
				Meta: *meta,
			}, nil
		},
		"status": func() (cli.Command, error) {
			return &command.StatusCommand{
				Meta: *meta,
			}, nil
		},
		"sync": func() (cli.Command, error) {
			return &command.SyncCommand{
				Meta: *meta,
			}, nil
		},
		"update": func() (cli.Command, error) {
			return &command.UpdateCommand{
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
