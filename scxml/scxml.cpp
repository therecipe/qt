// +build !minimal

#define protected public
#define private public

#include "scxml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIODevice>
#include <QLayout>
#include <QList>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QScxmlCompiler>
#include <QScxmlCppDataModel>
#include <QScxmlDataModel>
#include <QScxmlDynamicScxmlServiceFactory>
#include <QScxmlEcmaScriptDataModel>
#include <QScxmlError>
#include <QScxmlEvent>
#include <QScxmlInvokableService>
#include <QScxmlInvokableServiceFactory>
#include <QScxmlNullDataModel>
#include <QScxmlStateMachine>
#include <QScxmlStaticScxmlServiceFactory>
#include <QScxmlTableData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>
#include <QXmlStreamReader>

void* QScxmlCompiler_NewQScxmlCompiler(void* reader)
{
	return new QScxmlCompiler(static_cast<QXmlStreamReader*>(reader));
}

void* QScxmlCompiler_Compile(void* ptr)
{
	return static_cast<QScxmlCompiler*>(ptr)->compile();
}

void QScxmlCompiler_SetFileName(void* ptr, struct QtScxml_PackedString fileName)
{
	static_cast<QScxmlCompiler*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QScxmlCompiler_DestroyQScxmlCompiler(void* ptr)
{
	static_cast<QScxmlCompiler*>(ptr)->~QScxmlCompiler();
}

struct QtScxml_PackedString QScxmlCompiler_FileName(void* ptr)
{
	return ({ QByteArray t25588f = static_cast<QScxmlCompiler*>(ptr)->fileName().toUtf8(); QtScxml_PackedString { const_cast<char*>(t25588f.prepend("WHITESPACE").constData()+10), t25588f.size()-10 }; });
}

struct QtScxml_PackedList QScxmlCompiler_Errors(void* ptr)
{
	return ({ QVector<QScxmlError>* tmpValue = new QVector<QScxmlError>(static_cast<QScxmlCompiler*>(ptr)->errors()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlCompiler___errors_atList(void* ptr, int i)
{
	return new QScxmlError(static_cast<QVector<QScxmlError>*>(ptr)->at(i));
}

void QScxmlCompiler___errors_setList(void* ptr, void* i)
{
	static_cast<QVector<QScxmlError>*>(ptr)->append(*static_cast<QScxmlError*>(i));
}

void* QScxmlCompiler___errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlError>;
}

class MyQScxmlCppDataModel: public QScxmlCppDataModel
{
public:
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlCppDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlCppDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlCppDataModel_ScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked)); };
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlCppDataModel_HasScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	
	
	
	
	
	
	
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, stateMachine); };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlDataModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlDataModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlCppDataModel*)

int QScxmlCppDataModel_QScxmlCppDataModel_QRegisterMetaType(){qRegisterMetaType<QScxmlCppDataModel*>(); return qRegisterMetaType<MyQScxmlCppDataModel*>();}

char QScxmlCppDataModel_SetScxmlProperty(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlCppDataModel_SetScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
		return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlCppDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlCppDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
		return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void QScxmlCppDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void* QScxmlCppDataModel_ScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

void* QScxmlCppDataModel_ScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

char QScxmlCppDataModel_HasScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

char QScxmlCppDataModel_HasScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

char QScxmlCppDataModel_InState(void* ptr, struct QtScxml_PackedString stateName)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->inState(QString::fromUtf8(stateName.data, stateName.len));
}

void* QScxmlCppDataModel_ScxmlEvent(void* ptr)
{
	return const_cast<QScxmlEvent*>(&static_cast<QScxmlCppDataModel*>(ptr)->scxmlEvent());
}

class MyQScxmlDataModel: public QScxmlDataModel
{
public:
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, stateMachine); };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlDataModel_ScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked)); };
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlDataModel_HasScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlDataModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlDataModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlDataModel*)

int QScxmlDataModel_QScxmlDataModel_QRegisterMetaType(){qRegisterMetaType<QScxmlDataModel*>(); return qRegisterMetaType<MyQScxmlDataModel*>();}

char QScxmlDataModel_SetScxmlProperty(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
	return static_cast<QScxmlDataModel*>(ptr)->setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void QScxmlDataModel_SetStateMachine(void* ptr, void* stateMachine)
{
	static_cast<QScxmlDataModel*>(ptr)->setStateMachine(static_cast<QScxmlStateMachine*>(stateMachine));
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

void* QScxmlDataModel_StateMachine(void* ptr)
{
	return static_cast<QScxmlDataModel*>(ptr)->stateMachine();
}

void* QScxmlDataModel_ScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return new QVariant(static_cast<QScxmlDataModel*>(ptr)->scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

char QScxmlDataModel_HasScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return static_cast<QScxmlDataModel*>(ptr)->hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

void* QScxmlDataModel___setup_initialDataValues_atList(void* ptr, struct QtScxml_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QScxmlDataModel___setup_initialDataValues_setList(void* ptr, struct QtScxml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QScxmlDataModel___setup_initialDataValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
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

void QScxmlDataModel_____setup_keyList_setList(void* ptr, struct QtScxml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QScxmlDataModel_____setup_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
	return new QList<QObject*>;
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
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScxmlDataModel_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::event(static_cast<QEvent*>(e));
	}
}

char QScxmlDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QScxmlDataModel_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QScxmlDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QScxmlDataModel_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::customEvent(static_cast<QEvent*>(event));
	}
}

void QScxmlDataModel_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::deleteLater();
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::deleteLater();
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::deleteLater();
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::deleteLater();
	}
}

void QScxmlDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QScxmlDataModel_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QScxmlDataModel_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QScxmlNullDataModel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::metaObject());
	} else if (dynamic_cast<QScxmlEcmaScriptDataModel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::metaObject());
	} else if (dynamic_cast<QScxmlCppDataModel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::metaObject());
	}
}

class MyQScxmlDynamicScxmlServiceFactory: public QScxmlDynamicScxmlServiceFactory
{
public:
	QScxmlInvokableService * invoke(QScxmlStateMachine * parentStateMachine) { return static_cast<QScxmlInvokableService*>(callbackQScxmlDynamicScxmlServiceFactory_Invoke(this, parentStateMachine)); };
	bool event(QEvent * e) { return callbackQScxmlInvokableServiceFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlInvokableServiceFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlInvokableServiceFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlInvokableServiceFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlInvokableServiceFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlInvokableServiceFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlInvokableServiceFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlInvokableServiceFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlInvokableServiceFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlDynamicScxmlServiceFactory*)

int QScxmlDynamicScxmlServiceFactory_QScxmlDynamicScxmlServiceFactory_QRegisterMetaType(){qRegisterMetaType<QScxmlDynamicScxmlServiceFactory*>(); return qRegisterMetaType<MyQScxmlDynamicScxmlServiceFactory*>();}

void* QScxmlDynamicScxmlServiceFactory_Invoke(void* ptr, void* parentStateMachine)
{
	return static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->invoke(static_cast<QScxmlStateMachine*>(parentStateMachine));
}

void* QScxmlDynamicScxmlServiceFactory_InvokeDefault(void* ptr, void* parentStateMachine)
{
		return static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::invoke(static_cast<QScxmlStateMachine*>(parentStateMachine));
}

void* QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_names_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::StringId>;
}

void* QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::ParameterInfo>;
}

class MyQScxmlEcmaScriptDataModel: public QScxmlEcmaScriptDataModel
{
public:
	MyQScxmlEcmaScriptDataModel(QObject *parent = nullptr) : QScxmlEcmaScriptDataModel(parent) {QScxmlEcmaScriptDataModel_QScxmlEcmaScriptDataModel_QRegisterMetaType();};
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlEcmaScriptDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlEcmaScriptDataModel_SetScxmlEvent(this, const_cast<QScxmlEvent*>(&event)); };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlEcmaScriptDataModel_ScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked)); };
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, stateMachine); };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlDataModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlDataModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlEcmaScriptDataModel*)

int QScxmlEcmaScriptDataModel_QScxmlEcmaScriptDataModel_QRegisterMetaType(){qRegisterMetaType<QScxmlEcmaScriptDataModel*>(); return qRegisterMetaType<MyQScxmlEcmaScriptDataModel*>();}

void* QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQScxmlEcmaScriptDataModel(static_cast<QObject*>(parent));
	}
}

