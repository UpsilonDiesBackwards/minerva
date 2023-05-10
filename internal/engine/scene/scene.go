package scene

import (
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/renderer"
	"github.com/UpsilonDiesBackwards/minerva/internal/engine/window"
)

type Scene struct {
	entities []*renderer.Entity
}

func NewScene() *Scene {
	return &Scene{
		entities: []*renderer.Entity{},
	}
}

func (s *Scene) AddEntity(entity *renderer.Entity) {
	s.entities = append(s.entities, entity)
}

func (s *Scene) Update(deltaTime float32) {
	for _, entity := range s.entities {
		entity.Update(deltaTime)
	}
}

func (s *Scene) Draw(renderer *renderer.Renderer, win *window.Window) {
	renderer.Draw(s.entities, win)
}
