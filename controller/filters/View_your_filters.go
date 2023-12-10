package filters

// View_your_filters godoc
// @Summary View your filters
// @Description 
// @Tags mastodon,filters
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.List of V1::Filter
// @Router /api/v1/filters [get]
func View_your_filters(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}