package main

import (
	"context"
	"fmt"

	"mvdan.cc/sh/interp"
)

// vopt_*

func shVOptIf(ctx context.Context, path string, args []string) error {
	cmd, args := args[0], args[1:]
	if err := argRangeCheck(cmd, len(args), 2, 3); err != nil {
		return err
	}
	opt := args[0]
	if Options.Has(opt) {
		if len(args) > 1 {
			write(ctx, args[1])
		}
		return nil
	}
	if len(args) > 2 {
		write(ctx, args[2])
	}
	return nil
}

func shVOptWith(ctx context.Context, path string, args []string) error {
	cmd, args := args[0], args[1:]
	if err := argRangeCheck(cmd, len(args), 1, 2); err != nil {
		return err
	}
	opt := args[0]
	with, flag := "without", opt
	if len(args) > 1 {
		flag = args[1]
	}
	if Options.Has(opt) {
		with = "with"
	}
	write(ctx, "--"+with+"-"+flag)
	return nil
}

func shVOptEnable(ctx context.Context, path string, args []string) error {
	cmd, args := args[0], args[1:]
	if err := argRangeCheck(cmd, len(args), 1, 2); err != nil {
		return err
	}
	opt := args[0]
	enable, flag := "disable", opt
	if len(args) > 1 {
		flag = args[1]
	}
	if Options.Has(opt) {
		enable = "enable"
	}
	write(ctx, "--"+enable+"-"+flag)
	return nil
}

func shVOptConflict(ctx context.Context, path string, args []string) error {
	cmd, args := args[0], args[1:]
	if err := argRangeCheck(cmd, len(args), 2, 2); err != nil {
		return err
	}
	if Options.Has(args[0]) && Options.Has(args[1]) {
		return fmt.Errorf("%s: cannot set options %s and %s simultaneously",
			getPackage(ctx), args[0], args[1])
	}
	return nil
}

func shVOptBool(ctx context.Context, path string, args []string) error {
	cmd, args := args[0], args[1:]
	if err := argRangeCheck(cmd, len(args), 2, 2); err != nil {
		return err
	}
	opt := args[0]
	prop, val := args[1], "false"
	if Options.Has(opt) {
		val = "true"
	}
	write(ctx, "-D"+prop+"="+val)
	return nil
}

var shFuncs = map[string]interp.ModuleExec{
	"vopt_if":       shVOptIf,
	"vopt_with":     shVOptWith,
	"vopt_enable":   shVOptEnable,
	"vopt_conflict": shVOptConflict,
	"vopt_bool":     shVOptBool,
}
