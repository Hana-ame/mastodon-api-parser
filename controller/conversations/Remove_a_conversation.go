package conversations

// Remove_a_conversation godoc
// @Summary Remove a conversation
// @Description Removes a conversation from your list of conversations.
// @Tags mastodon,conversations
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Conversation in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/conversations/:id [delete]
func Remove_a_conversation(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}