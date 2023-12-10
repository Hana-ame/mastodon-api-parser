package domain_blocks

// Block_a_domain godoc
// @Summary Block a domain
// @Description Block a domain to:
// @Description 
// @Description hide all public posts from it
// @Description hide all notifications from it
// @Description remove all followers from it
// @Description prevent following new users from it (but does not remove existing follows)
// @Tags mastodon,domain_blocks
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param domain formData string true "REQUIRED String. Domain to block."
// @Success 200 object entities.empty object
// @Router /api/v1/domain_blocks [post]
func Block_a_domain(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}