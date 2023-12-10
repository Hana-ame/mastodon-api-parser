package suggestions

// Remove_a_suggestion godoc
// @Summary Remove a suggestion
// @Description Remove an account from follow suggestions.
// @Tags mastodon,suggestions
// @Produce json
// @Param account_id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.n/a
// @Router /api/v1/suggestions/:account_id [delete]
func Remove_a_suggestion(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}