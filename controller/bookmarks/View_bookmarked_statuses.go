package bookmarks

// View_bookmarked_statuses godoc
// @Summary View bookmarked statuses
// @Description Statuses the user has bookmarked.
// @Tags mastodon,bookmarks
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Success 200 array entities.Status
// @Router /api/v1/bookmarks [get]
func View_bookmarked_statuses(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}