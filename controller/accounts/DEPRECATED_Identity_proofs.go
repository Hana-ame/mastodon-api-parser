package accounts

// DEPRECATED_Identity_proofs godoc
// @Summary (DEPRECATED) Identity proofs
// @Description 
// @Tags mastodon,accounts
// @Produce json
// @Param id path string true "REQUIRED String. The ID of the Account in the database."
// @Success 200 array entities.IdentityProof
// @Router /api/v1/accounts/:id/identity_proofs [get]
func DEPRECATED_Identity_proofs(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}