package suggestions

// DEPRECATED_View_follow_suggestions_v1 godoc
// @Summary (DEPRECATED) View follow suggestions (v1)
// @Description Accounts the user has had past positive interactions with, but is not yet following.
// @Tags mastodon,suggestions
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
// @Success 200 array entities.Account
// @Router /api/v1/suggestions [get]
func DEPRECATED_View_follow_suggestions_v1(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}