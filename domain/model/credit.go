package domain_model

import "math"

type Credit struct {
	volume int
}

func NewCredit(performance *Performance, play *Play) *Credit {
	volumeCredits := int(math.Max(float64(performance.Audience-30), 0))

	if play.IsComedy() {
		volumeCredits += int(math.Trunc(float64(performance.Audience) / 5))
	}

	return &Credit{
		volume: volumeCredits,
	}
}

func (this Credit) Volume() int {
	return this.volume
}
