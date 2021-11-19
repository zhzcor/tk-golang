package asynctask

import (
	"context"
	"github.com/hibiken/asynq"
	"gserver/internal/store"
	"gserver/pkg/log"
	"sync"
)

var tasks sync.Map
var client *asynq.Client
var srv *asynq.Server

type Task struct {
	Handle  func(ctx context.Context, t *asynq.Task) error
	Options []asynq.Option
}

func RegisterTask(name string, t Task) {
	tasks.Store(name, t)
}

func Enqueue(name string, payload map[string]interface{}) error {
	t := asynq.NewTask(name, payload)
	res, err := client.Enqueue(t)
	if err != nil {
		return err
	}
	log.Info("task enqueue:", res.ID)
	return nil
}

func Background(ctx context.Context, redisAddr string) error {
	r := asynq.RedisClientOpt{Addr: redisAddr}
	srv = asynq.NewServer(r, asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 10,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	})
	client = asynq.NewClient(r)
	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.Use(func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(c context.Context, task *asynq.Task) error {
			// 注入ctx
			c = context.WithValue(c, store.CtxDatabaseKey, store.WithContext(ctx))
			return next.ProcessTask(c, task)
		})
	})
	// ...register other handlers...
	tasks.Range(func(key, value interface{}) bool {
		if task, o := value.(Task); o {
			mux.HandleFunc(key.(string), task.Handle)
		}
		return true
	})
	if err := srv.Start(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	return nil
}

func Stop() {
	if srv != nil {
		srv.Quiet()
		srv.Stop()
	}
	if client != nil {
		client.Close()
	}
}
