package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"github.com/grafov/bcast"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Topic  string    `json:"topic"`
	Sender string    `json:"sender"`
	Msg    string    `json:"message"`
	Date   time.Time `json:"-"`
}

func echoHandler(ws *websocket.Conn) {

	for {

		var message Message
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			log.Println("Unable to read message", err)
		}

		message.Sender = "Server"
		err = websocket.JSON.Send(ws, message)
		if err != nil {
			log.Println("Unable to send message", err)
		}
	}
}

func GetJson(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)["jsonContent"]
	tree := Node{
		Text: "Parent1",
		Href: "json/parent_1.json",
		Nodes: []*Node{
			&Node{
				Text: "Child 1",
				Href: "json/child_1.json",
				Nodes: []*Node{
					&Node{
						Text: "Grandchild 1",
					},
					&Node{
						Text: "Grandchild 2",
					},
				},
			},
			&Node{
				Text: "Parent2",
				Href: "json/parent_2.json",
			},
		},
	}
	switch q {
	case "tree.json":
		if err := json.NewEncoder(w).Encode(tree); err != nil {
			log.Println(err)
		}
	default:
		// strip the query
		type temp struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Text        string `json:"full_text"`
		}
		var result temp
		result = temp{
			Name:        "the name of the object",
			Description: "a description with _some markdown_ __inside__",
			Text:        "# Pressum manu \n\n## Nullos tecta\n\nLorem markdownum Achillem esses veretur superi non venisse labor mihi Tmolus\nomnes; sui. Me ad fuisse sidera; *pectora* alta candidus cubito oravere anxia\nfaciemque pectora, bos. Tua ille labitur delamentatur saxum lacrimasque\nrelinquit amplecti si mendaci praeteriit nobis ad euntem, tot uno.\n\n## Culpae testatus\n\nSi Thracia crimina voluisti, et prole. Tui non fallaciter crinis furit fluctus\nMavortia: sed siquid odium meae pro. Fecit protectus, et lacrimas illis\nMyrmidonasque ante ianua, toto mente videndo evinctus victricemque portus et\nconscia color, Astyanax.\n\nNervis qua iram eburno haeret offensus, veluti *videt*; quid nubes lacerto,\ngrave in ora desilit supposuique inquit. Quae aris rostris Saturnius.\n\n## Frondibus sit admovit tauri\n\nEst exul aries coloribus sagitta Nyctimene proles: et: obruor. Thymo anus: veste\nquae genetrici oret nostro patri et hic fit radice audire! Quae te quis petiit:\nalter usus agitant ponere, tot. Pyrois capulo, Phorbas hostes est; non venerat\nquod. Quaeque et eris tuetur hic conplevit quasque.\n\n> Phoebus quod ne virgo sideribus, memoratis *Epopeus et* culmine Mycenae\n> inscribit saepe saepe. Transtra protectum conspecta, hostes. Nodum fraterque\n> fatum tanget agmina cuius, ab cum lapis me vulgusque, mea dederat.\n\nIpse deo sanior verum omnia sunt area amores unde: cum litusque moram, fido, nec\nposuit, tot. *Actaeonis ab* acumine tanto *vidi* huic inania nostraque adspice\nfulicisque ille utque acumine meo, ore recuset inpune. Qui nec cibus, ab qui.\nEst dimittit equos; oraque, occidit, ipse nec rostrum scire nostrumque *et*.\nNegare silvasque.",
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Println(err)
		}
	}
}
