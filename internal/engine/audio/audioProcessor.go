package audio

import (
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func LoadAudio() {
	f, err := os.Open("res/audio/Echoes.flac")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := flac.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
}
