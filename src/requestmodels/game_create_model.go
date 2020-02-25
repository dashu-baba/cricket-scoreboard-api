package requestmodels

import (
	"cricket-scoreboard-api/src/models"
	"cricket-scoreboard-api/src/responsemodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GameCreateModel godoc
// @Summary Define Game create model
type GameCreateModel struct {
	Name     string          `json:"name" form:"name" xml:"name" binding:"required"`
	GameType models.GameType `json:"gameType" form:"gametype" xml:"gameType" binding:"required"`
	Teams    []string        `json:"teams" form:"teams" xml:"teams" binding:"required"`
}

//ValidateCreateTeamsRequests method validates the requests payload of CreateTeams method
func ValidateCreateTeamsRequests(model *GameCreateModel, c *gin.Context) responsemodels.ErrorModel {
	errorModel := responsemodels.ErrorModel{}
	if err := c.ShouldBind(&model); err != nil {
		errorModel = responsemodels.ErrorModel{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		}
	}

	return errorModel
}
