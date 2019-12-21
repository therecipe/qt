// +build !minimal

#define protected public
#define private public

#include "virtualkeyboard.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QInputMethodEvent>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QPointF>
#include <QQuickItem>
#include <QRadioData>
#include <QRectF>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QVirtualKeyboardAbstractInputMethod>
#include <QVirtualKeyboardExtensionPlugin>
#include <QVirtualKeyboardInputContext>
#include <QVirtualKeyboardInputEngine>
#include <QVirtualKeyboardSelectionListModel>
#include <QVirtualKeyboardTrace>
#include <QWidget>
#include <QWindow>

class MyQVirtualKeyboardAbstractInputMethod: public QVirtualKeyboardAbstractInputMethod
{
public:
	MyQVirtualKeyboardAbstractInputMethod(QObject *parent = Q_NULLPTR) : QVirtualKeyboardAbstractInputMethod(parent) {QVirtualKeyboardAbstractInputMethod_QVirtualKeyboardAbstractInputMethod_QRegisterMetaType();};
	bool clickPreeditText(int cursorPosition) { return callbackQVirtualKeyboardAbstractInputMethod_ClickPreeditText(this, cursorPosition) != 0; };
	QList<QVirtualKeyboardInputEngine::InputMode> inputModes(const QString & locale) { QByteArray* t0e038a = new QByteArray(locale.toUtf8()); QtVirtualKeyboard_PackedString localePacked = { const_cast<char*>(t0e038a->prepend("WHITESPACE").constData()+10), t0e038a->size()-10, t0e038a };return ({ QList<QVirtualKeyboardInputEngine::InputMode>* tmpP = static_cast<QList<QVirtualKeyboardInputEngine::InputMode>*>(callbackQVirtualKeyboardAbstractInputMethod_InputModes(this, localePacked)); QList<QVirtualKeyboardInputEngine::InputMode> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	bool keyEvent(Qt::Key key, const QString & text, Qt::KeyboardModifiers modifiers) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtVirtualKeyboard_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };return callbackQVirtualKeyboardAbstractInputMethod_KeyEvent(this, static_cast<qint64>(key), textPacked, static_cast<qint64>(modifiers)) != 0; };
	QList<QVirtualKeyboardInputEngine::PatternRecognitionMode> patternRecognitionModes() const { return ({ QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>* tmpP = static_cast<QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>*>(callbackQVirtualKeyboardAbstractInputMethod_PatternRecognitionModes(const_cast<void*>(static_cast<const void*>(this)))); QList<QVirtualKeyboardInputEngine::PatternRecognitionMode> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	bool reselect(int cursorPosition, const QVirtualKeyboardInputEngine::ReselectFlags & reselectFlags) { return callbackQVirtualKeyboardAbstractInputMethod_Reselect(this, cursorPosition, static_cast<qint64>(reselectFlags)) != 0; };
	void reset() { callbackQVirtualKeyboardAbstractInputMethod_Reset(this); };
	void Signal_SelectionListActiveItemChanged(QVirtualKeyboardSelectionListModel::Type ty, int index) { callbackQVirtualKeyboardAbstractInputMethod_SelectionListActiveItemChanged(this, static_cast<qint64>(ty), index); };
	void Signal_SelectionListChanged(QVirtualKeyboardSelectionListModel::Type ty) { callbackQVirtualKeyboardAbstractInputMethod_SelectionListChanged(this, static_cast<qint64>(ty)); };
	QVariant selectionListData(QVirtualKeyboardSelectionListModel::Type ty, int index, QVirtualKeyboardSelectionListModel::Role role) { return *static_cast<QVariant*>(callbackQVirtualKeyboardAbstractInputMethod_SelectionListData(this, static_cast<qint64>(ty), index, static_cast<qint64>(role))); };
	int selectionListItemCount(QVirtualKeyboardSelectionListModel::Type ty) { return callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemCount(this, static_cast<qint64>(ty)); };
	void selectionListItemSelected(QVirtualKeyboardSelectionListModel::Type ty, int index) { callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemSelected(this, static_cast<qint64>(ty), index); };
	bool selectionListRemoveItem(QVirtualKeyboardSelectionListModel::Type ty, int index) { return callbackQVirtualKeyboardAbstractInputMethod_SelectionListRemoveItem(this, static_cast<qint64>(ty), index) != 0; };
	QList<QVirtualKeyboardSelectionListModel::Type> selectionLists() { return ({ QList<QVirtualKeyboardSelectionListModel::Type>* tmpP = static_cast<QList<QVirtualKeyboardSelectionListModel::Type>*>(callbackQVirtualKeyboardAbstractInputMethod_SelectionLists(this)); QList<QVirtualKeyboardSelectionListModel::Type> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void Signal_SelectionListsChanged() { callbackQVirtualKeyboardAbstractInputMethod_SelectionListsChanged(this); };
	bool setInputMode(const QString & locale, QVirtualKeyboardInputEngine::InputMode inputMode) { QByteArray* t0e038a = new QByteArray(locale.toUtf8()); QtVirtualKeyboard_PackedString localePacked = { const_cast<char*>(t0e038a->prepend("WHITESPACE").constData()+10), t0e038a->size()-10, t0e038a };return callbackQVirtualKeyboardAbstractInputMethod_SetInputMode(this, localePacked, static_cast<qint64>(inputMode)) != 0; };
	bool setTextCase(QVirtualKeyboardInputEngine::TextCase textCase) { return callbackQVirtualKeyboardAbstractInputMethod_SetTextCase(this, static_cast<qint64>(textCase)) != 0; };
	QVirtualKeyboardTrace * traceBegin(int traceId, QVirtualKeyboardInputEngine::PatternRecognitionMode patternRecognitionMode, const QVariantMap & traceCaptureDeviceInfo, const QVariantMap & traceScreenInfo) { return static_cast<QVirtualKeyboardTrace*>(callbackQVirtualKeyboardAbstractInputMethod_TraceBegin(this, traceId, static_cast<qint64>(patternRecognitionMode), ({ QMap<QString, QVariant>* tmpValuef1ecb3 = new QMap<QString, QVariant>(traceCaptureDeviceInfo); QtVirtualKeyboard_PackedList { tmpValuef1ecb3, tmpValuef1ecb3->size() }; }), ({ QMap<QString, QVariant>* tmpValue56131b = new QMap<QString, QVariant>(traceScreenInfo); QtVirtualKeyboard_PackedList { tmpValue56131b, tmpValue56131b->size() }; }))); };
	bool traceEnd(QVirtualKeyboardTrace * trace) { return callbackQVirtualKeyboardAbstractInputMethod_TraceEnd(this, trace) != 0; };
	void update() { callbackQVirtualKeyboardAbstractInputMethod_Update(this); };
	 ~MyQVirtualKeyboardAbstractInputMethod() { callbackQVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethod(this); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardAbstractInputMethod_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardAbstractInputMethod_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardAbstractInputMethod_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardAbstractInputMethod_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardAbstractInputMethod_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardAbstractInputMethod_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardAbstractInputMethod_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardAbstractInputMethod_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardAbstractInputMethod_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardAbstractInputMethod_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardAbstractInputMethod_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardAbstractInputMethod*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardAbstractInputMethod*)

int QVirtualKeyboardAbstractInputMethod_QVirtualKeyboardAbstractInputMethod_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardAbstractInputMethod*>(); return qRegisterMetaType<MyQVirtualKeyboardAbstractInputMethod*>();}

void* QVirtualKeyboardAbstractInputMethod_NewQVirtualKeyboardAbstractInputMethod(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QWindow*>(parent));
	} else {
		return new MyQVirtualKeyboardAbstractInputMethod(static_cast<QObject*>(parent));
	}
}

char QVirtualKeyboardAbstractInputMethod_ClickPreeditText(void* ptr, int cursorPosition)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->clickPreeditText(cursorPosition);
}

char QVirtualKeyboardAbstractInputMethod_ClickPreeditTextDefault(void* ptr, int cursorPosition)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::clickPreeditText(cursorPosition);
}

void* QVirtualKeyboardAbstractInputMethod_InputContext(void* ptr)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->inputContext();
}

void* QVirtualKeyboardAbstractInputMethod_InputEngine(void* ptr)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->inputEngine();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod_InputModes(void* ptr, struct QtVirtualKeyboard_PackedString locale)
{
	return ({ QList<QVirtualKeyboardInputEngine::InputMode>* tmpValue805913 = new QList<QVirtualKeyboardInputEngine::InputMode>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->inputModes(QString::fromUtf8(locale.data, locale.len))); QtVirtualKeyboard_PackedList { tmpValue805913, tmpValue805913->size() }; });
}

char QVirtualKeyboardAbstractInputMethod_KeyEvent(void* ptr, long long key, struct QtVirtualKeyboard_PackedString text, long long modifiers)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->keyEvent(static_cast<Qt::Key>(key), QString::fromUtf8(text.data, text.len), static_cast<Qt::KeyboardModifier>(modifiers));
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod_PatternRecognitionModes(void* ptr)
{
	return ({ QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>* tmpValue9d8704 = new QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->patternRecognitionModes()); QtVirtualKeyboard_PackedList { tmpValue9d8704, tmpValue9d8704->size() }; });
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod_PatternRecognitionModesDefault(void* ptr)
{
		return ({ QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>* tmpValuef27220 = new QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::patternRecognitionModes()); QtVirtualKeyboard_PackedList { tmpValuef27220, tmpValuef27220->size() }; });
}

