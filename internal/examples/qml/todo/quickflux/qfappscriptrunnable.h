#ifndef QFAPPSCRIPTRUNNABLE_H
#define QFAPPSCRIPTRUNNABLE_H

#include <QObject>
#include <QJSValue>
#include <QQmlEngine>
#include <QPointer>

/// QFAppScriptRunnable handles registered callback in QFAppScript (Private class)

class QFAppScriptRunnable : public QObject
{
    Q_OBJECT
public:
    explicit QFAppScriptRunnable(QObject *parent = 0);
    ~QFAppScriptRunnable();

    QJSValue script() const;
    void setScript(const QJSValue &script);

    QString type() const;

    void run(QJSValue message);

    QFAppScriptRunnable *next() const;
    void setNext(QFAppScriptRunnable *next);

    void setCondition(QJSValue condition);

    void setEngine(QQmlEngine* engine);

    void release();

    bool isOnceOnly() const;
    void setIsOnceOnly(bool isOnceOnly);

signals:

public slots:
    QFAppScriptRunnable* then(QJSValue condition,QJSValue value);

private:
    void setType(const QString &type);

    QJSValue m_script;
    QString m_type;
    QFAppScriptRunnable* m_next;
    QPointer<QQmlEngine> m_engine;

    QJSValue m_condition;
    QJSValue m_callback;
    bool m_isSignalCondition;
    bool m_isOnceOnly;
};

#endif // QFAPPSCRIPTRUNNABLE_H
