package database

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/bdjuno/v3/types"
)

// SaveDeploymentParams allows to store the given params inside the database
func (db *Db) SaveDeploymentParams(params *types.DeploymentParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling deployment params: %s", err)
	}

	stmt := `
INSERT INTO deployment_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE deployment_params.height <= excluded.height`

	_, err = db.Sql.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing deployment params: %s", err)
	}

	return nil
}
