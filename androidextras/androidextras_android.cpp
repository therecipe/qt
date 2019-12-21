// +build android android_emulator

#define protected public
#define private public

#include "androidextras_android.h"
#include "_cgo_export.h"

#include <QAndroidActivityResultReceiver>
#include <QAndroidBinder>
#include <QAndroidIntent>
#include <QAndroidJniEnvironment>
#include <QAndroidJniExceptionCleaner>
#include <QAndroidJniObject>
#include <QAndroidParcel>
#include <QAndroidService>
#include <QAndroidServiceConnection>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QtAndroid>

class MyQAndroidActivityResultReceiver: public QAndroidActivityResultReceiver
{
public:
	void handleActivityResult(int receiverRequestCode, int resultCode, const QAndroidJniObject & data) { callbackQAndroidActivityResultReceiver_HandleActivityResult(this, receiverRequestCode, resultCode, const_cast<QAndroidJniObject*>(&data)); };
};

Q_DECLARE_METATYPE(QAndroidActivityResultReceiver*)
Q_DECLARE_METATYPE(MyQAndroidActivityResultReceiver*)

int QAndroidActivityResultReceiver_QAndroidActivityResultReceiver_QRegisterMetaType(){qRegisterMetaType<QAndroidActivityResultReceiver*>(); return qRegisterMetaType<MyQAndroidActivityResultReceiver*>();}

void QAndroidActivityResultReceiver_HandleActivityResult(void* ptr, int receiverRequestCode, int resultCode, void* data)
{
	static_cast<QAndroidActivityResultReceiver*>(ptr)->handleActivityResult(receiverRequestCode, resultCode, *static_cast<QAndroidJniObject*>(data));
}

