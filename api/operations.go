// Code generated by gqlclientgen - DO NOT EDIT

package api

import (
	"context"
	gqlclient "git.sr.ht/~emersion/gqlclient"
)

type Auth_api_key struct {
	Created           timestamptz  `json:"created"`
	Default_tenant_id *uuid        `json:"default_tenant_id,omitempty"`
	Expires_at        *timestamptz `json:"expires_at,omitempty"`
	Id                uuid         `json:"id"`
	Key_hash          *string      `json:"key_hash,omitempty"`
	Name              string       `json:"name"`
	Updated           timestamptz  `json:"updated"`
	User_id           uuid         `json:"user_id"`
}

type DateTime string

type String_comparison_exp struct {
	_similar *string `json:"_similar,omitempty"`
}

type UUID string

type account_type string

type auth_api_key_bool_exp struct {
	Id      *uuid_comparison_exp `json:"id,omitempty"`
	User_id *uuid_comparison_exp `json:"user_id,omitempty"`
}

type auth_role struct {
	Created     timestamptz `json:"created"`
	Id          uuid        `json:"id"`
	Name        string      `json:"name"`
	Permissions jsonb       `json:"permissions"`
	Tenant_id   *uuid       `json:"tenant_id,omitempty"`
	Updated     timestamptz `json:"updated"`
}

type create_api_key_input struct {
	User_id UUID   `json:"user_id"`
	Name    string `json:"name"`
	// Null = never expire
	Expires_at *DateTime `json:"expires_at,omitempty"`
	// Defaults to current tenant
	Default_tenant_id *UUID `json:"default_tenant_id,omitempty"`
}

type create_api_key_payload struct {
	// Key id
	Id *UUID `json:"id,omitempty"`
	// Secret value
	Key *string `json:"key,omitempty"`
}

type create_project_input struct {
	// The tenant
	Tenant_id *UUID `json:"tenant_id,omitempty"`
	// The name of the project to create
	Name string `json:"name"`
	// An optional description of this project - can be updated later with setProjectDescription
	Description *string `json:"description,omitempty"`
}

type create_service_account_input struct {
	// Tenant id. Must match current tenant otherwise an Unauthorized error will be returned.
	Tenant_id UUID `json:"tenant_id"`
	// Service account name. Populates the first_name field.
	Name string `json:"name"`
	// Role id. Defaults to TENANT_ADMIN if unspecified.
	Role_id *UUID `json:"role_id,omitempty"`
}

type delete_api_key_input struct {
	Key_id UUID `json:"key_id"`
}

type delete_project_input struct {
	// The ID of the project to delete
	Project_id UUID `json:"project_id"`
}

type delete_service_account_input struct {
	User_id UUID `json:"user_id"`
}

type jsonb string

type membership struct {
	Created     timestamptz `json:"created"`
	Id          uuid        `json:"id"`
	Role        *string     `json:"role,omitempty"`
	Role_detail *auth_role  `json:"role_detail,omitempty"`
	Role_id     *uuid       `json:"role_id,omitempty"`
	Tenant      *tenant     `json:"tenant"`
	Tenant_id   uuid        `json:"tenant_id"`
	Updated     timestamptz `json:"updated"`
	User_id     uuid        `json:"user_id"`
}

type membership_bool_exp struct {
	Id        *uuid_comparison_exp `json:"id,omitempty"`
	Role_id   *uuid_comparison_exp `json:"role_id,omitempty"`
	Tenant_id *uuid_comparison_exp `json:"tenant_id,omitempty"`
	User_id   *uuid_comparison_exp `json:"user_id,omitempty"`
}

type membership_id_payload struct {
	Id *UUID `json:"id,omitempty"`
}

type membership_role = string

const (
	Membership_roleReadOnlyUser membership_role = "READ_ONLY_USER"
	Membership_roleUser         membership_role = "USER"
	Membership_roleTenantAdmin  membership_role = "TENANT_ADMIN"
)

type project struct {
	Created     timestamptz  `json:"created"`
	Deleted_at  *timestamptz `json:"deleted_at,omitempty"`
	Description *string      `json:"description,omitempty"`
	Id          uuid         `json:"id"`
	Name        string       `json:"name"`
	Tenant_id   uuid         `json:"tenant_id"`
	Updated     timestamptz  `json:"updated"`
}

type project_id_payload struct {
	Id *UUID `json:"id,omitempty"`
}

type set_membership_role_input struct {
	Membership_id UUID             `json:"membership_id"`
	Role          *membership_role `json:"role,omitempty"`
	Role_id       *UUID            `json:"role_id,omitempty"`
}

type set_project_description_input struct {
	// The ID of the project to update
	Project_id UUID `json:"project_id"`
	// The new description of this project.
	Description *string `json:"description,omitempty"`
}

type set_project_name_input struct {
	// The ID of the project to update
	Project_id UUID `json:"project_id"`
	// The new name for this project
	Name string `json:"name"`
}

