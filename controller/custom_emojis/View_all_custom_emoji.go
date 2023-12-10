package custom_emojis

// View_all_custom_emoji godoc
// @Summary View all custom emoji
// @Description Returns custom emojis that are available on the server.
// @Tags mastodon,custom_emojis
// @Produce json
// @Success 200 array entities.CustomEmoji
// @Router /api/v1/custom_emojis [get]
func View_all_custom_emoji(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}