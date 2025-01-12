
�
#proto/environment/environment.protobucketeer.environment"�
EnvironmentV2
id (	Rid
name (	Rname
url_code (	RurlCode 
description (	Rdescription

project_id (	R	projectId
archived (Rarchived

created_at (R	createdAt

updated_at (R	updatedAt'
organization_id	 (	RorganizationIdB5Z3github.com/bucketeer-io/bucketeer/proto/environmentbproto3
�
proto/account/account.protobucketeer.account#proto/environment/environment.proto"�
Account
id (	Rid
email (	Remail
name (	Rname3
role (2.bucketeer.account.Account.RoleRrole
disabled (Rdisabled

created_at (R	createdAt

updated_at (R	updatedAt
deleted (Rdeleted"9
Role

VIEWER 

EDITOR	
OWNER

UNASSIGNEDc"�
	AccountV2
email (	Remail
name (	Rname(
avatar_image_url (	RavatarImageUrl'
organization_id (	RorganizationId[
organization_role (2..bucketeer.account.AccountV2.Role.OrganizationRorganizationRoleY
environment_roles (2,.bucketeer.account.AccountV2.EnvironmentRoleRenvironmentRoles
disabled (Rdisabled

created_at (R	createdAt

updated_at	 (R	updatedAt�
Role"Y
Environment
Environment_UNASSIGNED 
Environment_VIEWER
Environment_EDITOR"t
Organization
Organization_UNASSIGNED 
Organization_MEMBER
Organization_ADMIN
Organization_OWNER{
EnvironmentRole%
environment_id (	RenvironmentIdA
role (2-.bucketeer.account.AccountV2.Role.EnvironmentRrole"�
EnvironmentRoleV2F
environment (2$.bucketeer.environment.EnvironmentV2Renvironment3
role (2.bucketeer.account.Account.RoleRrole#
trial_project (RtrialProject(
trial_started_at (RtrialStartedAtB1Z/github.com/bucketeer-io/bucketeer/proto/accountbproto3
�
proto/account/api_key.protobucketeer.account"�
APIKey
id (	Rid
name (	Rname2
role (2.bucketeer.account.APIKey.RoleRrole
disabled (Rdisabled

created_at (R	createdAt

updated_at (R	updatedAt"
Role
SDK 
SERVICE"�
EnvironmentAPIKey3
environment_namespace (	RenvironmentNamespace2
api_key (2.bucketeer.account.APIKeyRapiKey1
environment_disabled (RenvironmentDisabled

project_id (	R	projectIdB1Z/github.com/bucketeer-io/bucketeer/proto/accountbproto3
�
proto/account/command.protobucketeer.accountproto/account/account.protoproto/account/api_key.proto"1
CreateAdminAccountCommand
email (	Remail"
EnableAdminAccountCommand"
DisableAdminAccountCommand"
ConvertAccountCommand"
DeleteAccountCommand"a
CreateAccountCommand
email (	Remail3
role (2.bucketeer.account.Account.RoleRrole"O
ChangeAccountRoleCommand3
role (2.bucketeer.account.Account.RoleRrole"
EnableAccountCommand"
DisableAccountCommand"�
CreateAccountV2Command
email (	Remail
name (	Rname(
avatar_image_url (	RavatarImageUrl'
organization_id (	RorganizationId[
organization_role (2..bucketeer.account.AccountV2.Role.OrganizationRorganizationRoleY
environment_roles (2,.bucketeer.account.AccountV2.EnvironmentRoleRenvironmentRoles"0
ChangeAccountV2NameCommand
name (	Rname"P
$ChangeAccountV2AvatarImageUrlCommand(
avatar_image_url (	RavatarImageUrl"l
&ChangeAccountV2OrganizationRoleCommandB
role (2..bucketeer.account.AccountV2.Role.OrganizationRrole"l
&ChangeAccountV2EnvironmentRolesCommandB
roles (2,.bucketeer.account.AccountV2.EnvironmentRoleRroles"
EnableAccountV2Command"
DisableAccountV2Command"
DeleteAccountV2Command"]
CreateAPIKeyCommand
name (	Rname2
role (2.bucketeer.account.APIKey.RoleRrole"-
ChangeAPIKeyNameCommand
name (	Rname"
EnableAPIKeyCommand"
DisableAPIKeyCommandB1Z/github.com/bucketeer-io/bucketeer/proto/accountbproto3
�
google/protobuf/wrappers.protogoogle.protobuf"#
DoubleValue
value (Rvalue""

FloatValue
value (Rvalue""

Int64Value
value (Rvalue"#
UInt64Value
value (Rvalue""

Int32Value
value (Rvalue"#
UInt32Value
value (Rvalue"!
	BoolValue
value (Rvalue"#
StringValue
value (	Rvalue""

BytesValue
value (RvalueB�
com.google.protobufBWrappersProtoPZ1google.golang.org/protobuf/types/known/wrapperspb��GPB�Google.Protobuf.WellKnownTypesbproto3
�P
proto/account/service.protobucketeer.accountgoogle/protobuf/wrappers.protoproto/account/account.protoproto/account/api_key.protoproto/account/command.proto"
GetMeV2Request"-
GetMeByEmailV2Request
email (	Remail"�
GetMeV2Response
email (	Remail
is_admin (RisAdminQ
environment_roles (2$.bucketeer.account.EnvironmentRoleV2RenvironmentRoles"c
CreateAdminAccountRequestF
command (2,.bucketeer.account.CreateAdminAccountCommandRcommand"
CreateAdminAccountResponse"s
EnableAdminAccountRequest
id (	RidF
command (2,.bucketeer.account.EnableAdminAccountCommandRcommand"
EnableAdminAccountResponse"u
DisableAdminAccountRequest
id (	RidG
command (2-.bucketeer.account.DisableAdminAccountCommandRcommand"
DisableAdminAccountResponse".
GetAdminAccountRequest
email (	Remail"O
GetAdminAccountResponse4
account (2.bucketeer.account.AccountRaccount"�
ListAdminAccountsRequest
	page_size (RpageSize
cursor (	RcursorN
order_by (23.bucketeer.account.ListAdminAccountsRequest.OrderByRorderByc
order_direction (2:.bucketeer.account.ListAdminAccountsRequest.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled"A
OrderBy
DEFAULT 	
EMAIL

CREATED_AT

UPDATED_AT"#
OrderDirection
ASC 
DESC"�
ListAdminAccountsResponse6
accounts (2.bucketeer.account.AccountRaccounts
cursor (	Rcursor
total_count (R
totalCount"k
ConvertAccountRequest
id (	RidB
command (2(.bucketeer.account.ConvertAccountCommandRcommand"
ConvertAccountResponse"�
CreateAccountRequestA
command (2'.bucketeer.account.CreateAccountCommandRcommand3
environment_namespace (	RenvironmentNamespace"
CreateAccountResponse"�
EnableAccountRequest
id (	RidA
command (2'.bucketeer.account.EnableAccountCommandRcommand3
environment_namespace (	RenvironmentNamespace"
EnableAccountResponse"�
DisableAccountRequest
id (	RidB
command (2(.bucketeer.account.DisableAccountCommandRcommand3
environment_namespace (	RenvironmentNamespace"
DisableAccountResponse"�
ChangeAccountRoleRequest
id (	RidE
command (2+.bucketeer.account.ChangeAccountRoleCommandRcommand3
environment_namespace (	RenvironmentNamespace"
ChangeAccountRoleResponse"^
GetAccountRequest
email (	Remail3
environment_namespace (	RenvironmentNamespace"J
GetAccountResponse4
account (2.bucketeer.account.AccountRaccount"�
ListAccountsRequest
	page_size (RpageSize
cursor (	Rcursor3
environment_namespace (	RenvironmentNamespaceI
order_by (2..bucketeer.account.ListAccountsRequest.OrderByRorderBy^
order_direction (25.bucketeer.account.ListAccountsRequest.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled/
role (2.google.protobuf.Int32ValueRrole"A
OrderBy
DEFAULT 	
EMAIL

CREATED_AT

UPDATED_AT"#
OrderDirection
ASC 
DESC"�
ListAccountsResponse6
accounts (2.bucketeer.account.AccountRaccounts
cursor (	Rcursor
total_count (R
totalCount"�
CreateAccountV2Request'
organization_id (	RorganizationIdC
command (2).bucketeer.account.CreateAccountV2CommandRcommand"Q
CreateAccountV2Response6
account (2.bucketeer.account.AccountV2Raccount"�
EnableAccountV2Request
email (	Remail'
organization_id (	RorganizationIdC
command (2).bucketeer.account.EnableAccountV2CommandRcommand"
EnableAccountV2Response"�
DisableAccountV2Request
email (	Remail'
organization_id (	RorganizationIdD
command (2*.bucketeer.account.DisableAccountV2CommandRcommand"
DisableAccountV2Response"�
DeleteAccountV2Request
email (	Remail'
organization_id (	RorganizationIdC
command (2).bucketeer.account.DeleteAccountV2CommandRcommand"
DeleteAccountV2Response"�
UpdateAccountV2Request
email (	Remail'
organization_id (	RorganizationId]
change_name_command (2-.bucketeer.account.ChangeAccountV2NameCommandRchangeNameCommandr
change_avatar_url_command (27.bucketeer.account.ChangeAccountV2AvatarImageUrlCommandRchangeAvatarUrlCommand�
 change_organization_role_command (29.bucketeer.account.ChangeAccountV2OrganizationRoleCommandRchangeOrganizationRoleCommand�
 change_environment_roles_command (29.bucketeer.account.ChangeAccountV2EnvironmentRolesCommandRchangeEnvironmentRolesCommand"
UpdateAccountV2Response"T
GetAccountV2Request
email (	Remail'
organization_id (	RorganizationId"N
GetAccountV2Response6
account (2.bucketeer.account.AccountV2Raccount"�
ListAccountsV2Request
	page_size (RpageSize
cursor (	Rcursor'
organization_id (	RorganizationIdK
order_by (20.bucketeer.account.ListAccountsV2Request.OrderByRorderBy`
order_direction (27.bucketeer.account.ListAccountsV2Request.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled/
role (2.google.protobuf.Int32ValueRrole"A
OrderBy
DEFAULT 	
EMAIL

CREATED_AT

UPDATED_AT"#
OrderDirection
ASC 
DESC"�
ListAccountsV2Response8
accounts (2.bucketeer.account.AccountV2Raccounts
cursor (	Rcursor
total_count (R
totalCount"�
CreateAPIKeyRequest@
command (2&.bucketeer.account.CreateAPIKeyCommandRcommand3
environment_namespace (	RenvironmentNamespace"J
CreateAPIKeyResponse2
api_key (2.bucketeer.account.APIKeyRapiKey"�
ChangeAPIKeyNameRequest
id (	RidD
command (2*.bucketeer.account.ChangeAPIKeyNameCommandRcommand3
environment_namespace (	RenvironmentNamespace"
ChangeAPIKeyNameResponse"�
EnableAPIKeyRequest
id (	Rid@
command (2&.bucketeer.account.EnableAPIKeyCommandRcommand3
environment_namespace (	RenvironmentNamespace"
EnableAPIKeyResponse"�
DisableAPIKeyRequest
id (	RidA
command (2'.bucketeer.account.DisableAPIKeyCommandRcommand3
environment_namespace (	RenvironmentNamespace"
DisableAPIKeyResponse"W
GetAPIKeyRequest
id (	Rid3
environment_namespace (	RenvironmentNamespace"G
GetAPIKeyResponse2
api_key (2.bucketeer.account.APIKeyRapiKey"�
ListAPIKeysRequest
	page_size (RpageSize
cursor (	Rcursor3
environment_namespace (	RenvironmentNamespaceH
order_by (2-.bucketeer.account.ListAPIKeysRequest.OrderByRorderBy]
order_direction (24.bucketeer.account.ListAPIKeysRequest.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled"@
OrderBy
DEFAULT 
NAME

CREATED_AT

UPDATED_AT"#
OrderDirection
ASC 
DESC"�
ListAPIKeysResponse4
api_keys (2.bucketeer.account.APIKeyRapiKeys
cursor (	Rcursor
total_count (R
totalCount"<
*GetAPIKeyBySearchingAllEnvironmentsRequest
id (	Rid"�
+GetAPIKeyBySearchingAllEnvironmentsResponseT
environment_api_key (2$.bucketeer.account.EnvironmentAPIKeyRenvironmentApiKey2�
AccountServiceP
GetMeV2!.bucketeer.account.GetMeV2Request".bucketeer.account.GetMeV2Response^
GetMeByEmailV2(.bucketeer.account.GetMeByEmailV2Request".bucketeer.account.GetMeV2Responseq
CreateAdminAccount,.bucketeer.account.CreateAdminAccountRequest-.bucketeer.account.CreateAdminAccountResponseq
EnableAdminAccount,.bucketeer.account.EnableAdminAccountRequest-.bucketeer.account.EnableAdminAccountResponset
DisableAdminAccount-.bucketeer.account.DisableAdminAccountRequest..bucketeer.account.DisableAdminAccountResponseh
GetAdminAccount).bucketeer.account.GetAdminAccountRequest*.bucketeer.account.GetAdminAccountResponsen
ListAdminAccounts+.bucketeer.account.ListAdminAccountsRequest,.bucketeer.account.ListAdminAccountsResponsee
ConvertAccount(.bucketeer.account.ConvertAccountRequest).bucketeer.account.ConvertAccountResponseb
CreateAccount'.bucketeer.account.CreateAccountRequest(.bucketeer.account.CreateAccountResponseb
EnableAccount'.bucketeer.account.EnableAccountRequest(.bucketeer.account.EnableAccountResponsee
DisableAccount(.bucketeer.account.DisableAccountRequest).bucketeer.account.DisableAccountResponsen
ChangeAccountRole+.bucketeer.account.ChangeAccountRoleRequest,.bucketeer.account.ChangeAccountRoleResponseY

GetAccount$.bucketeer.account.GetAccountRequest%.bucketeer.account.GetAccountResponse_
ListAccounts&.bucketeer.account.ListAccountsRequest'.bucketeer.account.ListAccountsResponseh
CreateAccountV2).bucketeer.account.CreateAccountV2Request*.bucketeer.account.CreateAccountV2Responseh
EnableAccountV2).bucketeer.account.EnableAccountV2Request*.bucketeer.account.EnableAccountV2Responsek
DisableAccountV2*.bucketeer.account.DisableAccountV2Request+.bucketeer.account.DisableAccountV2Responseh
UpdateAccountV2).bucketeer.account.UpdateAccountV2Request*.bucketeer.account.UpdateAccountV2Responseh
DeleteAccountV2).bucketeer.account.DeleteAccountV2Request*.bucketeer.account.DeleteAccountV2Response_
GetAccountV2&.bucketeer.account.GetAccountV2Request'.bucketeer.account.GetAccountV2Responsee
ListAccountsV2(.bucketeer.account.ListAccountsV2Request).bucketeer.account.ListAccountsV2Response_
CreateAPIKey&.bucketeer.account.CreateAPIKeyRequest'.bucketeer.account.CreateAPIKeyResponsek
ChangeAPIKeyName*.bucketeer.account.ChangeAPIKeyNameRequest+.bucketeer.account.ChangeAPIKeyNameResponse_
EnableAPIKey&.bucketeer.account.EnableAPIKeyRequest'.bucketeer.account.EnableAPIKeyResponseb
DisableAPIKey'.bucketeer.account.DisableAPIKeyRequest(.bucketeer.account.DisableAPIKeyResponseV
	GetAPIKey#.bucketeer.account.GetAPIKeyRequest$.bucketeer.account.GetAPIKeyResponse\
ListAPIKeys%.bucketeer.account.ListAPIKeysRequest&.bucketeer.account.ListAPIKeysResponse�
#GetAPIKeyBySearchingAllEnvironments=.bucketeer.account.GetAPIKeyBySearchingAllEnvironmentsRequest>.bucketeer.account.GetAPIKeyBySearchingAllEnvironmentsResponseB1Z/github.com/bucketeer-io/bucketeer/proto/accountbproto3