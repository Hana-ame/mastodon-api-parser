package filters

// View_a_specific_filter godoc
// @Summary View a specific filter
// @Description Obtain a single filter group owned by the current user.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Filter in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Filter
// @Router /api/v2/filters/:id [get]
func View_a_specific_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}