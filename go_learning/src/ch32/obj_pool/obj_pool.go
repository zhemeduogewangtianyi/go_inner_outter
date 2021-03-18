package obj_pool

import (
	"errors"
	"time"
)

type Reusable struct {
}

type ObjPool struct {
	bufChan chan *Reusable
}

func NewObjPool(numOfObj int) *ObjPool {

	objPool := ObjPool{}
	ch := make(chan *Reusable, numOfObj)

	for i := 0; i < numOfObj; i++ {
		ch <- &Reusable{}
	}

	objPool.bufChan = ch
	return &objPool
}

func (p *ObjPool) GetObj(duration time.Duration) (*Reusable, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(duration):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReturnObj(reusable *Reusable) error {
	select {
	case p.bufChan <- reusable:
		return nil
	default:
		return errors.New("overflow")

	}
}
