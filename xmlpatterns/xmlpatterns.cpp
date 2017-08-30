// +build !minimal

#define protected public
#define private public

#include "xmlpatterns.h"
#include "_cgo_export.h"

#include <QAbstractMessageHandler>
#include <QAbstractUriResolver>
#include <QAbstractXmlNodeModel>
#include <QAbstractXmlReceiver>
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
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QSimpleXmlNodeModel>
#include <QSourceLocation>
#include <QString>
#include <QStringList>
#include <QStringRef>
#include <QTextCodec>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>
#include <QXmlFormatter>
#include <QXmlItem>
#include <QXmlName>
#include <QXmlNamePool>
#include <QXmlNodeModelIndex>
#include <QXmlQuery>
#include <QXmlResultItems>
#include <QXmlSchema>
#include <QXmlSchemaValidator>
#include <QXmlSerializer>

class MyQAbstractMessageHandler: public QAbstractMessageHandler
{
public:
	 ~MyQAbstractMessageHandler() { callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler(this); };
	bool event(QEvent * e) { return callbackQAbstractMessageHandler_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractMessageHandler_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractMessageHandler_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractMessageHandler_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractMessageHandler_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractMessageHandler_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractMessageHandler_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractMessageHandler_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtXmlPatterns_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractMessageHandler_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractMessageHandler_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractMessageHandler_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAbstractMessageHandler*)

int QAbstractMessageHandler_QAbstractMessageHandler_QRegisterMetaType(){qRegisterMetaType<QAbstractMessageHandler*>(); return qRegisterMetaType<MyQAbstractMessageHandler*>();}

void QAbstractMessageHandler_DestroyQAbstractMessageHandler(void* ptr)
{
	static_cast<QAbstractMessageHandler*>(ptr)->~QAbstractMessageHandler();
}

void QAbstractMessageHandler_DestroyQAbstractMessageHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractMessageHandler___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QAbstractMessageHandler___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractMessageHandler___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QAbstractMessageHandler___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractMessageHandler___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractMessageHandler___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractMessageHandler___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractMessageHandler___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractMessageHandler___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractMessageHandler___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractMessageHandler___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractMessageHandler___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractMessageHandler___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QAbstractMessageHandler___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractMessageHandler___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QAbstractMessageHandler_EventDefault(void* ptr, void* e)
{
		return static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::event(static_cast<QEvent*>(e));
}

char QAbstractMessageHandler_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAbstractMessageHandler_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractMessageHandler_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractMessageHandler_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::customEvent(static_cast<QEvent*>(event));
}

void QAbstractMessageHandler_DeleteLaterDefault(void* ptr)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::deleteLater();
}

void QAbstractMessageHandler_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractMessageHandler_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QAbstractMessageHandler_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::metaObject());
}

class MyQAbstractUriResolver: public QAbstractUriResolver
{
public:
	MyQAbstractUriResolver(QObject *parent = Q_NULLPTR) : QAbstractUriResolver(parent) {QAbstractUriResolver_QAbstractUriResolver_QRegisterMetaType();};
	 ~MyQAbstractUriResolver() { callbackQAbstractUriResolver_DestroyQAbstractUriResolver(this); };
	QUrl resolve(const QUrl & relative, const QUrl & baseURI) const { return *static_cast<QUrl*>(callbackQAbstractUriResolver_Resolve(const_cast<void*>(static_cast<const void*>(this)), const_cast<QUrl*>(&relative), const_cast<QUrl*>(&baseURI))); };
	bool event(QEvent * e) { return callbackQAbstractUriResolver_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractUriResolver_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractUriResolver_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractUriResolver_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractUriResolver_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractUriResolver_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractUriResolver_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractUriResolver_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtXmlPatterns_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractUriResolver_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractUriResolver_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractUriResolver_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAbstractUriResolver*)

int QAbstractUriResolver_QAbstractUriResolver_QRegisterMetaType(){qRegisterMetaType<QAbstractUriResolver*>(); return qRegisterMetaType<MyQAbstractUriResolver*>();}

void* QAbstractUriResolver_NewQAbstractUriResolver(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractUriResolver(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractUriResolver(static_cast<QObject*>(parent));
	}
}

void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr)
{
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

void QAbstractUriResolver_DestroyQAbstractUriResolverDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractUriResolver_Resolve(void* ptr, void* relative, void* baseURI)
{
	return new QUrl(static_cast<QAbstractUriResolver*>(ptr)->resolve(*static_cast<QUrl*>(relative), *static_cast<QUrl*>(baseURI)));
}

void* QAbstractUriResolver___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QAbstractUriResolver___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractUriResolver___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QAbstractUriResolver___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractUriResolver___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractUriResolver___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractUriResolver___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractUriResolver___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractUriResolver___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractUriResolver___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractUriResolver___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractUriResolver___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractUriResolver___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QAbstractUriResolver___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractUriResolver___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QAbstractUriResolver_EventDefault(void* ptr, void* e)
{
		return static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::event(static_cast<QEvent*>(e));
}

char QAbstractUriResolver_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAbstractUriResolver_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractUriResolver_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractUriResolver_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::customEvent(static_cast<QEvent*>(event));
}

void QAbstractUriResolver_DeleteLaterDefault(void* ptr)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::deleteLater();
}

void QAbstractUriResolver_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractUriResolver_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QAbstractUriResolver_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::metaObject());
}

