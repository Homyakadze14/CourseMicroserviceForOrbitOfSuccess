package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/internal/services"
	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/pkg/postgres"
)

type CourseRepository struct {
	*postgres.Postgres
}

func NewCourseRepository(pg *postgres.Postgres) *CourseRepository {
	return &CourseRepository{pg}
}

func (r *CourseRepository) Create(ctx context.Context, obj *entities.Course) (id int, err error) {
	const op = "repositories.CourseRepository.Create"

	row := r.Pool.QueryRow(
		ctx,
		"INSERT INTO course(title, description, full_descritpion, work, difficulty, duration, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		obj.Title, obj.Description, obj.FullDescription, obj.Work, obj.Difficulty, obj.Duration, obj.Image)

	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return -1, services.ErrCourseAlreadyExists
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *CourseRepository) CreateTheme(ctx context.Context, obj *entities.Theme) (id int, err error) {
	const op = "repositories.CourseRepository.CreateTheme"

	row := r.Pool.QueryRow(
		ctx,
		"INSERT INTO theme(course_id, title) VALUES ($1, $2) RETURNING id",
		obj.CourseID, obj.Title)

	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return -1, services.ErrCourseAlreadyExists
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *CourseRepository) CreateLesson(ctx context.Context, obj *entities.Lesson) (id int, err error) {
	const op = "repositories.CourseRepository.CreateLesson"

	row := r.Pool.QueryRow(
		ctx,
		"INSERT INTO lesson(course_id, theme_id, title, type, duration, content, task) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		obj.CourseID, obj.ThemeID, obj.Title, obj.Type, obj.Duration, obj.Content, obj.Task)

	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return -1, services.ErrCourseAlreadyExists
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]*entities.Course, error) {
	const op = "repositories.CourseRepository.GetAllCourses"
	arraySize := 20
	rows, err := r.Pool.Query(ctx, "SELECT * FROM course")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	courses := make([]*entities.Course, 0, arraySize)
	for rows.Next() {
		var obj entities.Course
		err := rows.Scan(&obj.ID, &obj.Title, &obj.Description, &obj.FullDescription,
			&obj.Work, &obj.Difficulty, &obj.Duration, &obj.Image)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		courses = append(courses, &obj)
	}

	return courses, nil
}
