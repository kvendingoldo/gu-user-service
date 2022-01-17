package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type PetStore struct {
	Pets   map[int64]Pet
	NextId int64
	Lock   sync.Mutex
}

func NewPetStore() *PetStore {
	return &PetStore{
		Pets:   make(map[int64]Pet),
		NextId: 1000,
	}
}

func (p *PetStore) FindPetByID(c *gin.Context, petId int64) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	c.JSON(http.StatusOK, "ok")
}
