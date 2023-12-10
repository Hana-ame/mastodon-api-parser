package statuses

// Boost_a_status godoc
// @Summary Boost a status
// @Description Reshare a status on your own profile.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param visibility formData string false "String. Any visibility except limited or direct (i.e. public, unlisted, private). Defaults to public. Currently unused in UI."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id/reblog [post]
func Boost_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}