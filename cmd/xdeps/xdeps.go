package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

type depsResult struct {
	index int
	file  string
	deps  stringLists
	err   error
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("ERR ") // All stderr output is errors

	flag.Var(Options, "o", "")
	flag.Parse()

	inputs := flag.Args()
	results := make([]depsResult, len(inputs))

	out := make(chan depsResult)
	// Use a simple semaphore to avoid loading too many files / starting too many goroutines.
	// TODO: better management of nofiles / processes.
	sema := make(chan struct{}, runtime.NumCPU()*16)

	for i, f := range inputs {
		i, f := i, filepath.Clean(f)
		go func() {
			sema <- struct{}{}
			defer func() { <-sema }()
			deps, err := extractDeps(f)
			out <- depsResult{i, f, deps, err}
		}()
	}

	// Read results back
	for range inputs {
		r := <-out
		if r.err != nil {
			log.Fatalf("%s: %v", r.file, r.err)
		}
		results[r.index] = r
	}

	// Print results as json (so I can look at the data with jq)
	var buf bytes.Buffer
	type raw string
	enc := json.NewEncoder(&buf)

	push := func(args ...interface{}) {
		for _, arg := range args {
			switch arg := arg.(type) {
			case raw:
				buf.WriteString(string(arg))
			case rune:
				buf.WriteRune(arg)
			default:
				if err := enc.Encode(arg); err != nil {
					log.Fatalf("could not encode %T(%v): %v", arg, arg, err)
				}
				if sz := buf.Len() - 1; sz >= 0 {
					buf.Truncate(sz)
				}
			}
		}
	}

	for _, r := range results {
		push('{', "file", ':', r.file)
		for _, k := range r.deps.Keys() {
			push(',', k, ':', r.deps[k])
		}
		push(raw("}\n"))
		buf.WriteTo(os.Stdout)
	}
}

func parseFile(path string) (*syntax.File, error) {
	var r io.Reader = os.Stdin
	if path != "-" {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}

	parser := syntax.NewParser(syntax.Variant(syntax.LangBash))
	file, err := parser.Parse(r, path)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", path, err)
	} else if parser.Incomplete() {
		return nil, fmt.Errorf("%s: incomplete template", path)
	}
	return file, nil
}

func argRangeCheck(name string, n, min, max int) error {
	if n >= min && n <= max {
		return nil
	}
	if min == max {
		return fmt.Errorf("%s: expected %d arguments, got %d", name, min, n)
	}
	return fmt.Errorf("%s: expected %d..%d arguments, got %d", name, min, max, n)
}

// write writes a string to a module context's stdout (linked in ctx).
func write(ctx context.Context, s string) error {
	mod, ok := interp.FromModuleContext(ctx)
	if !ok {
		return fmt.Errorf("unable to acquire module context")
	}
	_, err := io.WriteString(mod.Stdout, s)
	return err
}

// getPackage returns the value of the pkgname environment variable in a module context.
func getPackage(ctx context.Context) string {
	mod, ok := interp.FromModuleContext(ctx)
	if !ok {
		return "<no-pkgname>"
	}
	name, ok := mod.Env.Get("pkgname").Value.(string)
	if !ok {
		return "<no-pkgname>"
	}
	return name
}

var permittedExecs = newStringSet(
	"awk",
	"egrep",
	"fgrep",
	"find",
	"grep",
	"sed",
	"date",
	"sha256sum",
	"sha1sum",
	"md5sum",
	"shasum",
)

func nopOpen(ctx context.Context, path string, flag int, perm os.FileMode) (io.ReadWriteCloser, error) {
	return nopStream(0), nil
}

func limitedExec(ctx context.Context, path string, args []string) error {
	cmd := args[0]
	if fn := shFuncs[cmd]; fn != nil {
		return fn(ctx, path, args)
	}

	if !permittedExecs.Has(cmd) {
		return fmt.Errorf("unrecognized function or permitted exec: %q", cmd)
	}

	mod, ok := interp.FromModuleContext(ctx)
	if !ok {
		return fmt.Errorf("unable to acquire module context")
	}

	pkgname := getPackage(ctx)

	cmdexec := exec.Command(cmd, args[1:]...)
	cmdexec.Stdin, cmdexec.Stdout, cmdexec.Stderr = mod.Stdin, mod.Stdout, newPrefixWriter(mod.Stderr, pkgname+": ")
	var env []string
	mod.Env.Each(func(name string, vr expand.Variable) bool {
		if s, ok := vr.Value.(string); ok {
			env = append(env, name+"="+s)
		}
		return true
	})

	err := cmdexec.Run()
	if err != nil {
		return fmt.Errorf("%s: %v", pkgname, err)
	}
	return nil
}

func extractDeps(path string) (deps stringLists, err error) {
	file, err := parseFile(path)
	if err != nil {
		return nil, err
	}

	// TODO: capture all IO and do something with it?
	// var (
	// 	stdin  = &bytes.Buffer{}
	// 	stdout = &bytes.Buffer{}
	// 	stderr = &bytes.Buffer{}
	// )

	runner, err := interp.New(interp.StdIO(nopStream(0), os.Stdout, os.Stderr))
	if err != nil {
		return nil, err
	}

	runner.Exec = limitedExec
	runner.Open = nopOpen

	ctx := context.Background()
	if err = runner.Run(ctx, file); err != nil {
		return nil, fmt.Errorf("%s: %v", path, err)
	}

	deps = stringLists{}

	for name, vr := range runner.Vars {
		if !strings.HasSuffix(name, "depends") {
			continue
		}

		switch val := vr.Value.(type) {
		case string:
			deps[name] = strings.Fields(val)
		case []string:
			// Array not used here
		case map[string]string:
			// Associative array not used here
		}
	}

	return deps, nil
}
