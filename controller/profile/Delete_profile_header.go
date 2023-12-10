package profile

// Delete_profile_header godoc
// @Summary Delete profile header
// @Description 
// @Tags mastodon,profile
// @Produce json
// @Success 200 object entities.CredentialAccount
// @Router /api/v1/profile/header [delete]
func Delete_profile_header(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}