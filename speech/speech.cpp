// +build !minimal

#define protected public
#define private public

#include "speech.h"
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
#include <QLayout>
#include <QList>
#include <QLocale>
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
#include <QSignalSpy>
#include <QString>
#include <QTextToSpeech>
#include <QTextToSpeechPlugin>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQTextToSpeech: public QTextToSpeech
{
public:
	MyQTextToSpeech(QObject *parent = nullptr) : QTextToSpeech(parent) {};
	MyQTextToSpeech(const QString &engine, QObject *parent = nullptr) : QTextToSpeech(engine, parent) {};
	void Signal_LocaleChanged(const QLocale & locale) { callbackQTextToSpeech_LocaleChanged(this, const_cast<QLocale*>(&locale)); };
	void pause() { callbackQTextToSpeech_Pause(this); };
	void Signal_PitchChanged(double pitch) { callbackQTextToSpeech_PitchChanged(this, pitch); };
	void Signal_RateChanged(double rate) { callbackQTextToSpeech_RateChanged(this, rate); };
	void resume() { callbackQTextToSpeech_Resume(this); };
	void say(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtSpeech_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQTextToSpeech_Say(this, textPacked); };
	void setLocale(const QLocale & locale) { callbackQTextToSpeech_SetLocale(this, const_cast<QLocale*>(&locale)); };
	void setPitch(double pitch) { callbackQTextToSpeech_SetPitch(this, pitch); };
	void setRate(double rate) { callbackQTextToSpeech_SetRate(this, rate); };
	void setVolume(double volume) { callbackQTextToSpeech_SetVolume(this, volume); };
	void Signal_StateChanged(QTextToSpeech::State state) { callbackQTextToSpeech_StateChanged(this, state); };
	void stop() { callbackQTextToSpeech_Stop(this); };
	void Signal_VolumeChanged(int volume) { callbackQTextToSpeech_VolumeChanged(this, volume); };
	bool event(QEvent * e) { return callbackQTextToSpeech_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextToSpeech_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQTextToSpeech_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextToSpeech_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextToSpeech_CustomEvent(this, event); };
	void deleteLater() { callbackQTextToSpeech_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextToSpeech_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextToSpeech_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSpeech_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQTextToSpeech_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextToSpeech_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTextToSpeech_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

struct QtSpeech_PackedString QTextToSpeech_QTextToSpeech_AvailableEngines()
{
	return ({ QByteArray tb1cc61 = QTextToSpeech::availableEngines().join("|").toUtf8(); QtSpeech_PackedString { const_cast<char*>(tb1cc61.prepend("WHITESPACE").constData()+10), tb1cc61.size()-10 }; });
}

void* QTextToSpeech_NewQTextToSpeech(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QWindow*>(parent));
	} else {
		return new MyQTextToSpeech(static_cast<QObject*>(parent));
	}
}

void* QTextToSpeech_NewQTextToSpeech2(char* engine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString(engine), static_cast<QWindow*>(parent));
	} else {
		return new MyQTextToSpeech(QString(engine), static_cast<QObject*>(parent));
	}
}

void QTextToSpeech_ConnectLocaleChanged(void* ptr)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(const QLocale &)>(&QTextToSpeech::localeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(const QLocale &)>(&MyQTextToSpeech::Signal_LocaleChanged));
}

void QTextToSpeech_DisconnectLocaleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(const QLocale &)>(&QTextToSpeech::localeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(const QLocale &)>(&MyQTextToSpeech::Signal_LocaleChanged));
}

void QTextToSpeech_LocaleChanged(void* ptr, void* locale)
{
	static_cast<QTextToSpeech*>(ptr)->localeChanged(*static_cast<QLocale*>(locale));
}

void QTextToSpeech_Pause(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "pause");
}

void QTextToSpeech_PauseDefault(void* ptr)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::pause();
}

void QTextToSpeech_ConnectPitchChanged(void* ptr)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::pitchChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_PitchChanged));
}

void QTextToSpeech_DisconnectPitchChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::pitchChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_PitchChanged));
}

void QTextToSpeech_PitchChanged(void* ptr, double pitch)
{
	static_cast<QTextToSpeech*>(ptr)->pitchChanged(pitch);
}

void QTextToSpeech_ConnectRateChanged(void* ptr)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::rateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_RateChanged));
}

void QTextToSpeech_DisconnectRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::rateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_RateChanged));
}

void QTextToSpeech_RateChanged(void* ptr, double rate)
{
	static_cast<QTextToSpeech*>(ptr)->rateChanged(rate);
}

void QTextToSpeech_Resume(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "resume");
}

void QTextToSpeech_ResumeDefault(void* ptr)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::resume();
}

void QTextToSpeech_Say(void* ptr, char* text)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "say", Q_ARG(QString, QString(text)));
}

void QTextToSpeech_SayDefault(void* ptr, char* text)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::say(QString(text));
}

