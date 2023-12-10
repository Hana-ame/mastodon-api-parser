package notifications

// Dismiss_all_notifications godoc
// @Summary Dismiss all notifications
// @Description Clear all notifications from the server.
// @Tags mastodon,notifications
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/notifications/clear [post]
func Dismiss_all_notifications(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}