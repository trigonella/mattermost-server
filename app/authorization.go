// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"net/http"
	"strings"

	"github.com/trigonella/mattermost-server/v5/mlog"
	"github.com/trigonella/mattermost-server/v5/model"
)

func (a *App) MakePermissionError(permissions []*model.Permission) *model.AppError {
	permissionsStr := "permission="
	for _, permission := range permissions {
		permissionsStr += permission.Id
		permissionsStr += ","
	}
	return model.NewAppError("Permissions", "api.context.permissions.app_error", nil, "userId="+a.Session().UserId+", "+permissionsStr, http.StatusForbidden)
}

func (a *App) SessionHasPermissionTo(session model.Session, permission *model.Permission) bool {
	if session.IsUnrestricted() {
		return true
	}
	return a.RolesGrantPermission(session.GetUserRoles(), permission.Id)
}

func (a *App) SessionHasPermissionToAny(session model.Session, permissions []*model.Permission) bool {
	for _, perm := range permissions {
		if a.SessionHasPermissionTo(session, perm) {
			return true
		}
	}
	return false
}

func (a *App) SessionHasPermissionToTeam(session model.Session, teamId string, permission *model.Permission) bool {
	if teamId == "" {
		return false
	}
	if session.IsUnrestricted() {
		return true
	}

	teamMember := session.GetTeamByTeamId(teamId)
	if teamMember != nil {
		if a.RolesGrantPermission(teamMember.GetRoles(), permission.Id) {
			return true
		}
	}

	return a.RolesGrantPermission(session.GetUserRoles(), permission.Id)
}

func (a *App) SessionHasPermissionToChannel(session model.Session, channelId string, permission *model.Permission) bool {
	if channelId == "" {
		return false
	}

	ids, err := a.Srv().Store.Channel().GetAllChannelMembersForUser(session.UserId, true, true)

	var channelRoles []string
	if err == nil {
		if roles, ok := ids[channelId]; ok {
			channelRoles = strings.Fields(roles)
			if a.RolesGrantPermission(channelRoles, permission.Id) {
				return true
			}
		}
	}

	channel, appErr := a.GetChannel(channelId)
	if appErr != nil && appErr.StatusCode == http.StatusNotFound {
		return false
	}

	if session.IsUnrestricted() {
		return true
	}

	if appErr == nil && channel.TeamId != "" {
		return a.SessionHasPermissionToTeam(session, channel.TeamId, permission)
	}

	return a.SessionHasPermissionTo(session, permission)
}

func (a *App) SessionHasPermissionToChannelByPost(session model.Session, postId string, permission *model.Permission) bool {
	if channelMember, err := a.Srv().Store.Channel().GetMemberForPost(postId, session.UserId); err == nil {

		if a.RolesGrantPermission(channelMember.GetRoles(), permission.Id) {
			return true
		}
	}

	if channel, err := a.Srv().Store.Channel().GetForPost(postId); err == nil {
		if channel.TeamId != "" {
			return a.SessionHasPermissionToTeam(session, channel.TeamId, permission)
		}
	}

	return a.SessionHasPermissionTo(session, permission)
}

func (a *App) SessionHasPermissionToCategory(session model.Session, userId, teamId, categoryId string) bool {
	if a.SessionHasPermissionTo(session, model.PERMISSION_EDIT_OTHER_USERS) {
		return true
	}
	category, err := a.GetSidebarCategory(categoryId)
	return err == nil && category != nil && category.UserId == session.UserId && category.UserId == userId && category.TeamId == teamId
}

func (a *App) SessionHasPermissionToUser(session model.Session, userId string) bool {
	if userId == "" {
		return false
	}
	if session.IsUnrestricted() {
		return true
	}

	if session.UserId == userId {
		return true
	}

	if a.SessionHasPermissionTo(session, model.PERMISSION_EDIT_OTHER_USERS) {
		return true
	}

	return false
}

func (a *App) SessionHasPermissionToUserOrBot(session model.Session, userId string) bool {
	if session.IsUnrestricted() {
		return true
	}
	if a.SessionHasPermissionToUser(session, userId) {
		return true
	}

	if err := a.SessionHasPermissionToManageBot(session, userId); err == nil {
		return true
	}

	return false
}

