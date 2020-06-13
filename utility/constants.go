package utility

const (
	POST_METHOD = "POST"
	GET_METHOD  = "GET"
	CODE_200    = "200"
	CODE_400    = "400"
	SUCCESS     = "Success"
	FAIL        = "Fail"
)

//end points
const (
	AUTHORIZE_END_POINT        = "/authorize"
	CALLBACK_END_POINT         = "/callback"
	GET_ACCESS_TOKEN_END_POINT = "/getToken"
	SPOTIFY_SEARCH             = "/search"
	REDIRECT_URL               = "http://localhost:8080/callback"
)
const (
	COLON           = ":"
	SLASH           = "/"
	EQUALTO         = "="
	QUESTION_MARK   = "?"
	AMPERSAND       = "&"
	UrlSlashReplace = "%2F"
	UrlColonReplace = "%3A"
)
const (
	ConstSpotify           = "spotify"
	ConstClientId          = "client_id"
	ConstClientSecret      = "client_secret"
	ConstCode              = "code"
	ConstFormEncoded       = "application/x-www-form-urlencoded"
	ConstAuthorization     = "Authorization"
	BasicAuthType          = "Basic "
	BearerAuthType         = "Bearer "
	ConstGrantType         = "grant_type"
	ConstAuthorizationCode = "authorization_code"
	ConstRedirectUri       = "redirect_uri"
	ConstContentType       = "Content-Type"
)
