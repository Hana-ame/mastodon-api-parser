package media

// DEPRECATED_Upload_media_as_an_attachment godoc
// @Summary (DEPRECATED) Upload media as an attachment
// @Description Creates an attachment to be used with a new status. This method will return after the full sized media is done processing.
// @Tags mastodon,media
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param file formData string true "REQUIRED Object. The file to be attached, encoded using multipart form data. The file must have a MIME type."
// @Param thumbnail formData string false "Object. The custom thumbnail of the media to be attached, encoded using multipart form data."
// @Param description formData string false "String. A plain-text description of the media, for accessibility purposes."
// @Param focus formData string false "String. Two floating points (x,y), comma-delimited, ranging from -1.0 to 1.0. See Focal points for cropping media thumbnails for more information."
// @Success 200 object entities.MediaAttachment
// @Router /api/v1/media [post]
func DEPRECATED_Upload_media_as_an_attachment(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}