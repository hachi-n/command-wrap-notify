package wrap

import (
	"context"
	"github.com/hachi-n/command-wrap-notify/internal/model"
	"github.com/hachi-n/command-wrap-notify/internal/notifier"
	"os/exec"
)

func ExecCommand(ctx context.Context, commandPath string, n notifier.Notifier) error {
	cmd := exec.CommandContext(ctx, commandPath)

	// ignore stdout
	_, err := cmd.CombinedOutput()
	exitCode := cmd.ProcessState.ExitCode()
	if exitCode != 0 {
		commandMessage := model.NewCommandMessage(commandPath, exitCode, err)
		_, err := n.Notify(commandMessage)
		if err != nil {
			return err
		}
	}
	return err
}
