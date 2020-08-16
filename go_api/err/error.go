package err

type Error struct {
	ErrCode int
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func New(code int, msg string) error {
	return &Error{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

const (
	Success = iota
	FileNotExists
	FilePermissionDenied
	FilePathIsDir
	FileDecodeFailed
	FileTypeUnsupported
	SpeakerInitFailed
	FileExistsInPool

	UnknownErr = -1
)
