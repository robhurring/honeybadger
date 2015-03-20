package honeybadger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Faults struct {
	Pagination
	Results []*Fault
}

type FaultNotices struct {
	Pagination
	Results []*FaultNotice
}

type Projects struct {
	Pagination
	Results []*Project
}

type Teams struct {
	Pagination
	Results []*Team
}

type TeamInvitations struct {
	Pagination
	Results []*TeamInvitation
}

type Deploys struct {
	Pagination
	Results []*Deploy
}

type Deploy struct {
	Env           string     `json:"environment"`
	Revision      string     `json:"revision"`
	Repository    string     `json:"repository"`
	LocalUsername string     `json:"local_username"`
	Url           string     `json:"url"`
	CreatedAt     *time.Time `json:"created_at"`
}

type Team struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Owner     *User      `json:"owner"`
	Members   []*User    `json:"members"`
	Projects  []*Project `json:"project"`
	CreatedAt *time.Time `json:"created_at"`
}

type TeamInvitation struct {
	Id         int        `json:"id"`
	Token      string     `json:"token"`
	Email      string     `json:"email"`
	Admin      bool       `json:"admin"`
	Message    string     `json:"message"`
	CreatedBy  *User      `json:"created_by"`
	AcceptedBy *User      `json:"accepted_by"`
	AcceptedAt *time.Time `json:"accepted_at"`
	CreatedAt  *time.Time `json:"created_at"`
}

type Project struct {
	Id                   int            `json:"id"`
	TeamId               int            `json:"team_id"`
	Name                 string         `json:"name"`
	Token                string         `json:"token"`
	Active               bool           `json:"active"`
	DisablePublicLinks   bool           `json:"disable_public_links"`
	GithubProject        string         `json:"github_project"`
	PivotalProjectId     string         `json:"pivotal_project_id"`
	AsanaWorkspaceId     string         `json:"asana_workspace_id"`
	Owner                *User          `json:"owner"`
	Users                []*User        `json:"users"`
	UnresolvedFaultCount int            `json:"unresolved_fault_count"`
	FaultCount           int            `json:"fault_count"`
	Sites                []interface{}  `json:"sites"`
	Environments         []*Environment `json:"environments"`
	CreatedAt            *time.Time     `json:"created_at"`
	LastNoticeAt         *time.Time     `json:"last_notice_at"`
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Environment struct {
	Id            int        `json:"id"`
	ProjectId     int        `json:"project_id"`
	Name          string     `json:"name"`
	Notifications bool       `json:"notifications"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type Fault struct {
	Id            int    `json:"id"`
	ProjectId     int    `json:"project_id"`
	Klass         string `json:"klass"`
	Component     string `json:"component"`
	Action        string `json:"string"`
	Env           string `json:"string"`
	Resolved      bool   `json:"resolved"`
	Ignored       bool   `json:"ignored"`
	CommentsCount int    `json:"comments_count"`
	Message       string `json:"message"`
	NoticesCount  int    `json:"notice_count"`
	Tags          []string
	Assignee      string     `json:"assignee"`
	CreatedAt     *time.Time `json:"created_at"`
	LatNoticeAt   *time.Time `json:"last_notice_at"`
}

type FaultNotice struct {
	Id        int              `json:"id"`
	FaultId   int              `json:"fault_id"`
	CreatedAt time.Time        `json:"created_at"`
	Message   string           `json:"message"`
	Request   interface{}      `json:"request"`
	Env       interface{}      `json:"environment"`
	WebEnv    interface{}      `json:"web_environment"`
	Backtrace []*BacktraceLine `json:"backtrace"`
}

type BacktraceLine struct {
	Number string `json:"number"`
	File   string `json:"file"`
	Method string `json:"method"`
}

type Pagination struct {
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	NumPages    int `json:"num_pages"`
}

type Honeybadger struct {
	BaseUrl    string
	ApiVersion string
	Token      string
	Client     *http.Client
}

type Params map[string]string

func (p Params) Query() string {
	var buffer bytes.Buffer

	for k, v := range p {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(url.QueryEscape(v))
		buffer.WriteString("&")
	}

	return strings.TrimRight(buffer.String(), "&")
}

func New(authToken string) *Honeybadger {
	return &Honeybadger{
		BaseUrl:    "https://api.honeybadger.io",
		ApiVersion: "v1",
		Token:      authToken,
		Client:     &http.Client{},
	}
}

func (h *Honeybadger) get(path string, params Params, response interface{}) (err error) {
	requestParams := Params{"auth_token": h.Token}

	if params != nil {
		for k, v := range params {
			requestParams[k] = v
		}
	}

	url := h.BaseUrl + "/" + h.ApiVersion + path
	url += "?" + requestParams.Query()

	req, err := http.NewRequest("GET", url, nil)
	resp, err := h.Client.Do(req)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, response)

	return
}

func (h *Honeybadger) Projects() (result *Projects, err error) {
	result = &Projects{}
	err = h.get("/projects", nil, result)
	return
}

func (h *Honeybadger) Project(projectId int) (result *Project, err error) {
	path := fmt.Sprintf("/projects/%d", projectId)
	result = &Project{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) Teams() (result *Teams, err error) {
	path := "/teams"
	result = &Teams{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) Team(teamId int) (result *Team, err error) {
	path := fmt.Sprintf("/teams/%d", teamId)
	result = &Team{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) TeamInvitations(teamId int) (result *TeamInvitations, err error) {
	path := fmt.Sprintf("/teams/%d/team_invitations", teamId)
	result = &TeamInvitations{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) TeamInvitation(teamId, teamInvitationId int) (result *TeamInvitation, err error) {
	path := fmt.Sprintf("/teams/%d/team_invitations/%d", teamId, teamInvitationId)
	result = &TeamInvitation{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) Deploys(projectId int) (result *Deploys, err error) {
	path := fmt.Sprintf("/projects/%d/deploys", projectId)
	result = &Deploys{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) Faults(projectId int) (result *Faults, err error) {
	path := fmt.Sprintf("/projects/%d/faults", projectId)
	result = &Faults{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) Fault(projectId, faultId int) (result *Fault, err error) {
	path := fmt.Sprintf("/projects/%d/faults/%d", projectId, faultId)
	result = &Fault{}
	err = h.get(path, nil, result)

	return
}

func (h *Honeybadger) FaultNotices(projectId, faultId int) (result *FaultNotices, err error) {
	path := fmt.Sprintf("/projects/%d/faults/%d/notices", projectId, faultId)
	result = &FaultNotices{}
	err = h.get(path, nil, result)

	return
}

func debug(obj ...interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}
