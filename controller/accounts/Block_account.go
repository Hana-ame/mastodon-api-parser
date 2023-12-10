package accounts

// Block_account godoc
// @Summary Block account
// @Description Block the given account. Clients should filter statuses from this account if received (e.g. due to a boost in the Home timeline)
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/block [post]
func Block_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}