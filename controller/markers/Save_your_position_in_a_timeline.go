package markers

// Save_your_position_in_a_timeline godoc
// @Summary Save your position in a timeline
// @Description 
// @Tags mastodon,markers
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param home[last_read_id] formData string false "String. ID of the last status read in the home timeline."
// @Param notifications[last_read_id] formData string false "String. ID of the last notification read."
// @Success 200 object entities.Marker
// @Router /api/v1/markers [post]
func Save_your_position_in_a_timeline(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}