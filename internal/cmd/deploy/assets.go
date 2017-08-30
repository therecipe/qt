package deploy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/therecipe/qt/internal/utils"
)

//linux

func linux_sh(target, name string) string {
	bb := new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprint(bb, "#!/bin/bash\n")
	fmt.Fprint(bb, "appname=`basename $0 | sed s,\\.sh$,,`\n\n")
	fmt.Fprint(bb, "dirname=`dirname $0`\n")
	fmt.Fprint(bb, "tmp=\"${dirname#?}\"\n\n")
	fmt.Fprint(bb, "if [ \"${dirname%$tmp}\" != \"/\" ]; then\n")
	fmt.Fprint(bb, "dirname=$PWD/$dirname\n")
	fmt.Fprint(bb, "fi\n")

	if strings.HasPrefix(target, "rpi") {
		fmt.Fprint(bb, "export DISPLAY=:0\n")
		fmt.Fprint(bb, "export LD_PRELOAD=\"/opt/vc/lib/libGLESv2.so /opt/vc/lib/libEGL.so\"\n")
	}

	if utils.QT_PKG_CONFIG() {
		libDir := strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), fmt.Sprintf("get lib dir for %v on %v", target, runtime.GOOS)))
		miscDir := utils.QT_MISC_DIR()

		fmt.Fprintf(bb, "export LD_LIBRARY_PATH=%v\n", libDir)
		fmt.Fprintf(bb, "export QT_PLUGIN_PATH=$%v\n", filepath.Join(miscDir, "plugins"))
		fmt.Fprintf(bb, "export QML_IMPORT_PATH=%v\n", filepath.Join(miscDir, "qml"))
		fmt.Fprintf(bb, "export QML2_IMPORT_PATH=%v\n", filepath.Join(miscDir, "qml"))
	} else {
		libDir := "lib"
		if name == libDir {
			libDir = "libs"
		}
		fmt.Fprint(bb, "export LD_LIBRARY_PATH=$dirname/"+libDir+"\n")
		fmt.Fprint(bb, "export QT_PLUGIN_PATH=$dirname/plugins\n")
		fmt.Fprint(bb, "export QML_IMPORT_PATH=$dirname/qml\n")
		fmt.Fprint(bb, "export QML2_IMPORT_PATH=$dirname/qml\n")
	}
	fmt.Fprint(bb, "$dirname/$appname \"$@\"\n")

	return bb.String()
}

//android

func android_config(target, path, depPath string) string {
	jsonStruct := &struct {
		Qt                            string `json:"qt"`
		Sdk                           string `json:"sdk"`
		SdkBuildToolsRevision         string `json:"sdkBuildToolsRevision"`
		Ndk                           string `json:"ndk"`
		Toolchainprefix               string `json:"toolchain-prefix"`
		Toolprefix                    string `json:"tool-prefix"`
		Toolchainversion              string `json:"toolchain-version"`
		Ndkhost                       string `json:"ndk-host"`
		Targetarchitecture            string `json:"target-architecture"`
		AndroidExtraLibs              string `json:"android-extra-libs"`
		AndroidPackageSourceDirectory string `json:"android-package-source-directory"`
		Qmlrootpath                   string `json:"qml-root-path"`
		Applicationbinary             string `json:"application-binary"`
	}{
		Qt:  filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7"),
		Sdk: utils.ANDROID_SDK_DIR(),
		SdkBuildToolsRevision: "26.0.0",
		Ndk:                           utils.ANDROID_NDK_DIR(),
		Toolchainprefix:               "arm-linux-androideabi",
		Toolprefix:                    "arm-linux-androideabi",
		Toolchainversion:              "4.9",
		Ndkhost:                       runtime.GOOS + "-x86_64",
		Targetarchitecture:            "armeabi-v7a",
		AndroidExtraLibs:              filepath.Join(depPath, "libgo_base.so"),
		AndroidPackageSourceDirectory: filepath.Join(path, target),
		Qmlrootpath:                   path,
		Applicationbinary:             filepath.Join(depPath, "libgo.so"),
	}

	if target == "android-emulator" {
		jsonStruct.Qt = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_x86")
		jsonStruct.Toolchainprefix = "x86"
		jsonStruct.Toolprefix = "i686-linux-android"
		jsonStruct.Targetarchitecture = "x86"
	}

	if utils.QT_DOCKER() {
		jsonStruct.AndroidExtraLibs += "," + filepath.Join(os.Getenv("HOME"), "openssl-1.0.2k", "libcrypto.so") + "," + filepath.Join(os.Getenv("HOME"), "openssl-1.0.2k", "libssl.so")
	}

	out, err := json.Marshal(jsonStruct)
	if err != nil {
		utils.Log.WithError(err).Panicf("failed to create json-config file for androiddeployqt on %v", runtime.GOOS)
	}
	return strings.Replace(string(out), `\\`, `/`, -1)
}

