package conflux_chain

import (
	"encoding/json"
	"strconv"
	"time"
)

type Metadata struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	TokenID      string `json:"token_id"`
	AnimationUrl string `json:"animation_url"`
	Date         string `json:"date"`
}

func NewMetadata() *Metadata {
	return new(Metadata)
}

func (p *Metadata) SetName(name string) {
	p.Name = name
}

func (p *Metadata) SetDescription(description string) {
	p.Description = description
}

func (p *Metadata) SetImage(imageURL string) {
	p.Image = imageURL
}

func (p *Metadata) SetAnimationUrl(animationUrl string) {
	p.AnimationUrl = animationUrl
}

func (p *Metadata) SetTokenID(tokenId int64) {
	tokenIdstr := strconv.Itoa(int(tokenId))
	p.TokenID = tokenIdstr
}

func (p *Metadata) SetDate() {
	dateStr := time.Now().Format("2006年01月02日")
	p.Date = string(dateStr)

}

func (p *Metadata) ToJson() string {
	r, _ := json.Marshal(p)
	return string(r)

}
