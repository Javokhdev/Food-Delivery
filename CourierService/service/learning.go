package service

import (
	"context"
	pb "learning-service/genproto/learning"
	s "learning-service/storage"
)

type LearningService struct {
	pb.UnimplementedLearningServiceServer
	stg s.InitRoot
}

func NewLearningService(stg s.InitRoot) *LearningService {
	return &LearningService{stg: stg}
}

func (s *LearningService) CreateLearningTopic(ctx context.Context, req *pb.CreateLearningTopicRequest) (*pb.CreateLearningTopicResponse, error) {
	res, err := s.stg.Learning().CreateLearningTopic(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetLearningTopics(ctx context.Context, req *pb.GetLearningTopicsRequest) (*pb.GetLearningTopicsResponse, error) {
	res, err := s.stg.Learning().GetLearningTopics(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) UpdateLearningTopic(ctx context.Context, req *pb.UpdateLearningTopicRequest) (*pb.UpdateLearningTopicResponse, error) {
	res, err := s.stg.Learning().UpdateLearningTopic(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) DeleteLearningTopic(ctx context.Context, req *pb.DeleteLearningTopicRequest) (*pb.DeleteLearningTopicResponse, error) {
	res, err := s.stg.Learning().DeleteLearningTopic(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CompletedTopics(ctx context.Context, req *pb.CompletedTopicsRequest) (*pb.CompletedTopicsResponse, error) {
	res, err := s.stg.Learning().CompletedTopics(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetCompletedTopics(ctx context.Context, req *pb.GetCompletedTopicsRequest) (*pb.GetCompletedTopicsResponse, error) {
	res, err := s.stg.Learning().GetCompletedTopics(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CreateQuiz(ctx context.Context, req *pb.CreateQuizRequest) (*pb.CreateQuizResponse, error) {
	res, err := s.stg.Learning().CreateQuiz(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetQuiz(ctx context.Context, req *pb.GetQuizRequest) (*pb.GetQuizResponse, error) {
	res, err := s.stg.Learning().GetQuiz(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) UpdateQuiz(ctx context.Context, req *pb.UpdateQuizRequest) (*pb.UpdateQuizResponse, error) {
	res, err := s.stg.Learning().UpdateQuiz(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) DeleteQuiz(ctx context.Context, req *pb.DeleteQuizRequest) (*pb.DeleteQuizResponse, error) {
	res, err := s.stg.Learning().DeleteQuiz(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) SubmitQuiz(ctx context.Context, req *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error) {
	res, err := s.stg.Learning().SubmitQuiz(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CreateExtraResourses(ctx context.Context, req *pb.CreateExtraResoursesRequest) (*pb.CreateExtraResoursesResponse, error) {
	res, err := s.stg.Learning().CreateExtraResourses(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetExtraResourses(ctx context.Context, req *pb.GetExtraResourcesRequest) (*pb.GetExtraResourcesResponse, error) {
	res, err := s.stg.Learning().GetExtraResourses(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) UpdateExtraResourses(ctx context.Context, req *pb.UpdateExtraResoursesRequest) (*pb.UpdateExtraResoursesResponse, error) {
	res, err := s.stg.Learning().UpdateExtraResourses(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) DeleteExtraResourses(ctx context.Context, req *pb.DeleteExtraResoursesRequest) (*pb.DeleteExtraResoursesResponse, error) {
	res, err := s.stg.Learning().DeleteExtraResourses(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CompletedExtraResources(ctx context.Context, req *pb.CompletedExtraResourcesRequest) (*pb.CompletedExtraResourcesResponse, error) {
	res, err := s.stg.Learning().CompletedExtraResources(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetLearningProgress(ctx context.Context, req *pb.GetLearningProgressRequest) (*pb.GetLearningProgressResponse, error) {
	res, err := s.stg.Learning().GetLearningProgress(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CreateLearningRecommendations(ctx context.Context, req *pb.CreateLearningRecommendationsRequest) (*pb.CreateLearningRecommendationsResponse, error) {
	res, err := s.stg.Learning().CreateLearningRecommendations(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetLearningRecommendations(ctx context.Context, req *pb.GetLearningRecommendationsRequest) (*pb.GetLearningRecommendationsResponse, error) {
	res, err := s.stg.Learning().GetLearningRecommendations(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CreateLearningFeedback(ctx context.Context, req *pb.CreateLearningFeedbackRequest) (*pb.CreateLearningFeedbackResponse, error) {
	res, err := s.stg.Learning().CreateLearningFeedback(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetLearningFeedback(ctx context.Context, req *pb.GetLearningFeedbackRequest) (*pb.GetLearningFeedbackResponse, error) {
	res, err := s.stg.Learning().GetLearningFeedback(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) CreateLearningHomeworks(ctx context.Context, req *pb.CreateLearningHomeworksRequest) (*pb.CreateLearningHomeworksResponse, error) {
	res, err := s.stg.Learning().CreateLearningHomeworks(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) GetLearningHomeworks(ctx context.Context, req *pb.GetLearningHomeworksRequest) (*pb.GetLearningHomeworksResponse, error) {
	res, err := s.stg.Learning().GetLearningHomeworks(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LearningService) SubmitHomework(ctx context.Context, req *pb.SubmitHomeworkRequest) (*pb.SubmitHomeworkResponse, error) {
	res, err := s.stg.Learning().SubmitHomework(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
