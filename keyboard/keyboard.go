package keyboard

import (
	"time"

	"github.com/notaud/gwintils/keyboard"
)

func ListenKey(key string) error {
	vk, err := keyboard.VirtualKey(key)
	if err != nil {
		return err
	}

	for {
		isPressed := keyboard.GetAsyncKeyState(vk)
		if isPressed {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	return nil
}
