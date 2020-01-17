// +build !minimal

#define protected public
#define private public

#include "speech.h"
#include "_cgo_export.h"

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
#include <QLayout>
#include <QLocale>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QString>
#include <QTextToSpeech>
#include <QTextToSpeechEngine>
#include <QTextToSpeechPlugin>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QVoice>
#include <QWidget>
#include <QWindow>

class MyQTextToSpeech: public QTextToSpeech
{
public:
	MyQTextToSpeech(QObject *parent = Q_NULLPTR) : QTextToSpeech(parent) {QTextToSpeech_QTextToSpeech_QRegisterMetaType();};
	MyQTextToSpeech(const QString &engine, QObject *parent = Q_NULLPTR) : QTextToSpeech(engine, parent) {QTextToSpeech_QTextToSpeech_QRegisterMetaType();};
	void Signal_LocaleChanged(const QLocale & locale) { callbackQTextToSpeech_LocaleChanged(this, const_cast<QLocale*>(&locale)); };
	void pause() { callbackQTextToSpeech_Pause(this); };
	void Signal_PitchChanged(double pitch) { callbackQTextToSpeech_PitchChanged(this, pitch); };
	void Signal_RateChanged(double rate) { callbackQTextToSpeech_RateChanged(this, rate); };
	void resume() { callbackQTextToSpeech_Resume(this); };
	void say(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtSpeech_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextToSpeech_Say(this, textPacked); };
	void setLocale(const QLocale & locale) { callbackQTextToSpeech_SetLocale(this, const_cast<QLocale*>(&locale)); };
	void setPitch(double pitch) { callbackQTextToSpeech_SetPitch(this, pitch); };
	void setRate(double rate) { callbackQTextToSpeech_SetRate(this, rate); };
	void setVoice(const QVoice & voice) { callbackQTextToSpeech_SetVoice(this, const_cast<QVoice*>(&voice)); };
	void setVolume(double volume) { callbackQTextToSpeech_SetVolume(this, volume); };
	void Signal_StateChanged(QTextToSpeech::State state) { callbackQTextToSpeech_StateChanged(this, state); };
	void stop() { callbackQTextToSpeech_Stop(this); };
	void Signal_VolumeChanged(int volume) { callbackQTextToSpeech_VolumeChanged(this, volume); };
	void Signal_VolumeChanged2(double volume) { callbackQTextToSpeech_VolumeChanged2(this, volume); };
	void childEvent(QChildEvent * event) { callbackQTextToSpeech_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextToSpeech_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextToSpeech_CustomEvent(this, event); };
	void deleteLater() { callbackQTextToSpeech_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextToSpeech_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextToSpeech_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextToSpeech_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextToSpeech_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTextToSpeech_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSpeech_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextToSpeech_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextToSpeech_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextToSpeech*)
Q_DECLARE_METATYPE(MyQTextToSpeech*)

int QTextToSpeech_QTextToSpeech_QRegisterMetaType(){qRegisterMetaType<QTextToSpeech*>(); return qRegisterMetaType<MyQTextToSpeech*>();}

void* QTextToSpeech_NewQTextToSpeech(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
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
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QMediaServiceProviderPlugin*>(parent));
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
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(static_cast<QWindow*>(parent));
	} else {
		return new MyQTextToSpeech(static_cast<QObject*>(parent));
	}
}

void* QTextToSpeech_NewQTextToSpeech2(struct QtSpeech_PackedString engine, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQTextToSpeech(QString::fromUtf8(engine.data, engine.len), static_cast<QObject*>(parent));
	}
}

struct QtSpeech_PackedString QTextToSpeech_QTextToSpeech_AvailableEngines()
{
	return ({ QByteArray* tb1cc61 = new QByteArray(QTextToSpeech::availableEngines().join("¡¦!").toUtf8()); QtSpeech_PackedString { const_cast<char*>(tb1cc61->prepend("WHITESPACE").constData()+10), tb1cc61->size()-10, tb1cc61 }; });
}

