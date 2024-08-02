package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	pb "learning-service/genproto/learning"

	"github.com/google/uuid"
)

type LearningStorage struct {
	db *sql.DB
}

func NewLearningStorage(db *sql.DB) *LearningStorage {
	return &LearningStorage{db: db}
}

func (c *LearningStorage) CreateLearningTopic(req *pb.CreateLearningTopicRequest) (*pb.CreateLearningTopicResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO topics(id, name, description, difficulty)
		VALUES($1, $2, $3, $4)`

	_, err := c.db.Exec(query, id, req.Name, req.Description, req.Difficulty)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateLearningTopicResponse{Id: id, Message: "success"}, nil
}

func (c *LearningStorage) GetLearningTopics(req *pb.GetLearningTopicsRequest) (*pb.GetLearningTopicsResponse, error) {
	var query string
	topics := []*pb.LearningTopic{}
	query = `
	SELECT id, name, description, difficulty FROM topics WHERE deleted_at = 0`
	var arr []interface{}
	var request pb.LearningTopic
	count := 1

	if len(request.Id) > 0 {
		query += fmt.Sprintf(" and id = $%d", count)
		count++
		arr = append(arr, request.Id)
	}
	if len(request.Name) > 0 {
		query += fmt.Sprintf(" and name = $%d", count)
		count++
		arr = append(arr, request.Name)
	}
	if len(request.Description) > 0 {
		query += fmt.Sprintf(" and description = $%d", count)
		count++
		arr = append(arr, request.Description)
	}
	if len(request.Difficulty) > 0 {
		query += fmt.Sprintf(" and phone_number = $%d", count)
		count++
		arr = append(arr, request.Difficulty)
	}

	rows, err := c.db.Query(query, arr...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&request.Id, &request.Name, &request.Description, &request.Difficulty)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		topics = append(topics, &request)
	}
	return &pb.GetLearningTopicsResponse{Topics: topics}, nil

}

func (c *LearningStorage) UpdateLearningTopic(req *pb.UpdateLearningTopicRequest) (*pb.UpdateLearningTopicResponse, error) {
	query := `
		UPDATE topics
		SET name = $1, description = $2, difficulty = $3
		WHERE id = $4`

	_, err := c.db.Exec(query, req.Name, req.Description, req.Difficulty, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateLearningTopicResponse{Message: "success"}, nil
}

func (c *LearningStorage) DeleteLearningTopic(req *pb.DeleteLearningTopicRequest) (*pb.DeleteLearningTopicResponse, error) {
	query := `
		UPDATE topics SET deleted_at = $1
		WHERE id = $2`

	_, err := c.db.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteLearningTopicResponse{Message: "success"}, nil
}


func (c *LearningStorage) CompletedTopics(req *pb.CompletedTopicsRequest) (*pb.CompletedTopicsResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO completed_topics (id, user_id, topic_id, xp_earned)
		VALUES ($1, $2, $3, $4)`

	_, err := c.db.Exec(query, id, req.UserId, req.TopicId, req.XpEarned)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		query := `UPDATE users SET xp = xp + $1 WHERE id = $2`
		_, err := c.db.Exec(query, req.XpEarned, req.UserId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &pb.CompletedTopicsResponse{Message: "success", XpEarned: req.XpEarned}, nil
	}

}

func (c *LearningStorage) GetCompletedTopics(req *pb.GetCompletedTopicsRequest) (*pb.GetCompletedTopicsResponse, error) {
	var query string
	topics := []*pb.CompletedTopics{}
	query = `SELECT id, user_id, topic_id, xp_earned FROM completed_topics`
	var request pb.CompletedTopics
	var arr []interface{}
	count := 1

	if len(request.Id) > 0 {
		query += fmt.Sprintf(" and id = $%d", count)
		count++
		arr = append(arr, request.Id)
	}
	if len(request.UserId) > 0 {
		query += fmt.Sprintf(" and user_id = $%d", count)
		count++
		arr = append(arr, request.UserId)
	}
	if len(request.TopicId) > 0 {
		query += fmt.Sprintf(" and topic_id = $%d", count)
		count++
		arr = append(arr, request.TopicId)
	}

	rows, err := c.db.Query(query, arr...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&request.Id, &request.UserId, &request.TopicId, &request.XpEarned)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		topics = append(topics, &request)
	}

	return &pb.GetCompletedTopicsResponse{Topics: topics}, nil
}

