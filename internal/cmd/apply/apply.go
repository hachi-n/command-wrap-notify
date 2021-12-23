package apply

import (
	"context"
	"github.com/hachi-n/command-wrap-notify/internal/config"
	"github.com/hachi-n/command-wrap-notify/internal/model"
	"github.com/hachi-n/command-wrap-notify/internal/wrap"
	"os/signal"
	"syscall"
)

func Apply(commandPath string) error {
	c := config.NewApplicationConfig()
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	s := model.NewSlack(c.Mention, c.Channel, c.Endpoint, c.Username)
	return wrap.ExecCommand(ctx, commandPath, s)
}
