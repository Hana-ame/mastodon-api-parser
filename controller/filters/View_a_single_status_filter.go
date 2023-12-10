package filters

// View_a_single_status_filter godoc
// @Summary View a single status filter
// @Description Obtain a single status filter.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterStatus in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.FilterStatus
// @Router /api/v2/filters/statuses/:id [get]
func View_a_single_status_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}