char QVirtualKeyboardAbstractInputMethod_Reselect(void* ptr, int cursorPosition, long long reselectFlags)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->reselect(cursorPosition, static_cast<QVirtualKeyboardInputEngine::ReselectFlag>(reselectFlags));
}

char QVirtualKeyboardAbstractInputMethod_ReselectDefault(void* ptr, int cursorPosition, long long reselectFlags)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::reselect(cursorPosition, static_cast<QVirtualKeyboardInputEngine::ReselectFlag>(reselectFlags));
}

void QVirtualKeyboardAbstractInputMethod_Reset(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), "reset");
}

void QVirtualKeyboardAbstractInputMethod_ResetDefault(void* ptr)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::reset();
}

void QVirtualKeyboardAbstractInputMethod_ConnectSelectionListActiveItemChanged(void* ptr, long long t)
{
	qRegisterMetaType<QVirtualKeyboardSelectionListModel::Type>();
	QObject::connect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type, int)>(&QVirtualKeyboardAbstractInputMethod::selectionListActiveItemChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type, int)>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListActiveItemChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListActiveItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type, int)>(&QVirtualKeyboardAbstractInputMethod::selectionListActiveItemChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type, int)>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListActiveItemChanged));
}

void QVirtualKeyboardAbstractInputMethod_SelectionListActiveItemChanged(void* ptr, long long ty, int index)
{
	static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListActiveItemChanged(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index);
}

void QVirtualKeyboardAbstractInputMethod_ConnectSelectionListChanged(void* ptr, long long t)
{
	qRegisterMetaType<QVirtualKeyboardSelectionListModel::Type>();
	QObject::connect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type)>(&QVirtualKeyboardAbstractInputMethod::selectionListChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type)>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type)>(&QVirtualKeyboardAbstractInputMethod::selectionListChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)(QVirtualKeyboardSelectionListModel::Type)>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListChanged));
}

void QVirtualKeyboardAbstractInputMethod_SelectionListChanged(void* ptr, long long ty)
{
	static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListChanged(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty));
}

void* QVirtualKeyboardAbstractInputMethod_SelectionListData(void* ptr, long long ty, int index, long long role)
{
	return new QVariant(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListData(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index, static_cast<QVirtualKeyboardSelectionListModel::Role>(role)));
}

void* QVirtualKeyboardAbstractInputMethod_SelectionListDataDefault(void* ptr, long long ty, int index, long long role)
{
		return new QVariant(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::selectionListData(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index, static_cast<QVirtualKeyboardSelectionListModel::Role>(role)));
}

int QVirtualKeyboardAbstractInputMethod_SelectionListItemCount(void* ptr, long long ty)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListItemCount(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty));
}

int QVirtualKeyboardAbstractInputMethod_SelectionListItemCountDefault(void* ptr, long long ty)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::selectionListItemCount(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty));
}

void QVirtualKeyboardAbstractInputMethod_SelectionListItemSelected(void* ptr, long long ty, int index)
{
	static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListItemSelected(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index);
}

void QVirtualKeyboardAbstractInputMethod_SelectionListItemSelectedDefault(void* ptr, long long ty, int index)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::selectionListItemSelected(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index);
}

char QVirtualKeyboardAbstractInputMethod_SelectionListRemoveItem(void* ptr, long long ty, int index)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListRemoveItem(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index);
}

char QVirtualKeyboardAbstractInputMethod_SelectionListRemoveItemDefault(void* ptr, long long ty, int index)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::selectionListRemoveItem(static_cast<QVirtualKeyboardSelectionListModel::Type>(ty), index);
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod_SelectionLists(void* ptr)
{
	return ({ QList<QVirtualKeyboardSelectionListModel::Type>* tmpValue1a2bc4 = new QList<QVirtualKeyboardSelectionListModel::Type>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionLists()); QtVirtualKeyboard_PackedList { tmpValue1a2bc4, tmpValue1a2bc4->size() }; });
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod_SelectionListsDefault(void* ptr)
{
		return ({ QList<QVirtualKeyboardSelectionListModel::Type>* tmpValue68d5c5 = new QList<QVirtualKeyboardSelectionListModel::Type>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::selectionLists()); QtVirtualKeyboard_PackedList { tmpValue68d5c5, tmpValue68d5c5->size() }; });
}

void QVirtualKeyboardAbstractInputMethod_ConnectSelectionListsChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)()>(&QVirtualKeyboardAbstractInputMethod::selectionListsChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)()>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListsChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (QVirtualKeyboardAbstractInputMethod::*)()>(&QVirtualKeyboardAbstractInputMethod::selectionListsChanged), static_cast<MyQVirtualKeyboardAbstractInputMethod*>(ptr), static_cast<void (MyQVirtualKeyboardAbstractInputMethod::*)()>(&MyQVirtualKeyboardAbstractInputMethod::Signal_SelectionListsChanged));
}

void QVirtualKeyboardAbstractInputMethod_SelectionListsChanged(void* ptr)
{
	static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->selectionListsChanged();
}

char QVirtualKeyboardAbstractInputMethod_SetInputMode(void* ptr, struct QtVirtualKeyboard_PackedString locale, long long inputMode)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->setInputMode(QString::fromUtf8(locale.data, locale.len), static_cast<QVirtualKeyboardInputEngine::InputMode>(inputMode));
}

char QVirtualKeyboardAbstractInputMethod_SetTextCase(void* ptr, long long textCase)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->setTextCase(static_cast<QVirtualKeyboardInputEngine::TextCase>(textCase));
}

void* QVirtualKeyboardAbstractInputMethod_TraceBegin(void* ptr, int traceId, long long patternRecognitionMode, void* traceCaptureDeviceInfo, void* traceScreenInfo)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->traceBegin(traceId, static_cast<QVirtualKeyboardInputEngine::PatternRecognitionMode>(patternRecognitionMode), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceCaptureDeviceInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceScreenInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void* QVirtualKeyboardAbstractInputMethod_TraceBeginDefault(void* ptr, int traceId, long long patternRecognitionMode, void* traceCaptureDeviceInfo, void* traceScreenInfo)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::traceBegin(traceId, static_cast<QVirtualKeyboardInputEngine::PatternRecognitionMode>(patternRecognitionMode), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceCaptureDeviceInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceScreenInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

char QVirtualKeyboardAbstractInputMethod_TraceEnd(void* ptr, void* trace)
{
	return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->traceEnd(static_cast<QVirtualKeyboardTrace*>(trace));
}

char QVirtualKeyboardAbstractInputMethod_TraceEndDefault(void* ptr, void* trace)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::traceEnd(static_cast<QVirtualKeyboardTrace*>(trace));
}

void QVirtualKeyboardAbstractInputMethod_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr), "update");
}

void QVirtualKeyboardAbstractInputMethod_UpdateDefault(void* ptr)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::update();
}

void QVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethod(void* ptr)
{
	static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->~QVirtualKeyboardAbstractInputMethod();
}

void QVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethodDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QVirtualKeyboardAbstractInputMethod___inputModes_atList(void* ptr, int i)
{
	return static_cast<qint64>(({QVirtualKeyboardInputEngine::InputMode tmp = static_cast<QList<QVirtualKeyboardInputEngine::InputMode>*>(ptr)->at(i); if (i == static_cast<QList<QVirtualKeyboardInputEngine::InputMode>*>(ptr)->size()-1) { static_cast<QList<QVirtualKeyboardInputEngine::InputMode>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___inputModes_setList(void* ptr, long long i)
{
	static_cast<QList<QVirtualKeyboardInputEngine::InputMode>*>(ptr)->append(static_cast<QVirtualKeyboardInputEngine::InputMode>(i));
}

void* QVirtualKeyboardAbstractInputMethod___inputModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVirtualKeyboardInputEngine::InputMode>();
}

long long QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_atList(void* ptr, int i)
{
	return static_cast<qint64>(({QVirtualKeyboardInputEngine::PatternRecognitionMode tmp = static_cast<QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>*>(ptr)->at(i); if (i == static_cast<QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>*>(ptr)->size()-1) { static_cast<QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_setList(void* ptr, long long i)
{
	static_cast<QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>*>(ptr)->append(static_cast<QVirtualKeyboardInputEngine::PatternRecognitionMode>(i));
}

void* QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVirtualKeyboardInputEngine::PatternRecognitionMode>();
}

long long QVirtualKeyboardAbstractInputMethod___selectionLists_atList(void* ptr, int i)
{
	return static_cast<qint64>(({QVirtualKeyboardSelectionListModel::Type tmp = static_cast<QList<QVirtualKeyboardSelectionListModel::Type>*>(ptr)->at(i); if (i == static_cast<QList<QVirtualKeyboardSelectionListModel::Type>*>(ptr)->size()-1) { static_cast<QList<QVirtualKeyboardSelectionListModel::Type>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___selectionLists_setList(void* ptr, long long i)
{
	static_cast<QList<QVirtualKeyboardSelectionListModel::Type>*>(ptr)->append(static_cast<QVirtualKeyboardSelectionListModel::Type>(i));
}

void* QVirtualKeyboardAbstractInputMethod___selectionLists_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVirtualKeyboardSelectionListModel::Type>();
}

void* QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_atList(void* ptr, struct QtVirtualKeyboard_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_setList(void* ptr, struct QtVirtualKeyboard_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_atList(void* ptr, struct QtVirtualKeyboard_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_setList(void* ptr, struct QtVirtualKeyboard_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_setList(void* ptr, struct QtVirtualKeyboard_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_setList(void* ptr, struct QtVirtualKeyboard_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QVirtualKeyboardAbstractInputMethod___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardAbstractInputMethod___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardAbstractInputMethod___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardAbstractInputMethod___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardAbstractInputMethod___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardAbstractInputMethod___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardAbstractInputMethod___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QVirtualKeyboardAbstractInputMethod_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardAbstractInputMethod_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardAbstractInputMethod_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardAbstractInputMethod_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::deleteLater();
}

void QVirtualKeyboardAbstractInputMethod_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardAbstractInputMethod_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardAbstractInputMethod_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardAbstractInputMethod_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::metaObject());
}

void QVirtualKeyboardAbstractInputMethod_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardAbstractInputMethod*>(ptr)->QVirtualKeyboardAbstractInputMethod::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVirtualKeyboardExtensionPlugin: public QVirtualKeyboardExtensionPlugin
{
public:
	void registerTypes(const char * uri) const { QtVirtualKeyboard_PackedString uriPacked = { const_cast<char*>(uri), -1, NULL };callbackQVirtualKeyboardExtensionPlugin_RegisterTypes(const_cast<void*>(static_cast<const void*>(this)), uriPacked); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardExtensionPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardExtensionPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardExtensionPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardExtensionPlugin_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardExtensionPlugin_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardExtensionPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardExtensionPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardExtensionPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardExtensionPlugin_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardExtensionPlugin_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardExtensionPlugin_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardExtensionPlugin*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardExtensionPlugin*)

int QVirtualKeyboardExtensionPlugin_QVirtualKeyboardExtensionPlugin_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardExtensionPlugin*>(); return qRegisterMetaType<MyQVirtualKeyboardExtensionPlugin*>();}

void QVirtualKeyboardExtensionPlugin_RegisterTypes(void* ptr, char* uri)
{
	static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

void QVirtualKeyboardExtensionPlugin_RegisterTypesDefault(void* ptr, char* uri)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::registerTypes(const_cast<const char*>(uri));
}

void* QVirtualKeyboardExtensionPlugin___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardExtensionPlugin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardExtensionPlugin___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardExtensionPlugin___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardExtensionPlugin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardExtensionPlugin___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardExtensionPlugin___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardExtensionPlugin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardExtensionPlugin___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QVirtualKeyboardExtensionPlugin_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardExtensionPlugin_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardExtensionPlugin_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::deleteLater();
}

void QVirtualKeyboardExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardExtensionPlugin_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardExtensionPlugin_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::metaObject());
}

void QVirtualKeyboardExtensionPlugin_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardExtensionPlugin*>(ptr)->QVirtualKeyboardExtensionPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVirtualKeyboardInputContext: public QVirtualKeyboardInputContext
{
public:
	void Signal_AnchorPositionChanged() { callbackQVirtualKeyboardInputContext_AnchorPositionChanged(this); };
	void Signal_AnchorRectIntersectsClipRectChanged() { callbackQVirtualKeyboardInputContext_AnchorRectIntersectsClipRectChanged(this); };
	void Signal_AnchorRectangleChanged() { callbackQVirtualKeyboardInputContext_AnchorRectangleChanged(this); };
	void Signal_AnimatingChanged() { callbackQVirtualKeyboardInputContext_AnimatingChanged(this); };
	void Signal_CapsLockActiveChanged() { callbackQVirtualKeyboardInputContext_CapsLockActiveChanged(this); };
	void Signal_CursorPositionChanged() { callbackQVirtualKeyboardInputContext_CursorPositionChanged(this); };
	void Signal_CursorRectIntersectsClipRectChanged() { callbackQVirtualKeyboardInputContext_CursorRectIntersectsClipRectChanged(this); };
	void Signal_CursorRectangleChanged() { callbackQVirtualKeyboardInputContext_CursorRectangleChanged(this); };
	void Signal_InputItemChanged() { callbackQVirtualKeyboardInputContext_InputItemChanged(this); };
	void Signal_InputMethodHintsChanged() { callbackQVirtualKeyboardInputContext_InputMethodHintsChanged(this); };
	void Signal_LocaleChanged() { callbackQVirtualKeyboardInputContext_LocaleChanged(this); };
	void Signal_PreeditTextChanged() { callbackQVirtualKeyboardInputContext_PreeditTextChanged(this); };
	void Signal_SelectedTextChanged() { callbackQVirtualKeyboardInputContext_SelectedTextChanged(this); };
	void Signal_SelectionControlVisibleChanged() { callbackQVirtualKeyboardInputContext_SelectionControlVisibleChanged(this); };
	void Signal_ShiftActiveChanged() { callbackQVirtualKeyboardInputContext_ShiftActiveChanged(this); };
	void Signal_SurroundingTextChanged() { callbackQVirtualKeyboardInputContext_SurroundingTextChanged(this); };
	void Signal_UppercaseChanged() { callbackQVirtualKeyboardInputContext_UppercaseChanged(this); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardInputContext_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardInputContext_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardInputContext_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardInputContext_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardInputContext_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardInputContext_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardInputContext_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardInputContext_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardInputContext_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardInputContext_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardInputContext_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardInputContext*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardInputContext*)

int QVirtualKeyboardInputContext_QVirtualKeyboardInputContext_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardInputContext*>(); return qRegisterMetaType<MyQVirtualKeyboardInputContext*>();}

int QVirtualKeyboardInputContext_AnchorPosition(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorPosition();
}

void QVirtualKeyboardInputContext_ConnectAnchorPositionChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorPositionChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorPositionChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectAnchorPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorPositionChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorPositionChanged));
}

void QVirtualKeyboardInputContext_AnchorPositionChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorPositionChanged();
}

char QVirtualKeyboardInputContext_AnchorRectIntersectsClipRect(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorRectIntersectsClipRect();
}

void QVirtualKeyboardInputContext_ConnectAnchorRectIntersectsClipRectChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorRectIntersectsClipRectChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorRectIntersectsClipRectChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectAnchorRectIntersectsClipRectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorRectIntersectsClipRectChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorRectIntersectsClipRectChanged));
}

void QVirtualKeyboardInputContext_AnchorRectIntersectsClipRectChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorRectIntersectsClipRectChanged();
}

void* QVirtualKeyboardInputContext_AnchorRectangle(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorRectangle(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QVirtualKeyboardInputContext_ConnectAnchorRectangleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorRectangleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorRectangleChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectAnchorRectangleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::anchorRectangleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnchorRectangleChanged));
}

void QVirtualKeyboardInputContext_AnchorRectangleChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->anchorRectangleChanged();
}

void QVirtualKeyboardInputContext_ConnectAnimatingChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::animatingChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnimatingChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectAnimatingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::animatingChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_AnimatingChanged));
}

