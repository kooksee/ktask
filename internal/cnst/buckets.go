package cnst

type oss struct {
	LogBucket  string
	TaskBucket string
}

var Oss = oss{
	LogBucket:  "yb-slogs",
	TaskBucket: "yb-tasks",
}
