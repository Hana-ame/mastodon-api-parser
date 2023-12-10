package tags

// View_information_about_a_single_tag godoc
// @Summary View information about a single tag
// @Description Show a hashtag and its associated information
// @Tags mastodon,tags
// @Produce json
// @Param id path string true "REQUIRED String. The name of the hashtag."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Tag
// @Router /api/v1/tags/:id [get]
func View_information_about_a_single_tag(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}