package mouse

import (
	"time"

	"github.com/notaud/gwintils/mouse"
)

type Point struct {
	X, Y int32
}

func Position() (*Point, error) {
	pos, err := mouse.GetPosition()
	if err != nil {
		return nil, err
	}

	return &Point{pos.X, pos.Y}, nil
}

func Move(targetX, targetY int32, durationOpt ...int) error {
	duration := 0
	if len(durationOpt) > 0 && durationOpt[0] > 0 {
		duration = durationOpt[0]
	}

	// If speed is 0, move instantly
	if duration <= 0 {
		return mouse.Move(targetX, targetY)
	}

	startTime := time.Now()
	duration = int(time.Duration(duration) * time.Millisecond)
	endTime := startTime.Add(time.Duration(duration))

	mousePos, err := mouse.GetPosition()
	if err != nil {
		return err
	}

	for time.Now().Before(endTime) {
		currentTime := time.Now()
		elapsedRatio := float64(currentTime.Sub(startTime)) / float64(duration)

		// Ensure we don't exceed 1.0 for the ratio
		if elapsedRatio > 1.0 {
			elapsedRatio = 1.0
		}

		newX := int32(float64(mousePos.X) + float64(targetX-mousePos.X)*elapsedRatio)
		newY := int32(float64(mousePos.Y) + float64(targetY-mousePos.Y)*elapsedRatio)

		err := mouse.Move(newX, newY)
		if err != nil {
			return err
		}

		time.Sleep(time.Millisecond * 10)
	}

	return nil
}
