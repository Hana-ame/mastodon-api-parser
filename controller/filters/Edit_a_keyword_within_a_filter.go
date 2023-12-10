package filters

// Edit_a_keyword_within_a_filter godoc
// @Summary Edit a keyword within a filter
// @Description Update the given filter keyword.
// @Tags mastodon,filters
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the FilterKeyword in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param keyword formData string true "REQUIRED String. The keyword to be added to the filter group."
// @Param whole_word formData boolean false "Boolean. Whether the keyword should consider word boundaries."
// @Success 200 object entities.FilterKeyword
// @Router /api/v2/filters/keywords/:id [put]
func Edit_a_keyword_within_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}