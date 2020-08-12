package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/therecipe/qt/internal/utils"
)

var buildVersion = "no build version"

func ParseFlags() bool {
	if runtime.GOOS == "windows" {
		utils.Log.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	}
	var (
		debug      = flag.Bool("debug", false, "print debug logs")
		help       = flag.Bool("help", false, "print help")
		p          = flag.Int("p", runtime.NumCPU(), "specify the number of cpu's to be used")
		qt_api     = flag.String("qt_api", "", "specify the api version to be used")
		qt_dir     = flag.String("qt_dir", utils.QT_DIR(), "export QT_DIR")
		qt_version = flag.String("qt_version", utils.QT_VERSION(), "export QT_VERSION")
		version    = flag.Bool("version", false, "print build version (if available)")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	runtime.GOMAXPROCS(*p)

	var forDocker bool
	if d := flag.Lookup("docker"); d != nil {
		forDocker = d.Value.String() == "true"
	}

	if api := *qt_api; api != "" {
		os.Setenv("QT_API", api)
	}

	_, err := exec.LookPath("go")
	if api := utils.QT_API(""); api != "" && err == nil && !utils.QT_DOCKER() && !forDocker {
		if utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+api, "-find", "get doc dir") == "" {
			utils.Log.Errorf("invalid api version provided: '%v'", api)
			fmt.Println("valid api versions are:") //TODO: if only one api version is available, use it
			if !utils.UseGOMOD("") {
				if o := utils.GoListOptional("{{join .Imports \"|\"}}", "github.com/therecipe/qt/internal/binding/files/docs", "get doc dir"); o != "" {
					for _, v := range strings.Split(o, "|") {
						fmt.Println(strings.TrimPrefix(strings.TrimSpace(strings.Replace(v, "'", "", -1)), "github.com/therecipe/qt/internal/binding/files/docs/"))
					}
				}
			} else {
				wg := new(sync.WaitGroup)
				for mid := 5; mid <= 15; mid++ {
					for min := 0; min <= 5; min++ {
						wg.Add(1)
						go func(mid, min int) {
							v := fmt.Sprintf("5.%v.%v", mid, min)
							if utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+v, "-find", "get doc dir") != "" {
								fmt.Println(v)
							}
							wg.Done()
						}(mid, min)
					}
				}
				wg.Wait()
			}
			os.Exit(1)
		}
	}

	if dir := *qt_dir; dir != utils.QT_DIR() {
		os.Setenv("QT_DIR", dir)
	}

	if version := *qt_version; version != utils.QT_VERSION() {
		os.Setenv("QT_VERSION", version)
	}

	if *version {
		println(buildVersion)
		os.Exit(0)
	}

	return help != nil && *help
}

