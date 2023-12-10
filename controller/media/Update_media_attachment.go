package media

// Update_media_attachment godoc
// @Summary Update media attachment
// @Description Update a MediaAttachment’s parameters, before it is attached to a status and posted.
// @Tags mastodon,media
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the MediaAttachment in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param thumbnail formData string false "Object. The custom thumbnail of the media to be attached, encoded using multipart form data."
// @Param description formData string false "String. A plain-text description of the media, for accessibility purposes."
// @Param focus formData string false "String. Two floating points (x,y), comma-delimited, ranging from -1.0 to 1.0. See Focal points for cropping media thumbnails for more information."
// @Success 200 object entities.MediaAttachment
// @Router /api/v1/media/:id [put]
func Update_media_attachment(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}