void QVirtualKeyboardInputContext_AnimatingChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->animatingChanged();
}

void QVirtualKeyboardInputContext_ConnectCapsLockActiveChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::capsLockActiveChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CapsLockActiveChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectCapsLockActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::capsLockActiveChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CapsLockActiveChanged));
}

void QVirtualKeyboardInputContext_CapsLockActiveChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->capsLockActiveChanged();
}

void QVirtualKeyboardInputContext_Clear(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->clear();
}

void QVirtualKeyboardInputContext_Commit(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->commit();
}

void QVirtualKeyboardInputContext_Commit2(void* ptr, struct QtVirtualKeyboard_PackedString text, int replaceFrom, int replaceLength)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->commit(QString::fromUtf8(text.data, text.len), replaceFrom, replaceLength);
}

int QVirtualKeyboardInputContext_CursorPosition(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorPosition();
}

void QVirtualKeyboardInputContext_ConnectCursorPositionChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorPositionChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorPositionChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectCursorPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorPositionChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorPositionChanged));
}

void QVirtualKeyboardInputContext_CursorPositionChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorPositionChanged();
}

char QVirtualKeyboardInputContext_CursorRectIntersectsClipRect(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorRectIntersectsClipRect();
}

void QVirtualKeyboardInputContext_ConnectCursorRectIntersectsClipRectChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorRectIntersectsClipRectChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorRectIntersectsClipRectChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectCursorRectIntersectsClipRectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorRectIntersectsClipRectChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorRectIntersectsClipRectChanged));
}

void QVirtualKeyboardInputContext_CursorRectIntersectsClipRectChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorRectIntersectsClipRectChanged();
}

void* QVirtualKeyboardInputContext_CursorRectangle(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorRectangle(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QVirtualKeyboardInputContext_ConnectCursorRectangleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorRectangleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorRectangleChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectCursorRectangleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::cursorRectangleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_CursorRectangleChanged));
}

void QVirtualKeyboardInputContext_CursorRectangleChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->cursorRectangleChanged();
}

void* QVirtualKeyboardInputContext_InputEngine(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->inputEngine();
}

void* QVirtualKeyboardInputContext_InputItem(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->inputItem();
}

void QVirtualKeyboardInputContext_ConnectInputItemChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::inputItemChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_InputItemChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectInputItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::inputItemChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_InputItemChanged));
}

void QVirtualKeyboardInputContext_InputItemChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->inputItemChanged();
}

long long QVirtualKeyboardInputContext_InputMethodHints(void* ptr)
{
	return static_cast<qint64>(static_cast<QVirtualKeyboardInputContext*>(ptr)->inputMethodHints());
}

void QVirtualKeyboardInputContext_ConnectInputMethodHintsChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::inputMethodHintsChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_InputMethodHintsChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectInputMethodHintsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::inputMethodHintsChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_InputMethodHintsChanged));
}

void QVirtualKeyboardInputContext_InputMethodHintsChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->inputMethodHintsChanged();
}

char QVirtualKeyboardInputContext_IsAnimating(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->isAnimating();
}

char QVirtualKeyboardInputContext_IsCapsLockActive(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->isCapsLockActive();
}

char QVirtualKeyboardInputContext_IsSelectionControlVisible(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->isSelectionControlVisible();
}

char QVirtualKeyboardInputContext_IsShiftActive(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->isShiftActive();
}

char QVirtualKeyboardInputContext_IsUppercase(void* ptr)
{
	return static_cast<QVirtualKeyboardInputContext*>(ptr)->isUppercase();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputContext_Locale(void* ptr)
{
	return ({ QByteArray* tb5677a = new QByteArray(static_cast<QVirtualKeyboardInputContext*>(ptr)->locale().toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(tb5677a->prepend("WHITESPACE").constData()+10), tb5677a->size()-10, tb5677a }; });
}

void QVirtualKeyboardInputContext_ConnectLocaleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::localeChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_LocaleChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectLocaleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::localeChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_LocaleChanged));
}

void QVirtualKeyboardInputContext_LocaleChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->localeChanged();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputContext_PreeditText(void* ptr)
{
	return ({ QByteArray* t2c7998 = new QByteArray(static_cast<QVirtualKeyboardInputContext*>(ptr)->preeditText().toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t2c7998->prepend("WHITESPACE").constData()+10), t2c7998->size()-10, t2c7998 }; });
}

void QVirtualKeyboardInputContext_ConnectPreeditTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::preeditTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_PreeditTextChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectPreeditTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::preeditTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_PreeditTextChanged));
}

void QVirtualKeyboardInputContext_PreeditTextChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->preeditTextChanged();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputContext_SelectedText(void* ptr)
{
	return ({ QByteArray* tafca69 = new QByteArray(static_cast<QVirtualKeyboardInputContext*>(ptr)->selectedText().toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(tafca69->prepend("WHITESPACE").constData()+10), tafca69->size()-10, tafca69 }; });
}

void QVirtualKeyboardInputContext_ConnectSelectedTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::selectedTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SelectedTextChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectSelectedTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::selectedTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SelectedTextChanged));
}

void QVirtualKeyboardInputContext_SelectedTextChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->selectedTextChanged();
}

void QVirtualKeyboardInputContext_ConnectSelectionControlVisibleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::selectionControlVisibleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SelectionControlVisibleChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectSelectionControlVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::selectionControlVisibleChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SelectionControlVisibleChanged));
}

void QVirtualKeyboardInputContext_SelectionControlVisibleChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->selectionControlVisibleChanged();
}

void QVirtualKeyboardInputContext_SendKeyClick(void* ptr, int key, struct QtVirtualKeyboard_PackedString text, int modifiers)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->sendKeyClick(key, QString::fromUtf8(text.data, text.len), modifiers);
}

void QVirtualKeyboardInputContext_SetAnimating(void* ptr, char isAnimating)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->setAnimating(isAnimating != 0);
}

void QVirtualKeyboardInputContext_ConnectShiftActiveChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::shiftActiveChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_ShiftActiveChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectShiftActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::shiftActiveChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_ShiftActiveChanged));
}

void QVirtualKeyboardInputContext_ShiftActiveChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->shiftActiveChanged();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputContext_SurroundingText(void* ptr)
{
	return ({ QByteArray* t710186 = new QByteArray(static_cast<QVirtualKeyboardInputContext*>(ptr)->surroundingText().toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t710186->prepend("WHITESPACE").constData()+10), t710186->size()-10, t710186 }; });
}

void QVirtualKeyboardInputContext_ConnectSurroundingTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::surroundingTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SurroundingTextChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectSurroundingTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::surroundingTextChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_SurroundingTextChanged));
}

void QVirtualKeyboardInputContext_SurroundingTextChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->surroundingTextChanged();
}

void QVirtualKeyboardInputContext_ConnectUppercaseChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::uppercaseChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_UppercaseChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputContext_DisconnectUppercaseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputContext*>(ptr), static_cast<void (QVirtualKeyboardInputContext::*)()>(&QVirtualKeyboardInputContext::uppercaseChanged), static_cast<MyQVirtualKeyboardInputContext*>(ptr), static_cast<void (MyQVirtualKeyboardInputContext::*)()>(&MyQVirtualKeyboardInputContext::Signal_UppercaseChanged));
}

void QVirtualKeyboardInputContext_UppercaseChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputContext*>(ptr)->uppercaseChanged();
}

