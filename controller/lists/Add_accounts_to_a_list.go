package lists

// Add_accounts_to_a_list godoc
// @Summary Add accounts to a list
// @Description Add accounts to the given list. Note that the user must be following these accounts.
// @Tags mastodon,lists
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the List in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param account_ids[] formData string true "REQUIRED Array of String. The accounts that should be added to the list."
// @Success 200 object entities.empty object
// @Router /api/v1/lists/:id/accounts [post]
func Add_accounts_to_a_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}