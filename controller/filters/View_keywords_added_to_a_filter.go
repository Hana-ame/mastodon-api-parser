package filters

// View_keywords_added_to_a_filter godoc
// @Summary View keywords added to a filter
// @Description List all keywords attached to the current filter group.
// @Tags mastodon,filters
// @Produce json
// @Param filter_id path string true "REQUIRED String. The ID of the Filter in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.FilterKeyword
// @Router /api/v2/filters/:filter_id/keywords [get]
func View_keywords_added_to_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}