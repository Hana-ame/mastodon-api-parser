package accounts

// Follow_account godoc
// @Summary Follow account
// @Description Follow the given account. Can also be used to update whether to show reblogs or enable notifications.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param reblogs formData boolean false "Boolean. Receive this account’s reblogs in home timeline? Defaults to true."
// @Param notify formData boolean false "Boolean. Receive notifications when this account posts a status? Defaults to false."
// @Param languages formData string false "Array of String (ISO 639-1 language two-letter code). Filter received statuses for these languages. If not provided, you will receive this account’s posts in all languages."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/follow [post]
func Follow_account(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}