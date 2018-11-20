package config

func logWriter() *zerologWriter {
	return &zerologWriter{}
}

type zerologWriter struct {
}

func (t *zerologWriter) Write(p []byte) (n int, err error) {
	return 0, cfg.OssSaveLog(p)
}
