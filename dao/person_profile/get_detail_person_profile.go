package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function retrieves the details of a specific person profile.
func (d *person_profilePostgresqlSQLDAO) GetDetailPersonProfile(
     ctx *context.ContextModel,
     id int64,
) (
     repository.PersonProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT id, uuid_key, person_profile_id, title_id, title_name, first_name, last_name, 
               sex, country_id, country_name, province_id, province_name, district_id, 
               district_name, sub_district_id, sub_district_name, urban_village_id, 
               urban_village_name, island_id, island_name, postal_code_id, postal_code_name, 
               phone, email, photo, last_sync, created_client, updated_client, created_at, 
               created_by, updated_at, updated_by, deleted
        FROM %s
        WHERE id = $1
    `, dao.GetDBTable(ctx, "person_profile"))

    var model repository.PersonProfileModel
    err := d.db.QueryRow(query, id).Scan(
        &model.Id, &model.UuidKey, &model.PersonProfileId, &model.TitleId, 
        &model.TitleName, &model.FirstName, &model.LastName, &model.Sex, 
        &model.Nik, &model.Address1, &model.Address2, &model.Address3, 
        &model.Hamlet, &model.Neighbour, &model.CountryId, &model.CountryName, 
        &model.ProvinceId, &model.ProvinceName, &model.DistrictId, &model.DistrictName, 
        &model.SubDistrictId, &model.SubDistrictName, &model.UrbanVillageId, 
        &model.UrbanVillageName, &model.IslandId, &model.IslandName, 
        &model.PostalCodeId, &model.PostalCodeName, &model.Phone, 
        &model.Email, &model.Photo, &model.LastSync, &model.CreatedClient, 
        &model.UpdatedClient, &model.CreatedAt, &model.CreatedBy, 
        &model.UpdatedAt, &model.UpdatedBy, &model.Deleted,
    )
    return model, err
}
