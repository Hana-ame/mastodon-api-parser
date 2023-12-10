package accounts

// Get_lists_containing_this_account godoc
// @Summary Get lists containing this account
// @Description User lists that you have added this account to.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.List
// @Router /api/v1/accounts/:id/lists [get]
func Get_lists_containing_this_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}