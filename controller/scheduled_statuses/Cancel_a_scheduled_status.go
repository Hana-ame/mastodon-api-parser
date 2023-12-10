package scheduled_statuses

// Cancel_a_scheduled_status godoc
// @Summary Cancel a scheduled status
// @Description 
// @Tags mastodon,scheduled_statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the ScheduledStatus in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/scheduled_statuses/:id [delete]
func Cancel_a_scheduled_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}