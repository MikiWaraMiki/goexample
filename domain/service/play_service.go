package domain_service

import (
	"errors"

	. "github.com/MikiWaraMiki/goexample/domain/model"
)

// TODO: infra層へ移動
type PlayService struct {
	plays []Play
}

func (ps PlayService) fetchByPlayId(playId string) (*Play, error) {
	for i := range ps.plays {
		if ps.plays[i].PlayID == playId {
			return &ps.plays[i], nil
		}
	}
	return nil, errors.New("not found")
}