func android_gradle_wrapper() string {
	return `distributionBase=GRADLE_USER_HOME
distributionPath=wrapper/dists
zipStoreBase=GRADLE_USER_HOME
zipStorePath=wrapper/dists
distributionUrl=https\://services.gradle.org/distributions/gradle-2.14.1-all.zip`
}

//darwin

func darwin_sh(name string) string {
	return fmt.Sprintf("#!/bin/bash\ncd \"${0%%/*}\"\n./%v_app \"$@\"", name)
}

func darwin_plist(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>NSPrincipalClass</key>
	<string>NSApplication</string>
	<key>CFBundleIconFile</key>
	<string>%[1]v.icns</string>
	<key>CFBundleName</key>
	<string>%[1]v</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleGetInfoString</key>
	<string>Created by Qt/QMake</string>
	<key>CFBundleSignature</key>
	<string>????</string>
	<key>CFBundleExecutable</key>
	<string>%[1]v</string>
	<key>CFBundleIdentifier</key>
	<string>com.ident.%[1]v</string>
	<key>NOTE</key>
	<string>This file was generated by Qt/QMake.</string>
</dict>
</plist>
`, name)
}

func darwin_pkginfo() string {
	return "APPL????\n"
}

//ios

func ios_plist(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>en</string>
	<key>CFBundleExecutable</key>
	<string>main</string>
	<key>CFBundleIdentifier</key>
	<string>com.identifier.%[1]v</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>%[1]v</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0</string>
	<key>CFBundleSignature</key>
	<string>????</string>
	<key>CFBundleVersion</key>
	<string>1.0</string>
	<key>LSRequiresIPhoneOS</key>
	<true/>
	<key>UILaunchStoryboardName</key>
	<string>LaunchScreen</string>
	<key>UIRequiredDeviceCapabilities</key>
	<array>
		<string>arm64</string>
	</array>
	<key>UISupportedInterfaceOrientations</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationPortraitUpsideDown</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>UISupportedInterfaceOrientations~ipad</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationPortraitUpsideDown</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>QtRunLoopIntegrationDisableSeparateStack</key>
	<true/>
	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSAllowsArbitraryLoads</key>
		<true/>
	</dict>
</dict>
</plist>
`, name)
}

