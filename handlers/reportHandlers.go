package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/authenticationmayursir/models"
)

var (
	mu      sync.Mutex
	reports = map[string]models.EmploymentReport{
		"1": {ID: "1", Content: "Annual Report 2021"},
		"2": {ID: "2", Content: "Quarterly Diversity Report Q2"},
		"3": {ID: "3", Content: "Employee Satisfaction Report 2022"},
	}
	nextID = 4 
)


func CreateReport(w http.ResponseWriter, r *http.Request) {
	var report models.EmploymentReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	report.ID = strconv.Itoa(nextID) 
	reports[report.ID] = report
	nextID++ 
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(report)
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	report, ok := reports[id]
	mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(report)
}

func UpdateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var updatedReport models.EmploymentReport
	if err := json.NewDecoder(r.Body).Decode(&updatedReport); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	report, exists := reports[id]
	if !exists {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}

	
	updatedReport.ID = report.ID
	reports[id] = updatedReport
	mu.Unlock()

	json.NewEncoder(w).Encode(updatedReport)
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	delete(reports, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
