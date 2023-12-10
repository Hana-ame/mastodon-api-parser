package accounts

// Set_private_note_on_profile godoc
// @Summary Set private note on profile
// @Description Sets a private note on a user.
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param comment formData string false "String. The comment to be set on that user. Provide an empty string or leave out this parameter to clear the currently set note."
// @Success 200 object entities.Relationship
// @Router /api/v1/accounts/:id/note [post]
func Set_private_note_on_profile(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}