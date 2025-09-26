package person_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function retrieves the details of a person profile by user id.
func (d *person_profilePostgresqlSQLDAO) GetPersonProfileByUserID(
     ctx *context.ContextModel,
     id int64,
) (
     repository.PersonProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT pp.id, pp.first_name, pp.last_name, 
               wp.id, wp.name, wp.username, 
               pp.created_at, pp.created_by, 
               pp.created_client, pp.updated_at, 
               pp.updated_by, pp.updated_client
        FROM %s pp
        JOIN %s wp ON pp.PersonProfileID = wp.id
        WHERE pp.PersonProfileID = $1
    `, dao.GetDBTable(ctx, "person_profile"), dao.GetDBTable(ctx, "wardes_profile"))

    var model repository.PersonProfileModel
    err := d.db.QueryRow(query, id).Scan(
        &model.ID, &model.FirstName, &model.LastName,
        &model.TitleID, &model.TitleName, &model.NIK,
        &model.CreatedAt, &model.CreatedBy, &model.CreatedClient,
        &model.UpdatedAt, &model.UpdatedBy, &model.UpdatedClient,
    )

    return model, err
}
