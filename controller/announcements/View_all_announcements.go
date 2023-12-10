package announcements

// View_all_announcements godoc
// @Summary View all announcements
// @Description See all currently active announcements set by admins.
// @Tags mastodon,announcements
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param with_dismissed query boolean false "Boolean. If true, response will include announcements dismissed by the user. Defaults to false."
// @Success 200 array entities.Announcement
// @Router /api/v1/announcements [get]
func View_all_announcements(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}