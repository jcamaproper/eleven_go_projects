package service

import "projects/crud_app_gingonic/model"

type VideoService interface {
	Save(model.Video) model.Video
	FindAll() []model.Video
}

type videoService struct {
	videos []model.Video
}

func New() VideoService {
	return &videoService{
		videos: []model.Video{},
	}
}

func (service *videoService) Save(video model.Video) model.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []model.Video {
	return service.videos
}
