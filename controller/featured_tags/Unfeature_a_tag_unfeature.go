package featured_tags

// Unfeature_a_tag_unfeature godoc
// @Summary Unfeature a tag {unfeature}
// @Description Stop promoting a hashtag on your profile.
// @Tags mastodon,featured_tags
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FeaturedTag in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.empty object
// @Router /api/v1/featured_tags/:id [delete]
func Unfeature_a_tag_unfeature(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}