func ios_launchscreen(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<document type="com.apple.InterfaceBuilder3.CocoaTouch.XIB" version="3.0" toolsVersion="10117" systemVersion="15E65" targetRuntime="iOS.CocoaTouch" propertyAccessControl="none" useAutolayout="YES" launchScreen="YES" useTraitCollections="YES">
	    <dependencies>
	        <deployment identifier="iOS"/>
	        <plugIn identifier="com.apple.InterfaceBuilder.IBCocoaTouchPlugin" version="10085"/>
	        <capability name="Constraints with non-1.0 multipliers" minToolsVersion="5.1"/>
	    </dependencies>
	    <objects>
	        <placeholder placeholderIdentifier="IBFilesOwner" id="-1" userLabel="File's Owner"/>
	        <placeholder placeholderIdentifier="IBFirstResponder" id="-2" customClass="UIResponder"/>
	        <view contentMode="scaleToFill" id="iN0-l3-epB">
	            <rect key="frame" x="0.0" y="0.0" width="480" height="480"/>
	            <autoresizingMask key="autoresizingMask" widthSizable="YES" heightSizable="YES"/>
	            <subviews>
	                <label opaque="NO" clipsSubviews="YES" userInteractionEnabled="NO" contentMode="left" horizontalHuggingPriority="251" verticalHuggingPriority="251" text="%v" textAlignment="center" lineBreakMode="middleTruncation" baselineAdjustment="alignBaselines" minimumFontSize="18" translatesAutoresizingMaskIntoConstraints="NO" id="kId-c2-rCX">
	                    <rect key="frame" x="20" y="140" width="441" height="43"/>
	                    <fontDescription key="fontDescription" type="boldSystem" pointSize="36"/>
	                    <color key="textColor" red="0.0" green="0.0" blue="0.0" alpha="1" colorSpace="calibratedRGB"/>
	                    <nil key="highlightedColor"/>
	                </label>
	            </subviews>
	            <color key="backgroundColor" white="1" alpha="1" colorSpace="custom" customColorSpace="calibratedWhite"/>
	            <constraints>
	                <constraint firstItem="kId-c2-rCX" firstAttribute="centerY" secondItem="iN0-l3-epB" secondAttribute="bottom" multiplier="1/3" constant="1" id="Kid-kn-2rF"/>
	                <constraint firstAttribute="centerX" secondItem="kId-c2-rCX" secondAttribute="centerX" id="Koa-jz-hwk"/>
	                <constraint firstItem="kId-c2-rCX" firstAttribute="leading" secondItem="iN0-l3-epB" secondAttribute="leading" constant="20" symbolic="YES" id="fvb-Df-36g"/>
	            </constraints>
	            <nil key="simulatedStatusBarMetrics"/>
	            <freeformSimulatedSizeMetrics key="simulatedDestinationMetrics"/>
	            <point key="canvasLocation" x="404" y="445"/>
	        </view>
	    </objects>
	</document>
	`, name)
}

func ios_appicon() string {
	return `{
  "images" : [
    {
      "idiom" : "iphone",
      "size" : "29x29",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "29x29",
      "scale" : "3x"
    },
    {
      "idiom" : "iphone",
      "size" : "40x40",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "40x40",
      "scale" : "3x"
    },
    {
      "idiom" : "iphone",
      "size" : "60x60",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "60x60",
      "scale" : "3x"
    },
    {
      "idiom" : "ipad",
      "size" : "29x29",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "29x29",
      "scale" : "2x"
    },
    {
      "idiom" : "ipad",
      "size" : "40x40",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "40x40",
      "scale" : "2x"
    },
    {
      "idiom" : "ipad",
      "size" : "76x76",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "76x76",
      "scale" : "2x"
    }
  ],
  "info" : {
    "version" : 1,
    "author" : "xcode"
  }
}
`
}

func ios_xcodeproject(depPath string) string {
	return fmt.Sprintf(`// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 46;
	objects = {

/* Begin PBXBuildFile section */
		254BB84F1B1FD08900C56DE9 /* Images.xcassets in Resources */ = {isa = PBXBuildFile; fileRef = 254BB84E1B1FD08900C56DE9 /* Images.xcassets */; };
		254BB8681B1FD16500C56DE9 /* main in Resources */ = {isa = PBXBuildFile; fileRef = 254BB8671B1FD16500C56DE9 /* main */; };
		25916F411CE65FF600695115 /* LaunchScreen.xib in Resources */ = {isa = PBXBuildFile; fileRef = 25916F401CE65FF600695115 /* LaunchScreen.xib */; };
		25F26AED1CE6675E0045FFBA /* Default-568h@2x.png in Resources */ = {isa = PBXBuildFile; fileRef = 25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */; };
/* End PBXBuildFile section */

/* Begin PBXFileReference section */
		254BB83E1B1FD08900C56DE9 /* main.app */ = {isa = PBXFileReference; explicitFileType = wrapper.application; includeInIndex = 0; path = main.app; sourceTree = BUILT_PRODUCTS_DIR; };
		254BB8421B1FD08900C56DE9 /* Info.plist */ = {isa = PBXFileReference; lastKnownFileType = text.plist.xml; path = Info.plist; sourceTree = "<group>"; };
		254BB84E1B1FD08900C56DE9 /* Images.xcassets */ = {isa = PBXFileReference; lastKnownFileType = folder.assetcatalog; path = Images.xcassets; sourceTree = "<group>"; };
		254BB8671B1FD16500C56DE9 /* main */ = {isa = PBXFileReference; lastKnownFileType = "compiled.mach-o.executable"; path = main; sourceTree = "<group>"; };
		25916F401CE65FF600695115 /* LaunchScreen.xib */ = {isa = PBXFileReference; fileEncoding = 4; lastKnownFileType = file.xib; path = LaunchScreen.xib; sourceTree = "<group>"; };
		25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */ = {isa = PBXFileReference; lastKnownFileType = image.png; path = "Default-568h@2x.png"; sourceTree = "<group>"; };
/* End PBXFileReference section */

/* Begin PBXGroup section */
		254BB8351B1FD08900C56DE9 = {
			isa = PBXGroup;
			children = (
				254BB8671B1FD16500C56DE9 /* main */,
				254BB8421B1FD08900C56DE9 /* Info.plist */,
				254BB84E1B1FD08900C56DE9 /* Images.xcassets */,
				25916F401CE65FF600695115 /* LaunchScreen.xib */,
				25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */,
				254BB83F1B1FD08900C56DE9 /* products */,
			);
			sourceTree = "<group>";
			usesTabs = 0;
		};
		254BB83F1B1FD08900C56DE9 /* products */ = {
			isa = PBXGroup;
			children = (
				254BB83E1B1FD08900C56DE9 /* main.app */,
			);
			name = products;
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXNativeTarget section */
		254BB83D1B1FD08900C56DE9 /* main */ = {
			isa = PBXNativeTarget;
			buildConfigurationList = 254BB8611B1FD08900C56DE9 /* Build configuration list for PBXNativeTarget "main" */;
			buildPhases = (
				254BB83C1B1FD08900C56DE9 /* Resources */,
				259BC5361CE6BA19005B5A05 /* ShellScript */,
			);
			buildRules = (
			);
			dependencies = (
			);
			name = main;
			productName = main;
			productReference = 254BB83E1B1FD08900C56DE9 /* main.app */;
			productType = "com.apple.product-type.application";
		};
/* End PBXNativeTarget section */

/* Begin PBXProject section */
		254BB8361B1FD08900C56DE9 /* Project object */ = {
			isa = PBXProject;
			attributes = {
				LastUpgradeCheck = 0630;
				ORGANIZATIONNAME = Developer;
				TargetAttributes = {
					254BB83D1B1FD08900C56DE9 = {
						CreatedOnToolsVersion = 6.3.1;
					};
				};
			};
			buildConfigurationList = 254BB8391B1FD08900C56DE9 /* Build configuration list for PBXProject "project" */;
			compatibilityVersion = "Xcode 3.2";
			developmentRegion = English;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
				Base,
			);
			mainGroup = 254BB8351B1FD08900C56DE9;
			productRefGroup = 254BB83F1B1FD08900C56DE9 /* products */;
			projectDirPath = "";
			projectRoot = "";
			targets = (
				254BB83D1B1FD08900C56DE9 /* main */,
			);
		};
/* End PBXProject section */

/* Begin PBXResourcesBuildPhase section */
		254BB83C1B1FD08900C56DE9 /* Resources */ = {
			isa = PBXResourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
				254BB8681B1FD16500C56DE9 /* main in Resources */,
				25F26AED1CE6675E0045FFBA /* Default-568h@2x.png in Resources */,
				25916F411CE65FF600695115 /* LaunchScreen.xib in Resources */,
				254BB84F1B1FD08900C56DE9 /* Images.xcassets in Resources */,
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
/* End PBXResourcesBuildPhase section */

/* Begin PBXShellScriptBuildPhase section */
		259BC5361CE6BA19005B5A05 /* ShellScript */ = {
			isa = PBXShellScriptBuildPhase;
			buildActionMask = 2147483647;
			files = (
			);
			inputPaths = (
				"$(TARGET_BUILD_DIR)/$(EXECUTABLE_PATH)",
			);
			outputPaths = (
			);
			runOnlyForDeploymentPostprocessing = 0;
			shellPath = /bin/sh;
			shellScript = "cp %v/qt.conf $CODESIGNING_FOLDER_PATH/qt.conf;  test -d $CODESIGNING_FOLDER_PATH/qt_qml && rm -r $CODESIGNING_FOLDER_PATH/qt_qml;  mkdir -p $CODESIGNING_FOLDER_PATH/qt_qml &&  for p in %v/%v/ios/qml; do rsync -r --exclude='*.a' --exclude='*.prl' --exclude='*.qmltypes'  $p/ $CODESIGNING_FOLDER_PATH/qt_qml; done";
			showEnvVarsInLog = 0;
		};
/* End PBXShellScriptBuildPhase section */

/* Begin XCBuildConfiguration section */
		254BB8601B1FD08900C56DE9 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ALWAYS_SEARCH_USER_PATHS = NO;
				CLANG_CXX_LANGUAGE_STANDARD = "gnu++0x";
				CLANG_CXX_LIBRARY = "libc++";
				CLANG_ENABLE_MODULES = YES;
				CLANG_ENABLE_OBJC_ARC = YES;
				CLANG_WARN_BOOL_CONVERSION = YES;
				CLANG_WARN_CONSTANT_CONVERSION = YES;
				CLANG_WARN_DIRECT_OBJC_ISA_USAGE = YES_ERROR;
				CLANG_WARN_EMPTY_BODY = YES;
				CLANG_WARN_ENUM_CONVERSION = YES;
				CLANG_WARN_INT_CONVERSION = YES;
				CLANG_WARN_OBJC_ROOT_CLASS = YES_ERROR;
				CLANG_WARN_UNREACHABLE_CODE = YES;
				CLANG_WARN__DUPLICATE_METHOD_MATCH = YES;
				"CODE_SIGN_IDENTITY[sdk=iphoneos*]" = "iPhone Developer";
				COPY_PHASE_STRIP = NO;
				DEBUG_INFORMATION_FORMAT = "dwarf-with-dsym";
				ENABLE_NS_ASSERTIONS = NO;
				ENABLE_STRICT_OBJC_MSGSEND = YES;
				GCC_C_LANGUAGE_STANDARD = gnu99;
				GCC_NO_COMMON_BLOCKS = YES;
				GCC_WARN_64_TO_32_BIT_CONVERSION = YES;
				GCC_WARN_ABOUT_RETURN_TYPE = YES_ERROR;
				GCC_WARN_UNDECLARED_SELECTOR = YES;
				GCC_WARN_UNINITIALIZED_AUTOS = YES_AGGRESSIVE;
				GCC_WARN_UNUSED_FUNCTION = YES;
				GCC_WARN_UNUSED_VARIABLE = YES;
				IPHONEOS_DEPLOYMENT_TARGET = 10.0;
				MTL_ENABLE_DEBUG_INFO = NO;
				SDKROOT = iphoneos;
				TARGETED_DEVICE_FAMILY = "1,2";
				VALIDATE_PRODUCT = YES;
			};
			name = Release;
		};
		254BB8631B1FD08900C56DE9 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ASSETCATALOG_COMPILER_APPICON_NAME = AppIcon;
				INFOPLIST_FILE = Info.plist;
				LD_RUNPATH_SEARCH_PATHS = "$(inherited) @executable_path/Frameworks";
				PRODUCT_NAME = "$(TARGET_NAME)";
			};
			name = Release;
		};
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		254BB8391B1FD08900C56DE9 /* Build configuration list for PBXProject "project" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				254BB8601B1FD08900C56DE9 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
		254BB8611B1FD08900C56DE9 /* Build configuration list for PBXNativeTarget "main" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				254BB8631B1FD08900C56DE9 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = 254BB8361B1FD08900C56DE9 /* Project object */;
}
`, depPath, utils.QT_DIR(), utils.QT_VERSION_MAJOR())
}

func ios_plugins() string {
	return `#include <QtPlugin>
