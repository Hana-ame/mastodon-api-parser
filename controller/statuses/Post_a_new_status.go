package statuses

// Post_a_new_status godoc
// @Summary Post a new status
// @Description Publish a status with the given parameters.
// @Tags mastodon,statuses
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param Idempotency-Key header string false "Provide this header with any arbitrary string to prevent duplicate submissions of the same status. Consider using a hash or UUID generated client-side. Idempotency keys are stored for up to 1 hour."
// @Param status formData string true "REQUIRED String. The text content of the status. If media_ids is provided, this becomes optional. Attaching a poll is optional while status is provided."
// @Param media_ids[] formData string true "REQUIRED Array of String. Include Attachment IDs to be attached as media. If provided, status becomes optional, and poll cannot be used."
// @Param poll[options][] formData string true "REQUIRED Array of String. Possible answers to the poll. If provided, media_ids cannot be used, and poll[expires_in] must be provided."
// @Param poll[expires_in] formData int true "REQUIRED Integer. Duration that the poll should be open, in seconds. If provided, media_ids cannot be used, and poll[options] must be provided."
// @Param poll[multiple] formData boolean false "Boolean. Allow multiple choices? Defaults to false."
// @Param poll[hide_totals] formData boolean false "Boolean. Hide vote counts until the poll ends? Defaults to false."
// @Param in_reply_to_id formData string false "String. ID of the status being replied to, if status is a reply."
// @Param sensitive formData boolean false "Boolean. Mark status and attached media as sensitive? Defaults to false."
// @Param spoiler_text formData string false "String. Text to be shown as a warning or subject before the actual content. Statuses are generally collapsed behind this field."
// @Param visibility formData string false "String. Sets the visibility of the posted status to public, unlisted, private, direct."
// @Param language formData string false "String. ISO 639 language code for this status."
// @Param scheduled_at formData string false "String. ISO 8601 Datetime at which to schedule a status. Providing this parameter will cause ScheduledStatus to be returned instead of Status. Must be at least 5 minutes in the future."
// @Success 200 object entities.Status. When scheduled_at is present, ScheduledStatus is returned instead.
// @Router /api/v1/statuses [post]
func Post_a_new_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}