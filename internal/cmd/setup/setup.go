package setup

func Setup(buildMode, buildTarget string) {
	switch buildMode {
	case "full":
		Prep()
		Check(buildTarget)
		Generate(buildTarget)
		Install(buildTarget)
		Test(buildTarget)
	case "prep":
		Prep()
	case "check":
		Check(buildTarget)
	case "generate":
		Generate(buildTarget)
	case "install":
		Install(buildTarget)
	case "test":
		Test(buildTarget)
	case "update":
		Update()
	case "upgrade":
		Upgrade()
	}
}
