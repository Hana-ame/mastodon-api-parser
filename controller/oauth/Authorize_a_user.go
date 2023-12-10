package oauth

// Authorize_a_user godoc
// @Summary Authorize a user
// @Description Displays an authorization form to the user. If approved, it will create and return an authorization code, then redirect to the desired redirect_uri, or show the authorization code if urn:ietf:wg:oauth:2.0:oob was requested. The authorization code can be used while requesting a token to obtain access to user-level methods.
// @Tags mastodon,oauth
// @Produce json
// @Param response_type query string true "REQUIRED String. Should be set equal to code."
// @Param client_id query string true "REQUIRED String. The client ID, obtained during app registration."
// @Param redirect_uri query string true "REQUIRED String. Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect_uris declared during app registration."
// @Param scope query string false "String. List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read."
// @Param force_login query boolean false "Boolean. Forces the user to re-login, which is necessary for authorizing with multiple accounts from the same instance."
// @Param lang query string false "String. The ISO 639-1 two-letter language code to use while rendering the authorization form."
// @Success 200 object entities.String (URL) or HTML response
// @Router /oauth/authorize [get]
func Authorize_a_user(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}