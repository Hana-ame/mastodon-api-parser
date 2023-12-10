package polls

// Vote_on_a_poll godoc
// @Summary Vote on a poll
// @Description 
// @Tags mastodon,polls
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Poll in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param choices[] formData int true "REQUIRED Array of Integer. Provide your own votes as an index for each option (starting from 0)."
// @Success 200 object entities.Poll
// @Router /api/v1/polls/:id/votes [post]
func Vote_on_a_poll(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}