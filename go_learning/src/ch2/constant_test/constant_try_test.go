package constant_test

import "testing"

//连续的节点
const (
	START = iota + 1
	PRODUCER
	CONSUMER
	PUSH_GATE_WAY
	END
)

//标识位
const (
	READER = 1 << iota
	WRITER
	EXECUTE
)

func TestConstant1(t *testing.T) {
	a := 7 //0111
	t.Log(a&READER == READER, a&WRITER == WRITER, a&EXECUTE == EXECUTE)
}

func TestConstantTry(t *testing.T) {

	t.Log(START, PRODUCER, CONSUMER, PUSH_GATE_WAY, END)

}
