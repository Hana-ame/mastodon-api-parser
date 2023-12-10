package polls

// View_a_poll godoc
// @Summary View a poll
// @Description 
// @Tags mastodon,polls
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Poll in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Poll
// @Router /api/v1/polls/:id [get]
func View_a_poll(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}