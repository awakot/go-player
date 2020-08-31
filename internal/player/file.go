package player

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

func (p *Player) DecodeMusicFile(path string) (beep.StreamSeekCloser, beep.Format, error) {
	musics := listMusic(path)
	p.musicTitle = musics[0]
	f, err := os.Open(path + "/" + p.musicTitle)
	if err != nil {
		fmt.Println(err)
	}
	return mp3.Decode(f)
}

func listMusic(path string) []string {
	cd, _ := os.Getwd()
	fileinfos, _ := ioutil.ReadDir(cd + "/" + path)
	var list []string
	r := regexp.MustCompile(`.*mp3`)
	for _, fileinfo := range fileinfos {
		if !r.MatchString(fileinfo.Name()) {
			continue
		}
		list = append(list, fileinfo.Name())
	}
	return list
}