Q_IMPORT_PLUGIN(QIOSIntegrationPlugin)
Q_IMPORT_PLUGIN(AVFServicePlugin)
Q_IMPORT_PLUGIN(AVFMediaPlayerServicePlugin)
Q_IMPORT_PLUGIN(AudioCaptureServicePlugin)
Q_IMPORT_PLUGIN(CoreAudioPlugin)
Q_IMPORT_PLUGIN(QM3uPlaylistPlugin)
//Q_IMPORT_PLUGIN(ContextPlugin)
//Q_IMPORT_PLUGIN(QDDSPlugin)
Q_IMPORT_PLUGIN(QGifPlugin)
Q_IMPORT_PLUGIN(QICNSPlugin)
Q_IMPORT_PLUGIN(QICOPlugin)
Q_IMPORT_PLUGIN(QJpegPlugin)
Q_IMPORT_PLUGIN(QMacJp2Plugin)
Q_IMPORT_PLUGIN(QTgaPlugin)
Q_IMPORT_PLUGIN(QTiffPlugin)
Q_IMPORT_PLUGIN(QWbmpPlugin)
Q_IMPORT_PLUGIN(QWebpPlugin)
Q_IMPORT_PLUGIN(QQmlDebuggerServiceFactory)
Q_IMPORT_PLUGIN(QQmlInspectorServiceFactory)
Q_IMPORT_PLUGIN(QLocalClientConnectionFactory)
Q_IMPORT_PLUGIN(QQmlNativeDebugConnectorFactory)
Q_IMPORT_PLUGIN(QQuickProfilerAdapterFactory)
Q_IMPORT_PLUGIN(QQmlProfilerServiceFactory)
Q_IMPORT_PLUGIN(QQmlDebugServerFactory)
Q_IMPORT_PLUGIN(QTcpServerConnectionFactory)
//Q_IMPORT_PLUGIN(QGenericEnginePlugin)
`
}

func ios_qmlplugins() string {
	return `#include <QtPlugin>
