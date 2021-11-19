package asynctask

import (
	"context"
	"github.com/hibiken/asynq"
	"gserver/pkg/log"
	"testing"
	"time"
)

func TestBackground(t *testing.T) {
	RegisterTask("test", Task{
		Handle: func(ctx context.Context, t *asynq.Task) error {
			log.Info("pong", t)
			return nil
		},
		Options: []asynq.Option{
			asynq.Timeout(3 * time.Second),
		},
	})
	go func() {
		err := Background(context.Background(), "10.16.0.23:6379")
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(1 * time.Second)
	err := Enqueue("test", map[string]interface{}{
		"data": "ping",
	})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	Stop()

}
