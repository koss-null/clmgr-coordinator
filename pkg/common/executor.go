package common

import (
	"errors"
	"os/exec"
	"reflect"
	"time"
)

type (
	executor struct {
		cmd []string
		time.Duration
	}

	Executor interface {
		Exec() (string, error)
		SetOp([]string)
		SetTimeout(time.Duration)
	}
)

func NewExecutor() Executor {
	return &executor{}
}

func (e *executor) SetOp(cmd []string) {
	e.cmd = cmd
}

func (e *executor) SetTimeout(t time.Duration) {
	e.Duration = t
}

func (e *executor) Exec() (string, error) {
	cmd := exec.Command(e.cmd[1], e.cmd[2:]...)

	res := make(chan interface{}, 1)
	go func() {
		out, err := cmd.Output()
		if err == nil {
			res <- string(out)
		} else {
			res <- err
		}
	}()
	select {
	case <-time.After(e.Duration):
		if err := cmd.Process.Kill(); err != nil {
			return "", err
		}
		return "", errors.New("operation timeout exceeded")
	case instance := <-res:
		inst := reflect.ValueOf(instance)
		if inst.Kind() == reflect.String {
			return instance.(string), nil
		}
		return "", instance.(error)
	}
}
