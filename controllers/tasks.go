package controllers

import "net/http"
import (
	"github.com/Caoimhin89/first-go-app/models"
	"github.com/Caoimhin89/first-go-app/common"
	"encoding/json"
	"github.com/gorilla/mux"
)

var Tasks = new(tasksController)

type tasksController struct{}

var t models.Task

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonValid(w, res, http.StatusCreated)
}
func (tc *tasksController) Get(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.Tasks.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(tasks)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonValid(w, res, http.StatusOK)
}
func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task, err := models.Tasks.FindOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonValid(w, res, http.StatusOK)
}
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	if err := models.Tasks.Update(id, t.Name, t.Desc); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)
}
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.Tasks.DeleteById(id); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)
}
