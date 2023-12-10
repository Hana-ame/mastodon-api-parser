package statuses

// View_edit_history_of_a_status godoc
// @Summary View edit history of a status
// @Description Get all known versions of a status, including the initial and current states.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The local ID of the Status in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.StatusEdit
// @Router /api/v1/statuses/:id/history [get]
func View_edit_history_of_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}