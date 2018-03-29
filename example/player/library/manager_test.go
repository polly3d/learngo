package library

import (
	"learngo/example/player/music"
	"testing"
)

func TestNewManager(t *testing.T) {
	manger := NewManager()
	if manger == nil {
		t.Error("New method is fail")
	}
}

func TestAddAndFind(t *testing.T) {
	music := music.Music{
		ID:   "1",
		Name: "hello muisc",
	}

	manager := NewManager()
	manager.Add(&music)

	findResult := manager.Find(music)

	if findResult == nil {
		t.Error("Can add and not find")
	}

	if music.Name != findResult.Name {
		t.Error("Not find")
	}
}

func TestRemove(t *testing.T) {
	music := music.Music{
		ID:   "1",
		Name: "hello music",
	}

	manager := NewManager()
	manager.Add(&music)
	manager.Remove(0)
	len := manager.Len()
	if len != 0 {
		t.Error("Not remove")
	}
}
