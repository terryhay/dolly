package file_proxy

// Opt contains mock methods of initialize mocked decorator object
type Opt struct {
	SlotClose       func() error
	SlotWriteString func(string) error
}

// Mock constructs mocked Proxy
func Mock(opt Opt) Proxy {
	return &impl{
		opt: opt,
	}
}
