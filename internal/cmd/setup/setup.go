package setup

func Setup(buildMode, buildTarget string) {
	switch buildMode {
	case "full":
		prep()
		check(buildTarget)
		generate(buildTarget)
		install(buildTarget)
		test(buildTarget)

	case "prep":
		prep()

	case "check":
		check(buildTarget)

	case "generate":
		generate(buildTarget)

	case "install":
		install(buildTarget)

	case "test":
		test(buildTarget)

	case "update":
		update()

	case "upgrade":
		upgrade()
	}
}
