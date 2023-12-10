package instance

// List_of_connected_domains godoc
// @Summary List of connected domains
// @Description Domains that this instance is aware of.
// @Tags mastodon,instance
// @Produce json
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 array entities.String
// @Router /api/v1/instance/peers [get]
func List_of_connected_domains(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}