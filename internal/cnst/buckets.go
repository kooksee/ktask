package cnst

type oss struct {
	LogBucket   string
	TaskBucket  string
	ImageBucket string
}

var Oss = oss{
	LogBucket:   "yb-slogs",
	TaskBucket:  "yb-tasks",
	ImageBucket: "yb-galaxy",
}
