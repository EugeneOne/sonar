package audio

import (
	"go_api/err"
	"go_api/file"
)

type Audio interface {
	Play() error
	Pause() error
	Resume() error
	Stop() error
	Done()
	Seek(position int) error
	Position() (int, error)
	SetVolume(volume int) error
	Close() error
}

const (
	fileTypeMp3  = "mp3"
	fileTypeWav  = "wav"
	fileTypeFlac = "flac"
	fileTypeOgg  = "ogg"
)

func NewAudio(filePath string) (Audio, error) {
	absPath, e := file.GetAbsolutePath(filePath)
	if nil != e {
		return nil, e
	}

	switch file.GetExt(absPath) {
	case fileTypeMp3, fileTypeWav, fileTypeFlac, fileTypeOgg:
		return newBeepAudio(filePath)

	default:
		return nil, err.New(err.FileTypeUnsupported, "unsupported file type")
	}
}
