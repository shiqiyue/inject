package inject

// 注入前执行
type InjectBefore interface {
	Before()
}

// 注入后执行
type InjectAfter interface {
	After()
}

// 注入后执行,比after后执行
type InjectSetup interface {
	SetUp()
}
