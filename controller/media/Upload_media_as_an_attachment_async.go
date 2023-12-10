package media

// Upload_media_as_an_attachment_async godoc
// @Summary Upload media as an attachment (async)
// @Description Creates a media attachment to be used with a new status. The full sized media will be processed asynchronously in the background for large uploads.
// @Tags mastodon,media
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param file formData string true "REQUIRED Object. The file to be attached, encoded using multipart form data. The file must have a MIME type."
// @Param thumbnail formData string false "Object. The custom thumbnail of the media to be attached, encoded using multipart form data."
// @Param description formData string false "String. A plain-text description of the media, for accessibility purposes."
// @Param focus formData string false "String. Two floating points (x,y), comma-delimited, ranging from -1.0 to 1.0. See Focal points for cropping media thumbnails for more information."
// @Success 200 object entities.MediaAttachment, but without a URL
// @Router /api/v2/media [post]
func Upload_media_as_an_attachment_async(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}