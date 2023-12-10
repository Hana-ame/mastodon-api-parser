package notifications

// Get_all_notifications godoc
// @Summary Get all notifications
// @Description Notifications concerning the user. This API returns Link headers containing links to the next/previous page. However, the links can also be constructed dynamically using query params and id values.
// @Description 
// @Description Types to filter include:
// @Description 
// @Description mention = Someone mentioned you in their status
// @Description status = Someone you enabled notifications for has posted a status
// @Description reblog = Someone boosted one of your statuses
// @Description follow = Someone followed you
// @Description follow_request = Someone requested to follow you
// @Description favourite = Someone favourited one of your statuses
// @Description poll = A poll you have voted in or created has ended
// @Description update = A status you boosted with has been edited
// @Description admin.sign_up = Someone signed up (optionally sent to admins)
// @Description admin.report = A new report has been filed
// @Tags mastodon,notifications
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param max_id query string false "String. All results returned will be lesser than this ID. In effect, sets an upper bound on results."
// @Param since_id query string false "String. All results returned will be greater than this ID. In effect, sets a lower bound on results."
// @Param min_id query string false "String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
// @Param limit query int false "Integer. Maximum number of results to return. Defaults to 15 notifications. Max 30 notifications."
// @Param types[] query string false "Array of String. Types to include in the result."
// @Param exclude_types[] query string false "Array of String. Types to exclude from the results."
// @Param account_id query string false "String. Return only notifications received from the specified account."
// @Success 200 array entities.Notification
// @Router /api/v1/notifications [get]
func Get_all_notifications(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}