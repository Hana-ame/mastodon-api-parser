package filters

// Delete_a_filter godoc
// @Summary Delete a filter
// @Description Delete a filter group with the given id.
// @Tags mastodon,filters
// @Produce json
// @Success 200 object entities.empty object
// @Router /api/v2/filters/:id [delete]
func Delete_a_filter(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}