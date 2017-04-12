#include <QtCore>
#include <QUuid>
#include <QFAppDispatcher>
#include "qfappscriptrunnable.h"
#include "qfappscriptdispatcherwrapper.h"

QFAppScriptRunnable::QFAppScriptRunnable(QObject *parent) : QObject(parent)
{
    m_engine = 0;
    m_next = 0;
    m_isSignalCondition = false;
    m_isOnceOnly = true;
}

QFAppScriptRunnable::~QFAppScriptRunnable()
{
    release();
}

QJSValue QFAppScriptRunnable::script() const
{
    return m_script;
}

void QFAppScriptRunnable::setScript(const QJSValue &script)
{
    m_script = script;
}

QString QFAppScriptRunnable::type() const
{
    return m_type;
}

void QFAppScriptRunnable::setType(const QString &type)
{
    m_type = type;
}

bool QFAppScriptRunnable::isOnceOnly() const
{
    return m_isOnceOnly;
}

void QFAppScriptRunnable::setIsOnceOnly(bool isOnceOnly)
{
    m_isOnceOnly = isOnceOnly;
}

void QFAppScriptRunnable::setEngine(QQmlEngine* engine)
{
    m_engine = engine;
}

void QFAppScriptRunnable::release()
{
    if (!m_condition.isNull() &&
            m_condition.isObject() &&
        m_condition.hasProperty("disconnect")) {

        QJSValue disconnect = m_condition.property("disconnect");
        QJSValueList args;
        args << m_callback;

        disconnect.callWithInstance(m_condition,args);
    }

    m_condition = QJSValue();
    m_callback = QJSValue();
}

void QFAppScriptRunnable::run(QJSValue message)
{
    QJSValueList args;
    if (m_isSignalCondition &&
        message.hasProperty("length")) {
        int count = message.property("length").toInt();
        for (int i = 0 ; i < count;i++) {
            args << message.property(i);
        }
    } else {
        args << message;
    }
    QJSValue ret = m_script.call(args);

    if (ret.isError()) {
        QString message = QString("%1:%2: %3: %4")
                          .arg(ret.property("fileName").toString())
                          .arg(ret.property("lineNumber").toString())
                          .arg(ret.property("name").toString())
                          .arg(ret.property("message").toString());
        qWarning() << message;
    }

}

QFAppScriptRunnable *QFAppScriptRunnable::then(QJSValue condition,QJSValue script)
{
    QFAppScriptRunnable* runnable = new QFAppScriptRunnable(this);
    runnable->setEngine(m_engine.data());
    runnable->setCondition(condition);
    runnable->setScript(script);
    setNext(runnable);
    return runnable;
}

QFAppScriptRunnable *QFAppScriptRunnable::next() const
{
    return m_next;
}

void QFAppScriptRunnable::setNext(QFAppScriptRunnable *next)
{
    m_next = next;
}

void QFAppScriptRunnable::setCondition(QJSValue condition)
{
    m_condition = condition;

    if (condition.isString()) {
        setType(condition.toString());
        m_isSignalCondition = false;
    } else if (condition.isObject() && condition.hasProperty("connect")) {
        Q_ASSERT(!m_engine.isNull());

        QString type = QString("QuickFlux.AppScript.%1").arg(QUuid::createUuid().toString());
        setType(type);

        QString generator = "function(dispatcher) { return function() {dispatcher.dispatch(arguments)}}";
        QFAppDispatcher* dispatcher = QFAppDispatcher::instance(m_engine);
        QFAppScriptDispatcherWrapper * wrapper = new QFAppScriptDispatcherWrapper();
        wrapper->setType(type);
        wrapper->setDispatcher(dispatcher);

        QJSValue generatorFunc = m_engine->evaluate(generator);

        QJSValueList args;
        args << m_engine->newQObject(wrapper);
        QJSValue callback = generatorFunc.call(args);

        args.clear();
        args << callback;

        QJSValue connect = condition.property("connect");
        connect.callWithInstance(condition,args);

        m_callback = callback;
        m_isSignalCondition = true;
    } else {
        qWarning() << "AppScript: Invalid condition type";
    }
}

