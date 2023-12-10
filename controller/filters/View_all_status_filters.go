package filters

// View_all_status_filters godoc
// @Summary View all status filters
// @Description Obtain a list of all status filters within this filter group.
// @Tags mastodon,filters
// @Produce json
// @Param filter_id path string true "REQUIRED String. The ID of the Filter in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.FilterStatus
// @Router /api/v2/filters/:filter_id/statuses [get]
func View_all_status_filters(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}