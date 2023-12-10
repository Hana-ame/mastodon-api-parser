package statuses

// View_status_source godoc
// @Summary View status source
// @Description Obtain the source properties for a status so that it can be edited.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The local ID of the Status in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.StatusSource
// @Router /api/v1/statuses/:id/source [get]
func View_status_source(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}