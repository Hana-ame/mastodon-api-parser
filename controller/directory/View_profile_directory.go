package directory

// View_profile_directory godoc
// @Summary View profile directory
// @Description List accounts visible in the directory.
// @Tags mastodon,directory
// @Produce json
// @Param offset query string false "Number. Skip the first n results."
// @Param limit query string false "Number. How many accounts to load. Defaults to 40 accounts. Max 80 accounts."
// @Param order query string false "String. Use active to sort by most recently posted statuses (default) or new to sort by most recently created profiles."
// @Param local query boolean false "Boolean. If true, returns only local accounts."
// @Success 200 array entities.Account
// @Router /api/v1/directory [get]
func View_profile_directory(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}