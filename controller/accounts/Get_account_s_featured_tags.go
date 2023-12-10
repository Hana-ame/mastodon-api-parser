package accounts

// Get_account_s_featured_tags godoc
// @Summary Get account’s featured tags
// @Description Tags featured by this account.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.FeaturedTag
// @Router /api/v1/accounts/:id/featured_tags [get]
func Get_account_s_featured_tags(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}