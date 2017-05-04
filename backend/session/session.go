package session

import (
	"io/ioutil"

	"github.com/gorilla/sessions"
)

var (
	// GlobalSession for session object
	GlobalSession *sessions.Session
	// UniversalSessionStore for creating a new cookie
	UniversalSessionStore = sessions.NewCookieStore([]byte("f?:<@-r'lojerfipehsfi-esoeg5reoushlg/zhresi-rygte6rtf7syuftgsuyfeusj,gfjyergfjhr^gvhnmzm bvzn,^>//bds,grbfxdjhvm,ba[wreotie$#eo49ut@3u84t9u85t459ty4uoty9urgofe9ty5wgdhbfcudkb]"))
)

var (
	// SigningKey to be used
	SigningKey = "Hello Hexerent"
	// PrivateKey to be used
	PrivateKey []byte
	// PublicKey to be used
	PublicKey []byte
)

func init() {
	PrivateKey, _ = ioutil.ReadFile("../Keys/Hexerent_Private_Key.rsa")
	PublicKey, _ = ioutil.ReadFile("../Keys/Hexerent_Public_Key.rsa")
}
