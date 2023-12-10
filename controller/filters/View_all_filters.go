package filters

// View_all_filters godoc
// @Summary View all filters
// @Description Obtain a list of all filter groups for the current user.
// @Tags mastodon,filters
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.Filter
// @Router /api/v2/filters [get]
func View_all_filters(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}