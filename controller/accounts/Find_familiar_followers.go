package accounts

// Find_familiar_followers godoc
// @Summary Find familiar followers
// @Description Obtain a list of all accounts that follow a given account, filtered for accounts you follow.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param id[] query string false "Array of String. Find familiar followers for the provided account IDs."
// @Success 200 array entities.FamiliarFollowers
// @Router /api/v1/accounts/familiar_followers [get]
func Find_familiar_followers(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}