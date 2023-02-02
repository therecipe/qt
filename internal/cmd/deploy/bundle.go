package deploy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/bluszcz/cutego/internal/binding/parser"
	"github.com/bluszcz/cutego/internal/binding/templater"

	"github.com/bluszcz/cutego/internal/cmd"
	"github.com/bluszcz/cutego/internal/cmd/moc"
	"github.com/bluszcz/cutego/internal/cmd/rcc"
	"github.com/bluszcz/cutego/internal/utils"
)

func bundle(mode, target, path, name, depPath string, tagsCustom string, fast bool) {
	copy := func(src, dst string) {
		copy := "cp"

		if runtime.GOOS == "windows" && utils.MSYSTEM() == "" {
			copy = "xcopy"

			//TODO: -->
			src = strings.TrimSuffix(src, "/")
			src = strings.TrimSuffix(src, "\\")

			dst = strings.TrimSuffix(dst, "/")
			dst = strings.TrimSuffix(dst, "\\")
			//<--
		}

		var args []string
		if utils.ExistsDir(src) {
			if runtime.GOOS != "windows" || utils.MSYSTEM() != "" {
				args = append(args, "-R")
			}
		}

		var suffix string
		if !utils.ExistsDir(dst) {
			if runtime.GOOS == "windows" && utils.MSYSTEM() == "" {
				suffix = "*"
			}
		}

		utils.RunCmd(exec.Command(copy, append(args, src, dst+suffix)...), fmt.Sprintf("copy %v to %v for %v on %v", filepath.Base(src), filepath.Base(dst), target, runtime.GOOS))
	}

	switch target {
	case "darwin":

		//copy default assets
		utils.Save(filepath.Join(depPath, name+".app", "Contents", "Info.plist"), darwin_plist(name))
		utils.Save(filepath.Join(depPath, name+".app", "Contents", "PkgInfo"), darwin_pkginfo())
		utils.MkdirAll(filepath.Join(depPath, name+".app", "Contents", "Resources"))
		utils.Save(filepath.Join(depPath, name+".app", "Contents", "Resources", "empty.lproj"), "")

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", filepath.Join(depPath, name+".app"))

		if cmd.ImportsFlutter() {
			utils.MkdirAll(filepath.Join(depPath, name+".app", "Contents", "Frameworks"))

			copy(filepath.Join(path, "build", "flutter_assets"), filepath.Join(depPath, name+".app", "Contents", "MacOS", "flutter_assets"))
			copy(filepath.Join(path, "FlutterEmbedder.framework"), filepath.Join(depPath, name+".app", "Contents", "Frameworks", "FlutterEmbedder.framework"))

			if utils.QT_STATIC() {
				utils.RunCmd(exec.Command("install_name_tool", "-add_rpath", "@executable_path/../Frameworks", filepath.Join(depPath, name+".app", "Contents", "MacOS", name)), "add an rpath to the binary")
			}
		}

		if utils.QT_STATIC() {
			break
		}

		if utils.QT_NIX() {
			/*
				TODO:
				no self containing deployments possible because
				macdeployqt can't find qmlimportscanner and
				icu libs are still partially linked against
				the store libs after macdeployqt
			*/

			//workaround to make bundled applications start when they are double clicked from within the finder
			os.Rename(filepath.Join(depPath, name+".app", "Contents", "MacOS", name), filepath.Join(depPath, name+".app", "Contents", "MacOS", name+"_bin"))
			utils.SaveExec(filepath.Join(depPath, name+".app", "Contents", "MacOS", name), darwin_nix_script(name))
			break
		}

		dep := exec.Command(utils.ToolPath("macdeployqt", target))
		dep.Args = append(dep.Args, filepath.Join(depPath, name+".app"), "-qmldir="+path)
		dep.Dir = filepath.Dir(dep.Path)
		utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))

		//break the rpath
		pPath := "/broken/"
		fn := filepath.Join(depPath, name+".app", "Contents", "MacOS", name)
		data, err := ioutil.ReadFile(fn)
		if err != nil {
			utils.Log.WithError(err).Warn("couldn't find", fn)
			break
		}

		prefPath := utils.QT_INSTALL_PREFIX(target)

		start := bytes.Index(data, []byte(prefPath))
		if start == -1 {
			break
		}

		end := bytes.IndexByte(data[start:], byte(0))
		if end == -1 {
			break
		}

		rep := []byte(pPath)
		if lendiff := end - len(rep); lendiff < 0 {
			end -= lendiff
		} else {
			rep = append(rep, bytes.Repeat([]byte{0}, lendiff)...)
		}
		data = bytes.Replace(data, data[start:start+end], rep, -1)

		if err := ioutil.WriteFile(fn, data, 0755); err != nil {
			utils.Log.WithError(err).Warn("couldn't patch", fn)
		} else {
			utils.Log.Debugln("patched", fn)
		}

	case "linux", "rpi1", "rpi2", "rpi3", "freebsd":
		var rpiToolPath string
		if strings.HasPrefix(target, "rpi") {
			env, _, _, _ := cmd.BuildEnv(target, "", "")
			rpiToolPath = strings.TrimSuffix(env["CC"], "gcc")
		}

		defer func() {
			filepath.Walk(depPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if strings.HasPrefix(filepath.Base(path), "lib") { //TODO: only strip binaries
					if strings.HasPrefix(target, "rpi") {
						utils.RunCmdOptional(exec.Command(rpiToolPath+"strip", "-s", path), "strip binaries on linux")
					} else {
						utils.RunCmdOptional(exec.Command("strip", "-s", path), "strip binaries on linux")
					}
				}
				return nil
			})
		}()

		//copy default assets
		if !(target == "linux" || target == "freebsd") || name == "lib" {
			utils.SaveExec(filepath.Join(depPath, fmt.Sprintf("%v.sh", name)), linux_sh(target, name))
		}

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", depPath)

		if cmd.ImportsFlutter() {
			libDir := "lib"
			if name == libDir {
				libDir = "libs"
			}
			utils.MkdirAll(filepath.Join(depPath, libDir))

			copy(filepath.Join(path, "build", "flutter_assets"), filepath.Join(depPath, "flutter_assets"))
			copy(filepath.Join(path, "libflutter_engine.so"), filepath.Join(depPath, "lib", "libflutter_engine.so"))
			copy(filepath.Join(path, "icudtl.dat"), filepath.Join(depPath, "icudtl.dat"))
		}

		if utils.QT_STATIC() || utils.QT_PKG_CONFIG() {
			break
		}

		//TODO: -->
		{
			libDir := "lib"
			if name == libDir {
				libDir = "libs"
			}
			utils.MkdirAll(filepath.Join(depPath, libDir))

			var (
				libraryPath   = filepath.Join(utils.QT_INSTALL_PREFIX(target), "lib")
				lddOutput     string
				usesWebEngine bool
				usesQml       bool
			)

			if strings.HasPrefix(target, "rpi") {
				lddOutput = utils.RunCmd(exec.Command(rpiToolPath+"ldd", "--root=/", filepath.Join(depPath, name)), fmt.Sprintf("ldd binary for %v on %v", target, runtime.GOOS))
			} else {
				lddOutput = utils.RunCmd(exec.Command("ldd", filepath.Join(depPath, name)), fmt.Sprintf("ldd binary for %v on %v", target, runtime.GOOS))
			}

			for _, dep := range strings.Split(lddOutput, "\n") {
				if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
					var libName string

					if strings.HasPrefix(target, "rpi") {
						libName = strings.TrimSpace(strings.Replace(strings.Split(dep, "=>")[0], "not found", "", -1))
					} else {
						if libraryPath == "" {
							libraryPath, libName = filepath.Split(strings.Split(dep, " ")[2])
						} else {
							_, libName = filepath.Split(strings.Split(dep, " ")[2])
						}
					}

					if utils.ExistsFile(filepath.Join(libraryPath, libName)) {
						utils.RunCmd(exec.Command("cp", "-L", strings.TrimSuffix(filepath.Join(libraryPath, libName), ".5"), filepath.Join(depPath, libDir, libName)), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))

						if strings.HasPrefix(libName, "libQt5Multimedia") {
							libName = "libQt5MultimediaWidgets.so.5"
							utils.RunCmdOptional(exec.Command("cp", "-L", strings.TrimSuffix(filepath.Join(libraryPath, libName), ".5"), filepath.Join(depPath, libDir, libName)), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
						}
						if strings.HasPrefix(libName, "libQt5WebEngine") {
							libName = "libQt5WebEngineWidgets.so.5"
							utils.RunCmdOptional(exec.Command("cp", "-L", strings.TrimSuffix(filepath.Join(libraryPath, libName), ".5"), filepath.Join(depPath, libDir, libName)), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
						}
						if strings.HasPrefix(libName, "libQt5Quick") {
							libName = "libQt5QuickWidgets.so.5"
							utils.RunCmdOptional(exec.Command("cp", "-L", strings.TrimSuffix(filepath.Join(libraryPath, libName), ".5"), filepath.Join(depPath, libDir, libName)), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
						}
					}

					if strings.Contains(dep, "WebEngine") || strings.Contains(dep, "WebView") {
						usesWebEngine = true
					}
					if strings.Contains(dep, "Quick") || strings.Contains(dep, "Qml") {
						usesQml = true
					}
				}
			}

			libs := []string{"DBus", "XcbQpa", "Quick", "Widgets", "EglDeviceIntegration", "EglFsKmsSupport", "OpenGL", "WaylandClient", "WaylandCompositor", "QuickControls2", "QuickTemplates2", "QuickWidgets", "QuickParticles", "CLucene", "Concurrent", "Svg", "MultimediaGstTools"}
			if usesQml {
				libs = append(libs, []string{"3DCore", "3DExtras", "3DInput", "3DLogic", "3DQuick", "3DQuickExtras", "3DQuickInput", "3DQuickRender", "3DRender", "Gamepad", "Multimedia", "MultimediaQuick", "QmlModels", "QmlWorkerScript"}...)
			}
			if usesWebEngine {
				libs = append(libs, []string{"WebEngine", "WebEngineCore", "WebChannel", "Positioning"}...)
			}
			for _, libName := range libs {
				if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so", libName))) {
					utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so", libName)), filepath.Join(depPath, libDir, fmt.Sprintf("libQt5%v.so.5", libName))), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
				}
			}
			if utils.ExistsFile(filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0")) {
				utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0"), filepath.Join(depPath, libDir, "libqgsttools_p.so.1")), fmt.Sprintf("copy libqgsttools_p.so.1 for %v on %v", target, runtime.GOOS))
			}

			libraryPath = filepath.Dir(libraryPath)
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))
			//TODO: use rsync with exclude instead ...
			filepath.Walk(depPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if filepath.Ext(info.Name()) == ".debug" || filepath.Ext(info.Name()) == ".qmlc" || filepath.Ext(info.Name()) == ".jsc" {
					utils.RemoveAll(path)
				}
				return nil
			})

			if usesWebEngine {
				utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, "libexec", "QtWebEngineProcess"), depPath), fmt.Sprintf("copy QtWebEngineProcess for %v on %v", target, runtime.GOOS))
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "resources/"), depPath), fmt.Sprintf("copy resource dir for %v on %v", target, runtime.GOOS))
				utils.MkdirAll(filepath.Join(depPath, "translations/"))
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "translations/qtwebengine_locales/"), filepath.Join(depPath, "translations/")), fmt.Sprintf("copy qtwebengine_locales dir for %v on %v", target, runtime.GOOS))

				//patch QtWebEngineProcess rpath
				pPath := "lib"
				fn := filepath.Join(depPath, "QtWebEngineProcess")
				data, err := ioutil.ReadFile(fn)
				if err != nil {
					utils.Log.WithError(err).Warn("couldn't find", fn)
					break
				}

				prefPath := "$ORIGIN/"

				start := bytes.Index(data, []byte(prefPath))
				if start == -1 {
					break
				}

				end := bytes.IndexByte(data[start:], byte(0))
				if end == -1 {
					break
				}

				rep := append([]byte(prefPath), []byte(pPath)...)
				if lendiff := end - len(rep); lendiff < 0 {
					end -= lendiff
				} else {
					rep = append(rep, bytes.Repeat([]byte{0}, lendiff)...)
				}
				data = bytes.Replace(data, data[start:start+end], rep, -1)

				if err := ioutil.WriteFile(fn, data, 0755); err != nil {
					utils.Log.WithError(err).Warn("couldn't patch", fn)
				} else {
					utils.Log.Debugln("patched", fn)
				}
			}

			//patch QtCore path
			pPath := "."
			fn := filepath.Join(depPath, "/lib/", "libQt5Core.so.5")
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				utils.Log.WithError(err).Warn("couldn't find", fn)
				break
			}

			prefPath := "qt_prfxpath="

			start := bytes.Index(data, []byte(prefPath))
			if start == -1 {
				break
			}

			end := bytes.IndexByte(data[start:], byte(0))
			if end == -1 {
				break
			}

			rep := append([]byte(prefPath), []byte(pPath)...)
			if lendiff := end - len(rep); lendiff < 0 {
				end -= lendiff
			} else {
				rep = append(rep, bytes.Repeat([]byte{0}, lendiff)...)
			}
			data = bytes.Replace(data, data[start:start+end], rep, -1)

			if err := ioutil.WriteFile(fn, data, 0644); err != nil {
				utils.Log.WithError(err).Warn("couldn't patch", fn)
			} else {
				utils.Log.Debugln("patched", fn)
			}
		}
		//<--

	case "windows":

		//copy default assets
		//TODO: windres icon

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+string(filepath.Separator)+".", depPath)

		if cmd.ImportsFlutter() {
			if runtime.GOOS == "windows" {
				utils.MkdirAll(filepath.Join(depPath, "flutter_assets"))
				copy(filepath.Join(path, "build", "flutter_assets"+string(filepath.Separator)+"."), filepath.Join(depPath, "flutter_assets"))
			} else {
				copy(filepath.Join(path, "build", "flutter_assets"), filepath.Join(depPath, "flutter_assets"))
			}

			//TODO: new generic copy method; xcopy is somewhat broken under wine
			utils.Save(filepath.Join(depPath, "flutter_engine.dll"), utils.Load(filepath.Join(path, "flutter_engine.dll")))
			utils.Save(filepath.Join(depPath, "icudtl.dat"), utils.Load(filepath.Join(path, "icudtl.dat")))
		}

		//TODO: -->
		switch {
		case runtime.GOOS != target:
			if utils.QT_MXE_STATIC() {
				break
			}

			var libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "bin")
			for _, d := range []string{"libbz2", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libgraphite2", "libeay32", "ssleay32", "libcrypto-1_1-x64", "libpcre2-16-0", "libssl-1_1-x64", "libzstd"} {
				if utils.QT_MXE_ARCH() == "386" {
					d = strings.TrimSuffix(d, "-x64")
				}
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}
			for _, icu := range []string{"icudt", "icuin", "icuuc"} {
				for i := 55; i < 70; i++ {
					lib := filepath.Join(libraryPath, fmt.Sprintf("%v%v.dll", icu, i))
					if utils.ExistsFile(lib) {
						utils.RunCmd(exec.Command("cp", lib, depPath), fmt.Sprintf("copy %v for %v on %v", filepath.Base(lib), target, runtime.GOOS))
						break
					}
				}
			}
			for _, d := range []string{"libjasper", "libjpeg-9", "libmng-2", "libtiff-5", "libwebp-7", "liblcms2-2", "liblzma-5", "libwebpdemux-2"} {
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}

			var gccDep = "libgcc_s_sjlj-1"
			if utils.QT_MXE_ARCH() == "amd64" {
				gccDep = "libgcc_s_seh-1"
			}
			utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5")
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5", "bin")
			var output = utils.RunCmd(exec.Command(utils.QT_MXE_BIN("objdump"), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
			for lib, deps := range parser.LibDeps {
				if strings.Contains(output, lib) && lib != parser.MOC {
					for _, lib := range append(deps, lib) {
						if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib))) && !utils.ExistsFile(filepath.Join(depPath, fmt.Sprintf("Qt5%v.dll", lib))) {
							utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
						}
					}
				}
			}
			for _, d := range []string{"Qt5OpenGL", "Qt5Quick", "Qt5QuickControls2", "Qt5QuickTemplates2"} {
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}

		case utils.QT_MSYS2():
			if utils.QT_MSYS2_STATIC() {
				break
			}

			copyCmd := "xcopy"
			if utils.MSYSTEM() != "" {
				copyCmd = "cp"
			}

			deploy := exec.Command(utils.ToolPath("windeployqt", target))
			deploy.Args = append(deploy.Args, "--verbose=2", "--force", fmt.Sprintf("--qmldir=%v", path))
			if utils.QT_DOCKER() {
				deploy.Args = append(deploy.Args, "--no-translations") //TODO:
			}
			deploy.Args = append(deploy.Args, filepath.Join(depPath, name+".exe"))
			env, _, _, _ := cmd.BuildEnv(target, "", "")
			for key, value := range env {
				deploy.Env = append(deploy.Env, fmt.Sprintf("%v=%v", key, value))
			}

			utils.RunCmd(deploy, fmt.Sprintf("depoy %v on %v", target, runtime.GOOS))

			var libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
			for _, d := range []string{"libbz2-1", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libgraphite2", "libeay32", "ssleay32", "libcrypto-1_1-x64", "libpcre2-16-0", "libssl-1_1-x64", "libdouble-conversion", "libzstd"} {
				if utils.QT_MSYS2_ARCH() == "386" {
					d = strings.TrimSuffix(d, "-x64")
				}

				if d == "libeay32" || d == "ssleay32" {
					utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(utils.QT_MSYS2_DIR(), "..", "opt", "bin", fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
				} else {
					utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
				}
			}
			for _, icu := range []string{"icudt", "icuin", "icuuc"} {
				for i := 55; i < 70; i++ {
					lib := filepath.Join(libraryPath, fmt.Sprintf("lib%v%v.dll", icu, i))
					if utils.ExistsFile(lib) {
						utils.RunCmd(exec.Command(copyCmd, lib, depPath), fmt.Sprintf("copy %v for %v on %v", filepath.Base(lib), target, runtime.GOOS))
						break
					}
				}
			}

			var gccDep = "libgcc_s_dw2-1"
			if utils.QT_MSYS2_ARCH() == "amd64" {
				gccDep = "libgcc_s_seh-1"
			}

			utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "share", "qt5")
			if utils.MSYSTEM() != "" {
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))
			} else {
				utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
				utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))
			}

			libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
			var output = utils.RunCmd(exec.Command(utils.ToolPath("objdump", target), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
			for lib, deps := range parser.LibDeps {
				if strings.Contains(output, lib) && lib != parser.MOC {
					for _, lib := range append(deps, lib) {
						if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib))) && !utils.ExistsFile(filepath.Join(depPath, fmt.Sprintf("Qt5%v.dll", lib))) {
							if utils.MSYSTEM() != "" {
								utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
							} else {
								utils.RunCmd(exec.Command("xcopy", "/Y", filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
							}
						}
					}
				}
			}

			deps := []string{"Qt5OpenGL", "Qt5Quick", "Qt5QuickControls2", "Qt5QuickTemplates2"}
			if utils.QT_WEBKIT() {
				deps = append(deps, []string{"libjpeg-8", "libsqlite3-0", "libwebp-7", "libxml2-2", "liblzma-5", "libxslt-1"}...)
			}
			for _, lib := range deps {
				if utils.MSYSTEM() != "" {
					utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
				} else {
					utils.RunCmd(exec.Command("xcopy", "/Y", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
				}
			}

			filepath.Walk(depPath, func(path string, info os.FileInfo, err error) error {
				if strings.HasSuffix(info.Name(), "d.dll") && !strings.HasSuffix(info.Name(), "board.dll") {
					utils.RemoveAll(path)
				}
				return nil
			})

		default:

			if utils.QT_WEBKIT() {
				libraryPath := filepath.Dir(utils.ToolPath("qmake", target))
				output := utils.RunCmd(exec.Command(filepath.Join("objdump"), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
				for lib, deps := range parser.LibDeps {
					if strings.Contains(output, lib) && lib == "WebKit" {
						for _, lib := range append(deps, lib) {
							for _, pref := range []string{"lib", ""} {
								libName := filepath.Join(libraryPath, fmt.Sprintf("%vQt5%v.dll", pref, lib))
								if utils.ExistsFile(libName) {
									copy(libName, depPath)
								}
							}
						}
					}
				}

				for _, lib := range []string{"libxml2-2", "libxslt-1", "Qt5MultimediaWidgets", "Qt5OpenGL", "Qt5PrintSupport"} {
					copy(filepath.Join(libraryPath, lib+".dll"), depPath)
				}
				for _, icu := range []string{"icudt", "icuin", "icuuc"} {
					for i := 55; i < 70; i++ {
						lib := filepath.Join(libraryPath, fmt.Sprintf("%v%v.dll", icu, i))
						if utils.ExistsFile(lib) {
							copy(lib, depPath)
							break
						}
					}
				}
			}

			dep := exec.Command(utils.ToolPath("windeployqt", target))
			if utils.QT_MSVC() {
				env, _, _, _ := cmd.BuildEnv(target, "", "")
				for k, v := range env {
					if strings.Contains(v, "mingw") {
						v = strings.Replace(v, "mingw", "_break_", -1)
					}
					dep.Env = append(dep.Env, fmt.Sprintf("%v=%v", k, v))
				}
			}
			dep.Args = append(dep.Args, "--verbose=2", "--force", fmt.Sprintf("--qmldir=%v", path), filepath.Join(depPath, name+".exe"))
			utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))

			if utils.QT_MSVC() {
				cmd.PatchBinary(filepath.Join(depPath, name+".exe"))
			}
		}
		//<--

	case "android", "android-emulator":

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+string(filepath.Separator)+".", filepath.Join(depPath, "build"))

		//wrap exported go main inside c main
		env, _, _, _ := cmd.BuildEnv(target, "", "")
		compiler := env["CXX"]

		libPath := filepath.Join(depPath, "build", "libs", "armeabi-v7a")
		if target == "android" && utils.GOARCH() == "arm64" {
			libPath = filepath.Join(depPath, "build", "libs", "arm64-v8a")
		}
		if target == "android-emulator" {
			libPath = filepath.Join(depPath, "build", "libs", "x86")
		}
		utils.MkdirAll(libPath)

		finalLibGoName := "libgo.so"
		if utils.QT_VERSION_NUM() >= 5140 {
			finalLibGoName = "libgo_" + filepath.Base(libPath) + ".so"
		}

		wrapper := filepath.Join(depPath, "c_main_wrapper.cpp")
		utils.Save(wrapper, "#include \"libgo_base.h\"\nint main(int argc, char *argv[]) { go_main_wrapper(argc, argv); }")
		cmd := exec.Command(compiler, "c_main_wrapper.cpp", "-o", filepath.Join(depPath, "libgo.so"), "-I../..", "-L.", "-lgo_base", "-Wl,-soname,"+finalLibGoName, "-shared")
		if target == "android-emulator" {
			cmd = exec.Command(compiler, "c_main_wrapper.cpp", "-o", filepath.Join(depPath, "libgo.so"), "-I../..", "-L.", "-lgo_base", "-Wl,-soname,"+finalLibGoName, "-shared")
		}
		if utils.ANDROID_NDK_REQUIRE_NOSTDLIBPP_LDFLAG() {
			cmd.Args = append(cmd.Args, "-nostdlib++")
		}
		cmd.Args = append(cmd.Args, strings.Split(env["CGO_CPPFLAGS"], " ")...)
		cmd.Args = append(cmd.Args, "-I"+filepath.Join(utils.ANDROID_NDK_DIR(), "sysroot", "usr", "include"))
		cmd.Args = append(cmd.Args, strings.Split(env["CGO_LDFLAGS"], " ")...)
		cmd.Dir = depPath
		utils.RunCmd(cmd, fmt.Sprintf("compile wrapper for %v on %v", target, runtime.GOOS))
		utils.RemoveAll(wrapper)

		strip := exec.Command(filepath.Join(filepath.Dir(compiler), "llvm-strip"), "--strip-all", "libgo.so")
		strip.Dir = depPath
		utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", target, runtime.GOOS))

		if utils.QT_VAGRANT() {
			libPath = strings.Replace(libPath, "C:\\media\\sf_GOPATH", "C:\\media\\UNC\\vboxsrv\\media_sf_GOPATH", -1)
			utils.RemoveAll(libPath)
			utils.MkdirAll(libPath)
		}

		//trick androiddeployqt into checking dependencies from libgo_base.so 1/2 ->
		copy(filepath.Join(depPath, "libgo_base.so"), filepath.Join(libPath, finalLibGoName))
		copy(filepath.Join(depPath, "libgo_base.so"), filepath.Join(libPath, "libgo_base.so"))

		os.Rename(filepath.Join(depPath, "libgo.so"), filepath.Join(depPath, finalLibGoName))
		os.Chtimes(filepath.Join(depPath, finalLibGoName), time.Now(), time.Now())
		//<-

		utils.Save(filepath.Join(depPath, "android-libgo.so-deployment-settings.json"), android_config(target, path, depPath))

		dep := exec.Command(filepath.Join(utils.QT_INSTALL_PREFIX(target), "bin", "androiddeployqt"))
		dep.Dir = filepath.Join(utils.QT_INSTALL_PREFIX(target), "bin")
		dep.Env = append(dep.Env, "JAVA_HOME="+utils.JDK_DIR(), "ANDROID_NDK_HOME="+utils.ANDROID_NDK_DIR())
		dep.Args = append(dep.Args,
			"--input", filepath.Join(depPath, "android-libgo.so-deployment-settings.json"),
			"--output", filepath.Join(depPath, "build"),
			"--deployment", "bundled",
			"--android-platform", "android-28",
			"--jdk", utils.JDK_DIR(),
			"--gradle",
			"--verbose",
		)
		if !utils.QT_DEBUG_QML() {
			dep.Args = append(dep.Args, "--no-gdbserver")
		}

		if utils.ExistsFile(filepath.Join(path, target, name+".jks")) {
			dep.Args = append(dep.Args,
				"--sign", filepath.Join(path, target, name+".jks"), strings.TrimSpace(utils.Load(filepath.Join(path, target, "jks_alias"))),
				"--storepass", strings.TrimSpace(utils.Load(filepath.Join(path, target, "jks_pass"))))
		}

		if runtime.GOOS == "windows" {
			//TODO: -->
			utils.SaveExec(filepath.Join(depPath, "build.bat"), fmt.Sprintf("set JAVA_HOME=%v\r\n%v", utils.JDK_DIR(), strings.Join(dep.Args, " ")))
			utils.RunCmd(exec.Command(filepath.Join(depPath, "build.bat")), fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))
			utils.RemoveAll(filepath.Join(depPath, "build.bat"))
			//<--
		} else {
			utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))
		}

		if utils.QT_VAGRANT() {
			depPathUNC := strings.Replace(depPath, "C:\\media\\sf_GOPATH", "C:\\media\\UNC\\vboxsrv\\media_sf_GOPATH", -1)
			if utils.ExistsFile(filepath.Join(path, target, name+".jks")) {
				copy(filepath.Join(depPathUNC, "build", "build", "outputs", "apk", "release", "build-release-signed.apk"), depPath)
			} else {
				copy(filepath.Join(depPathUNC, "build", "build", "outputs", "apk", "debug", "build-debug.apk"), depPath)
			}
		} else {
			if utils.ExistsFile(filepath.Join(path, target, name+".jks")) {
				copy(filepath.Join(depPath, "build", "build", "outputs", "apk", "release", "build-release-signed.apk"), depPath)
			} else {
				copy(filepath.Join(depPath, "build", "build", "outputs", "apk", "debug", "build-debug.apk"), depPath)
			}
		}

	case "ios", "ios-simulator":

		//copy default assets
		buildPath := filepath.Join(depPath, "build")
		utils.MkdirAll(filepath.Join(buildPath, "project.xcodeproj"))
		utils.MkdirAll(filepath.Join(buildPath, "project.xcodeproj", "project.xcworkspace", "xcshareddata"))
		utils.MkdirAll(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset"))
		utils.Save(filepath.Join(buildPath, "Info.plist"), ios_plist(name))
		utils.Save(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset", "Contents.json"), ios_appicon())
		utils.Save(filepath.Join(buildPath, "LaunchScreen.xib"), ios_launchscreen(name))
		utils.Save(filepath.Join(buildPath, "project.xcodeproj", "project.pbxproj"), ios_xcodeproject())
		utils.Save(filepath.Join(buildPath, "project.xcodeproj", "project.xcworkspace", "xcshareddata", "WorkspaceSettings.xcsettings"), ios_xcsettings())
		copy(filepath.Join(utils.QT_INSTALL_PREFIX(target), "mkspecs", "macx-ios-clang", "Default-568h@2x.png"), buildPath)

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", buildPath)

		var t string
		switch target {
		case "ios":
			t = "arm64"
		case "ios-simulator":
			t = "x86_64"
		}

		utils.Save(filepath.Join(depPath, "c_main_wrapper_"+t+".cpp"), ios_c_main_wrapper())
		rcc.ResourceNames = make(map[string]string)
		moc.ResourceNames = make(map[string]string)
		cmdC := exec.Command("xcrun", "clang++", "c_main_wrapper_"+t+".cpp", target+"_plugin_import.cpp")
		cmdC.Dir = depPath
		newArgs := templater.GetiOSClang(target, t, depPath)
		if utils.ExistsFile(filepath.Join(depPath, target+"_qml_plugin_import.cpp")) {
			cmdC.Args = append(cmdC.Args, target+"_qml_plugin_import.cpp")
		}

		env, tags, _, _ := cmd.BuildEnv(target, "", "")
		if (!fast || utils.QT_STUB()) && !utils.QT_FAT() {
			tags = append(tags, "minimal")
		}
		if tagsCustom != "" {
			tags = append(tags, strings.Split(tagsCustom, " ")...)
		}
		cmdF := utils.GoList("{{if not .Standard}}|{{join .CgoCPPFLAGS \"|\"}}|{{join .CgoCFLAGS \"|\"}}|{{join .CgoCXXFLAGS \"|\"}}|{{join .CgoLDFLAGS \"|\"}}|{{end}}", "-deps", utils.BuildTags(tags))
		cmdF.Dir = path
		for k, v := range env {
			cmdF.Env = append(cmdF.Env, fmt.Sprintf("%v=%v", k, v))
		}
		var nextNeedsFramework bool
		for _, f := range strings.Split(strings.TrimSpace(utils.RunCmd(cmdF, "go list flags")), "|") {
			f = strings.TrimSpace(f)
			var found bool
			for _, n := range newArgs {
				if f == n {
					found = true
					break
				}
			}
			if f == "-framework" {
				nextNeedsFramework = true
			} else {
				if !found {
					if nextNeedsFramework {
						newArgs = append(newArgs, []string{"-framework", f}...)
					} else {
						newArgs = append(newArgs, f)
					}
				}
				nextNeedsFramework = false
			}
		}

		cmdC.Args = append(cmdC.Args, "-o", "build/main", "-u", "_qt_registerPlatformPlugin", "-Wl,-e,_qt_main_wrapper", "-I../..", "-L.", "-lgo")
		cmdC.Args = append(cmdC.Args, newArgs...)
		utils.RunCmd(cmdC, fmt.Sprintf("compile wrapper for %v (%v) on %v", target, t, runtime.GOOS))

		strip := exec.Command("strip", "main")
		strip.Dir = filepath.Join(depPath, "build")
		utils.RunCmd(strip, fmt.Sprintf("strip binary for %v (%v) on %v", target, t, runtime.GOOS))

		//run xcodebuild
		utils.RunCmd(exec.Command("xcrun", "xcodebuild", "clean", "build", "CODE_SIGN_IDENTITY=", "CODE_SIGNING_REQUIRED=NO", "CONFIGURATION_BUILD_DIR="+depPath, "-configuration", "Release", "-project", filepath.Join(depPath, "build", "project.xcodeproj")), fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))

	case "sailfish", "sailfish-emulator":

		//copy default assets
		utils.MkdirAll(filepath.Join(depPath, "rpm"))
		utils.Save(filepath.Join(depPath, "rpm", name+".spec"), sailfish_spec(name))
		utils.Save(filepath.Join(depPath, name+".desktop"), sailfish_desktop(name))
		if utils.QT_SAILFISH() {
			copy("/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/usr/share/themes/sailfish-default/meegotouch/z1.0/icons/icon-launcher-default.png", filepath.Join(depPath, fmt.Sprintf("harbour-%v.png", name)))
		} else {
			copy(filepath.Join(utils.SAILFISH_DIR(), "tutorials", "stocqt", "stocqt.png"), filepath.Join(depPath, fmt.Sprintf("harbour-%v.png", name)))
		}

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", depPath)

		if utils.QT_SAILFISH() {
			utils.RemoveAll(filepath.Join("/home", "user", target))
			copy(strings.Replace(depPath, "\\", "/", -1), filepath.Join("/home", "user", target))

			arch, template := "i486", "i486-meego-linux-gnu"
			if target == "sailfish" {
				arch, template = "armv7hl", "armv7hl-meego-linux"
			}

			pack := exec.Command("mb2", "-t", template, "build")
			pack.Dir = filepath.Join("/home", "user", target)
			utils.RunCmd(pack, fmt.Sprintf("deploy for %v (%v) on %v", target, arch, runtime.GOOS))

			copy(filepath.Join("/home", "user", target, "RPMS")+"/.", strings.Replace(depPath, "\\", "/", -1))
		} else {
			err := sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk", "&&", "rm", "-R", target)
			if err != nil {
				utils.Log.WithError(err).Warnf("failed to cleanup for %v on %v", target, runtime.GOOS)
			}

			err = sailfish_ssh("2222", "mersdk", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH/", -1)+"/deploy", "\\", "/", -1), "&&", "cp", "-R", target, "/home/mersdk")
			if err != nil {
				utils.Log.WithError(err).Panicf("failed to copy project for %v on %v", target, runtime.GOOS)
			}

			arch := "i486"
			if target == "sailfish" {
				arch = "armv7hl"
			}
			err = sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk/"+target, "&&", "mb2", "-t", "SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-"+arch, "build")
			if err != nil {
				utils.Log.WithError(err).Errorf("failed to deploy for %v (%v) on %v", target, arch, runtime.GOOS)
			}

			err = sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk/"+target+"/RPMS", "&&", "cp", "*", strings.Replace(strings.Replace(depPath, utils.MustGoPath(), "/media/sf_GOPATH/", -1), "\\", "/", -1))
			if err != nil {
				utils.Log.WithError(err).Panicf("failed to receive project for %v on %v", target, runtime.GOOS)
			}
		}

	case "ubports":

		//copy default assets
		copy("/usr/share/icons/suru/apps/sources/placeholder-app-icon.svg", filepath.Join(depPath, "logo.svg"))
		utils.Save(filepath.Join(depPath, "manifest.json"), ubports_manifest(name))
		utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.desktop", name)), ubports_desktop(name))
		utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.apparmor", name)), ubports_apparmor())

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", depPath)

		click := exec.Command("click", "build", "--no-validate", depPath)
		click.Dir = depPath
		utils.RunCmd(click, fmt.Sprintf("deploy for %v (%v) on %v", target, utils.QT_UBPORTS_ARCH(), runtime.GOOS))

	case "js", "wasm":

		//copy default assets
		copy(filepath.Join(utils.QT_INSTALL_PREFIX(target), "plugins", "platforms", "wasm_shell.html"), filepath.Join(depPath, "index.html"))
		copy(filepath.Join(utils.QT_INSTALL_PREFIX(target), "plugins", "platforms", "qtloader.js"), depPath)
		copy(filepath.Join(utils.QT_INSTALL_PREFIX(target), "plugins", "platforms", "qtlogo.svg"), depPath)
		if parser.UseWasm() {
			copy(filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.js"), filepath.Join(depPath, "go.js"))
		}

		//patch default assets
		index := utils.Load(filepath.Join(depPath, "index.html"))
		index = strings.Replace(index, "@APPNAME@", "main", -1)
		utils.Save(filepath.Join(depPath, "index.html"), strings.Replace(index, "  </body>", "    <script type=\"text/javascript\" src=\"go.js\"></script>\n  </body>", -1))

		if parser.UseWasm() {
			gojs := utils.Load(filepath.Join(depPath, "go.js"))
			if utils.GOVERSION_NUM() >= 113 {
				gojs = strings.Replace(gojs, "\"syscall/js.valueCall\": (sp) => {", "\"syscall/js.valueCall\": (sp) => {\n\t\t\t\t\t\t\tconst s = loadString(sp + 16); if (s.search(\"_Exec\") != -1) { const v = loadValue(sp + 8); Reflect.apply(Reflect.get(v, s), v, loadSliceOfValues(sp + 32)); }", -1)
				gojs = strings.Replace(gojs, "v, loadString(sp + 16));", "v, s);", -1)
			}
			gojs = strings.Replace(gojs, "})();", wasm_js(), -1)
			utils.Save(filepath.Join(depPath, "go.js"), gojs)
		} else {
			gojs := utils.Load(filepath.Join(depPath, "go.js"))
			gojs = strings.Replace(gojs, "(function() {", "Module._goMain = function() {", -1)
			utils.Save(filepath.Join(depPath, "go.js"), strings.Replace(gojs, "}).call(this);", "};", -1))
		}

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/.", depPath)

		if fast {
			break
		}

		utils.Save(filepath.Join(depPath, "c_main_wrapper_js.cpp"), js_c_main_wrapper(target))
		env, _, _, _ := cmd.BuildEnv(target, "", "")
		cmdD := exec.Command(filepath.Join(env["EMSCRIPTEN"], "em++"), "c_main_wrapper_js.cpp", target+".js_plugin_import.cpp")
		cmdD.Dir = depPath
		newArgs := templater.GetiOSClang(target, "", depPath)
		if utils.ExistsFile(filepath.Join(depPath, target+".js_qml_plugin_import.cpp")) {
			cmdD.Args = append(cmdD.Args, target+".js_qml_plugin_import.cpp")
		}

		for rccFile := range rcc.ResourceNames {
			cmdD.Args = append(cmdD.Args, rccFile)
		}
		rcc.ResourceNames = make(map[string]string)

		for mocFile := range moc.ResourceNames {
			cmdD.Args = append(cmdD.Args, mocFile)
		}
		moc.ResourceNames = make(map[string]string)

		//TODO: use "go list" deps instead ? and get rid of "build_static" ->
		//also remove GOCACHE stale trickery once done
		for _, l := range parser.LibDeps["build_static"] {
			for _, ml := range parser.GetLibs() {
				if strings.ToLower(l) == strings.ToLower(ml) {
					if (strings.ToLower(l) == "qml" || strings.ToLower(l) == "quick") && !cmd.ImportsQmlOrQuick() {
						continue
					}
					cmdD.Args = append(cmdD.Args, utils.GoQtPkgPath(strings.ToLower(l), strings.ToLower(l)+"-minimal.cpp"))
					break
				}
			}
		}

		for _, l := range parser.LibDeps[parser.MOC] {
			for _, ml := range parser.GetLibs() {
				if strings.ToLower(l) == strings.ToLower(ml) {
					cmdD.Args = append(cmdD.Args, utils.GoQtPkgPath(strings.ToLower(l), strings.ToLower(l)+"-minimal.cpp"))
					break
				}
			}
		}
		/*
			if !utils.QT_FAT() {
				tags = append(tags, "minimal")
			}
			if tagsCustom != "" {
				tags = append(tags, strings.Split(tagsCustom, " ")...)
			}
			lcmd := utils.GoList("{{join .Deps \"|\"}}", utils.BuildTags(tags))
			lcmd.Dir = path
			for k, v := range env {
				lcmd.Env = append(lcmd.Env, fmt.Sprintf("%v=%v", k, v))
			}
			for _, l := range strings.Split(strings.TrimSpace(utils.RunCmd(lcmd, "go list deps")), "|") {
				for _, ml := range parser.GetLibs() {
					if strings.HasSuffix(strings.ToLower(l), "github.com/bluszcz/cutego/"+strings.ToLower(ml)) {
						cmd.Args = append(cmd.Args, utils.GoQtPkgPath(strings.ToLower(ml), strings.ToLower(ml)+"-minimal.cpp"))
						break
					}
				}
			}
		*/
		//<-

		//TODO: check if minimal packages are stale and skip main.js rebuild this if they aren't
		cmdD.Args = append(cmdD.Args, newArgs...)
		cmdD.Args = append(cmdD.Args, []string{"-o", "main.js"}...)
		for key, value := range env {
			cmdD.Env = append(cmdD.Env, fmt.Sprintf("%v=%v", key, value))
		}
		utils.RunCmd(cmdD, fmt.Sprintf("compile wrapper for %v (%v) on %v", target, target, runtime.GOOS))

		utils.RemoveAll(filepath.Join(depPath, "c_main_wrapper_js.cpp"))
		utils.RemoveAll(filepath.Join(depPath, target+".js_plugin_import.cpp"))
		utils.RemoveAll(filepath.Join(depPath, target+".js_qml_plugin_import.cpp"))
		utils.RemoveAll(filepath.Join(depPath, "go.js.map"))
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			if utils.UseGOMOD(path) {
				utils.RunCmd(exec.Command("chown", "-R", idug, filepath.Dir(utils.GOMOD(path))), "chown files to user")
			} else {
				utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
			}
		}
	}
}
