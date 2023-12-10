package statuses

// Unmute_a_conversation godoc
// @Summary Unmute a conversation
// @Description Start receiving notifications again for the thread that this status is part of.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id/unmute [post]
func Unmute_a_conversation(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}