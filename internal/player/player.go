package player

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/gdamore/tcell/views"
)

//Player go-player Player
type Player struct {
	sampleRate   beep.SampleRate
	streamer     beep.StreamSeeker
	ctrl         *beep.Ctrl
	resampler    *beep.Resampler
	volume       *effects.Volume
	gain         *effects.Gain
	leftDoppler  beep.Streamer
	rightDoppler beep.Streamer
	cuePoint     int
	musicTitle   string
	musicDirPath string

	//front
	viewApp *views.Application
	view    views.View
	panel   views.Widget
}

// PlayPause is
func (p *Player) PlayPause() {
	speaker.Lock()
	p.ctrl.Paused = !p.ctrl.Paused
	speaker.Unlock()
}
