
�

proto/environment/command.protobucketeer.environment"�
CreateEnvironmentV2Command
name (	Rname
url_code (	RurlCode 
description (	Rdescription

project_id (	R	projectId"0
RenameEnvironmentV2Command
name (	Rname"I
%ChangeDescriptionEnvironmentV2Command 
description (	Rdescription"
ArchiveEnvironmentV2Command"
UnarchiveEnvironmentV2Command"{
CreateProjectCommand
id (	BRid 
description (	Rdescription
name (	Rname
url_code (	RurlCode"t
CreateTrialProjectCommand
id (	BRid
email (	Remail
name (	Rname
url_code (	RurlCode"*
RenameProjectCommand
name (	Rname"C
ChangeDescriptionProjectCommand 
description (	Rdescription"
EnableProjectCommand"
DisableProjectCommand"
ConvertTrialProjectCommand"�
CreateOrganizationCommand
name (	Rname
url_code (	RurlCode 
description (	Rdescription
is_trial (RisTrial"3
ChangeNameOrganizationCommand
name (	Rname"H
$ChangeDescriptionOrganizationCommand 
description (	Rdescription"
EnableOrganizationCommand"
DisableOrganizationCommand"
ArchiveOrganizationCommand"
UnarchiveOrganizationCommand"!
ConvertTrialOrganizationCommandB5Z3github.com/bucketeer-io/bucketeer/proto/environmentbproto3
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
�
$proto/environment/organization.protobucketeer.environment"�
Organization
id (	Rid
name (	Rname
url_code (	RurlCode 
description (	Rdescription
disabled (Rdisabled
archived (Rarchived
trial (Rtrial

created_at (R	createdAt

updated_at	 (R	updatedAtB5Z3github.com/bucketeer-io/bucketeer/proto/environmentbproto3
�
proto/environment/project.protobucketeer.environment"�
Project
id (	Rid 
description (	Rdescription
disabled (Rdisabled
trial (Rtrial#
creator_email (	RcreatorEmail

created_at (R	createdAt

updated_at (R	updatedAt
name (	Rname
url_code	 (	RurlCode'
organization_id
 (	RorganizationIdB5Z3github.com/bucketeer-io/bucketeer/proto/environmentbproto3
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
�C
proto/environment/service.protobucketeer.environmentgoogle/protobuf/wrappers.proto#proto/environment/environment.protoproto/environment/project.proto$proto/environment/organization.protoproto/environment/command.proto")
GetEnvironmentV2Request
id (	Rid"b
GetEnvironmentV2ResponseF
environment (2$.bucketeer.environment.EnvironmentV2Renvironment"�
ListEnvironmentsV2Request
	page_size (RpageSize
cursor (	RcursorS
order_by (28.bucketeer.environment.ListEnvironmentsV2Request.OrderByRorderByh
order_direction (2?.bucketeer.environment.ListEnvironmentsV2Request.OrderDirectionRorderDirection

project_id (	R	projectId6
archived (2.google.protobuf.BoolValueRarchived%
search_keyword (	RsearchKeyword"V
OrderBy
DEFAULT 
ID
NAME
URL_CODE

CREATED_AT

UPDATED_AT"#
OrderDirection
ASC 
DESC"�
ListEnvironmentsV2ResponseH
environments (2$.bucketeer.environment.EnvironmentV2Renvironments
cursor (	Rcursor
total_count (R
totalCount"i
CreateEnvironmentV2RequestK
command (21.bucketeer.environment.CreateEnvironmentV2CommandRcommand"e
CreateEnvironmentV2ResponseF
environment (2$.bucketeer.environment.EnvironmentV2Renvironment"�
UpdateEnvironmentV2Request
id (	RidX
rename_command (21.bucketeer.environment.RenameEnvironmentV2CommandRrenameCommandz
change_description_command (2<.bucketeer.environment.ChangeDescriptionEnvironmentV2CommandRchangeDescriptionCommand"
UpdateEnvironmentV2Response"{
ArchiveEnvironmentV2Request
id (	RidL
command (22.bucketeer.environment.ArchiveEnvironmentV2CommandRcommand"
ArchiveEnvironmentV2Response"
UnarchiveEnvironmentV2Request
id (	RidN
command (24.bucketeer.environment.UnarchiveEnvironmentV2CommandRcommand" 
UnarchiveEnvironmentV2Response"#
GetProjectRequest
id (	Rid"N
GetProjectResponse8
project (2.bucketeer.environment.ProjectRproject"�
ListProjectsRequest
	page_size (RpageSize
cursor (	RcursorM
order_by (22.bucketeer.environment.ListProjectsRequest.OrderByRorderByb
order_direction (29.bucketeer.environment.ListProjectsRequest.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled"V
OrderBy
DEFAULT 
ID

CREATED_AT

UPDATED_AT
NAME
URL_CODE"#
OrderDirection
ASC 
DESC"�
ListProjectsResponse:
projects (2.bucketeer.environment.ProjectRprojects
cursor (	Rcursor
total_count (R
totalCount"]
CreateProjectRequestE
command (2+.bucketeer.environment.CreateProjectCommandRcommand"Q
CreateProjectResponse8
project (2.bucketeer.environment.ProjectRproject"g
CreateTrialProjectRequestJ
command (20.bucketeer.environment.CreateTrialProjectCommandRcommand"
CreateTrialProjectResponse"�
UpdateProjectRequest
id (	Ridt
change_description_command (26.bucketeer.environment.ChangeDescriptionProjectCommandRchangeDescriptionCommandR
rename_command (2+.bucketeer.environment.RenameProjectCommandRrenameCommand"
UpdateProjectResponse"m
EnableProjectRequest
id (	RidE
command (2+.bucketeer.environment.EnableProjectCommandRcommand"
EnableProjectResponse"o
DisableProjectRequest
id (	RidF
command (2,.bucketeer.environment.DisableProjectCommandRcommand"
DisableProjectResponse"y
ConvertTrialProjectRequest
id (	RidK
command (21.bucketeer.environment.ConvertTrialProjectCommandRcommand"
ConvertTrialProjectResponse"(
GetOrganizationRequest
id (	Rid"b
GetOrganizationResponseG
organization (2#.bucketeer.environment.OrganizationRorganization"�
ListOrganizationsRequest
	page_size (RpageSize
cursor (	RcursorR
order_by (27.bucketeer.environment.ListOrganizationsRequest.OrderByRorderByg
order_direction (2>.bucketeer.environment.ListOrganizationsRequest.OrderDirectionRorderDirection%
search_keyword (	RsearchKeyword6
disabled (2.google.protobuf.BoolValueRdisabled6
archived (2.google.protobuf.BoolValueRarchived"V
OrderBy
DEFAULT 
ID

CREATED_AT

UPDATED_AT
NAME
URL_CODE"#
OrderDirection
ASC 
DESC"�
ListOrganizationsResponseI
Organizations (2#.bucketeer.environment.OrganizationROrganizations
cursor (	Rcursor
total_count (R
totalCount"g
CreateOrganizationRequestJ
command (20.bucketeer.environment.CreateOrganizationCommandRcommand"e
CreateOrganizationResponseG
Organization (2#.bucketeer.environment.OrganizationROrganization"�
UpdateOrganizationRequest
id (	Rid[
rename_command (24.bucketeer.environment.ChangeNameOrganizationCommandRrenameCommandy
change_description_command (2;.bucketeer.environment.ChangeDescriptionOrganizationCommandRchangeDescriptionCommand"
UpdateOrganizationResponse"w
EnableOrganizationRequest
id (	RidJ
command (20.bucketeer.environment.EnableOrganizationCommandRcommand"
EnableOrganizationResponse"y
DisableOrganizationRequest
id (	RidK
command (21.bucketeer.environment.DisableOrganizationCommandRcommand"
DisableOrganizationResponse"y
ArchiveOrganizationRequest
id (	RidK
command (21.bucketeer.environment.ArchiveOrganizationCommandRcommand"
ArchiveOrganizationResponse"}
UnarchiveOrganizationRequest
id (	RidM
command (23.bucketeer.environment.UnarchiveOrganizationCommandRcommand"
UnarchiveOrganizationResponse"�
ConvertTrialOrganizationRequest
id (	RidP
command (26.bucketeer.environment.ConvertTrialOrganizationCommandRcommand""
 ConvertTrialOrganizationResponse2�
EnvironmentServiceu
GetEnvironmentV2..bucketeer.environment.GetEnvironmentV2Request/.bucketeer.environment.GetEnvironmentV2Response" {
ListEnvironmentsV20.bucketeer.environment.ListEnvironmentsV2Request1.bucketeer.environment.ListEnvironmentsV2Response" ~
CreateEnvironmentV21.bucketeer.environment.CreateEnvironmentV2Request2.bucketeer.environment.CreateEnvironmentV2Response" ~
UpdateEnvironmentV21.bucketeer.environment.UpdateEnvironmentV2Request2.bucketeer.environment.UpdateEnvironmentV2Response" �
ArchiveEnvironmentV22.bucketeer.environment.ArchiveEnvironmentV2Request3.bucketeer.environment.ArchiveEnvironmentV2Response" �
UnarchiveEnvironmentV24.bucketeer.environment.UnarchiveEnvironmentV2Request5.bucketeer.environment.UnarchiveEnvironmentV2Response" c

GetProject(.bucketeer.environment.GetProjectRequest).bucketeer.environment.GetProjectResponse" i
ListProjects*.bucketeer.environment.ListProjectsRequest+.bucketeer.environment.ListProjectsResponse" l
CreateProject+.bucketeer.environment.CreateProjectRequest,.bucketeer.environment.CreateProjectResponse" {
CreateTrialProject0.bucketeer.environment.CreateTrialProjectRequest1.bucketeer.environment.CreateTrialProjectResponse" l
UpdateProject+.bucketeer.environment.UpdateProjectRequest,.bucketeer.environment.UpdateProjectResponse" l
EnableProject+.bucketeer.environment.EnableProjectRequest,.bucketeer.environment.EnableProjectResponse" o
DisableProject,.bucketeer.environment.DisableProjectRequest-.bucketeer.environment.DisableProjectResponse" ~
ConvertTrialProject1.bucketeer.environment.ConvertTrialProjectRequest2.bucketeer.environment.ConvertTrialProjectResponse" r
GetOrganization-.bucketeer.environment.GetOrganizationRequest..bucketeer.environment.GetOrganizationResponse" x
ListOrganizations/.bucketeer.environment.ListOrganizationsRequest0.bucketeer.environment.ListOrganizationsResponse" {
CreateOrganization0.bucketeer.environment.CreateOrganizationRequest1.bucketeer.environment.CreateOrganizationResponse" {
UpdateOrganization0.bucketeer.environment.UpdateOrganizationRequest1.bucketeer.environment.UpdateOrganizationResponse" {
EnableOrganization0.bucketeer.environment.EnableOrganizationRequest1.bucketeer.environment.EnableOrganizationResponse" ~
DisableOrganization1.bucketeer.environment.DisableOrganizationRequest2.bucketeer.environment.DisableOrganizationResponse" ~
ArchiveOrganization1.bucketeer.environment.ArchiveOrganizationRequest2.bucketeer.environment.ArchiveOrganizationResponse" �
UnarchiveOrganization3.bucketeer.environment.UnarchiveOrganizationRequest4.bucketeer.environment.UnarchiveOrganizationResponse" �
ConvertTrialOrganization6.bucketeer.environment.ConvertTrialOrganizationRequest7.bucketeer.environment.ConvertTrialOrganizationResponse" B5Z3github.com/bucketeer-io/bucketeer/proto/environmentbproto3