package dashboard

import (
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard/about"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard/databases"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard/destinations"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard/profile"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard/summary"
	"github.com/labstack/echo/v4"
)

func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	summary.MountRouter(parent.Group(""), mids, servs)
	databases.MountRouter(parent.Group("/databases"), mids, servs)
	destinations.MountRouter(parent.Group("/destinations"), mids, servs)
	profile.MountRouter(parent.Group("/profile"), mids, servs)
	about.MountRouter(parent.Group("/about"), mids, servs)
}
