package cnst

type tx struct {
	Send     string
	Result   string
	Callback string
	Done     string
}

var TaskTx = tx{
	Send:     "Send",
	Result:   "Result",
	Callback: "Callback",
	Done:     "Done",
}
