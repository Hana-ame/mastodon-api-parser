package featured_tags

// Feature_a_tag godoc
// @Summary Feature a tag
// @Description Promote a hashtag on your profile.
// @Tags mastodon,featured_tags
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param name formData string true "REQUIRED String. The hashtag to be featured, without the hash sign."
// @Success 200 object entities.FeaturedTag
// @Router /api/v1/featured_tags [post]
func Feature_a_tag(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}