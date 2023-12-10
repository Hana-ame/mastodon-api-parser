package announcements

// Remove_a_reaction_from_an_announcement godoc
// @Summary Remove a reaction from an announcement
// @Description Undo a react emoji to an announcement.
// @Tags mastodon,announcements
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Announcement in the database."
// @Param name path string true "REQUIRED String. Unicode emoji, or the shortcode of a custom emoji."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Empty
// @Router /api/v1/announcements/:id/reactions/:name [delete]
func Remove_a_reaction_from_an_announcement(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}