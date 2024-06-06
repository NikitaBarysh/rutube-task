package handler

import (
	"encoding/json"
	"net/http"

	"rutube-task/internal/entity"
)

func (h *Handler) setEmployees(rw http.ResponseWriter, r *http.Request) {
	employeesList := make([]entity.Employee, 0)

	if err := json.NewDecoder(r.Body).Decode(&employeesList); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.EmployeeServiceInterface.SetEmployeeList(r.Context(), employeesList)
	if err != nil {
		http.Error(rw, "Err to add, check your request and try again", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) getEmployees(rw http.ResponseWriter, r *http.Request) {
	list, err := h.service.EmployeeServiceInterface.GetEmployee(r.Context())
	if err != nil {
		http.Error(rw, "Err to get, try again", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(list)
	if err != nil {
		http.Error(rw, "Err to decode list with employees", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