class MyQAbstractXmlNodeModel: public QAbstractXmlNodeModel
{
public:
	 ~MyQAbstractXmlNodeModel() { callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(this); };
	QString stringValue(const QXmlNodeModelIndex & n) const { return ({ QtXmlPatterns_PackedString tempVal = callbackQAbstractXmlNodeModel_StringValue(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n)); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QUrl baseUri(const QXmlNodeModelIndex & n) const { return *static_cast<QUrl*>(callbackQAbstractXmlNodeModel_BaseUri(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QUrl documentUri(const QXmlNodeModelIndex & n) const { return *static_cast<QUrl*>(callbackQAbstractXmlNodeModel_DocumentUri(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QVariant typedValue(const QXmlNodeModelIndex & node) const { return *static_cast<QVariant*>(callbackQAbstractXmlNodeModel_TypedValue(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&node))); };
	QVector<QXmlName> namespaceBindings(const QXmlNodeModelIndex & n) const { return *static_cast<QVector<QXmlName>*>(callbackQAbstractXmlNodeModel_NamespaceBindings(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QVector<QXmlNodeModelIndex> nodesByIdref(const QXmlName & idref) const { return *static_cast<QVector<QXmlNodeModelIndex>*>(callbackQAbstractXmlNodeModel_NodesByIdref(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlName*>(&idref))); };
	QXmlName name(const QXmlNodeModelIndex & ni) const { return *static_cast<QXmlName*>(callbackQAbstractXmlNodeModel_Name(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni))); };
	QXmlNodeModelIndex elementById(const QXmlName & id) const { return *static_cast<QXmlNodeModelIndex*>(callbackQAbstractXmlNodeModel_ElementById(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlName*>(&id))); };
	QXmlNodeModelIndex nextFromSimpleAxis(QAbstractXmlNodeModel::SimpleAxis axis, const QXmlNodeModelIndex & origin) const { return *static_cast<QXmlNodeModelIndex*>(callbackQAbstractXmlNodeModel_NextFromSimpleAxis(const_cast<void*>(static_cast<const void*>(this)), axis, const_cast<QXmlNodeModelIndex*>(&origin))); };
	QXmlNodeModelIndex root(const QXmlNodeModelIndex & n) const { return *static_cast<QXmlNodeModelIndex*>(callbackQAbstractXmlNodeModel_Root(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QXmlNodeModelIndex::DocumentOrder compareOrder(const QXmlNodeModelIndex & ni1, const QXmlNodeModelIndex & ni2) const { return static_cast<QXmlNodeModelIndex::DocumentOrder>(callbackQAbstractXmlNodeModel_CompareOrder(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni1), const_cast<QXmlNodeModelIndex*>(&ni2))); };
	QXmlNodeModelIndex::NodeKind kind(const QXmlNodeModelIndex & ni) const { return static_cast<QXmlNodeModelIndex::NodeKind>(callbackQAbstractXmlNodeModel_Kind(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni))); };
};

void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr)
{
	static_cast<QAbstractXmlNodeModel*>(ptr)->~QAbstractXmlNodeModel();
}

void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractXmlNodeModel_SourceLocation(void* ptr, void* index)
{
	return new QSourceLocation(static_cast<QAbstractXmlNodeModel*>(ptr)->sourceLocation(*static_cast<QXmlNodeModelIndex*>(index)));
}

struct QtXmlPatterns_PackedString QAbstractXmlNodeModel_StringValue(void* ptr, void* n)
{
	return ({ QByteArray t02d1c5 = static_cast<QAbstractXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(n)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(t02d1c5.prepend("WHITESPACE").constData()+10), t02d1c5.size()-10 }; });
}

void* QAbstractXmlNodeModel_BaseUri(void* ptr, void* n)
{
	return new QUrl(static_cast<QAbstractXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QAbstractXmlNodeModel_DocumentUri(void* ptr, void* n)
{
	return new QUrl(static_cast<QAbstractXmlNodeModel*>(ptr)->documentUri(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node)
{
	return new QVariant(static_cast<QAbstractXmlNodeModel*>(ptr)->typedValue(*static_cast<QXmlNodeModelIndex*>(node)));
}

struct QtXmlPatterns_PackedList QAbstractXmlNodeModel_NamespaceBindings(void* ptr, void* n)
{
	return ({ QVector<QXmlName>* tmpValue = new QVector<QXmlName>(static_cast<QAbstractXmlNodeModel*>(ptr)->namespaceBindings(*static_cast<QXmlNodeModelIndex*>(n))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtXmlPatterns_PackedList QAbstractXmlNodeModel_NodesByIdref(void* ptr, void* idref)
{
	return ({ QVector<QXmlNodeModelIndex>* tmpValue = new QVector<QXmlNodeModelIndex>(static_cast<QAbstractXmlNodeModel*>(ptr)->nodesByIdref(*static_cast<QXmlName*>(idref))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

void* QAbstractXmlNodeModel_Name(void* ptr, void* ni)
{
	return new QXmlName(static_cast<QAbstractXmlNodeModel*>(ptr)->name(*static_cast<QXmlNodeModelIndex*>(ni)));
}

void* QAbstractXmlNodeModel_CreateIndex(void* ptr, long long data)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->createIndex(data));
}

void* QAbstractXmlNodeModel_CreateIndex3(void* ptr, long long data, long long additionalData)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->createIndex(data, additionalData));
}

void* QAbstractXmlNodeModel_CreateIndex2(void* ptr, void* pointer, long long additionalData)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->createIndex(pointer, additionalData));
}

void* QAbstractXmlNodeModel_ElementById(void* ptr, void* id)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->elementById(*static_cast<QXmlName*>(id)));
}

void* QAbstractXmlNodeModel_NextFromSimpleAxis(void* ptr, long long axis, void* origin)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->nextFromSimpleAxis(static_cast<QAbstractXmlNodeModel::SimpleAxis>(axis), *static_cast<QXmlNodeModelIndex*>(origin)));
}

void* QAbstractXmlNodeModel_Root(void* ptr, void* n)
{
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->root(*static_cast<QXmlNodeModelIndex*>(n)));
}

long long QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2)
{
	return static_cast<QAbstractXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

long long QAbstractXmlNodeModel_Kind(void* ptr, void* ni)
{
	return static_cast<QAbstractXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

void* QAbstractXmlNodeModel___namespaceBindings_atList(void* ptr, int i)
{
	return new QXmlName(static_cast<QVector<QXmlName>*>(ptr)->at(i));
}

void QAbstractXmlNodeModel___namespaceBindings_setList(void* ptr, void* i)
{
	static_cast<QVector<QXmlName>*>(ptr)->append(*static_cast<QXmlName*>(i));
}

void* QAbstractXmlNodeModel___namespaceBindings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QXmlName>;
}

void* QAbstractXmlNodeModel___nodesByIdref_atList(void* ptr, int i)
{
	return new QXmlNodeModelIndex(static_cast<QVector<QXmlNodeModelIndex>*>(ptr)->at(i));
}

void QAbstractXmlNodeModel___nodesByIdref_setList(void* ptr, void* i)
{
	static_cast<QVector<QXmlNodeModelIndex>*>(ptr)->append(*static_cast<QXmlNodeModelIndex*>(i));
}

void* QAbstractXmlNodeModel___nodesByIdref_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QXmlNodeModelIndex>;
}

class MyQAbstractXmlReceiver: public QAbstractXmlReceiver
{
public:
	MyQAbstractXmlReceiver() : QAbstractXmlReceiver() {};
	void atomicValue(const QVariant & value) { callbackQAbstractXmlReceiver_AtomicValue(this, const_cast<QVariant*>(&value)); };
	void attribute(const QXmlName & name, const QStringRef & value) { callbackQAbstractXmlReceiver_Attribute(this, const_cast<QXmlName*>(&name), const_cast<QStringRef*>(&value)); };
	void characters(const QStringRef & value) { callbackQAbstractXmlReceiver_Characters(this, const_cast<QStringRef*>(&value)); };
	void comment(const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQAbstractXmlReceiver_Comment(this, valuePacked); };
	void endDocument() { callbackQAbstractXmlReceiver_EndDocument(this); };
	void endElement() { callbackQAbstractXmlReceiver_EndElement(this); };
	void endOfSequence() { callbackQAbstractXmlReceiver_EndOfSequence(this); };
	void namespaceBinding(const QXmlName & name) { callbackQAbstractXmlReceiver_NamespaceBinding(this, const_cast<QXmlName*>(&name)); };
	void processingInstruction(const QXmlName & target, const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQAbstractXmlReceiver_ProcessingInstruction(this, const_cast<QXmlName*>(&target), valuePacked); };
	void startDocument() { callbackQAbstractXmlReceiver_StartDocument(this); };
	void startElement(const QXmlName & name) { callbackQAbstractXmlReceiver_StartElement(this, const_cast<QXmlName*>(&name)); };
	void startOfSequence() { callbackQAbstractXmlReceiver_StartOfSequence(this); };
	 ~MyQAbstractXmlReceiver() { callbackQAbstractXmlReceiver_DestroyQAbstractXmlReceiver(this); };
};

void* QAbstractXmlReceiver_NewQAbstractXmlReceiver()
{
	return new MyQAbstractXmlReceiver();
}

void QAbstractXmlReceiver_AtomicValue(void* ptr, void* value)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->atomicValue(*static_cast<QVariant*>(value));
}

void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Characters(void* ptr, void* value)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Comment(void* ptr, struct QtXmlPatterns_PackedString value)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->comment(QString::fromUtf8(value.data, value.len));
}

void QAbstractXmlReceiver_EndDocument(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->endDocument();
}

void QAbstractXmlReceiver_EndElement(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->endElement();
}

void QAbstractXmlReceiver_EndOfSequence(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->endOfSequence();
}

void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, struct QtXmlPatterns_PackedString value)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->processingInstruction(*static_cast<QXmlName*>(target), QString::fromUtf8(value.data, value.len));
}

void QAbstractXmlReceiver_StartDocument(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->startDocument();
}

void QAbstractXmlReceiver_StartElement(void* ptr, void* name)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_StartOfSequence(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->startOfSequence();
}

void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr)
{
	static_cast<QAbstractXmlReceiver*>(ptr)->~QAbstractXmlReceiver();
}

void QAbstractXmlReceiver_DestroyQAbstractXmlReceiverDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQSimpleXmlNodeModel: public QSimpleXmlNodeModel
{
public:
	 ~MyQSimpleXmlNodeModel() { callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(this); };
	QString stringValue(const QXmlNodeModelIndex & node) const { return ({ QtXmlPatterns_PackedString tempVal = callbackQSimpleXmlNodeModel_StringValue(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&node)); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QUrl baseUri(const QXmlNodeModelIndex & node) const { return *static_cast<QUrl*>(callbackQSimpleXmlNodeModel_BaseUri(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&node))); };
	QVector<QXmlName> namespaceBindings(const QXmlNodeModelIndex & node) const { return *static_cast<QVector<QXmlName>*>(callbackQSimpleXmlNodeModel_NamespaceBindings(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&node))); };
	QVector<QXmlNodeModelIndex> nodesByIdref(const QXmlName & idref) const { return *static_cast<QVector<QXmlNodeModelIndex>*>(callbackQSimpleXmlNodeModel_NodesByIdref(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlName*>(&idref))); };
	QXmlNodeModelIndex elementById(const QXmlName & id) const { return *static_cast<QXmlNodeModelIndex*>(callbackQSimpleXmlNodeModel_ElementById(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlName*>(&id))); };
	QUrl documentUri(const QXmlNodeModelIndex & n) const { return *static_cast<QUrl*>(callbackQSimpleXmlNodeModel_DocumentUri(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QVariant typedValue(const QXmlNodeModelIndex & node) const { return *static_cast<QVariant*>(callbackQSimpleXmlNodeModel_TypedValue(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&node))); };
	QXmlName name(const QXmlNodeModelIndex & ni) const { return *static_cast<QXmlName*>(callbackQSimpleXmlNodeModel_Name(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni))); };
	QXmlNodeModelIndex nextFromSimpleAxis(QAbstractXmlNodeModel::SimpleAxis axis, const QXmlNodeModelIndex & origin) const { return *static_cast<QXmlNodeModelIndex*>(callbackQSimpleXmlNodeModel_NextFromSimpleAxis(const_cast<void*>(static_cast<const void*>(this)), axis, const_cast<QXmlNodeModelIndex*>(&origin))); };
	QXmlNodeModelIndex root(const QXmlNodeModelIndex & n) const { return *static_cast<QXmlNodeModelIndex*>(callbackQSimpleXmlNodeModel_Root(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&n))); };
	QXmlNodeModelIndex::DocumentOrder compareOrder(const QXmlNodeModelIndex & ni1, const QXmlNodeModelIndex & ni2) const { return static_cast<QXmlNodeModelIndex::DocumentOrder>(callbackQSimpleXmlNodeModel_CompareOrder(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni1), const_cast<QXmlNodeModelIndex*>(&ni2))); };
	QXmlNodeModelIndex::NodeKind kind(const QXmlNodeModelIndex & ni) const { return static_cast<QXmlNodeModelIndex::NodeKind>(callbackQSimpleXmlNodeModel_Kind(const_cast<void*>(static_cast<const void*>(this)), const_cast<QXmlNodeModelIndex*>(&ni))); };
};

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr)
{
	static_cast<QSimpleXmlNodeModel*>(ptr)->~QSimpleXmlNodeModel();
}

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtXmlPatterns_PackedString QSimpleXmlNodeModel_StringValue(void* ptr, void* node)
{
	return ({ QByteArray t8d99fb = static_cast<QSimpleXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(t8d99fb.prepend("WHITESPACE").constData()+10), t8d99fb.size()-10 }; });
}

