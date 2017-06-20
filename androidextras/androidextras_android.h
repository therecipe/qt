// +build android android_emulator

#pragma once

#ifndef GO_QTANDROIDEXTRAS_H
#define GO_QTANDROIDEXTRAS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtAndroidExtras_PackedString { char* data; long long len; };
struct QtAndroidExtras_PackedList { void* data; long long len; };
void QAndroidActivityResultReceiver_HandleActivityResult(void* ptr, int receiverRequestCode, int resultCode, void* data);
void* QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM();
void* QAndroidJniEnvironment_NewQAndroidJniEnvironment();
void QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(void* ptr);
char QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionCheck();
void QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionDescribe();
void QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionClear();
void* QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionOccurred();
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(char* className, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethodCaught(char* className, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(char* className, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodStringCaught(char* className, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(void* clazz, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3Caught(void* clazz, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(void* clazz, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3Caught(void* clazz, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_FromLocalRef(void* localRef);
void* QAndroidJniObject_QAndroidJniObject_FromString(struct QtAndroidExtras_PackedString stri);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(char* className, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectFieldCaught(char* className, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(char* className, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldStringCaught(char* className, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(char* className, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2Caught(char* className, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(char* className, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2Caught(char* className, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3Caught(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3Caught(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(void* clazz, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4Caught(void* clazz, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(void* clazz, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4Caught(void* clazz, char* fieldName, char* signature);
void* QAndroidJniObject_NewQAndroidJniObject();
void* QAndroidJniObject_NewQAndroidJniObject2(char* className);
void* QAndroidJniObject_NewQAndroidJniObject3(char* className, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_NewQAndroidJniObject4(void* clazz);
void* QAndroidJniObject_NewQAndroidJniObject5(void* clazz, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_NewQAndroidJniObject6(void* object);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(char* className, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodIntCaught(char* className, char* methodName);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(char* className, char* methodName);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBooleanCaught(char* className, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(char* className, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoidCaught(char* className, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(void* clazz, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3Caught(void* clazz, char* methodName);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(void* clazz, char* methodName);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3Caught(void* clazz, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(void* clazz, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3Caught(void* clazz, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(char* className, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldIntCaught(char* className, char* fieldName);
char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(char* className, char* fieldName);
char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBooleanCaught(char* className, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(void* clazz, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2Caught(void* clazz, char* fieldName);
char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(void* clazz, char* fieldName);
char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2Caught(void* clazz, char* fieldName);
char QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className);
void QAndroidJniObject_SetField(void* ptr, char* fieldName, void* value);
void QAndroidJniObject_SetField2(void* ptr, char* fieldName, char* signature, void* value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(char* className, char* fieldName, int value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2Caught(char* className, char* fieldName, int value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(char* className, char* fieldName, char value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2Caught(char* className, char* fieldName, char value);
void QAndroidJniObject_QAndroidJniObject_SetStaticField(char* className, char* fieldName, char* signature, void* value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(void* clazz, char* fieldName, int value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4Caught(void* clazz, char* fieldName, int value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(void* clazz, char* fieldName, char value);
void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4Caught(void* clazz, char* fieldName, char value);
void QAndroidJniObject_QAndroidJniObject_SetStaticField3(void* clazz, char* fieldName, char* signature, void* value);
void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr);
void* QAndroidJniObject_CallObjectMethod(void* ptr, char* methodName);
void* QAndroidJniObject_CallObjectMethodCaught(void* ptr, char* methodName);
void* QAndroidJniObject_CallMethodString(void* ptr, char* methodName);
void* QAndroidJniObject_CallMethodStringCaught(void* ptr, char* methodName);
void* QAndroidJniObject_CallObjectMethod2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_CallObjectMethod2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_CallMethodString2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_CallMethodString2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_GetObjectField(void* ptr, char* fieldName);
void* QAndroidJniObject_GetObjectFieldCaught(void* ptr, char* fieldName);
void* QAndroidJniObject_GetFieldString(void* ptr, char* fieldName);
void* QAndroidJniObject_GetFieldStringCaught(void* ptr, char* fieldName);
void* QAndroidJniObject_GetObjectField2(void* ptr, char* fieldName, char* signature);
void* QAndroidJniObject_GetObjectField2Caught(void* ptr, char* fieldName, char* signature);
void* QAndroidJniObject_GetFieldString2(void* ptr, char* fieldName, char* signature);
void* QAndroidJniObject_GetFieldString2Caught(void* ptr, char* fieldName, char* signature);
struct QtAndroidExtras_PackedString QAndroidJniObject_ToString(void* ptr);
int QAndroidJniObject_CallMethodInt(void* ptr, char* methodName);
int QAndroidJniObject_CallMethodIntCaught(void* ptr, char* methodName);
char QAndroidJniObject_CallMethodBoolean(void* ptr, char* methodName);
char QAndroidJniObject_CallMethodBooleanCaught(void* ptr, char* methodName);
void QAndroidJniObject_CallMethodVoid(void* ptr, char* methodName);
void QAndroidJniObject_CallMethodVoidCaught(void* ptr, char* methodName);
int QAndroidJniObject_CallMethodInt2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_CallMethodInt2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_CallMethodBoolean2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
char QAndroidJniObject_CallMethodBoolean2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_CallMethodVoid2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_CallMethodVoid2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_GetFieldInt(void* ptr, char* fieldName);
int QAndroidJniObject_GetFieldIntCaught(void* ptr, char* fieldName);
char QAndroidJniObject_GetFieldBoolean(void* ptr, char* fieldName);
char QAndroidJniObject_GetFieldBooleanCaught(void* ptr, char* fieldName);
void* QAndroidJniObject_Object(void* ptr);
char QAndroidJniObject_IsValid(void* ptr);
void* QtAndroid_QtAndroid_AndroidActivity();
void* QtAndroid_QtAndroid_AndroidContext();
void* QtAndroid_QtAndroid_AndroidService();
int QtAndroid_QtAndroid_AndroidSdkVersion();
void QtAndroid_QtAndroid_HideSplashScreen();
void QtAndroid_QtAndroid_StartActivity(void* intent, int receiverRequestCode, void* resultReceiver);
void QtAndroid_QtAndroid_StartIntentSender(void* intentSender, int receiverRequestCode, void* resultReceiver);

#ifdef __cplusplus
}
#endif

#endif