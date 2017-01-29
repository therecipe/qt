// +build !minimal

#define protected public
#define private public

#include "scxml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QJSEngine>
#include <QList>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QScxmlCppDataModel>
#include <QScxmlDataModel>
#include <QScxmlEcmaScriptDataModel>
#include <QScxmlError>
#include <QScxmlEvent>
#include <QScxmlEventFilter>
#include <QScxmlNullDataModel>
#include <QScxmlParser>
#include <QScxmlStateMachine>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QXmlStreamReader>

class MyQScxmlCppDataModel: public QScxmlCppDataModel
{
public:
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlCppDataModel_HasScxmlProperty(const_cast<MyQScxmlCppDataModel*>(this), namePacked) != 0; };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlCppDataModel_ScxmlProperty(const_cast<MyQScxmlCppDataModel*>(this), namePacked)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlCppDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlCppDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	
	
	
	
	
	
	void timerEvent(QTimerEvent * event) { callbackQScxmlCppDataModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScxmlCppDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlCppDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlCppDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlCppDataModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlCppDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlCppDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlCppDataModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlCppDataModel_MetaObject(const_cast<MyQScxmlCppDataModel*>(this))); };
};

char QScxmlCppDataModel_In(void* ptr, char* stateName)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->In(QString(stateName));
}

char QScxmlCppDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

char QScxmlCppDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlCppDataModel_ScxmlEvent(void* ptr)
{
	return const_cast<QScxmlEvent*>(&static_cast<QScxmlCppDataModel*>(ptr)->scxmlEvent());
}

void* QScxmlCppDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlCppDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::scxmlProperty(QString(name)));
}

void QScxmlCppDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

char QScxmlCppDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlCppDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlCppDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlCppDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void* QScxmlCppDataModel___setup_initialDataValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlCppDataModel___setup_initialDataValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlCppDataModel___setup_initialDataValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlCppDataModel___setup_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedString QScxmlCppDataModel_____setup_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlCppDataModel_____setup_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlCppDataModel_____setup_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

void QScxmlCppDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlCppDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlCppDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlCppDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlCppDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlCppDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlCppDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlCppDataModel*>(ptr), "deleteLater");
}

void QScxmlCppDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::deleteLater();
}

void QScxmlCppDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScxmlCppDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScxmlCppDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::event(static_cast<QEvent*>(e));
}

char QScxmlCppDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScxmlCppDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlCppDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlCppDataModel*>(ptr)->metaObject());
}

void* QScxmlCppDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::metaObject());
}

class MyQScxmlDataModel: public QScxmlDataModel
{
public:
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlDataModel_HasScxmlProperty(const_cast<MyQScxmlDataModel*>(this), namePacked) != 0; };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlDataModel_ScxmlProperty(const_cast<MyQScxmlDataModel*>(this), namePacked)); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlDataModel_SetScxmlEvent(this, const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, stateMachine); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<MyQScxmlDataModel*>(this))); };
};

char QScxmlDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

void* QScxmlDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void QScxmlDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

char QScxmlDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

void QScxmlDataModel_SetStateMachine(void* ptr, void* stateMachine)
{
	static_cast<QScxmlDataModel*>(ptr)->setStateMachine(static_cast<QScxmlStateMachine*>(stateMachine));
}

char QScxmlDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void* QScxmlDataModel_StateMachine(void* ptr)
{
	return static_cast<QScxmlDataModel*>(ptr)->stateMachine();
}

void QScxmlDataModel_ConnectStateMachineChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlDataModel*>(ptr), static_cast<void (QScxmlDataModel::*)(QScxmlStateMachine *)>(&QScxmlDataModel::stateMachineChanged), static_cast<MyQScxmlDataModel*>(ptr), static_cast<void (MyQScxmlDataModel::*)(QScxmlStateMachine *)>(&MyQScxmlDataModel::Signal_StateMachineChanged));
}

void QScxmlDataModel_DisconnectStateMachineChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlDataModel*>(ptr), static_cast<void (QScxmlDataModel::*)(QScxmlStateMachine *)>(&QScxmlDataModel::stateMachineChanged), static_cast<MyQScxmlDataModel*>(ptr), static_cast<void (MyQScxmlDataModel::*)(QScxmlStateMachine *)>(&MyQScxmlDataModel::Signal_StateMachineChanged));
}

