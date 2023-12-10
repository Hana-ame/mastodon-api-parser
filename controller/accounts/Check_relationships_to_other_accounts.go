package accounts

// Check_relationships_to_other_accounts godoc
// @Summary Check relationships to other accounts
// @Description Find out whether a given account is followed, blocked, muted, etc.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param id[] query string false "Array. Check relationships for the provided account IDs."
// @Success 200 array entities.Relationship
// @Router /api/v1/accounts/relationships [get]
func Check_relationships_to_other_accounts(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}