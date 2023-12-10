package timelines

// DEPRECATED_View_direct_timeline godoc
// @Summary (DEPRECATED) View direct timeline
// @Description View statuses with a “direct” privacy, from your account or in your notifications.
// @Tags mastodon,timelines
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param since_id query string false "String. All results returned will be greater than this ID. In effect, sets a lower bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
// @Success 200 array entities.Status
// @Router /api/v1/timelines/direct [get]
func DEPRECATED_View_direct_timeline(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}