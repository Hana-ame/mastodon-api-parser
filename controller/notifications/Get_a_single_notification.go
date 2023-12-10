package notifications

// Get_a_single_notification godoc
// @Summary Get a single notification
// @Description View information about a notification with a given ID.
// @Tags mastodon,notifications
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Notification in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Notification
// @Router /api/v1/notifications/:id [get]
func Get_a_single_notification(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}