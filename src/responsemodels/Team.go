package responsemodels

// Team godoc
// @Summary Define Team model
type Team struct {
	ID      string   `json:"id" form:"id" xml:"id"`
	Name    string   `json:"name" form:"name" xml:"name"`
	Logo    string   `json:"logo" form:"logo" xml:"logo"`
	Players []Player `json:"players" form:"players" xml:"players"`
}
