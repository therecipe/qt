#include <android/log.h>
#include <dlfcn.h>
#include <jni.h>
#include <pthread.h>
#include "_cgo_export.h"

#define LOG_INFO(...) __android_log_print(ANDROID_LOG_INFO, "Go", __VA_ARGS__)
#define LOG_FATAL(...) __android_log_print(ANDROID_LOG_FATAL, "Go", __VA_ARGS__)

JNIEnv* globalEnv;
JavaVM* globalVm;

void* startMainMethod(void *data)
{
	void* mainLibHandler = dlopen((const char*)getGoLibPath(), 0);
	if (mainLibHandler == NULL) {
		LOG_FATAL("dlopen failed");
		return NULL;
	}

	uintptr_t mainPC = (uintptr_t)dlsym(mainLibHandler, "main.main");
	if (!mainPC) {
    LOG_FATAL("missing main.main");
	} else {
  	callMain(mainPC);
	}

	/*
	if (mainLibHandler) {
		int res = dlclose(mainLibHandler);
		if (res < 0)
			LOG_FATAL("dlclose failed");
	}

	//TODO: quit java app

	if (globalVm != NULL)
		(*globalVm)->DetachCurrentThread(globalEnv);
	*/

  return NULL;
}

void Java_org_qtproject_qt5_android_QtNative_startGoApplication(JNIEnv *env, jobject object, jstring paramsString, jstring environmentString)
{
  //Register java env and vm and context class
  globalEnv = env;
  (*globalEnv)->GetJavaVM(globalEnv, &globalVm);

  //Set params and env
  const char* nativeParamsString = (*globalEnv)->GetStringUTFChars(globalEnv, paramsString, (jboolean*)0);
  const char* nativeEnvironmentString = (*globalEnv)->GetStringUTFChars(globalEnv, environmentString, (jboolean*)0);
  setAndroidParamsAndEnv((char*)nativeParamsString, (char*)nativeEnvironmentString);
  (*globalEnv)->ReleaseStringUTFChars(globalEnv, paramsString, nativeParamsString);
	(*globalEnv)->ReleaseStringUTFChars(globalEnv, environmentString, nativeEnvironmentString);

  pthread_t appThread;
  pthread_create(&appThread, NULL, startMainMethod, NULL);
}
