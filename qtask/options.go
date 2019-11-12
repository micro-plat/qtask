package qtask

type opts map[string]interface{}

//Option 配置选项
type Option func(opts)

//WithDeadline 秒数，设置任务截止时间
func WithDeadline(second int) Option {
	return func(o opts) {
		o["max_timeout"] = second
	}
}
