package controller

import (
	"context"

	"github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/internal/entities"
	coursev1 "github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/proto/gen/course"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	GetAllCourses(ctx context.Context) ([]*entities.Course, error)
	GetCourse(ctx context.Context, id int) (*entities.Course, error)
	GetThemes(ctx context.Context, cid int) ([]*entities.Theme, error)
	GetLessons(ctx context.Context, cid, tid int) ([]*entities.Lesson, error)
	DeleteCourse(ctx context.Context, cid int) error
	UpdateCourse(ctx context.Context, obj *entities.Course) (int, error)
	UpdateTheme(ctx context.Context, obj *entities.Theme) (int, error)
	UpdateLesson(ctx context.Context, obj *entities.Lesson) (int, error)
}

func Register(gRPCServer *grpc.Server, course Course) {
	coursev1.RegisterCourseServiceServer(gRPCServer, &serverAPI{course: course})
}

func toCourseEntitie(obj *coursev1.CreateRequest) *entities.Course {
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

func toThemeEntitie(obj *coursev1.CreateTheme) *entities.Theme {
	return &entities.Theme{
		Title: obj.Title,
	}
}

func toLessonEntitie(obj *coursev1.CreateLesson) *entities.Lesson {
	return &entities.Lesson{
		Title:    obj.Title,
		Type:     obj.Type,
		Duration: obj.Duration,
		Content:  obj.Content,
		Task:     obj.Task,
	}
}

func toCourseDTO(obj *entities.Course) *coursev1.Course {
	return &coursev1.Course{
		Id:              int32(obj.ID),
		Title:           obj.Title,
		Description:     obj.Description,
		FullDescription: obj.FullDescription,
		Work:            obj.Work,
		Difficulty:      obj.Difficulty,
		Duration:        obj.Duration,
		Image:           obj.Image,
	}
}

func (s *serverAPI) Create(
	ctx context.Context,
	in *coursev1.CreateRequest,
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

func (s *serverAPI) GetAll(
	ctx context.Context,
	in *emptypb.Empty,
) (*coursev1.GetResponse, error) {
	courses, err := s.course.GetAllCourses(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	response := make([]*coursev1.Course, len(courses))
	for i, course := range courses {
		response[i] = toCourseDTO(course)
	}

	return &coursev1.GetResponse{
		Courses: response,
	}, nil
}

func toThemeDTO(obj *entities.Theme, les []*coursev1.Lesson) *coursev1.Theme {
	return &coursev1.Theme{
		Id:      int32(obj.ID),
		Title:   obj.Title,
		Lessons: les,
	}
}

func toLessonDTO(obj *entities.Lesson) *coursev1.Lesson {
	return &coursev1.Lesson{
		Id:       int32(obj.ID),
		Title:    obj.Title,
		Type:     obj.Type,
		Duration: obj.Duration,
		Content:  obj.Content,
		Task:     obj.Task,
	}
}

func (s *serverAPI) Get(
	ctx context.Context,
	in *coursev1.GetCourseRequest,
) (*coursev1.GetCourseResponse, error) {
	course, err := s.course.GetCourse(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	themes, err := s.course.GetThemes(ctx, course.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	themesResp := make([]*coursev1.Theme, len(themes))
	for i, theme := range themes {
		lessons, err := s.course.GetLessons(ctx, course.ID, theme.ID)
		if err != nil {
			return nil, status.Error(codes.Internal, ErrInternalServerError)
		}
		lsResp := make([]*coursev1.Lesson, len(lessons))
		for j, lesson := range lessons {
			lsResp[j] = toLessonDTO(lesson)
		}
		themesResp[i] = toThemeDTO(theme, lsResp)
	}

	return &coursev1.GetCourseResponse{
		Id:              int32(course.ID),
		Title:           course.Title,
		Description:     course.Description,
		FullDescription: course.FullDescription,
		Work:            course.Work,
		Difficulty:      course.Difficulty,
		Duration:        course.Duration,
		Image:           course.Image,
		Themes:          themesResp,
	}, nil
}

func (s *serverAPI) Delete(
	ctx context.Context,
	in *coursev1.DeleteCourseRequest,
) (*coursev1.SuccessResponse, error) {
	err := s.course.DeleteCourse(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	return &coursev1.SuccessResponse{
		Success: true,
	}, nil
}

func toCourseEntitieUpd(obj *coursev1.UpdateCourseRequest) *entities.Course {
	return &entities.Course{
		ID:              int(obj.Id),
		Title:           obj.Title,
		Description:     obj.Description,
		FullDescription: obj.FullDescription,
		Work:            obj.Work,
		Difficulty:      obj.Difficulty,
		Duration:        obj.Duration,
		Image:           obj.Image,
	}
}

func toThemeEntitieUpd(obj *coursev1.UpdateTheme) *entities.Theme {
	theme := &entities.Theme{
		Title: obj.Title,
	}

	if obj.Id != nil {
		theme.ID = int(*obj.Id)
	}

	return theme
}

func toLessonEntitieUpd(obj *coursev1.UpdateLesson) *entities.Lesson {
	lesson := &entities.Lesson{
		Title:    obj.Title,
		Type:     obj.Type,
		Duration: obj.Duration,
		Content:  obj.Content,
		Task:     obj.Task,
	}

	if obj.Id != nil {
		lesson.ID = int(*obj.Id)
	}

	return lesson
}

func (s *serverAPI) Update(
	ctx context.Context,
	in *coursev1.UpdateCourseRequest,
) (*coursev1.SuccessResponse, error) {
	course := toCourseEntitieUpd(in)
	id, err := s.course.UpdateCourse(ctx, course)
	if err != nil {
		return nil, status.Error(codes.Internal, ErrInternalServerError)
	}

	for _, th := range in.Themes {
		theme := toThemeEntitieUpd(th)
		theme.CourseID = id
		var themeID int
		var err error
		if th.Id == nil {
			themeID, err = s.course.CreateTheme(ctx, theme)
		} else {
			themeID, err = s.course.UpdateTheme(ctx, theme)
		}
		if err != nil {
			return nil, status.Error(codes.Internal, ErrInternalServerError)
		}
		for _, ls := range th.Lessons {
			lesson := toLessonEntitieUpd(ls)
			lesson.CourseID = id
			lesson.ThemeID = themeID
			if ls.Id == nil {
				_, err = s.course.CreateLesson(ctx, lesson)
			} else {
				_, err = s.course.UpdateLesson(ctx, lesson)
			}
			if err != nil {
				return nil, status.Error(codes.Internal, ErrInternalServerError)
			}
		}
	}

	return &coursev1.SuccessResponse{
		Success: true,
	}, nil
}
