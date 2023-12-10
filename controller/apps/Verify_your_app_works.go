package apps

// Verify_your_app_works godoc
// @Summary Verify your app works
// @Description Confirm that the app’s OAuth2 credentials work.
// @Tags mastodon,apps
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Application, but without client_id or client_secret
// @Router /api/v1/apps/verify_credentials [get]
func Verify_your_app_works(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}