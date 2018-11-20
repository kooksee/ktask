package kts

type TaskResult struct {
	TaskID     string `json:"task_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Output     string `json:"output,omitempty"`
	Code       string `json:"code,omitempty"`
	UpdateTime int    `json:"update_time,omitempty"`
}

func (t *TaskResult) Encode() []byte {
	dt, _ := json.Marshal(t)
	return dt
}

func (t *TaskResult) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}
