package search

// REMOVED_Search_results_v1 godoc
// @Summary (REMOVED) Search results (v1)
// @Description 
// @Tags mastodon,search
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param q query string true "REQUIRED String. The search query."
// @Param type query string false "String. Specify whether to search for only accounts, hashtags, statuses"
// @Param resolve query boolean false "Boolean. Attempt WebFinger lookup? Defaults to false."
// @Param account_id query string false "String. If provided, will only return statuses authored by this account."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return, per type. Defaults to 20 results per category. Max 40 results per category."
// @Param offset query int false "Integer. Offset in search results, used for pagination. Defaults to 0."
// @Success 200 object entities.Search, but hashtags is an array of strings instead of an array of Tag.
// @Router /api/v1/search [get]
func REMOVED_Search_results_v1(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}