func (a *App) HasPermissionTo(askingUserId string, permission *model.Permission) bool {
	user, err := a.GetUser(askingUserId)
	if err != nil {
		return false
	}

	roles := user.GetRoles()

	return a.RolesGrantPermission(roles, permission.Id)
}

func (a *App) HasPermissionToTeam(askingUserId string, teamId string, permission *model.Permission) bool {
	if teamId == "" || askingUserId == "" {
		return false
	}
	teamMember, _ := a.GetTeamMember(teamId, askingUserId)
	if teamMember != nil && teamMember.DeleteAt == 0 {
		if a.RolesGrantPermission(teamMember.GetRoles(), permission.Id) {
			return true
		}
	}
	return a.HasPermissionTo(askingUserId, permission)
}

func (a *App) HasPermissionToChannel(askingUserId string, channelId string, permission *model.Permission) bool {
	if channelId == "" || askingUserId == "" {
		return false
	}

	channelMember, err := a.GetChannelMember(channelId, askingUserId)
	if err == nil {
		roles := channelMember.GetRoles()
		if a.RolesGrantPermission(roles, permission.Id) {
			return true
		}
	}

	var channel *model.Channel
	channel, err = a.GetChannel(channelId)
	if err == nil {
		return a.HasPermissionToTeam(askingUserId, channel.TeamId, permission)
	}

	return a.HasPermissionTo(askingUserId, permission)
}

func (a *App) HasPermissionToChannelByPost(askingUserId string, postId string, permission *model.Permission) bool {
	if channelMember, err := a.Srv().Store.Channel().GetMemberForPost(postId, askingUserId); err == nil {
		if a.RolesGrantPermission(channelMember.GetRoles(), permission.Id) {
			return true
		}
	}

	if channel, err := a.Srv().Store.Channel().GetForPost(postId); err == nil {
		return a.HasPermissionToTeam(askingUserId, channel.TeamId, permission)
	}

	return a.HasPermissionTo(askingUserId, permission)
}

func (a *App) HasPermissionToUser(askingUserId string, userId string) bool {
	if askingUserId == userId {
		return true
	}

	if a.HasPermissionTo(askingUserId, model.PERMISSION_EDIT_OTHER_USERS) {
		return true
	}

	return false
}

func (a *App) RolesGrantPermission(roleNames []string, permissionId string) bool {
	roles, err := a.GetRolesByNames(roleNames)
	if err != nil {
		// This should only happen if something is very broken. We can't realistically
		// recover the situation, so deny permission and log an error.
		mlog.Error("Failed to get roles from database with role names: "+strings.Join(roleNames, ",")+" ", mlog.Err(err))
		return false
	}

	for _, role := range roles {
		if role.DeleteAt != 0 {
			continue
		}

		permissions := role.Permissions
		for _, permission := range permissions {
			if permission == permissionId {
				return true
			}
		}
	}

	return false
}

// SessionHasPermissionToManageBot returns nil if the session has access to manage the given bot.
// This function deviates from other authorization checks in returning an error instead of just
// a boolean, allowing the permission failure to be exposed with more granularity.
func (a *App) SessionHasPermissionToManageBot(session model.Session, botUserId string) *model.AppError {
	existingBot, err := a.GetBot(botUserId, true)
	if err != nil {
		return err
	}
	if session.IsUnrestricted() {
		return nil
	}

	if existingBot.OwnerId == session.UserId {
		if !a.SessionHasPermissionTo(session, model.PERMISSION_MANAGE_BOTS) {
			if !a.SessionHasPermissionTo(session, model.PERMISSION_READ_BOTS) {
				// If the user doesn't have permission to read bots, pretend as if
				// the bot doesn't exist at all.
				return model.MakeBotNotFoundError(botUserId)
			}
			return a.MakePermissionError([]*model.Permission{model.PERMISSION_MANAGE_BOTS})
		}
	} else {
		if !a.SessionHasPermissionTo(session, model.PERMISSION_MANAGE_OTHERS_BOTS) {
			if !a.SessionHasPermissionTo(session, model.PERMISSION_READ_OTHERS_BOTS) {
				// If the user doesn't have permission to read others' bots,
				// pretend as if the bot doesn't exist at all.
				return model.MakeBotNotFoundError(botUserId)
			}
			return a.MakePermissionError([]*model.Permission{model.PERMISSION_MANAGE_OTHERS_BOTS})
		}
	}

	return nil
}
