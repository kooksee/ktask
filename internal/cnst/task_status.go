package cnst

type status struct {
	Pending string
	Failed  string
	Success string
}

var TaskStatus = status{
	Pending: "PENDING",
	Failed:  "FAILED",
	Success: "SUCCESS",
}
