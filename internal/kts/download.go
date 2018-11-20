package kts

type Download struct {
	Type     string            `json:"type,omitempty"`
	Url      string            `json:"type,omitempty"`
	Header   map[string]string `json:"header,omitempty"`
	TimeWait int               `json:"time_wait,omitempty"`
}

func (t *Download) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}
