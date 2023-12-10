package oauth

// Obtain_a_token godoc
// @Summary Obtain a token
// @Description Obtain an access token, to be used during API calls that are not public.
// @Tags mastodon,oauth
// @Produce json
// @Param grant_type formData string true "REQUIRED String. Set equal to authorization_code if code is provided in order to gain user-level access. Otherwise, set equal to client_credentials to obtain app-level access only."
// @Param code formData string false "String. A user authorization code, obtained via GET /oauth/authorize."
// @Param client_id formData string true "REQUIRED String. The client ID, obtained during app registration."
// @Param client_secret formData string true "REQUIRED String. The client secret, obtained during app registration."
// @Param redirect_uri formData string true "REQUIRED String. Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the token will be shown instead. Must match one of the redirect_uris declared during app registration."
// @Param scope formData string false "String. List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). If code was provided, then this must be equal to the scope requested from the user. Otherwise, it must be a subset of scopes declared during app registration. If not provided, defaults to read."
// @Success 200 object entities.Token
// @Router /oauth/token [post]
func Obtain_a_token(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}