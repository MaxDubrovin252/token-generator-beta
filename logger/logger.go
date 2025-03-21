package logger

import (
	"log/slog"
	"os"
	"token-generator/logger/pretty"
)

const (
	ENVLOCAL = "local"
	ENVDEV   = "dev"
	ENVPROD  = "prod"
)

func SetLogger(env string) (*slog.Logger, error) {
	var log *slog.Logger

	switch env {
	case ENVLOCAL:
		log = setupPrettySlog()

	case ENVDEV:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case ENVPROD:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log, nil
}
func setupPrettySlog() *slog.Logger {
	opts := pretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