void* QVirtualKeyboardInputContext___preeditTextAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QInputMethodEvent::Attribute>();
}

void* QVirtualKeyboardInputContext___setPreeditText_attributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QInputMethodEvent::Attribute>();
}

void* QVirtualKeyboardInputContext___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputContext___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputContext___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardInputContext___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardInputContext___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardInputContext___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardInputContext___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputContext___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputContext___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardInputContext___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputContext___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputContext___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QVirtualKeyboardInputContext_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardInputContext_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardInputContext_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardInputContext_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::deleteLater();
}

void QVirtualKeyboardInputContext_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardInputContext_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardInputContext_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardInputContext_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::metaObject());
}

void QVirtualKeyboardInputContext_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputContext*>(ptr)->QVirtualKeyboardInputContext::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVirtualKeyboardInputEngine: public QVirtualKeyboardInputEngine
{
public:
	void Signal_ActiveKeyChanged(Qt::Key key) { callbackQVirtualKeyboardInputEngine_ActiveKeyChanged(this, static_cast<qint64>(key)); };
	void Signal_InputMethodChanged() { callbackQVirtualKeyboardInputEngine_InputMethodChanged(this); };
	void Signal_InputMethodReset() { callbackQVirtualKeyboardInputEngine_InputMethodReset(this); };
	void Signal_InputMethodUpdate() { callbackQVirtualKeyboardInputEngine_InputMethodUpdate(this); };
	void Signal_InputModeChanged() { callbackQVirtualKeyboardInputEngine_InputModeChanged(this); };
	void Signal_InputModesChanged() { callbackQVirtualKeyboardInputEngine_InputModesChanged(this); };
	void Signal_PatternRecognitionModesChanged() { callbackQVirtualKeyboardInputEngine_PatternRecognitionModesChanged(this); };
	void Signal_PreviousKeyChanged(Qt::Key key) { callbackQVirtualKeyboardInputEngine_PreviousKeyChanged(this, static_cast<qint64>(key)); };
	void Signal_VirtualKeyClicked(Qt::Key key, const QString & text, Qt::KeyboardModifiers modifiers, bool isAutoRepeat) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtVirtualKeyboard_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQVirtualKeyboardInputEngine_VirtualKeyClicked(this, static_cast<qint64>(key), textPacked, static_cast<qint64>(modifiers), isAutoRepeat); };
	void Signal_WordCandidateListModelChanged() { callbackQVirtualKeyboardInputEngine_WordCandidateListModelChanged(this); };
	void Signal_WordCandidateListVisibleHintChanged() { callbackQVirtualKeyboardInputEngine_WordCandidateListVisibleHintChanged(this); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardInputEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardInputEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardInputEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardInputEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardInputEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardInputEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardInputEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardInputEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardInputEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardInputEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardInputEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardInputEngine*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardInputEngine*)

int QVirtualKeyboardInputEngine_QVirtualKeyboardInputEngine_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardInputEngine*>(); return qRegisterMetaType<MyQVirtualKeyboardInputEngine*>();}

long long QVirtualKeyboardInputEngine_ActiveKey(void* ptr)
{
	return static_cast<qint64>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->activeKey());
}

void QVirtualKeyboardInputEngine_ConnectActiveKeyChanged(void* ptr, long long t)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key)>(&QVirtualKeyboardInputEngine::activeKeyChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key)>(&MyQVirtualKeyboardInputEngine::Signal_ActiveKeyChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectActiveKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key)>(&QVirtualKeyboardInputEngine::activeKeyChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key)>(&MyQVirtualKeyboardInputEngine::Signal_ActiveKeyChanged));
}

void QVirtualKeyboardInputEngine_ActiveKeyChanged(void* ptr, long long key)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->activeKeyChanged(static_cast<Qt::Key>(key));
}

void* QVirtualKeyboardInputEngine_InputContext(void* ptr)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputContext();
}

void* QVirtualKeyboardInputEngine_InputMethod(void* ptr)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputMethod();
}

void QVirtualKeyboardInputEngine_ConnectInputMethodChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectInputMethodChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodChanged));
}

void QVirtualKeyboardInputEngine_InputMethodChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputMethodChanged();
}

void QVirtualKeyboardInputEngine_ConnectInputMethodReset(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodReset), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodReset), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectInputMethodReset(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodReset), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodReset));
}

void QVirtualKeyboardInputEngine_InputMethodReset(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputMethodReset();
}

void QVirtualKeyboardInputEngine_ConnectInputMethodUpdate(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodUpdate), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodUpdate), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectInputMethodUpdate(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputMethodUpdate), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputMethodUpdate));
}

void QVirtualKeyboardInputEngine_InputMethodUpdate(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputMethodUpdate();
}

long long QVirtualKeyboardInputEngine_InputMode(void* ptr)
{
	return static_cast<qint64>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputMode());
}

void QVirtualKeyboardInputEngine_ConnectInputModeChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputModeChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputModeChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectInputModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputModeChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputModeChanged));
}

void QVirtualKeyboardInputEngine_InputModeChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputModeChanged();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardInputEngine_InputModes(void* ptr)
{
	return ({ QList<int>* tmpValueefe75f = new QList<int>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputModes()); QtVirtualKeyboard_PackedList { tmpValueefe75f, tmpValueefe75f->size() }; });
}

void QVirtualKeyboardInputEngine_ConnectInputModesChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputModesChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputModesChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectInputModesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::inputModesChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_InputModesChanged));
}

void QVirtualKeyboardInputEngine_InputModesChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->inputModesChanged();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardInputEngine_PatternRecognitionModes(void* ptr)
{
	return ({ QList<int>* tmpValueaec2e9 = new QList<int>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->patternRecognitionModes()); QtVirtualKeyboard_PackedList { tmpValueaec2e9, tmpValueaec2e9->size() }; });
}

void QVirtualKeyboardInputEngine_ConnectPatternRecognitionModesChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::patternRecognitionModesChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_PatternRecognitionModesChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectPatternRecognitionModesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::patternRecognitionModesChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_PatternRecognitionModesChanged));
}

void QVirtualKeyboardInputEngine_PatternRecognitionModesChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->patternRecognitionModesChanged();
}

long long QVirtualKeyboardInputEngine_PreviousKey(void* ptr)
{
	return static_cast<qint64>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->previousKey());
}

void QVirtualKeyboardInputEngine_ConnectPreviousKeyChanged(void* ptr, long long t)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key)>(&QVirtualKeyboardInputEngine::previousKeyChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key)>(&MyQVirtualKeyboardInputEngine::Signal_PreviousKeyChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectPreviousKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key)>(&QVirtualKeyboardInputEngine::previousKeyChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key)>(&MyQVirtualKeyboardInputEngine::Signal_PreviousKeyChanged));
}

void QVirtualKeyboardInputEngine_PreviousKeyChanged(void* ptr, long long key)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->previousKeyChanged(static_cast<Qt::Key>(key));
}

char QVirtualKeyboardInputEngine_Reselect(void* ptr, int cursorPosition, long long reselectFlags)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->reselect(cursorPosition, static_cast<QVirtualKeyboardInputEngine::ReselectFlag>(reselectFlags));
}

void QVirtualKeyboardInputEngine_SetInputMethod(void* ptr, void* inputMethod)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->setInputMethod(static_cast<QVirtualKeyboardAbstractInputMethod*>(inputMethod));
}

void QVirtualKeyboardInputEngine_SetInputMode(void* ptr, long long inputMode)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->setInputMode(static_cast<QVirtualKeyboardInputEngine::InputMode>(inputMode));
}

void* QVirtualKeyboardInputEngine_TraceBegin(void* ptr, int traceId, long long patternRecognitionMode, void* traceCaptureDeviceInfo, void* traceScreenInfo)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->traceBegin(traceId, static_cast<QVirtualKeyboardInputEngine::PatternRecognitionMode>(patternRecognitionMode), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceCaptureDeviceInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(traceScreenInfo); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

char QVirtualKeyboardInputEngine_TraceEnd(void* ptr, void* trace)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->traceEnd(static_cast<QVirtualKeyboardTrace*>(trace));
}

void QVirtualKeyboardInputEngine_VirtualKeyCancel(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->virtualKeyCancel();
}

