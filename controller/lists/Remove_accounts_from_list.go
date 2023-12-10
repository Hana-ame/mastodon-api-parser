package lists

// Remove_accounts_from_list godoc
// @Summary Remove accounts from list
// @Description Remove accounts from the given list.
// @Tags mastodon,lists
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the List in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param account_ids[] formData string true "REQUIRED Array of String. The accounts that should be removed from the list."
// @Success 200 object entities.empty object
// @Router /api/v1/lists/:id/accounts [delete]
func Remove_accounts_from_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}