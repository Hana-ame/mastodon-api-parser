package tags

// Follow_a_hashtag godoc
// @Summary Follow a hashtag
// @Description Follow a hashtag. Posts containing a followed hashtag will be inserted into your home timeline.
// @Tags mastodon,tags
// @Produce json
// @Param id path string true "REQUIRED String. The name of the hashtag."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Tag
// @Router /api/v1/tags/:id/follow [post]
func Follow_a_hashtag(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}