char QScxmlEcmaScriptDataModel_SetScxmlProperty(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlEcmaScriptDataModel_SetScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
		return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlEcmaScriptDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlEcmaScriptDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
		return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void QScxmlEcmaScriptDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlEcmaScriptDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void* QScxmlEcmaScriptDataModel_ScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

void* QScxmlEcmaScriptDataModel_ScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

char QScxmlEcmaScriptDataModel_HasScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

char QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

void* QScxmlError_NewQScxmlError()
{
	return new QScxmlError();
}

void* QScxmlError_NewQScxmlError3(void* other)
{
	return new QScxmlError(*static_cast<QScxmlError*>(other));
}

void* QScxmlError_NewQScxmlError2(struct QtScxml_PackedString fileName, int line, int column, struct QtScxml_PackedString description)
{
	return new QScxmlError(QString::fromUtf8(fileName.data, fileName.len), line, column, QString::fromUtf8(description.data, description.len));
}

void QScxmlError_DestroyQScxmlError(void* ptr)
{
	static_cast<QScxmlError*>(ptr)->~QScxmlError();
}

struct QtScxml_PackedString QScxmlError_Description(void* ptr)
{
	return ({ QByteArray td8220c = static_cast<QScxmlError*>(ptr)->description().toUtf8(); QtScxml_PackedString { const_cast<char*>(td8220c.prepend("WHITESPACE").constData()+10), td8220c.size()-10 }; });
}

struct QtScxml_PackedString QScxmlError_FileName(void* ptr)
{
	return ({ QByteArray t71e3d6 = static_cast<QScxmlError*>(ptr)->fileName().toUtf8(); QtScxml_PackedString { const_cast<char*>(t71e3d6.prepend("WHITESPACE").constData()+10), t71e3d6.size()-10 }; });
}

struct QtScxml_PackedString QScxmlError_ToString(void* ptr)
{
	return ({ QByteArray t6a9dc4 = static_cast<QScxmlError*>(ptr)->toString().toUtf8(); QtScxml_PackedString { const_cast<char*>(t6a9dc4.prepend("WHITESPACE").constData()+10), t6a9dc4.size()-10 }; });
}

char QScxmlError_IsValid(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->isValid();
}

int QScxmlError_Column(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->column();
}

int QScxmlError_Line(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->line();
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

void QScxmlEvent_SetData(void* ptr, void* data)
{
	static_cast<QScxmlEvent*>(ptr)->setData(*static_cast<QVariant*>(data));
}

void QScxmlEvent_SetDelay(void* ptr, int delayInMiliSecs)
{
	static_cast<QScxmlEvent*>(ptr)->setDelay(delayInMiliSecs);
}

void QScxmlEvent_SetErrorMessage(void* ptr, struct QtScxml_PackedString message)
{
	static_cast<QScxmlEvent*>(ptr)->setErrorMessage(QString::fromUtf8(message.data, message.len));
}

void QScxmlEvent_SetEventType(void* ptr, long long ty)
{
	static_cast<QScxmlEvent*>(ptr)->setEventType(static_cast<QScxmlEvent::EventType>(ty));
}

void QScxmlEvent_SetInvokeId(void* ptr, struct QtScxml_PackedString invokeid)
{
	static_cast<QScxmlEvent*>(ptr)->setInvokeId(QString::fromUtf8(invokeid.data, invokeid.len));
}

void QScxmlEvent_SetName(void* ptr, struct QtScxml_PackedString name)
{
	static_cast<QScxmlEvent*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QScxmlEvent_SetOrigin(void* ptr, struct QtScxml_PackedString origin)
{
	static_cast<QScxmlEvent*>(ptr)->setOrigin(QString::fromUtf8(origin.data, origin.len));
}

void QScxmlEvent_SetOriginType(void* ptr, struct QtScxml_PackedString origintype)
{
	static_cast<QScxmlEvent*>(ptr)->setOriginType(QString::fromUtf8(origintype.data, origintype.len));
}

void QScxmlEvent_SetSendId(void* ptr, struct QtScxml_PackedString sendid)
{
	static_cast<QScxmlEvent*>(ptr)->setSendId(QString::fromUtf8(sendid.data, sendid.len));
}

void QScxmlEvent_DestroyQScxmlEvent(void* ptr)
{
	static_cast<QScxmlEvent*>(ptr)->~QScxmlEvent();
}

long long QScxmlEvent_EventType(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->eventType();
}

struct QtScxml_PackedString QScxmlEvent_ErrorMessage(void* ptr)
{
	return ({ QByteArray tc788d7 = static_cast<QScxmlEvent*>(ptr)->errorMessage().toUtf8(); QtScxml_PackedString { const_cast<char*>(tc788d7.prepend("WHITESPACE").constData()+10), tc788d7.size()-10 }; });
}

struct QtScxml_PackedString QScxmlEvent_InvokeId(void* ptr)
{
	return ({ QByteArray t0fa333 = static_cast<QScxmlEvent*>(ptr)->invokeId().toUtf8(); QtScxml_PackedString { const_cast<char*>(t0fa333.prepend("WHITESPACE").constData()+10), t0fa333.size()-10 }; });
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

void* QScxmlEvent_Data(void* ptr)
{
	return new QVariant(static_cast<QScxmlEvent*>(ptr)->data());
}

char QScxmlEvent_IsErrorEvent(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->isErrorEvent();
}

int QScxmlEvent_Delay(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->delay();
}

class MyQScxmlInvokableService: public QScxmlInvokableService
{
public:
	MyQScxmlInvokableService(QScxmlStateMachine *parentStateMachine, QScxmlInvokableServiceFactory *parent) : QScxmlInvokableService(parentStateMachine, parent) {QScxmlInvokableService_QScxmlInvokableService_QRegisterMetaType();};
	bool start() { return callbackQScxmlInvokableService_Start(this) != 0; };
	void postEvent(QScxmlEvent * event) { callbackQScxmlInvokableService_PostEvent(this, event); };
	QString id() const { return ({ QtScxml_PackedString tempVal = callbackQScxmlInvokableService_Id(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString name() const { return ({ QtScxml_PackedString tempVal = callbackQScxmlInvokableService_Name(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool event(QEvent * e) { return callbackQScxmlInvokableService_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlInvokableService_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlInvokableService_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableService_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlInvokableService_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlInvokableService_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlInvokableService_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableService_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlInvokableService_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlInvokableService_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlInvokableService_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlInvokableService*)

int QScxmlInvokableService_QScxmlInvokableService_QRegisterMetaType(){qRegisterMetaType<QScxmlInvokableService*>(); return qRegisterMetaType<MyQScxmlInvokableService*>();}

void* QScxmlInvokableService_NewQScxmlInvokableService(void* parentStateMachine, void* parent)
{
	return new MyQScxmlInvokableService(static_cast<QScxmlStateMachine*>(parentStateMachine), static_cast<QScxmlInvokableServiceFactory*>(parent));
}

char QScxmlInvokableService_Start(void* ptr)
{
	return static_cast<QScxmlInvokableService*>(ptr)->start();
}

void QScxmlInvokableService_PostEvent(void* ptr, void* event)
{
	static_cast<QScxmlInvokableService*>(ptr)->postEvent(static_cast<QScxmlEvent*>(event));
}

void* QScxmlInvokableService_ParentStateMachine(void* ptr)
{
	return static_cast<QScxmlInvokableService*>(ptr)->parentStateMachine();
}

struct QtScxml_PackedString QScxmlInvokableService_Id(void* ptr)
{
	return ({ QByteArray t66c5a0 = static_cast<QScxmlInvokableService*>(ptr)->id().toUtf8(); QtScxml_PackedString { const_cast<char*>(t66c5a0.prepend("WHITESPACE").constData()+10), t66c5a0.size()-10 }; });
}

struct QtScxml_PackedString QScxmlInvokableService_Name(void* ptr)
{
	return ({ QByteArray t320c29 = static_cast<QScxmlInvokableService*>(ptr)->name().toUtf8(); QtScxml_PackedString { const_cast<char*>(t320c29.prepend("WHITESPACE").constData()+10), t320c29.size()-10 }; });
}

void* QScxmlInvokableService___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlInvokableService___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlInvokableService___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QScxmlInvokableService___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableService___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableService___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableService___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableService___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableService___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableService___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableService___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableService___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableService___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlInvokableService___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableService___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScxmlInvokableService_EventDefault(void* ptr, void* e)
{
		return static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::event(static_cast<QEvent*>(e));
}

char QScxmlInvokableService_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QScxmlInvokableService_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlInvokableService_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlInvokableService_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::customEvent(static_cast<QEvent*>(event));
}

void QScxmlInvokableService_DeleteLaterDefault(void* ptr)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::deleteLater();
}

void QScxmlInvokableService_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlInvokableService_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QScxmlInvokableService_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QScxmlInvokableService*>(ptr)->QScxmlInvokableService::metaObject());
}

class MyQScxmlInvokableServiceFactory: public QScxmlInvokableServiceFactory
{
public:
	QScxmlInvokableService * invoke(QScxmlStateMachine * parentStateMachine) { return static_cast<QScxmlInvokableService*>(callbackQScxmlInvokableServiceFactory_Invoke(this, parentStateMachine)); };
	bool event(QEvent * e) { return callbackQScxmlInvokableServiceFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlInvokableServiceFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlInvokableServiceFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlInvokableServiceFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlInvokableServiceFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlInvokableServiceFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlInvokableServiceFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlInvokableServiceFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlInvokableServiceFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlInvokableServiceFactory*)

int QScxmlInvokableServiceFactory_QScxmlInvokableServiceFactory_QRegisterMetaType(){qRegisterMetaType<QScxmlInvokableServiceFactory*>(); return qRegisterMetaType<MyQScxmlInvokableServiceFactory*>();}

void* QScxmlInvokableServiceFactory_Invoke(void* ptr, void* parentStateMachine)
{
	return static_cast<QScxmlInvokableServiceFactory*>(ptr)->invoke(static_cast<QScxmlStateMachine*>(parentStateMachine));
}

void* QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_names_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::StringId>;
}

void* QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::ParameterInfo>;
}

void* QScxmlInvokableServiceFactory___parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::ParameterInfo>;
}

void* QScxmlInvokableServiceFactory___names_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::StringId>;
}

void* QScxmlInvokableServiceFactory___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QScxmlInvokableServiceFactory___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScxmlInvokableServiceFactory___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QScxmlInvokableServiceFactory___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableServiceFactory___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableServiceFactory___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableServiceFactory___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableServiceFactory___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableServiceFactory___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableServiceFactory___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QScxmlInvokableServiceFactory___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableServiceFactory___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QScxmlInvokableServiceFactory___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QScxmlInvokableServiceFactory___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScxmlInvokableServiceFactory___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScxmlInvokableServiceFactory_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::event(static_cast<QEvent*>(e));
	}
}

char QScxmlInvokableServiceFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QScxmlInvokableServiceFactory_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QScxmlInvokableServiceFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QScxmlInvokableServiceFactory_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::customEvent(static_cast<QEvent*>(event));
	}
}

void QScxmlInvokableServiceFactory_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::deleteLater();
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::deleteLater();
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::deleteLater();
	}
}

void QScxmlInvokableServiceFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QScxmlInvokableServiceFactory_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QScxmlInvokableServiceFactory_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QScxmlStaticScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::metaObject());
	} else if (dynamic_cast<QScxmlDynamicScxmlServiceFactory*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScxmlDynamicScxmlServiceFactory*>(ptr)->QScxmlDynamicScxmlServiceFactory::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QScxmlInvokableServiceFactory*>(ptr)->QScxmlInvokableServiceFactory::metaObject());
	}
}

class MyQScxmlNullDataModel: public QScxmlNullDataModel
{
public:
	MyQScxmlNullDataModel(QObject *parent = nullptr) : QScxmlNullDataModel(parent) {QScxmlNullDataModel_QScxmlNullDataModel_QRegisterMetaType();};
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tec2727 = context.toUtf8(); QtScxml_PackedString contextPacked = { const_cast<char*>(tec2727.prepend("WHITESPACE").constData()+10), tec2727.size()-10 };return callbackQScxmlNullDataModel_SetScxmlProperty(this, namePacked, const_cast<QVariant*>(&value), contextPacked) != 0; };
	bool setup(const QVariantMap & initialDataValues) { return callbackQScxmlNullDataModel_Setup(this, ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(initialDataValues); QtScxml_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlNullDataModel_SetScxmlEvent(this, const_cast<QScxmlEvent*>(&event)); };
	QVariant scxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return *static_cast<QVariant*>(callbackQScxmlNullDataModel_ScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked)); };
	bool hasScxmlProperty(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtScxml_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQScxmlNullDataModel_HasScxmlProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, stateMachine); };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlDataModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlDataModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlNullDataModel*)

