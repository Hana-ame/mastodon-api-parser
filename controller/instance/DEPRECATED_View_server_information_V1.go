package instance

// DEPRECATED_View_server_information_V1 godoc
// @Summary (DEPRECATED) View server information (V1)
// @Description Obtain general information about the server.
// @Tags mastodon,instance
// @Produce json
// @Success 200 object entities.V1::Instance
// @Router /api/v1/instance [get]
func DEPRECATED_View_server_information_V1(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}