package featured_tags

// View_suggested_tags_to_feature godoc
// @Summary View suggested tags to feature
// @Description Shows up to 10 recently-used tags.
// @Tags mastodon,featured_tags
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.Tag
// @Router /api/v1/featured_tags/suggestions [get]
func View_suggested_tags_to_feature(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}