package parameter

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
)

type parameterPostgresqlSQLDAO struct {
    dao.GetListDataDAO
    db *sql.DB
}

func NewParameterPostgresqlSQLDAO(
    db *sql.DB,
) ParameterDAO {
    return &parameterPostgresqlSQLDAO{
        GetListDataDAO: dao.GetListDataDAO{
            DB: db,
        },
        db: db,
    }
}