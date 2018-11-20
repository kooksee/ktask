package kts

import (
	"github.com/google/uuid"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/utils"
	"io"
	"time"
)

type Task struct {
	Callback    string `json:"callback,omitempty"`
	CreateTime  int    `json:"create_time,omitempty"`
	UpdateTime  int    `json:"update_time,omitempty"`
	Input       string `json:"input,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	Status      string `json:"status,omitempty"`
	TaskID      string `json:"task_id,omitempty"`
	TopicName   string `json:"topic_name,omitempty"`
	Output      string `json:"output,omitempty"`
	Code        string `json:"code,omitempty"`
	Event       string `json:"event,omitempty"`
	Tx          string `json:"tx,omitempty"`
}

func (t *Task) Mock() *Task {
	return &Task{
		CreateTime:  int(time.Now().Unix()),
		Input:       "hello",
		ServiceName: "test",
		Event:       "test",
		TaskID:      uuid.New().String(),
		TopicName:   "log",
		Status:      cnst.TaskStatus.Pending,
		Tx:          cnst.TaskTx.Send,
	}
}

func (t *Task) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Task) DecodeFromReader(data io.Reader) error {
	dec := json.NewDecoder(data)
	return dec.Decode(t)
}

func (t *Task) Encode() []byte {
	dt, err := json.Marshal(t)
	utils.MustNotError(err)
	return dt
}
