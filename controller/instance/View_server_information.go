package instance

// View_server_information godoc
// @Summary View server information
// @Description Obtain general information about the server.
// @Tags mastodon,instance
// @Produce json
// @Success 200 object entities.Instance
// @Router /api/v2/instance [get]
func View_server_information(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}