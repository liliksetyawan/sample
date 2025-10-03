package wardes_profile

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
)

type wardes_profilePostgresqlSQLDAO struct {
    dao.GetListDataDAO
    db *sql.DB
}

func NewWardesProfilePostgresqlSQLDAO(
    db *sql.DB,
) WardesProfileDAO {
    return &wardes_profilePostgresqlSQLDAO{
        GetListDataDAO: dao.GetListDataDAO{
            DB: db,
        },
        db: db,
    }
}




go