class MyQAndroidBinder: public QAndroidBinder
{
public:
	MyQAndroidBinder() : QAndroidBinder() {QAndroidBinder_QAndroidBinder_QRegisterMetaType();};
	MyQAndroidBinder(const QAndroidJniObject &binder) : QAndroidBinder(binder) {QAndroidBinder_QAndroidBinder_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QAndroidBinder*)
Q_DECLARE_METATYPE(MyQAndroidBinder*)

int QAndroidBinder_QAndroidBinder_QRegisterMetaType(){qRegisterMetaType<QAndroidBinder*>(); return qRegisterMetaType<MyQAndroidBinder*>();}

void* QAndroidBinder_NewQAndroidBinder()
{
	return new MyQAndroidBinder();
}

void* QAndroidBinder_NewQAndroidBinder2(void* binder)
{
	return new MyQAndroidBinder(*static_cast<QAndroidJniObject*>(binder));
}

void* QAndroidBinder_Handle(void* ptr)
{
	return new QAndroidJniObject(static_cast<QAndroidBinder*>(ptr)->handle().object());
}

char QAndroidBinder_Transact(void* ptr, int code, void* data, void* reply, long long flags)
{
	return static_cast<QAndroidBinder*>(ptr)->transact(code, *static_cast<QAndroidParcel*>(data), static_cast<QAndroidParcel*>(reply), static_cast<QAndroidBinder::CallType>(flags));
}

Q_DECLARE_METATYPE(QAndroidIntent)
Q_DECLARE_METATYPE(QAndroidIntent*)
void* QAndroidIntent_NewQAndroidIntent()
{
	return new QAndroidIntent();
}

void* QAndroidIntent_NewQAndroidIntent2(void* intent)
{
	return new QAndroidIntent(*static_cast<QAndroidJniObject*>(intent));
}

void* QAndroidIntent_NewQAndroidIntent3(struct QtAndroidExtras_PackedString action)
{
	return new QAndroidIntent(QString::fromUtf8(action.data, action.len));
}

void* QAndroidIntent_NewQAndroidIntent4(void* packageContext, char* className)
{
	return new QAndroidIntent(*static_cast<QAndroidJniObject*>(packageContext), const_cast<const char*>(className));
}

void* QAndroidIntent_ExtraBytes(void* ptr, struct QtAndroidExtras_PackedString key)
{
	return new QByteArray(static_cast<QAndroidIntent*>(ptr)->extraBytes(QString::fromUtf8(key.data, key.len)));
}

void* QAndroidIntent_ExtraVariant(void* ptr, struct QtAndroidExtras_PackedString key)
{
	return new QVariant(static_cast<QAndroidIntent*>(ptr)->extraVariant(QString::fromUtf8(key.data, key.len)));
}

void* QAndroidIntent_Handle(void* ptr)
{
	return new QAndroidJniObject(static_cast<QAndroidIntent*>(ptr)->handle().object());
}

void QAndroidIntent_PutExtra(void* ptr, struct QtAndroidExtras_PackedString key, void* data)
{
	static_cast<QAndroidIntent*>(ptr)->putExtra(QString::fromUtf8(key.data, key.len), *static_cast<QByteArray*>(data));
}

void QAndroidIntent_PutExtra2(void* ptr, struct QtAndroidExtras_PackedString key, void* value)
{
	static_cast<QAndroidIntent*>(ptr)->putExtra(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(value));
}

Q_DECLARE_METATYPE(QAndroidJniEnvironment*)
void* QAndroidJniEnvironment_NewQAndroidJniEnvironment()
{
	return new QAndroidJniEnvironment();
}

void* QAndroidJniEnvironment_FindClass(void* ptr, char* className)
{
	return static_cast<QAndroidJniEnvironment*>(ptr)->findClass(const_cast<const char*>(className));
}

void* QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM()
{
	return QAndroidJniEnvironment::javaVM();
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

Q_DECLARE_METATYPE(QAndroidJniExceptionCleaner*)
void* QAndroidJniExceptionCleaner_NewQAndroidJniExceptionCleaner(long long outputMode)
{
	return new QAndroidJniExceptionCleaner(static_cast<QAndroidJniExceptionCleaner::OutputMode>(outputMode));
}

void QAndroidJniExceptionCleaner_Clean(void* ptr)
{
	static_cast<QAndroidJniExceptionCleaner*>(ptr)->clean();
}

void QAndroidJniExceptionCleaner_DestroyQAndroidJniExceptionCleaner(void* ptr)
{
	static_cast<QAndroidJniExceptionCleaner*>(ptr)->~QAndroidJniExceptionCleaner();
}

Q_DECLARE_METATYPE(QAndroidJniObject)
Q_DECLARE_METATYPE(QAndroidJniObject*)
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

int QAndroidJniObject_CallMethodInt2(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_CallMethodInt2Caught(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_CallMethodBoolean2(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_CallMethodBoolean2Caught(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_CallMethodVoid2(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_CallMethodVoid2Caught(void* ptr, char* methodName, char* sig, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName), const_cast<const char*>(sig), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
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

char QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className)
{
	return QAndroidJniObject::isClassAvailable(const_cast<const char*>(className));
}

char QAndroidJniObject_IsValid(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->isValid();
}

void* QAndroidJniObject_Object(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->object();
}

void QAndroidJniObject_SetField(void* ptr, char* fieldName, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), static_cast<jobject>(value));
}

void QAndroidJniObject_SetField2(void* ptr, char* fieldName, char* signature, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_QAndroidJniObject_SetStaticField(char* className, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
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

void QAndroidJniObject_QAndroidJniObject_SetStaticField3(void* clazz, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
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

struct QtAndroidExtras_PackedString QAndroidJniObject_ToString(void* ptr)
{
	return ({ QByteArray* t3150b4 = new QByteArray(static_cast<QAndroidJniObject*>(ptr)->toString().toUtf8()); QtAndroidExtras_PackedString { const_cast<char*>(t3150b4->prepend("WHITESPACE").constData()+10), t3150b4->size()-10, t3150b4 }; });
}

void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr)
{
	static_cast<QAndroidJniObject*>(ptr)->~QAndroidJniObject();
}

Q_DECLARE_METATYPE(QAndroidParcel)
Q_DECLARE_METATYPE(QAndroidParcel*)
void* QAndroidParcel_NewQAndroidParcel()
{
	return new QAndroidParcel();
}

void* QAndroidParcel_NewQAndroidParcel2(void* parcel)
{
	return new QAndroidParcel(*static_cast<QAndroidJniObject*>(parcel));
}

void* QAndroidParcel_Handle(void* ptr)
{
	return new QAndroidJniObject(static_cast<QAndroidParcel*>(ptr)->handle().object());
}

void* QAndroidParcel_ReadBinder(void* ptr)
{
	return new QAndroidBinder(static_cast<QAndroidParcel*>(ptr)->readBinder());
}

void* QAndroidParcel_ReadData(void* ptr)
{
	return new QByteArray(static_cast<QAndroidParcel*>(ptr)->readData());
}

int QAndroidParcel_ReadFileDescriptor(void* ptr)
{
	return static_cast<QAndroidParcel*>(ptr)->readFileDescriptor();
}

void* QAndroidParcel_ReadVariant(void* ptr)
{
	return new QVariant(static_cast<QAndroidParcel*>(ptr)->readVariant());
}

void QAndroidParcel_WriteBinder(void* ptr, void* binder)
{
	static_cast<QAndroidParcel*>(ptr)->writeBinder(*static_cast<QAndroidBinder*>(binder));
}

void QAndroidParcel_WriteData(void* ptr, void* data)
{
	static_cast<QAndroidParcel*>(ptr)->writeData(*static_cast<QByteArray*>(data));
}

void QAndroidParcel_WriteFileDescriptor(void* ptr, int fd)
{
	static_cast<QAndroidParcel*>(ptr)->writeFileDescriptor(fd);
}

void QAndroidParcel_WriteVariant(void* ptr, void* value)
{
	static_cast<QAndroidParcel*>(ptr)->writeVariant(*static_cast<QVariant*>(value));
}

class MyQAndroidService: public QAndroidService
{
public:
	MyQAndroidService(int &argc, char **argv) : QAndroidService(argc, argv) {QAndroidService_QAndroidService_QRegisterMetaType();};
	QAndroidBinder * onBind(const QAndroidIntent & intent) { return static_cast<QAndroidBinder*>(callbackQAndroidService_OnBind(this, const_cast<QAndroidIntent*>(&intent))); };
	void Signal_AboutToQuit() { callbackQAndroidService_AboutToQuit(this); };
	void Signal_ApplicationNameChanged() { callbackQAndroidService_ApplicationNameChanged(this); };
	void Signal_ApplicationVersionChanged() { callbackQAndroidService_ApplicationVersionChanged(this); };
	bool event(QEvent * e) { return callbackQAndroidService_Event(this, e) != 0; };
	void Signal_OrganizationDomainChanged() { callbackQAndroidService_OrganizationDomainChanged(this); };
	void Signal_OrganizationNameChanged() { callbackQAndroidService_OrganizationNameChanged(this); };
	void quit() { callbackQAndroidService_Quit(this); };
	void childEvent(QChildEvent * event) { callbackQAndroidService_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAndroidService_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAndroidService_CustomEvent(this, event); };
	void deleteLater() { callbackQAndroidService_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAndroidService_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAndroidService_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAndroidService_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAndroidService_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtAndroidExtras_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQAndroidService_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAndroidService_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QAndroidService*)
Q_DECLARE_METATYPE(MyQAndroidService*)

int QAndroidService_QAndroidService_QRegisterMetaType(){qRegisterMetaType<QAndroidService*>(); return qRegisterMetaType<MyQAndroidService*>();}

void* QAndroidService_NewQAndroidService(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQAndroidService(argcs, argvs);
}

void* QAndroidService_OnBind(void* ptr, void* intent)
{
	return static_cast<QAndroidService*>(ptr)->onBind(*static_cast<QAndroidIntent*>(intent));
}

void* QAndroidService_OnBindDefault(void* ptr, void* intent)
{
		return static_cast<QAndroidService*>(ptr)->QAndroidService::onBind(*static_cast<QAndroidIntent*>(intent));
}

void* QAndroidService___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAndroidService___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAndroidService___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAndroidService___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAndroidService___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAndroidService___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAndroidService___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAndroidService___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAndroidService___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAndroidService___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAndroidService___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAndroidService___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

char QAndroidService_EventDefault(void* ptr, void* e)
{
		return static_cast<QAndroidService*>(ptr)->QAndroidService::event(static_cast<QEvent*>(e));
}

void QAndroidService_QuitDefault(void* ptr)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::quit();
}

void QAndroidService_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::childEvent(static_cast<QChildEvent*>(event));
}

void QAndroidService_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAndroidService_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::customEvent(static_cast<QEvent*>(event));
}

void QAndroidService_DeleteLaterDefault(void* ptr)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::deleteLater();
}

void QAndroidService_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAndroidService_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QAndroidService*>(ptr)->QAndroidService::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAndroidService_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QAndroidService*>(ptr)->QAndroidService::metaObject());
}

void QAndroidService_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QAndroidService*>(ptr)->QAndroidService::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQAndroidServiceConnection: public QAndroidServiceConnection
{
public:
	MyQAndroidServiceConnection() : QAndroidServiceConnection() {QAndroidServiceConnection_QAndroidServiceConnection_QRegisterMetaType();};
	MyQAndroidServiceConnection(const QAndroidJniObject &serviceConnection) : QAndroidServiceConnection(serviceConnection) {QAndroidServiceConnection_QAndroidServiceConnection_QRegisterMetaType();};
	void onServiceConnected(const QString & name, const QAndroidBinder & serviceBinder) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtAndroidExtras_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQAndroidServiceConnection_OnServiceConnected(this, namePacked, const_cast<QAndroidBinder*>(&serviceBinder)); };
	void onServiceDisconnected(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtAndroidExtras_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQAndroidServiceConnection_OnServiceDisconnected(this, namePacked); };
};