type success_payload struct {
	// A boolean specifying whether this mutation was successful
	Success *bool `json:"success,omitempty"`
	// If populated, contains an error message explaining what went wrong with the mutation
	Error *string `json:"error,omitempty"`
}

type tenant struct {
	Id   uuid   `json:"id"`
	Name string `json:"name"`
}

type timestamptz = string

type user struct {
	Id                 uuid        `json:"id"`
	Username           string      `json:"username"`
	First_name         *string     `json:"first_name,omitempty"`
	Default_membership *membership `json:"default_membership,omitempty"`
}

type user_id_payload struct {
	Id *UUID `json:"id,omitempty"`
}

type user_view_same_tenant struct {
	Account_type *account_type `json:"account_type,omitempty"`
	Email        *string       `json:"email,omitempty"`
	First_name   *string       `json:"first_name,omitempty"`
	Id           *uuid         `json:"id,omitempty"`
	Last_name    *string       `json:"last_name,omitempty"`
	Memberships  []membership  `json:"memberships"`
	Username     *string       `json:"username,omitempty"`
}

type user_view_same_tenant_bool_exp struct {
	Id       *uuid_comparison_exp   `json:"id,omitempty"`
	Username *String_comparison_exp `json:"username,omitempty"`
}

type uuid = string

type uuid_comparison_exp struct {
	_eq      *uuid  `json:"_eq,omitempty"`
	_gt      *uuid  `json:"_gt,omitempty"`
	_gte     *uuid  `json:"_gte,omitempty"`
	_in      []uuid `json:"_in,omitempty"`
	_is_null *bool  `json:"_is_null,omitempty"`
	_lt      *uuid  `json:"_lt,omitempty"`
	_lte     *uuid  `json:"_lte,omitempty"`
	_neq     *uuid  `json:"_neq,omitempty"`
	_nin     []uuid `json:"_nin,omitempty"`
}

