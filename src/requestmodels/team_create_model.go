package requestmodels

// TeamCreateModel godoc
// @Summary Define Team create model
type TeamCreateModel struct {
	Name    string              `json:"name" form:"name" xml:"name" binding:"required"`
	Logo    string              `json:"logo" form:"logo" xml:"logo"`
	Players []PlayerCreateModel `json:"players" form:"players" xml:"players" binding:"required"`
}
