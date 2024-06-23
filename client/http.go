package client

import (
	"fmt"
	"log"
	"net/http"

	hckit "github.com/hashicorp-demoapp/go-hckit"
	"github.com/hashicorp-demoapp/product-api-go/data/model"
)

// HTTP contains all client details
type HTTP struct {
	client  *http.Client
	baseURL string
}

// NewHTTP creates a new HTTP client
func NewHTTP(baseURL string) *HTTP {
	c := &http.Client{Transport: hckit.TracingRoundTripper{Proxied: http.DefaultTransport}}
	return &HTTP{c, baseURL}
}

/ GetCafes retrieves a list of coffees
func (h *HTTP) GetFriends() ([]model.Friend, error) {
	log.Print("INFO: Executing GetFriends")
	resp, err := h.client.Get(fmt.Sprintf("%s/friends", h.baseURL))
	if err != nil {
		return nil, err
	}

	friends := model.Friends{}
	friends.FromJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

// GetCoffee retrieves a single coffee
func (h *HTTP) GetFriend(friendID int) (*model.Friend, error) {
	resp, err := h.client.Get(fmt.Sprintf("%s/friends/%d", h.baseURL, friendID))
	if err != nil {
		return nil, err
	}

	friend := model.Friend{}
	err = friend.FromJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return &friend, nil
}



// GetCoffees retrieves a list of coffees
func (h *HTTP) GetCoffees() ([]model.Coffee, error) {
	log.Print("INFO: Executing GetCoffees")
	resp, err := h.client.Get(fmt.Sprintf("%s/coffees", h.baseURL))
	if err != nil {
		return nil, err
	}

	coffees := model.Coffees{}
	coffees.FromJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// GetCoffee retrieves a single coffee
func (h *HTTP) GetCoffee(coffeeID int) (*model.Coffee, error) {
	resp, err := h.client.Get(fmt.Sprintf("%s/coffees/%d", h.baseURL, coffeeID))
	if err != nil {
		return nil, err
	}

	coffee := model.Coffee{}
	err = coffee.FromJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return &coffee, nil
}

// GetIngredientsForCoffee retrieves a list of ingredients that go into a particular coffee
func (h *HTTP) GetIngredientsForCoffee(coffeeID int) ([]model.Ingredient, error) {
	resp, err := h.client.Get(fmt.Sprintf("%s/coffees/%d/ingredients", h.baseURL, coffeeID))
	if err != nil {
		return nil, err
	}

	ingredients := model.Ingredients{}
	err = ingredients.FromJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}
