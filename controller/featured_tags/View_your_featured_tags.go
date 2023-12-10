package featured_tags

// View_your_featured_tags godoc
// @Summary View your featured tags
// @Description List all hashtags featured on your profile.
// @Tags mastodon,featured_tags
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.FeaturedTag
// @Router /api/v1/featured_tags [get]
func View_your_featured_tags(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}