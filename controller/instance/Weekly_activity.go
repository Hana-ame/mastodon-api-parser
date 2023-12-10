package instance

// Weekly_activity godoc
// @Summary Weekly activity
// @Description Instance activity over the last 3 months, binned weekly.
// @Tags mastodon,instance
// @Produce json
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.Hash
// @Router /api/v1/instance/activity [get]
func Weekly_activity(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}