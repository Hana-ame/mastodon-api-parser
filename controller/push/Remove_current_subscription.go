package push

// Remove_current_subscription godoc
// @Summary Remove current subscription
// @Description Removes the current Web Push API subscription.
// @Tags mastodon,push
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.none
// @Router /api/v1/push/subscription [delete]
func Remove_current_subscription(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}