struct QtSpeech_PackedList QTextToSpeech_AvailableLocales(void* ptr)
{
	return ({ QVector<QLocale>* tmpValue3aad13 = new QVector<QLocale>(static_cast<QTextToSpeech*>(ptr)->availableLocales()); QtSpeech_PackedList { tmpValue3aad13, tmpValue3aad13->size() }; });
}

struct QtSpeech_PackedList QTextToSpeech_AvailableVoices(void* ptr)
{
	return ({ QVector<QVoice>* tmpValue05b1bf = new QVector<QVoice>(static_cast<QTextToSpeech*>(ptr)->availableVoices()); QtSpeech_PackedList { tmpValue05b1bf, tmpValue05b1bf->size() }; });
}

void* QTextToSpeech_Locale(void* ptr)
{
	return new QLocale(static_cast<QTextToSpeech*>(ptr)->locale());
}

void QTextToSpeech_ConnectLocaleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(const QLocale &)>(&QTextToSpeech::localeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(const QLocale &)>(&MyQTextToSpeech::Signal_LocaleChanged), static_cast<Qt::ConnectionType>(t));
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

double QTextToSpeech_Pitch(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->pitch();
}

void QTextToSpeech_ConnectPitchChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::pitchChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_PitchChanged), static_cast<Qt::ConnectionType>(t));
}

void QTextToSpeech_DisconnectPitchChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::pitchChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_PitchChanged));
}

void QTextToSpeech_PitchChanged(void* ptr, double pitch)
{
	static_cast<QTextToSpeech*>(ptr)->pitchChanged(pitch);
}

double QTextToSpeech_Rate(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->rate();
}

void QTextToSpeech_ConnectRateChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::rateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_RateChanged), static_cast<Qt::ConnectionType>(t));
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

void QTextToSpeech_Say(void* ptr, struct QtSpeech_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "say", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextToSpeech_SayDefault(void* ptr, struct QtSpeech_PackedString text)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::say(QString::fromUtf8(text.data, text.len));
}

void QTextToSpeech_SetLocale(void* ptr, void* locale)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setLocale", Q_ARG(const QLocale, *static_cast<QLocale*>(locale)));
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

void QTextToSpeech_SetVoice(void* ptr, void* voice)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setVoice", Q_ARG(const QVoice, *static_cast<QVoice*>(voice)));
}

void QTextToSpeech_SetVoiceDefault(void* ptr, void* voice)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setVoice(*static_cast<QVoice*>(voice));
}

void QTextToSpeech_SetVolume(void* ptr, double volume)
{
	QMetaObject::invokeMethod(static_cast<QTextToSpeech*>(ptr), "setVolume", Q_ARG(double, volume));
}

void QTextToSpeech_SetVolumeDefault(void* ptr, double volume)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::setVolume(volume);
}

long long QTextToSpeech_State(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->state();
}

void QTextToSpeech_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QTextToSpeech::State>();
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(QTextToSpeech::State)>(&QTextToSpeech::stateChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(QTextToSpeech::State)>(&MyQTextToSpeech::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
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

void* QTextToSpeech_Voice(void* ptr)
{
	return new QVoice(static_cast<QTextToSpeech*>(ptr)->voice());
}

double QTextToSpeech_Volume(void* ptr)
{
	return static_cast<QTextToSpeech*>(ptr)->volume();
}

void QTextToSpeech_ConnectVolumeChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(int)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(int)>(&MyQTextToSpeech::Signal_VolumeChanged), static_cast<Qt::ConnectionType>(t));
}

void QTextToSpeech_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(int)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(int)>(&MyQTextToSpeech::Signal_VolumeChanged));
}

void QTextToSpeech_VolumeChanged(void* ptr, int volume)
{
	static_cast<QTextToSpeech*>(ptr)->volumeChanged(volume);
}

void QTextToSpeech_ConnectVolumeChanged2(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_VolumeChanged2), static_cast<Qt::ConnectionType>(t));
}

