package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

// TODO: infra層へ移動
type PlayService struct {
	plays []Play
}

func NewPlayService(plays []Play) *PlayService {
	return &PlayService{
		plays: plays,
	}
}
func (ps PlayService) FetchByPlayId(playId string) (*Play, error) {
	for i := range ps.plays {
		if ps.plays[i].PlayID == playId {
			return &ps.plays[i], nil
		}
	}
	return nil, errors.New("not found")
}
