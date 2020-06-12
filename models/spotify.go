package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Spotify struct {
	Id           int    `orm:"column(id);auto"`
	ClientId     string `orm:"column(client_id);size(100)"`
	AccessToken  string `orm:"column(access_token);size(500)"`
	TokenType    string `orm:"column(token_type);size(45);null"`
	ExpiresIn    int    `orm:"column(expires_in);null"`
	RefreshToken string `orm:"column(refresh_token);size(500);null"`
	Scope        string `orm:"column(scope);size(45);null"`
}

func (t *Spotify) TableName() string {
	return "spotify"
}

func init() {
	orm.RegisterModel(new(Spotify))
}

// AddSpotify insert a new Spotify into database and returns
// last inserted Id on success.
func AddSpotify(m *Spotify) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// UpdateSpotify updates Spotify by ClientId and returns error if
// the record to be updated doesn't exist
func UpdateSpotifyByClientId(m *Spotify) (err error) {
	o := orm.NewOrm()
	var updateMap = orm.Params{
		"access_token":  m.AccessToken,
		"token_type":    m.TokenType,
		"expires_in":    m.ExpiresIn,
		"refresh_token": m.RefreshToken,
		"scope":         m.Scope,
	}
	count, err := o.QueryTable(new(Spotify)).Filter("client_id", m.ClientId).Update(updateMap)
	fmt.Println("No of Rows updated = ", count)
	return
}

func GetSpotifyByClientId(clientId string) (v *Spotify, err error) {
	o := orm.NewOrm()
	v = &Spotify{}
	err = o.QueryTable(new(Spotify)).Filter("client_id", clientId).One(v)

	if err == orm.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return v, nil
}

func UpdateIfExistElseInsert(m *Spotify) (err error) {

	v, err := GetSpotifyByClientId(m.ClientId)
	if err != nil {
		fmt.Println("Error in  UpdateIfExistElseInsert ", err)
		return
	}

	if v == nil {
		fmt.Println("Inserting")
		_, err = AddSpotify(m)
		return
	}
	fmt.Println("Updating")
	err = UpdateSpotifyByClientId(m)
	return

}
