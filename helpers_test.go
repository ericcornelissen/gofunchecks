package main

type mockPrinter struct {
	callCount *uint
	calls     *[][]interface{}
}

func (p mockPrinter) Print(msgs ...interface{}) {
	if p.callCount != nil {
		*(p.callCount)++
	}

	if p.calls != nil {
		*(p.calls) = append(*(p.calls), msgs)
	}
}
