package controller

import (
	coursev1 "github.com/Homyakadze14/CourseMicroserviceForOrbitOfSuccess/proto/gen/course"
	"google.golang.org/grpc"
)

type serverAPI struct {
	coursev1.UnimplementedCourseServiceServer
	course Course
}

type Course interface {
}

func Register(gRPCServer *grpc.Server, course Course) {
	coursev1.RegisterCourseServiceServer(gRPCServer, &serverAPI{course: course})
}
