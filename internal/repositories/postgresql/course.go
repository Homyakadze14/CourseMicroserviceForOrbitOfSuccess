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

func (r *CourseRepository) GetCourse(ctx context.Context, id int) (*entities.Course, error) {
	const op = "repositories.CourseRepository.GetCourse"

	row := r.Pool.QueryRow(
		ctx,
		"SELECT id, title, description, full_descritpion, work, difficulty, duration, image FROM course WHERE id=$1",
		id)

	var obj entities.Course
	err := row.Scan(&obj.ID, &obj.Title, &obj.Description, &obj.FullDescription,
		&obj.Work, &obj.Difficulty, &obj.Duration, &obj.Image)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &obj, nil
}

func (r *CourseRepository) GetThemes(ctx context.Context, cid int) ([]*entities.Theme, error) {
	const op = "repositories.CourseRepository.GetThemes"
	arraySize := 20
	rows, err := r.Pool.Query(ctx, "SELECT * FROM theme WHERE course_id=$1", cid)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	themes := make([]*entities.Theme, 0, arraySize)
	for rows.Next() {
		var obj entities.Theme
		err := rows.Scan(&obj.ID, &obj.CourseID, &obj.Title)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		themes = append(themes, &obj)
	}

	return themes, nil
}

func (r *CourseRepository) GetLessons(ctx context.Context, cid, tid int) ([]*entities.Lesson, error) {
	const op = "repositories.CourseRepository.GetLessons"
	arraySize := 20
	rows, err := r.Pool.Query(ctx, "SELECT * FROM lesson WHERE course_id=$1 AND theme_id=$2", cid, tid)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	lessons := make([]*entities.Lesson, 0, arraySize)
	for rows.Next() {
		var obj entities.Lesson
		err := rows.Scan(&obj.ID, &obj.CourseID, &obj.ThemeID, &obj.Title,
			&obj.Type, &obj.Duration, &obj.Content, &obj.Task)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		lessons = append(lessons, &obj)
	}

	return lessons, nil
}

func (r *CourseRepository) DeleteCourse(ctx context.Context, id int) (err error) {
	const op = "repositories.CourseRepository.DeleteCourse"

	_, err = r.Pool.Exec(ctx,
		"DELETE FROM course WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *CourseRepository) UpdateCourse(ctx context.Context, obj *entities.Course) (err error) {
	const op = "repositories.CourseRepository.UpdateCourse"

	_, err = r.Pool.Exec(
		ctx,
		(`UPDATE course SET title=$1 AND description=$2 AND full_descritpion=$3 AND work=$4
		AND difficulty=$5 AND duration=$6 AND image=$7`),
		obj.Title, obj.Description, obj.FullDescription, obj.Work, obj.Difficulty, obj.Duration, obj.Image)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *CourseRepository) UpdateTheme(ctx context.Context, obj *entities.Theme) (err error) {
	const op = "repositories.CourseRepository.UpdateTheme"

	_, err = r.Pool.Exec(
		ctx,
		"INSERT INTO theme(course_id, title) VALUES ($1, $2) ON CONFLICT(id) DO UPDATE SET coures_id=$1 AND title=$2",
		obj.CourseID, obj.Title)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *CourseRepository) UpdateLesson(ctx context.Context, obj *entities.Lesson) (err error) {
	const op = "repositories.CourseRepository.UpdateLesson"

	_, err = r.Pool.Exec(
		ctx,
		(`INSERT INTO lesson(course_id, theme_id, title, type, duration, content, task) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT(id) DO UPDATE SET course_id=$1 AND theme_id=$2 AND title=$3 AND type=$4 AND duration=$5 AND content=$6 AND task=$7`),
		obj.CourseID, obj.ThemeID, obj.Title, obj.Type, obj.Duration, obj.Content, obj.Task)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
