// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mgrun

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var ErrNonZeroExit = errors.New("process exited with non-zero status")

type Executor struct {
	command   string
	dir       string
	extraEnvs []string
	cmd       *exec.Cmd

	onStdout func(string)
	onStderr func(string)

	exitCode int
	sMutex   sync.Mutex
}

func New(command string) *Executor {
	return &Executor{
		command:  command,
		exitCode: -1,
	}
}

func (e *Executor) SetDir(path string) {
	e.sMutex.Lock()
	defer e.sMutex.Unlock()
	e.dir = path
}

func (e *Executor) AddEnv(key, value string) {
	e.sMutex.Lock()
	defer e.sMutex.Unlock()
	e.extraEnvs = append(e.extraEnvs, fmt.Sprintf("%s=%s", key, value))
}

func (e *Executor) OnStdout(fn func(string)) {
	e.onStdout = fn
}

func (e *Executor) OnStderr(fn func(string)) {
	e.onStderr = fn
}

func (e *Executor) ExitCode() int {
	e.sMutex.Lock()
	defer e.sMutex.Unlock()
	return e.exitCode
}

func (e *Executor) Run() error {
	e.cmd = buildShellCommand(e.command)
	if e.dir != "" {
		e.cmd.Dir = e.dir
	}
	e.cmd.Env = append(os.Environ(), e.extraEnvs...)

	stdout, err := e.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := e.cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := e.cmd.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go e.readStream(stdout, os.Stdout, e.onStdout, &wg)
	go e.readStream(stderr, os.Stderr, e.onStderr, &wg)

	wg.Wait()

	err = e.cmd.Wait()

	if e.cmd.ProcessState != nil {
		e.sMutex.Lock()
		e.exitCode = e.cmd.ProcessState.ExitCode()
		e.sMutex.Unlock()
	}

	if err != nil {
		return ErrNonZeroExit
	}

	return nil
}

func (e *Executor) readStream(
	r io.Reader,
	mirror io.Writer,
	callback func(string),
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if mirror != nil {
			fmt.Fprintln(mirror, line)
		}

		if callback != nil {
			callback(line)
		}
	}
}

func buildShellCommand(command string) *exec.Cmd {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/C", "powershell", "-Command", command)
	default:
		// linux / darwin
		return exec.Command("sh", "-c", command)
	}
}
