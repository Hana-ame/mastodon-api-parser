package preferences

// View_user_preferences godoc
// @Summary View user preferences
// @Description Preferences defined by the user in their account settings.
// @Tags mastodon,preferences
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Preferences by key and value
// @Router /api/v1/preferences [get]
func View_user_preferences(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}