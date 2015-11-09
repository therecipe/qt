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

  //Start main
  uintptr_t mainPC = (uintptr_t)dlsym(RTLD_DEFAULT, "main.main");

  if (!mainPC)
    LOG_FATAL("missing main.main");
  else
    callMain(mainPC);


  //Close lib handler


  //Quit java app


  //Detach java thread


  return NULL;
}

void Java_org_qtproject_qt5_android_QtNative_startGoApplication(JNIEnv *env, jobject object, jstring paramsString, jstring environmentString)
{

  //Register java env and vm and context class
  globalEnv = env;
  (*globalEnv)->GetJavaVM(globalEnv, &globalVm);


  //Set env
  const char* nativeParamsString = (*env)->GetStringUTFChars(env, paramsString, (jboolean *)0);
  const char* nativeEnvString = (*env)->GetStringUTFChars(env, environmentString, (jboolean *)0);
  setAndroidParamsAndEnv((char*)nativeParamsString, (char*)nativeEnvString);


  //Get libHandler and main from params
  

  pthread_t appThread;
  pthread_create(&appThread, NULL, startMainMethod, NULL);
}
