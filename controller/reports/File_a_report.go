package reports

// File_a_report godoc
// @Summary File a report
// @Description 
// @Tags mastodon,reports
// @Produce json
// @Param Authorization header string true "REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
// @Param account_id formData string true "REQUIRED String. ID of the account to report."
// @Param status_ids[] formData string false "Array of String. You can attach statuses to the report to provide additional context."
// @Param comment formData string false "String. The reason for the report. Default maximum of 1000 characters."
// @Param forward formData boolean false "Boolean. If the account is remote, should the report be forwarded to the remote admin? Defaults to false."
// @Param category formData string false "String. Specify if the report is due to spam, violation of enumerated instance rules, or some other reason. Defaults to other. Will be set to violation if rule_ids[] is provided (regardless of any category value you provide)."
// @Param rule_ids[] formData int false "Array of Integer. For violation category reports, specify the ID of the exact rules broken. Rules and their IDs are available via GET /api/v1/instance/rules and GET /api/v1/instance."
// @Success 200 object entities.Report
// @Router /api/v1/reports [post]
func File_a_report(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}