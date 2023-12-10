package followed_tags

// View_all_followed_tags godoc
// @Summary View all followed tags
// @Description 
// @Tags mastodon,followed_tags
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 100 tags. Max 200 tags."
// @Success 200 array entities.Tag
// @Router /api/v1/followed_tags [get]
func View_all_followed_tags(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}