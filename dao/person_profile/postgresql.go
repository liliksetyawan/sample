package person_profile

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
)

type person_profilePostgresqlSQLDAO struct {
    dao.GetListDataDAO
    db *sql.DB
}

func NewPersonProfilePostgresqlSQLDAO(
    db *sql.DB,
) PersonProfileDAO {
    return &person_profilePostgresqlSQLDAO{
        GetListDataDAO: dao.GetListDataDAO{
            DB: db,
        },
        db: db,
    }
}