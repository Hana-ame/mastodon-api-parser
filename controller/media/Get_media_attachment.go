package media

// Get_media_attachment godoc
// @Summary Get media attachment
// @Description Get a media attachment, before it is attached to a status and posted, but after it is accepted for processing. Use this method to check that the full-sized media has finished processing.
// @Tags mastodon,media
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the MediaAttachment in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.MediaAttachment
// @Router /api/v1/media/:id [get]
func Get_media_attachment(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}