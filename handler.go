package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
	"github.com/whatsauth/ws"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var WAKeyword = "wh4t5auth0"
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if ws.IsLoginRequest(msg, WAKeyword) { //untuk whatsauth request login
			dt := &ws.WhatsauthRequest{
				Uuid:        ws.GetUUID(msg, WAKeyword),
				Phonenumber: msg.Phone_number,
				Delay:       msg.From_link_delay,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/whatsauth/request")
		} else { //untuk membalas pesan masuk
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: "HELLO!" + msg.Alias_name + "Ini adalah Tutorial Untuk penggunaan whatsauth ini. Terima Kasih" + "Link :",
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
