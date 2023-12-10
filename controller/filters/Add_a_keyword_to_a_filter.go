package filters

// Add_a_keyword_to_a_filter godoc
// @Summary Add a keyword to a filter
// @Description Add the given keyword to the specified filter group
// @Tags mastodon,filters
// @Produce json
// @Param filter_id path string true "REQUIRED String. The ID of the Filter in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param keyword formData string true "REQUIRED String. The keyword to be added to the filter group."
// @Param whole_word formData boolean false "Boolean. Whether the keyword should consider word boundaries."
// @Success 200 object entities.FilterKeyword
// @Router /api/v2/filters/:filter_id/keywords [post]
func Add_a_keyword_to_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}