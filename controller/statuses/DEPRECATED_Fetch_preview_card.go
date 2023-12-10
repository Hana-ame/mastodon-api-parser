package statuses

// DEPRECATED_Fetch_preview_card godoc
// @Summary (DEPRECATED) Fetch preview card
// @Description 
// @Tags mastodon,statuses
// @Produce json
// @Param id path string true "REQUIRED String. The local ID of the Status in the database."
// @Param Authorization header string false "Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.PreviewCard
// @Router /api/v1/statuses/:id/card [get]
func DEPRECATED_Fetch_preview_card(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}