void QScxmlDataModel_StateMachineChanged(void* ptr, void* stateMachine)
{
	static_cast<QScxmlDataModel*>(ptr)->stateMachineChanged(static_cast<QScxmlStateMachine*>(stateMachine));
}

void* QScxmlDataModel___setup_initialDataValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlDataModel___setup_initialDataValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlDataModel___setup_initialDataValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlDataModel___setup_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedString QScxmlDataModel_____setup_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlDataModel_____setup_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlDataModel_____setup_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

void* QScxmlDataModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlDataModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlDataModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QScxmlDataModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlDataModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlDataModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QScxmlDataModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlDataModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlDataModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlDataModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlDataModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlDataModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlDataModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlDataModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlDataModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QScxmlDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlDataModel*>(ptr), "deleteLater");
}

void QScxmlDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::deleteLater();
}

void QScxmlDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScxmlDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScxmlDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::event(static_cast<QEvent*>(e));
}

char QScxmlDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScxmlDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlDataModel*>(ptr)->metaObject());
}

void* QScxmlDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::metaObject());
}

class MyQScxmlEcmaScriptDataModel: public QScxmlEcmaScriptDataModel
{
public:
	MyQScxmlEcmaScriptDataModel(QObject *parent) : QScxmlEcmaScriptDataModel(parent) {};
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(const_cast<MyQScxmlEcmaScriptDataModel*>(this), namePacked) != 0; };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlEcmaScriptDataModel_ScxmlProperty(const_cast<MyQScxmlEcmaScriptDataModel*>(this), namePacked)); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlEcmaScriptDataModel_SetScxmlEvent(this, const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlEcmaScriptDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQScxmlEcmaScriptDataModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScxmlEcmaScriptDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlEcmaScriptDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlEcmaScriptDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlEcmaScriptDataModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlEcmaScriptDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlEcmaScriptDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlEcmaScriptDataModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlEcmaScriptDataModel_MetaObject(const_cast<MyQScxmlEcmaScriptDataModel*>(this))); };
};

void* QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(void* parent)
{
	return new MyQScxmlEcmaScriptDataModel(static_cast<QObject*>(parent));
}

void* QScxmlEcmaScriptDataModel_Engine(void* ptr)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->engine();
}

char QScxmlEcmaScriptDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

char QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlEcmaScriptDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlEcmaScriptDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::scxmlProperty(QString(name)));
}

void QScxmlEcmaScriptDataModel_SetEngine(void* ptr, void* engine)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setEngine(static_cast<QJSEngine*>(engine));
}

void QScxmlEcmaScriptDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlEcmaScriptDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

char QScxmlEcmaScriptDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlEcmaScriptDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlEcmaScriptDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlEcmaScriptDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void* QScxmlEcmaScriptDataModel___setup_initialDataValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlEcmaScriptDataModel___setup_initialDataValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlEcmaScriptDataModel___setup_initialDataValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlEcmaScriptDataModel___setup_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedString QScxmlEcmaScriptDataModel_____setup_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlEcmaScriptDataModel_____setup_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlEcmaScriptDataModel_____setup_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

void* QScxmlEcmaScriptDataModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlEcmaScriptDataModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlEcmaScriptDataModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QScxmlEcmaScriptDataModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlEcmaScriptDataModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlEcmaScriptDataModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QScxmlEcmaScriptDataModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlEcmaScriptDataModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlEcmaScriptDataModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlEcmaScriptDataModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlEcmaScriptDataModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlEcmaScriptDataModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlEcmaScriptDataModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlEcmaScriptDataModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlEcmaScriptDataModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QScxmlEcmaScriptDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlEcmaScriptDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlEcmaScriptDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlEcmaScriptDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlEcmaScriptDataModel*>(ptr), "deleteLater");
}

void QScxmlEcmaScriptDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::deleteLater();
}

void QScxmlEcmaScriptDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScxmlEcmaScriptDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScxmlEcmaScriptDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::event(static_cast<QEvent*>(e));
}

char QScxmlEcmaScriptDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScxmlEcmaScriptDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlEcmaScriptDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->metaObject());
}

void* QScxmlEcmaScriptDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::metaObject());
}

void* QScxmlError_NewQScxmlError()
{
	return new QScxmlError();
}

void* QScxmlError_NewQScxmlError3(void* other)
{
	return new QScxmlError(*static_cast<QScxmlError*>(other));
}

void* QScxmlError_NewQScxmlError2(char* fileName, int line, int column, char* description)
{
	return new QScxmlError(QString(fileName), line, column, QString(description));
}

int QScxmlError_Column(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->column();
}

struct QtScxml_PackedString QScxmlError_Description(void* ptr)
{
	return ({ QByteArray td8220c = static_cast<QScxmlError*>(ptr)->description().toUtf8(); QtScxml_PackedString { const_cast<char*>(td8220c.prepend("WHITESPACE").constData()+10), td8220c.size()-10 }; });
}

struct QtScxml_PackedString QScxmlError_FileName(void* ptr)
{
	return ({ QByteArray t71e3d6 = static_cast<QScxmlError*>(ptr)->fileName().toUtf8(); QtScxml_PackedString { const_cast<char*>(t71e3d6.prepend("WHITESPACE").constData()+10), t71e3d6.size()-10 }; });
}

char QScxmlError_IsValid(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->isValid();
}

int QScxmlError_Line(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->line();
}

struct QtScxml_PackedString QScxmlError_ToString(void* ptr)
{
	return ({ QByteArray t6a9dc4 = static_cast<QScxmlError*>(ptr)->toString().toUtf8(); QtScxml_PackedString { const_cast<char*>(t6a9dc4.prepend("WHITESPACE").constData()+10), t6a9dc4.size()-10 }; });
}

void QScxmlError_DestroyQScxmlError(void* ptr)
{
	static_cast<QScxmlError*>(ptr)->~QScxmlError();
}

void* QScxmlEvent_NewQScxmlEvent()
{
	return new QScxmlEvent();
}

void* QScxmlEvent_NewQScxmlEvent2(void* other)
{
	return new QScxmlEvent(*static_cast<QScxmlEvent*>(other));
}

void QScxmlEvent_Clear(void* ptr)
{
	static_cast<QScxmlEvent*>(ptr)->clear();
}

void* QScxmlEvent_Data(void* ptr)
{
	return new QVariant(static_cast<QScxmlEvent*>(ptr)->data());
}

int QScxmlEvent_Delay(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->delay();
}

struct QtScxml_PackedString QScxmlEvent_ErrorMessage(void* ptr)
{
	return ({ QByteArray tc788d7 = static_cast<QScxmlEvent*>(ptr)->errorMessage().toUtf8(); QtScxml_PackedString { const_cast<char*>(tc788d7.prepend("WHITESPACE").constData()+10), tc788d7.size()-10 }; });
}

long long QScxmlEvent_EventType(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->eventType();
}

struct QtScxml_PackedString QScxmlEvent_InvokeId(void* ptr)
{
	return ({ QByteArray t0fa333 = static_cast<QScxmlEvent*>(ptr)->invokeId().toUtf8(); QtScxml_PackedString { const_cast<char*>(t0fa333.prepend("WHITESPACE").constData()+10), t0fa333.size()-10 }; });
}

char QScxmlEvent_IsErrorEvent(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->isErrorEvent();
}

struct QtScxml_PackedString QScxmlEvent_Name(void* ptr)
{
	return ({ QByteArray t0b09c8 = static_cast<QScxmlEvent*>(ptr)->name().toUtf8(); QtScxml_PackedString { const_cast<char*>(t0b09c8.prepend("WHITESPACE").constData()+10), t0b09c8.size()-10 }; });
}

struct QtScxml_PackedString QScxmlEvent_Origin(void* ptr)
{
	return ({ QByteArray tb77f0c = static_cast<QScxmlEvent*>(ptr)->origin().toUtf8(); QtScxml_PackedString { const_cast<char*>(tb77f0c.prepend("WHITESPACE").constData()+10), tb77f0c.size()-10 }; });
}

