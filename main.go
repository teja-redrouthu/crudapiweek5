package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "strconv"
)

type Ticket struct {
    ID          int     `json:"id"`
    Passenger   string  `json:"passenger"`
    Origin      string  `json:"origin"`
    Destination string  `json:"destination"`
    Price       float64 `json:"price"`
    Date        string  `json:"date"` // Format: YYYY-MM-DD
}

var (
    tickets  = []Ticket{}
    nextID   = 1
    mu       sync.Mutex
)

func main() {
    http.HandleFunc("/tickets", handleTickets)
    http.HandleFunc("/tickets/", handleTicketByID)

    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func handleTickets(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(tickets)
    case "POST":
        var ticket Ticket
        if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
            http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
            return
        }
        mu.Lock()
        ticket.ID = nextID
        nextID++
        tickets = append(tickets, ticket)
        mu.Unlock()
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(ticket)
    default:
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func handleTicketByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/tickets/"):]

    mu.Lock()
    defer mu.Unlock()
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 || id >= nextID {
        http.Error(w, "Ticket not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
        for _, ticket := range tickets {
            if ticket.ID == id {
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(ticket)
                return
            }
        }
    case "PUT":
        var updatedTicket Ticket
        if err := json.NewDecoder(r.Body).Decode(&updatedTicket); err != nil {
            http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
            return
        }
        for i, ticket := range tickets {
            if ticket.ID == id {
                tickets[i] = updatedTicket
                tickets[i].ID = id // Ensure the ID remains unchanged
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(tickets[i])
                return
            }
        }
    case "DELETE":
        for i, ticket := range tickets {
            if ticket.ID == id {
                tickets = append(tickets[:i], tickets[i+1:]...)
                w.WriteHeader(http.StatusNoContent)
                return
            }
        }
    default:
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
    http.Error(w, "Ticket not found", http.StatusNotFound)
}
