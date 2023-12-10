package accounts

// Register_an_account godoc
// @Summary Register an account
// @Description Creates a user and account records. Returns an account access token for the app that initiated the request. The app should save this token for later, and should wait for the user to confirm their account by clicking a link in their email inbox.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param username formData string true "REQUIRED String. The desired username for the account"
// @Param email formData string true "REQUIRED String. The email address to be used for login"
// @Param password formData string true "REQUIRED String. The password to be used for login"
// @Param agreement formData boolean true "REQUIRED Boolean. Whether the user agrees to the local rules, terms, and policies. These should be presented to the user in order to allow them to consent before setting this parameter to TRUE."
// @Param locale formData string true "REQUIRED String. The language of the confirmation email that will be sent."
// @Param reason formData string false "String. If registrations require manual approval, this text will be reviewed by moderators."
// @Success 200 object entities.Token
// @Router /api/v1/accounts [post]
func Register_an_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}