func InitEnv(target string, docker bool, path string) {
	if docker || target != runtime.GOOS || ((runtime.GOARCH != "amd64" || utils.GOARCH() != "amd64") && runtime.GOOS != "windows") {
		return
	}

	switch runtime.GOOS {
	case "freebsd":
		return
	case "linux":
		if utils.QT_PKG_CONFIG() || utils.QT_MXE() || utils.QT_FELGO() || utils.QT_STATIC() {
			return
		}
	case "darwin":
		if utils.QT_PKG_CONFIG() || utils.QT_HOMEBREW() || utils.QT_MACPORTS() || utils.QT_NIX() || utils.QT_FELGO() || utils.QT_STATIC() {
			return
		}
	case "windows":
		if utils.QT_MSYS2() {
			return
		}

		defer func() {
			if utils.QT_MSVC() {
				var data string
				if utils.QT_DOCKER() {
					msvc_version := "14.16.27023"
					msvc_path := fmt.Sprintf("C:\\Program Files (x86)\\Microsoft Visual Studio\\2017\\BuildTools\\VC\\Tools\\MSVC\\%v", msvc_version)
					mswk_version := "10.0.17763.0"
					mswk_path := func(m string) string {
						return fmt.Sprintf("C:\\Program Files (x86)\\Windows Kits\\10\\%v\\%v", m, mswk_version)
					}
					msvc_arch := "x64"
					if utils.GOARCH() == "386" {
						msvc_arch = "x86"
					}
					data = fmt.Sprintf("@echo INCLUDE=%[1]v\\include;%[2]v\\ucrt;%[2]v\\shared;%[2]v\\um\r\n", msvc_path, mswk_path("include"))
					data += fmt.Sprintf("@echo LIB=%[1]v\\lib\\%[2]v;%[3]v\\ucrt\\%[2]v;%[3]v\\um\\%[2]v\r\n", msvc_path, msvc_arch, mswk_path("lib"))
					if utils.GOARCH() == "386" {
						data += fmt.Sprintf("@echo PATH=%[1]v\\bin\\HostX64\\%[2]v;%[1]v\\bin\\HostX64\\x64;%[3]v;%%PATH%%\r\n", msvc_path, msvc_arch, utils.MINGWTOOLSDIR())
					} else {
						data += fmt.Sprintf("@echo PATH=%[1]v\\bin\\HostX64\\%[2]v;%[3]v;%%PATH%%\r\n", msvc_path, msvc_arch, utils.MINGWTOOLSDIR())
					}
				} else {
					data = fmt.Sprintf("set PATH=%v;%%PATH%%\r\ncall \"%v\"\r\nset", utils.MINGWTOOLSDIR(), utils.GOVSVARSPATH())
				}
				utils.Save(filepath.Join(utils.GOBIN(), "qtenvexc.bat"), data)
				for _, s := range strings.Split(utils.RunCmdOptional(exec.Command(filepath.Join(utils.GOBIN(), "qtenvexc.bat")), fmt.Sprintf("run qtenvexc for %v on %v", target, runtime.GOOS)), "\r\n") {
					if !strings.Contains(s, "=") {
						continue
					}
					es := strings.Split(s, "=")
					os.Setenv(es[0], strings.Join(es[1:], "="))
				}
				utils.RemoveAll(filepath.Join(utils.GOBIN(), "qtenvexc.bat"))
			} else {
				qtenvPath := filepath.Join(filepath.Dir(utils.ToolPath("qmake", target)), "qtenv2.bat")
				for _, s := range strings.Split(utils.Load(qtenvPath), "\r\n") {
					if strings.HasPrefix(s, "set PATH") {
						os.Setenv("PATH", strings.TrimPrefix(strings.Replace(s, "%PATH%", os.Getenv("PATH"), -1), "set PATH="))
						break
					}
				}
				for i, dPath := range []string{filepath.Join(runtime.GOROOT(), "bin", "qtenv.bat"), filepath.Join(utils.GOBIN(), "qtenv.bat")} {
					sPath := qtenvPath
					existed := utils.ExistsFile(dPath)
					if existed {
						utils.RemoveAll(dPath)
					}
					err := os.Link(sPath, dPath)
					if i != 0 {
						continue
					}
					if err == nil {
						if !existed {
							utils.Log.Infof("successfully created %v symlink in your PATH (%v)", filepath.Base(dPath), dPath)
						}
					} else {
						utils.Log.Warnf("failed to create %v symlink in your PATH (%v); please use %v instead", filepath.Base(dPath), dPath, sPath)
					}
				}
			}
		}()
	default:
		return
	}

	if !utils.ExistsDir(utils.QT_DIR()) {
		qt_dir := strings.TrimSpace(utils.RunCmd(utils.GoList("{{.Dir}}", "github.com/therecipe/env_"+runtime.GOOS+"_amd64_"+strconv.Itoa(utils.QT_VERSION_NUM())[:3], "-find"), "get env dir"))

		switch runtime.GOOS {
		case "linux", "darwin", "windows":
			os.Setenv("QT_DIR", qt_dir)
		}

		var err error
		var link string
		switch runtime.GOOS {
		case "linux":
			link = filepath.Join("/var", "tmp", ".env_linux_amd64")
			utils.RemoveAll(link)
			err = os.Symlink(qt_dir, link)
		case "darwin":
			link = filepath.Join("/Applications", ".env_darwin_amd64")
			utils.RemoveAll(link)
			err = os.Symlink(qt_dir, link)
		case "windows":
			link = filepath.Join("C:\\", "Users", "Public", "env_windows_amd64")
			utils.RemoveAll(link)
			_, err = utils.RunCmdOptionalError(exec.Command("cmd", "/C", "mklink", "/J", link, qt_dir), "create symlink for env")

			if err == nil {
				link = filepath.Join("C:\\", "Users", "Public", "env_windows_amd64_Tools")
				utils.RemoveAll(link)
				_, err = utils.RunCmdOptionalError(exec.Command("cmd", "/C", "mklink", "/J", link, strings.TrimSpace(utils.RunCmd(utils.GoList("{{.Dir}}", "github.com/therecipe/env_"+runtime.GOOS+"_amd64_"+strconv.Itoa(utils.QT_VERSION_NUM())[:3]+"/Tools", "-find"), "get env dir"))), "create symlink for env")
			}
		}
		if err != nil && target == runtime.GOOS {
			if !(utils.ExistsFile(link) || utils.ExistsDir(link)) {
				if s := "failed to create env symlink; fallback to patching binaries instead"; utils.UseGOMOD(path) {
					utils.Log.WithError(err).Debug(s)
				}
				cmd := exec.Command("go", "run", "patch.go", qt_dir)
				cmd.Dir = qt_dir
				_, err = utils.RunCmdOptionalError(cmd, "patch env binaries")
				if err != nil {
					utils.Log.WithError(err).Warn("failed to patch binaries\r\nyou won't be able to simply \"go build/run\" your application, but qtdeploy-ing applications should work nevertheless")
				}
			}
		}

		var webenginearchive string
		var bindir string
		switch runtime.GOOS {
		case "linux":
			webenginearchive = filepath.Join(qt_dir, utils.QT_VERSION(), "gcc_64", "lib", "libQt5WebEngineCore.so."+utils.QT_VERSION()+".gz")
			bindir = filepath.Join(qt_dir, utils.QT_VERSION(), "gcc_64", "bin")
		case "darwin":
			webenginearchive = filepath.Join(qt_dir, utils.QT_VERSION(), "clang_64", "lib", "QtWebEngineCore.framework", "Versions", "Current", "QtWebEngineCore.gz")
			bindir = filepath.Join(qt_dir, utils.QT_VERSION(), "clang_64", "bin")
		case "windows":
			return
		}

		if utils.ExistsFile(webenginearchive) {
			_, err := utils.RunCmdOptionalError(exec.Command("gzip", "-d", webenginearchive), "uncompress webengine")
			if err != nil {
				utils.Log.WithError(err).Warn("failed to uncompress webengine, (do you use go modules?)")
			} else {
				utils.RemoveAll(webenginearchive)
			}
		}

		if len(bindir) != 0 && utils.UseGOMOD(path) && strings.Contains(qt_dir, "@") {
			filepath.Walk(bindir, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return err
				}
				if filepath.Ext(info.Name()) != ".conf" {
					utils.RunCmd(exec.Command("chmod", "+x", path), "restore exec permission")
				}
				return nil
			})
		}
	}
}

