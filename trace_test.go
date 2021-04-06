package trace

import (
	"fmt"
	"testing"
)

func TestNewtraceID(*testing.T) {
	tr := NewtraceID("192.169.1.1")
	ip, t := ParseTrace(tr)
	fmt.Println(ip, t)
	ip, t = ParseTrace("1613552838058100200050")
	fmt.Println(ip, t)
}
