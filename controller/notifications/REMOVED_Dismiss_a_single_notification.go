package notifications

// REMOVED_Dismiss_a_single_notification godoc
// @Summary (REMOVED) Dismiss a single notification
// @Description Dismiss a single notification from the server.
// @Tags mastodon,notifications
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param id formData string true "REQUIRED String. The ID of the notification in the database."
// @Success 200 object entities.empty object
// @Router /api/v1/notifications/dismiss [post]
func REMOVED_Dismiss_a_single_notification(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}