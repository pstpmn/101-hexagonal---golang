package env

func GetMysql() (env map[string]string) {
	env = make(map[string]string)
	env["HOST"] = "0.0.0.0"
	env["PORT"] = "3306"
	env["USER"] = "superuser"
	env["PASS"] = "superuser"
	env["DBNAME"] = "members_service"
	env["TIMEZONE"] = "Asia/Bangkok"
	return env
}
