package audio

import (
	"go_api/err"
	"go_api/file"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

type BeepAudio struct {
	filePath        string
	format          beep.Format
	WrappedStreamer beep.Streamer
	Streamer        beep.StreamSeekCloser
	ctrl            *beep.Ctrl
	done            chan struct{}
}

func newBeepAudio(filePath string) (Audio, error) {
	fp, e := file.Open(filePath)
	if e != nil {
		return nil, e
	}

	var (
		streamer beep.StreamSeekCloser
		format   beep.Format
	)

	switch file.GetExt(filePath) {
	case fileTypeMp3:
		streamer, format, e = mp3.Decode(fp)
	case fileTypeWav:
		streamer, format, e = wav.Decode(fp)
	case fileTypeFlac:
		streamer, format, e = flac.Decode(fp)
	case fileTypeOgg:
		streamer, format, e = vorbis.Decode(fp)

	default:
		fp.Close()
		return nil, err.New(err.FileTypeUnsupported, "unsupported file type")
	}

	audio := &BeepAudio{
		filePath: filePath,
		format:   format,
		Streamer: streamer,
	}

	audio.ctrl = &beep.Ctrl{
		Streamer: audio.Streamer,
		Paused:   false,
	}

	audio.WrappedStreamer = audio.ctrl

	return audio, nil
}

func (audio *BeepAudio) Play() error {
	speaker.Clear()
	if e := speaker.Init(audio.format.SampleRate, audio.format.SampleRate.N(time.Second/4)); nil != e {
		return err.New(err.SpeakerInitFailed, e.Error())
	}

	audio.done = make(chan struct{}, 1)
	speaker.Play(beep.Seq(audio.WrappedStreamer, beep.Callback(func() {
		audio.done <- struct{}{}
	})))
	return nil
}

func (audio *BeepAudio) Pause() error {
	speaker.Lock()
	defer speaker.Unlock()

	audio.ctrl.Paused = true
	return nil
}

func (audio *BeepAudio) Resume() error {
	speaker.Lock()
	defer speaker.Unlock()

	audio.ctrl.Paused = false
	return nil
}

func (audio *BeepAudio) Stop() error {
	audio.Pause()
	audio.Seek(0)
	return nil
}

func (audio *BeepAudio) Done() {
	<-audio.done
}

func (audio *BeepAudio) Seek(position int) error {
	speaker.Lock()
	defer speaker.Unlock()

	audio.Streamer.Seek(audio.Streamer.Len() / 100 * position)
	return nil
}

func (audio *BeepAudio) Position() (int, error) {
	speaker.Lock()
	defer speaker.Unlock()

	pos := float64(audio.Streamer.Position()) / float64(audio.Streamer.Len()) * 100
	return int(pos), nil
}

func (audio *BeepAudio) SetVolume(volume int) error {

	return nil
}

func (audio *BeepAudio) Close() error {
	speaker.Clear()
	audio.Stop()

	if closer, ok := audio.WrappedStreamer.(beep.StreamCloser); ok {
		return closer.Close()
	}

	return nil
}
