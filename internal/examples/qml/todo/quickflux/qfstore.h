#ifndef QFSTORE_H
#define QFSTORE_H

#include <QObject>
#include <QQmlListProperty>
#include <QJSValue>
#include <QPointer>
#include <QQmlParserStatus>
#include "qfactioncreator.h"
#include "qfdispatcher.h"

class QFStore : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* bindSource READ bindSource WRITE setBindSource NOTIFY bindSourceChanged)
    Q_PROPERTY(QQmlListProperty<QObject> children READ children)
    Q_PROPERTY(QQmlListProperty<QObject> redispatchTargets READ redispatchTargets)
    Q_PROPERTY(bool filterFunctionEnabled MEMBER m_filterFunctionEnabled NOTIFY filterFunctionEnabledChanged)

    Q_CLASSINFO("DefaultProperty", "children")


public:
    explicit QFStore(QObject *parent = 0);
    QQmlListProperty<QObject> children();

    QObject* bindSource() const;
    void setBindSource(QObject* source);

    QQmlListProperty<QObject> redispatchTargets();

signals:
    void dispatched(QString type, QJSValue message);

    void bindSourceChanged();

    void filterFunctionEnabledChanged();

public slots:
    void dispatch(QString type, QJSValue message = QJSValue());

    void bind(QObject* source);


protected:
    void classBegin();
    void componentComplete();

private slots:
    void setup();

private:
    QObjectList m_children;

    QPointer<QObject> m_bindSource;

    QPointer<QFActionCreator> m_actionCreator;

    QPointer<QFDispatcher> m_dispatcher;

    QObjectList m_redispatchTargets;

    bool m_filterFunctionEnabled;

};

#endif // QFSTORE_H
