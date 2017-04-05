#ifndef QFAPPSCRIPT_H
#define QFAPPSCRIPT_H

#include <QObject>
#include <QJSValue>
#include <QQmlScriptString>
#include <QQuickItem>
#include <QQmlParserStatus>
#include <QPointer>
#include <QFAppDispatcher>
#include "qfappscriptrunnable.h"

class QFAppScript : public QQuickItem
{
    Q_OBJECT
    Q_PROPERTY(QQmlScriptString script READ script WRITE setScript NOTIFY scriptChanged)
    Q_PROPERTY(bool running READ running WRITE setRunning NOTIFY runningChanged)
    Q_PROPERTY(QString runWhen READ runWhen WRITE setRunWhen NOTIFY runWhenChanged)
    Q_PROPERTY(QJSValue message READ message NOTIFY messageChanged)
    Q_PROPERTY(int listenerId READ listenerId WRITE setListenerId NOTIFY listenerIdChanged)
    Q_PROPERTY(QList<int> waitFor READ waitFor WRITE setWaitFor NOTIFY waitForChanged)
    Q_PROPERTY(bool autoExit READ autoExit WRITE setAutoExit NOTIFY autoExitChanged)

public:
    explicit QFAppScript(QQuickItem *parent = 0);

    QQmlScriptString script() const;
    void setScript(const QQmlScriptString &script);

    bool running() const;

    QString runWhen() const;
    void setRunWhen(const QString &runWhen);

    QJSValue message() const;
    void setMessage(const QJSValue &message);

    int listenerId() const;
    void setListenerId(int listenerId);

    QList<int> waitFor() const;
    void setWaitFor(const QList<int> &waitFor);

    bool autoExit() const;
    void setAutoExit(bool autoExit);

signals:
    void started();
    void finished(int returnCode);

    void scriptChanged();
    void runningChanged();
    void runWhenChanged();
    void messageChanged();
    void listenerIdChanged();
    void waitForChanged();
    void autoExitChanged();

public slots:
    void exit(int returnCode = 0);
    void run(QJSValue message = QJSValue());

    QFAppScriptRunnable* once(QJSValue condition, QJSValue script);
    void on(QJSValue condition, QJSValue script);

private slots:
    void onDispatched(QString type,QJSValue message);

private:
    virtual void componentComplete();
    void abort();
    void clear();
    void setRunning(bool running);

    void setListenerWaitFor();

    QQmlScriptString m_script;
    QList<QFAppScriptRunnable*> m_runnables;
    QPointer<QFAppDispatcher> m_dispatcher;
    QString m_runWhen;

    bool m_running;
    bool m_processing;

    int m_listenerId;

    bool m_autoExit;

    // The message object passed to run()
    QJSValue m_message;
    QFListener* m_listener;

    QList<int> m_waitFor;
};

#endif // QFAPPSCRIPT_H
