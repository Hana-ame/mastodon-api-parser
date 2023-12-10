package endorsements

// View_currently_featured_profiles godoc
// @Summary View currently featured profiles
// @Description Accounts that the user is currently featuring on their profile.
// @Tags mastodon,endorsements
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
// @Success 200 array entities.Account
// @Router /api/v1/endorsements [get]
func View_currently_featured_profiles(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}