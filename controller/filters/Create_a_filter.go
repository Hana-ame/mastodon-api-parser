package filters

// Create_a_filter godoc
// @Summary Create a filter
// @Description 
// @Tags mastodon,filters
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param phrase formData string true "REQUIRED String. The text to be filtered."
// @Param context[] formData string true "REQUIRED Array of String. Where the filter should be applied. Specify at least one of home, notifications, public, thread, account."
// @Param irreversible formData boolean false "Boolean. Should the server irreversibly drop matching entities from home and notifications? Defaults to false."
// @Param whole_word formData boolean false "Boolean. Should the filter consider word boundaries for this keyword? Defaults to false."
// @Param expires_in formData int false "Integer. Number of seconds from now that the filter should expire. Otherwise, null for a filter that doesn’t expire."
// @Success 200 object entities.V1::Filter
// @Router /api/v1/filters [post]
func Create_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}