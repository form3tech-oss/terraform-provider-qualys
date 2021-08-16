package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/google/uuid"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/gorilla/mux"
)

type connectorState map[string]*gcp.Connector

func testServer() *httptest.Server {
	connectors := make(connectorState)

	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("%s/{connector_id}", gcp.ApiPath), getConnectorHandlerFunc(connectors)).Methods("GET")
	router.HandleFunc(gcp.ApiPath, createConnectorHandlerFunc(connectors)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{connector_id}", gcp.ApiPath), updateConnectorHandlerFunc(connectors)).Methods("PUT")
	router.HandleFunc(gcp.ApiPath, deleteConnectorHandlerFunc(connectors)).Methods("DELETE")

	ts := httptest.NewServer(router)

	must(os.Setenv("QUALYS_URL", ts.URL))
	must(os.Setenv("QUALYS_USERNAME", "foo"))
	must(os.Setenv("QUALYS_PASSWORD", "bar"))

	return ts
}

func getConnectorHandlerFunc(state connectorState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["connector_id"]

		conn, ok := state[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			_, err := fmt.Fprint(w, "Not Found")
			must(err)
			return
		}

		bs, err := json.Marshal(conn)
		must(err)

		w.Header().Set("Content-Type", "application/json")
		_, err = fmt.Fprint(w, string(bs))
		must(err)
	}
}

func createConnectorHandlerFunc(state connectorState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := gcp.Connector{
			ConnectorID: uuid.New().String(),
			Description: r.FormValue("description"),
			Name:        r.FormValue("name"),
			Project:     r.FormValue("projectId"),
			Provider:    "gcp",
		}
		state[conn.ConnectorID] = &conn

		bs, err := json.Marshal(conn)
		must(err)

		w.Header().Set("Content-Type", "application/json")
		_, err = fmt.Fprint(w, string(bs))
		must(err)
	}
}

func updateConnectorHandlerFunc(state connectorState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["connector_id"]

		conn := gcp.Connector{}
		dec := json.NewDecoder(r.Body)
		must(dec.Decode(&conn))
		state[id] = &conn

		bs, err := json.Marshal(conn)
		must(err)

		_, err = fmt.Fprint(w, string(bs))
	}
}

func deleteConnectorHandlerFunc(state connectorState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ids gcp.ConnectorIds
		dec := json.NewDecoder(r.Body)
		must(dec.Decode(&ids))

		for _, id := range ids.ConnectorIds {
			delete(state, id)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
