package instance

// View_extended_description godoc
// @Summary View extended description
// @Description Obtain an extended description of this server
// @Tags mastodon,instance
// @Produce json
// @Success 200 object entities.ExtendedDescription
// @Router /api/v1/instance/extended_description [get]
func View_extended_description(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}