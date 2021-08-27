package project

import "github.com/rajihawa/mask-off/app/entity"

type Inmem struct {
	m map[string]*entity.Project
}

func NewInmem() *Inmem {
	var m = map[string]*entity.Project{}
	return &Inmem{
		m: m,
	}
}

func (r *Inmem) Get(id string) (*entity.Project, error) {
	p := r.m[id]
	if p == nil {
		return nil, entity.ErrNotFound
	}
	return p, nil
}

func (r *Inmem) List() ([]*entity.Project, error) {
	var pl []*entity.Project
	for _, v := range r.m {
		if v != nil {
			pl = append(pl, v)
		}
	}
	return pl, nil
}

func (r *Inmem) Create(p *entity.Project) (string, error) {
	r.m[p.ID] = p
	return p.ID, nil
}

func (r *Inmem) Update(p *entity.Project) error {
	_, err := r.Get(p.ID)
	if err != nil {
		return nil
	}
	r.m[p.ID] = p
	return nil
}

func (r *Inmem) Delete(id string) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
