package accounts

// Mute_account godoc
// @Summary Mute account
// @Description Mute the given account. Clients should filter statuses and notifications from this account, if received (e.g. due to a boost in the Home timeline).
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param notifications formData boolean false "Boolean. Mute notifications in addition to statuses? Defaults to true."
// @Param duration formData string false "Number. How long the mute should last, in seconds. Defaults to 0 (indefinite)."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/mute [post]
func Mute_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}