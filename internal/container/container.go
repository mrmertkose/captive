package container

import "captive/internal/controller"

type Repositories struct {

	//UserRepo      repository.UserRepositoryI
	//UserLoginRepo repository.UserLoginRepositoryI
	//UserAgentRepo repository.UserAgentRepositoryI
	//FileRepo      repository.FileRepositoryI
	//FileLogRepo   repository.FileLogRepositoryI
	//ProfileRepo   repository.ProfileRepositoryI
	//SystemRepo    repository.SystemRepositoryI
}

type Services struct {
}

type Controllers struct {
	HomeCntrl *controller.HomeController
	//AuthCntrl    *controller.AuthController
	//UserCntrl    *controller.UserController
	//SystemCntrl  *controller.SystemController
	//FileCntrl    *controller.FileController
	//ProfileCntrl *controller.ProfileController
	//LogCntrl     *controller.LogController
}

type Container struct {
	Repositories Repositories
	Services     Services
	Controllers  Controllers
}

// db inject
func NewContainer() *Container {

	repos := Repositories{
		//UserRepo:      repository.NewUserRepository(mysql),
		//UserLoginRepo: repository.NewUserLoginRepository(mysql),
		//UserAgentRepo: repository.NewUserAgentRepository(mysql),
		//FileRepo:      repository.NewFileRepository(mysql),
		//FileLogRepo:   repository.NewFileLogRepository(mysql),
		//ProfileRepo:   repository.NewProfileRepository(mysql),
		//SystemRepo:    repository.NewSystemRepository(mysql),
	}

	services := Services{}

	controllers := Controllers{
		HomeCntrl: controller.NewHomeController(),
		//AuthCntrl:    controller.NewAuthController(authDeps),
		//UserCntrl:    controller.NewUserController(repos.UserRepo),
		//SystemCntrl:  controller.NewSystemController(repos.SystemRepo),
		//FileCntrl:    controller.NewFileController(repos.FileRepo),
		//ProfileCntrl: controller.NewProfileController(repos.ProfileRepo, redisCli),
		//LogCntrl:     controller.NewLogController(repos.FileLogRepo),
	}

	return &Container{
		Repositories: repos,
		Services:     services,
		Controllers:  controllers,
	}
}
