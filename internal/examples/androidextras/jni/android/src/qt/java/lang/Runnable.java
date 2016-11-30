package qt.java.lang;

public class Runnable implements java.lang.Runnable {
	static {
		System.loadLibrary("go");
	}
	public String ObjectName;

	public static native void qtrun(qt.java.lang.Runnable me);
	@Override
	public void run() {
		qtrun(this);
	}
}
