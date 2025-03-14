package station

import (
	"encoding/json"
	"github.com/mfajrihusaini02/mrt-schedules/common/client"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (response []Response, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []Response, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	for _, item := range stations {
		response = append(response, Response{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}
