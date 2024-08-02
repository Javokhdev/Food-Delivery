package handler

import (
	"context"
	"net/http"

	pb "api-gateway/genproto/game"

	"github.com/gin-gonic/gin"

	pbu "api-gateway/genproto/user"
)

// CreateGameLevel creates a new game level
// @Summary Create game level
// @Description Create Game level
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.CreateGameLevelRequest true "Create game level"
// @Success 200 {object} pb.CreateGameLevelResponse
// @Failure 400 {string} string "Error while creating game level"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/create [post]
func (h *Handler) CreateGameLevel(ctx *gin.Context) {
	req := pb.CreateGameLevelRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateGameLevelResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.CreateGameLevel(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateGameLevelResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetGameLevels gets all game levels
// @Summary Get game levels
// @Description Get all game levels
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetGameLevelsResponse
// @Failure 400 {string} string "Error while getting game levels"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/get [get]
func (h *Handler) GetGameLevels(ctx *gin.Context) {
	req := pb.GetGameLevelsRequest{}

	res, err := h.Game.GetGameLevels(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetGameLevelsResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateGameLevel updates a game level
// @Summary Update game level
// @Description Update Game level
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.UpdateGameLevelRequest true "Update game level"
// @Success 200 {object} pb.UpdateGameLevelResponse
// @Failure 400 {string} string "Error while updating game level"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/update [put]
func (h *Handler) UpdateGameLevel(ctx *gin.Context) {
	req := pb.UpdateGameLevelRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.UpdateGameLevelResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.UpdateGameLevel(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.UpdateGameLevelResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteGameLevel deletes a game level
// @Summary Delete game level
// @Description Delete Game level
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Game Level ID"
// @Success 200 {object} pb.DeleteGameLevelResponse
// @Failure 400 {string} string "Error while deleting game level"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/delete/{id} [delete]
func (h *Handler) DeleteGameLevel(ctx *gin.Context) {
	req := pb.DeleteGameLevelRequest{}
	id := ctx.Param("id")
	req.Id = id

	res, err := h.Game.DeleteGameLevel(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.DeleteGameLevelResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// BeginGameLevel begins a game level
// @Summary Begin game level
// @Description Begin Game level
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.BeginGameLevelRequest true "Begin game level"
// @Success 200 {object} pb.BeginGameLevelResponse
// @Failure 400 {string} string "Error while beginning game level"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/begin [post]
func (h *Handler) BeginGameLevel(ctx *gin.Context) {
	req := pb.BeginGameLevelRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.BeginGameLevelResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.BeginGameLevel(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.BeginGameLevelResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// CompleteGameLevel completes a game level
// @Summary Complete game level
// @Description Complete Game level
// @Tags game_level
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.CompleteGameLevelRequest true "Complete game level"
// @Success 200 {object} pb.CompleteGameLevelResponse
// @Failure 400 {string} string "Error while completing game level"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /level/complete [post]
func (h *Handler) CompleteGameLevel(c *gin.Context) {
	req := pb.CompleteGameLevelRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &pb.CompleteGameLevelResponse{Message: "Invalid input"})
		return
	}

	// Complete the game level
	res, err := h.Game.CompleteGameLevel(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &pb.CompleteGameLevelResponse{Message: err.Error()})
		return
	}

	// Create GetXpRequest and call GetXp
	var xpReq pbu.GetXpRequest
	xpReq.UserId = req.GetUserId()
	xpReq.Xp = res.Xp // Adjust according to your response struct field containing the earned XP

	xpRes, err := h.User.GetXp(context.Background(), &xpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to give XP to user"})
		return
	}

	// Ensure XP update was successful
	if xpRes.Message != "success" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to give XP to user"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateGameChallenge creates a new game challenge
// @Summary Create game challenge
// @Description Create Game challenge
// @Tags game_challenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.CreateGameChallengeRequest true "Create game challenge"
// @Success 200 {object} pb.CreateGameChallengeResponse
// @Failure 400 {string} string "Error while creating game challenge"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /challenge/create [post]
func (h *Handler) CreateGameChallenge(ctx *gin.Context) {
	req := pb.CreateGameChallengeRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.CreateGameChallengeResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.CreateGameChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.CreateGameChallengeResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetGameChallenge gets a game challenge
// @Summary Get game challenge
// @Description Get Game challenge
// @Tags game_challenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Game Challenge ID"
// @Success 200 {object} pb.GetGameChallengeResponse
// @Failure 400 {string} string "Error while getting game challenge"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /challenge/get/{id} [get]
func (h *Handler) GetGameChallenge(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.GetGameChallengeRequest{Id: id}
	req.Id = id

	res, err := h.Game.GetGameChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetGameChallengeResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateGameChallenge updates a game challenge
// @Summary Update game challenge
// @Description Update Game challenge
// @Tags game_challenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.UpdateGameChallengeRequest true "Update game challenge"
// @Success 200 {object} pb.UpdateGameChallengeResponse
// @Failure 400 {string} string "Error while updating game challenge"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /challenge/update [put]
func (h *Handler) UpdateGameChallenge(ctx *gin.Context) {
	req := pb.UpdateGameChallengeRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.UpdateGameChallengeResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.UpdateGameChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.UpdateGameChallengeResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// DeleteGameChallenge deletes a game challenge
// @Summary Delete game challenge
// @Description Delete Game challenge
// @Tags game_challenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Game Challenge ID"
// @Success 200 {object} pb.DeleteGameChallengeResponse
// @Failure 400 {string} string "Error while deleting game challenge"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /challenge/delete/{id} [delete]
func (h *Handler) DeleteGameChallenge(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.DeleteGameChallengeRequest{Id: id}
	req.Id = id

	res, err := h.Game.DeleteGameChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.DeleteGameChallengeResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// SubmitGameChallengeAnswer submits an answer to a game challenge
// @Summary Submit game challenge
// @Description Submit Game challenge
// @Tags game_challenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param game body pb.SubmitChallengeRequest true "Submit game challenge answer"
// @Success 200 {object} pb.SubmitChallengeResponse
// @Failure 400 {string} string "Error while submitting game challenge answer"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /challenge/submit [post]
func (h *Handler) SubmitChallenge(ctx *gin.Context) {
	req := pb.SubmitChallengeRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &pb.SubmitChallengeResponse{Message: "Invalid input"})
		return
	}

	res, err := h.Game.SubmitChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.SubmitChallengeResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetGameLeaderboard gets the leaderboard
// @Summary Get Game leaderboard
// @Description Get the leaderboard
// @Tags leaderboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetGameLeaderboardResponse
// @Failure 400 {string} string "Error while getting leaderboard"
// @Failure 500 {string} string "500 – Internal Server Error"
// @Router /game/leaderboard [get]
func (h *Handler) GetGameLeaderboard(ctx *gin.Context) {
	req := pb.GetGameLeaderboardRequest{}

	res, err := h.Game.GetGameLeaderboard(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &pb.GetGameLeaderboardResponse{})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
