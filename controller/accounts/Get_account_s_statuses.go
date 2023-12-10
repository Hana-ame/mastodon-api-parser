package accounts

// Get_account_s_statuses godoc
// @Summary Get account’s statuses
// @Description Statuses posted to the given account.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param since_id query string false "String. All results returned will be greater than this ID. In effect, sets a lower bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Param only_media query boolean false "Boolean. Filter out statuses without attachments."
// @Param exclude_replies query boolean false "Boolean. Filter out statuses in reply to a different account."
// @Param exclude_reblogs query boolean false "Boolean. Filter out boosts from the response."
// @Param pinned query boolean false "Boolean. Filter for pinned statuses only. Defaults to false, which includes all statuses. Pinned statuses do not receive special priority in the order of the returned results."
// @Param tagged query string false "String. Filter for statuses using a specific hashtag."
// @Success 200 array entities.Status
// @Router /api/v1/accounts/:id/statuses [get]
func Get_account_s_statuses(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}