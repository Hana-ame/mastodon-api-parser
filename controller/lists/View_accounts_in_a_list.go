package lists

// View_accounts_in_a_list godoc
// @Summary View accounts in a list
// @Description 
// @Tags mastodon,lists
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the List in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results. Defaults to 40 accounts. Max 80 accounts. Set to 0 in order to get all accounts without pagination."
// @Success 200 array entities.Account
// @Router /api/v1/lists/:id/accounts [get]
func View_accounts_in_a_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}