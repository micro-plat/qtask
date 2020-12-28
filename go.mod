module github.com/micro-plat/qtask

go 1.14

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/hydra v0.13.2
	github.com/micro-plat/lib4go v1.0.9
	github.com/zkfy/go-oci8 v1.0.0
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go
