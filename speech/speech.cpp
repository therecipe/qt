// +build !minimal

#define protected public
#define private public

#include "speech.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QList>
#include <QLocale>
#include <QMap>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTextToSpeech>
#include <QTextToSpeechPlugin>
#include <QVariant>
#include <QVector>

class MyQTextToSpeech: public QTextToSpeech
{
public:
	MyQTextToSpeech(QObject *parent) : QTextToSpeech(parent) {};
	MyQTextToSpeech(const QString &engine, QObject *parent) : QTextToSpeech(engine, parent) {};
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
};

struct QtSpeech_PackedString QTextToSpeech_QTextToSpeech_AvailableEngines()
{
	return ({ QByteArray tb1cc61 = QTextToSpeech::availableEngines().join("|").toUtf8(); QtSpeech_PackedString { const_cast<char*>(tb1cc61.prepend("WHITESPACE").constData()+10), tb1cc61.size()-10 }; });
}

void* QTextToSpeech_NewQTextToSpeech(void* parent)
{
	return new MyQTextToSpeech(static_cast<QObject*>(parent));
}

void* QTextToSpeech_NewQTextToSpeech2(char* engine, void* parent)
{
	return new MyQTextToSpeech(QString(engine), static_cast<QObject*>(parent));
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

