package filters

// Remove_a_filter godoc
// @Summary Remove a filter
// @Description 
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Filter in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/filters/:id [delete]
func Remove_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}