char QVirtualKeyboardInputEngine_VirtualKeyClick(void* ptr, long long key, struct QtVirtualKeyboard_PackedString text, long long modifiers)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->virtualKeyClick(static_cast<Qt::Key>(key), QString::fromUtf8(text.data, text.len), static_cast<Qt::KeyboardModifier>(modifiers));
}

void QVirtualKeyboardInputEngine_ConnectVirtualKeyClicked(void* ptr, long long t)
{
	qRegisterMetaType<Qt::Key>();
	qRegisterMetaType<Qt::KeyboardModifiers>();
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key, const QString &, Qt::KeyboardModifiers, bool)>(&QVirtualKeyboardInputEngine::virtualKeyClicked), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key, const QString &, Qt::KeyboardModifiers, bool)>(&MyQVirtualKeyboardInputEngine::Signal_VirtualKeyClicked), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectVirtualKeyClicked(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)(Qt::Key, const QString &, Qt::KeyboardModifiers, bool)>(&QVirtualKeyboardInputEngine::virtualKeyClicked), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)(Qt::Key, const QString &, Qt::KeyboardModifiers, bool)>(&MyQVirtualKeyboardInputEngine::Signal_VirtualKeyClicked));
}

void QVirtualKeyboardInputEngine_VirtualKeyClicked(void* ptr, long long key, struct QtVirtualKeyboard_PackedString text, long long modifiers, char isAutoRepeat)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->virtualKeyClicked(static_cast<Qt::Key>(key), QString::fromUtf8(text.data, text.len), static_cast<Qt::KeyboardModifier>(modifiers), isAutoRepeat != 0);
}

char QVirtualKeyboardInputEngine_VirtualKeyPress(void* ptr, long long key, struct QtVirtualKeyboard_PackedString text, long long modifiers, char repeat)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->virtualKeyPress(static_cast<Qt::Key>(key), QString::fromUtf8(text.data, text.len), static_cast<Qt::KeyboardModifier>(modifiers), repeat != 0);
}

char QVirtualKeyboardInputEngine_VirtualKeyRelease(void* ptr, long long key, struct QtVirtualKeyboard_PackedString text, long long modifiers)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->virtualKeyRelease(static_cast<Qt::Key>(key), QString::fromUtf8(text.data, text.len), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QVirtualKeyboardInputEngine_WordCandidateListModel(void* ptr)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->wordCandidateListModel();
}

void QVirtualKeyboardInputEngine_ConnectWordCandidateListModelChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::wordCandidateListModelChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_WordCandidateListModelChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectWordCandidateListModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::wordCandidateListModelChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_WordCandidateListModelChanged));
}

void QVirtualKeyboardInputEngine_WordCandidateListModelChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->wordCandidateListModelChanged();
}

char QVirtualKeyboardInputEngine_WordCandidateListVisibleHint(void* ptr)
{
	return static_cast<QVirtualKeyboardInputEngine*>(ptr)->wordCandidateListVisibleHint();
}

void QVirtualKeyboardInputEngine_ConnectWordCandidateListVisibleHintChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::wordCandidateListVisibleHintChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_WordCandidateListVisibleHintChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardInputEngine_DisconnectWordCandidateListVisibleHintChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardInputEngine*>(ptr), static_cast<void (QVirtualKeyboardInputEngine::*)()>(&QVirtualKeyboardInputEngine::wordCandidateListVisibleHintChanged), static_cast<MyQVirtualKeyboardInputEngine*>(ptr), static_cast<void (MyQVirtualKeyboardInputEngine::*)()>(&MyQVirtualKeyboardInputEngine::Signal_WordCandidateListVisibleHintChanged));
}

void QVirtualKeyboardInputEngine_WordCandidateListVisibleHintChanged(void* ptr)
{
	static_cast<QVirtualKeyboardInputEngine*>(ptr)->wordCandidateListVisibleHintChanged();
}

int QVirtualKeyboardInputEngine___inputModes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputEngine___inputModes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardInputEngine___inputModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QVirtualKeyboardInputEngine___patternRecognitionModes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputEngine___patternRecognitionModes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardInputEngine___patternRecognitionModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_atList(void* ptr, struct QtVirtualKeyboard_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_setList(void* ptr, struct QtVirtualKeyboard_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_atList(void* ptr, struct QtVirtualKeyboard_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_setList(void* ptr, struct QtVirtualKeyboard_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_setList(void* ptr, struct QtVirtualKeyboard_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_setList(void* ptr, struct QtVirtualKeyboard_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QVirtualKeyboardInputEngine___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardInputEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardInputEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardInputEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardInputEngine___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardInputEngine___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardInputEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardInputEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QVirtualKeyboardInputEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardInputEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardInputEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardInputEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::deleteLater();
}

void QVirtualKeyboardInputEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardInputEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardInputEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardInputEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::metaObject());
}

void QVirtualKeyboardInputEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardInputEngine*>(ptr)->QVirtualKeyboardInputEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVirtualKeyboardSelectionListModel: public QVirtualKeyboardSelectionListModel
{
public:
	void Signal_ActiveItemChanged(int index) { callbackQVirtualKeyboardSelectionListModel_ActiveItemChanged(this, index); };
	void Signal_ItemSelected(int index) { callbackQVirtualKeyboardSelectionListModel_ItemSelected(this, index); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQVirtualKeyboardSelectionListModel_DropMimeData(this, const_cast<QMimeData*>(data), static_cast<qint64>(action), row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQVirtualKeyboardSelectionListModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQVirtualKeyboardSelectionListModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQVirtualKeyboardSelectionListModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQVirtualKeyboardSelectionListModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQVirtualKeyboardSelectionListModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), static_cast<qint64>(action), row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQVirtualKeyboardSelectionListModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQVirtualKeyboardSelectionListModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQVirtualKeyboardSelectionListModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQVirtualKeyboardSelectionListModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQVirtualKeyboardSelectionListModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtVirtualKeyboard_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackQVirtualKeyboardSelectionListModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQVirtualKeyboardSelectionListModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQVirtualKeyboardSelectionListModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, static_cast<qint64>(orientation), role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQVirtualKeyboardSelectionListModel_HeaderDataChanged(this, static_cast<qint64>(orientation), first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQVirtualKeyboardSelectionListModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQVirtualKeyboardSelectionListModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQVirtualKeyboardSelectionListModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQVirtualKeyboardSelectionListModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtVirtualKeyboard_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), static_cast<qint64>(hint)); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQVirtualKeyboardSelectionListModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtVirtualKeyboard_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), static_cast<qint64>(hint)); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQVirtualKeyboardSelectionListModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, static_cast<qint64>(flags))); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQVirtualKeyboardSelectionListModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtVirtualKeyboard_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtVirtualKeyboard_PackedString tempVal = callbackQVirtualKeyboardSelectionListModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQVirtualKeyboardSelectionListModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQVirtualKeyboardSelectionListModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQVirtualKeyboardSelectionListModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQVirtualKeyboardSelectionListModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQVirtualKeyboardSelectionListModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQVirtualKeyboardSelectionListModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQVirtualKeyboardSelectionListModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQVirtualKeyboardSelectionListModel_ResetInternalData(this); };
	void revert() { callbackQVirtualKeyboardSelectionListModel_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQVirtualKeyboardSelectionListModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQVirtualKeyboardSelectionListModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQVirtualKeyboardSelectionListModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQVirtualKeyboardSelectionListModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQVirtualKeyboardSelectionListModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQVirtualKeyboardSelectionListModel_SetHeaderData(this, section, static_cast<qint64>(orientation), const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQVirtualKeyboardSelectionListModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtVirtualKeyboard_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQVirtualKeyboardSelectionListModel_Sort(this, column, static_cast<qint64>(order)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQVirtualKeyboardSelectionListModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQVirtualKeyboardSelectionListModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQVirtualKeyboardSelectionListModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQVirtualKeyboardSelectionListModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardSelectionListModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardSelectionListModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardSelectionListModel_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardSelectionListModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardSelectionListModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardSelectionListModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardSelectionListModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardSelectionListModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardSelectionListModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardSelectionListModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardSelectionListModel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardSelectionListModel*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardSelectionListModel*)

int QVirtualKeyboardSelectionListModel_QVirtualKeyboardSelectionListModel_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardSelectionListModel*>(); return qRegisterMetaType<MyQVirtualKeyboardSelectionListModel*>();}

void QVirtualKeyboardSelectionListModel_ConnectActiveItemChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (QVirtualKeyboardSelectionListModel::*)(int)>(&QVirtualKeyboardSelectionListModel::activeItemChanged), static_cast<MyQVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (MyQVirtualKeyboardSelectionListModel::*)(int)>(&MyQVirtualKeyboardSelectionListModel::Signal_ActiveItemChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardSelectionListModel_DisconnectActiveItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (QVirtualKeyboardSelectionListModel::*)(int)>(&QVirtualKeyboardSelectionListModel::activeItemChanged), static_cast<MyQVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (MyQVirtualKeyboardSelectionListModel::*)(int)>(&MyQVirtualKeyboardSelectionListModel::Signal_ActiveItemChanged));
}

