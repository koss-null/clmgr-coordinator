package resource

import (
	"errors"
	"sync"
)

type (
	pool struct {
		key sync.Locker
		// resources are in map to make search quicker
		res map[string]Resource
	}

	Pool interface {
		Add(Resource) error
		Remove(string) error
		Change(Resource) error
		Get(string) Resource
		Contains(string) bool
	}
)

var LocalPool Pool
var GlobalPool Pool

func NewPool() Pool {
	return &pool{key: &sync.Mutex{}}
}

func (p *pool) Add(res Resource) error {
	p.key.Lock()
	defer p.key.Unlock()

	if _, ok := p.res[res.Name]; ok {
		return errors.New("this resource have already been added")
	}

	p.res[res.Name] = res
	return nil
}

func (p *pool) Remove(res string) error {
	p.key.Lock()
	defer p.key.Unlock()

	if _, ok := p.res[res]; !ok {
		return errors.New("there is no such resource")
	}

	delete(p.res, res)
	return nil
}

func (p *pool) Change(res Resource) error {
	p.key.Lock()
	defer p.key.Unlock()

	if _, ok := p.res[res.Name]; !ok {
		return errors.New("there is no such resource")
	}

	p.res[res.Name] = res
	return nil
}

func (p *pool) Get(res string) Resource {
	if r, ok := p.res[res]; !ok {
		return Resource{}
	} else {
		return r
	}
}

func (p *pool) Contains(res string) bool {
	_, ok := p.res[res]
	return ok
}
