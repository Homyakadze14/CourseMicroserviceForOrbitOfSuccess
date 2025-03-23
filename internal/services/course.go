package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/internal/entities"
)

var (
	ErrCourseAlreadyExists = errors.New("course with this data already exists")
	ErrThemeAlreadyExists  = errors.New("theme with this data already exists")
	ErrLessonAlreadyExists = errors.New("lesson with this data already exists")
)

type CourseService struct {
	log     *slog.Logger
	crsRepo CourseRepo
}

type CourseRepo interface {
	Create(ctx context.Context, course *entities.Course) (id int, err error)
	CreateTheme(ctx context.Context, theme *entities.Theme) (id int, err error)
	CreateLesson(ctx context.Context, lesson *entities.Lesson) (id int, err error)
}

func NewCourseService(
	log *slog.Logger,
	crsRepo CourseRepo,
) *CourseService {
	return &CourseService{
		log:     log,
		crsRepo: crsRepo,
	}
}

func (s *CourseService) Create(ctx context.Context, obj *entities.Course) (int, error) {
	const op = "Course.Create"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to create course")
	id, err := s.crsRepo.Create(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully created course")

	return id, err
}

func (s *CourseService) CreateTheme(ctx context.Context, obj *entities.Theme) (int, error) {
	const op = "Course.CreateTheme"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to create theme")
	id, err := s.crsRepo.CreateTheme(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully created theme")

	return id, err
}

func (s *CourseService) CreateLesson(ctx context.Context, obj *entities.Lesson) (int, error) {
	const op = "Course.CreateLesson"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to create lesson")
	id, err := s.crsRepo.CreateLesson(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully created lesson")

	return id, err
}