void QVirtualKeyboardSelectionListModel_ActiveItemChanged(void* ptr, int index)
{
	static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->activeItemChanged(index);
}

void QVirtualKeyboardSelectionListModel_ConnectItemSelected(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (QVirtualKeyboardSelectionListModel::*)(int)>(&QVirtualKeyboardSelectionListModel::itemSelected), static_cast<MyQVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (MyQVirtualKeyboardSelectionListModel::*)(int)>(&MyQVirtualKeyboardSelectionListModel::Signal_ItemSelected), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardSelectionListModel_DisconnectItemSelected(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (QVirtualKeyboardSelectionListModel::*)(int)>(&QVirtualKeyboardSelectionListModel::itemSelected), static_cast<MyQVirtualKeyboardSelectionListModel*>(ptr), static_cast<void (MyQVirtualKeyboardSelectionListModel::*)(int)>(&MyQVirtualKeyboardSelectionListModel::Signal_ItemSelected));
}

void QVirtualKeyboardSelectionListModel_ItemSelected(void* ptr, int index)
{
	static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->itemSelected(index);
}

void QVirtualKeyboardSelectionListModel_RemoveItem(void* ptr, int index)
{
	static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->removeItem(index);
}

void QVirtualKeyboardSelectionListModel_SelectItem(void* ptr, int index)
{
	static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->selectItem(index);
}

void* QVirtualKeyboardSelectionListModel___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardSelectionListModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QVirtualKeyboardSelectionListModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QVirtualKeyboardSelectionListModel_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QVirtualKeyboardSelectionListModel___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QVirtualKeyboardSelectionListModel___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardSelectionListModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QVirtualKeyboardSelectionListModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QVirtualKeyboardSelectionListModel___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QVirtualKeyboardSelectionListModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtVirtualKeyboard_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

int QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QVirtualKeyboardSelectionListModel___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardSelectionListModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardSelectionListModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardSelectionListModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardSelectionListModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardSelectionListModel___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardSelectionListModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardSelectionListModel___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardSelectionListModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardSelectionListModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

char QVirtualKeyboardSelectionListModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long QVirtualKeyboardSelectionListModel_FlagsDefault(void* ptr, void* index)
{
		return static_cast<qint64>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::flags(*static_cast<QModelIndex*>(index)));
}

void* QVirtualKeyboardSelectionListModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QVirtualKeyboardSelectionListModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
		return new QModelIndex(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QVirtualKeyboardSelectionListModel_BuddyDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char QVirtualKeyboardSelectionListModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QVirtualKeyboardSelectionListModel_CanFetchMoreDefault(void* ptr, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QVirtualKeyboardSelectionListModel_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* QVirtualKeyboardSelectionListModel_DataDefault(void* ptr, void* index, int role)
{
		return new QVariant(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::data(*static_cast<QModelIndex*>(index), role));
}

void QVirtualKeyboardSelectionListModel_FetchMoreDefault(void* ptr, void* parent)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char QVirtualKeyboardSelectionListModel_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QVirtualKeyboardSelectionListModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
		return new QVariant(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char QVirtualKeyboardSelectionListModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QVirtualKeyboardSelectionListModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel_ItemDataDefault(void* ptr, void* index)
{
		return ({ QMap<int, QVariant>* tmpValue644fa8 = new QMap<int, QVariant>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::itemData(*static_cast<QModelIndex*>(index))); QtVirtualKeyboard_PackedList { tmpValue644fa8, tmpValue644fa8->size() }; });
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValue1190c9 = new QList<QModelIndex>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtVirtualKeyboard_PackedList { tmpValue1190c9, tmpValue1190c9->size() }; });
}

void* QVirtualKeyboardSelectionListModel_MimeDataDefault(void* ptr, void* indexes)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardSelectionListModel_MimeTypesDefault(void* ptr)
{
		return ({ QByteArray* t30b194 = new QByteArray(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::mimeTypes().join("¡¦!").toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t30b194->prepend("WHITESPACE").constData()+10), t30b194->size()-10, t30b194 }; });
}

char QVirtualKeyboardSelectionListModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QVirtualKeyboardSelectionListModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QVirtualKeyboardSelectionListModel_ParentDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::parent(*static_cast<QModelIndex*>(index)));
}

char QVirtualKeyboardSelectionListModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QVirtualKeyboardSelectionListModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QVirtualKeyboardSelectionListModel_ResetInternalDataDefault(void* ptr)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::resetInternalData();
}

void QVirtualKeyboardSelectionListModel_RevertDefault(void* ptr)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::revert();
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardSelectionListModel_RoleNamesDefault(void* ptr)
{
		return ({ QHash<int, QByteArray>* tmpValue9e1c1f = new QHash<int, QByteArray>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::roleNames()); QtVirtualKeyboard_PackedList { tmpValue9e1c1f, tmpValue9e1c1f->size() }; });
}

int QVirtualKeyboardSelectionListModel_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::rowCount(*static_cast<QModelIndex*>(parent));
}

char QVirtualKeyboardSelectionListModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QVirtualKeyboardSelectionListModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QVirtualKeyboardSelectionListModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void QVirtualKeyboardSelectionListModel_SortDefault(void* ptr, int column, long long order)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QVirtualKeyboardSelectionListModel_SpanDefault(void* ptr, void* index)
{
		return ({ QSize tmpValue = static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QVirtualKeyboardSelectionListModel_SubmitDefault(void* ptr)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::submit();
}

long long QVirtualKeyboardSelectionListModel_SupportedDragActionsDefault(void* ptr)
{
		return static_cast<qint64>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::supportedDragActions());
}

long long QVirtualKeyboardSelectionListModel_SupportedDropActionsDefault(void* ptr)
{
		return static_cast<qint64>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::supportedDropActions());
}

void QVirtualKeyboardSelectionListModel_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardSelectionListModel_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardSelectionListModel_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardSelectionListModel_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::deleteLater();
}

void QVirtualKeyboardSelectionListModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardSelectionListModel_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardSelectionListModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardSelectionListModel_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::metaObject());
}

void QVirtualKeyboardSelectionListModel_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardSelectionListModel*>(ptr)->QVirtualKeyboardSelectionListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVirtualKeyboardTrace: public QVirtualKeyboardTrace
{
public:
	void Signal_CanceledChanged(bool isCanceled) { callbackQVirtualKeyboardTrace_CanceledChanged(this, isCanceled); };
	void Signal_ChannelsChanged() { callbackQVirtualKeyboardTrace_ChannelsChanged(this); };
	void Signal_FinalChanged(bool isFinal) { callbackQVirtualKeyboardTrace_FinalChanged(this, isFinal); };
	void Signal_LengthChanged(int length) { callbackQVirtualKeyboardTrace_LengthChanged(this, length); };
	void Signal_OpacityChanged(qreal opacity) { callbackQVirtualKeyboardTrace_OpacityChanged(this, opacity); };
	void Signal_TraceIdChanged(int traceId) { callbackQVirtualKeyboardTrace_TraceIdChanged(this, traceId); };
	void childEvent(QChildEvent * event) { callbackQVirtualKeyboardTrace_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardTrace_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVirtualKeyboardTrace_CustomEvent(this, event); };
	void deleteLater() { callbackQVirtualKeyboardTrace_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVirtualKeyboardTrace_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVirtualKeyboardTrace_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVirtualKeyboardTrace_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVirtualKeyboardTrace_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVirtualKeyboardTrace_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtVirtualKeyboard_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQVirtualKeyboardTrace_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVirtualKeyboardTrace_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QVirtualKeyboardTrace*)
Q_DECLARE_METATYPE(MyQVirtualKeyboardTrace*)

int QVirtualKeyboardTrace_QVirtualKeyboardTrace_QRegisterMetaType(){qRegisterMetaType<QVirtualKeyboardTrace*>(); return qRegisterMetaType<MyQVirtualKeyboardTrace*>();}

int QVirtualKeyboardTrace_AddPoint(void* ptr, void* point)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->addPoint(*static_cast<QPointF*>(point));
}