Q_IMPORT_PLUGIN(QtQuick2Plugin)
Q_IMPORT_PLUGIN(QMultimediaDeclarativeModule)
Q_IMPORT_PLUGIN(QtQuickLayoutsPlugin)
Q_IMPORT_PLUGIN(QtQuickControls2Plugin)
Q_IMPORT_PLUGIN(QtQuickControls2MaterialStylePlugin)
Q_IMPORT_PLUGIN(QtQuickControls2UniversalStylePlugin)
Q_IMPORT_PLUGIN(QtQuick2DialogsPlugin)
Q_IMPORT_PLUGIN(QtQuickControls1Plugin)
Q_IMPORT_PLUGIN(QmlFolderListModelPlugin)
Q_IMPORT_PLUGIN(QmlSettingsPlugin)
Q_IMPORT_PLUGIN(QtQuickTemplates2Plugin)
Q_IMPORT_PLUGIN(QtQuick2DialogsPrivatePlugin)
Q_IMPORT_PLUGIN(QtQuick2WindowPlugin)
Q_IMPORT_PLUGIN(QtQmlModelsPlugin)
Q_IMPORT_PLUGIN(QtQuickExtrasPlugin)
Q_IMPORT_PLUGIN(QtGraphicalEffectsPlugin)
Q_IMPORT_PLUGIN(QtGraphicalEffectsPrivatePlugin)
Q_IMPORT_PLUGIN(QWebViewModule)
`
}

func ios_qtconf() string {
	return `[Paths]
