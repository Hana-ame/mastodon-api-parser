package timelines

// View_hashtag_timeline godoc
// @Summary View hashtag timeline
// @Description View public statuses containing the given hashtag.
// @Tags mastodon,timelines
// @Produce json
// @Param hashtag path string true "REQUIRED String. The name of the hashtag (not including the # symbol)."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param any[] query string false "Array of String. Return statuses that contain any of these additional tags."
// @Param all[] query string false "Array of String. Return statuses that contain all of these additional tags."
// @Param none[] query string false "Array of String. Return statuses that contain none of these additional tags."
// @Param local query boolean false "Boolean. Return only local statuses? Defaults to false."
// @Param remote query boolean false "Boolean. Return only remote statuses? Defaults to false."
// @Param only_media query boolean false "Boolean. Return only statuses with media attachments? Defaults to false."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param since_id query string false "String. All results returned will be greater than this ID. In effect, sets a lower bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Success 200 array entities.Status
// @Router /api/v1/timelines/tag/:hashtag [get]
func View_hashtag_timeline(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}