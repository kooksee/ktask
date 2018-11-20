package cnst

type consumer struct {
	Task       string
	TaskResult string
	Callback   string
	TaskIds    string
	Log        string
	HttpGet    string
}

var Consumer = consumer{
	Task:       "tasks",
	TaskIds:    "task-ids",
	TaskResult: "tasks-result",
	Callback:   "callback",
	Log:        "log",
	HttpGet:    "http_get",
}
