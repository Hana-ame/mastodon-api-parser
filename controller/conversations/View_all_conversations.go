package conversations

// View_all_conversations godoc
// @Summary View all conversations
// @Description 
// @Tags mastodon,conversations
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 conversations. Max 40 conversations."
// @Success 200 array entities.Conversation
// @Router /api/v1/conversations [get]
func View_all_conversations(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}