package library

import (
	"errors"
	"learngo/example/player/music"
)

// MusicManager is
type MusicManager struct {
	music []music.Music
}

// NewManager is
func NewManager() *MusicManager {
	return &MusicManager{make([]music.Music, 0)}
}

// Add is
func (m *MusicManager) Add(music *music.Music) {
	m.music = append(m.music, *music)
}

// Find ...
func (m *MusicManager) Find(music music.Music) *music.Music {
	for _, m := range m.music {
		if m.Name == music.Name {
			return &m
		}
	}
	return nil
}

// Get ...
func (m *MusicManager) Get(index int) (music *music.Music, err error) {
	if index < 0 || index >= len(m.music) {
		return nil, errors.New("out of range")
	}
	return &m.music[index], nil
}

// Remove ...
func (m *MusicManager) Remove(index int) (musicRemove *music.Music, err error) {
	if index < 0 || index >= len(m.music) {
		return nil, errors.New("out of range")
	}

	removeMusic := m.music[index]

	if index < len(m.music)-1 {
		m.music = append(m.music[:index-1], m.music[index+1:]...)
	} else if index == 0 {
		m.music = make([]music.Music, 0)
	}

	return &removeMusic, nil
}

// Len ...
func (m *MusicManager) Len() int {
	return len(m.music)
}
