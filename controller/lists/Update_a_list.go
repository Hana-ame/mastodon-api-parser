package lists

// Update_a_list godoc
// @Summary Update a list
// @Description Change the title of a list, or which replies to show.
// @Tags mastodon,lists
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the List in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param title formData string true "REQUIRED String. The title of the list to be created."
// @Param replies_policy formData string false "String. One of followed, list, or none. Defaults to list."
// @Success 200 object entities.List
// @Router /api/v1/lists/:id [put]
func Update_a_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}