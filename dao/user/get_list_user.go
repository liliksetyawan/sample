package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get a list of users with pagination and search
func (d *userPostgresqlSQLDAO) GetListUser(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, id, auth_user_id, person_profile_id, client_id, user_code, username,
               last_token, signature_key, ip_whitelist, device, level, locale, additional_info,
               alias, is_user_nex_acc, mobile_access, website_access, wdi_portal_access,
               domain_nd6, username_nd6, password_nd6, domain_nexvin, username_nexvin,
               password_nexvin, created_by, created_client, created_at, updated_by,
               updated_at, updated_client, deleted
        FROM %s
    `, dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.UserModel
            err := rows.Scan(
                &temp.UUIDKey, &temp.ID, &temp.AuthUserID, &temp.PersonProfileID, &temp.ClientID,
                &temp.UserCode, &temp.Username, &temp.LastToken, &temp.SignatureKey, &temp.IPWhitelist,
                &temp.Device, &temp.Level, &temp.Locale, &temp.AdditionalInfo, &temp.Alias,
                &temp.IsUserNexAcc, &temp.MobileAccess, &temp.WebsiteAccess, &temp.WDIPortalAccess,
                &temp.DomainNd6, &temp.UsernameNd6, &temp.PasswordNd6, &temp.DomainNexvin,
                &temp.UsernameNexvin, &temp.PasswordNexvin, &temp.CreatedBy, &temp.CreatedClient,
                &temp.CreatedAt, &temp.UpdatedBy, &temp.UpdatedAt, &temp.UpdatedClient, &temp.Deleted,
            )
            return temp, err
        },
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
