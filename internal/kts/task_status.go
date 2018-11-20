package kts

type TaskStatus struct {
	Status string `json:"status,omitempty"`
	TaskID string `json:"task_id,omitempty"`
	Code   string `json:"code,omitempty"`
}

func (t *TaskStatus) Encode() []byte {
	dt, _ := json.Marshal(t)
	return dt
}

func (t *TaskStatus) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}
