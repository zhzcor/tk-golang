package app

import (
	"context"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent/importtask"
)

type ImportTask struct {
}

func (i ImportTask) UpdateImportTask(ctx context.Context, id, status int, remark string, total int) error {
	if errorno := i.ImportTaskIdExist(ctx, id); errorno != nil {
		return errorno
	}
	res := store.WithContext(ctx).ImportTask.UpdateOneID(id).
		SetStatus(uint8(status)).
		SetRemark(remark)
	if total > 0 {
		res = res.SetTotal(total)
	}
	_, err := res.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

//添加导入任务
func (i ImportTask) AddImportTask(ctx context.Context, createdAdminId int, importName, path string) (int, error) {
	task, err := store.WithContext(ctx).ImportTask.Create().
		SetStatus(1).
		SetCreatedAdminID(createdAdminId).
		SetImportName(importName).
		SetPath(path).Save(ctx)
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}

//id导入任务是否存在
func (i ImportTask) ImportTaskIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.ImportTask.
		Query().SoftDelete().
		Where(importtask.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "导入任务不存在",
		})
	}
	return nil
}
