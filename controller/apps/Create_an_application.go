package apps

// Create_an_application godoc
// @Summary Create an application
// @Description Create a new application to obtain OAuth2 credentials.
// @Tags mastodon,apps
// @Produce json
// @Param client_name formData string true "REQUIRED String. A name for your application"
// @Param redirect_uris formData string true "REQUIRED String. Where the user should be redirected after authorization. To display the authorization code to the user instead of redirecting to a web page, use urn:ietf:wg:oauth:2.0:oob in this parameter."
// @Param scopes formData string false "String. Space separated list of scopes. If none is provided, defaults to read. See OAuth Scopes for a list of possible scopes."
// @Param website formData string false "String. A URL to the homepage of your app"
// @Success 200 object entities.Application
// @Router /api/v1/apps [post]
func Create_an_application(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}