package statuses

// Unpin_status_from_profile godoc
// @Summary Unpin status from profile
// @Description Unfeature a status from the top of your profile.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The local ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id/unpin [post]
func Unpin_status_from_profile(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}