Q_DECLARE_METATYPE(QAndroidServiceConnection*)
Q_DECLARE_METATYPE(MyQAndroidServiceConnection*)

int QAndroidServiceConnection_QAndroidServiceConnection_QRegisterMetaType(){qRegisterMetaType<QAndroidServiceConnection*>(); return qRegisterMetaType<MyQAndroidServiceConnection*>();}

void* QAndroidServiceConnection_NewQAndroidServiceConnection()
{
	return new MyQAndroidServiceConnection();
}

void* QAndroidServiceConnection_NewQAndroidServiceConnection2(void* serviceConnection)
{
	return new MyQAndroidServiceConnection(*static_cast<QAndroidJniObject*>(serviceConnection));
}

void* QAndroidServiceConnection_Handle(void* ptr)
{
	return new QAndroidJniObject(static_cast<QAndroidServiceConnection*>(ptr)->handle().object());
}

void QAndroidServiceConnection_OnServiceConnected(void* ptr, struct QtAndroidExtras_PackedString name, void* serviceBinder)
{
	static_cast<QAndroidServiceConnection*>(ptr)->onServiceConnected(QString::fromUtf8(name.data, name.len), *static_cast<QAndroidBinder*>(serviceBinder));
}

void QAndroidServiceConnection_OnServiceDisconnected(void* ptr, struct QtAndroidExtras_PackedString name)
{
	static_cast<QAndroidServiceConnection*>(ptr)->onServiceDisconnected(QString::fromUtf8(name.data, name.len));
}

