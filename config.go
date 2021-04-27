package auth

// Config values for authentication
type Config struct {
	ProviderUrl      string `kong:"prefix='oauth2-',help='Url of oauth2 provider',alias='oidc',env='OAUTH2_PROVIDER_URL'"`
	ClientID         string `kong:"prefix='oauth2-',help='Oauth2 client id to use',env='OAUTH2_CLIENTID',alias='cid'"`
	ClientSecret     string `kong:"prefix='oauth2-',help='Oauth2 client secret to use',env='OAUTH2_CLIENT_SECRET',alias='cis'"`
	LoginCallbackUrl string `kong:"prefix='oauth2-',help='Redirect url after successful login',env='OAUTH2_LOGIN_CALLBACK_URL'"`
}
