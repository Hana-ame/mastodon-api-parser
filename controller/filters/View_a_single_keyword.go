package filters

// View_a_single_keyword godoc
// @Summary View a single keyword
// @Description Get one filter keyword by the given id.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterKeyword in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.FilterKeyword
// @Router /api/v2/filters/keywords/:id [get]
func View_a_single_keyword(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}