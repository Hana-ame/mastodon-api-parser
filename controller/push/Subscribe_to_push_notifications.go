package push

// Subscribe_to_push_notifications godoc
// @Summary Subscribe to push notifications
// @Description Add a Web Push API subscription to receive notifications. Each access token can have one push subscription. If you create a new subscription, the old subscription is deleted.
// @Tags mastodon,push
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param subscription[endpoint] formData string true "REQUIRED String. The endpoint URL that is called when a notification event occurs."
// @Param subscription[keys][p256dh] formData string true "REQUIRED String. User agent public key. Base64 encoded string of a public key from a ECDH keypair using the prime256v1 curve."
// @Param subscription[keys][auth] formData string true "REQUIRED String. Auth secret. Base64 encoded string of 16 bytes of random data."
// @Param data[alerts][mention] formData boolean false "Boolean. Receive mention notifications? Defaults to false."
// @Param data[alerts][status] formData boolean false "Boolean. Receive new subscribed account notifications? Defaults to false."
// @Param data[alerts][reblog] formData boolean false "Boolean. Receive reblog notifications? Defaults to false."
// @Param data[alerts][follow] formData boolean false "Boolean. Receive follow notifications? Defaults to false."
// @Param data[alerts][follow_request] formData boolean false "Boolean. Receive follow request notifications? Defaults to false."
// @Param data[alerts][favourite] formData boolean false "Boolean. Receive favourite notifications? Defaults to false."
// @Param data[alerts][poll] formData boolean false "Boolean. Receive poll notifications? Defaults to false."
// @Param data[alerts][update] formData boolean false "Boolean. Receive status edited notifications? Defaults to false."
// @Param data[alerts][update] formData string false "data[alerts][admin.sign_up]"
// @Param data[alerts][update] formData boolean false "Boolean. Receive new user signup notifications? Defaults to false. Must have a role with the appropriate permissions."
// @Param data[alerts][update] formData string false "data[alerts][admin.report]"
// @Param data[alerts][update] formData boolean false "Boolean. Receive new report notifications? Defaults to false. Must have a role with the appropriate permissions."
// @Param data[policy] formData string false "String. Specify whether to receive push notifications from all, followed, follower, or none users."
// @Success 200 object entities.WebPushSubscription
// @Router /api/v1/push/subscription [post]
func Subscribe_to_push_notifications(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}