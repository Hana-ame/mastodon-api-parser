package emails

// Resend_confirmation_email godoc
// @Summary Resend confirmation email
// @Description 
// @Tags mastodon,emails
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param email formData string false "If provided, updates the unconfirmed user’s email before resending the confirmation email."
// @Success 200 object entities.Empty object
// @Router /api/v1/emails/confirmation [post]
func Resend_confirmation_email(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}