#include "qandroidactivityresultreceiver_android.h"
#include <QAndroidJniObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAndroidActivityResultReceiver>
#include "_cgo_export.h"

class MyQAndroidActivityResultReceiver: public QAndroidActivityResultReceiver {
public:
};

void QAndroidActivityResultReceiver_HandleActivityResult(void* ptr, int receiverRequestCode, int resultCode, void* data){
	static_cast<QAndroidActivityResultReceiver*>(ptr)->handleActivityResult(receiverRequestCode, resultCode, *static_cast<QAndroidJniObject*>(data));
}

#include "qandroidjniobject_android.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAndroidJniObject>
#include "_cgo_export.h"

class MyQAndroidJniObject: public QAndroidJniObject {
public:
};

void* QAndroidJniObject_NewQAndroidJniObject(){
	return new QAndroidJniObject();
}

void* QAndroidJniObject_NewQAndroidJniObject2(char* className){
	return new QAndroidJniObject(const_cast<const char*>(className));
}

void* QAndroidJniObject_NewQAndroidJniObject4(void* clazz){
	return new QAndroidJniObject(static_cast<jclass>(clazz));
}

void* QAndroidJniObject_NewQAndroidJniObject6(void* object){
	return new QAndroidJniObject(static_cast<jobject>(object));
}

int QAndroidJniObject_CallMethodInt(void* ptr, char* methodName){
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName));
}
int QAndroidJniObject_CallMethodBoolean(void* ptr, char* methodName){
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName));
}
void QAndroidJniObject_CallMethodVoid(void* ptr, char* methodName){
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName));
}


void* QAndroidJniObject_CallObjectMethod(void* ptr, char* methodName){
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(char* className, char* methodName){
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(char* className, char* methodName){
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(char* className, char* methodName){
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}


int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(void* clazz, char* methodName){
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(void* clazz, char* methodName){
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(void* clazz, char* methodName){
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}


void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(char* className, char* methodName){
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(void* clazz, char* methodName){
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_FromString(char* stri){
	return new QAndroidJniObject(QAndroidJniObject::fromString(QString(stri)).object());
}

int QAndroidJniObject_GetFieldInt(void* ptr, char* fieldName){
	return static_cast<QAndroidJniObject*>(ptr)->getField<jint>(const_cast<const char*>(fieldName));
}
int QAndroidJniObject_GetFieldBoolean(void* ptr, char* fieldName){
	return static_cast<QAndroidJniObject*>(ptr)->getField<jboolean>(const_cast<const char*>(fieldName));
}


void* QAndroidJniObject_GetObjectField(void* ptr, char* fieldName){
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetObjectField2(void* ptr, char* fieldName, char* signature){
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(char* className, char* fieldName){
	return QAndroidJniObject::getStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(char* className, char* fieldName){
	return QAndroidJniObject::getStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}


int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(void* clazz, char* fieldName){
	return QAndroidJniObject::getStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(void* clazz, char* fieldName){
	return QAndroidJniObject::getStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}


void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(char* className, char* fieldName){
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(char* className, char* fieldName, char* signature){
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(void* clazz, char* fieldName){
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(void* clazz, char* fieldName, char* signature){
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

int QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className){
	return QAndroidJniObject::isClassAvailable(const_cast<const char*>(className));
}

int QAndroidJniObject_IsValid(void* ptr){
	return static_cast<QAndroidJniObject*>(ptr)->isValid();
}

char* QAndroidJniObject_ToString(void* ptr){
	return static_cast<QAndroidJniObject*>(ptr)->toString().toUtf8().data();
}

void* QAndroidJniObject_Object(void* ptr){
	return static_cast<QAndroidJniObject*>(ptr)->object();
}

void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr){
	static_cast<QAndroidJniObject*>(ptr)->~QAndroidJniObject();
}

#include "qandroidjnienvironment_android.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAndroidJniEnvironment>
#include "_cgo_export.h"

class MyQAndroidJniEnvironment: public QAndroidJniEnvironment {
public:
};

void* QAndroidJniEnvironment_NewQAndroidJniEnvironment(){
	return new QAndroidJniEnvironment();
}

void* QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM(){
	return QAndroidJniEnvironment::javaVM();
}

void QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(void* ptr){
	static_cast<QAndroidJniEnvironment*>(ptr)->~QAndroidJniEnvironment();
}