Imports = qt_qml
Qml2Imports = qt_qml
`
}

//sailfish

func sailfish_spec(name string) string {
	return fmt.Sprintf(`#
# Do NOT Edit the Auto-generated Part!
# Generated by: spectacle version 0.27
#

Name:       harbour-%[1]v

# >> macros
# << macros

Summary:    Put your summary here
Version:    0.1
Release:    1
Group:      Qt/Qt
License:    MIT
Source0:    %%{name}-%%{version}.tar.bz2
#Requires:  mapplauncherd-booster-silica-qt5
#Requires:  nemo-qml-plugin-thumbnailer-qt5
Requires:   sailfishsilica-qt5
#Requires:  qt5-qtdocgallery
BuildRequires:  pkgconfig(sailfishapp)
BuildRequires:  pkgconfig(Qt5Quick)
BuildRequires:  pkgconfig(Qt5Qml)
BuildRequires:  pkgconfig(Qt5Core)
#BuildRequires: pkgconfig(qdeclarative5-boostable)
BuildRequires:  desktop-file-utils

%%description
Put your description here


%%prep
%%setup -q -n %%{name}-%%{version}

# >> setup
# << setup

%%build
# >> build pre
# << build pre

# >> build post
# << build post

%%install
rm -rf %%{buildroot}
# >> install pre
# << install pre
install -d %%{buildroot}%%{_bindir}
install -p -m 0755 %%(pwd)/%%{name} %%{buildroot}%%{_bindir}/%%{name}
install -d %%{buildroot}%%{_datadir}/applications
install -d %%{buildroot}%%{_datadir}/%%{name}
install -d %%{buildroot}%%{_datadir}/icons/hicolor/86x86/apps
install -m 0444 -t %%{buildroot}%%{_datadir}/icons/hicolor/86x86/apps %%{name}.png
install -p %%(pwd)/%[1]v.desktop %%{buildroot}%%{_datadir}/applications/%%{name}.desktop

