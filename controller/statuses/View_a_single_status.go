package statuses

// View_a_single_status godoc
// @Summary View a single status
// @Description Obtain information about a status.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id [get]
func View_a_single_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}