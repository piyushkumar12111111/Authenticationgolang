package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/piyushkumar/authenticationmayursir/models"
)

var (
	mu      sync.Mutex
	reports = make(map[string]models.EmploymentReport)
)

func CreateReport(w http.ResponseWriter, r *http.Request) {
	var report models.EmploymentReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//! Generate a new UUID for the report ID
	uuid := uuid.New().String()
	report.ID = uuid
	report.UpdatedAt = time.Now()

	mu.Lock()
	reports[report.ID] = report
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

	// Maintain the ID of the original report
	updatedReport.ID = report.ID
	reports[id] = updatedReport
	updatedReport.UpdatedAt = time.Now()
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

func GetAllReports(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var reportsSlice []models.EmploymentReport
	for _, report := range reports {
		reportsSlice = append(reportsSlice, report)
	}

	sort.Slice(reportsSlice, func(i, j int) bool {
		return reportsSlice[i].UpdatedAt.After(reportsSlice[j].UpdatedAt)
	})

	json.NewEncoder(w).Encode(reportsSlice)
}
