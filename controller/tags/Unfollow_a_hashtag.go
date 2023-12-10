package tags

// Unfollow_a_hashtag godoc
// @Summary Unfollow a hashtag
// @Description Unfollow a hashtag. Posts containing this hashtag will no longer be inserted into your home timeline.
// @Tags mastodon,tags
// @Produce json
// @Param id path string true "REQUIRED String. The name of the hashtag."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Tag
// @Router /api/v1/tags/:id/unfollow [post]
func Unfollow_a_hashtag(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}