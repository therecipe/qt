# Android

## Release and Branding

An Android application should be structured like the following:

```
project_name
├── android
│   ├── AndroidManifest.xml
│   ├── alias.txt
│   ├── build.gradle
│   ├── gradle
│   │   └── wrapper
│   │       └── gradle-wrapper.properties
│   ├── icon.png
│   ├── password.txt
│   ├── project_name.keystore
│   └── res
│       ├── drawable-hdpi
│       │   └── icon.png
│       ├── drawable-ldpi
│       │   └── icon.png
│       └── drawable-mdpi
│           └── icon.png
├── deploy
│   └── android
│       └── build
├── project.go
└── qml
    └── project.qml
```


The content of the `android` folder will be copied into
`deploy/android/build` before using `gradle` to bundle the binary inside
a `.apk`.

You can use the `AndroidManifest.xml` found inside
`deploy/android/build` as a template to build upon. (Set package
identifier, icon, permissions, ...)

For example: `android:icon="@drawable/icon"` attribute can be added to
`application` in `AndroidManifest.xml` to use the `icon.png`.

## Java native interface

You can call android's java runtime directly from go. See [java native interface example](../examples/androidextras/jni)


