package user

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
)

type userPostgresqlSQLDAO struct {
    dao.GetListDataDAO
    db *sql.DB
}

func NewUserPostgresqlSQLDAO(
    db *sql.DB,
) UserDAO {
    return &userPostgresqlSQLDAO{
        GetListDataDAO: dao.GetListDataDAO{
            DB: db,
        },
        db: db,
    }
}