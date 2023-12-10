package notifications

// Dismiss_a_single_notification godoc
// @Summary Dismiss a single notification
// @Description Dismiss a single notification from the server.
// @Tags mastodon,notifications
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Notification in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/notifications/:id/dismiss [post]
func Dismiss_a_single_notification(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}