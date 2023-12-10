package announcements

// Add_a_reaction_to_an_announcement godoc
// @Summary Add a reaction to an announcement
// @Description React to an announcement with an emoji.
// @Tags mastodon,announcements
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Announcement in the database."
// @Param name path string true "REQUIRED String. Unicode emoji, or the shortcode of a custom emoji."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Empty
// @Router /api/v1/announcements/:id/reactions/:name [put]
func Add_a_reaction_to_an_announcement(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}