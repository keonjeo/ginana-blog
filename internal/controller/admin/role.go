package admin

import "ginana-blog/internal/model"

func (c *CAdmin) GetRoleList() (err error) {
	roles, err := c.Svc.GetRoles(c.Pager)
	if err != nil {
		return
	}
	c.Ctx.ViewData("data", roles)
	c.setHeadMetas("角色列表")
	c.Ctx.View("admin/role/list.html")
	return
}

func (c *CAdmin) GetRoleAdd() (err error) {
	polices, err := c.Svc.GetAllPolices()
	if err != nil {
		return
	}
	c.Ctx.ViewData("polices", polices)
	c.setHeadMetas("角色创建")
	c.Ctx.View("admin/role/add.html")
	return
}

func (c *CAdmin) PostRoleAdd() (err error) {
	req := new(model.CreateRoleReq)
	if err = c.Ctx.ReadForm(req); err != nil {
		return
	}
	if err = c.Valid(req); err != nil {
		return
	}
	if _, err = c.Svc.CreateRole(req); err != nil {
		return
	}
	c.setHeadMetas("角色创建")
	c.ShowMsg("角色已创建")
	return
}

func (c *CAdmin) GetRoleEditBy(id int64) (err error) {
	if id == 1 && c.UserID != 1 {
		return c.Hm.GetMessage(500, "禁止修改管理员角色")
	}
	role, err := c.Svc.GetRole(id)
	if err != nil {
		return
	}
	c.Ctx.ViewData("data", role)
	polices, err := c.Svc.GetAllPolices()
	if err != nil {
		return
	}
	c.Ctx.ViewData("polices", polices)
	c.setHeadMetas("角色编辑")
	c.Ctx.View("admin/role/edit.html")
	return
}

func (c *CAdmin) PostRoleEditBy(id int64) (err error) {
	if id == 1 && c.UserID != 1 {
		return c.Hm.GetMessage(500, "禁止修改管理员角色")
	}
	req := new(model.UpdateRoleReq)
	if err = c.Ctx.ReadForm(req); err != nil {
		return
	}
	req.ID = id
	if err = c.Valid(req); err != nil {
		return
	}
	if _, err = c.Svc.UpdateRole(req); err != nil {
		return
	}
	c.setHeadMetas("角色更新")
	c.ShowMsg("角色已更新")
	return
}

func (c *CAdmin) GetRoleDeleteBy(id int64) (err error) {
	if id == 1 && c.UserID != 1 {
		return c.Hm.GetMessage(500, "禁止删除管理员角色")
	}
	if err = c.Svc.DeleteRole(id); err != nil {
		return
	}
	c.setHeadMetas("删除角色")
	c.ShowMsg("角色已删除")
	return
}
