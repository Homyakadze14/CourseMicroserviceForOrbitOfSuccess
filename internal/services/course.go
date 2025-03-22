package services

import "log/slog"

type CourseService struct {
	log *slog.Logger
}

func NewCourseService(
	log *slog.Logger,
) *CourseService {
	return &CourseService{
		log: log,
	}
}
