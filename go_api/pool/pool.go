package pool

import (
	"go_api/audio"
	"go_api/err"
	"sync"
)

var lock *sync.RWMutex
var pool map[string]audio.Audio

func init() {
	lock = &sync.RWMutex{}
	pool = make(map[string]audio.Audio)
}

func Put(hash string, audio audio.Audio) (string, error) {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := pool[hash]; ok {
		return "", err.New(err.FileExistsInPool, "file already exists in pool")
	}

	pool[hash] = audio
	return hash, nil
}

func Get(hash string) audio.Audio {
	lock.RLock()
	defer lock.RUnlock()

	if audio, ok := pool[hash]; ok {
		return audio
	}

	return nil
}

func Del(hash string) {
	lock.Lock()
	defer lock.Unlock()

	delete(pool, hash)
}

func Pop(hash string) audio.Audio {
	lock.Lock()
	defer lock.Unlock()

	if audio, ok := pool[hash]; ok {
		delete(pool, hash)
		return audio
	}

	return nil
}