void QVirtualKeyboardTrace_ConnectCanceledChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(bool)>(&QVirtualKeyboardTrace::canceledChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(bool)>(&MyQVirtualKeyboardTrace::Signal_CanceledChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectCanceledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(bool)>(&QVirtualKeyboardTrace::canceledChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(bool)>(&MyQVirtualKeyboardTrace::Signal_CanceledChanged));
}

void QVirtualKeyboardTrace_CanceledChanged(void* ptr, char isCanceled)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->canceledChanged(isCanceled != 0);
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardTrace_ChannelData(void* ptr, struct QtVirtualKeyboard_PackedString channel, int pos, int count)
{
	return ({ QList<QVariant>* tmpValue1cce99 = new QList<QVariant>(static_cast<QVirtualKeyboardTrace*>(ptr)->channelData(QString::fromUtf8(channel.data, channel.len), pos, count)); QtVirtualKeyboard_PackedList { tmpValue1cce99, tmpValue1cce99->size() }; });
}

struct QtVirtualKeyboard_PackedString QVirtualKeyboardTrace_Channels(void* ptr)
{
	return ({ QByteArray* t8a9301 = new QByteArray(static_cast<QVirtualKeyboardTrace*>(ptr)->channels().join("¡¦!").toUtf8()); QtVirtualKeyboard_PackedString { const_cast<char*>(t8a9301->prepend("WHITESPACE").constData()+10), t8a9301->size()-10, t8a9301 }; });
}

void QVirtualKeyboardTrace_ConnectChannelsChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)()>(&QVirtualKeyboardTrace::channelsChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)()>(&MyQVirtualKeyboardTrace::Signal_ChannelsChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectChannelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)()>(&QVirtualKeyboardTrace::channelsChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)()>(&MyQVirtualKeyboardTrace::Signal_ChannelsChanged));
}

void QVirtualKeyboardTrace_ChannelsChanged(void* ptr)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->channelsChanged();
}

void QVirtualKeyboardTrace_ConnectFinalChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(bool)>(&QVirtualKeyboardTrace::finalChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(bool)>(&MyQVirtualKeyboardTrace::Signal_FinalChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectFinalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(bool)>(&QVirtualKeyboardTrace::finalChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(bool)>(&MyQVirtualKeyboardTrace::Signal_FinalChanged));
}

void QVirtualKeyboardTrace_FinalChanged(void* ptr, char isFinal)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->finalChanged(isFinal != 0);
}

char QVirtualKeyboardTrace_IsCanceled(void* ptr)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->isCanceled();
}

char QVirtualKeyboardTrace_IsFinal(void* ptr)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->isFinal();
}

int QVirtualKeyboardTrace_Length(void* ptr)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->length();
}

void QVirtualKeyboardTrace_ConnectLengthChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(int)>(&QVirtualKeyboardTrace::lengthChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(int)>(&MyQVirtualKeyboardTrace::Signal_LengthChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectLengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(int)>(&QVirtualKeyboardTrace::lengthChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(int)>(&MyQVirtualKeyboardTrace::Signal_LengthChanged));
}

void QVirtualKeyboardTrace_LengthChanged(void* ptr, int length)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->lengthChanged(length);
}

double QVirtualKeyboardTrace_Opacity(void* ptr)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->opacity();
}

void QVirtualKeyboardTrace_ConnectOpacityChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(qreal)>(&QVirtualKeyboardTrace::opacityChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(qreal)>(&MyQVirtualKeyboardTrace::Signal_OpacityChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectOpacityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(qreal)>(&QVirtualKeyboardTrace::opacityChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(qreal)>(&MyQVirtualKeyboardTrace::Signal_OpacityChanged));
}

void QVirtualKeyboardTrace_OpacityChanged(void* ptr, double opacity)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->opacityChanged(opacity);
}

struct QtVirtualKeyboard_PackedList QVirtualKeyboardTrace_Points(void* ptr, int pos, int count)
{
	return ({ QList<QVariant>* tmpValued2cf96 = new QList<QVariant>(static_cast<QVirtualKeyboardTrace*>(ptr)->points(pos, count)); QtVirtualKeyboard_PackedList { tmpValued2cf96, tmpValued2cf96->size() }; });
}

void QVirtualKeyboardTrace_SetCanceled(void* ptr, char canceled)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setCanceled(canceled != 0);
}

void QVirtualKeyboardTrace_SetChannelData(void* ptr, struct QtVirtualKeyboard_PackedString channel, int index, void* data)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setChannelData(QString::fromUtf8(channel.data, channel.len), index, *static_cast<QVariant*>(data));
}

void QVirtualKeyboardTrace_SetChannels(void* ptr, struct QtVirtualKeyboard_PackedString channels)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setChannels(QString::fromUtf8(channels.data, channels.len).split("¡¦!", QString::SkipEmptyParts));
}

void QVirtualKeyboardTrace_SetFinal(void* ptr, char final)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setFinal(final != 0);
}

void QVirtualKeyboardTrace_SetOpacity(void* ptr, double opacity)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setOpacity(opacity);
}

void QVirtualKeyboardTrace_SetTraceId(void* ptr, int id)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->setTraceId(id);
}

int QVirtualKeyboardTrace_TraceId(void* ptr)
{
	return static_cast<QVirtualKeyboardTrace*>(ptr)->traceId();
}

void QVirtualKeyboardTrace_ConnectTraceIdChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(int)>(&QVirtualKeyboardTrace::traceIdChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(int)>(&MyQVirtualKeyboardTrace::Signal_TraceIdChanged), static_cast<Qt::ConnectionType>(t));
}

void QVirtualKeyboardTrace_DisconnectTraceIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVirtualKeyboardTrace*>(ptr), static_cast<void (QVirtualKeyboardTrace::*)(int)>(&QVirtualKeyboardTrace::traceIdChanged), static_cast<MyQVirtualKeyboardTrace*>(ptr), static_cast<void (MyQVirtualKeyboardTrace::*)(int)>(&MyQVirtualKeyboardTrace::Signal_TraceIdChanged));
}

void QVirtualKeyboardTrace_TraceIdChanged(void* ptr, int traceId)
{
	static_cast<QVirtualKeyboardTrace*>(ptr)->traceIdChanged(traceId);
}

void* QVirtualKeyboardTrace___channelData_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardTrace___channelData_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVirtualKeyboardTrace___channelData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVirtualKeyboardTrace___points_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardTrace___points_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVirtualKeyboardTrace___points_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVirtualKeyboardTrace___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardTrace___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardTrace___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QVirtualKeyboardTrace___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVirtualKeyboardTrace___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVirtualKeyboardTrace___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVirtualKeyboardTrace___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardTrace___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardTrace___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVirtualKeyboardTrace___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVirtualKeyboardTrace___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVirtualKeyboardTrace___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QVirtualKeyboardTrace_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::childEvent(static_cast<QChildEvent*>(event));
}

void QVirtualKeyboardTrace_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVirtualKeyboardTrace_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::customEvent(static_cast<QEvent*>(event));
}

void QVirtualKeyboardTrace_DeleteLaterDefault(void* ptr)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::deleteLater();
}

void QVirtualKeyboardTrace_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVirtualKeyboardTrace_EventDefault(void* ptr, void* e)
{
		return static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::event(static_cast<QEvent*>(e));
}

char QVirtualKeyboardTrace_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVirtualKeyboardTrace_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::metaObject());
}

void QVirtualKeyboardTrace_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVirtualKeyboardTrace*>(ptr)->QVirtualKeyboardTrace::timerEvent(static_cast<QTimerEvent*>(event));
}

