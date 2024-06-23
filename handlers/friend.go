package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"bytes"
	"io"

	"github.com/gorilla/mux"
	"github.com/hashicorp-demoapp/product-api-go/data"
	"github.com/hashicorp-demoapp/product-api-go/data/model"
	"github.com/hashicorp/go-hclog"
)

// Coffee -
type Friend struct {
	con data.Connection
	log hclog.Logger
}

// NewFriend
func NewFriend(con data.Connection, l hclog.Logger) *Friend {
	return &Friend{con, l}
}

func (c *Friend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Friend")

	vars := mux.Vars(r)

	var friendID *int

	if vars["id"] != "" {
		cId, err := strconv.Atoi(vars["id"])
		if err != nil {
			c.log.Error("Cafe provided could not be converted to an integer", "error", err)
			http.Error(rw, "Unable to list ingredients", http.StatusInternalServerError)
			return
		}
		friendID = &cId
	}

	friends, err := c.con.GetFriends(friendID)
	if err != nil {
		c.log.Error("Unable to get products from database", "error", err)
		http.Error(rw, "Unable to list products", http.StatusInternalServerError)
		return
	}

	var d []byte
	d, err = json.Marshal(friends)
	if err != nil {
		c.log.Error("Unable to convert products to JSON", "error", err)
		http.Error(rw, "Unable to list products", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// CreateCafe creates a new cafe
func (c *Friend) CreateFriend(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafe | CreateFriend")

	var friends []model.Friend

	// 요청 본문을 읽고 출력합니다.
	reqBody, _ := io.ReadAll(r.Body)
	c.log.Info("Request Body", "body", string(reqBody))
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // 요청 본문을 리셋합니다.

	err := json.NewDecoder(r.Body).Decode(&friends)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	c.log.Info("Decoded Body", "body", friends)

	// 단일 카페 객체로 처리
	friend := friends[0]

	createdFriend, err := c.con.CreateFriend(friend)
	if err != nil {
		c.log.Error("Unable to create new cafe", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to create new cafe: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := createdFriend.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert cafe to JSON", "error", err)
		http.Error(rw, "Unable to create new cafe", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

func (c *Friend) UpdateFriend(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Friend | UpdateFriend")

	vars := mux.Vars(r)

	body := model.Friend{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	friendID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("friendID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	friend, err := c.con.UpdateFriend(friendID, body)
	if err != nil {
		c.log.Error("Unable to update friend", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to update friend: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := friend.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert friend to JSON", "error", err)
		http.Error(rw, "Unable to create new friend", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// DeleteFriend deletes a user Friend
func (c *Friend) DeleteFriend(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Cafes | DeleteFriend")

	vars := mux.Vars(r)

	friendID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("cafeID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	err = c.con.DeleteFriend(friendID)
	if err != nil {
		c.log.Error("Unable to delete Friend from database", "error", err)
		http.Error(rw, "Unable to delete Friend", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%s", "Deleted Friend")
}