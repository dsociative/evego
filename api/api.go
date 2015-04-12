package api

import "net/http"
import "net/url"
import "io/ioutil"
import "github.com/dsociative/evego/parser"

type APIFace interface {
	Characters() []parser.Character
}

type API struct {
	vcode string
	keyid string
}

func (api *API) Request(urlPrefix string, values url.Values) []byte {
	api_host := "https://api.eveonline.com"
	resp, err := http.PostForm(api_host+urlPrefix, values)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	} else {
		panic(err)
	}
}

func (api *API) Do(urlPrefix string, values url.Values, parseFunc func(raw []byte) parser.Model) parser.Model {
	return parseFunc(api.Request(urlPrefix, values))
}

func (api *API) SkillTree() parser.Tree {
	return api.Do(
		"/eve/SkillTree.xml.aspx", url.Values{}, parser.ParseSkillTree,
	).(parser.Tree)
}

func (api API) Characters() []parser.Character {
	values := url.Values{"keyID": {api.keyid}, "vCode": {api.vcode}}
	return parser.ParseCharacters(
		api.Request("/account/Characters.xml.aspx", values),
	)
}

func (api *API) Queue(character *parser.Character) parser.SkillQueue {
	values := url.Values{
		"keyID":       {api.keyid},
		"vCode":       {api.vcode},
		"characterID": {character.CharacterID},
	}
	return parser.ParseSkillQueue(
		api.Request("/char/SkillQueue.xml.aspx", values),
	)
}

func New(vcode string, keyid string) API {
	return API{vcode, keyid}
}
