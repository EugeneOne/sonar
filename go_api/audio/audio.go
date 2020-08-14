package audio

type Audio interface {
	Play() error
	Pause() error
	Stop() error
	Seek(position int) error
	Position() (int, error)
	Info() (*AudioInfo, error)
	SetVolume(volume int) error
	Close() error
}

type AudioInfo struct {
	FilePath string
}
