package main

import (
	"101-hexagonal-golang/env"
	infra "101-hexagonal-golang/infrastructure/mysql"
	repo "101-hexagonal-golang/infrastructure/mysql/repositories"
	"101-hexagonal-golang/internal/core/usecase"
	handles "101-hexagonal-golang/internal/handles"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	env := env.GetMysql()
	conn, _ := infra.ConnMysql(env)
	infra.Migration(conn)
	repo := repo.NewMembersRepo(conn)
	usecase := usecase.NewMembersUseCase(repo)

	handles.NewMembersHandles(usecase, app)
	// router
	app.Listen(":3000")
}