func Docker(arg []string, target, path string, writeCacheToHost bool) {
	virtual(arg, target, path, writeCacheToHost, true, "")
}

func Vagrant(arg []string, target, path string, writeCacheToHost bool, system string) {
	virtual(arg, target, path, writeCacheToHost, false, system)
}

func virtual(arg []string, target, path string, writeCacheToHost bool, docker bool, system string) {
	dUser := "therecipe/qt"
	if strings.Contains(target, ":") {
		dUser = strings.Split(target, ":")[0]
		target = strings.Split(target, ":")[1]
	}
	arg = append(arg, target)
	if system == "" {
		system = "linux"
	}
	android_arm64 := target == "android_arm64"

	var image string
	switch target {
	case "windows":
		if utils.QT_MXE_ARCH() == "amd64" || utils.QT_MSYS2_ARCH() == "amd64" {
			if utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC() {
				image = "windows_64_static"
			} else {
				image = "windows_64_shared"
			}
		} else {
			if utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC() {
				image = "windows_32_static"
			} else {
				image = "windows_32_shared"
			}
		}
		if !docker {
			image = target
		}

	case "linux", "rpi1", "rpi2", "rpi3", "js", "wasm", "darwin":
		image = target

	case "android", "android-emulator", "android_arm64":
		image = "android"

	case "sailfish", "sailfish-emulator":
		image = "sailfish"

	default:
		switch system {
		case "darwin":
		case "linux":
			if strings.Contains(target, "_") {
				image = target
			}
		case "windows":
		}
	}
	if image != "" {
		target = image
	}

	var args []string
	if docker {
		args = []string{"run", "--rm"}
	} else {
		if target == "pkg_config" {
			args = []string{"source /home/vagrant/.profile"}
		}
	}
	if runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
		u, err := user.Current()
		if err != nil {
			utils.Log.WithError(err).Error("failed to lookup current user")
		} else {
			args = append(args, "-e", fmt.Sprintf("IDUG=%v:%v", u.Uid, u.Gid))
		}
	}

	paths := make([]string, 0)

	if utils.UseGOMOD(path) {
		args = append(args, []string{"-v", fmt.Sprintf("%v:/media/%v", filepath.Dir(utils.GOMOD(path)), filepath.Base(filepath.Dir(utils.GOMOD(path))))}...)
	}

	for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
		if runtime.GOOS == "windows" {
			gp = "//" + strings.ToLower(gp[:1]) + gp[1:]
		}
		gp = strings.Replace(gp, "\\", "/", -1)
		gp = strings.Replace(gp, ":", "", -1)

		var pathprefix string
		if docker {
			args = append(args, []string{"-v", fmt.Sprintf("%v:/media/sf_GOPATH%v", gp, i)}...)
		} else if system == "windows" {
			pathprefix = "C:"
		}
		paths = append(paths, fmt.Sprintf(pathprefix+"/media/sf_GOPATH%v", i))
	}

	var gpfs string
	if docker {
		gpfs = "/home/user/work"
	} else {
		switch system {
		case "linux", "docker":
			gpfs = "/home/vagrant/gopath"
		case "darwin":
			gpfs = "/Users/vagrant/gopath"
		case "windows":
			gpfs = "C:/gopath"
		}
	}

	pathseperator := ":"
	if system == "windows" {
		pathseperator = ";"
	}
	gpath := strings.Join(paths, pathseperator)
	if utils.UseGOMOD(path) {
		gpath += pathseperator + gpfs
	} else {
		if writeCacheToHost {
			gpath += pathseperator + gpfs
			args = append(args, []string{"-e", "QT_STUB=true"}...) //TODO: won't work with wine images atm
		} else {
			if strings.Contains(path, "github.com/therecipe/qt/internal/examples") && !strings.Contains(path, "github.com/therecipe/qt/internal/examples/androidextras") {
				gpath += pathseperator + gpfs
			} else {
				gpath = gpfs + pathseperator + gpath
			}
		}
	}

	if utils.QT_DEBUG() {
		args = append(args, []string{"-e", "QT_DEBUG=true"}...) //TODO: won't work with wine images atm
	}

	if utils.QT_DEBUG_QML() {
		args = append(args, []string{"-e", "QT_DEBUG_QML=true"}...) //TODO: won't work with wine images atm
	}

	if utils.QT_DEBUG_CONSOLE() {
		args = append(args, []string{"-e", "QT_DEBUG_CONSOLE=true"}...) //TODO: won't work with wine images atm
	}

	if utils.QT_FAT() {
		args = append(args, []string{"-e", "QT_FAT=true"}...) //TODO: won't work with wine images atm
	}

	if api := utils.QT_API(""); api != "" {
		args = append(args, []string{"-e", "QT_API=" + api}...) //TODO: won't work with wine images atm
	}

	if utils.CI() {
		args = append(args, []string{"-e", "CI=true"}...)
	}

	if p, ok := os.LookupEnv("PKG_CONFIG_PATH"); ok {
		args = append(args, []string{"-e", "PKG_CONFIG_PATH=" + p}...) //TODO: won't work with wine images atm
	}

	if utils.QT_WEBKIT() {
		args = append(args, []string{"-e", "QT_WEBKIT=true"}...) //TODO: won't work with wine images atm
	}

	if utils.QT_NOT_CACHED() {
		args = append(args, []string{"-e", "QT_NOT_CACHED=true"}...) //TODO: won't work with wine images atm
	}

	if (target == "android" && utils.GOARCH() == "arm64") || android_arm64 {
		args = append(args, []string{"-e", "GOARCH=arm64"}...)
	}

	if v, ok := os.LookupEnv("GO111MODULE"); ok {
		args = append(args, []string{"-e", "GO111MODULE=" + v}...)
	}

	if v, ok := os.LookupEnv("GOPROXY"); ok {
		args = append(args, []string{"-e", "GOPROXY=" + v}...)
	}

	if v, ok := os.LookupEnv("GONOPROXY"); ok {
		args = append(args, []string{"-e", "GONOPROXY=" + v}...)
	}

	if v, ok := os.LookupEnv("GOSUMDB"); ok {
		args = append(args, []string{"-e", "GOSUMDB=" + v}...)
	}

	if v, ok := os.LookupEnv("GONOSUMBDB"); ok {
		args = append(args, []string{"-e", "GONOSUMBDB=" + v}...)
	}

	if v, ok := os.LookupEnv("GOPRIVATE"); ok {
		args = append(args, []string{"-e", "GOPRIVATE=" + v}...)
	}

	if docker {
		args = append(args, []string{"-e", "GOPATH=" + gpath}...)
	} else {
		args = append(args, []string{"-e", "QT_VAGRANT=true"}...) //TODO: won't work with wine images atm
		args = append(args, []string{"-e", "GOPATH='" + gpath + "'"}...)
	}

	if docker {
		args = append(args, []string{"-i", fmt.Sprintf("%v:%v", dUser, image)}...)
	} else {
		for i, a := range args {
			if a == "-e" {
				args[i] = "&&"
				args[i+1] = "export " + args[i+1]
			}
		}
		if args[0] == "&&" {
			args = args[1:]
		}
		args = append(args, "&&")

		if system == "windows" {
			if strings.Contains(target, "_") {
				args = append(args, []string{"C:/msys64/usr/bin/bash -l -c"}...)
				arg[0] = "'C:/gopath/bin/" + arg[0]
			} else {
				args = append(args, []string{"C:/windows/system32/cmd /C"}...)
			}
		}
	}

	args = append(args, arg...)

	if utils.UseGOMOD(path) {
		if docker {
			args = append(args, strings.Replace(path, filepath.Dir(filepath.Dir(utils.GOMOD(path))), "/media", -1))
		} else {
			//TODO: vagrant
		}
	} else {
		var found bool
		for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
			gp = filepath.Clean(gp)
			if strings.HasPrefix(path, gp) {
				if !docker && system == "windows" && strings.Contains(target, "_") {
					args = append(args, strings.Replace(strings.Replace(path, gp, fmt.Sprintf("C:/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
				} else {
					args = append(args, strings.Replace(strings.Replace(path, gp, fmt.Sprintf("/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
				}
				found = true
				break
			}
		}

		if !found && path != "" {
			utils.Log.Panicln("Project needs to be inside GOPATH; have:", path, "want:", utils.GOPATH())
		}
	}

	if docker {
		for i := range args {
			for _, t := range []string{"windows", "ubports", "linux", "js", "wasm", "android", "darwin"} {
				if strings.HasPrefix(args[i], t+"_") {
					args[i] = t
					break
				}
			}
		}

		utils.RunCmd(exec.Command("docker", args...), fmt.Sprintf("deploy binary for %v on %v with docker", target, runtime.GOOS))
	} else {
		switch target {
		case "ios-simulator":
			target = "ios"
		case "android-emulator":
			target = "android"
		}

		upCmd := exec.Command("vagrant", "up", "--no-provision", target)
		upCmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(upCmd, fmt.Sprintf("vagrant up for %v/%v on %v", system, target, runtime.GOOS))

		command := strings.Join(args, " ")
		command = strings.Replace(command, " pkg_config ", " linux ", -1)
		command = strings.Replace(command, " homebrew ", " desktop ", -1)
		for _, v := range []string{"windows_32_shared", "windows_32_static", "windows_64_shared", "windows_64_static"} {
			command = strings.Replace(command, " "+v+" ", " windows ", -1)
		}
		if system == "windows" && strings.Contains(target, "_") {
			command += "'"
		}

		var cmd *exec.Cmd
		if system == "windows" {
			cmd = exec.Command("vagrant", []string{"ssh", target, "--", command}...)
		} else {
			cmd = exec.Command("vagrant", []string{"ssh", "-c", command, target}...)
		}

		cmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(cmd, fmt.Sprintf("deploy binary for %v/%v on %v with vagrant", system, target, runtime.GOOS))

		haltCmd := exec.Command("vagrant", "halt", target)
		haltCmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(haltCmd, fmt.Sprintf("vagrant halt for %v/%v on %v", system, target, runtime.GOOS))
	}
}

func BuildEnv(target, name, depPath string) (map[string]string, []string, []string, string) {

	var (
		tags    []string
		ldFlags []string
		out     string
		env     map[string]string
	)

	androidAPI := utils.ANDROID_NDK_PLATFORM()
	androidEABI := strings.Replace(androidAPI, "-", "eabi", -1)

	switch target {
	case "android":
		tags = []string{target}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "libgo_base.so")
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "android",
			"GOARCH": "arm",
			"GOARM":  "7",

			"CGO_ENABLED": "1",
			"CC":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "llvm", "prebuilt", runtime.GOOS+"-x86_64", "bin", "clang"),
			"CXX":         filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "llvm", "prebuilt", runtime.GOOS+"-x86_64", "bin", "clang++"),
		}
		if utils.QT_VERSION_NUM() >= 5140 {
			env["CGO_CPPFLAGS"] = "-target armv7a-linux-" + androidEABI
			env["CGO_LDFLAGS"] = "-target armv7a-linux-" + androidEABI
		} else {
			env["CGO_CPPFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target armv7-none-linux-androideabi -gcc-toolchain %v --sysroot=%v -isystem %v",
				androidAPI,
				filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot", "usr", "include", "arm-linux-androideabi"))
			env["CGO_LDFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target armv7-none-linux-androideabi -gcc-toolchain %v --sysroot=%v",
				androidAPI,
				filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", androidAPI, "arch-arm"))
		}
		if runtime.GOOS == "windows" {
			env["TMP"] = os.Getenv("TMP")
			env["TEMP"] = os.Getenv("TEMP")
		}

		if utils.GOARCH() == "arm64" {
			env["GOARCH"] = "arm64"

			if utils.QT_VERSION_NUM() >= 5140 {
				env["CGO_CPPFLAGS"] = "-target aarch64-linux-" + androidEABI
				env["CGO_LDFLAGS"] = "-target aarch64-linux-" + androidEABI
			} else {
				env["CGO_CPPFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target aarch64-none-linux-android -gcc-toolchain %v --sysroot=%v -isystem %v",
					androidAPI,
					filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "aarch64-linux-android-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
					filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot"),
					filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot", "usr", "include", "aarch64-linux-android"))
				env["CGO_LDFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target aarch64-none-linux-android -gcc-toolchain %v --sysroot=%v",
					androidAPI,
					filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "aarch64-linux-android-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
					filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", androidAPI, "arch-arm64"))
			}
		}

	case "android-emulator":
		tags = []string{strings.Replace(target, "-", "_", -1)}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "libgo_base.so")
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "android",
			"GOARCH": "386",

			"CGO_ENABLED": "1",
			"CC":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "llvm", "prebuilt", runtime.GOOS+"-x86_64", "bin", "clang"),
			"CXX":         filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "llvm", "prebuilt", runtime.GOOS+"-x86_64", "bin", "clang++"),
		}
		if utils.QT_VERSION_NUM() >= 5140 {
			env["CGO_CPPFLAGS"] = "-target i686-linux-" + androidEABI
			env["CGO_LDFLAGS"] = "-target i686-linux-" + androidEABI
		} else {
			env["CGO_CPPFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target i686-none-linux-android -mstackrealign -gcc-toolchain %v --sysroot=%v -isystem %v",
				androidAPI,
				filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "x86-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot", "usr", "include", "i686-linux-android"))
			env["CGO_LDFLAGS"] = fmt.Sprintf("-Wno-unused-command-line-argument -D__ANDROID_API__=%v -target i686-none-linux-android -mstackrealign -gcc-toolchain %v --sysroot=%v",
				androidAPI,
				filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "x86-4.9", "prebuilt", runtime.GOOS+"-x86_64"),
				filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", androidAPI, "arch-x86"))
		}
		if runtime.GOOS == "windows" {
			env["TMP"] = os.Getenv("TMP")
			env["TEMP"] = os.Getenv("TEMP")
		}

		//TODO: x86_64

	case "ios", "ios-simulator":
		tags = []string{"ios"}
		ldFlags = []string{"-w"}
		out = filepath.Join(depPath, "libgo.a")

		GOARCH := "amd64"
		CLANGDIR := "iPhoneSimulator"
		CLANGPLAT := utils.IPHONESIMULATOR_SDK_DIR()
		CLANGFLAG := "ios-simulator"
		CLANGARCH := "x86_64"
		if target == "ios" {
			GOARCH = "arm64"
			CLANGDIR = "iPhoneOS"
			CLANGPLAT = utils.IPHONEOS_SDK_DIR()
			CLANGFLAG = "iphoneos"
			CLANGARCH = "arm64"
		}

		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   runtime.GOOS,
			"GOARCH": GOARCH,

			"CGO_ENABLED": "1",

			//TODO: move flags into template_cgo_qmake ?
			"CGO_CPPFLAGS": fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=11.0 -arch %v", utils.XCODE_DIR(), CLANGDIR, CLANGPLAT, CLANGFLAG, CLANGARCH),
			"CGO_LDFLAGS":  fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=11.0 -arch %v", utils.XCODE_DIR(), CLANGDIR, CLANGPLAT, CLANGFLAG, CLANGARCH),
		}

	case "darwin":
		ldFlags = []string{"-w"}
		out = filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", name, name))
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "darwin",
			"GOARCH": "amd64",

			"CGO_ENABLED": "1",
		}

	case "windows":
		ldFlags = []string{"-s", "-w"}
		if !utils.QT_DEBUG_CONSOLE() {
			ldFlags = append(ldFlags, "-H=windowsgui")
		}
		if runtime.GOOS != target {
			tags = []string{"windows"}
		}
		out = filepath.Join(depPath, name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"TMP":  os.Getenv("TMP"),
			"TEMP": os.Getenv("TEMP"),

			"GOOS":   "windows",
			"GOARCH": "386",

			"CGO_ENABLED": "1",
		}

		if runtime.GOOS == target {
			if utils.QT_VERSION_NUM() >= 5120 {
				env["GOARCH"] = "amd64"

				if utils.QT_VERSION_NUM() >= 5122 {
					env["GOARCH"] = utils.GOARCH()
				}

				//TODO: support 32 and 64 bit; fix InitEnv as well (for 32 bit support)
				if strings.Contains(utils.QT_INSTALL_PREFIX(target), "env_windows_amd64") && utils.QT_VERSION_NUM() < 5130 { //TODO: fix env_windows_amd64_512 instead
					env["CGO_LDFLAGS"] = filepath.Join("C:\\", "Users", "Public", "env_windows_amd64", "Tools", "mingw730_64", "x86_64-w64-mingw32", "lib", "libmsvcrt.a")
				}
			}

			if utils.QT_MSYS2() {
				env["GOARCH"] = utils.QT_MSYS2_ARCH()
				env["PATH"] = filepath.Join(utils.QT_MSYS2_DIR(), "bin") + ";" + env["PATH"]
			} else {
				env["PATH"] = utils.MINGWTOOLSDIR() + ";" + env["PATH"]
			}

			if utils.QT_MSVC() {
				env["CC_FOR_CGO"] = "gcc"
				env["CC"] = "cl"
			}
		} else {
			delete(env, "TMP")
			delete(env, "TEMP")

			env["GOARCH"] = utils.QT_MXE_ARCH()
			env["CC"] = utils.QT_MXE_BIN("gcc")
			env["CXX"] = utils.QT_MXE_BIN("g++")
		}

	case "linux", "ubports":
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "linux",
			"GOARCH": utils.GOARCH(),

			"CGO_ENABLED": "1",
		}

		if utils.QT_VERSION_NUM() <= 5051 {
			env["CGO_CXXFLAGS"] = "-std=c++11"
		}

		if strings.HasPrefix(env["GOARCH"], "arm") {
			if env["GOARCH"] == "arm" {
				if arm, ok := utils.GOARM(); ok {
					env["GOARM"] = arm
				}
			}
			if cc, ok := os.LookupEnv("CC"); ok {
				env["CC"] = cc
			}
			if cxx, ok := os.LookupEnv("CXX"); ok {
				env["CXX"] = cxx
			}
		}

		if utils.QT_UBPORTS() || target == "ubports" {
			tags = []string{"ubports", utils.QT_UBPORTS_VERSION()}

			if utils.QT_UBPORTS_ARCH() == "arm" {
				env["GOARCH"] = "arm"

				if env["GOARM"] == "" {
					env["GOARM"] = "7"
				}

				if env["CC"] == "" || env["CXX"] == "" {
					env["CC"] = "arm-linux-gnueabihf-gcc"
					env["CXX"] = "arm-linux-gnueabihf-g++"
				}
			}
		}

		if _, ok := os.LookupEnv("CGO_LDFLAGS"); target == "linux" && !((utils.QT_STATIC() || utils.QT_PKG_CONFIG() || ok) && !ImportsFlutter()) {
			env["CGO_LDFLAGS"] = "-Wl,-rpath,$ORIGIN/lib -Wl,--disable-new-dtags"
		}

	case "freebsd":
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "freebsd",
			"GOARCH": utils.GOARCH(),

			"CGO_ENABLED": "1",
		}

		if strings.HasPrefix(env["GOARCH"], "arm") {
			if env["GOARCH"] == "arm" {
				if arm, ok := utils.GOARM(); ok {
					env["GOARM"] = arm
				}
			}
			if cc, ok := os.LookupEnv("CC"); ok {
				env["CC"] = cc
			}
			if cxx, ok := os.LookupEnv("CXX"); ok {
				env["CXX"] = cxx
			}
		}

	case "rpi1", "rpi2", "rpi3":
		tags = []string{target}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, name)

		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "linux",
			"GOARCH": "arm",
			"GOARM":  "7",

			"CGO_ENABLED": "1",
			"CC":          fmt.Sprintf("%v/arm-bcm2708/%v/bin/arm-linux-gnueabihf-gcc", utils.RPI_TOOLS_DIR(), utils.RPI_COMPILER()),
			"CXX":         fmt.Sprintf("%v/arm-bcm2708/%v/bin/arm-linux-gnueabihf-g++", utils.RPI_TOOLS_DIR(), utils.RPI_COMPILER()),
		}

		if target == "rpi1" {
			env["GOARM"] = "6"
		}

	case "sailfish":
		tags = []string{target}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "harbour-"+name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "linux",
			"GOARCH": "arm",
			"GOARM":  "7",

			"CGO_ENABLED": "1",
			"CC":          "/srv/mer/toolings/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "/opt/cross/bin/armv7hl-meego-linux-gnueabi-gcc",
			"CXX":         "/srv/mer/toolings/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "/opt/cross/bin/armv7hl-meego-linux-gnueabi-g++",

			//TODO: use plain -I and -L instead ?
			"CPATH":        "/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-armv7hl/usr/include",
			"LIBRARY_PATH": "/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-armv7hl/usr/lib:/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-armv7hl/lib:/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-armv7hl/usr/lib/pulseaudio",

			//TODO: move flags into template_cgo_qmake ?
			"CGO_LDFLAGS": "--sysroot=/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-armv7hl/",
		}

	case "sailfish-emulator":
		tags = []string{strings.Replace(target, "-", "_", -1)}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "harbour-"+name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "linux",
			"GOARCH": "386",

			"CGO_ENABLED": "1",
			"CC":          "/srv/mer/toolings/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "/opt/cross/bin/i486-meego-linux-gnu-gcc",
			"CXX":         "/srv/mer/toolings/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "/opt/cross/bin/i486-meego-linux-gnu-g++",

			//TODO: use plain -I and -L instead ?
			"CPATH":        "/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-i486/usr/include",
			"LIBRARY_PATH": "/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-i486/usr/lib:/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-i486/lib:/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-i486/usr/lib/pulseaudio",

			//TODO: move flags into template_cgo_qmake ?
			"CGO_LDFLAGS": "--sysroot=/srv/mer/targets/SailfishOS-" + utils.QT_SAILFISH_VERSION() + "-i486/",
		}

	case "js", "wasm":
		tags = []string{target}
		if target == "wasm" {
			ldFlags = []string{"-s", "-w"}
		}
		out = filepath.Join(depPath, "go")
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,

			"CGO_ENABLED": "1",
		}

		if target == "wasm" {
			env["GOOS"] = "js"
			env["GOARCH"] = "wasm"
		}

		env["EM_CONFIG"] = filepath.Join(os.Getenv("HOME"), ".emscripten")
		for _, l := range strings.Split(utils.Load(env["EM_CONFIG"]), "\n") {
			l = strings.Replace(l, "'", "", -1)
			l = strings.Replace(l, " ", "", -1)
			switch {
			case strings.HasPrefix(l, "LLVM_ROOT"):
				env["LLVM_ROOT"] = strings.Split(l, "=")[1]
			case strings.HasPrefix(l, "BINARYEN_ROOT"):
				env["BINARYEN_ROOT"] = strings.Split(l, "=")[1]
			case strings.HasPrefix(l, "EMSCRIPTEN_NATIVE_OPTIMIZER"):
				env["EMSCRIPTEN_NATIVE_OPTIMIZER"] = strings.Split(l, "=")[1]
			case strings.HasPrefix(l, "NODE_JS"):
				env["NODE_JS"] = strings.TrimSuffix(strings.Split(l, "=")[1], "/node")
			case strings.HasPrefix(l, "EMSCRIPTEN_ROOT"):
				env["EMSCRIPTEN"] = strings.Split(l, "=")[1]
				env["EMSDK"] = strings.Split(env["EMSCRIPTEN"], "/emscripten/1.")[0]
			}
		}
		if _, ok := env["EMSCRIPTEN"]; !ok {
			env["EMSCRIPTEN"] = filepath.Join(env["BINARYEN_ROOT"], "emscripten")
		}
		env["PATH"] = env["EMSDK"] + ":" + env["LLVM_ROOT"] + ":" + env["NODE_JS"] + ":" + env["EMSCRIPTEN"] + ":" + env["PATH"]
	}

	if runtime.GOOS != target || utils.GOVERSION_NUM() == 110 {
		env["CGO_CFLAGS_ALLOW"] = utils.CGO_CFLAGS_ALLOW()
		env["CGO_CXXFLAGS_ALLOW"] = utils.CGO_CXXFLAGS_ALLOW()
		env["CGO_LDFLAGS_ALLOW"] = utils.CGO_LDFLAGS_ALLOW()
	}

	if utils.QT_MSVC() {
		env["CGO_MSCFLAGS_ALLOW"] = utils.CGO_MSCFLAGS_ALLOW()
		env["CGO_MSCXXFLAGS_ALLOW"] = utils.CGO_MSCXXFLAGS_ALLOW()
		env["CGO_MSLDFLAGS_ALLOW"] = utils.CGO_MSLDFLAGS_ALLOW()
	}

	if flags := utils.GOFLAGS(); len(flags) != 0 {
		env["GOFLAGS"] = flags
	}

	for _, e := range os.Environ() {
		es := strings.Split(e, "=")
		if _, ok := env[es[0]]; !ok {
			env[es[0]] = strings.Join(es[1:], "=")
		}
	}

	return env, tags, ldFlags, out
}
