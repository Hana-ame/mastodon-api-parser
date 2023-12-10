package accounts

// Lookup_account_ID_from_Webfinger_address godoc
// @Summary Lookup account ID from Webfinger address
// @Description Quickly lookup a username to see if it is available, skipping WebFinger resolution.
// @Tags mastodon,accounts
// @Produce json
// @Param acct query string true "REQUIRED String. The username or Webfinger address to lookup."
// @Success 200 object entities.Account
// @Router /api/v1/accounts/lookup [get]
func Lookup_account_ID_from_Webfinger_address(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}