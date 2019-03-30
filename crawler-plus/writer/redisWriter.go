package writer

type RedisWriter struct {
}

func (RedisWriter) Write(interface{}, string, string) error {
	panic("implement me")
}