struct QtScxml_PackedString QScxmlEvent_OriginType(void* ptr)
{
	return ({ QByteArray t2089d9 = static_cast<QScxmlEvent*>(ptr)->originType().toUtf8(); QtScxml_PackedString { const_cast<char*>(t2089d9.prepend("WHITESPACE").constData()+10), t2089d9.size()-10 }; });
}

struct QtScxml_PackedString QScxmlEvent_ScxmlType(void* ptr)
{
	return ({ QByteArray tf8852d = static_cast<QScxmlEvent*>(ptr)->scxmlType().toUtf8(); QtScxml_PackedString { const_cast<char*>(tf8852d.prepend("WHITESPACE").constData()+10), tf8852d.size()-10 }; });
}

struct QtScxml_PackedString QScxmlEvent_SendId(void* ptr)
{
	return ({ QByteArray tc83807 = static_cast<QScxmlEvent*>(ptr)->sendId().toUtf8(); QtScxml_PackedString { const_cast<char*>(tc83807.prepend("WHITESPACE").constData()+10), tc83807.size()-10 }; });
}

void QScxmlEvent_SetData(void* ptr, void* data)
{
	static_cast<QScxmlEvent*>(ptr)->setData(*static_cast<QVariant*>(data));
}

void QScxmlEvent_SetDelay(void* ptr, int delayInMiliSecs)
{
	static_cast<QScxmlEvent*>(ptr)->setDelay(delayInMiliSecs);
}

void QScxmlEvent_SetErrorMessage(void* ptr, char* message)
{
	static_cast<QScxmlEvent*>(ptr)->setErrorMessage(QString(message));
}

void QScxmlEvent_SetEventType(void* ptr, long long ty)
{
	static_cast<QScxmlEvent*>(ptr)->setEventType(static_cast<QScxmlEvent::EventType>(ty));
}

void QScxmlEvent_SetInvokeId(void* ptr, char* invokeid)
{
	static_cast<QScxmlEvent*>(ptr)->setInvokeId(QString(invokeid));
}

void QScxmlEvent_SetName(void* ptr, char* name)
{
	static_cast<QScxmlEvent*>(ptr)->setName(QString(name));
}

void QScxmlEvent_SetOrigin(void* ptr, char* origin)
{
	static_cast<QScxmlEvent*>(ptr)->setOrigin(QString(origin));
}

void QScxmlEvent_SetOriginType(void* ptr, char* origintype)
{
	static_cast<QScxmlEvent*>(ptr)->setOriginType(QString(origintype));
}

void QScxmlEvent_SetSendId(void* ptr, char* sendid)
{
	static_cast<QScxmlEvent*>(ptr)->setSendId(QString(sendid));
}

void QScxmlEvent_DestroyQScxmlEvent(void* ptr)
{
	static_cast<QScxmlEvent*>(ptr)->~QScxmlEvent();
}

class MyQScxmlEventFilter: public QScxmlEventFilter
{
public:
	bool handle(QScxmlEvent * event, QScxmlStateMachine * stateMachine) { return callbackQScxmlEventFilter_Handle(this, event, stateMachine) != 0; };
	 ~MyQScxmlEventFilter() { callbackQScxmlEventFilter_DestroyQScxmlEventFilter(this); };
};

char QScxmlEventFilter_Handle(void* ptr, void* event, void* stateMachine)
{
	return static_cast<QScxmlEventFilter*>(ptr)->handle(static_cast<QScxmlEvent*>(event), static_cast<QScxmlStateMachine*>(stateMachine));
}

void QScxmlEventFilter_DestroyQScxmlEventFilter(void* ptr)
{
	static_cast<QScxmlEventFilter*>(ptr)->~QScxmlEventFilter();
}

void QScxmlEventFilter_DestroyQScxmlEventFilterDefault(void* ptr)
{

}

