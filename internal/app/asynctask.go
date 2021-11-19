package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gserver/httpapi/appapi/request"
	"gserver/internal/errorno"
	"gserver/pkg/asynctask"
	"gserver/pkg/baijiayun"
	"gserver/pkg/log"
	"time"
)

func init() {
	asynctask.RegisterTask("order_video_upload", asynctask.Task{
		Handle: OrderUploadSyncTask,
		Options: []asynq.Option{
			asynq.Timeout(20 * time.Second),
		},
	})
	asynctask.RegisterTask("user_question_record", asynctask.Task{
		Handle: UserQuestionRecord,
		Options: []asynq.Option{
			asynq.Timeout(20 * time.Second),
		},
	})
}

//视频上传
func OrderUploadSyncTask(ctx context.Context, t *asynq.Task) error {
	payLoad := t.Payload
	id, err := payLoad.GetInt("id")
	if err != nil {
		log.Error(fmt.Sprintf("order upload param error id: %s", err.Error()))
		return errorno.NewParamErr(err)
	}
	videoUrl, err := payLoad.GetString("video_url")
	if err != nil {
		log.Error(fmt.Sprintf("order upload param error video_url: %s", err.Error()))
		return errorno.NewParamErr(err)
	}
	uploadUrl, err := payLoad.GetString("upload_url")
	if err != nil {
		log.Error(fmt.Sprintf("order upload param error upload_url: %s", err.Error()))
		return errorno.NewParamErr(err)
	}
	filename, err := payLoad.GetString("filename")
	if err != nil {
		log.Error(fmt.Sprintf("order upload param error filename: %s", err.Error()))
		return errorno.NewParamErr(err)
	}
	b := baijiayun.BjyCloud{}
	b.OrderUpload(ctx, id, videoUrl, uploadUrl, filename)
	return nil
}

//用户题目统计
func UserQuestionRecord(ctx context.Context, t *asynq.Task) error {
	payLoad := t.Payload
	questionList, err := payLoad.GetString("question_list")
	if err != nil {
		log.Error(fmt.Sprintf("user question record error : %s", err.Error()))
		return errorno.NewParamErr(err)
	}
	fmt.Println(questionList)
	var list []request.UserRecodeAdd
	err = json.Unmarshal([]byte(questionList), &list)
	if err != nil {
		log.Error(fmt.Sprintf("user question record format error : %s", err.Error()))
		return errorno.NewParamErr(err)
	}

	questionBankId, err := payLoad.GetInt("question_bank_id")
	if err != nil {
		log.Error(fmt.Sprintf("user question record question bank id error : %s", err.Error()))
		return errorno.NewParamErr(err)
	}

	userId, err := payLoad.GetInt("user_id")
	if err != nil {
		log.Error(fmt.Sprintf("user question record user id error : %s", err.Error()))
		return errorno.NewParamErr(err)
	}

	tk := TkQuestionBank{}
	err = tk.SetUserQuestionBankRecord(ctx, questionBankId, userId, list)
	if err != nil {
		return err
	}
	return nil
}
