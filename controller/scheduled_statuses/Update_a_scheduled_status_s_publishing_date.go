package scheduled_statuses

// Update_a_scheduled_status_s_publishing_date godoc
// @Summary Update a scheduled status’s publishing date
// @Description 
// @Tags mastodon,scheduled_statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the ScheduledStatus in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param scheduled_at formData string false "String. ISO 8601 Datetime at which the status will be published. Must be at least 5 minutes into the future."
// @Success 200 object entities.ScheduledStatus
// @Router /api/v1/scheduled_statuses/:id [put]
func Update_a_scheduled_status_s_publishing_date(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}