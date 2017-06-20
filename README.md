# go-victorops

Unofficial VictorOps API for Golang

## Usage

```go
var cli = victorops.New("login", "api-id", "api-key")
incidents, _ := cli.Incidents()
fmt.Println(incidents)
```

## Current state

### On-call

- [ ] GET /api-public/v1/user/{user}/oncall/schedule
- [ ] GET /api-public/v1/team/{team}/oncall/schedule
- [ ] PATCH /api-public/v1/team/{team}/oncall/user

### Incidents

- [x] GET /api-public/v1/incidents
- [x] POST /api-public/v1/incidents
- [x] PATCH /api-public/v1/incidents/ack
- [x] PATCH /api-public/v1/incidents/resolve
- [ ] PATCH /api-public/v1/incidents/byUser/ack
- [ ] PATCH /api-public/v1/incidents/byUser/resolve

### Alerts

- [ ] GET /api-public/v1/alerts/{uuid}

### Reporting

- [ ] GET /api-reporting/v1/team/{team}/oncall/log
- [ ] GET /api-reporting/v1/incidents
- [ ] GET /api-reporting/v2/incidents

### Users

- [ ] GET /api-public/v1/user
- [ ] POST /api-public/v1/user
- [ ] DELETE /api-public/v1/user/{user}
- [ ] GET /api-public/v1/user/{user}
- [ ] PUT /api-public/v1/user/{user}

### User Contact Methods
- [ ] GET /api-public/v1/user/{user}/contact-methods
- [ ] GET /api-public/v1/user/{user}/contact-methods/devices
- [ ] DELETE /api-public/v1/user/{user}/contact-methods/devices/{contactId}
- [ ] GET /api-public/v1/user/{user}/contact-methods/devices/{contactId}
- [ ] PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}
- [ ] GET /api-public/v1/user/{user}/contact-methods/emails
- [ ] POST /api-public/v1/user/{user}/contact-methods/emails
- [ ] DELETE /api-public/v1/user/{user}/contact-methods/emails/{contactId}
- [ ] GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}
- [ ] PUT /api-public/v1/user/{user}/contact-methods/emails/{contactId}
- [ ] GET /api-public/v1/user/{user}/contact-methods/phones
- [ ] POST /api-public/v1/user/{user}/contact-methods/phones
- [ ] DELETE /api-public/v1/user/{user}/contact-methods/phones/{contactId}
- [ ] GET /api-public/v1/user/{user}/contact-methods/phones/{contactId}
- [ ] PUT /api-public/v1/user/{user}/contact-methods/phones/{contactId}

### User Paging Policies

- [ ] GET /api-public/v1/user/{user}/policies

### Routing keys

- [ ] GET /api-public/v1/org/routing-keys

### Teams
- [ ] GET /api-public/v1/team
- [ ] POST /api-public/v1/team
- [ ] DELETE /api-public/v1/team/{team}
- [ ] GET /api-public/v1/team/{team}
- [ ] PUT /api-public/v1/team/{team}
- [ ] GET /api-public/v1/team/{team}/members
- [ ] POST /api-public/v1/team/{team}/members
- [ ] DELETE /api-public/v1/team/{team}/members/{user}


## F.A.Q.

- Will all the endpoints be implemented?
  Dunno! I implemented first the endpoints I needed for [victorops-notifier](https://github.com/caarlos0/victorops-notifier)
