package Assos

import (
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Assos(option string, event *Struct.Message) bool {
	var assos = `<#bde>
<#42genesys> :42genesys:
<#42entrepreneurs> :42entrepreneurs:
<#42_ai> :42ai:
<#42born2music> :musical_note:
<#association-sans-nom> :asn-party2:
<#esport42> :esport42:
<#escape-games>
<#airsoft42>  :ak-42:
<#code_her>  :codeher:
<#42zen>`
	event.API.PostMessage(event.Channel, slack.MsgOptionText(assos, false))
	return true
}
