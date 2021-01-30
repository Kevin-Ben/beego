package oauth

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"net/url"
	"sogo/models"
)

// SendController operations for Send
type GithubController struct {
	beego.Controller
}

func GetTokenAuthUrl(code string, ClientID string, Secret string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		ClientID, Secret, code,
	)
}

type Token struct {
	access_token string
	token_type   string
	scope        string
}

func GetAccessToken(httpUrl string) string {
	resp, _ := http.Get(httpUrl)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	data, _ := url.ParseQuery(string(body))

	return data.Get("access_token")
}

func GetUserInfo(AccessToken string) []byte {

	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", AccessToken))

	var client = http.Client{}
	var res *http.Response
	res, _ = client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	return body

}

type GitHubUser struct {
	id int64 `json:"id"`
}

func JsonToMap(jsonStr string) (m map[string]interface{}, err error) {
	a := make(map[string]interface{})
	unmarsha1Err := json.Unmarshal([]byte(jsonStr), &a)
	if unmarsha1Err != nil {
		return nil, unmarsha1Err
	}
	return a, nil
}

func (c *GithubController) Get() {

	ClientID, _ := beego.AppConfig.String("github::ClientID")
	RedirectURL, _ := beego.AppConfig.String("github::RedirectURL")
	authCode := c.GetString("code")
	authToken := c.GetString("token")

	if authCode == "" && authToken == "" {
		httpUrl := "https://github.com/login/oauth/authorize?client_id=" + ClientID + "&redirect_uri=" + RedirectURL + "&scope=user"
		c.Redirect(httpUrl, 302)
		return
	}

	if authCode != "" {
		Secret, _ := beego.AppConfig.String("github::Secret")

		httpUrl := GetTokenAuthUrl(authCode, ClientID, Secret)

		token := GetAccessToken(httpUrl)

		c.Ctx.WriteString(token)
	}

	if authToken != "" {
		UserInfo := GetUserInfo(authToken)
		HashMap, _ := JsonToMap(string(UserInfo))
		gitHubId := int64(HashMap["id"].(float64))

		fmt.Println("gitHubId: " + string(gitHubId))
		user, _ := models.ReadOrCreate(gitHubId)
		c.Data["json"] = user
		c.ServeJSON()
	}

}
