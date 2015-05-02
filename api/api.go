package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	HOST       = "https://api.eveonline.com"
	TREE       = "/eve/SkillTree.xml.aspx"
	CHARACTERS = "/account/Characters.xml.aspx"
	QUEUE      = "/char/SkillQueue.xml.aspx"
	KILLS      = "/char/KillLog.xml.aspx"
)

type APIFace interface {
	Characters() ([]Character, error)
}

type API struct {
	vcode string
	keyid string
}

func (api *API) Request(urlPrefix string, values url.Values, model Model) (err error) {
	var resp *http.Response

	if resp, err = http.PostForm(HOST+urlPrefix, values); err == nil {
		defer resp.Body.Close()
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err == nil {
			err = Parse(body, model)
		}
	}
	return err
}

func (api *API) SkillTree() (tree Tree, err error) {
	err = api.Request(TREE, url.Values{}, &tree)
	return tree, err
}

func (api API) Characters() ([]Character, error) {
	characters := Characters{}
	err := api.Request(
		CHARACTERS,
		url.Values{"keyID": {api.keyid}, "vCode": {api.vcode}},
		&characters,
	)
	return characters.Character, err
}

func (api *API) Queue(character *Character) (queue SkillQueue, err error) {
	err = api.Request(
		QUEUE,
		url.Values{
			"keyID":       {api.keyid},
			"vCode":       {api.vcode},
			"characterID": {character.CharacterID},
		},
		&queue,
	)
	return queue, err
}

func (api *API) KillLog(character *Character) (kills Kills, err error) {
	err = api.Request(
		KILLS,
		url.Values{
			"keyID":       {api.keyid},
			"vCode":       {api.vcode},
			"characterID": {character.CharacterID},
		},
		&kills,
	)
	return kills, err
}

func New(vcode string, keyid string) API {
	return API{vcode, keyid}
}
