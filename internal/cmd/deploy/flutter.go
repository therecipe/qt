package deploy

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func flutter(target string, path string) {
	var ext string
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	fPath, err := exec.LookPath("flutter" + ext)
	if err != nil {
		utils.Log.WithError(err).Error("failed to find the flutter binary in your PATH")
		return
	}

	//symlink engine libs

	defer func() {
		if utils.QT_DOCKER() {
			for _, hiddenDir := range []string{".packages", ".dart_tool"} {
				utils.RemoveAll(filepath.Join(path, hiddenDir))
			}
		}

		switch target {
		case "windows":
			for _, f := range []string{"flutter_engine.dll", "flutter_engine.dll.lib"} {
				utils.RemoveAll(utils.GoQtPkgPath("flutter", f))
				os.Symlink(filepath.Join(path, f), utils.GoQtPkgPath("flutter", f))
			}

		case "darwin":
			f := "FlutterEmbedder.framework"
			utils.RemoveAll(utils.GoQtPkgPath("flutter", f))
			os.Symlink(filepath.Join(path, f), utils.GoQtPkgPath("flutter", f))

		case "linux":
			f := "libflutter_engine.so"
			utils.RemoveAll(utils.GoQtPkgPath("flutter", f))
			os.Symlink(filepath.Join(path, f), utils.GoQtPkgPath("flutter", f))
		}
	}()

	//build the flutter binaries

	if (target == "linux" && utils.QT_DOCKER()) || !utils.QT_DOCKER() {
		cmd := exec.Command(fPath, "build", "bundle")
		cmd.Dir = path
		utils.RunCmd(cmd, "build the flutter bundle")
	}

	//check cache

	switch target {
	case "windows":
		if utils.ExistsFile(filepath.Join(path, "flutter_engine.dll")) {
			utils.Log.Debugf("flutter deps already cached; skipping download for %v", target)
			return
		}
	case "darwin":
		if utils.ExistsDir(filepath.Join(path, "FlutterEmbedder.framework")) {
			utils.Log.Debugf("flutter deps already cached; skipping download for %v", target)
			return
		}
	case "linux":
		if utils.ExistsFile(filepath.Join(path, "libflutter_engine.so")) {
			utils.Log.Debugf("flutter deps already cached; skipping download for %v", target)
			return
		}
	}

	//get engine version

	var vs struct {
		Channel        string `json:"channel"`
		EngineRevision string `json:"engineRevision"`
	}
	err = json.Unmarshal([]byte(utils.RunCmd(exec.Command(fPath, "--machine", "--version"), "readout the flutter engine version")), &vs)
	if err != nil {
		utils.Log.WithError(err).Error("failed to unmarshal the flutter engine version response")
	}

	//download engine libs

	for _, t := range []string{"windows", "darwin", "linux"} {
		if target != t && !utils.QT_DOCKER() {
			continue
		}

		url := fmt.Sprintf("https://storage.googleapis.com/flutter_infra/flutter/%v/%v-x64/", vs.EngineRevision, t)
		switch t {
		case "windows":
			url += "windows-x64-embedder.zip"
		case "darwin":
			url += "FlutterEmbedder.framework.zip"
		case "linux":
			url += "linux-x64-embedder"
		}

		unzip(download(url), path, t)
	}

	//download additional artifacts

	if target != "darwin" {
		unzip(download(fmt.Sprintf("https://storage.googleapis.com/flutter_infra/flutter/%v/%v-x64/artifacts.zip", vs.EngineRevision, target)), path, target)
	}
}

func download(url string) []byte {
	utils.Log.WithField("url", url).Debug("download file")

	resp, err := http.Get(url)
	if err != nil {
		utils.Log.WithError(err).Errorf("failed to download %v", url)
		return nil
	}

	bb := new(bytes.Buffer)
	copyWithProgress(bb, resp.Body, func(off int64) {
		utils.Log.Debugf("%v => %v%%", filepath.Base(url), off/(resp.ContentLength/100))
	})
	resp.Body.Close()

	return bb.Bytes()
}

func copyWithProgress(w io.Writer, r io.Reader, callback func(off int64)) error {
	tee := io.TeeReader(r, w)
	buf := make([]byte, bytes.MinRead*10)
	off := int64(0)
	for count := 0; ; count++ {
		n, err := tee.Read(buf)
		off += int64(n)
		if count%100 == 0 {
			callback(off)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	callback(off)
	return nil
}

func unzip(data []byte, dst string, target string) {
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		utils.Log.WithError(err).Errorf("failed to read zip archive for %v", target)
	}
	for _, f := range r.File {

		if f.FileInfo().IsDir() {
			os.MkdirAll(filepath.Join(dst, f.Name), f.Mode())
			continue
		}

		if f.Mode()&os.ModeSymlink != 0 {

			fr, errR := f.Open()
			if errR != nil {
				utils.Log.WithError(errR).Error("failed to read file")
			}

			bb := new(bytes.Buffer)
			io.Copy(bb, fr)
			fr.Close()

			os.Symlink(filepath.Join(filepath.Dir(filepath.Join(dst, f.Name)), bb.String()), filepath.Join(dst, f.Name))
			continue
		}

		if strings.HasSuffix(f.Name, ".zip") {
			dst = filepath.Join(dst, strings.TrimSuffix(f.Name, ".zip"))
			utils.MkdirAll(dst)

			fr, errR := f.Open()
			if errR != nil {
				utils.Log.WithError(errR).Error("failed to read file")
			}

			bb := new(bytes.Buffer)
			io.Copy(bb, fr)
			fr.Close()

			unzip(bb.Bytes(), dst, target)
			continue
		}

		if ((f.Name == "libflutter_engine.so" || f.Name == "flutter_engine.dll" ||
			f.Name == "flutter_engine.dll.lib" || f.Name == "icudtl.dat") && target != "darwin") || target == "darwin" {

			fw, errC := os.OpenFile(filepath.Join(dst, f.Name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if errC != nil {
				utils.Log.WithError(errC).Error("failed to create file")
			}

			fr, errR := f.Open()
			if errR != nil {
				utils.Log.WithError(errR).Error("failed to read file")
			}

			io.Copy(fw, fr)

			fw.Close()
			fr.Close()
		}
	}
}
