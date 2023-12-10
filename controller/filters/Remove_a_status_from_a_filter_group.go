package filters

// Remove_a_status_from_a_filter_group godoc
// @Summary Remove a status from a filter group
// @Description Add a status filter to the current filter group.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterStatus in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.FilterStatus
// @Router /api/v2/filters/statuses/:id [delete]
func Remove_a_status_from_a_filter_group(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}