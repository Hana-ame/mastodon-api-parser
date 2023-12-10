package oauth

// Revoke_a_token godoc
// @Summary Revoke a token
// @Description Revoke an access token to make it no longer valid for use.
// @Tags mastodon,oauth
// @Produce json
// @Param client_id formData string true "REQUIRED String. The client ID, obtained during app registration."
// @Param client_secret formData string true "REQUIRED String. The client secret, obtained during app registration."
// @Param token formData string true "REQUIRED String. The previously obtained token, to be invalidated."
// @Success 200 object entities.empty object
// @Router /oauth/revoke [post]
func Revoke_a_token(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}