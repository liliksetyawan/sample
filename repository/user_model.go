package repository

import "database/sql"

type UserModel struct {
    ID             sql.NullInt64
    UUIDKey        sql.NullString
    AuthUserID     sql.NullInt64
    PersonProfileID sql.NullInt64
    ClientID       sql.NullString
    UserCode       sql.NullString
    Username       sql.NullString
    LastToken      sql.NullTime
    SignatureKey   sql.NullString
    IPWhitelist    sql.NullString
    Device         sql.NullString
    Level          sql.NullInt64
    Locale         sql.NullString
    AdditionalInfo sql.NullString
    Alias          sql.NullString
    IsUserNexAcc   sql.NullBool
    MobileAccess   sql.NullBool
    WebsiteAccess  sql.NullBool
    WdiPortalAccess sql.NullBool
    DomainNd6      sql.NullString
    UsernameNd6    sql.NullString
    PasswordNd6    sql.NullString
    DomainNexvin   sql.NullString
    UsernameNexvin sql.NullString
    PasswordNexvin sql.NullString
    CreatedBy      sql.NullInt64
    CreatedClient  sql.NullString
    CreatedAt      sql.NullTime
    UpdatedBy      sql.NullInt64
    UpdatedAt      sql.NullTime
    UpdatedClient  sql.NullString
    Deleted        sql.NullBool
}