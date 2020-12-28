package main

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	// app.Initializing(func(c component.IContainer) error {
	// 	//检查db配置是否正确
	// 	if _, err := c.GetDB(); err != nil {
	// 		return err
	// 	}
	// 	//检查消息队列配置
	// 	if _, err := c.GetQueue(); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

	// app.Micro("/order/request", order.NewRequestHandler)

	// app.MQC("/order/pay", order.NewPayHandler)

	// queryChan := app.GetDynamicQueue()
	// queryChan <- &conf.Queue{
	// 	Queue:   "QTASK:TEST:ORDER-PAY",
	// 	Service: "/order/pay",
	// }

	// qtask.Bind(app, 1)
}
