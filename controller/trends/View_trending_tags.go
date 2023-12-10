package trends

// View_trending_tags godoc
// @Summary View trending tags
// @Description Tags that are being used more frequently within the past week.
// @Tags mastodon,trends
// @Produce json
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 10 tags. Max 20 tags."
// @Param offset query int false "Integer. Skip the first n results."
// @Success 200 array entities.Tag
// @Router /api/v1/trends/tags [get]
func View_trending_tags(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}