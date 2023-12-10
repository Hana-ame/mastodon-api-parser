package lists

// Show_a_single_list godoc
// @Summary Show a single list
// @Description Fetch the list with the given ID. Used for verifying the title of a list, and which replies to show within that list.
// @Tags mastodon,lists
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the List in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.List
// @Router /api/v1/lists/:id [get]
func Show_a_single_list(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}