package timelines

// View_public_timeline godoc
// @Summary View public timeline
// @Description 
// @Tags mastodon,timelines
// @Produce json
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param local query boolean false "Boolean. Show only local statuses? Defaults to false."
// @Param remote query boolean false "Boolean. Show only remote statuses? Defaults to false."
// @Param only_media query boolean false "Boolean. Show only statuses with media attached? Defaults to false."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param since_id query string false "String. All results returned will be greater than this ID. In effect, sets a lower bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Success 200 array entities.Status
// @Router /api/v1/timelines/public [get]
func View_public_timeline(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}