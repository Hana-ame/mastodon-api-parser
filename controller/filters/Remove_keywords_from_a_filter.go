package filters

// Remove_keywords_from_a_filter godoc
// @Summary Remove keywords from a filter
// @Description Deletes the given filter keyword.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterKeyword in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v2/filters/keywords/:id [delete]
func Remove_keywords_from_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}