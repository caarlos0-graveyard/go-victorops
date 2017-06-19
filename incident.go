package victorops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Transition represents a state changes of an incident
type Transition struct {
	Name    string    `json:",omitempty"`
	At      time.Time `json:",omitempty"`
	Message string    `json:",omitempty"`
}

// Incident represents an incident on victorops
type Incident struct {
	AlertCount        int          `json:",omitempty"`
	CurrentPhase      string       `json:",omitempty"`
	EntityDisplayName string       `json:",omitempty"`
	EntityID          string       `json:",omitempty"`
	EntityState       string       `json:",omitempty"`
	EntityType        string       `json:",omitempty"`
	Host              string       `json:",omitempty"`
	IncidentNumber    string       `json:",omitempty"`
	LastAlertID       string       `json:",omitempty"`
	LastAlertTime     time.Time    `json:",omitempty"`
	Service           string       `json:",omitempty"`
	StartTime         time.Time    `json:",omitempty"`
	PagedTeams        []string     `json:",omitempty"`
	PagedUsers        []string     `json:",omitempty"`
	Transitions       []Transition `json:",omitempty"`
	Error             string       `json:",omitempty"`
}

// IncidentTarget User or EscalationPolicy
type IncidentTarget struct {
	Type string `json:"type,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// IncidentBody request body to create an incident
type IncidentBody struct {
	Summary  string           `json:"summary,omitempty"`
	Details  string           `json:"details,omitempty"`
	Username string           `json:"userName,omitempty"`
	Targets  []IncidentTarget `json:"targets,omitempty"`
}

// Incidents get a list of the currently open, acknowledged and
// recently resolved incidents
func (c *Client) Incidents() (incidents []Incident, err error) {
	resp, err := c.request(http.MethodGet, "incidents", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var result = struct {
		Incidents []Incident
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Incidents, err
}

// CreateIncident Create a new incident.
func (c *Client) CreateIncident(body IncidentBody) (incident Incident, err error) {
	bts, err := json.Marshal(body)
	if err != nil {
		return
	}
	resp, err := c.request(http.MethodPost, "incidents", bytes.NewBuffer(bts))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&incident)
	return
}
