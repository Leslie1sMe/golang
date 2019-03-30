package writer

type MysqlWriter struct {
}

func (MysqlWriter) Write(interface{}, string, string) error {
	panic("implement me")
}
