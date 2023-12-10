package statuses

// Delete_a_status godoc
// @Summary Delete a status
// @Description Delete one of your own statuses.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status with source text and poll or media_attachments
// @Router /api/v1/statuses/:id [delete]
func Delete_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}