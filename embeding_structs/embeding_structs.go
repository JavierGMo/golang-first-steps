package embeding_structs

import "fmt"

type ExampleImp struct {
	num int64
}

type Base struct {
	num int64
}

func (b Base) describe() string {
	return fmt.Sprintf("Base witj num=%v", b.num)
}

type Container struct {
	Base
	str string
}
