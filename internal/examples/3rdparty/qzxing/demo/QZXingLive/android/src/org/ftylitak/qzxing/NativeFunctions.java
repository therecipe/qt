package org.ftylitak.qzxing;

public class NativeFunctions {
	static {
		System.loadLibrary("go");
	}
	public static native void onPermissionsGranted();
	public static native void onPermissionsDenied();
}