int QScxmlNullDataModel_QScxmlNullDataModel_QRegisterMetaType(){qRegisterMetaType<QScxmlNullDataModel*>(); return qRegisterMetaType<MyQScxmlNullDataModel*>();}

void* QScxmlNullDataModel_NewQScxmlNullDataModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlNullDataModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQScxmlNullDataModel(static_cast<QObject*>(parent));
	}
}

char QScxmlNullDataModel_SetScxmlProperty(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlNullDataModel_SetScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name, void* value, struct QtScxml_PackedString context)
{
		return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), QString::fromUtf8(context.data, context.len));
}

char QScxmlNullDataModel_Setup(void* ptr, void* initialDataValues)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

char QScxmlNullDataModel_SetupDefault(void* ptr, void* initialDataValues)
{
		return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setup(*static_cast<QMap<QString, QVariant>*>(initialDataValues));
}

void QScxmlNullDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlNullDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlNullDataModel_DestroyQScxmlNullDataModel(void* ptr)
{
	static_cast<QScxmlNullDataModel*>(ptr)->~QScxmlNullDataModel();
}

void* QScxmlNullDataModel_ScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

void* QScxmlNullDataModel_ScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::scxmlProperty(QString::fromUtf8(name.data, name.len)));
}

char QScxmlNullDataModel_HasScxmlProperty(void* ptr, struct QtScxml_PackedString name)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

char QScxmlNullDataModel_HasScxmlPropertyDefault(void* ptr, struct QtScxml_PackedString name)
{
		return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::hasScxmlProperty(QString::fromUtf8(name.data, name.len));
}

