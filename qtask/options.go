package qtask

type opts map[string]interface{}

//Option 配置选项
type Option func(opts)

//WithFirstTry 秒数，设置首次重试放入队列时间
func WithFirstTry(second int) Option {
	return func(o opts) {
		o["first_timeout"] = second
	}
}

//WithDeadline 秒数，设置任务截止时间
func WithDeadline(second int) Option {
	return func(o opts) {
		o["max_timeout"] = second
	}
}
