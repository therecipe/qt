#ifndef QFMIDDLEWARE_H
#define QFMIDDLEWARE_H

#include <QQuickItem>
#include <QJSValue>

class QFMiddleware : public QQuickItem
{
    Q_OBJECT
    Q_PROPERTY(bool filterFunctionEnabled MEMBER m_filterFunctionEnabled NOTIFY filterFunctionEnabledChanged)
    Q_PROPERTY(QJSValue _nextCallback READ nextCallback WRITE setNextCallback NOTIFY _nextCallbackChanged)

public:
    QFMiddleware(QQuickItem* parent = 0);

    QJSValue nextCallback() const;
    void setNextCallback(const QJSValue &nextCallback);

signals:
    void dispatched(QString type, QJSValue message);
    void filterFunctionEnabledChanged();
    void _nextCallbackChanged();

public slots:
    void next(QString type, QJSValue message = QJSValue());

private:
    bool m_filterFunctionEnabled;

    QJSValue m_nextCallback;

};

#endif // QFMIDDLEWARE_H
