package accounts

// Verify_account_credentials godoc
// @Summary Verify account credentials
// @Description Test to make sure that the user token works.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.CredentialAccount
// @Router /api/v1/accounts/verify_credentials [get]
func Verify_account_credentials(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}