class MyQScxmlNullDataModel: public QScxmlNullDataModel
{
public:
	MyQScxmlNullDataModel(QObject *parent) : QScxmlNullDataModel(parent) {};
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlNullDataModel_HasScxmlProperty(const_cast<MyQScxmlNullDataModel*>(this), namePacked) != 0; };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlNullDataModel_ScxmlProperty(const_cast<MyQScxmlNullDataModel*>(this), namePacked)); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlNullDataModel_SetScxmlEvent(this, const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlNullDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlNullDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQScxmlNullDataModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScxmlNullDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlNullDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlNullDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlNullDataModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlNullDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlNullDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlNullDataModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlNullDataModel_MetaObject(const_cast<MyQScxmlNullDataModel*>(this))); };
};

void* QScxmlNullDataModel_NewQScxmlNullDataModel(void* parent)
{
	return new MyQScxmlNullDataModel(static_cast<QObject*>(parent));
}

char QScxmlNullDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

char QScxmlNullDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlNullDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlNullDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::scxmlProperty(QString(name)));
}

void QScxmlNullDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlNullDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

char QScxmlNullDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlNullDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

char QScxmlNullDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlNullDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void* QScxmlNullDataModel___setup_initialDataValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlNullDataModel___setup_initialDataValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlNullDataModel___setup_initialDataValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlNullDataModel___setup_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedString QScxmlNullDataModel_____setup_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlNullDataModel_____setup_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlNullDataModel_____setup_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

void* QScxmlNullDataModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlNullDataModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlNullDataModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QScxmlNullDataModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlNullDataModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlNullDataModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QScxmlNullDataModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlNullDataModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlNullDataModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlNullDataModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlNullDataModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlNullDataModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlNullDataModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlNullDataModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlNullDataModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QScxmlNullDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlNullDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlNullDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlNullDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlNullDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlNullDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlNullDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlNullDataModel*>(ptr), "deleteLater");
}

void QScxmlNullDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::deleteLater();
}

void QScxmlNullDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScxmlNullDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScxmlNullDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::event(static_cast<QEvent*>(e));
}

char QScxmlNullDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScxmlNullDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlNullDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlNullDataModel*>(ptr)->metaObject());
}

void* QScxmlNullDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::metaObject());
}

void* QScxmlParser_NewQScxmlParser(void* reader)
{
	return new QScxmlParser(static_cast<QXmlStreamReader*>(reader));
}

void QScxmlParser_AddError(void* ptr, char* msg)
{
	static_cast<QScxmlParser*>(ptr)->addError(QString(msg));
}

