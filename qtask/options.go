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

// WithDeleteDeadline 秒数，设置任务删除截止时间
func WithDeleteDeadline(second int) Option {
	return func(o opts) {
		o["delete_interval"] = second
	}
}

// WithMaxCount 任务执行最大次数
func WithMaxCount(count int) Option {
	return func(o opts) {
		o["max_count"] = count
	}
}

// WithOrderNO 业务单号
func WithOrderNO(orderNO string) Option {
	return func(o opts) {
		o["order_no"] = orderNO
	}
}
