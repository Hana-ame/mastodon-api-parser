package trends

// View_trending_statuses godoc
// @Summary View trending statuses
// @Description Statuses that have been interacted with more than others.
// @Tags mastodon,trends
// @Produce json
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Param offset query int false "Integer. Skip the first n results."
// @Success 200 array entities.Status
// @Router /api/v1/trends/statuses [get]
func View_trending_statuses(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}