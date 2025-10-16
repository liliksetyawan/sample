package wardes_profile_image

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
)

type wardes_profile_imagePostgresqlSQLDAO struct {
    dao.GetListDataDAO
    db *sql.DB
}

func NewWardesProfileImagePostgresqlSQLDAO(
    db *sql.DB,
) WardesProfileImageDAO {
    return &wardes_profile_imagePostgresqlSQLDAO{
        GetListDataDAO: dao.GetListDataDAO{
            DB: db,
        },
        db: db,
    }
}