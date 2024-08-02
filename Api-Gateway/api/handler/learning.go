package handler

import (
	"encoding/json"
	"log"
	"net/http"

	pb "api-gateway/genproto/learning"

	"github.com/gin-gonic/gin"
)

// CreateLearningTopic creates a new topic
// @Summary Create learning topic
// @Description Create Learning topic
// @Tags topic
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param company body pb.CreateLearningTopicRequest true "Create topic"
// @Success 200 {object} pb.CreateLearningTopicResponse
// @Failure 400 {string} string "Error while creating company"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/create [post]
func (h *Handler) CreateLearningTopic(ctx *gin.Context) {
	req := pb.CreateLearningTopicRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateLearningTopicResponse{Message: "Invalid input"})
		return
	}
	input, err := json.Marshal(&req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = h.Kaf.ProduceMessages("app-c", input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err.Error())
		return
	}

	res, err := h.Learning.CreateLearningTopic(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateLearningTopicResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetLearningTopics retrieves all learning topics
// @Summary Get learning topics
// @Description Get Learning topics
// @Tags topic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Topic ID"
// @Param name query string false "Topic Name"
// @Param description query string false "Topic Description"
// @Param difficulty query string false "Topic Difficulty"
// @Success 200 {object} pb.GetLearningTopicsResponse
// @Failure 400 {string} string "Error while getting topics"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/topics [get]
func (h *Handler) GetLearningTopics(ctx *gin.Context) {
	req := pb.GetLearningTopicsRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.GetLearningTopicsResponse{})
		return
	}

	res, err := h.Learning.GetLearningTopics(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetLearningTopicsResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateLearningTopic updates a learning topic
// @Summary Update learning topic
// @Description Update Learning topic
// @Tags topic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic ID"
// @Param topic body pb.UpdateLearningTopicRequest true "Update topic"
// @Success 200 {object} pb.UpdateLearningTopicResponse
// @Failure 400 {string} string "Error while updating topic"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/update/{id} [put]
func (h *Handler) UpdateLearningTopic(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.UpdateLearningTopicRequest{Id: id}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.UpdateLearningTopicResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.UpdateLearningTopic(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.UpdateLearningTopicResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteLearningTopic deletes a learning topic
// @Summary Delete learning topic
// @Description Delete Learning topic
// @Tags topic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Topic ID"
// @Success 200 {object} pb.DeleteLearningTopicResponse
// @Failure 400 {string} string "Error while deleting topic"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/delete/{id} [delete]
func (h *Handler) DeleteLearningTopic(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.DeleteLearningTopicRequest{Id: id}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.DeleteLearningTopicResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.DeleteLearningTopic(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.DeleteLearningTopicResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CompletedTopics marks a topic as completed
// @Summary Mark topic as completed
// @Description Mark Learning topic as completed
// @Tags topic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param completion body pb.CompletedTopicsRequest true "Mark topic as completed"
// @Success 200 {object} pb.CompletedTopicsResponse
// @Failure 400 {string} string "Error while marking topic as completed"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/completed [post]
func (h *Handler) CompletedTopics(ctx *gin.Context) {
	req := pb.CompletedTopicsRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CompletedTopicsResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CompletedTopics(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CompletedTopicsResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetCompletedTopics retrieves completed topics
// @Summary Get completed topics
// @Description Get completed Learning topics
// @Tags topic
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param  query  query   pb.GetCompletedTopicsRequest true  "Query parameter"
// @Success 200 {object} pb.GetCompletedTopicsResponse
// @Failure 400 {string} string "Error while getting completed topics"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /topic/getcompleted [get]
func (h *Handler) GetCompletedTopics(ctx *gin.Context) {
	req := &pb.GetCompletedTopicsRequest{}
	req.Id = ctx.Query("id")
	req.TopicId = ctx.Query("topic_id")
	req.UserId = ctx.Query("user_id")

	res, err := h.Learning.GetCompletedTopics(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting completed topics"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CreateQuiz creates a new quiz
// @Summary Create quiz
// @Description Create quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param quiz body pb.CreateQuizRequest true "Create quiz"
// @Success 200 {object} pb.CreateQuizResponse
// @Failure 400 {string} string "Error while creating quiz"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /quiz/create [post]
func (h *Handler) CreateQuiz(ctx *gin.Context) {
	req := pb.CreateQuizRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateQuizResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CreateQuiz(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateQuizResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetQuiz retrieves all quizzes
// @Summary Get quizzes
// @Description Get quizzes
// @Tags quiz
// @Accept json
// @Produce json
// @Param id query string false "Quiz ID"
// @Param topic_id query string false "Topic id"
// @Security BearerAuth
// @Success 200 {object} pb.GetQuizResponse
// @Failure 400 {string} string "Error while getting quizzes"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /quiz/quizzes [get]
func (h *Handler) GetQuiz(ctx *gin.Context) {
	req := &pb.GetQuizRequest{}
	req.Id = ctx.Query("id")
	req.TopicId = ctx.Query("topic_id")
	res, err := h.Learning.GetQuiz(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetQuizResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateQuiz updates a quiz
// @Summary Update quiz
// @Description Update quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Quiz ID"
// @Param quiz body pb.UpdateQuizRequest true "Update quiz"
// @Success 200 {object} pb.UpdateQuizResponse
// @Failure 400 {string} string "Error while updating quiz"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /quiz/update/{id} [put]
func (h *Handler) UpdateQuiz(ctx *gin.Context) {
	req := pb.UpdateQuizRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.UpdateQuizResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.UpdateQuiz(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.UpdateQuizResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteQuiz deletes a quiz
// @Summary Delete quiz
// @Description Delete quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Quiz ID"
// @Success 200 {object} pb.DeleteQuizResponse
// @Failure 400 {string} string "Error while deleting quiz"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /quiz/delete/{id} [delete]
func (h *Handler) DeleteQuiz(ctx *gin.Context) {
	req := pb.DeleteQuizRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.DeleteQuizResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.DeleteQuiz(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.DeleteQuizResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// SubmitQuiz submits a quiz
// @Summary Submit quiz
// @Description Submit quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param submission body pb.SubmitQuizRequest true "Submit quiz"
// @Success 200 {object} pb.SubmitQuizResponse
// @Failure 400 {string} string "Error while submitting quiz"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /quiz/submit [post]
func (h *Handler) SubmitQuiz(ctx *gin.Context) {
	req := pb.SubmitQuizRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.SubmitQuizResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.SubmitQuiz(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.SubmitQuizResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CreateExtraResources creates a new extra resource
// @Summary Create extra resource
// @Description Create extra resource
// @Tags extra_resources
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param resource body pb.CreateExtraResoursesRequest true "Create extra resource"
// @Success 200 {object} pb.CreateExtraResoursesResponse
// @Failure 400 {string} string "Error while creating extra resource"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /extra_resources/create [post]
func (h *Handler) CreateExtraResourses(ctx *gin.Context) {
	req := pb.CreateExtraResoursesRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateExtraResoursesResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CreateExtraResourses(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateExtraResoursesResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetExtraResources retrieves all extra resources
// @Summary Get extra resources
// @Description Get extra resources
// @Tags extra_resources
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetExtraResourcesResponse
// @Failure 400 {string} string "Error while getting extra resources"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /extra_resources/get [get]
func (h *Handler) GetExtraResourses(ctx *gin.Context) {
	req := pb.GetExtraResourcesRequest{}
	res, err := h.Learning.GetExtraResourses(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetExtraResourcesResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateExtraResources updates an extra resource
// @Summary Update extra resource
// @Description Update extra resource
// @Tags extra_resources
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Resource ID"
// @Param resource body pb.UpdateExtraResoursesRequest true "Update extra resource"
// @Success 200 {object} pb.UpdateExtraResoursesResponse
// @Failure 400 {string} string "Error while updating extra resource"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /extra_resources/update/{id} [put]
func (h *Handler) UpdateExtraResourses(ctx *gin.Context) {
	req := pb.UpdateExtraResoursesRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.UpdateExtraResoursesResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.UpdateExtraResourses(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.UpdateExtraResoursesResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteExtraResources deletes an extra resource
// @Summary Delete extra resource
// @Description Delete extra resource
// @Tags extra_resources
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Resource ID"
// @Success 200 {object} pb.DeleteExtraResoursesResponse
// @Failure 400 {string} string "Error while deleting extra resource"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /extra_resources/delete/{id} [delete]
func (h *Handler) DeleteExtraResourses(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.DeleteExtraResoursesRequest{Id: id}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.DeleteExtraResoursesResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.DeleteExtraResourses(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.DeleteExtraResoursesResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CompletedExtraResources marks an extra resource as completed
// @Summary Mark extra resource as completed
// @Description Mark extra resource as completed
// @Tags extra_resources
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param completion body pb.CompletedExtraResourcesRequest true "Mark extra resource as completed"
// @Success 200 {object} pb.CompletedExtraResourcesResponse
// @Failure 400 {string} string "Error while marking extra resource as completed"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /extra_resources/completed [post]
func (h *Handler) CompletedExtraResources(ctx *gin.Context) {
	req := pb.CompletedExtraResourcesRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CompletedExtraResourcesResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CompletedExtraResources(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CompletedExtraResourcesResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetLearningProgress retrieves learning progress
// @Summary Get learning progress
// @Description Get learning progress
// @Tags progress
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.GetLearningProgressResponse
// @Failure 400 {string} string "Error while getting learning progress"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /progress/get [get]
func (h *Handler) GetLearningProgress(ctx *gin.Context) {
	id := ctx.Query("user_id")
	req := pb.GetLearningProgressRequest{UserId: id}
	res, err := h.Learning.GetLearningProgress(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetLearningProgressResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CreateLearningRecommendations creates a new learning recommendation
// @Summary Create learning recommendation
// @Description Create learning recommendation
// @Tags recommendations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param recommendation body pb.CreateLearningRecommendationsRequest true "Create learning recommendation"
// @Success 200 {object} pb.CreateLearningRecommendationsResponse
// @Failure 400 {string} string "Error while creating learning recommendation"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /recommendations/create [post]
func (h *Handler) CreateLearningRecommendations(ctx *gin.Context) {
	req := pb.CreateLearningRecommendationsRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateLearningRecommendationsResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CreateLearningRecommendations(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateLearningRecommendationsResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetLearningRecommendations retrieves all learning recommendations
// @Summary Get learning recommendations
// @Description Get learning recommendations
// @Tags recommendations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetLearningRecommendationsResponse
// @Failure 400 {string} string "Error while getting learning recommendations"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /recommendations/get [get]
func (h *Handler) GetLearningRecommendations(ctx *gin.Context) {
	req := pb.GetLearningRecommendationsRequest{}
	res, err := h.Learning.GetLearningRecommendations(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetLearningRecommendationsResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CreateLearningFeedback creates new learning feedback
// @Summary Create learning feedback
// @Description Create learning feedback
// @Tags feedback
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param feedback body pb.CreateLearningFeedbackRequest true "Create learning feedback"
// @Success 200 {object} pb.CreateLearningFeedbackResponse
// @Failure 400 {string} string "Error while creating learning feedback"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /feedback/create [post]
func (h *Handler) CreateLearningFeedback(ctx *gin.Context) {
	req := pb.CreateLearningFeedbackRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateLearningFeedbackResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CreateLearningFeedback(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateLearningFeedbackResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetLearningFeedback retrieves all learning feedback
// @Summary Get learning feedback
// @Description Get learning feedback
// @Tags feedback
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetLearningFeedbackResponse
// @Failure 400 {string} string "Error while getting learning feedback"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /feedback/get [get]
func (h *Handler) GetLearningFeedback(ctx *gin.Context) {
	req := pb.GetLearningFeedbackRequest{}
	res, err := h.Learning.GetLearningFeedback(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetLearningFeedbackResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CreateLearningHomeworks creates a new homework
// @Summary Create homework
// @Description Create homework
// @Tags homeworks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param homework body pb.CreateLearningHomeworksRequest true "Create homework"
// @Success 200 {object} pb.CreateLearningHomeworksResponse
// @Failure 400 {string} string "Error while creating homework"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /homeworks/create [post]
func (h *Handler) CreateLearningHomeworks(ctx *gin.Context) {
	req := pb.CreateLearningHomeworksRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateLearningHomeworksResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.CreateLearningHomeworks(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateLearningHomeworksResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetLearningHomeworks retrieves all homeworks
// @Summary Get homeworks
// @Description Get homeworks
// @Tags homeworks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetLearningHomeworksResponse
// @Failure 400 {string} string "Error while getting homeworks"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /homeworks/get [get]
func (h *Handler) GetLearningHomeworks(ctx *gin.Context) {
	req := pb.GetLearningHomeworksRequest{}
	res, err := h.Learning.GetLearningHomeworks(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetLearningHomeworksResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// SubmitHomework submits a homework
// @Summary Submit homework
// @Description Submit homework
// @Tags homeworks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param submission body pb.SubmitHomeworkRequest true "Submit homework"
// @Success 200 {object} pb.SubmitHomeworkResponse
// @Failure 400 {string} string "Error while submitting homework"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /homeworks/submit [post]
func (h *Handler) SubmitHomework(ctx *gin.Context) {
	req := pb.SubmitHomeworkRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.SubmitHomeworkResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Learning.SubmitHomework(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.SubmitHomeworkResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
