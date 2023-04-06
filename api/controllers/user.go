package controllers

type UserController struct {
	BaseController
}

func (c *UserController) GetUserInfo() {
	rq := middleware.LoadContext(c.Ctx)
	type RequestData struct {
		ToUserId int `form:"toUserId" valid:"Required;" default:"0"`
	}
	d := new(RequestData)
	if !c.CheckFormAndGet(d) {
		return
	}

	rq.DataOrError(nil, nil)
}
