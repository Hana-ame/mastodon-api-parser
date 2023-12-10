package filters

// View_a_single_filter godoc
// @Summary View a single filter
// @Description 
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterKeyword in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.V1::Filter
// @Router /api/v1/filters/:id [get]
func View_a_single_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}