class MyQScxmlStateMachine: public QScxmlStateMachine
{
public:
	MyQScxmlStateMachine(const QMetaObject *metaObject, QObject *parent = nullptr) : QScxmlStateMachine(metaObject, parent) {QScxmlStateMachine_QScxmlStateMachine_QRegisterMetaType();};
	bool init() { return callbackQScxmlStateMachine_Init(this) != 0; };
	void Signal_DataModelChanged(QScxmlDataModel * model) { callbackQScxmlStateMachine_DataModelChanged(this, model); };
	void Signal_Finished() { callbackQScxmlStateMachine_Finished(this); };
	void Signal_InitializedChanged(bool initialized) { callbackQScxmlStateMachine_InitializedChanged(this, initialized); };
	void Signal_InvokedServicesChanged(const QVector<QScxmlInvokableService *> & invokedServices) { callbackQScxmlStateMachine_InvokedServicesChanged(this, ({ QVector<QScxmlInvokableService *>* tmpValue = const_cast<QVector<QScxmlInvokableService *>*>(&invokedServices); QtScxml_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Log(const QString & label, const QString & msg) { QByteArray t64c653 = label.toUtf8(); QtScxml_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t19f34e = msg.toUtf8(); QtScxml_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQScxmlStateMachine_Log(this, labelPacked, msgPacked); };
	void Signal_ReachedStableState() { callbackQScxmlStateMachine_ReachedStableState(this); };
	void Signal_RunningChanged(bool running) { callbackQScxmlStateMachine_RunningChanged(this, running); };
	void start() { callbackQScxmlStateMachine_Start(this); };
	void stop() { callbackQScxmlStateMachine_Stop(this); };
	void Signal_TableDataChanged(QScxmlTableData * tableData) { callbackQScxmlStateMachine_TableDataChanged(this, tableData); };
	bool event(QEvent * e) { return callbackQScxmlStateMachine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlStateMachine_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlStateMachine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlStateMachine_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlStateMachine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlStateMachine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlStateMachine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlStateMachine_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlStateMachine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlStateMachine*)

int QScxmlStateMachine_QScxmlStateMachine_QRegisterMetaType(){qRegisterMetaType<QScxmlStateMachine*>(); return qRegisterMetaType<MyQScxmlStateMachine*>();}

void* QScxmlStateMachine_QScxmlStateMachine_FromData(void* data, struct QtScxml_PackedString fileName)
{
	return QScxmlStateMachine::fromData(static_cast<QIODevice*>(data), QString::fromUtf8(fileName.data, fileName.len));
}

void* QScxmlStateMachine_QScxmlStateMachine_FromFile(struct QtScxml_PackedString fileName)
{
	return QScxmlStateMachine::fromFile(QString::fromUtf8(fileName.data, fileName.len));
}

void* QScxmlStateMachine_NewQScxmlStateMachine(void* metaObject, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QWindow*>(parent));
	} else {
		return new MyQScxmlStateMachine(static_cast<QMetaObject*>(metaObject), static_cast<QObject*>(parent));
	}
}

struct QtScxml_PackedList QScxmlStateMachine_InitialValues(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QScxmlStateMachine*>(ptr)->initialValues()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
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

void QScxmlStateMachine_CancelDelayedEvent(void* ptr, struct QtScxml_PackedString sendId)
{
	static_cast<QScxmlStateMachine*>(ptr)->cancelDelayedEvent(QString::fromUtf8(sendId.data, sendId.len));
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

void QScxmlStateMachine_ConnectInvokedServicesChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QVector<QScxmlInvokableService *> &)>(&QScxmlStateMachine::invokedServicesChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QVector<QScxmlInvokableService *> &)>(&MyQScxmlStateMachine::Signal_InvokedServicesChanged));
}

void QScxmlStateMachine_DisconnectInvokedServicesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QVector<QScxmlInvokableService *> &)>(&QScxmlStateMachine::invokedServicesChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QVector<QScxmlInvokableService *> &)>(&MyQScxmlStateMachine::Signal_InvokedServicesChanged));
}

void QScxmlStateMachine_InvokedServicesChanged(void* ptr, void* invokedServices)
{
	static_cast<QScxmlStateMachine*>(ptr)->invokedServicesChanged(*static_cast<QVector<QScxmlInvokableService *>*>(invokedServices));
}

void QScxmlStateMachine_ConnectLog(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_DisconnectLog(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_Log(void* ptr, struct QtScxml_PackedString label, struct QtScxml_PackedString msg)
{
	static_cast<QScxmlStateMachine*>(ptr)->log(QString::fromUtf8(label.data, label.len), QString::fromUtf8(msg.data, msg.len));
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

void QScxmlStateMachine_SetDataModel(void* ptr, void* model)
{
	static_cast<QScxmlStateMachine*>(ptr)->setDataModel(static_cast<QScxmlDataModel*>(model));
}

void QScxmlStateMachine_SetInitialValues(void* ptr, void* initialValues)
{
	static_cast<QScxmlStateMachine*>(ptr)->setInitialValues(*static_cast<QMap<QString, QVariant>*>(initialValues));
}

void QScxmlStateMachine_SetRunning(void* ptr, char running)
{
	static_cast<QScxmlStateMachine*>(ptr)->setRunning(running != 0);
}

void QScxmlStateMachine_SetTableData(void* ptr, void* tableData)
{
	static_cast<QScxmlStateMachine*>(ptr)->setTableData(static_cast<QScxmlTableData*>(tableData));
}

void QScxmlStateMachine_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "start");
}

void QScxmlStateMachine_StartDefault(void* ptr)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::start();
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

void QScxmlStateMachine_SubmitEvent2(void* ptr, struct QtScxml_PackedString eventName)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString::fromUtf8(eventName.data, eventName.len));
}

void QScxmlStateMachine_SubmitEvent3(void* ptr, struct QtScxml_PackedString eventName, void* data)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString::fromUtf8(eventName.data, eventName.len), *static_cast<QVariant*>(data));
}

