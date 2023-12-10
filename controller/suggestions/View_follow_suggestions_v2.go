package suggestions

// View_follow_suggestions_v2 godoc
// @Summary View follow suggestions (v2)
// @Description Accounts that are promoted by staff, or that the user has had past positive interactions with, but is not yet following.
// @Tags mastodon,suggestions
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
// @Success 200 array entities.Suggestion
// @Router /api/v2/suggestions [get]
func View_follow_suggestions_v2(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}