struct QtScxml_PackedList QScxmlParser_Errors(void* ptr)
{
	return ({ QVector<QScxmlError>* tmpValue = new QVector<QScxmlError>(static_cast<QScxmlParser*>(ptr)->errors()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedString QScxmlParser_FileName(void* ptr)
{
	return ({ QByteArray t0fecdd = static_cast<QScxmlParser*>(ptr)->fileName().toUtf8(); QtScxml_PackedString { const_cast<char*>(t0fecdd.prepend("WHITESPACE").constData()+10), t0fecdd.size()-10 }; });
}

void QScxmlParser_InstantiateDataModel(void* ptr, void* stateMachine)
{
	static_cast<QScxmlParser*>(ptr)->instantiateDataModel(static_cast<QScxmlStateMachine*>(stateMachine));
}

void* QScxmlParser_InstantiateStateMachine(void* ptr)
{
	return static_cast<QScxmlParser*>(ptr)->instantiateStateMachine();
}

void QScxmlParser_Parse(void* ptr)
{
	static_cast<QScxmlParser*>(ptr)->parse();
}

long long QScxmlParser_QtMode(void* ptr)
{
	return static_cast<QScxmlParser*>(ptr)->qtMode();
}

void QScxmlParser_SetFileName(void* ptr, char* fileName)
{
	static_cast<QScxmlParser*>(ptr)->setFileName(QString(fileName));
}

void QScxmlParser_SetQtMode(void* ptr, long long mode)
{
	static_cast<QScxmlParser*>(ptr)->setQtMode(static_cast<QScxmlParser::QtMode>(mode));
}

void QScxmlParser_DestroyQScxmlParser(void* ptr)
{
	static_cast<QScxmlParser*>(ptr)->~QScxmlParser();
}

void* QScxmlParser___errors_atList(void* ptr, int i)
{
	return new QScxmlError(static_cast<QVector<QScxmlError>*>(ptr)->at(i));
}

void QScxmlParser___errors_setList(void* ptr, void* i)
{
	static_cast<QVector<QScxmlError>*>(ptr)->append(*static_cast<QScxmlError*>(i));
}

void* QScxmlParser___errors_newList(void* ptr)
{
	return new QVector<QScxmlError>;
}

class MyQScxmlStateMachine: public QScxmlStateMachine
{
public:
	void Signal_DataModelChanged(QScxmlDataModel * model) { callbackQScxmlStateMachine_DataModelChanged(this, model); };
	void Signal_EventOccurred(const QScxmlEvent & event) { callbackQScxmlStateMachine_EventOccurred(this, const_cast<QScxmlEvent*>(&event)); };
	void Signal_ExternalEventOccurred(const QScxmlEvent & event) { callbackQScxmlStateMachine_ExternalEventOccurred(this, const_cast<QScxmlEvent*>(&event)); };
	void Signal_Finished() { callbackQScxmlStateMachine_Finished(this); };
	bool init() { return callbackQScxmlStateMachine_Init(this) != 0; };
	void Signal_InitializedChanged(bool initialized) { callbackQScxmlStateMachine_InitializedChanged(this, initialized); };
	void Signal_Log(const QString & label, const QString & msg) { QByteArray t64c653 = label.toUtf8(); QtScxml_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t19f34e = msg.toUtf8(); QtScxml_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQScxmlStateMachine_Log(this, labelPacked, msgPacked); };
	void Signal_ReachedStableState() { callbackQScxmlStateMachine_ReachedStableState(this); };
	void Signal_RunningChanged(bool running) { callbackQScxmlStateMachine_RunningChanged(this, running); };
	void start() { callbackQScxmlStateMachine_Start(this); };
	void stop() { callbackQScxmlStateMachine_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlStateMachine_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQScxmlStateMachine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlStateMachine_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlStateMachine_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlStateMachine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlStateMachine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlStateMachine_MetaObject(const_cast<MyQScxmlStateMachine*>(this))); };
};

struct QtScxml_PackedList QScxmlStateMachine_InitialValues(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QScxmlStateMachine*>(ptr)->initialValues()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

char QScxmlStateMachine_IsInitialized(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInitialized();
}

void QScxmlStateMachine_SetInitialValues(void* ptr, void* initialValues)
{
	static_cast<QScxmlStateMachine*>(ptr)->setInitialValues(*static_cast<QMap<QString, QVariant>*>(initialValues));
}

struct QtScxml_PackedString QScxmlStateMachine_ActiveStateNames(void* ptr, char compress)
{
	return ({ QByteArray t57b518 = static_cast<QScxmlStateMachine*>(ptr)->activeStateNames(compress != 0).join("|").toUtf8(); QtScxml_PackedString { const_cast<char*>(t57b518.prepend("WHITESPACE").constData()+10), t57b518.size()-10 }; });
}

void QScxmlStateMachine_CancelDelayedEvent(void* ptr, char* sendId)
{
	static_cast<QScxmlStateMachine*>(ptr)->cancelDelayedEvent(QString(sendId));
}

long long QScxmlStateMachine_DataBinding(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->dataBinding();
}

void* QScxmlStateMachine_DataModel(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->dataModel();
}

void QScxmlStateMachine_ConnectDataModelChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlDataModel *)>(&QScxmlStateMachine::dataModelChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlDataModel *)>(&MyQScxmlStateMachine::Signal_DataModelChanged));
}

void QScxmlStateMachine_DisconnectDataModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlDataModel *)>(&QScxmlStateMachine::dataModelChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlDataModel *)>(&MyQScxmlStateMachine::Signal_DataModelChanged));
}

void QScxmlStateMachine_DataModelChanged(void* ptr, void* model)
{
	static_cast<QScxmlStateMachine*>(ptr)->dataModelChanged(static_cast<QScxmlDataModel*>(model));
}

void QScxmlStateMachine_ConnectEventOccurred(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::eventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_EventOccurred));
}

void QScxmlStateMachine_DisconnectEventOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::eventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_EventOccurred));
}

