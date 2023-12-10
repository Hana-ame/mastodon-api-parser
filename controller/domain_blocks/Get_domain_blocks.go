package domain_blocks

// Get_domain_blocks godoc
// @Summary Get domain blocks
// @Description View domains the user has blocked.
// @Tags mastodon,domain_blocks
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param since_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param min_id query string false "Internal parameter. Use HTTP Link header for pagination."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 100 domain blocks. Max 200 domain blocks."
// @Success 200 array entities.String
// @Router /api/v1/domain_blocks [get]
func Get_domain_blocks(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}