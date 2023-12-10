package instance

// List_of_rules godoc
// @Summary List of rules
// @Description Rules that the users of this service should follow.
// @Tags mastodon,instance
// @Produce json
// @Success 200 array entities.Rule
// @Router /api/v1/instance/rules [get]
func List_of_rules(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}