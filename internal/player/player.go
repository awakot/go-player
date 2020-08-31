package player

import (
	"fmt"
	"time"

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

	// // Filter
	// filter *Filter
}

func (p *Player) New() {
	p.viewApp = &views.Application{}
	p.musicDirPath = "music"
}

// LoadMusic load a music to PunosPanel
func (p *Player) LoadMusic() {
	speaker.Lock()
	// p.ctrl.Paused = true
	speaker.Unlock()

	streamer, format, err := p.DecodeMusicFile(p.musicDirPath)
	if err != nil {
		fmt.Println(err)
	}

	p.sampleRate = format.SampleRate
	p.streamer = streamer
	p.ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, p.streamer)}
	p.resampler = beep.ResampleRatio(4, 1, p.ctrl)
	p.volume = &effects.Volume{Streamer: p.resampler, Base: 2}
	p.gain = &effects.Gain{Streamer: p.resampler, Gain: 2}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	speaker.Play(p.volume)
	// first, pause music
	speaker.Lock()
	p.ctrl.Paused = !p.ctrl.Paused
	speaker.Unlock()
}

// PlayPause is
func (p *Player) PlayPause() {
	speaker.Lock()
	p.ctrl.Paused = !p.ctrl.Paused
	speaker.Unlock()
}

// SetVol set volume
func (p *Player) SetVol(volume float64) {
	speaker.Lock()
	p.volume.Volume = volume
	speaker.Unlock()
}

// Volup increase volume of music
func (p *Player) Volup() {
	speaker.Lock()
	p.volume.Volume += 0.1
	speaker.Unlock()
}

// Voldown decrease volume of music
func (p *Player) Voldown() {
	speaker.Lock()
	p.volume.Volume -= 0.1
	speaker.Unlock()
}