void QScxmlStateMachine_EventOccurred(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->eventOccurred(*static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_ConnectExternalEventOccurred(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::externalEventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_ExternalEventOccurred));
}

void QScxmlStateMachine_DisconnectExternalEventOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::externalEventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_ExternalEventOccurred));
}

void QScxmlStateMachine_ExternalEventOccurred(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->externalEventOccurred(*static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::finished), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_Finished));
}

void QScxmlStateMachine_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::finished), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_Finished));
}

void QScxmlStateMachine_Finished(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->finished();
}

void* QScxmlStateMachine_QScxmlStateMachine_FromData(void* data, char* fileName)
{
	return QScxmlStateMachine::fromData(static_cast<QIODevice*>(data), QString(fileName));
}

void* QScxmlStateMachine_QScxmlStateMachine_FromFile(char* fileName)
{
	return QScxmlStateMachine::fromFile(QString(fileName));
}

struct QtScxml_PackedString QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(char* prefix)
{
	return ({ QByteArray tec49a6 = QScxmlStateMachine::generateSessionId(QString(prefix)).toUtf8(); QtScxml_PackedString { const_cast<char*>(tec49a6.prepend("WHITESPACE").constData()+10), tec49a6.size()-10 }; });
}

char QScxmlStateMachine_Init(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "init", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QScxmlStateMachine_InitDefault(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::init();
}

void QScxmlStateMachine_ConnectInitializedChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::initializedChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_InitializedChanged));
}

void QScxmlStateMachine_DisconnectInitializedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::initializedChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_InitializedChanged));
}

void QScxmlStateMachine_InitializedChanged(void* ptr, char initialized)
{
	static_cast<QScxmlStateMachine*>(ptr)->initializedChanged(initialized != 0);
}

char QScxmlStateMachine_IsActive(void* ptr, char* scxmlStateName)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isActive(QString(scxmlStateName));
}

char QScxmlStateMachine_IsDispatchableTarget(void* ptr, char* target)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isDispatchableTarget(QString(target));
}

char QScxmlStateMachine_IsInvoked(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInvoked();
}

char QScxmlStateMachine_IsRunning(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isRunning();
}

void QScxmlStateMachine_ConnectLog(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_DisconnectLog(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_Log(void* ptr, char* label, char* msg)
{
	static_cast<QScxmlStateMachine*>(ptr)->log(QString(label), QString(msg));
}

struct QtScxml_PackedString QScxmlStateMachine_Name(void* ptr)
{
	return ({ QByteArray t9dc494 = static_cast<QScxmlStateMachine*>(ptr)->name().toUtf8(); QtScxml_PackedString { const_cast<char*>(t9dc494.prepend("WHITESPACE").constData()+10), t9dc494.size()-10 }; });
}

struct QtScxml_PackedList QScxmlStateMachine_ParseErrors(void* ptr)
{
	return ({ QVector<QScxmlError>* tmpValue = new QVector<QScxmlError>(static_cast<QScxmlStateMachine*>(ptr)->parseErrors()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void QScxmlStateMachine_ConnectReachedStableState(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::reachedStableState), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_ReachedStableState));
}

void QScxmlStateMachine_DisconnectReachedStableState(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::reachedStableState), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_ReachedStableState));
}

void QScxmlStateMachine_ReachedStableState(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->reachedStableState();
}

void QScxmlStateMachine_ConnectRunningChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::runningChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_RunningChanged));
}

void QScxmlStateMachine_DisconnectRunningChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::runningChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_RunningChanged));
}

void QScxmlStateMachine_RunningChanged(void* ptr, char running)
{
	static_cast<QScxmlStateMachine*>(ptr)->runningChanged(running != 0);
}

void* QScxmlStateMachine_ScxmlEventFilter(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->scxmlEventFilter();
}

struct QtScxml_PackedString QScxmlStateMachine_SessionId(void* ptr)
{
	return ({ QByteArray t536cb5 = static_cast<QScxmlStateMachine*>(ptr)->sessionId().toUtf8(); QtScxml_PackedString { const_cast<char*>(t536cb5.prepend("WHITESPACE").constData()+10), t536cb5.size()-10 }; });
}

