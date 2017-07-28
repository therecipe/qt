package deploy

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

func build(mode, target, path, ldFlagsCustom, tagsCustom, name, depPath string, fast bool) {
	env, tags, ldFlags, out := cmd.BuildEnv(target, name, depPath)
	if (!fast || utils.QT_STUB()) && !utils.QT_FAT() {
		tags = append(tags, "minimal")
	}
	if ldFlagsCustom != "" {
		ldFlags = append(ldFlags, strings.Split(ldFlagsCustom, " ")...)
	}
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}
	if utils.QT_DEBUG_QML() {
		out = filepath.Join(depPath, name)
	}

	var ending string
	switch target {
	case "android", "android-emulator", "ios", "ios-simulator":
		utils.Save(filepath.Join(path, "cgo_main_wrapper.go"), "package main\nimport \"C\"\n//export go_main_wrapper\nfunc go_main_wrapper() { main() }")
	case "windows":
		ending = ".exe"
	case "sailfish", "sailfish-emulator":
		build_sailfish(mode, target, path, ldFlagsCustom, name, depPath)
		return
	}

	builds := []struct{}{{}}
	if target == "ios" || target == "ios-simulator" {
		builds = append(builds, struct{}{})
	}
	for i := range builds {
		if i == 1 {
			switch target {
			case "ios":
				out = strings.Replace(out, "libgo.a", "libgo_armv7.a", -1)
			case "ios-simulator":
				out = strings.Replace(out, "libgo.a", "libgo_386.a", -1)
			}
		}

		cmd := exec.Command("go", "build", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v", fmt.Sprintf("-ldflags=\"%v\"", strings.Join(ldFlags, "\" \"")), "-o", out+ending)
		cmd.Dir = path

		if fast && !utils.QT_STUB() {
			cmd.Args = append(cmd.Args, "-i")
		}

		cmd.Args = append(cmd.Args, fmt.Sprintf("-tags=\"%v\"", strings.Join(tags, "\" \"")))

		if target != runtime.GOOS {
			cmd.Args = append(cmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", strings.Replace(target, "-", "_", -1), env["GOOS"], env["GOARCH"]))}...)
		}

		switch target {
		case "android", "android-emulator":
			cmd.Args = append(cmd.Args, "-buildmode", "c-shared")
		case "ios", "ios-simulator":
			cmd.Args = append(cmd.Args, "-buildmode", "c-archive")
		}

		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}

		if i == 1 {
			switch target {
			case "ios":
				var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
				tmp = strings.Replace(tmp, "GOARCH=arm64", "GOARCH=arm", -1)
				cmd.Env = append(strings.Split(tmp, "|"), "GOARM=7")
			case "ios-simulator":
				var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch x86_64", "-arch i386", -1)
				tmp = strings.Replace(tmp, "GOARCH=amd64", "GOARCH=386", -1)
				cmd.Env = strings.Split(tmp, "|")
			}
		}

		utils.RunCmd(cmd, fmt.Sprintf("build for %v on %v", target, runtime.GOOS))

		if target == "darwin" {
			strip := exec.Command("strip", "-x", out)
			strip.Dir = path
			utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", target, runtime.GOOS))
		}
	}

	utils.RemoveAll(filepath.Join(path, "cgo_main_wrapper.go"))
}

func build_sailfish(mode, target, path, ldFlagsCustom, name, depPath string) {
	if !strings.Contains(path, utils.MustGoPath()) {
		utils.Log.Panicln("Project needs to be inside GOPATH; have:", path, "want:", utils.MustGoPath())
	}

	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "registervm", filepath.Join(utils.SAILFISH_DIR(), "mersdk", "MerSDK", "MerSDK.vbox")), fmt.Sprintf("register mersdk for %v on %v", target, runtime.GOOS))
	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "MerSDK", "--name", "GOROOT", "--hostpath", runtime.GOROOT(), "--automount"), fmt.Sprintf("share GOROOT dir for %v on %v", target, runtime.GOOS))
	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "MerSDK", "--name", "GOPATH", "--hostpath", utils.MustGoPath(), "--automount"), fmt.Sprintf("share GOPATH dir for %v on %v", target, runtime.GOOS))

	if runtime.GOOS == "windows" {
		utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "MerSDK"), fmt.Sprintf("start vbox mersdk for %v on %v", target, runtime.GOOS))
	} else {
		utils.RunCmdOptional(exec.Command("nohup", filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "MerSDK"), fmt.Sprintf("start vbox mersdk for %v on %v", target, runtime.GOOS))
	}

	time.Sleep(10 * time.Second)

	arch := "i486"
	if target == "sailfish" {
		arch = "armv7hl"
	}

	sailfish_ssh("2222", "root", "ln", "-s", fmt.Sprintf("/opt/cross/bin/%v-meego-linux-gnu-as", arch), fmt.Sprintf("/opt/cross/libexec/gcc/%v-meego-linux-gnu/4.8.3/as", arch))
	sailfish_ssh("2222", "root", "ln", "-s", fmt.Sprintf("/opt/cross/bin/%v-meego-linux-gnu-ld", arch), fmt.Sprintf("/opt/cross/libexec/gcc/%v-meego-linux-gnu/4.8.3/ld", arch))

	//TODO:
	var err error
	if target == "sailfish-emulator" {
		err = sailfish_ssh("2222", "root", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=386", "CGO_ENABLED=1", "CC=/opt/cross/bin/i486-meego-linux-gnu-gcc", "CXX=/opt/cross/bin/i486-meego-linux-gnu-g++", "CPATH=/srv/mer/targets/SailfishOS-i486/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-i486/usr/lib:/srv/mer/targets/SailfishOS-i486/lib", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-i486/", "go", "build", "-ldflags=\"-s -w\"", "-tags=\"minimal sailfish_emulator\"", "-o", "deploy/"+target+"/harbour-"+name)
	} else {
		err = sailfish_ssh("2222", "root", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=arm", "GOARM=7", "CGO_ENABLED=1", "CC=/opt/cross/bin/armv7hl-meego-linux-gnueabi-gcc", "CXX=/opt/cross/bin/armv7hl-meego-linux-gnueabi-g++", "CPATH=/srv/mer/targets/SailfishOS-armv7hl/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-armv7hl/usr/lib:/srv/mer/targets/SailfishOS-armv7hl/lib", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-armv7hl/", "go", "build", "-ldflags=\"-s -w\"", "-tags=\"minimal sailfish\"", "-o", "deploy/"+target+"/harbour-"+name)
	}
	if err != nil {
		println(err.Error())
		utils.Log.Panicf("failed to build for %v on %v", target, runtime.GOOS)
	}
}
