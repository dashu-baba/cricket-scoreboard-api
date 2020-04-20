package validators

import (
	"cricket-scoreboard-api/src/requestmodels"
	"errors"

	"github.com/gin-gonic/gin"
)

//ValidateOverUpdateModel godoc
// @Summary validates the OverUpdateModel model
func ValidateOverUpdateModel(c *gin.Context) (requestmodels.OverUpdateModel, error) {
	model := requestmodels.OverUpdateModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	if model.Wicket != (requestmodels.WicketDetailsModel{}) && (model.Wicket.BatsmanID == model.Wicket.BowlerID || model.Wicket.BatsmanID == model.Wicket.SupportID) {
		return model, errors.New("Player should not be same")
	}

	return model, nil
}

//ValidateCreateOverModel godoc
// @Summary validates the CreateOverModel model
func ValidateCreateOverModel(c *gin.Context) (requestmodels.CreateOverModel, error) {
	model := requestmodels.CreateOverModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}

//ValidateNextBatsmanModel godoc
// @Summary validates the NextBatsmanModel model
func ValidateNextBatsmanModel(c *gin.Context) (requestmodels.NextBatsmanModel, error) {
	model := requestmodels.NextBatsmanModel{}
	if err := c.ShouldBind(&model); err != nil {
		return model, err
	}

	return model, nil
}
