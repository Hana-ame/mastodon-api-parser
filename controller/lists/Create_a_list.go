package lists

// Create_a_list godoc
// @Summary Create a list
// @Description Create a new list.
// @Tags mastodon,lists
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param title formData string true "REQUIRED String. The title of the list to be created."
// @Param replies_policy formData string false "String. One of followed, list, or none. Defaults to list."
// @Param exclusive formData boolean false "Boolean. Whether members of this list need to get removed from the “Home” feed"
// @Success 200 object entities.List
// @Router /api/v1/lists [post]
func Create_a_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}