void* QtAndroid_QtAndroid_AndroidActivity()
{
	return new QAndroidJniObject(QtAndroid::androidActivity().object());
}

void* QtAndroid_QtAndroid_AndroidContext()
{
	return new QAndroidJniObject(QtAndroid::androidContext().object());
}

int QtAndroid_QtAndroid_AndroidSdkVersion()
{
	return QtAndroid::androidSdkVersion();
}

void* QtAndroid_QtAndroid_AndroidService()
{
	return new QAndroidJniObject(QtAndroid::androidService().object());
}

char QtAndroid_QtAndroid_BindService(void* serviceIntent, void* serviceConnection, long long flags)
{
	return QtAndroid::bindService(*static_cast<QAndroidIntent*>(serviceIntent), *static_cast<QAndroidServiceConnection*>(serviceConnection), static_cast<QtAndroid::BindFlag>(flags));
}

void QtAndroid_QtAndroid_HideSplashScreen()
{
	QtAndroid::hideSplashScreen();
}

void QtAndroid_QtAndroid_HideSplashScreen2(int duration)
{
	QtAndroid::hideSplashScreen(duration);
}

char QtAndroid_QtAndroid_ShouldShowRequestPermissionRationale(struct QtAndroidExtras_PackedString permission)
{
	return QtAndroid::shouldShowRequestPermissionRationale(QString::fromUtf8(permission.data, permission.len));
}

void QtAndroid_QtAndroid_StartActivity(void* intent, int receiverRequestCode, void* resultReceiver)
{
	QtAndroid::startActivity(*static_cast<QAndroidJniObject*>(intent), receiverRequestCode, static_cast<QAndroidActivityResultReceiver*>(resultReceiver));
}

void QtAndroid_QtAndroid_StartActivity2(void* intent, int receiverRequestCode, void* resultReceiver)
{
	QtAndroid::startActivity(*static_cast<QAndroidIntent*>(intent), receiverRequestCode, static_cast<QAndroidActivityResultReceiver*>(resultReceiver));
}

void QtAndroid_QtAndroid_StartIntentSender(void* intentSender, int receiverRequestCode, void* resultReceiver)
{
	QtAndroid::startIntentSender(*static_cast<QAndroidJniObject*>(intentSender), receiverRequestCode, static_cast<QAndroidActivityResultReceiver*>(resultReceiver));
}

