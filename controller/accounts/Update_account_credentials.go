package accounts

// Update_account_credentials godoc
// @Summary Update account credentials
// @Description Update the user’s display and preferences.
// @Tags mastodon,accounts
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param display_name formData string false "String. The display name to use for the profile."
// @Param note formData string false "String. The account bio."
// @Param locked formData boolean false "Boolean. Whether manual approval of follow requests is required."
// @Param bot formData boolean false "Boolean. Whether the account has a bot flag."
// @Param discoverable formData boolean false "Boolean. Whether the account should be shown in the profile directory."
// @Param fields_attributes formData string false "Hash. The profile fields to be set. Inside this hash, the key is an integer cast to a string (although the exact integer does not matter), and the value is another hash including name and value. By default, max 4 fields."
// @Param fields_attributes[:index][name] formData string false "String. The name of the profile field. By default, max 255 characters."
// @Param fields_attributes[:index][value] formData string false "String. The value of the profile field. By default, max 255 characters."
// @Param source[privacy] formData string false "String. Default post privacy for authored statuses. Can be public, unlisted, or private."
// @Param source[sensitive] formData boolean false "Boolean. Whether to mark authored statuses as sensitive by default."
// @Param source[language] formData string false "String. Default language to use for authored statuses (ISO 6391)"
// @Success 200 object entities.the user’s own Account with source attribute
// @Router /api/v1/accounts/update_credentials [patch]
func Update_account_credentials(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}