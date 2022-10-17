package handles

import (
	"101-hexagonal-golang/internal/core/ports"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

type Member struct {
	MemberId int `json:"memberId" xml:"memberId" form:"memberId" validate:"required|int"`
}

type handles struct {
	MemberUseCase ports.MembersUseCase
}

func NewMembersHandles(memberUseCase ports.MembersUseCase, app *fiber.App) *handles {
	handles := &handles{
		MemberUseCase: memberUseCase,
	}
	app.Get("/", handles.HelloWorld)
	app.Post("/member", handles.FindMemberById)
	return handles
}

func (h *handles) HelloWorld(c *fiber.Ctx) error {
	data, err := h.MemberUseCase.GetById(1)
	fmt.Println(data, err)
	return c.SendString("I'm a GET request!")
}

func (h *handles) FindMemberById(c *fiber.Ctx) error {
	m := new(Member)
	c.BodyParser(m)
	v := validate.Struct(m)
	if !v.Validate() { // validate ok
		return c.Status(400).JSON(map[string]interface{}{
			"status":  false,
			"message": v.Errors.One(),
		})
	}

	// find member by id
	data, err := h.MemberUseCase.GetById(m.MemberId)
	if err != nil {
		return c.Status(200).JSON(map[string]interface{}{
			"status":  false,
			"message": fmt.Sprint(err),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"status":  true,
		"message": "Successful",
		"result": map[string]interface{}{
			"memberData": data,
		},
	})
}
