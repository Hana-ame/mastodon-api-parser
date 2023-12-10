package statuses

// Get_parent_and_child_statuses_in_context godoc
// @Summary Get parent and child statuses in context
// @Description View statuses above and below this status in the thread.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Context
// @Router /api/v1/statuses/:id/context [get]
func Get_parent_and_child_statuses_in_context(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}