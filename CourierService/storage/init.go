package storage

import (
	pb "learning-service/genproto/learning"
)

type InitRoot interface {
	Learning() Learning
}

type Learning interface {
	CreateLearningTopic(request *pb.CreateLearningTopicRequest) (*pb.CreateLearningTopicResponse, error)
	GetLearningTopics(request *pb.GetLearningTopicsRequest) (*pb.GetLearningTopicsResponse, error)
	UpdateLearningTopic(request *pb.UpdateLearningTopicRequest) (*pb.UpdateLearningTopicResponse, error)
	DeleteLearningTopic(request *pb.DeleteLearningTopicRequest) (*pb.DeleteLearningTopicResponse, error)

	CompletedTopics(request *pb.CompletedTopicsRequest) (*pb.CompletedTopicsResponse, error)
	GetCompletedTopics(request *pb.GetCompletedTopicsRequest) (*pb.GetCompletedTopicsResponse, error)

	CreateQuiz(request *pb.CreateQuizRequest) (*pb.CreateQuizResponse, error)
	GetQuiz(request *pb.GetQuizRequest) (*pb.GetQuizResponse, error)
	UpdateQuiz(request *pb.UpdateQuizRequest) (*pb.UpdateQuizResponse, error)
	DeleteQuiz(request *pb.DeleteQuizRequest) (*pb.DeleteQuizResponse, error)

	SubmitQuiz(request *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error)

	CreateExtraResourses(request *pb.CreateExtraResoursesRequest) (*pb.CreateExtraResoursesResponse, error)
	GetExtraResourses(request *pb.GetExtraResourcesRequest) (*pb.GetExtraResourcesResponse, error)
	UpdateExtraResourses(request *pb.UpdateExtraResoursesRequest) (*pb.UpdateExtraResoursesResponse, error)
	DeleteExtraResourses(request *pb.DeleteExtraResoursesRequest) (*pb.DeleteExtraResoursesResponse, error)

	CompletedExtraResources(request *pb.CompletedExtraResourcesRequest) (*pb.CompletedExtraResourcesResponse, error)

	GetLearningProgress(request *pb.GetLearningProgressRequest) (*pb.GetLearningProgressResponse, error)

	CreateLearningRecommendations(request *pb.CreateLearningRecommendationsRequest) (*pb.CreateLearningRecommendationsResponse, error)
	GetLearningRecommendations(request *pb.GetLearningRecommendationsRequest) (*pb.GetLearningRecommendationsResponse, error)

	CreateLearningFeedback(request *pb.CreateLearningFeedbackRequest) (*pb.CreateLearningFeedbackResponse, error)
	GetLearningFeedback(request *pb.GetLearningFeedbackRequest) (*pb.GetLearningFeedbackResponse, error)

	CreateLearningHomeworks(request *pb.CreateLearningHomeworksRequest) (*pb.CreateLearningHomeworksResponse, error)
	GetLearningHomeworks(request *pb.GetLearningHomeworksRequest) (*pb.GetLearningHomeworksResponse, error)

	SubmitHomework(request *pb.SubmitHomeworkRequest) (*pb.SubmitHomeworkResponse, error)

}