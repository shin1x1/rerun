package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Rerun struct {
	command string
	args    []string
	sleep   time.Duration
}

func NewRerun(command string, args []string, sleep int) *Rerun {
	return &Rerun{
		command: command,
		args:    args,
		sleep:   time.Duration(sleep),
	}
}

func (r *Rerun) Run() {
	var wg sync.WaitGroup

	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(r.command, r.args)

			cmd := exec.Command(r.command, r.args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Duration(r.sleep) * time.Second)

		}()

		wg.Wait()
	}
}
