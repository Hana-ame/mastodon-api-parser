package statuses

// Undo_bookmark_of_a_status godoc
// @Summary Undo bookmark of a status
// @Description Remove a status from your private bookmarks.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Status
// @Router /api/v1/statuses/:id/unbookmark [post]
func Undo_bookmark_of_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}