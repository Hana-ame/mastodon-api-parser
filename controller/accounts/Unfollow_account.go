package accounts

// Unfollow_account godoc
// @Summary Unfollow account
// @Description Unfollow the given account.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/unfollow [post]
func Unfollow_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}