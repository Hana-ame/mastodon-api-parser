// https://docs.joinmastodon.org/methods/accounts

package mastodon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateCredentials godoc
//
//	@Summary		Update account credentials
//	@Description	Update the user’s display and preferences.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#update_credentials
//	@Tags			mastodon/accounts
//	@Accept			mpfd
//	@Produce		json
//	@Param			Authorization							header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@param			account[display_name]					formData	string	false	"The display name to use for the profile."
//	@Param			account[note]							formData	string	false	"The account bio."
//	@Param			account[avatar]							formData	file	false	"Avatar image encoded using multipart/form-data"
//	@Param			account[header]							formData	file	false	"Header image encoded using multipart/form-data"
//	@Param			account[locked]							formData	boolean	false	"Whether manual approval of follow requests is required."
//	@Param			account[bot]							formData	boolean	false	"Whether the account has a bot flag."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[fields_attributes][0][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][0][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][1][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][1][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][2][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][2][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][3][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][3][value]	formData	string	false	"hash"
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Success		200										{object}	entities.CredentialAccount
//	@Failure		401										{object}	errorMsg
//	@Failure		403										{object}	errorMsg
//	@Failure		404										{object}	errorMsg
//	@Failure		422										{object}	errorMsg
//	@Failure		500										{object}	errorMsg
//	@Router			/api/v1/accounts/update_credentials [patch]
func UpdateCredentials(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "The access token is invalid"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}

	fe := formExactor{Form: form}

	// mastodon api
	account, err := mastodon.UpdateCredentials(
		authorizationID,
		fe.String("account[display_name]"),    //displayName,
		fe.String("account[note]"),            // 	note,
		fe.URLString("account[avatar]"),       // 	avatar, // todo
		fe.URLString("account[header]"),       // 	header, // todo this is file
		fe.Bool("account[locked]"),            // 	locked,
		fe.Bool("account[bot]"),               // 	bot,
		fe.String("account[discoverable]"),    // 	discoverable,
		fe.Hash("account[fields_attributes]"), // 	fieldsAttributes,
		fe.String("account[display_name]"),    // 	sourcePrivacy, // todo, check key name
		fe.Bool("source[sensitive]"),          // 	sourceSensitive, // todo, check key name
		fe.String("source[language]"),         // 	sourceLanguage) // todo, check key name
	)

	// error handling
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return
	c.JSON(http.StatusOK, account)
}
