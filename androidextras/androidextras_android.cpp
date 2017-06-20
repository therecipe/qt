// +build android android_emulator

#define protected public
#define private public

#include "androidextras_android.h"
#include "_cgo_export.h"

#include <QAndroidActivityResultReceiver>
#include <QAndroidJniEnvironment>
#include <QAndroidJniObject>
#include <QByteArray>
#include <QString>
#include <QtAndroid>

class MyQAndroidActivityResultReceiver: public QAndroidActivityResultReceiver
{
public:
	void handleActivityResult(int receiverRequestCode, int resultCode, const QAndroidJniObject & data) { callbackQAndroidActivityResultReceiver_HandleActivityResult(this, receiverRequestCode, resultCode, const_cast<QAndroidJniObject*>(&data)); };
};

void QAndroidActivityResultReceiver_HandleActivityResult(void* ptr, int receiverRequestCode, int resultCode, void* data)
{
	static_cast<QAndroidActivityResultReceiver*>(ptr)->handleActivityResult(receiverRequestCode, resultCode, *static_cast<QAndroidJniObject*>(data));
}

void* QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM()
{
	return QAndroidJniEnvironment::javaVM();
}

void* QAndroidJniEnvironment_NewQAndroidJniEnvironment()
{
	return new QAndroidJniEnvironment();
}

void QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(void* ptr)
{
	static_cast<QAndroidJniEnvironment*>(ptr)->~QAndroidJniEnvironment();
}

char QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionCheck()
{
	return ({ QAndroidJniEnvironment env; env->ExceptionCheck(); });
}

void QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionDescribe()
{
	({ QAndroidJniEnvironment env; env->ExceptionDescribe(); });
}

void QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionClear()
{
	({ QAndroidJniEnvironment env; env->ExceptionClear(); });
}

void* QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionOccurred()
{
	return ({ QAndroidJniEnvironment env; env->ExceptionOccurred(); });
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethodCaught(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodStringCaught(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3Caught(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3Caught(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_FromLocalRef(void* localRef)
{
	return new QAndroidJniObject(QAndroidJniObject::fromLocalRef(static_cast<jobject>(localRef)).object());
}

void* QAndroidJniObject_QAndroidJniObject_FromString(struct QtAndroidExtras_PackedString stri)
{
	return new QAndroidJniObject(QAndroidJniObject::fromString(QString::fromUtf8(stri.data, stri.len)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectFieldCaught(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldStringCaught(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2Caught(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2Caught(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3Caught(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3Caught(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4Caught(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4Caught(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_NewQAndroidJniObject()
{
	return new QAndroidJniObject();
}

void* QAndroidJniObject_NewQAndroidJniObject2(char* className)
{
	return new QAndroidJniObject(const_cast<const char*>(className));
}

void* QAndroidJniObject_NewQAndroidJniObject3(char* className, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(const_cast<const char*>(className), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_NewQAndroidJniObject4(void* clazz)
{
	return new QAndroidJniObject(static_cast<jclass>(clazz));
}

void* QAndroidJniObject_NewQAndroidJniObject5(void* clazz, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<jclass>(clazz), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_NewQAndroidJniObject6(void* object)
{
	return new QAndroidJniObject(static_cast<jobject>(object));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodIntCaught(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBooleanCaught(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(char* className, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoidCaught(char* className, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2Caught(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3Caught(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3Caught(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(void* clazz, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3Caught(void* clazz, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4Caught(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldIntCaught(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBooleanCaught(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2Caught(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2Caught(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className)
{
	return QAndroidJniObject::isClassAvailable(const_cast<const char*>(className));
}

void QAndroidJniObject_SetField(void* ptr, char* fieldName, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), static_cast<jobject>(value));
}

void QAndroidJniObject_SetField2(void* ptr, char* fieldName, char* signature, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(char* className, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2Caught(char* className, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(char* className, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2Caught(char* className, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticField(char* className, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(void* clazz, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4Caught(void* clazz, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(void* clazz, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4Caught(void* clazz, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticField3(void* clazz, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr)
{
	static_cast<QAndroidJniObject*>(ptr)->~QAndroidJniObject();
}

void* QAndroidJniObject_CallObjectMethod(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallObjectMethodCaught(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallMethodString(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallMethodStringCaught(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallObjectMethod2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_CallObjectMethod2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_CallMethodString2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_CallMethodString2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_GetObjectField(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetObjectFieldCaught(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetFieldString(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetFieldStringCaught(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetObjectField2(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_GetObjectField2Caught(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_GetFieldString2(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_GetFieldString2Caught(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

struct QtAndroidExtras_PackedString QAndroidJniObject_ToString(void* ptr)
{
	return ({ QByteArray t3150b4 = static_cast<QAndroidJniObject*>(ptr)->toString().toUtf8(); QtAndroidExtras_PackedString { const_cast<char*>(t3150b4.prepend("WHITESPACE").constData()+10), t3150b4.size()-10 }; });
}

int QAndroidJniObject_CallMethodInt(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName));
}

int QAndroidJniObject_CallMethodIntCaught(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName));
}

char QAndroidJniObject_CallMethodBoolean(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName));
}

char QAndroidJniObject_CallMethodBooleanCaught(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName));
}

void QAndroidJniObject_CallMethodVoid(void* ptr, char* methodName)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName));
}

void QAndroidJniObject_CallMethodVoidCaught(void* ptr, char* methodName)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName));
}

int QAndroidJniObject_CallMethodInt2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_CallMethodInt2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_CallMethodBoolean2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_CallMethodBoolean2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_CallMethodVoid2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_CallMethodVoid2Caught(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_GetFieldInt(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jint>(const_cast<const char*>(fieldName));
}

int QAndroidJniObject_GetFieldIntCaught(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jint>(const_cast<const char*>(fieldName));
}

char QAndroidJniObject_GetFieldBoolean(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jboolean>(const_cast<const char*>(fieldName));
}

char QAndroidJniObject_GetFieldBooleanCaught(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jboolean>(const_cast<const char*>(fieldName));
}

void* QAndroidJniObject_Object(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->object();
}

char QAndroidJniObject_IsValid(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->isValid();
}

void* QtAndroid_QtAndroid_AndroidActivity()
{
	return new QAndroidJniObject(QtAndroid::androidActivity().object());
}

void* QtAndroid_QtAndroid_AndroidContext()
{
	return new QAndroidJniObject(QtAndroid::androidContext().object());
}

void* QtAndroid_QtAndroid_AndroidService()
{
	return new QAndroidJniObject(QtAndroid::androidService().object());
}

int QtAndroid_QtAndroid_AndroidSdkVersion()
{
	return QtAndroid::androidSdkVersion();
}

void QtAndroid_QtAndroid_HideSplashScreen()
{
	QtAndroid::hideSplashScreen();
}

void QtAndroid_QtAndroid_StartActivity(void* intent, int receiverRequestCode, void* resultReceiver)
{
	QtAndroid::startActivity(*static_cast<QAndroidJniObject*>(intent), receiverRequestCode, static_cast<QAndroidActivityResultReceiver*>(resultReceiver));
}

void QtAndroid_QtAndroid_StartIntentSender(void* intentSender, int receiverRequestCode, void* resultReceiver)
{
	QtAndroid::startIntentSender(*static_cast<QAndroidJniObject*>(intentSender), receiverRequestCode, static_cast<QAndroidActivityResultReceiver*>(resultReceiver));
}

