package statuses

// Translate_a_status godoc
// @Summary Translate a status
// @Description Translate the status content into some language.
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Status in the database."
// @Param lang formData string false "String (ISO 639 language code). The status content will be translated into this language. Defaults to the user’s current locale."
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.Translation
// @Router /api/v1/statuses/:id/translate [post]
func Translate_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}