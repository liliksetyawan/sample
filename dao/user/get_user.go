package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get a user by ID
func (d *userPostgresqlSQLDAO) GetUser(
     ctx *context.ContextModel,
     uuid_key string,
) (
     repository.UserModel, 
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
        WHERE uuid_key = $1 AND deleted = false
    `, dao.GetDBTable(ctx, "user"))

    args := []interface{}{uuid_key}
    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2"
        args = append(args, ctx.Limitation.UserID)
    }

    var model repository.UserModel
    err := d.db.QueryRow(query, args...).Scan(
        &model.UUIDKey, &model.ID, &model.AuthUserID, &model.PersonProfileID, &model.ClientID,
        &model.UserCode, &model.Username, &model.LastToken, &model.SignatureKey, &model.IPWhitelist,
        &model.Device, &model.Level, &model.Locale, &model.AdditionalInfo, &model.Alias,
        &model.IsUserNexAcc, &model.MobileAccess, &model.WebsiteAccess, &model.WDIPortalAccess,
        &model.DomainNd6, &model.UsernameNd6, &model.PasswordNd6, &model.DomainNexvin,
        &model.UsernameNexvin, &model.PasswordNexvin, &model.CreatedBy, &model.CreatedClient,
        &model.CreatedAt, &model.UpdatedBy, &model.UpdatedAt, &model.UpdatedClient, &model.Deleted,
    )

    return model, err
}
