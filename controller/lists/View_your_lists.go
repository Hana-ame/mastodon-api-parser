package lists

// View_your_lists godoc
// @Summary View your lists
// @Description Fetch all lists that the user owns.
// @Tags mastodon,lists
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.List
// @Router /api/v1/lists [get]
func View_your_lists(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}