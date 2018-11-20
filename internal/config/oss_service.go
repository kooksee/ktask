package config

import (
	"bytes"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"io/ioutil"
	"strings"
)

func (t *config) OssSaveLog(data []byte) error {
	bk, _ := cfg.GetOss().Bucket(cnst.Oss.LogBucket)

	// 计算出key
	level := json.Get(data, "level").ToString()
	time := json.Get(data, "time").ToString()
	serviceMame := json.Get(data, "service_name").ToString()
	serviceId := json.Get(data, "service_id").ToString()
	ks := []string{serviceMame, serviceId, "logs", level}
	if mth := json.Get(data, "method").ToString(); mth != "" {
		ks = append(ks, mth)
	}

	if mth := json.Get(data, "mth").ToString(); mth != "" {
		ks = append(ks, mth)
	}

	ks = append(ks, time)
	k := strings.Join(ks, "/") + ".json"

	return utils.Retry(3, func() error {
		return bk.PutObject(k, bytes.NewBuffer(data))
	})
}

func (t *config) OssSaveTask(taskId string, task []byte) error {
	bk, _ := cfg.GetOss().Bucket(cnst.Oss.TaskBucket)
	return bk.PutObject(t.taskKey(taskId), bytes.NewReader(task))
}

func (t *config) taskKey(k string) string {
	return "tasks/" + k + ".json"
}

func (t *config) OssUpdateTaskById(result *kts.Task) error {
	bk, _ := cfg.GetOss().Bucket(cnst.Oss.TaskBucket)

	idt, err := bk.GetObject(t.taskKey(result.TaskID))
	if err != nil {
		return err
	}

	task := &kts.Task{}
	if err := task.DecodeFromReader(idt); err != nil {
		return err
	}

	task.Status = result.Status
	task.Output = result.Output
	task.Code = result.Code
	task.UpdateTime = result.UpdateTime

	return t.OssSaveTask(task.TaskID, task.Encode())
}

func (t *config) OssGetTaskById(taskId string) (dt []byte, err error) {
	bk, _ := cfg.GetOss().Bucket(cnst.Oss.TaskBucket)

	k := t.taskKey(taskId)
	if dt, ok := t.Cache.Get(k); ok {
		return dt.([]byte), nil
	}

	return dt, func() error {
		idt, err := bk.GetObject(k)
		if err != nil {
			return err
		}

		dt, err = ioutil.ReadAll(idt)
		if err != nil {
			return err
		}

		t.Cache.SetDefault(k, dt)
		return nil
	}()
}
