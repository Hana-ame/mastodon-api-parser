package markers

// Get_saved_timeline_positions godoc
// @Summary Get saved timeline positions
// @Description 
// @Tags mastodon,markers
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param timeline[] query string false "Array of String. Specify the timeline(s) for which markers should be fetched. Possible values: home, notifications. If not provided, an empty object will be returned."
// @Success 200 object entities.Hash of timeline key and associated Marker
// @Router /api/v1/markers [get]
func Get_saved_timeline_positions(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}