struct QtXmlPatterns_PackedString QSimpleXmlNodeModel_StringValueDefault(void* ptr, void* node)
{
		return ({ QByteArray tf9d087 = static_cast<QSimpleXmlNodeModel*>(ptr)->QSimpleXmlNodeModel::stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(tf9d087.prepend("WHITESPACE").constData()+10), tf9d087.size()-10 }; });
}

void* QSimpleXmlNodeModel_BaseUri(void* ptr, void* node)
{
	return new QUrl(static_cast<QSimpleXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(node)));
}

void* QSimpleXmlNodeModel_BaseUriDefault(void* ptr, void* node)
{
		return new QUrl(static_cast<QSimpleXmlNodeModel*>(ptr)->QSimpleXmlNodeModel::baseUri(*static_cast<QXmlNodeModelIndex*>(node)));
}

struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NamespaceBindings(void* ptr, void* node)
{
	return ({ QVector<QXmlName>* tmpValue = new QVector<QXmlName>(static_cast<QSimpleXmlNodeModel*>(ptr)->namespaceBindings(*static_cast<QXmlNodeModelIndex*>(node))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NamespaceBindingsDefault(void* ptr, void* node)
{
		return ({ QVector<QXmlName>* tmpValue = new QVector<QXmlName>(static_cast<QSimpleXmlNodeModel*>(ptr)->QSimpleXmlNodeModel::namespaceBindings(*static_cast<QXmlNodeModelIndex*>(node))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NodesByIdref(void* ptr, void* idref)
{
	return ({ QVector<QXmlNodeModelIndex>* tmpValue = new QVector<QXmlNodeModelIndex>(static_cast<QSimpleXmlNodeModel*>(ptr)->nodesByIdref(*static_cast<QXmlName*>(idref))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NodesByIdrefDefault(void* ptr, void* idref)
{
		return ({ QVector<QXmlNodeModelIndex>* tmpValue = new QVector<QXmlNodeModelIndex>(static_cast<QSimpleXmlNodeModel*>(ptr)->QSimpleXmlNodeModel::nodesByIdref(*static_cast<QXmlName*>(idref))); QtXmlPatterns_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSimpleXmlNodeModel_NamePool(void* ptr)
{
	return new QXmlNamePool(static_cast<QSimpleXmlNodeModel*>(ptr)->namePool());
}

void* QSimpleXmlNodeModel_ElementById(void* ptr, void* id)
{
	return new QXmlNodeModelIndex(static_cast<QSimpleXmlNodeModel*>(ptr)->elementById(*static_cast<QXmlName*>(id)));
}

void* QSimpleXmlNodeModel_ElementByIdDefault(void* ptr, void* id)
{
		return new QXmlNodeModelIndex(static_cast<QSimpleXmlNodeModel*>(ptr)->QSimpleXmlNodeModel::elementById(*static_cast<QXmlName*>(id)));
}

void* QSimpleXmlNodeModel_DocumentUri(void* ptr, void* n)
{
	return new QUrl(static_cast<QSimpleXmlNodeModel*>(ptr)->documentUri(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QSimpleXmlNodeModel_DocumentUriDefault(void* ptr, void* n)
{
	Q_UNUSED(ptr);
	Q_UNUSED(n);
	
}

void* QSimpleXmlNodeModel_TypedValue(void* ptr, void* node)
{
	return new QVariant(static_cast<QSimpleXmlNodeModel*>(ptr)->typedValue(*static_cast<QXmlNodeModelIndex*>(node)));
}

void* QSimpleXmlNodeModel_TypedValueDefault(void* ptr, void* node)
{
	Q_UNUSED(ptr);
	Q_UNUSED(node);
	
}

void* QSimpleXmlNodeModel_Name(void* ptr, void* ni)
{
	return new QXmlName(static_cast<QSimpleXmlNodeModel*>(ptr)->name(*static_cast<QXmlNodeModelIndex*>(ni)));
}

void* QSimpleXmlNodeModel_NameDefault(void* ptr, void* ni)
{
	Q_UNUSED(ptr);
	Q_UNUSED(ni);
	
}

void* QSimpleXmlNodeModel_NextFromSimpleAxis(void* ptr, long long axis, void* origin)
{
	return new QXmlNodeModelIndex(static_cast<QSimpleXmlNodeModel*>(ptr)->nextFromSimpleAxis(static_cast<QAbstractXmlNodeModel::SimpleAxis>(axis), *static_cast<QXmlNodeModelIndex*>(origin)));
}

void* QSimpleXmlNodeModel_NextFromSimpleAxisDefault(void* ptr, long long axis, void* origin)
{
	Q_UNUSED(ptr);
	Q_UNUSED(axis);
	Q_UNUSED(origin);
	
}

void* QSimpleXmlNodeModel_Root(void* ptr, void* n)
{
	return new QXmlNodeModelIndex(static_cast<QSimpleXmlNodeModel*>(ptr)->root(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QSimpleXmlNodeModel_RootDefault(void* ptr, void* n)
{
	Q_UNUSED(ptr);
	Q_UNUSED(n);
	
}

long long QSimpleXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2)
{
	return static_cast<QSimpleXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

long long QSimpleXmlNodeModel_CompareOrderDefault(void* ptr, void* ni1, void* ni2)
{
	Q_UNUSED(ptr);
	Q_UNUSED(ni1);
	Q_UNUSED(ni2);
	
}

long long QSimpleXmlNodeModel_Kind(void* ptr, void* ni)
{
	return static_cast<QSimpleXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

long long QSimpleXmlNodeModel_KindDefault(void* ptr, void* ni)
{
	Q_UNUSED(ptr);
	Q_UNUSED(ni);
	
}

void* QSourceLocation_NewQSourceLocation()
{
	return new QSourceLocation();
}

void* QSourceLocation_NewQSourceLocation2(void* other)
{
	return new QSourceLocation(*static_cast<QSourceLocation*>(other));
}

void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c)
{
	return new QSourceLocation(*static_cast<QUrl*>(u), l, c);
}

void QSourceLocation_SetColumn(void* ptr, long long newColumn)
{
	static_cast<QSourceLocation*>(ptr)->setColumn(newColumn);
}

void QSourceLocation_SetLine(void* ptr, long long newLine)
{
	static_cast<QSourceLocation*>(ptr)->setLine(newLine);
}

void QSourceLocation_SetUri(void* ptr, void* newUri)
{
	static_cast<QSourceLocation*>(ptr)->setUri(*static_cast<QUrl*>(newUri));
}

void QSourceLocation_DestroyQSourceLocation(void* ptr)
{
	static_cast<QSourceLocation*>(ptr)->~QSourceLocation();
}

void* QSourceLocation_Uri(void* ptr)
{
	return new QUrl(static_cast<QSourceLocation*>(ptr)->uri());
}

char QSourceLocation_IsNull(void* ptr)
{
	return static_cast<QSourceLocation*>(ptr)->isNull();
}

long long QSourceLocation_Column(void* ptr)
{
	return static_cast<QSourceLocation*>(ptr)->column();
}

long long QSourceLocation_Line(void* ptr)
{
	return static_cast<QSourceLocation*>(ptr)->line();
}

class MyQXmlFormatter: public QXmlFormatter
{
public:
	MyQXmlFormatter(const QXmlQuery &query, QIODevice *outputDevice) : QXmlFormatter(query, outputDevice) {};
	void atomicValue(const QVariant & value) { callbackQXmlSerializer_AtomicValue(this, const_cast<QVariant*>(&value)); };
	void attribute(const QXmlName & name, const QStringRef & value) { callbackQXmlSerializer_Attribute(this, const_cast<QXmlName*>(&name), const_cast<QStringRef*>(&value)); };
	void characters(const QStringRef & value) { callbackQXmlSerializer_Characters(this, const_cast<QStringRef*>(&value)); };
	void comment(const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQXmlSerializer_Comment(this, valuePacked); };
	void endDocument() { callbackQXmlSerializer_EndDocument(this); };
	void endElement() { callbackQXmlSerializer_EndElement(this); };
	void endOfSequence() { callbackQXmlSerializer_EndOfSequence(this); };
	void processingInstruction(const QXmlName & name, const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQXmlSerializer_ProcessingInstruction(this, const_cast<QXmlName*>(&name), valuePacked); };
	void startDocument() { callbackQXmlSerializer_StartDocument(this); };
	void startElement(const QXmlName & name) { callbackQXmlSerializer_StartElement(this, const_cast<QXmlName*>(&name)); };
	void startOfSequence() { callbackQXmlSerializer_StartOfSequence(this); };
	void namespaceBinding(const QXmlName & nb) { callbackQXmlSerializer_NamespaceBinding(this, const_cast<QXmlName*>(&nb)); };
};

void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice)
{
	return new MyQXmlFormatter(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlFormatter_SetIndentationDepth(void* ptr, int depth)
{
	static_cast<QXmlFormatter*>(ptr)->setIndentationDepth(depth);
}

int QXmlFormatter_IndentationDepth(void* ptr)
{
	return static_cast<QXmlFormatter*>(ptr)->indentationDepth();
}

void* QXmlItem_NewQXmlItem()
{
	return new QXmlItem();
}

void* QXmlItem_NewQXmlItem4(void* atomicValue)
{
	return new QXmlItem(*static_cast<QVariant*>(atomicValue));
}

void* QXmlItem_NewQXmlItem2(void* other)
{
	return new QXmlItem(*static_cast<QXmlItem*>(other));
}

void* QXmlItem_NewQXmlItem3(void* node)
{
	return new QXmlItem(*static_cast<QXmlNodeModelIndex*>(node));
}

void QXmlItem_DestroyQXmlItem(void* ptr)
{
	static_cast<QXmlItem*>(ptr)->~QXmlItem();
}

void* QXmlItem_ToAtomicValue(void* ptr)
{
	return new QVariant(static_cast<QXmlItem*>(ptr)->toAtomicValue());
}

void* QXmlItem_ToNodeModelIndex(void* ptr)
{
	return new QXmlNodeModelIndex(static_cast<QXmlItem*>(ptr)->toNodeModelIndex());
}

char QXmlItem_IsAtomicValue(void* ptr)
{
	return static_cast<QXmlItem*>(ptr)->isAtomicValue();
}

char QXmlItem_IsNode(void* ptr)
{
	return static_cast<QXmlItem*>(ptr)->isNode();
}

char QXmlItem_IsNull(void* ptr)
{
	return static_cast<QXmlItem*>(ptr)->isNull();
}

void* QXmlName_QXmlName_FromClarkName(struct QtXmlPatterns_PackedString clarkName, void* namePool)
{
	return new QXmlName(QXmlName::fromClarkName(QString::fromUtf8(clarkName.data, clarkName.len), *static_cast<QXmlNamePool*>(namePool)));
}

void* QXmlName_NewQXmlName()
{
	return new QXmlName();
}

void* QXmlName_NewQXmlName2(void* namePool, struct QtXmlPatterns_PackedString localName, struct QtXmlPatterns_PackedString namespaceURI, struct QtXmlPatterns_PackedString prefix)
{
	return new QXmlName(*static_cast<QXmlNamePool*>(namePool), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(prefix.data, prefix.len));
}

void* QXmlName_NewQXmlName3(void* other)
{
	return new QXmlName(*static_cast<QXmlName*>(other));
}

char QXmlName_QXmlName_IsNCName(struct QtXmlPatterns_PackedString candidate)
{
	return QXmlName::isNCName(QString::fromUtf8(candidate.data, candidate.len));
}

struct QtXmlPatterns_PackedString QXmlName_LocalName(void* ptr, void* namePool)
{
	return ({ QByteArray tf2a5c3 = static_cast<QXmlName*>(ptr)->localName(*static_cast<QXmlNamePool*>(namePool)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(tf2a5c3.prepend("WHITESPACE").constData()+10), tf2a5c3.size()-10 }; });
}

struct QtXmlPatterns_PackedString QXmlName_NamespaceUri(void* ptr, void* namePool)
{
	return ({ QByteArray tbb7dfe = static_cast<QXmlName*>(ptr)->namespaceUri(*static_cast<QXmlNamePool*>(namePool)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(tbb7dfe.prepend("WHITESPACE").constData()+10), tbb7dfe.size()-10 }; });
}

struct QtXmlPatterns_PackedString QXmlName_Prefix(void* ptr, void* namePool)
{
	return ({ QByteArray t451eec = static_cast<QXmlName*>(ptr)->prefix(*static_cast<QXmlNamePool*>(namePool)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(t451eec.prepend("WHITESPACE").constData()+10), t451eec.size()-10 }; });
}

struct QtXmlPatterns_PackedString QXmlName_ToClarkName(void* ptr, void* namePool)
{
	return ({ QByteArray t611ba4 = static_cast<QXmlName*>(ptr)->toClarkName(*static_cast<QXmlNamePool*>(namePool)).toUtf8(); QtXmlPatterns_PackedString { const_cast<char*>(t611ba4.prepend("WHITESPACE").constData()+10), t611ba4.size()-10 }; });
}

char QXmlName_IsNull(void* ptr)
{
	return static_cast<QXmlName*>(ptr)->isNull();
}

void* QXmlNamePool_NewQXmlNamePool()
{
	return new QXmlNamePool();
}

void* QXmlNamePool_NewQXmlNamePool2(void* other)
{
	return new QXmlNamePool(*static_cast<QXmlNamePool*>(other));
}

void QXmlNamePool_DestroyQXmlNamePool(void* ptr)
{
	static_cast<QXmlNamePool*>(ptr)->~QXmlNamePool();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex()
{
	return new QXmlNodeModelIndex();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other)
{
	return new QXmlNodeModelIndex(*static_cast<QXmlNodeModelIndex*>(other));
}

char QXmlNodeModelIndex_IsNull(void* ptr)
{
	return static_cast<QXmlNodeModelIndex*>(ptr)->isNull();
}

void* QXmlNodeModelIndex_Model(void* ptr)
{
	return const_cast<QAbstractXmlNodeModel*>(static_cast<QXmlNodeModelIndex*>(ptr)->model());
}

long long QXmlNodeModelIndex_AdditionalData(void* ptr)
{
	return static_cast<QXmlNodeModelIndex*>(ptr)->additionalData();
}

long long QXmlNodeModelIndex_Data(void* ptr)
{
	return static_cast<QXmlNodeModelIndex*>(ptr)->data();
}

void* QXmlNodeModelIndex_InternalPointer(void* ptr)
{
	return static_cast<QXmlNodeModelIndex*>(ptr)->internalPointer();
}

void* QXmlNodeModelIndex___namespaceBindings_atList(void* ptr, int i)
{
	return new QXmlName(static_cast<QVector<QXmlName>*>(ptr)->at(i));
}

void QXmlNodeModelIndex___namespaceBindings_setList(void* ptr, void* i)
{
	static_cast<QVector<QXmlName>*>(ptr)->append(*static_cast<QXmlName*>(i));
}

void* QXmlNodeModelIndex___namespaceBindings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QXmlName>;
}

void* QXmlQuery_NewQXmlQuery()
{
	return new QXmlQuery();
}

void* QXmlQuery_NewQXmlQuery4(long long queryLanguage, void* np)
{
	return new QXmlQuery(static_cast<QXmlQuery::QueryLanguage>(queryLanguage), *static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery3(void* np)
{
	return new QXmlQuery(*static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery2(void* other)
{
	return new QXmlQuery(*static_cast<QXmlQuery*>(other));
}

char QXmlQuery_SetFocus3(void* ptr, void* document)
{
	return static_cast<QXmlQuery*>(ptr)->setFocus(static_cast<QIODevice*>(document));
}

char QXmlQuery_SetFocus4(void* ptr, struct QtXmlPatterns_PackedString focus)
{
	return static_cast<QXmlQuery*>(ptr)->setFocus(QString::fromUtf8(focus.data, focus.len));
}

char QXmlQuery_SetFocus2(void* ptr, void* documentURI)
{
	return static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QUrl*>(documentURI));
}

void QXmlQuery_BindVariable4(void* ptr, struct QtXmlPatterns_PackedString localName, void* device)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString::fromUtf8(localName.data, localName.len), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable2(void* ptr, struct QtXmlPatterns_PackedString localName, void* value)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString::fromUtf8(localName.data, localName.len), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable6(void* ptr, struct QtXmlPatterns_PackedString localName, void* query)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString::fromUtf8(localName.data, localName.len), *static_cast<QXmlQuery*>(query));
}

void QXmlQuery_BindVariable3(void* ptr, void* name, void* device)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable(void* ptr, void* name, void* value)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable5(void* ptr, void* name, void* query)
{
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlQuery*>(query));
}

void QXmlQuery_SetFocus(void* ptr, void* item)
{
	static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QXmlItem*>(item));
}

void QXmlQuery_SetInitialTemplateName2(void* ptr, struct QtXmlPatterns_PackedString localName)
{
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(QString::fromUtf8(localName.data, localName.len));
}

void QXmlQuery_SetInitialTemplateName(void* ptr, void* name)
{
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(*static_cast<QXmlName*>(name));
}

void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler)
{
	static_cast<QXmlQuery*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(aMessageHandler));
}

void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager)
{
	static_cast<QXmlQuery*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(newManager));
}

void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI)
{
	static_cast<QXmlQuery*>(ptr)->setQuery(static_cast<QIODevice*>(sourceCode), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery2(void* ptr, struct QtXmlPatterns_PackedString sourceCode, void* documentURI)
{
	static_cast<QXmlQuery*>(ptr)->setQuery(QString::fromUtf8(sourceCode.data, sourceCode.len), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery3(void* ptr, void* queryURI, void* baseURI)
{
	static_cast<QXmlQuery*>(ptr)->setQuery(*static_cast<QUrl*>(queryURI), *static_cast<QUrl*>(baseURI));
}

void QXmlQuery_SetUriResolver(void* ptr, void* resolver)
{
	static_cast<QXmlQuery*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void QXmlQuery_DestroyQXmlQuery(void* ptr)
{
	static_cast<QXmlQuery*>(ptr)->~QXmlQuery();
}

void* QXmlQuery_MessageHandler(void* ptr)
{
	return static_cast<QXmlQuery*>(ptr)->messageHandler();
}

void* QXmlQuery_NetworkAccessManager(void* ptr)
{
	return static_cast<QXmlQuery*>(ptr)->networkAccessManager();
}

void* QXmlQuery_InitialTemplateName(void* ptr)
{
	return new QXmlName(static_cast<QXmlQuery*>(ptr)->initialTemplateName());
}

void* QXmlQuery_NamePool(void* ptr)
{
	return new QXmlNamePool(static_cast<QXmlQuery*>(ptr)->namePool());
}

long long QXmlQuery_QueryLanguage(void* ptr)
{
	return static_cast<QXmlQuery*>(ptr)->queryLanguage();
}

char QXmlQuery_EvaluateTo2(void* ptr, void* callback)
{
	return static_cast<QXmlQuery*>(ptr)->evaluateTo(static_cast<QAbstractXmlReceiver*>(callback));
}

char QXmlQuery_EvaluateTo4(void* ptr, void* target)
{
	return static_cast<QXmlQuery*>(ptr)->evaluateTo(static_cast<QIODevice*>(target));
}

char QXmlQuery_EvaluateTo5(void* ptr, struct QtXmlPatterns_PackedString output)
{
	return static_cast<QXmlQuery*>(ptr)->evaluateTo(new QString(QString::fromUtf8(output.data, output.len)));
}

char QXmlQuery_EvaluateTo3(void* ptr, struct QtXmlPatterns_PackedString target)
{
	return static_cast<QXmlQuery*>(ptr)->evaluateTo(new QStringList(QString::fromUtf8(target.data, target.len).split("|", QString::SkipEmptyParts)));
}

char QXmlQuery_IsValid(void* ptr)
{
	return static_cast<QXmlQuery*>(ptr)->isValid();
}

void* QXmlQuery_UriResolver(void* ptr)
{
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlQuery*>(ptr)->uriResolver());
}

void QXmlQuery_EvaluateTo(void* ptr, void* result)
{
	static_cast<QXmlQuery*>(ptr)->evaluateTo(static_cast<QXmlResultItems*>(result));
}

class MyQXmlResultItems: public QXmlResultItems
{
public:
	MyQXmlResultItems() : QXmlResultItems() {};
	 ~MyQXmlResultItems() { callbackQXmlResultItems_DestroyQXmlResultItems(this); };
};

void* QXmlResultItems_Next(void* ptr)
{
	return new QXmlItem(static_cast<QXmlResultItems*>(ptr)->next());
}

void* QXmlResultItems_NewQXmlResultItems()
{
	return new MyQXmlResultItems();
}

void QXmlResultItems_DestroyQXmlResultItems(void* ptr)
{
	static_cast<QXmlResultItems*>(ptr)->~QXmlResultItems();
}

void QXmlResultItems_DestroyQXmlResultItemsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QXmlResultItems_Current(void* ptr)
{
	return new QXmlItem(static_cast<QXmlResultItems*>(ptr)->current());
}

char QXmlResultItems_HasError(void* ptr)
{
	return static_cast<QXmlResultItems*>(ptr)->hasError();
}

void* QXmlSchema_NewQXmlSchema()
{
	return new QXmlSchema();
}

void* QXmlSchema_NewQXmlSchema2(void* other)
{
	return new QXmlSchema(*static_cast<QXmlSchema*>(other));
}

char QXmlSchema_Load2(void* ptr, void* source, void* documentUri)
{
	return static_cast<QXmlSchema*>(ptr)->load(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

char QXmlSchema_Load3(void* ptr, void* data, void* documentUri)
{
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

char QXmlSchema_Load(void* ptr, void* source)
{
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QUrl*>(source));
}

void QXmlSchema_SetMessageHandler(void* ptr, void* handler)
{
	static_cast<QXmlSchema*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager)
{
	static_cast<QXmlSchema*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchema_SetUriResolver(void* ptr, void* resolver)
{
	static_cast<QXmlSchema*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void QXmlSchema_DestroyQXmlSchema(void* ptr)
{
	static_cast<QXmlSchema*>(ptr)->~QXmlSchema();
}

void* QXmlSchema_MessageHandler(void* ptr)
{
	return static_cast<QXmlSchema*>(ptr)->messageHandler();
}

void* QXmlSchema_NetworkAccessManager(void* ptr)
{
	return static_cast<QXmlSchema*>(ptr)->networkAccessManager();
}

void* QXmlSchema_DocumentUri(void* ptr)
{
	return new QUrl(static_cast<QXmlSchema*>(ptr)->documentUri());
}

void* QXmlSchema_NamePool(void* ptr)
{
	return new QXmlNamePool(static_cast<QXmlSchema*>(ptr)->namePool());
}

char QXmlSchema_IsValid(void* ptr)
{
	return static_cast<QXmlSchema*>(ptr)->isValid();
}

void* QXmlSchema_UriResolver(void* ptr)
{
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchema*>(ptr)->uriResolver());
}

void* QXmlSchemaValidator_NewQXmlSchemaValidator()
{
	return new QXmlSchemaValidator();
}

void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema)
{
	return new QXmlSchemaValidator(*static_cast<QXmlSchema*>(schema));
}

void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler)
{
	static_cast<QXmlSchemaValidator*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager)
{
	static_cast<QXmlSchemaValidator*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchemaValidator_SetSchema(void* ptr, void* schema)
{
	static_cast<QXmlSchemaValidator*>(ptr)->setSchema(*static_cast<QXmlSchema*>(schema));
}

void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver)
{
	static_cast<QXmlSchemaValidator*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr)
{
	static_cast<QXmlSchemaValidator*>(ptr)->~QXmlSchemaValidator();
}

void* QXmlSchemaValidator_MessageHandler(void* ptr)
{
	return static_cast<QXmlSchemaValidator*>(ptr)->messageHandler();
}

void* QXmlSchemaValidator_NetworkAccessManager(void* ptr)
{
	return static_cast<QXmlSchemaValidator*>(ptr)->networkAccessManager();
}

void* QXmlSchemaValidator_NamePool(void* ptr)
{
	return new QXmlNamePool(static_cast<QXmlSchemaValidator*>(ptr)->namePool());
}

void* QXmlSchemaValidator_Schema(void* ptr)
{
	return new QXmlSchema(static_cast<QXmlSchemaValidator*>(ptr)->schema());
}

char QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri)
{
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

char QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri)
{
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

char QXmlSchemaValidator_Validate(void* ptr, void* source)
{
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QUrl*>(source));
}

void* QXmlSchemaValidator_UriResolver(void* ptr)
{
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchemaValidator*>(ptr)->uriResolver());
}

class MyQXmlSerializer: public QXmlSerializer
{
public:
	MyQXmlSerializer(const QXmlQuery &query, QIODevice *outputDevice) : QXmlSerializer(query, outputDevice) {};
	void atomicValue(const QVariant & value) { callbackQXmlSerializer_AtomicValue(this, const_cast<QVariant*>(&value)); };
	void attribute(const QXmlName & name, const QStringRef & value) { callbackQXmlSerializer_Attribute(this, const_cast<QXmlName*>(&name), const_cast<QStringRef*>(&value)); };
	void characters(const QStringRef & value) { callbackQXmlSerializer_Characters(this, const_cast<QStringRef*>(&value)); };
	void comment(const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQXmlSerializer_Comment(this, valuePacked); };
	void endDocument() { callbackQXmlSerializer_EndDocument(this); };
	void endElement() { callbackQXmlSerializer_EndElement(this); };
	void endOfSequence() { callbackQXmlSerializer_EndOfSequence(this); };
	void namespaceBinding(const QXmlName & nb) { callbackQXmlSerializer_NamespaceBinding(this, const_cast<QXmlName*>(&nb)); };
	void processingInstruction(const QXmlName & name, const QString & value) { QByteArray tf32b67 = value.toUtf8(); QtXmlPatterns_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQXmlSerializer_ProcessingInstruction(this, const_cast<QXmlName*>(&name), valuePacked); };
	void startDocument() { callbackQXmlSerializer_StartDocument(this); };
	void startElement(const QXmlName & name) { callbackQXmlSerializer_StartElement(this, const_cast<QXmlName*>(&name)); };
	void startOfSequence() { callbackQXmlSerializer_StartOfSequence(this); };
};

void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice)
{
	return new MyQXmlSerializer(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlSerializer_AtomicValue(void* ptr, void* value)
{
	static_cast<QXmlSerializer*>(ptr)->atomicValue(*static_cast<QVariant*>(value));
}

void QXmlSerializer_AtomicValueDefault(void* ptr, void* value)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::atomicValue(*static_cast<QVariant*>(value));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::atomicValue(*static_cast<QVariant*>(value));
	}
}

void QXmlSerializer_Attribute(void* ptr, void* name, void* value)
{
	static_cast<QXmlSerializer*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlSerializer_AttributeDefault(void* ptr, void* name, void* value)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
	}
}

void QXmlSerializer_Characters(void* ptr, void* value)
{
	static_cast<QXmlSerializer*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_CharactersDefault(void* ptr, void* value)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::characters(*static_cast<QStringRef*>(value));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::characters(*static_cast<QStringRef*>(value));
	}
}

void QXmlSerializer_Comment(void* ptr, struct QtXmlPatterns_PackedString value)
{
	static_cast<QXmlSerializer*>(ptr)->comment(QString::fromUtf8(value.data, value.len));
}

void QXmlSerializer_CommentDefault(void* ptr, struct QtXmlPatterns_PackedString value)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::comment(QString::fromUtf8(value.data, value.len));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::comment(QString::fromUtf8(value.data, value.len));
	}
}

void QXmlSerializer_EndDocument(void* ptr)
{
	static_cast<QXmlSerializer*>(ptr)->endDocument();
}

void QXmlSerializer_EndDocumentDefault(void* ptr)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endDocument();
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endDocument();
	}
}

void QXmlSerializer_EndElement(void* ptr)
{
	static_cast<QXmlSerializer*>(ptr)->endElement();
}

void QXmlSerializer_EndElementDefault(void* ptr)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endElement();
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endElement();
	}
}

void QXmlSerializer_EndOfSequence(void* ptr)
{
	static_cast<QXmlSerializer*>(ptr)->endOfSequence();
}

void QXmlSerializer_EndOfSequenceDefault(void* ptr)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endOfSequence();
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endOfSequence();
	}
}

void QXmlSerializer_NamespaceBinding(void* ptr, void* nb)
{
	static_cast<QXmlSerializer*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(nb));
}

void QXmlSerializer_NamespaceBindingDefault(void* ptr, void* nb)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::namespaceBinding(*static_cast<QXmlName*>(nb));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::namespaceBinding(*static_cast<QXmlName*>(nb));
	}
}

void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, struct QtXmlPatterns_PackedString value)
{
	static_cast<QXmlSerializer*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString::fromUtf8(value.data, value.len));
}

void QXmlSerializer_ProcessingInstructionDefault(void* ptr, void* name, struct QtXmlPatterns_PackedString value)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::processingInstruction(*static_cast<QXmlName*>(name), QString::fromUtf8(value.data, value.len));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::processingInstruction(*static_cast<QXmlName*>(name), QString::fromUtf8(value.data, value.len));
	}
}

void QXmlSerializer_SetCodec(void* ptr, void* outputCodec)
{
	static_cast<QXmlSerializer*>(ptr)->setCodec(static_cast<QTextCodec*>(outputCodec));
}

void QXmlSerializer_StartDocument(void* ptr)
{
	static_cast<QXmlSerializer*>(ptr)->startDocument();
}

void QXmlSerializer_StartDocumentDefault(void* ptr)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::startDocument();
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::startDocument();
	}
}

void QXmlSerializer_StartElement(void* ptr, void* name)
{
	static_cast<QXmlSerializer*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlSerializer_StartElementDefault(void* ptr, void* name)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::startElement(*static_cast<QXmlName*>(name));
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::startElement(*static_cast<QXmlName*>(name));
	}
}

void QXmlSerializer_StartOfSequence(void* ptr)
{
	static_cast<QXmlSerializer*>(ptr)->startOfSequence();
}

void QXmlSerializer_StartOfSequenceDefault(void* ptr)
{
	if (dynamic_cast<QXmlFormatter*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::startOfSequence();
	} else {
		static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::startOfSequence();
	}
}

void* QXmlSerializer_OutputDevice(void* ptr)
{
	return static_cast<QXmlSerializer*>(ptr)->outputDevice();
}

void* QXmlSerializer_Codec(void* ptr)
{
	return const_cast<QTextCodec*>(static_cast<QXmlSerializer*>(ptr)->codec());
}

