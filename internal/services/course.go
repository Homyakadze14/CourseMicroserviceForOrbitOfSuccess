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
	GetAllCourses(ctx context.Context) ([]*entities.Course, error)
	GetCourse(ctx context.Context, id int) (*entities.Course, error)
	GetThemes(ctx context.Context, cid int) ([]*entities.Theme, error)
	GetLessons(ctx context.Context, cid, tid int) ([]*entities.Lesson, error)
	DeleteCourse(ctx context.Context, id int) (err error)
	UpdateCourse(ctx context.Context, obj *entities.Course) (id int, err error)
	UpdateTheme(ctx context.Context, obj *entities.Theme) (id int, err error)
	UpdateLesson(ctx context.Context, obj *entities.Lesson) (id int, err error)
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

func (s *CourseService) GetAllCourses(ctx context.Context) ([]*entities.Course, error) {
	const op = "Course.GetAllCourses"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to get courses")
	courses, err := s.crsRepo.GetAllCourses(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("courses successfully geted")

	return courses, err
}

func (s *CourseService) GetCourse(ctx context.Context, id int) (*entities.Course, error) {
	const op = "Course.GetCourse"

	log := s.log.With(
		slog.String("op", op),
		slog.Int("id", id),
	)

	log.Info("trying to get course")
	course, err := s.crsRepo.GetCourse(ctx, id)
	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("course successfully geted")

	return course, err
}

func (s *CourseService) GetThemes(ctx context.Context, cid int) ([]*entities.Theme, error) {
	const op = "Course.GetThemes"

	log := s.log.With(
		slog.String("op", op),
		slog.Int("cid", cid),
	)

	log.Info("trying to get themes")
	themes, err := s.crsRepo.GetThemes(ctx, cid)
	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("themes successfully geted")

	return themes, err
}

func (s *CourseService) GetLessons(ctx context.Context, cid, tid int) ([]*entities.Lesson, error) {
	const op = "Course.GetLessons"

	log := s.log.With(
		slog.String("op", op),
		slog.Int("cid", cid),
	)

	log.Info("trying to get lessons")
	lessons, err := s.crsRepo.GetLessons(ctx, cid, tid)
	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("lessons successfully geted")

	return lessons, err
}

func (s *CourseService) DeleteCourse(ctx context.Context, cid int) error {
	const op = "Course.DeleteCourse"

	log := s.log.With(
		slog.String("op", op),
		slog.Int("cid", cid),
	)

	log.Info("trying to delete course")
	err := s.crsRepo.DeleteCourse(ctx, cid)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("course successfully deleted")

	return nil
}

func (s *CourseService) UpdateCourse(ctx context.Context, obj *entities.Course) (int, error) {
	const op = "Course.UpdateCourse"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to update course")
	id, err := s.crsRepo.UpdateCourse(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully updated course")

	return id, err
}

func (s *CourseService) UpdateTheme(ctx context.Context, obj *entities.Theme) (int, error) {
	const op = "Course.UpdateTheme"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to update theme")
	id, err := s.crsRepo.UpdateTheme(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully created theme")

	return id, err
}

func (s *CourseService) UpdateLesson(ctx context.Context, obj *entities.Lesson) (int, error) {
	const op = "Course.UpdateLesson"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("trying to update lesson")
	id, err := s.crsRepo.UpdateLesson(ctx, obj)
	if err != nil {
		log.Error(err.Error())
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully updated lesson")

	return id, err
}
