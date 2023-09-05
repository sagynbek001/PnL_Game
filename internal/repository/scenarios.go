package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type ScenarioRepo struct {
	conn *pgx.Conn
}

func NewScenarioRepo(conn *pgx.Conn) models.ScenarioRepo {
	return &ScenarioRepo{conn: conn}
}

func (u *ScenarioRepo) Create(ctx context.Context, scenario *models.Scenario) error {
	query := `INSERT INTO scenarios (name, owner_id, steps_count, deleted_at, data) VALUES ($1, $2, $3, $4, $5)`
	_, err := u.conn.Exec(ctx, query, scenario.Name, scenario.OwnerID, scenario.StepsCount, scenario.DeletedAt, scenario.Data)
	if err != nil {
		return err
	}

	return nil
}

func (u *ScenarioRepo) GetAll(ctx context.Context) ([]models.Scenario, error) {
	sql := `SELECT * FROM scenarios`
	rows, err := u.conn.Query(ctx, sql)
	if err != nil {
		return []models.Scenario{}, fmt.Errorf("query: %w", err)
	}
	var scenarios []models.Scenario
	for rows.Next() {
		var scenario models.Scenario
		err := rows.Scan(&scenario.ID, &scenario.Name, &scenario.OwnerID, &scenario.StepsCount, &scenario.DeletedAt, &scenario.Data)
		if err != nil {
			return scenarios, fmt.Errorf("scan: %w", err)
		}
		scenarios = append(scenarios, scenario)
	}
	return scenarios, err
}

func (u *ScenarioRepo) GetByID(ctx context.Context, id uint64) (models.Scenario, error) {
	var scenario models.Scenario

	query := `SELECT * FROM scenarios WHERE id=$1`

	row := u.conn.QueryRow(ctx, query, id)
	err := row.Scan(&scenario.ID, &scenario.Name, &scenario.OwnerID, &scenario.StepsCount, &scenario.DeletedAt, &scenario.Data)
	if err != nil {
		return models.Scenario{}, err
	}

	return scenario, nil
}

func (u *ScenarioRepo) Update(ctx context.Context, scenario *models.Scenario) error {
	query := `UPDATE scenarios SET  name=$2, owner_id=$3, steps_count=$4, deleted_at=$5, data=$6 WHERE id=$1`

	_, err := u.conn.Exec(ctx, query, scenario.ID, scenario.Name, scenario.OwnerID, scenario.StepsCount, scenario.DeletedAt, scenario.Data)
	if err != nil {
		return err
	}

	return nil
}

func (u *ScenarioRepo) Delete(ctx context.Context, id uint64) error {
	sql := `DELETE FROM scenarios WHERE id=$1`
	_, err := u.conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}
