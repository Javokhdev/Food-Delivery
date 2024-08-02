package api

import (
	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/docs"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	// ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// err = ca.LoadPolicy()
	// if err != nil {
	// 	log.Fatal("casbin error load policy: ", err)
	// 	panic(err)
	// }

	r := gin.Default()

	// r.Group("/")
	// router.Use(middleware.NewAuth(ca))

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	c := r.Group("/topic")
	c.POST("/create", h.CreateLearningTopic)
	c.GET("/topics", h.GetLearningTopics)
	c.PUT("/update/:id", h.UpdateLearningTopic)
	c.DELETE("/delete/:id", h.DeleteLearningTopic)
	c.POST("/completed", h.CompletedTopics)
	c.GET("/getcompleted", h.GetCompletedTopics)

	q := r.Group("/quiz")
	q.POST("/create", h.CreateQuiz)
	q.GET("/quizzes", h.GetQuiz)
	q.PUT("/update/:id", h.UpdateQuiz)
	q.DELETE("/delete/:id", h.DeleteQuiz)
	q.POST("/submit", h.SubmitQuiz)

	rs := r.Group("/extra_resources")
	rs.POST("/create", h.CreateExtraResourses)
	rs.GET("/get", h.GetExtraResourses)
	rs.PUT("/update/:id", h.UpdateExtraResourses)
	rs.DELETE("/delete/:id", h.DeleteExtraResourses)
	rs.POST("/completed", h.CompletedExtraResources)

	ps := r.Group("/progress")
	ps.GET("/get", h.GetLearningProgress)

	rn := r.Group("/recommendations")
	rn.POST("/create", h.CreateLearningRecommendations)
	rn.GET("/get", h.GetLearningRecommendations)

	f := r.Group("/feedback")
	f.POST("/create", h.CreateLearningFeedback)
	f.GET("/get", h.GetLearningFeedback)

	hw := r.Group("/homeworks")
	hw.POST("/create", h.CreateLearningHomeworks)
	hw.GET("/get", h.GetLearningHomeworks)
	hw.POST("/submit", h.SubmitHomework)

	l := r.Group("/level")
	l.POST("/create", h.CreateGameLevel)
	l.GET("/get", h.GetGameLevels)
	l.PUT("/update", h.UpdateGameLevel)
	l.DELETE("/delete/:id", h.DeleteGameLevel)
	l.POST("/begin", h.BeginGameLevel)
	l.POST("/complete", h.CompleteGameLevel)

	ch := r.Group("/challenge")
	ch.POST("/create", h.CreateGameChallenge)
	ch.GET("/get/:id", h.GetGameChallenge)
	ch.PUT("/update/", h.UpdateGameChallenge)
	ch.DELETE("/delete/:id", h.DeleteGameChallenge)
	ch.POST("/submit", h.SubmitChallenge)

	g := r.Group("/game")
	g.GET("/leaderboard", h.GetGameLeaderboard)

	return r
}
