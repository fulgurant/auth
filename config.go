package auth

type Config struct {
	ClientID     string `kong:"prefix:'oauth2-',help:'Oauth2 client id to use',env:'OAUTH2_CLIENTID'"`
	ClientSecret string `kong:"prefix:'oauth2-',help:'Oauth2 client secret to use',env:'OAUTH2_CLIENT_SECRET'"`
}