func CreateProject(client *gqlclient.Client, ctx context.Context, name string, description *string) (create_project *project_id_payload, err error) {
	op := gqlclient.NewOperation("mutation createProject ($name: String!, $description: String) {\n\tcreate_project(input: {name:$name,description:$description}) {\n\t\tid\n\t}\n}\n")
	op.Var("name", name)
	op.Var("description", description)
	var respData struct {
		Create_project *project_id_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Create_project, err
}

func SetProjectName(client *gqlclient.Client, ctx context.Context, project_id UUID, name string) (set_project_name *project_id_payload, err error) {
	op := gqlclient.NewOperation("mutation setProjectName ($project_id: UUID!, $name: String!) {\n\tset_project_name(input: {project_id:$project_id,name:$name}) {\n\t\tid\n\t}\n}\n")
	op.Var("project_id", project_id)
	op.Var("name", name)
	var respData struct {
		Set_project_name *project_id_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Set_project_name, err
}

func SetProjectDescription(client *gqlclient.Client, ctx context.Context, project_id UUID, description *string) (set_project_description *project_id_payload, err error) {
	op := gqlclient.NewOperation("mutation setProjectDescription ($project_id: UUID!, $description: String) {\n\tset_project_description(input: {project_id:$project_id,description:$description}) {\n\t\tid\n\t}\n}\n")
	op.Var("project_id", project_id)
	op.Var("description", description)
	var respData struct {
		Set_project_description *project_id_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Set_project_description, err
}

func DeleteProject(client *gqlclient.Client, ctx context.Context, projectId UUID) (delete_project *success_payload, err error) {
	op := gqlclient.NewOperation("mutation deleteProject ($projectId: UUID!) {\n\tdelete_project(input: {project_id:$projectId}) {\n\t\tsuccess\n\t\terror\n\t}\n}\n")
	op.Var("projectId", projectId)
	var respData struct {
		Delete_project *success_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Delete_project, err
}

func GetProject(client *gqlclient.Client, ctx context.Context, id uuid) (project_by_pk *project, err error) {
	op := gqlclient.NewOperation("query getProject ($id: uuid!) {\n\tproject_by_pk(id: $id) {\n\t\tid\n\t\tname\n\t\tdescription\n\t\ttenant_id\n\t\tupdated\n\t}\n}\n")
	op.Var("id", id)
	var respData struct {
		Project_by_pk *project
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Project_by_pk, err
}

func CreateServiceAccount(client *gqlclient.Client, ctx context.Context, tenant_id UUID, name string, role_id *UUID) (create_service_account *user_id_payload, err error) {
	op := gqlclient.NewOperation("mutation CreateServiceAccount ($tenant_id: UUID!, $name: String!, $role_id: UUID) {\n\tcreate_service_account(input: {tenant_id:$tenant_id,name:$name,role_id:$role_id}) {\n\t\tid\n\t}\n}\n")
	op.Var("tenant_id", tenant_id)
	op.Var("name", name)
	op.Var("role_id", role_id)
	var respData struct {
		Create_service_account *user_id_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Create_service_account, err
}

func SetMembershipRole(client *gqlclient.Client, ctx context.Context, membership_id UUID, role membership_role) (set_membership_role *membership_id_payload, err error) {
	op := gqlclient.NewOperation("mutation SetMembershipRole ($membership_id: UUID!, $role: membership_role!) {\n\tset_membership_role(input: {membership_id:$membership_id,role:$role}) {\n\t\tid\n\t}\n}\n")
	op.Var("membership_id", membership_id)
	op.Var("role", role)
	var respData struct {
		Set_membership_role *membership_id_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Set_membership_role, err
}

func DeleteServiceAccount(client *gqlclient.Client, ctx context.Context, id UUID) (delete_service_account *success_payload, err error) {
	op := gqlclient.NewOperation("mutation DeleteServiceAccount ($id: UUID!) {\n\tdelete_service_account(input: {user_id:$id}) {\n\t\tsuccess\n\t}\n}\n")
	op.Var("id", id)
	var respData struct {
		Delete_service_account *success_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Delete_service_account, err
}

func GetTenantUser(client *gqlclient.Client, ctx context.Context, id uuid) (tenant_user []user_view_same_tenant, err error) {
	op := gqlclient.NewOperation("query GetTenantUser ($id: uuid!) {\n\tuser_view_same_tenant(where: {id:{_eq:$id}}) {\n\t\tid\n\t\tfirst_name\n\t\tlast_name\n\t\taccount_type\n\t\tmemberships {\n\t\t\tid\n\t\t\ttenant_id\n\t\t\trole_detail {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n}\n")
	op.Var("id", id)
	var respData struct {
		User_view_same_tenant []user_view_same_tenant
	}
	err = client.Execute(ctx, op, &respData)
	return respData.User_view_same_tenant, err
}

func GetTenantUserIdsWithSimilarUsername(client *gqlclient.Client, ctx context.Context, similar string) (tenant_user []user_view_same_tenant, err error) {
	op := gqlclient.NewOperation("query GetTenantUserIdsWithSimilarUsername ($similar: String!) {\n\tuser_view_same_tenant(where: {username:{_similar:$similar}}) {\n\t\tid\n\t}\n}\n")
	op.Var("similar", similar)
	var respData struct {
		User_view_same_tenant []user_view_same_tenant
	}
	err = client.Execute(ctx, op, &respData)
	return respData.User_view_same_tenant, err
}

func AuthRoles(client *gqlclient.Client, ctx context.Context) (auth_roles []auth_role, err error) {
	op := gqlclient.NewOperation("query AuthRoles {\n\tauth_role {\n\t\tname\n\t\tid\n\t}\n}\n")
	var respData struct {
		Auth_role []auth_role
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Auth_role, err
}

func CurrentUser(client *gqlclient.Client, ctx context.Context) (current_user []user, err error) {
	op := gqlclient.NewOperation("query CurrentUser {\n\tuser {\n\t\tid\n\t\tusername\n\t\tfirst_name\n\t\tdefault_membership {\n\t\t\tid\n\t\t\ttenant {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n}\n")
	var respData struct {
		User []user
	}
	err = client.Execute(ctx, op, &respData)
	return respData.User, err
}

func CreateAPIKey(client *gqlclient.Client, ctx context.Context, user_id UUID, name string, expires_at *DateTime, default_tenant_id *UUID) (create_api_key *create_api_key_payload, err error) {
	op := gqlclient.NewOperation("mutation CreateAPIKey ($user_id: UUID!, $name: String!, $expires_at: DateTime, $default_tenant_id: UUID) {\n\tcreate_api_key(input: {user_id:$user_id,name:$name,expires_at:$expires_at,default_tenant_id:$default_tenant_id}) {\n\t\tid\n\t\tkey\n\t}\n}\n")
	op.Var("user_id", user_id)
	op.Var("name", name)
	op.Var("expires_at", expires_at)
	op.Var("default_tenant_id", default_tenant_id)
	var respData struct {
		Create_api_key *create_api_key_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Create_api_key, err
}

func DeleteAPIKey(client *gqlclient.Client, ctx context.Context, id UUID) (delete_api_key *success_payload, err error) {
	op := gqlclient.NewOperation("mutation DeleteAPIKey ($id: UUID!) {\n\tdelete_api_key(input: {key_id:$id}) {\n\t\tsuccess\n\t}\n}\n")
	op.Var("id", id)
	var respData struct {
		Delete_api_key *success_payload
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Delete_api_key, err
}

func APIKeysByUser(client *gqlclient.Client, ctx context.Context, user_id uuid) (auth_api_key []Auth_api_key, err error) {
	op := gqlclient.NewOperation("query APIKeysByUser ($user_id: uuid!) {\n\tauth_api_key(where: {user_id:{_eq:$user_id}}) {\n\t\tid\n\t\tname\n\t\tcreated\n\t\texpires_at\n\t\tdefault_tenant_id\n\t\tuser_id\n\t\tkey_hash\n\t}\n}\n")
	op.Var("user_id", user_id)
	var respData struct {
		Auth_api_key []Auth_api_key
	}
	err = client.Execute(ctx, op, &respData)
	return respData.Auth_api_key, err
}