# >> install post
# << install post

desktop-file-install --delete-original       \
  --dir %%{buildroot}%%{_datadir}/applications             \
   %%{buildroot}%%{_datadir}/applications/*.desktop

%%files
%%defattr(-,root,root,-)
%%{_bindir}
%%{_datadir}/%%{name}
%%{_datadir}/icons/hicolor/86x86/apps
%%{_datadir}/applications/%%{name}.desktop

# >> files
# << files`, name)
}

func sailfish_desktop(name string) string {
	return fmt.Sprintf(`[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
X-Nemo-Application-Type=generic
Comment=Put your comment here
Name=%[1]v
Icon=harbour-%[1]v
Exec=harbour-%[1]v`, name)
}

func sailfish_ssh(port, login string, cmd ...string) error {

	typ := "SailfishOS_Emulator"
	if port == "2222" {
		typ = "engine"
	}

	signer, err := ssh.ParsePrivateKey([]byte(utils.Load(filepath.Join(utils.SAILFISH_DIR(), "vmshare", "ssh", "private_keys", typ, login))))
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", "localhost:"+port, &ssh.ClientConfig{User: login, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()})
	if err != nil {
		return err
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		return err
	}

	output, err := sess.CombinedOutput(strings.Join(cmd, " "))
	if err != nil {
		utils.Log.WithField("cmd", strings.Join(cmd, " ")).Debugf("failed to run ssh cmd for %v on %v", typ, runtime.GOOS)
		return errors.New(string(output))
	}

	return nil
}
