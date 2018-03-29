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
		if !utils.QT_SAILFISH() {
			build_sailfish(mode, target, path, ldFlagsCustom, name, depPath)
			return
		}
	}

	var pattern string
	if strings.Contains(runtime.Version(), "1.1") {
		pattern = "all="
	}

	cmd := exec.Command("go", "build", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v", fmt.Sprintf("-ldflags=%v\"%v\"", pattern, strings.Join(ldFlags, "\" \"")), "-o", out+ending)
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

	utils.RunCmd(cmd, fmt.Sprintf("build for %v on %v", target, runtime.GOOS))

	if target == "darwin" {
		strip := exec.Command("strip", "-x", out) //TODO: -u -r
		strip.Dir = path
		utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", target, runtime.GOOS))
	}

	utils.RemoveAll(filepath.Join(path, "cgo_main_wrapper.go"))
}

func build_sailfish(mode, target, path, ldFlagsCustom, name, depPath string) {
	if !strings.Contains(path, utils.MustGoPath()) {
		utils.Log.Panicln("Project needs to be inside GOPATH; have:", path, "want:", utils.MustGoPath())
	}

	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "registervm", filepath.Join(utils.SAILFISH_DIR(), "mersdk", "Sailfish OS Build Engine", "Sailfish OS Build Engine.vbox")), fmt.Sprintf("register mersdk for %v on %v", target, runtime.GOOS))
	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "Sailfish OS Build Engine", "--name", "GOROOT", "--hostpath", runtime.GOROOT(), "--automount"), fmt.Sprintf("share GOROOT dir for %v on %v", target, runtime.GOOS))
	utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "Sailfish OS Build Engine", "--name", "GOPATH", "--hostpath", utils.MustGoPath(), "--automount"), fmt.Sprintf("share GOPATH dir for %v on %v", target, runtime.GOOS))

	if runtime.GOOS == "windows" {
		utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "Sailfish OS Build Engine"), fmt.Sprintf("start vbox mersdk for %v on %v", target, runtime.GOOS))
	} else {
		utils.RunCmdOptional(exec.Command("nohup", filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "Sailfish OS Build Engine"), fmt.Sprintf("start vbox mersdk for %v on %v", target, runtime.GOOS))
	}

	time.Sleep(10 * time.Second)

	for _, l := range []string{"libmpc.so.3", "libmpfr.so.4", "libgmp.so.10", "libpthread_nonshared.a", "libc_nonshared.a"} {
		sailfish_ssh("2222", "root", "ln", "-s", fmt.Sprintf("/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/usr/lib/%v", l), fmt.Sprintf("/usr/lib/%v", l))
	}

	arch, gcc := "i486", "gnu"
	if target == "sailfish" {
		arch, gcc = "armv7hl", "gnueabi"
	}

	sailfish_ssh("2222", "root", "ln", "-s", fmt.Sprintf("/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/%v-meego-linux-%v-as", arch, gcc), fmt.Sprintf("/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/libexec/gcc/%v-meego-linux-%v/4.8.3/as", arch, gcc))
	sailfish_ssh("2222", "root", "ln", "-s", fmt.Sprintf("/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/%v-meego-linux-%v-ld", arch, gcc), fmt.Sprintf("/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/libexec/gcc/%v-meego-linux-%v/4.8.3/ld", arch, gcc))

	var pattern string
	if strings.Contains(runtime.Version(), "1.1") {
		pattern = "all="
	}

	//TODO:
	var err error
	if target == "sailfish-emulator" {
		err = sailfish_ssh("2222", "root", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH/", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=386", "CGO_ENABLED=1", "CGO_CFLAGS_ALLOW=.*", "CGO_CXXFLAGS_ALLOW=.*", "CGO_LDFLAGS_ALLOW=.*", "CC=/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/i486-meego-linux-gnu-gcc", "CXX=/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/i486-meego-linux-gnu-g++", "CPATH=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/usr/lib:/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/lib:/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/usr/lib/pulseaudio", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-i486/", "go", "build", fmt.Sprintf("-ldflags=%v\"-s -w\"", pattern), "-tags=\"minimal sailfish_emulator\"", "-o", "deploy/"+target+"/harbour-"+name)
	} else {
		err = sailfish_ssh("2222", "root", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH/", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=arm", "GOARM=7", "CGO_ENABLED=1", "CGO_CFLAGS_ALLOW=.*", "CGO_CXXFLAGS_ALLOW=.*", "CGO_LDFLAGS_ALLOW=.*", "CC=/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/armv7hl-meego-linux-gnueabi-gcc", "CXX=/srv/mer/toolings/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"/opt/cross/bin/armv7hl-meego-linux-gnueabi-g++", "CPATH=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-armv7hl/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-armv7hl/usr/lib:/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-armv7hl/lib:/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-armv7hl/usr/lib/pulseaudio", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-"+utils.QT_SAILFISH_VERSION()+"-armv7hl/", "go", "build", "-x", "-v", fmt.Sprintf("-ldflags=%v\"-s -w\"", pattern), "-tags=\"minimal sailfish\"", "-o", "deploy/"+target+"/harbour-"+name)
	}
	if err != nil {
		println(err.Error())
		utils.Log.Panicf("failed to build for %v on %v", target, runtime.GOOS)
	}
}
