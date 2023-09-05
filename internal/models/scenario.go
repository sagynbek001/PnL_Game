package models

import (
	"context"
	"time"
)

// scenarios ...

type Scenario struct {
	ID         int64
	OwnerID    int64
	Name       string
	StepsCount int64
	DeletedAt  *time.Time
	Data       ScenarioData
}
type ScenarioData struct {
	Managers []Manager `json:"managers"`
	Projects []Project `json:"projects"`
}

type Manager struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	UsersID   int64      `json:"user_id"`
	Employees []Employee `json:"employees"`
	Events    []Event    `json:"events"`
}

type Event struct {
	ID          int64  `json:"id"`
	EventTypeID int64  `json:"event_type_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Employee struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Seniority      int64   `json:"seniority"`
	Salary         float64 `json:"salary"`
	ProjectID      int64   `json:"project_id"`
	EmployeeStatus string  `json:"employee_status"`
}

type Project struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Rates []Rate `json:"rates"`
}

type Rate struct {
	ID                   int64   `json:"id"`
	ProjectID            int64   `json:"project_id"`
	Type                 string  `json:"type"`
	Seniority            int64   `json:"seniority"`
	Rate                 float64 `json:"rate"`
	IllCompensation      float64 `json:"ill_compensation"`
	VacationCompensation float64 `json:"vacation_compensation"`
}

// scenariosRepo represents scenarios repository contract
type ScenarioUsecase interface {
	Create(ctx context.Context, scenario *Scenario) error
	GetAll(ctx context.Context) ([]Scenario, error)
	GetByID(ctx context.Context, id uint64) (Scenario, error)
	Update(ctx context.Context, scenarios *Scenario) error
	Delete(ctx context.Context, id uint64) error
}

// scenariosRepo represents scenarios repository contract
type ScenarioRepo interface {
	Create(ctx context.Context, scenarios *Scenario) error
	GetAll(ctx context.Context) ([]Scenario, error)
	GetByID(ctx context.Context, id uint64) (Scenario, error)
	Update(ctx context.Context, scenarios *Scenario) error
	Delete(ctx context.Context, id uint64) error
}