void QScxmlStateMachine_SetDataModel(void* ptr, void* model)
{
	static_cast<QScxmlStateMachine*>(ptr)->setDataModel(static_cast<QScxmlDataModel*>(model));
}

void QScxmlStateMachine_SetRunning(void* ptr, char running)
{
	static_cast<QScxmlStateMachine*>(ptr)->setRunning(running != 0);
}

void QScxmlStateMachine_SetScxmlEventFilter(void* ptr, void* newFilter)
{
	static_cast<QScxmlStateMachine*>(ptr)->setScxmlEventFilter(static_cast<QScxmlEventFilter*>(newFilter));
}

void QScxmlStateMachine_SetSessionId(void* ptr, char* id)
{
	static_cast<QScxmlStateMachine*>(ptr)->setSessionId(QString(id));
}

void QScxmlStateMachine_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "start");
}

void QScxmlStateMachine_StartDefault(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::start();
}

struct QtScxml_PackedString QScxmlStateMachine_StateNames(void* ptr, char compress)
{
	return ({ QByteArray td672b5 = static_cast<QScxmlStateMachine*>(ptr)->stateNames(compress != 0).join("|").toUtf8(); QtScxml_PackedString { const_cast<char*>(td672b5.prepend("WHITESPACE").constData()+10), td672b5.size()-10 }; });
}

void QScxmlStateMachine_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "stop");
}

void QScxmlStateMachine_StopDefault(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::stop();
}

void QScxmlStateMachine_SubmitEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_SubmitEvent2(void* ptr, char* eventName)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString(eventName));
}

void QScxmlStateMachine_SubmitEvent3(void* ptr, char* eventName, void* data)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString(eventName), *static_cast<QVariant*>(data));
}

void* QScxmlStateMachine___initialValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlStateMachine___initialValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___initialValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___initialValues_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlStateMachine___setInitialValues_initialValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlStateMachine___setInitialValues_initialValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___setInitialValues_initialValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___setInitialValues_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlStateMachine___initialValuesChanged_initialValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QScxmlStateMachine___initialValuesChanged_initialValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___initialValuesChanged_initialValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___initialValuesChanged_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlStateMachine___parseErrors_atList(void* ptr, int i)
{
	return new QScxmlError(static_cast<QVector<QScxmlError>*>(ptr)->at(i));
}

void QScxmlStateMachine___parseErrors_setList(void* ptr, void* i)
{
	static_cast<QVector<QScxmlError>*>(ptr)->append(*static_cast<QScxmlError*>(i));
}

void* QScxmlStateMachine___parseErrors_newList(void* ptr)
{
	return new QVector<QScxmlError>;
}

struct QtScxml_PackedString QScxmlStateMachine_____initialValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____initialValues_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlStateMachine_____initialValues_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

struct QtScxml_PackedString QScxmlStateMachine_____setInitialValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____setInitialValues_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlStateMachine_____setInitialValues_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

struct QtScxml_PackedString QScxmlStateMachine_____initialValuesChanged_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____initialValuesChanged_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QScxmlStateMachine_____initialValuesChanged_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

void* QScxmlStateMachine___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlStateMachine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlStateMachine___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QScxmlStateMachine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlStateMachine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlStateMachine___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QScxmlStateMachine___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlStateMachine___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlStateMachine___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlStateMachine___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlStateMachine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlStateMachine___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QScxmlStateMachine___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlStateMachine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlStateMachine___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QScxmlStateMachine_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlStateMachine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlStateMachine_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlStateMachine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlStateMachine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlStateMachine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::customEvent(static_cast<QEvent*>(event));
}

void QScxmlStateMachine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "deleteLater");
}

void QScxmlStateMachine_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::deleteLater();
}

void QScxmlStateMachine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScxmlStateMachine_Event(void* ptr, void* e)
{
	return static_cast<QScxmlStateMachine*>(ptr)->event(static_cast<QEvent*>(e));
}

char QScxmlStateMachine_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::event(static_cast<QEvent*>(e));
}

char QScxmlStateMachine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlStateMachine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QScxmlStateMachine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlStateMachine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlStateMachine*>(ptr)->metaObject());
}

void* QScxmlStateMachine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::metaObject());
}