void QScxmlStateMachine_ConnectTableDataChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlTableData *)>(&QScxmlStateMachine::tableDataChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlTableData *)>(&MyQScxmlStateMachine::Signal_TableDataChanged));
}

void QScxmlStateMachine_DisconnectTableDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlTableData *)>(&QScxmlStateMachine::tableDataChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlTableData *)>(&MyQScxmlStateMachine::Signal_TableDataChanged));
}

void QScxmlStateMachine_TableDataChanged(void* ptr, void* tableData)
{
	static_cast<QScxmlStateMachine*>(ptr)->tableDataChanged(static_cast<QScxmlTableData*>(tableData));
}

void* QScxmlStateMachine_DataModel(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->dataModel();
}

void* QScxmlStateMachine_TableData(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->tableData();
}

struct QtScxml_PackedString QScxmlStateMachine_Name(void* ptr)
{
	return ({ QByteArray t9dc494 = static_cast<QScxmlStateMachine*>(ptr)->name().toUtf8(); QtScxml_PackedString { const_cast<char*>(t9dc494.prepend("WHITESPACE").constData()+10), t9dc494.size()-10 }; });
}

struct QtScxml_PackedString QScxmlStateMachine_SessionId(void* ptr)
{
	return ({ QByteArray t536cb5 = static_cast<QScxmlStateMachine*>(ptr)->sessionId().toUtf8(); QtScxml_PackedString { const_cast<char*>(t536cb5.prepend("WHITESPACE").constData()+10), t536cb5.size()-10 }; });
}

struct QtScxml_PackedString QScxmlStateMachine_ActiveStateNames(void* ptr, char compress)
{
	return ({ QByteArray t57b518 = static_cast<QScxmlStateMachine*>(ptr)->activeStateNames(compress != 0).join("|").toUtf8(); QtScxml_PackedString { const_cast<char*>(t57b518.prepend("WHITESPACE").constData()+10), t57b518.size()-10 }; });
}

struct QtScxml_PackedString QScxmlStateMachine_StateNames(void* ptr, char compress)
{
	return ({ QByteArray td672b5 = static_cast<QScxmlStateMachine*>(ptr)->stateNames(compress != 0).join("|").toUtf8(); QtScxml_PackedString { const_cast<char*>(td672b5.prepend("WHITESPACE").constData()+10), td672b5.size()-10 }; });
}

