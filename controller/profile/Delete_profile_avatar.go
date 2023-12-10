package profile

// Delete_profile_avatar godoc
// @Summary Delete profile avatar
// @Description 
// @Tags mastodon,profile
// @Produce json
// @Success 200 object entities.CredentialAccount
// @Router /api/v1/profile/avatar [delete]
func Delete_profile_avatar(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}