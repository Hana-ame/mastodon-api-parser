package accounts

// Remove_account_from_followers godoc
// @Summary Remove account from followers
// @Description Remove the given account from your followers.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/remove_from_followers [post]
func Remove_account_from_followers(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}