struct QtScxml_PackedList QScxmlStateMachine_ParseErrors(void* ptr)
{
	return ({ QVector<QScxmlError>* tmpValue = new QVector<QScxmlError>(static_cast<QScxmlStateMachine*>(ptr)->parseErrors()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtScxml_PackedList QScxmlStateMachine_InvokedServices(void* ptr)
{
	return ({ QVector<QScxmlInvokableService *>* tmpValue = new QVector<QScxmlInvokableService *>(static_cast<QScxmlStateMachine*>(ptr)->invokedServices()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

char QScxmlStateMachine_IsActive(void* ptr, struct QtScxml_PackedString scxmlStateName)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isActive(QString::fromUtf8(scxmlStateName.data, scxmlStateName.len));
}

char QScxmlStateMachine_IsActive2(void* ptr, int stateIndex)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isActive(stateIndex);
}

char QScxmlStateMachine_IsDispatchableTarget(void* ptr, struct QtScxml_PackedString target)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isDispatchableTarget(QString::fromUtf8(target.data, target.len));
}

char QScxmlStateMachine_IsInitialized(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInitialized();
}

char QScxmlStateMachine_IsInvoked(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInvoked();
}

char QScxmlStateMachine_IsRunning(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isRunning();
}

void* QScxmlStateMachine___initialValues_atList(void* ptr, struct QtScxml_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QScxmlStateMachine___initialValues_setList(void* ptr, struct QtScxml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___initialValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___initialValues_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlStateMachine___initialValuesChanged_initialValues_atList(void* ptr, struct QtScxml_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QScxmlStateMachine___initialValuesChanged_initialValues_setList(void* ptr, struct QtScxml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___initialValuesChanged_initialValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___initialValuesChanged_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtScxml_PackedList { tmpValue, tmpValue->size() }; });
}

void* QScxmlStateMachine___invokedServicesChanged_invokedServices_atList(void* ptr, int i)
{
	return const_cast<QScxmlInvokableService*>(static_cast<QVector<QScxmlInvokableService *>*>(ptr)->at(i));
}

void QScxmlStateMachine___invokedServicesChanged_invokedServices_setList(void* ptr, void* i)
{
	static_cast<QVector<QScxmlInvokableService *>*>(ptr)->append(static_cast<QScxmlInvokableService*>(i));
}

void* QScxmlStateMachine___invokedServicesChanged_invokedServices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlInvokableService *>;
}

void* QScxmlStateMachine___setInitialValues_initialValues_atList(void* ptr, struct QtScxml_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QScxmlStateMachine___setInitialValues_initialValues_setList(void* ptr, struct QtScxml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QScxmlStateMachine___setInitialValues_initialValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtScxml_PackedList QScxmlStateMachine___setInitialValues_keyList(void* ptr)
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
	Q_UNUSED(ptr);
	return new QVector<QScxmlError>;
}

void* QScxmlStateMachine___invokedServices_atList(void* ptr, int i)
{
	return const_cast<QScxmlInvokableService*>(static_cast<QVector<QScxmlInvokableService *>*>(ptr)->at(i));
}

void QScxmlStateMachine___invokedServices_setList(void* ptr, void* i)
{
	static_cast<QVector<QScxmlInvokableService *>*>(ptr)->append(static_cast<QScxmlInvokableService*>(i));
}

void* QScxmlStateMachine___invokedServices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlInvokableService *>;
}

struct QtScxml_PackedString QScxmlStateMachine_____initialValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____initialValues_keyList_setList(void* ptr, struct QtScxml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QScxmlStateMachine_____initialValues_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtScxml_PackedString QScxmlStateMachine_____initialValuesChanged_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____initialValuesChanged_keyList_setList(void* ptr, struct QtScxml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QScxmlStateMachine_____initialValuesChanged_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtScxml_PackedString QScxmlStateMachine_____setInitialValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtScxml_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QScxmlStateMachine_____setInitialValues_keyList_setList(void* ptr, struct QtScxml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QScxmlStateMachine_____setInitialValues_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
	return new QList<QObject*>;
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
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QScxmlStateMachine_EventDefault(void* ptr, void* e)
{
		return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::event(static_cast<QEvent*>(e));
}

char QScxmlStateMachine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QScxmlStateMachine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlStateMachine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::customEvent(static_cast<QEvent*>(event));
}

void QScxmlStateMachine_DeleteLaterDefault(void* ptr)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::deleteLater();
}

void QScxmlStateMachine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QScxmlStateMachine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::metaObject());
}

class MyQScxmlStaticScxmlServiceFactory: public QScxmlStaticScxmlServiceFactory
{
public:
	QScxmlInvokableService * invoke(QScxmlStateMachine * parentStateMachine) { return static_cast<QScxmlInvokableService*>(callbackQScxmlStaticScxmlServiceFactory_Invoke(this, parentStateMachine)); };
	bool event(QEvent * e) { return callbackQScxmlInvokableServiceFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlInvokableServiceFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQScxmlInvokableServiceFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlInvokableServiceFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQScxmlInvokableServiceFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScxmlInvokableServiceFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlInvokableServiceFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtScxml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScxmlInvokableServiceFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlInvokableServiceFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlInvokableServiceFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQScxmlStaticScxmlServiceFactory*)

int QScxmlStaticScxmlServiceFactory_QScxmlStaticScxmlServiceFactory_QRegisterMetaType(){qRegisterMetaType<QScxmlStaticScxmlServiceFactory*>(); return qRegisterMetaType<MyQScxmlStaticScxmlServiceFactory*>();}

void* QScxmlStaticScxmlServiceFactory_Invoke(void* ptr, void* parentStateMachine)
{
	return static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->invoke(static_cast<QScxmlStateMachine*>(parentStateMachine));
}

void* QScxmlStaticScxmlServiceFactory_InvokeDefault(void* ptr, void* parentStateMachine)
{
		return static_cast<QScxmlStaticScxmlServiceFactory*>(ptr)->QScxmlStaticScxmlServiceFactory::invoke(static_cast<QScxmlStateMachine*>(parentStateMachine));
}

void* QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_nameList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::StringId>;
}

void* QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QScxmlExecutableContent::ParameterInfo>;
}

class MyQScxmlTableData: public QScxmlTableData
{
public:
	 ~MyQScxmlTableData() { callbackQScxmlTableData_DestroyQScxmlTableData(this); };
	
	
	
	QScxmlInvokableServiceFactory * serviceFactory(int id) const { return static_cast<QScxmlInvokableServiceFactory*>(callbackQScxmlTableData_ServiceFactory(const_cast<void*>(static_cast<const void*>(this)), id)); };
	QString name() const { return ({ QtScxml_PackedString tempVal = callbackQScxmlTableData_Name(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	
};

void QScxmlTableData_DestroyQScxmlTableData(void* ptr)
{
	static_cast<QScxmlTableData*>(ptr)->~QScxmlTableData();
}

void QScxmlTableData_DestroyQScxmlTableDataDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScxmlTableData_ServiceFactory(void* ptr, int id)
{
	return static_cast<QScxmlTableData*>(ptr)->serviceFactory(id);
}

struct QtScxml_PackedString QScxmlTableData_Name(void* ptr)
{
	return ({ QByteArray t814496 = static_cast<QScxmlTableData*>(ptr)->name().toUtf8(); QtScxml_PackedString { const_cast<char*>(t814496.prepend("WHITESPACE").constData()+10), t814496.size()-10 }; });
}

