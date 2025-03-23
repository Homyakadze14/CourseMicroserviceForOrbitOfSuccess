package controller

import (
	"context"

	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/internal/entities"
	coursev1 "github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/proto/gen/course"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternalServerError = "internal server error"
)

type serverAPI struct {
	coursev1.UnimplementedCourseServiceServer
	course Course
}

type Course interface {
	Create(ctx context.Context, obj *entities.Course) (int, error)
	CreateTheme(ctx context.Context, obj *entities.Theme) (int, error)
	CreateLesson(ctx context.Context, obj *entities.Lesson) (int, error)
}

func Register(gRPCServer *grpc.Server, course Course) {
	coursev1.RegisterCourseServiceServer(gRPCServer, &serverAPI{course: course})
}

func toCourseEntitie(obj *coursev1.Course) *entities.Course {
	return &entities.Course{
		Title:           obj.Title,
		Description:     obj.Description,
		FullDescription: obj.FullDescription,
		Work:            obj.Work,
		Difficulty:      obj.Difficulty,
		Duration:        obj.Duration,
		Image:           obj.Image,
	}
}

func toThemeEntitie(obj *coursev1.Theme) *entities.Theme {
	return &entities.Theme{
		Title: obj.Title,
	}
}

func toLessonEntitie(obj *coursev1.Lesson) *entities.Lesson {
	return &entities.Lesson{
		Title:        obj.Title,
		Type:         obj.Type,
		Duration:     obj.Duration,
		Completed:    obj.Completed,
		Content:      obj.Content,
		PracticeType: obj.PracticeType,
		Task:         obj.Task,
	}
}

func (s *serverAPI) Create(
	ctx context.Context,
	in *coursev1.Course,
) (*coursev1.SuccessResponse, error) {

	course := toCourseEntitie(in)
	id, err := s.course.Create(ctx, course)
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	for _, th := range in.Themes {
		theme := toThemeEntitie(th)
		theme.CourseID = id
		themeID, err := s.course.CreateTheme(ctx, theme)
		if err != nil {
			return nil, status.Error(codes.Internal, ErrInternalServerError)
		}
		for _, ls := range th.Lessons {
			lesson := toLessonEntitie(ls)
			lesson.CourseID = id
			lesson.ThemeID = themeID
			_, err := s.course.CreateLesson(ctx, lesson)
			if err != nil {
				return nil, status.Error(codes.Internal, ErrInternalServerError)
			}
		}
	}

	return &coursev1.SuccessResponse{
		Success: true,
	}, nil
}