void QTextToSpeech_SetLocale(void* ptr, void* locale)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setLocale", Q_ARG(QLocale, *static_cast<QLocale*>(locale)));
}

void QTextToSpeech_SetLocaleDefault(void* ptr, void* locale)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setLocale(*static_cast<QLocale*>(locale));
}

void QTextToSpeech_SetPitch(void* ptr, double pitch)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setPitch", Q_ARG(double, pitch));
}

void QTextToSpeech_SetPitchDefault(void* ptr, double pitch)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setPitch(pitch);
}

void QTextToSpeech_SetRate(void* ptr, double rate)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setRate", Q_ARG(double, rate));
}

void QTextToSpeech_SetRateDefault(void* ptr, double rate)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setRate(rate);
}

void QTextToSpeech_SetVolume(void* ptr, double volume)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setVolume", Q_ARG(double, volume));
}

void QTextToSpeech_SetVolumeDefault(void* ptr, double volume)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setVolume(volume);
}

void QTextToSpeech_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(QTextToSpeech::State)>(&QTextToSpeech::stateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(QTextToSpeech::State)>(&MyQTextToSpeech::Signal_StateChanged));
}

void QTextToSpeech_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(QTextToSpeech::State)>(&QTextToSpeech::stateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(QTextToSpeech::State)>(&MyQTextToSpeech::Signal_StateChanged));
}

void QTextToSpeech_StateChanged(void* ptr, long long state)
{
	static_cast<QTextToSpeech*>(ptr)->stateChanged(static_cast<QTextToSpeech::State>(state));
}

void QTextToSpeech_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "stop");
}

void QTextToSpeech_StopDefault(void* ptr)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::stop();
}

void QTextToSpeech_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(int)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(int)>(&MyQTextToSpeech::Signal_VolumeChanged));
}

void QTextToSpeech_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(int)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(int)>(&MyQTextToSpeech::Signal_VolumeChanged));
}

void QTextToSpeech_VolumeChanged(void* ptr, int volume)
{
	static_cast<QTextToSpeech*>(ptr)->volumeChanged(volume);
}

void* QTextToSpeech_Locale(void* ptr)
{
	return new QLocale(static_cast<QTextToSpeech*>(ptr)->locale());
}

struct QtSpeech_PackedList QTextToSpeech_AvailableLocales(void* ptr)
{
	return ({ QVector<QLocale>* tmpValue = new QVector<QLocale>(static_cast<QTextToSpeech*>(ptr)->availableLocales()); QtSpeech_PackedList { tmpValue, tmpValue->size() }; });
}

long long QTextToSpeech_State(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->state();
}

double QTextToSpeech_Pitch(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->pitch();
}

double QTextToSpeech_Rate(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->rate();
}

double QTextToSpeech_Volume(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->volume();
}

void* QTextToSpeech___availableLocales_atList(void* ptr, int i)
{
	return new QLocale(static_cast<QVector<QLocale>*>(ptr)->at(i));
}

void QTextToSpeech___availableLocales_setList(void* ptr, void* i)
{
	static_cast<QVector<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QTextToSpeech___availableLocales_newList(void* ptr)
{
	return new QVector<QLocale>;
}

void* QTextToSpeech___availableVoices_newList(void* ptr)
{
	return new QVector<QVoice>;
}

void* QTextToSpeech___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QTextToSpeech___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextToSpeech___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QTextToSpeech___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTextToSpeech___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QTextToSpeech___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTextToSpeech___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QTextToSpeech___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTextToSpeech___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTextToSpeech___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QTextToSpeech___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QTextToSpeech_EventDefault(void* ptr, void* e)
{
		return static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::event(static_cast<QEvent*>(e));
}

char QTextToSpeech_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QTextToSpeech_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::childEvent(static_cast<QChildEvent*>(event));
}

void QTextToSpeech_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextToSpeech_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::customEvent(static_cast<QEvent*>(event));
}

void QTextToSpeech_DeleteLaterDefault(void* ptr)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::deleteLater();
}

void QTextToSpeech_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextToSpeech_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QTextToSpeech_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::metaObject());
}

class MyQTextToSpeechPlugin: public QTextToSpeechPlugin
{
public:
	 ~MyQTextToSpeechPlugin() { callbackQTextToSpeechPlugin_DestroyQTextToSpeechPlugin(this); };
};

void QTextToSpeechPlugin_DestroyQTextToSpeechPlugin(void* ptr)
{
	static_cast<QTextToSpeechPlugin*>(ptr)->~QTextToSpeechPlugin();
}

void QTextToSpeechPlugin_DestroyQTextToSpeechPluginDefault(void* ptr)
{

}

void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QTextToSpeechPlugin___createTextToSpeechEngine_parameters_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtSpeech_PackedList QTextToSpeechPlugin___createTextToSpeechEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtSpeech_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSpeech_PackedString QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtSpeech_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

