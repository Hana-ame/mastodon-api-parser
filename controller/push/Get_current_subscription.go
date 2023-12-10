package push

// Get_current_subscription godoc
// @Summary Get current subscription
// @Description View the PushSubscription currently associated with this access token.
// @Tags mastodon,push
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Success 200 object entities.WebPushSubscription
// @Router /api/v1/push/subscription [get]
func Get_current_subscription(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}