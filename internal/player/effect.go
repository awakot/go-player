package player

import (
	"time"

	"github.com/faiface/beep/speaker"
)

// Fforward fast-forward music
func (p *Player) Fforward() {
	speaker.Lock()
	newPos := p.streamer.Position() + p.sampleRate.N(time.Millisecond*200)
	if newPos >= p.streamer.Len() {
		newPos = p.streamer.Len() - 1
	}
	if err := p.streamer.Seek(newPos); err != nil {
		report(err)
	}
	speaker.Unlock()
}

// Rewind rewind music
func (p *Player) Rewind() {
	speaker.Lock()
	newPos := p.streamer.Position() - p.sampleRate.N(time.Millisecond*200)
	if newPos < 0 {
		newPos = 0
	}
	if err := p.streamer.Seek(newPos); err != nil {
		report(err)
	}
	speaker.Unlock()
}

// Cue set and return cue point
func (p *Player) Cue() {
	if p.ctrl.Paused {
		speaker.Lock()
		p.cuePoint = p.streamer.Position()
		speaker.Unlock()
	} else {
		speaker.Lock()
		p.streamer.Seek(p.cuePoint)
		speaker.Unlock()
	}
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

// SetVol set volume
func (p *Player) SetVol(volume float64) {
	speaker.Lock()
	p.volume.Volume = volume
	speaker.Unlock()
}

// GainUp increase gain of music
func (p *Player) GainUp() {
	speaker.Lock()
	p.gain.Gain += 0.1
	speaker.Unlock()
}

// GainDown increase gain of music
func (p *Player) GainDown() {
	speaker.Lock()
	p.gain.Gain -= 0.1
	speaker.Unlock()
}

// SetGain set gain of music
func (p *Player) SetGain(gain float64) {
	speaker.Lock()
	p.gain.Gain = gain
	speaker.Unlock()
}

// Spdup increase speed controll
func (p *Player) Spdup() {
	speaker.Lock()
	p.resampler.SetRatio(p.resampler.Ratio() * 16 / 15)
	speaker.Unlock()
}

// Spddown decrease volume controll
func (p *Player) Spddown() {
	speaker.Lock()
	p.resampler.SetRatio(p.resampler.Ratio() * 15 / 16)
	speaker.Unlock()
}

// SetSpd set speed
func (p *Player) SetSpd(speed float64) {
	speaker.Lock()
	p.resampler.SetRatio(speed)
	speaker.Unlock()
}

// Cutoffup is
func (p *Player) Cutoffup() {
	speaker.Lock()
	p.filter.Freq += 100
	speaker.Unlock()
}

// Cutoffdown is
func (p *Player) Cutoffdown() {
	speaker.Lock()
	p.filter.Freq -= 100
	speaker.Unlock()
}
