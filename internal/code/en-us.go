package code

var enUSText = map[int]string{
	ServerError:        "Internal server error",
	TooManyRequests:    "Too many requests",
	ParamBindError:     "Parameter error",
	AuthorizationError: "Authorization error",
	UrlSignError:       "URL signature error",
	CacheSetError:      "Failed to set cache",
	CacheGetError:      "Failed to get cache",
	CacheDelError:      "Failed to del cache",
	CacheNotExist:      "Cache does not exist",
	ResubmitError:      "Please do not submit repeatedly",
	HashIdsEncodeError: "HashID encryption failed",
	HashIdsDecodeError: "HashID decryption failed",
	RBACError:          "No access",
	RedisConnectError:  "Failed to connection Redis",
	MySQLConnectError:  "Failed to connection MySQL",
	WriteConfigError:   "Failed to write configuration file",
	SendEmailError:     "Failed to send mail",
	MySQLExecError:     "SQL execution failed",
	GoVersionError:     "Go Version mismatch",
	SocketConnectError: "Socket not connected",
	SocketSendError:    "Socket message sending failed",

	AuthorizedCreateError:    "Failed to create caller",
	AuthorizedListError:      "Failed to get caller list",
	AuthorizedDeleteError:    "Failed to delete caller",
	AuthorizedUpdateError:    "Failed to update caller",
	AuthorizedDetailError:    "Failed to get caller details",
	AuthorizedCreateAPIError: "Failed to create caller API address",
	AuthorizedListAPIError:   "Failed to get caller API address list",
	AuthorizedDeleteAPIError: "Failed to delete caller API address",

	AdminCreateError:             "Failed to create administrator",
	AdminListError:               "Failed to get administrator list",
	AdminDeleteError:             "Failed to delete administrator",
	AdminUpdateError:             "Failed to update administrator",
	AdminResetPasswordError:      "Reset password failed",
	AdminLoginError:              "Login failed",
	AdminLogOutError:             "Exit failed",
	AdminModifyPasswordError:     "Failed to modify password",
	AdminModifyPersonalInfoError: "Failed to modify personal information",
	AdminMenuListError:           "Failed to get administrator menu authorization list",
	AdminMenuCreateError:         "Administrator menu authorization failed",
	AdminOfflineError:            "Offline administrator failed",
	AdminDetailError:             "Failed to get personal information",

	MenuCreateError:       "Failed to create menu",
	MenuUpdateError:       "Failed to update menu",
	MenuDeleteError:       "Failed to delete menu",
	MenuListError:         "Failed to get menu list",
	MenuDetailError:       "Failed to get menu details",
	MenuCreateActionError: "Failed to create menu action",
	MenuListActionError:   "Failed to get menu action list",
	MenuDeleteActionError: "Failed to delete menu action",

	CronCreateError:  "Failed to create cron",
	CronUpdateError:  "Failed to update menu",
	CronListError:    "Failed to get cron list",
	CronDetailError:  "Failed to get cron detail",
	CronExecuteError: "Failed to execute cron",

	UmsAdminRegisterError:     "Failed to register user",
	UmsAdminAuthorizedError:   "Failed to authorized user",
	UmsAdminRefreshTokenError: "Failed to refresh token",
	UmsAdminGetUsernameError:  "Failed to get username",
	UmsAdminGetMenuListError:  "Failed to get menu list",
	UmsAdminGetRoleListError:  "Failed to get role list",
	UmsAdminGetListError:      "Failed to get user list",
	UmsAdminGetItemError:      "Failed to get user item",
	UmsAdminUpdateError:       "Failed to update admin",
}
