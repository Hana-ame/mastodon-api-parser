package accounts

// Search_for_matching_accounts godoc
// @Summary Search for matching accounts
// @Description Search for matching accounts by username or display name.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param q query string true "REQUIRED String. Search query for accounts."
// @Param limit query int false "Integer. Maximum number of results. Defaults to 40 accounts. Max 80 accounts."
// @Param offset query int false "Integer. Skip the first n results."
// @Param resolve query boolean false "Boolean. Attempt WebFinger lookup. Defaults to false. Use this when q is an exact address."
// @Param following query boolean false "Boolean. Limit the search to users you are following. Defaults to false."
// @Success 200 array entities.Account
// @Router /api/v1/accounts/search [get]
func Search_for_matching_accounts(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}