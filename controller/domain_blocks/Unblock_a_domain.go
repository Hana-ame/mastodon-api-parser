package domain_blocks

// Unblock_a_domain godoc
// @Summary Unblock a domain
// @Description Remove a domain block, if it exists in the user’s array of blocked domains.
// @Tags mastodon,domain_blocks
// @Produce json
// @Success 200 object entities.empty object
// @Router /api/v1/domain_blocks [delete]
func Unblock_a_domain(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}