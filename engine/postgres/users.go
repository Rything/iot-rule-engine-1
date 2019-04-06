package postgres

import (
	"database/sql"

	"github.com/nattaponra/iot-rule-engine/engine"

	"github.com/lib/pq"
)

var _ engine.RuleEngineRepository = (*ruleEngineRepository)(nil)

const errDuplicate = "unique_violation"

type ruleEngineRepository struct {
	db *sql.DB
}

func New(db *sql.DB) engine.RuleEngineRepository {
	return &ruleEngineRepository{db}
}

func (ur ruleEngineRepository) Save(rule engine.Rule) error {
	q := `INSERT INTO users (email, password) VALUES ($1, $2)`

	if _, err := ur.db.Exec(q, rule.Email, rule.Password); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && errDuplicate == pqErr.Code.Name() {
			return engine.ErrConflict
		}
		return err
	}

	return nil
}
