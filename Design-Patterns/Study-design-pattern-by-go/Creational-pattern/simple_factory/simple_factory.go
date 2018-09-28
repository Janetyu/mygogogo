package simple_factory

import (
	"fmt"
)

/**
 * 简单工厂模式
 */

// CellPhone is interface
type CellPhone interface {
	Call(name string) string
}

// NewCall return CellPhone instance by warning_signal
// warning_signal 警号
func NewCall(warning_signal int) CellPhone {
	if warning_signal == 1 {
		return &police_1{}
	} else if warning_signal == 2 {
		return &police_2{}
	}
	return nil
}

// police_1 is one of CellPhone implement
type police_1 struct{}

// notify name to work
func (*police_1) Call(name string) string {
	return fmt.Sprintf("Hi %s , it's time to work", name)
}

// police_2 is another CellPhone implement
type police_2 struct{}

// notify name to work
func (*police_2) Call(name string) string {
	return fmt.Sprintf("Hello %s , it's time to work", name)
}
