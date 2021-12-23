package wrap

import (
	"context"
	"fmt"
	"github.com/hachi-n/command-wrap-notify/internal/model"
	"os/exec"
)

func ExecCommand(ctx context.Context, commandPath string, slack *model.Slack) error {
	cmd := exec.CommandContext(ctx, commandPath)

	// ignore stdout
	_, err := cmd.CombinedOutput()
	exitCode := cmd.ProcessState.ExitCode()
	if exitCode != 0 {
		message := generateNotifyDefaultMessage(commandPath, exitCode, err)
		_, err := slack.Notify(message)
		if err != nil {
			return err
		}
	}
	return err
}

func generateNotifyDefaultMessage(commandPath string, exitCode int, err error) string {
	return fmt.Sprintf(
		"command: %s, exit status: %d, error: %v",
		commandPath, exitCode, err,
	)
}
