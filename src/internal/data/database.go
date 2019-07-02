package data

import (
	"context"
	"example/github.com/TaskManager/src/tools/status_bar"
)

type Database interface {
	CreateTask(ctx context.Context)
	UpdateTask(ctx context.Context)
	GetTask(ctx context.Context)
	CreateStatusBar(ctx context.Context, statBar status_bar.StatusBar)
	UpdateStatusBar(ctx context.Context)
	GetStatusBar(ctx context.Context)
}
