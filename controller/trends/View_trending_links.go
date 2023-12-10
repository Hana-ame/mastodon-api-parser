package trends

// View_trending_links godoc
// @Summary View trending links
// @Description Links that have been shared more than others.
// @Tags mastodon,trends
// @Produce json
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 10 links. Max 20 links."
// @Param offset query int false "Integer. Skip the first n results."
// @Success 200 array entities.Trends::Link
// @Router /api/v1/trends/links [get]
func View_trending_links(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}