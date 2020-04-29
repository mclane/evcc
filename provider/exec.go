package provider

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/andig/evcc/util"
	"github.com/kballard/go-shellquote"
)

// Script implements shell script-based providers and setters
type Script struct {
	log     *util.Logger
	script  string
	args    []string
	timeout time.Duration
}

// NewScriptProvider creates a script provider.
// Script execution is aborted after given timeout.
func NewScriptProvider(log *util.Logger, script string, timeout time.Duration) *Script {
	if script == "" {
		log.FATAL.Fatalf("config: missing script")
	}

	args, err := shellquote.Split(script)
	if err != nil {
		log.FATAL.Fatalf("config: cannot parse script: %s", script)
	}

	return &Script{
		log:     util.NewLogger("exec"),
		script:  script,
		args:    args,
		timeout: timeout,
	}
}

// StringGetter returns string from exec result. Only STDOUT is considered.
func (e *Script) StringGetter() (string, error) {
	ctx := context.Background()

	if e.timeout > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, e.timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, e.args[0], e.args[1:]...)
	b, err := cmd.Output()

	s := strings.TrimSpace(string(b))

	if err != nil {
		// use STDOUT if available
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			s = strings.TrimSpace(string(ee.Stderr))
		}

		e.log.ERROR.Printf("%s: %s", strings.Join(e.args, " "), s)
		return "", err
	}

	e.log.TRACE.Printf("%s: %s", strings.Join(e.args, " "), s)
	return s, nil
}

// IntSetter invokes script with parameter replaced by int value
func (e *Script) IntSetter(param string) IntSetter {
	// return func to access cached value
	return func(i int64) error {
		cmd, err := util.ReplaceFormatted(e.script, map[string]interface{}{
			param: i,
		})

		if err == nil {
			exec := NewScriptProvider(e.log, cmd, e.timeout)
			_, err = exec.StringGetter()
		}

		return err
	}
}

// BoolSetter invokes script with parameter replaced by bool value
func (e *Script) BoolSetter(param string) BoolSetter {
	// return func to access cached value
	return func(b bool) error {
		cmd, err := util.ReplaceFormatted(e.script, map[string]interface{}{
			param: b,
		})

		if err == nil {
			exec := NewScriptProvider(e.log, cmd, e.timeout)
			_, err = exec.StringGetter()
		}

		return err
	}
}
