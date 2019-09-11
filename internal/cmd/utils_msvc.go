package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

var rep = map[string]string{
	"@@MEAA": "@@UEAA",
	"@@MEBA": "@@UEBA",
	"@@IEAA": "@@QEAA",
	"@@AEBA": "@@QEBA",
	"@@EEBA": "@@UEBA",
	"@@EEAA": "@@UEAA",
	"@@IEBA": "@@QEBA",
	"@@AEAA": "@@QEAA",

	"@@CA": "@@SA",
	"@@KA": "@@SA",
	"@@1U": "@@2U",

	"@QCoreApplication@@0PEAV": "@QCoreApplication@@2PEAV",
}

var block = map[string]string{"@@AEBA": "@Connection@QMetaObject"}

func libs() (libs []string) {

	filepath.Walk(filepath.Join(utils.QT_INSTALL_PREFIX("windows"), "lib"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		if filepath.Ext(info.Name()) == ".lib" && strings.HasPrefix(info.Name(), "Qt5") && !strings.HasSuffix(info.Name(), ".lib.orig") &&
			(!strings.HasSuffix(info.Name(), "d.lib") || strings.HasSuffix(info.Name(), "board.lib")) {
			libs = append(libs, path)
		}
		return nil
	})

	return
}

func PatchLibs() {
	for _, l := range libs() {
		patchLib(l)
	}
}

func patchLib(l string) {
	utils.Log.WithField("lib", l).Debugln("patching lib")

	if utils.ExistsFile(strings.Replace(l, ".lib", ".lib.orig", -1)) {
		utils.RemoveAll(l)
		os.Rename(strings.Replace(l, ".lib", ".lib.orig", -1), l)
	}
	data, _ := ioutil.ReadFile(l)

	for _, prr := range symbolsSliceFor(l, false, true) {
		for pr, pa := range rep {
			if !bytes.Contains(prr, []byte(pr)) {
				continue
			}
			data = bytes.Replace(data, prr, bytes.Replace(prr, []byte(pr), []byte(pa), -1), -1)
			if pref, ok := block[pr]; ok && bytes.Contains(prr, []byte(pref)) {
				data = bytes.Replace(data, append([]byte(pref), pa...), append([]byte(pref), pr...), -1)
			}
		}
	}
	utils.SaveBytes(l, data)
}

func PatchBinary(fn string) {
	utils.Log.WithField("file", fn).Debugln("patching binary")

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		utils.Log.WithError(err).Errorln("failed to load binary")
	}

	for _, l := range libs() {
		sb := symbolsSliceFor(strings.Replace(l, ".lib", ".lib.orig", -1), true, false)
		for i, s := range symbolsSliceFor(l, true, false) {
			if bytes.Equal(s, sb[i]) {
				continue
			}

			if bytes.Contains(data, s) {
				//println(string(s), "patch->", string(sb[i]))
				data = bytes.Replace(data, s, sb[i], -1)
			}
		}
	}
	utils.SaveBytes(fn, data)
}

func symbolsSliceFor(fn string, skipDll bool, backup bool) (symbols [][]byte) {

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		utils.Log.WithError(err).Debugln("failed to load")
		return nil
	}
	if backup {
		utils.SaveBytes(strings.Replace(fn, ".lib", ".lib.orig", -1), data)
	}

	for _, s := range bytes.Split(data, []byte("?")) {
		if !bytes.Contains(s, []byte("@@")) || (skipDll && bytes.Contains(s, []byte(".dll"))) {
			continue
		}
		if bytes.Contains(s, []byte("__")) {
			symbols = append(symbols, bytes.Split(s, []byte("__"))[0])
		} else {
			symbols = append(symbols, s)
		}
	}

	return
}

func testLibs() {
	for _, l := range libs() {
		test(l)
	}
}

func test(fn string) {
	skipDll := true

	patchLib(fn)
	PatchBinary(fn)

	var j int
	sb := symbolsSliceFor(strings.Replace(fn, ".lib", ".lib.orig", -1), skipDll, false)
	for i, s := range symbolsSliceFor(fn, skipDll, false) {
		if bytes.Equal(s, sb[i]) {
			continue
		}
		j++
	}
	if j != 0 {
		println("symbol diff", j)
	}

	c, _ := ioutil.ReadFile(fn)
	a := sha256.New()
	a.Write(c)
	ch := hex.EncodeToString(a.Sum(nil))

	c, _ = ioutil.ReadFile(strings.Replace(fn, ".lib", ".lib.orig", -1))
	a = sha256.New()
	a.Write(c)
	bh := hex.EncodeToString(a.Sum(nil))

	if ch != bh {
		println("hash diff", ch, bh)
	}
}
