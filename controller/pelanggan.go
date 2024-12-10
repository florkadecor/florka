package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostDataPelanggan(w http.ResponseWriter, r *http.Request) {
	// Auth
	var respn model.Response
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(r))
	if err != nil {
		respn.Status = "Error : Token Tidak Valid "
		respn.Info = at.GetSecretFromHeader(r)
		respn.Location = "Decode Token Error: " + at.GetLoginFromHeader(r)
		respn.Response = err.Error()
		at.WriteJSON(w, http.StatusForbidden, respn)
		return
	}
	_, err = atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		respn.Status = "Error : User tidak ada di database "
		respn.Info = payload.Alias
		respn.Location = payload.Id
		respn.Response = err.Error()
		at.WriteJSON(w, http.StatusNotFound, respn)
		return
	}
	// Auth
	// Parsing Body
	var pelanggandata model.DataPelanggan
	err = json.NewDecoder(r.Body).Decode(&pelanggandata)
	if err != nil {
		respn.Status = "Error : Parsing data pelanggan gagal"
		respn.Info = payload.Alias
		respn.Location = payload.Id
		respn.Response = err.Error()
		at.WriteJSON(w, http.StatusNotFound, respn)
		return
	}
	//input database
	id, err := atdb.InsertOneDoc(config.Mongoconn, "pelanggan", pelanggandata)
	if err != nil {
		respn.Status = "Error : gagal input database"
		respn.Response = err.Error()
		at.WriteJSON(w, http.StatusNotFound, respn)
		return
	}
	pelanggandata.ID = id
	at.WriteJSON(w, http.StatusOK, pelanggandata)
}