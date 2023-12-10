package accounts

// Get_account_s_following godoc
// @Summary Get account’s following
// @Description Accounts which the given account is following, if network is not hidden by the account owner.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
// @Success 200 array entities.Account
// @Router /api/v1/accounts/:id/following [get]
func Get_account_s_following(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}