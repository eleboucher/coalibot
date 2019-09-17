package assos

import (
	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Assos(option string, event *utils.Message) bool {
	var assos = `<#C04GT8U3Y|bde>
<#C8C0EGZUY|42genesys> :42genesys:
<#C03BN553Z|42entrepreneurs> :42entrepreneurs:
<#C4LF6DS82|42_ai> :42ai:
<#C03EZET39|42born2music> :musical_note:
<#C4VMBA7P1|association-sans-nom> :asn-party2:
<#C0HTYJR99|esport42> :esport42:
<#C736Z2BNZ|escape-games>
<#C7X7V8HV5|airsoft42>  :ak-42:
<#C4W29G7N2|code_her>  :codeher:
<#C8U8Z88TB|42zen>`
	utils.PostMsg(event, slack.MsgOptionText(assos, false))
	return true
}
