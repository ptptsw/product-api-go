package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ptptsw/product-api-go/data"
	"github.com/ptptsw/product-api-go/data/model"
	"github.com/hashicorp/go-hclog"
)

// Food -
type Food struct {
	con data.Connection
	log hclog.Logger
}

// NewFood -
func NewFood(con data.Connection, l hclog.Logger) *Food {
	return &Food{con, l}
}

func (c *Food) ServeHTTP(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Food | unknown", "path", r.URL.Path)
	http.NotFound(rw, r)
}

// GetUserFoods gets all user foods for a specific user
func (c *Food) GetUserFoods(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Foods | GetUserFoods")

	foods, err := c.con.GetFoods(userID, nil)
	if err != nil {
		// c.log.Error("Unable to get food from database", "error", err)
		http.Error(rw, "Unable to get food from database", http.StatusInternalServerError)
		http.Error(rw, "Unable to list foods", http.StatusInternalServerError)
		return
	}

	d, err := foods.ToJSON()
	if err != nil {
		// c.log.Error("Unable to convert foods to JSON", "error", err)
		http.Error(rw, "Unable to convert foods to JSON", http.StatusInternalServerError)
		http.Error(rw, "Unable to get food from database", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// CreateFood creates a new food
func (c *Food) CreateFood(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Foods | CreateFood")

	body := []model.FoodItems{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	food, err := c.con.CreateFood(userID, body)
	if err != nil {
		c.log.Error("Unable to create new food", "error", err)
		http.Error(rw, "Unable to create new food", http.StatusInternalServerError)
		return
	}

	d, err := food.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert food to JSON", "error", err)
		http.Error(rw, "Unable to create new food", http.StatusInternalServerError)
	}

	rw.Write(d)
}

// GetUserFood gets a specific user food
func (c *Food) GetUserFood(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Foods | GetUserFood")

	vars := mux.Vars(r)

	foodID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("foodID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to list food", http.StatusInternalServerError)
		return
	}

	foods, err := c.con.GetFoods(userID, &foodID)
	if err != nil {
		c.log.Error("Unable to get food from database", "error", err)
		http.Error(rw, "Unable to list food", http.StatusInternalServerError)
		return
	}

	food := model.Food{}

	if len(foods) > 0 {
		food = foods[0]
	}

	d, err := food.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert foods to JSON", "error", err)
		http.Error(rw, "Unable to list food", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

// UpdateFood updates an food
func (c *Food) UpdateFood(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Foods | UpdateFood")

	// Get foodID
	vars := mux.Vars(r)
	foodID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("foodID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to update food", http.StatusInternalServerError)
		return
	}

	body := []model.FoodItems{}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	food, err := c.con.UpdateFood(userID, foodID, body)
	if err != nil {
		c.log.Error("Unable to create new food", "error", err)
		http.Error(rw, "Unable to update food", http.StatusInternalServerError)
		return
	}

	d, err := food.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert food to JSON", "error", err)
		http.Error(rw, "Unable to update food", http.StatusInternalServerError)
	}

	rw.Write(d)
}

// DeleteFood deletes a user food
func (c *Food) DeleteFood(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Foods | DeleteFood")

	vars := mux.Vars(r)

	foodID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("foodID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete food", http.StatusInternalServerError)
		return
	}

	err = c.con.DeleteFood(userID, foodID)
	if err != nil {
		c.log.Error("Unable to delete food from database", "error", err)
		http.Error(rw, "Unable to delete food", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%s", "Deleted food")
}
