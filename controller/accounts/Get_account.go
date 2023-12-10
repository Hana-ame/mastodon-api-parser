package accounts

// Get_account godoc
// @Summary Get account
// @Description View information about a profile.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Account
// @Router /api/v1/accounts/:id [get]
func Get_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}