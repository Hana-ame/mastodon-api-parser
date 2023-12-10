package accounts

// Unfeature_account_from_profile godoc
// @Summary Unfeature account from profile
// @Description Remove the given account from the user’s featured profiles.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/unpin [post]
func Unfeature_account_from_profile(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}