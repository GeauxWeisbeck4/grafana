---
aliases:
  - /docs/grafana/latest/administration/manage-users-and-permissions/manage-dashboard-permissions/
  - /docs/grafana/latest/permissions/dashboard_folder_permissions/
title: Manage dashboard permissions
weight: 500
---

# Manage dashboard permissions

Dashboard and dasboard folder permissions enable you to grant a viewer the ability to edit and save dashboard changes, or limit an editor's permission to modify a dashboard.

For more information about dashboard permissions, refer to [Dashboard permissions]({{< relref "../about-users-and-permissions/#dashboard-permissions" >}}).

## Grant dashboard folder permissions

When you grant user permissions for folders, that setting applies to all dashboards contained in the folder. Consider using this approach to assigning dashboard permissions when you have users or teams who require access to groups of related dashboards.

### Before you begin

- Ensure you have organization administrator privileges
- Identify the dashboard folder permissions you want to modify and the users or teams to which you want to grant access. For more information about dashboard permissions, refer to [Dashboard permissions]({{< relref "../about-users-and-permissions/#dashboard-permissions" >}}).

**To grant dashboard folder permissions**:

1. Sign in to Grafana as an organization administrator.
2. In the sidebar, hover your mouse over the **Dashboards** (squares) icon and click **Browse**.
3. Hover your mouse cursor over a folder and click **Go to folder**.
4. Click the **Permissions** tab, and then click **Add Permission**.
5. In the **Add Permission For** dropdown menu, select **User**, **Team**, or one of the role options.
6. Select the user or team.

   If you select a role option, you do not select a user or team.

7. Select the permission and click **Save**.

## Grant dashboard permissions

When you grant dashboard folder permissions, that setting applies to all dashboards in the folder. For a more granular approach to assigning permissions, you can also assign user permissions to individual dashboards.

For example, if a user with the viewer organization role requires editor (or admin) access to a dashboard, you can assign those elevated permissions on an individual basis.

> **Note**: If you have assigned a user dashboard folder permissions, you cannot also assign the user permission to dashboards contained in the folder.

Grant dashboard permissions when you want to restrict or enhance dashboard access for users who do not have permissions defined in the associated dashboard folder.

### Before you begin

- Ensure you have organization administrator privileges
- Identify the dashboard permissions you want to modify and the users or teams to which you want to grant access

**To grant dashboard permissions**:

1. Sign in to Grafana as an organization administrator.
1. In the sidebar, hover your mouse over the **Dashboards** (squares) icon and click **Browse**.
1. Open a dashboard.
1. In the top right corner of the dashboard, click **Dashboard settings** (the cog icon).
1. Click **Permissions** and then click **Add Permission**.
1. In the **Add Permission For** dropdown menu, select **User** or **Team**.
1. Select the user or team.
1. Select the permission and click **Save**.

## Enable viewers to edit (but not save) dashboards and use Explore

By default, the viewer organization role does not allow viewers to create dashboards or use the Explore feature. However, by modifying a configuration setting, you can allow viewers to edit a panel and make changes to a dashboard but not save those changes. This setting also enables viewers to use the Explore feature.

This modification is useful for public Grafana installations where you want anonymous users to be able to edit panels and queries but not save or create new dashboards.

> **Note**: If you use Grafana Enterprise and customize users' permissions using RBAC, the RBAC permissions override the functionality enabled by the `viewers_can_edit` flag.

### Before you begin

- Ensure that you have access to the Grafana server

**To enable viewers to preview dashboards and use Explore**:

1. Open the Grafana configuration file.

   For more information about the Grafana configuration file and its location, refer to [Configuration]({{< relref "../../configuration/" >}}).

1. Locate the `viewers_can_edit` parameter.
1. Set the `viewers_can_edit` value to `true`.
1. Save your changes and restart Grafana.

## Edit dashboard permissions

Edit dashboard permissions when you are want to enhance or restrict a user's access to a dashboard. For more information about dashboard permissions, refer to [Dashboard permissions]({{< relref "../about-users-and-permissions/#dashboard-permissions" >}}).

### Before you begin

- Identify the dashboard and user permission you want to change
- Ensure you have organization administrator privileges

**To edit dashboard permissions**:

1. Sign in to Grafana as an organization administrator.
1. In the sidebar, hover your mouse over the **Dashboards** (squares) icon and click **Browse**.
1. Open a dashboard.
1. In the top-right corner of the dashboard, click **Dashboard settings** (the cog icon).
1. Click **Permissions**.
1. In the dropdown, update the permissions, and click **Save**.

## Restrict access to dashboards

Grafana applies the highest permission a given user has to access a resource like a dashboard, so if you want to prevent a user from accessing a folder or dashboard you need to consider the user's organization role, folder permissions, and dashboard permissions.

- You cannot override organization administrator permissions. Organization administrators have access to all organization resources.
- User permissions set for a dashboard folder propagate to dashboards contained in the folder.
- A lower permission level does not affect access if a more general rule exists with a higher permission.

Refer to the following examples to understand how organization and dashboard permissions impact a user's access to dashboards.

### Example 1

In this example, user1 has the editor organization role.

Dashboard permissions settings:

- Everyone with Editor role can edit
- user1 is set to `view`

Result: User1 has edit permissions because the user's organization role is Editor.

### Example 2

In this example, user1 has the viewer organization role and is a member of team1.

Dashboard permissions settings:

- Everyone with Viewer role can view
- user1 is set to `edit`
- team1 is set to `admin`

Result: User1 has administrator permissions for the dashboard because user1 is a member of team1.

### Example 3

In this example, user1 has the viewer organization role.

Dashboard permissions settings:

- user1 is set to `admin`, which is inherited from the permissions set in parent folder
- user1 is set to `edit`

Result: You receive an error message that cannot override a higher permission with a lower permission in the same dashboard. User1 has administrator permissions.

> Refer to [Role-based access Control]({{< relref "../../../enterprise/access-control/_index.md" >}}) in Grafana Enterprise to understand how to use RBAC permissions to restrict access to dashboards, folders, administrative functions, and other resources.
