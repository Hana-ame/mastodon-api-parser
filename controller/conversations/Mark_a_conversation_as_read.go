package conversations

// Mark_a_conversation_as_read godoc
// @Summary Mark a conversation as read
// @Description 
// @Tags mastodon,conversations
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Conversation in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Conversation
// @Router /api/v1/conversations/:id/read [post]
func Mark_a_conversation_as_read(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}