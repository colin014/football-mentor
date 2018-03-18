package model

import "github.com/jinzhu/gorm"

type Club struct {
	gorm.Model            `json:"-"`
	Name           string `json:"name" binding:"required"`
	LogoUrl        string `json:"logo_url" binding:"required"`
	LeagueName     string `json:"league_name" binding:"required"`
	LeagueLogoUrl  string `json:"league_logo_url" binding:"required"`
	CurrentPlace   int    `json:"current_place"`
	StadiumName    string `json:"stadium_name"`
	FacebookPageId string `json:"facebook_page_id"`
	Website        string `json:"website"`
}

func (Club) TableName() string {
	return "club"
}

func (c *Club) Update(r *Club) {
	c.Name = r.Name
	c.LogoUrl = r.LogoUrl
	c.LeagueName = r.LeagueName
	c.LeagueLogoUrl = r.LeagueLogoUrl
	c.CurrentPlace = r.CurrentPlace
	c.StadiumName = r.StadiumName
	c.FacebookPageId = r.FacebookPageId
	c.Website = r.Website
}
