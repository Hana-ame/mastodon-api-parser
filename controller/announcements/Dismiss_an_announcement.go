package announcements

// Dismiss_an_announcement godoc
// @Summary Dismiss an announcement
// @Description Allows a user to mark the announcement as read.
// @Tags mastodon,announcements
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Announcement in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Empty
// @Router /api/v1/announcements/:id/dismiss [post]
func Dismiss_an_announcement(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}