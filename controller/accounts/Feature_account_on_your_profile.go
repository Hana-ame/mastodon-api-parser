package accounts

// Feature_account_on_your_profile godoc
// @Summary Feature account on your profile
// @Description Add the given account to the user’s featured profiles. (Featured profiles are currently shown on the user’s own public profile.)
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/pin [post]
func Feature_account_on_your_profile(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}