func (c *LearningStorage) CreateQuiz(req *pb.CreateQuizRequest) (*pb.CreateQuizResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO quizzes(id, topic_id, question, options, answer)
		VALUES($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(query, id, req.TopicId, req.Question, req.Options, req.Answer)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateQuizResponse{Id: id, Message: "success"}, nil
}

func (c *LearningStorage) GetQuiz(req *pb.GetQuizRequest) (*pb.GetQuizResponse, error) {
	query := `SELECT id, topic_id, question, options FROM quizzes`

	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var request pb.Quiz

	var quizzes []*pb.Quiz
	var arr []interface{}
	count := 1

	if len(request.Id) > 0 {
		query += fmt.Sprintf(" and id = $%d", count)
		count++
		arr = append(arr, request.Id)
	}
	if len(request.TopicId) > 0 {
		query += fmt.Sprintf(" and topic_id = $%d", count)
		count++
		arr = append(arr, request.TopicId)
	}

	rows, err = c.db.Query(query, arr...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var quiz pb.Quiz
		err := rows.Scan(&quiz.Id, &quiz.TopicId, &quiz.Question, &quiz.Options)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		quizzes = append(quizzes, &quiz)
	}
	return &pb.GetQuizResponse{Quiz: quizzes}, nil
}

func (c *LearningStorage) UpdateQuiz(req *pb.UpdateQuizRequest) (*pb.UpdateQuizResponse, error) {
	query := `
		UPDATE quizzes
		SET topic_id = $1, question = $2, options = $3, answer = $4
		WHERE id = $5`

	_, err := c.db.Exec(query, req.TopicId, req.Question, req.Options, req.Answer, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateQuizResponse{Message: "success"}, nil
}

func (c *LearningStorage) DeleteQuiz(req *pb.DeleteQuizRequest) (*pb.DeleteQuizResponse, error) {
	query := `DELETE FROM quizzes WHERE id = $1`

	_, err := c.db.Exec(query, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteQuizResponse{Message: "success"}, nil
}

func (c *LearningStorage) SubmitQuiz(req *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error) {
	query := `SELECT id, answer FROM quizzes WHERE id = $1 AND answer = $2`
	err := c.db.QueryRow(query, req.QuizId, req.Answer).Scan(&req.QuizId, &req.Answer)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("Incorrect answer or quiz not found")
			return &pb.SubmitQuizResponse{Message: "Incorrect answer or quiz not found", XpEarned: 0}, nil
		}
		log.Println(err)
		return nil, err
	}
	updateQuery := `UPDATE users SET xp = xp + $1 WHERE id = $2`
	_, err = c.db.Exec(updateQuery, 10, req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.SubmitQuizResponse{Message: "success", XpEarned: 10}, nil
}



func (c *LearningStorage) CreateExtraResourses(req *pb.CreateExtraResoursesRequest) (*pb.CreateExtraResoursesResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO extra_resources(id, title, type, url)
		VALUES($1, $2, $3, $4)`

	_, err := c.db.Exec(query, id, req.Title, req.Type, req.Url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateExtraResoursesResponse{Id: id, Message: "success"}, nil
}

func (c *LearningStorage) GetExtraResourses(req *pb.GetExtraResourcesRequest) (*pb.GetExtraResourcesResponse, error) {
	query := `SELECT id, title, type, url FROM extra_resources`
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var resources []*pb.CreateExtraResourses
	for rows.Next() {
		var resource pb.CreateExtraResourses
		err := rows.Scan(&resource.Id, &resource.Title, &resource.Type, &resource.Url)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resources = append(resources, &resource)
	}
	return &pb.GetExtraResourcesResponse{ExtraResources: resources}, nil
}

func (c *LearningStorage) UpdateExtraResourses(req *pb.UpdateExtraResoursesRequest) (*pb.UpdateExtraResoursesResponse, error) {
	query := `
		UPDATE extra_resources
		SET title = $1, type = $2, url = $3
		WHERE id = $4`

	_, err := c.db.Exec(query, req.Title, req.Type, req.Url, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateExtraResoursesResponse{Message: "success"}, nil
}

func (c *LearningStorage) DeleteExtraResourses(req *pb.DeleteExtraResoursesRequest) (*pb.DeleteExtraResoursesResponse, error) {
	query := `DELETE FROM extra_resources WHERE id = $1`

	_, err := c.db.Exec(query, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteExtraResoursesResponse{Message: "success"}, nil
}

func (c *LearningStorage) CompletedExtraResources(req *pb.CompletedExtraResourcesRequest) (*pb.CompletedExtraResourcesResponse, error) {
	query := `
		INSERT INTO completed_extra_resources (user_id, extra_resource_id)
		VALUES ($1, $2)`

	_, err := c.db.Exec(query, req.UserId, req.ExtraResourceId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	updateQuery := `UPDATE users SET xp = xp + $1 WHERE id = $2`
	_, err = c.db.Exec(updateQuery, 10, req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CompletedExtraResourcesResponse{Message: "success", XpEarned: 10}, nil
}


func (c *LearningStorage) GetLearningProgress(req *pb.GetLearningProgressRequest) (*pb.GetLearningProgressResponse, error) {
	fmt.Println(req.UserId)
	totalQuery := `SELECT 
		(SELECT COUNT(*) FROM topics) as total_topics, 
		(SELECT COUNT(*) FROM quizzes) as total_quizzes,
		(SELECT COUNT(*) FROM extra_resources) as total_resources,
		(SELECT COUNT(*) FROM completed_topics WHERE user_id = $1) as completed_topics,
		(SELECT COUNT(*) FROM completed_extra_resources WHERE user_id = $1) as completed_resources`
	var progress pb.GetLearningProgressResponse
	err := c.db.QueryRow(totalQuery, req.UserId).Scan(
		&progress.TotalTopics, &progress.TotalQuizzes, &progress.TotalResourses,
		&progress.CompletedTopics, &progress.CompletedResourses)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Calculate overall progress
	progress.OverallProgress = (float32(progress.CompletedTopics+progress.CompletedQuizzes+progress.CompletedResourses) /
		float32(progress.TotalTopics+progress.TotalQuizzes+progress.TotalResourses)) * 100

	return &progress, nil
}

func (c *LearningStorage) CreateLearningRecommendations(req *pb.CreateLearningRecommendationsRequest) (*pb.CreateLearningRecommendationsResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO recommendations(id, type, name, user_id, reason)
		VALUES($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(query, id, req.Type, req.Name, req.UserId, req.Reason)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateLearningRecommendationsResponse{Id: id, Message: "success"}, nil
}

func (c *LearningStorage) GetLearningRecommendations(req *pb.GetLearningRecommendationsRequest) (*pb.GetLearningRecommendationsResponse, error) {
	query := `SELECT id, type, name, user_id, reason FROM recommendations`
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var recommendations []*pb.Recommendation
	for rows.Next() {
		var recommendation pb.Recommendation
		err := rows.Scan(&recommendation.Id, &recommendation.Type, &recommendation.Name, &recommendation.UserId, &recommendation.Reason)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		recommendations = append(recommendations, &recommendation)
	}
	return &pb.GetLearningRecommendationsResponse{Recommendations: recommendations}, nil
}

func (c *LearningStorage) CreateLearningFeedback(req *pb.CreateLearningFeedbackRequest) (*pb.CreateLearningFeedbackResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO feedback(id, user_id, topic_id, rating, comment)
		VALUES($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(query, id, req.UserId, req.TopicId, req.Rating, req.Comment)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	updateQuery := `UPDATE users SET xp = xp + $1 WHERE id = $2`

	_, err = c.db.Exec(updateQuery, 10, req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateLearningFeedbackResponse{Message: "success", XpEarned: 10}, nil
}

func (c *LearningStorage) GetLearningFeedback(req *pb.GetLearningFeedbackRequest) (*pb.GetLearningFeedbackResponse, error) {
	query := `SELECT id, user_id, topic_id, rating, comment FROM feedback`
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*pb.LearningFeedback
	for rows.Next() {
		var feedback pb.LearningFeedback
		err := rows.Scan(&feedback.Id, &feedback.UserId, &feedback.TopicId, &feedback.Rating, &feedback.Comment)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		feedbacks = append(feedbacks, &feedback)
	}
	return &pb.GetLearningFeedbackResponse{Feedback: feedbacks}, nil
}

func (c *LearningStorage) CreateLearningHomeworks(req *pb.CreateLearningHomeworksRequest) (*pb.CreateLearningHomeworksResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO homeworks(id, user_id, title, description, difficulty)
		VALUES($1, $2, $3, $4, $5)`

	_, err := c.db.Exec(query, id, req.UserId, req.Title, req.Description, req.Difficulty)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateLearningHomeworksResponse{Id: id, Message: "success"}, nil
}

func (c *LearningStorage) GetLearningHomeworks(req *pb.GetLearningHomeworksRequest) (*pb.GetLearningHomeworksResponse, error) {
	query := `SELECT id, user_id, title, description, difficulty FROM homeworks`
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var homeworks []*pb.LearningHomeworks
	for rows.Next() {
		var homework pb.LearningHomeworks
		err := rows.Scan(&homework.Id, &homework.UserId, &homework.Title, &homework.Description, &homework.Difficulty)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		homeworks = append(homeworks, &homework)
	}
	return &pb.GetLearningHomeworksResponse{Homeworks: homeworks}, nil
}

func (c *LearningStorage) SubmitHomework(req *pb.SubmitHomeworkRequest) (*pb.SubmitHomeworkResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO submitted_homeworks (id, user_id, homework_id, xp_earned)
		VALUES ($1, $2, $3, $4)`

	_, err := c.db.Exec(query, id, req.UserId, req.HomeworkId, req.XpEarned)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	updateQuery := `UPDATE users SET xp = xp + $1 WHERE id = $2`

	_, err = c.db.Exec(updateQuery, req.XpEarned, req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.SubmitHomeworkResponse{Message: "success", XpEarned: req.XpEarned}, nil
}


