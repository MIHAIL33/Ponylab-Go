package handler

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/websocket"
)

func (h *Handler) getBasePage(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/test.page.gohtml")
}

func (h *Handler) wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://" + r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go h.getAllUniq(conn)
}

type responsePayload struct {
	Uid string `json:"uid"`
	Time string `json:"last_date"`
}

func (h *Handler) getAllUniq(conn *websocket.Conn) {
	for {
		devS, err := h.service.GetAllFromCache()
		if err != nil {
			log.Println("failed loading data from cache")
		}

		var resp []responsePayload
		var temp responsePayload
		for _, val := range *devS {
			temp.Uid = val.UID
			temp.Time = val.CreatedAt.Format(time.RFC3339)
			resp = append(resp, temp)
		}

		sort.Slice(resp, func(i, j int) bool { return resp[i].Uid < resp[j].Uid })

		if err = conn.WriteJSON(resp); err != nil {
			log.Println("failed write json")
		}

		time.Sleep(1 * time.Second)
	}
}

func render(w http.ResponseWriter, content string) {

	templates := []string {
		content,
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
		"templates/base.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}