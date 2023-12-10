package favourites

// View_favourited_statuses godoc
// @Summary View favourited statuses
// @Description Statuses the user has favourited.
// @Tags mastodon,favourites
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Success 200 array entities.Status
// @Router /api/v1/favourites [get]
func View_favourited_statuses(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}