package search

// Perform_a_search godoc
// @Summary Perform a search
// @Description 
// @Tags mastodon,search
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param q query string true "REQUIRED String. The search query."
// @Param type query string false "String. Specify whether to search for only accounts, hashtags, statuses"
// @Param resolve query boolean false "Boolean. Attempt WebFinger lookup? Defaults to false."
// @Param following query boolean false "Boolean. Only include accounts that the user is following? Defaults to false."
// @Param account_id query string false "String. If provided, will only return statuses authored by this account."
// @Param exclude_unreviewed query boolean false "Boolean. Filter out unreviewed tags? Defaults to false. Use true when trying to find trending tags."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return, per type. Defaults to 20 results per category. Max 40 results per category."
// @Param offset query int false "Integer. Skip the first n results."
// @Success 200 object entities.Search
// @Router /api/v2/search [get]
func Perform_a_search(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}