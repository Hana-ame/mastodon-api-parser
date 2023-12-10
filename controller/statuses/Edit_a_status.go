package statuses

// Edit_a_status godoc
// @Summary Edit a status
// @Description Edit a given status to change its text, sensitivity, media attachments, or poll. Note that editing a poll’s options will reset the votes.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param status formData string false "String. The plain text content of the status."
// @Param spoiler_text formData string false "String. The plain text subject or content warning of the status."
// @Param sensitive formData boolean false "Boolean. Whether the status should be marked as sensitive."
// @Param language formData string false "String. ISO 639 language code for the status."
// @Param media_ids[] formData string false "Array of String. Include Attachment IDs to be attached as media. If provided, status becomes optional, and poll cannot be used."
// @Param poll[options][] formData string false "Array of String. Possible answers to the poll. If provided, media_ids cannot be used, and poll[expires_in] must be provided."
// @Param poll[expires_in] formData int false "Integer. Duration that the poll should be open, in seconds. If provided, media_ids cannot be used, and poll[options] must be provided."
// @Param poll[multiple] formData boolean false "Boolean. Allow multiple choices? Defaults to false."
// @Param poll[hide_totals] formData boolean false "Boolean. Hide vote counts until the poll ends? Defaults to false."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id [put]
func Edit_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}