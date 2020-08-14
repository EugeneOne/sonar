package main

func main() {}

func Open(filePath string) (string, int) {
	return "", 0
}

func Play(fp string) int {
	return 0
}

func Pause(fp string) int {
	return 0
}

func Stop(fp string) int {
	return 0
}

func Seek(fp string, position int) int {
	return 0
}

func Position(fp string) int {
	return 0
}

func Info(fp string) {
	// todo: return a struct to describe the audio info?
}

func SetVolume(fp string, volume int) int {
	return 0
}

func Close(fp string) int {
	return 0
}