void QTextToSpeech_DisconnectVolumeChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeech*>(ptr), static_cast<void (QTextToSpeech::*)(double)>(&QTextToSpeech::volumeChanged), static_cast<MyQTextToSpeech*>(ptr), static_cast<void (MyQTextToSpeech::*)(double)>(&MyQTextToSpeech::Signal_VolumeChanged2));
}

void QTextToSpeech_VolumeChanged2(void* ptr, double volume)
{
	static_cast<QTextToSpeech*>(ptr)->volumeChanged(volume);
}

void* QTextToSpeech___availableLocales_atList(void* ptr, int i)
{
	return new QLocale(({QLocale tmp = static_cast<QVector<QLocale>*>(ptr)->at(i); if (i == static_cast<QVector<QLocale>*>(ptr)->size()-1) { static_cast<QVector<QLocale>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextToSpeech___availableLocales_setList(void* ptr, void* i)
{
	static_cast<QVector<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QTextToSpeech___availableLocales_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLocale>();
}

void* QTextToSpeech___availableVoices_atList(void* ptr, int i)
{
	return new QVoice(({QVoice tmp = static_cast<QVector<QVoice>*>(ptr)->at(i); if (i == static_cast<QVector<QVoice>*>(ptr)->size()-1) { static_cast<QVector<QVoice>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextToSpeech___availableVoices_setList(void* ptr, void* i)
{
	static_cast<QVector<QVoice>*>(ptr)->append(*static_cast<QVoice*>(i));
}

void* QTextToSpeech___availableVoices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QVoice>();
}

void* QTextToSpeech___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeech___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QTextToSpeech___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextToSpeech___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextToSpeech___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTextToSpeech___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeech___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTextToSpeech___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeech___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeech___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
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

char QTextToSpeech_EventDefault(void* ptr, void* e)
{
		return static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::event(static_cast<QEvent*>(e));
}

char QTextToSpeech_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextToSpeech_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::metaObject());
}

void QTextToSpeech_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeech*>(ptr)->QTextToSpeech::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQTextToSpeechEngine: public QTextToSpeechEngine
{
public:
	MyQTextToSpeechEngine(QObject *parent = Q_NULLPTR) : QTextToSpeechEngine(parent) {QTextToSpeechEngine_QTextToSpeechEngine_QRegisterMetaType();};
	QVector<QLocale> availableLocales() const { return ({ QVector<QLocale>* tmpP = static_cast<QVector<QLocale>*>(callbackQTextToSpeechEngine_AvailableLocales(const_cast<void*>(static_cast<const void*>(this)))); QVector<QLocale> tmpV = *tmpP; tmpP->~QVector(); free(tmpP); tmpV; }); };
	QVector<QVoice> availableVoices() const { return ({ QVector<QVoice>* tmpP = static_cast<QVector<QVoice>*>(callbackQTextToSpeechEngine_AvailableVoices(const_cast<void*>(static_cast<const void*>(this)))); QVector<QVoice> tmpV = *tmpP; tmpP->~QVector(); free(tmpP); tmpV; }); };
	QLocale locale() const { return *static_cast<QLocale*>(callbackQTextToSpeechEngine_Locale(const_cast<void*>(static_cast<const void*>(this)))); };
	void pause() { callbackQTextToSpeechEngine_Pause(this); };
	double pitch() const { return callbackQTextToSpeechEngine_Pitch(const_cast<void*>(static_cast<const void*>(this))); };
	double rate() const { return callbackQTextToSpeechEngine_Rate(const_cast<void*>(static_cast<const void*>(this))); };
	void resume() { callbackQTextToSpeechEngine_Resume(this); };
	void say(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtSpeech_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextToSpeechEngine_Say(this, textPacked); };
	bool setLocale(const QLocale & locale) { return callbackQTextToSpeechEngine_SetLocale(this, const_cast<QLocale*>(&locale)) != 0; };
	bool setPitch(double pitch) { return callbackQTextToSpeechEngine_SetPitch(this, pitch) != 0; };
	bool setRate(double rate) { return callbackQTextToSpeechEngine_SetRate(this, rate) != 0; };
	bool setVoice(const QVoice & voice) { return callbackQTextToSpeechEngine_SetVoice(this, const_cast<QVoice*>(&voice)) != 0; };
	bool setVolume(double volume) { return callbackQTextToSpeechEngine_SetVolume(this, volume) != 0; };
	QTextToSpeech::State state() const { return static_cast<QTextToSpeech::State>(callbackQTextToSpeechEngine_State(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_StateChanged(QTextToSpeech::State state) { callbackQTextToSpeechEngine_StateChanged(this, state); };
	void stop() { callbackQTextToSpeechEngine_Stop(this); };
	QVoice voice() const { return *static_cast<QVoice*>(callbackQTextToSpeechEngine_Voice(const_cast<void*>(static_cast<const void*>(this)))); };
	double volume() const { return callbackQTextToSpeechEngine_Volume(const_cast<void*>(static_cast<const void*>(this))); };
	void childEvent(QChildEvent * event) { callbackQTextToSpeechEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextToSpeechEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextToSpeechEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQTextToSpeechEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextToSpeechEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextToSpeechEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextToSpeechEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextToSpeechEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTextToSpeechEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSpeech_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextToSpeechEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextToSpeechEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextToSpeechEngine*)
Q_DECLARE_METATYPE(MyQTextToSpeechEngine*)

int QTextToSpeechEngine_QTextToSpeechEngine_QRegisterMetaType(){qRegisterMetaType<QTextToSpeechEngine*>(); return qRegisterMetaType<MyQTextToSpeechEngine*>();}

void* QTextToSpeechEngine_NewQTextToSpeechEngine(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextToSpeechEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQTextToSpeechEngine(static_cast<QObject*>(parent));
	}
}

struct QtSpeech_PackedList QTextToSpeechEngine_AvailableLocales(void* ptr)
{
	return ({ QVector<QLocale>* tmpValued4fc62 = new QVector<QLocale>(static_cast<QTextToSpeechEngine*>(ptr)->availableLocales()); QtSpeech_PackedList { tmpValued4fc62, tmpValued4fc62->size() }; });
}

struct QtSpeech_PackedList QTextToSpeechEngine_AvailableVoices(void* ptr)
{
	return ({ QVector<QVoice>* tmpValue7c0049 = new QVector<QVoice>(static_cast<QTextToSpeechEngine*>(ptr)->availableVoices()); QtSpeech_PackedList { tmpValue7c0049, tmpValue7c0049->size() }; });
}

void* QTextToSpeechEngine_QTextToSpeechEngine_CreateVoice(struct QtSpeech_PackedString name, long long gender, long long age, void* data)
{
	return new QVoice(QTextToSpeechEngine::createVoice(QString::fromUtf8(name.data, name.len), static_cast<QVoice::Gender>(gender), static_cast<QVoice::Age>(age), *static_cast<QVariant*>(data)));
}

void* QTextToSpeechEngine_Locale(void* ptr)
{
	return new QLocale(static_cast<QTextToSpeechEngine*>(ptr)->locale());
}

void QTextToSpeechEngine_Pause(void* ptr)
{
	static_cast<QTextToSpeechEngine*>(ptr)->pause();
}

double QTextToSpeechEngine_Pitch(void* ptr)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->pitch();
}

double QTextToSpeechEngine_Rate(void* ptr)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->rate();
}

void QTextToSpeechEngine_Resume(void* ptr)
{
	static_cast<QTextToSpeechEngine*>(ptr)->resume();
}

void QTextToSpeechEngine_Say(void* ptr, struct QtSpeech_PackedString text)
{
	static_cast<QTextToSpeechEngine*>(ptr)->say(QString::fromUtf8(text.data, text.len));
}

char QTextToSpeechEngine_SetLocale(void* ptr, void* locale)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

char QTextToSpeechEngine_SetPitch(void* ptr, double pitch)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->setPitch(pitch);
}

char QTextToSpeechEngine_SetRate(void* ptr, double rate)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->setRate(rate);
}

char QTextToSpeechEngine_SetVoice(void* ptr, void* voice)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->setVoice(*static_cast<QVoice*>(voice));
}

char QTextToSpeechEngine_SetVolume(void* ptr, double volume)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->setVolume(volume);
}

long long QTextToSpeechEngine_State(void* ptr)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->state();
}

void QTextToSpeechEngine_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QTextToSpeech::State>();
	QObject::connect(static_cast<QTextToSpeechEngine*>(ptr), static_cast<void (QTextToSpeechEngine::*)(QTextToSpeech::State)>(&QTextToSpeechEngine::stateChanged), static_cast<MyQTextToSpeechEngine*>(ptr), static_cast<void (MyQTextToSpeechEngine::*)(QTextToSpeech::State)>(&MyQTextToSpeechEngine::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
}

void QTextToSpeechEngine_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextToSpeechEngine*>(ptr), static_cast<void (QTextToSpeechEngine::*)(QTextToSpeech::State)>(&QTextToSpeechEngine::stateChanged), static_cast<MyQTextToSpeechEngine*>(ptr), static_cast<void (MyQTextToSpeechEngine::*)(QTextToSpeech::State)>(&MyQTextToSpeechEngine::Signal_StateChanged));
}

void QTextToSpeechEngine_StateChanged(void* ptr, long long state)
{
	static_cast<QTextToSpeechEngine*>(ptr)->stateChanged(static_cast<QTextToSpeech::State>(state));
}

void QTextToSpeechEngine_Stop(void* ptr)
{
	static_cast<QTextToSpeechEngine*>(ptr)->stop();
}

void* QTextToSpeechEngine_Voice(void* ptr)
{
	return new QVoice(static_cast<QTextToSpeechEngine*>(ptr)->voice());
}

void* QTextToSpeechEngine_QTextToSpeechEngine_VoiceData(void* voice)
{
	return new QVariant(QTextToSpeechEngine::voiceData(*static_cast<QVoice*>(voice)));
}

double QTextToSpeechEngine_Volume(void* ptr)
{
	return static_cast<QTextToSpeechEngine*>(ptr)->volume();
}

void* QTextToSpeechEngine___availableLocales_atList(void* ptr, int i)
{
	return new QLocale(({QLocale tmp = static_cast<QVector<QLocale>*>(ptr)->at(i); if (i == static_cast<QVector<QLocale>*>(ptr)->size()-1) { static_cast<QVector<QLocale>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextToSpeechEngine___availableLocales_setList(void* ptr, void* i)
{
	static_cast<QVector<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QTextToSpeechEngine___availableLocales_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLocale>();
}

void* QTextToSpeechEngine___availableVoices_atList(void* ptr, int i)
{
	return new QVoice(({QVoice tmp = static_cast<QVector<QVoice>*>(ptr)->at(i); if (i == static_cast<QVector<QVoice>*>(ptr)->size()-1) { static_cast<QVector<QVoice>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextToSpeechEngine___availableVoices_setList(void* ptr, void* i)
{
	static_cast<QVector<QVoice>*>(ptr)->append(*static_cast<QVoice*>(i));
}

void* QTextToSpeechEngine___availableVoices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QVoice>();
}

void* QTextToSpeechEngine___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeechEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeechEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QTextToSpeechEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextToSpeechEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextToSpeechEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTextToSpeechEngine___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeechEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeechEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTextToSpeechEngine___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextToSpeechEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QTextToSpeechEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QTextToSpeechEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QTextToSpeechEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextToSpeechEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::customEvent(static_cast<QEvent*>(event));
}

void QTextToSpeechEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::deleteLater();
}

void QTextToSpeechEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QTextToSpeechEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::event(static_cast<QEvent*>(e));
}

char QTextToSpeechEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextToSpeechEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::metaObject());
}

void QTextToSpeechEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QTextToSpeechEngine*>(ptr)->QTextToSpeechEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQTextToSpeechPlugin: public QTextToSpeechPlugin
{
public:
	QTextToSpeechEngine * createTextToSpeechEngine(const QVariantMap & parameters, QObject * parent, QString * errorString) const { QByteArray* tc8b6bd = new QByteArray(errorString->toUtf8()); QtSpeech_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd->prepend("WHITESPACE").constData()+10), tc8b6bd->size()-10, tc8b6bd };return static_cast<QTextToSpeechEngine*>(callbackQTextToSpeechPlugin_CreateTextToSpeechEngine(const_cast<void*>(static_cast<const void*>(this)), ({ QMap<QString, QVariant>* tmpValuef1f142 = new QMap<QString, QVariant>(parameters); QtSpeech_PackedList { tmpValuef1f142, tmpValuef1f142->size() }; }), parent, errorStringPacked)); };
};

Q_DECLARE_METATYPE(QTextToSpeechPlugin*)
Q_DECLARE_METATYPE(MyQTextToSpeechPlugin*)

int QTextToSpeechPlugin_QTextToSpeechPlugin_QRegisterMetaType(){qRegisterMetaType<QTextToSpeechPlugin*>(); return qRegisterMetaType<MyQTextToSpeechPlugin*>();}

void* QTextToSpeechPlugin_CreateTextToSpeechEngine(void* ptr, void* parameters, void* parent, struct QtSpeech_PackedString errorString)
{
	return static_cast<QTextToSpeechPlugin*>(ptr)->createTextToSpeechEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QObject*>(parent), new QString(QString::fromUtf8(errorString.data, errorString.len)));
}

void* QTextToSpeechPlugin_CreateTextToSpeechEngineDefault(void* ptr, void* parameters, void* parent, struct QtSpeech_PackedString errorString)
{
		return static_cast<QTextToSpeechPlugin*>(ptr)->QTextToSpeechPlugin::createTextToSpeechEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QObject*>(parent), new QString(QString::fromUtf8(errorString.data, errorString.len)));
}

void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(void* ptr, struct QtSpeech_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QTextToSpeechPlugin___createTextToSpeechEngine_parameters_setList(void* ptr, struct QtSpeech_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtSpeech_PackedList QTextToSpeechPlugin___createTextToSpeechEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtSpeech_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtSpeech_PackedString QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtSpeech_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_setList(void* ptr, struct QtSpeech_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

long long QVoice_Age(void* ptr)
{
	return static_cast<QVoice*>(ptr)->age();
}

struct QtSpeech_PackedString QVoice_QVoice_AgeName(long long age)
{
	return ({ QByteArray* t3dc0ed = new QByteArray(QVoice::ageName(static_cast<QVoice::Age>(age)).toUtf8()); QtSpeech_PackedString { const_cast<char*>(t3dc0ed->prepend("WHITESPACE").constData()+10), t3dc0ed->size()-10, t3dc0ed }; });
}

long long QVoice_Gender(void* ptr)
{
	return static_cast<QVoice*>(ptr)->gender();
}

struct QtSpeech_PackedString QVoice_QVoice_GenderName(long long gender)
{
	return ({ QByteArray* t94d66c = new QByteArray(QVoice::genderName(static_cast<QVoice::Gender>(gender)).toUtf8()); QtSpeech_PackedString { const_cast<char*>(t94d66c->prepend("WHITESPACE").constData()+10), t94d66c->size()-10, t94d66c }; });
}

struct QtSpeech_PackedString QVoice_Name(void* ptr)
{
	return ({ QByteArray* t889245 = new QByteArray(static_cast<QVoice*>(ptr)->name().toUtf8()); QtSpeech_PackedString { const_cast<char*>(t889245->prepend("WHITESPACE").constData()+10), t889245->size()-10, t889245 }; });
}

