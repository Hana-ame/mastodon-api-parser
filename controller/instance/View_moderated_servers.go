package instance

// View_moderated_servers godoc
// @Summary View moderated servers
// @Description Obtain a list of domains that have been blocked.
// @Tags mastodon,instance
// @Produce json
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.DomainBlock
// @Router /api/v1/instance/domain_blocks [get]
func View_moderated_servers(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}