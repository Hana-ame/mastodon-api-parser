package statuses

// Pin_status_to_profile godoc
// @Summary Pin status to profile
// @Description Feature one of your own public statuses at the top of your profile.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The local ID of the Status in the database. The status should be authored by the authorized account."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id/pin [post]
func Pin_status_to_profile(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}