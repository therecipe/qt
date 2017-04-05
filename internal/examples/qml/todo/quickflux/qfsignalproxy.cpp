#include <QtCore>
#include <QMetaObject>
#include "qfsignalproxy.h"

QFSignalProxy::QFSignalProxy(QObject *parent) : QObject(parent)
{

}

void QFSignalProxy::bind(QObject *source, int signalIdx, QQmlEngine* engine, QFDispatcher* dispatcher)
{
    const int memberOffset = QObject::staticMetaObject.methodCount();

    QMetaMethod method = source->metaObject()->method(signalIdx);

    parameterTypes = QVector<int>(method.parameterCount());
    parameterNames = QVector<QString>(method.parameterCount());
    type = method.name();
    m_engine = engine;
    m_dispatcher = dispatcher;

    for (int i = 0 ; i < method.parameterCount() ; i++) {
        parameterTypes[i] = method.parameterType(i);
        parameterNames[i] = QString(method.parameterNames().at(i));
    }

    if (!QMetaObject::connect(source, signalIdx, this, memberOffset, Qt::AutoConnection, 0)) {
        qWarning() << "Failed to bind signal";
    }
}

int QFSignalProxy::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    int methodId = QObject::qt_metacall(_c, _id, _a);

    if (methodId < 0) {
        return methodId;
    }

    if (_c == QMetaObject::InvokeMetaMethod) {

        if (methodId == 0) {
            QVariantMap message;

            for (int i = 0 ; i < parameterTypes.count() ; i++) {
                const QMetaType::Type type = static_cast<QMetaType::Type>(parameterTypes.at(i));
                QVariant v;

                if (type == QMetaType::QVariant) {
                    v = *reinterpret_cast<QVariant *>(_a[i + 1]);
                } else {
                    v = QVariant(type, _a[i + 1]);
                }

                message[parameterNames.at(i)] = v;
            }

            dispatch(message);
        }
        methodId--;
    }

    return methodId;
}

void QFSignalProxy::dispatch(const QVariantMap &message)
{
    if (m_engine.isNull() || m_dispatcher.isNull()) {
        return;
    }

    QJSValue value = m_engine->newObject();

    QMapIterator<QString, QVariant> iter(message);
    while (iter.hasNext()) {
        iter.next();
        QJSValue v = m_engine->toScriptValue<QVariant>(iter.value());
        value.setProperty(iter.key(), v);
    }

    m_dispatcher->dispatch(type, value);
}

QFDispatcher *QFSignalProxy::dispatcher() const
{
    return m_dispatcher;
}

void QFSignalProxy::setDispatcher(QFDispatcher *dispatcher)
{
    m_dispatcher = dispatcher;
}
