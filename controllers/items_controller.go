package controllers

import (
	"encoding/json"
	"github.com/voicurobert/bookstore_items-api/domain/items"
	"github.com/voicurobert/bookstore_items-api/services"
	"github.com/voicurobert/bookstore_items-api/utils/http_utils"
	"github.com/voicurobert/bookstore_oauth-go/oauth"
	"github.com/voicurobert/bookstore_utils-go/rest_errors"
	"io/ioutil"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//http_utils.RespondError(w, err)
		return
	}
	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.RespondError(w, *rest_errors.NewBadRequestError("invalid request body"))
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		http_utils.RespondError(w, *rest_errors.NewBadRequestError("invalid item json body"))
		return
	}
	itemRequest.Seller = oauth.GetCallerID(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
