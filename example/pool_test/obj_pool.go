// Package object_pool 对象池的使用，用来换成难以创建可重复使用的对象
package object_pool

import (
	"errors"
	"time"
)

type ReuseableObj struct {
}

type ObjPool struct {
	bufChan chan *ReuseableObj // 用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReuseableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReuseableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReuseableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}

}

func (p *ObjPool) ReleaseObj(obj *ReuseableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
