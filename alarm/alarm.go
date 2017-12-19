package alarm

import (
	"github.com/faiface/beep/mp3"
	"os"
	"time"
)

type Alarm struct {
}

func (a Alarm) openFile(path) os.File {
	file, err := os.Open("sounds/fire.mp3")
	return file
}

func (a Alarm) Ring(t time.Duration) {

}
