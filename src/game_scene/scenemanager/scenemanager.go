package scenemanager

import "github.com/hajimehoshi/ebiten"

type scenemanager struct {
}

var manager *scenemanager

func init() {
	manager = &scenemanager{}
}

// Update 하기
func Update(screen *ebiten.Image) error {
	return nil
}
