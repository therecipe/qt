#include "qfappscriptgroup.h"

/*! \qmltype AppScriptGroup
    \inqmlmodule QuickFlux

    AppScriptGroup hold a group of AppScript objects which are mutually exclusive in execution.
    Whatever a AppScript is going to start, it will terminate all other AppScript objects.
    So that only one AppScript is running at a time.

\code

Item {

    AppScript {
        id: script1
        script: {
            // write script here
        }
    }

    AppScript {
        id: script2
        script: {
            // write script here
        }
    }

    AppScriptGroup {
        scripts: [script1, script2]
    }

    Component.onCompleted: {

        script1.run();

        script2.run();

        // At this point, AppScriptGroup will force script1 to terminate since script2 has been started.

    }

}

\endcode

 */

QFAppScriptGroup::QFAppScriptGroup(QQuickItem* parent) : QQuickItem(parent)
{

}

/*! \qmlproperty array AppScriptGroup::scripts
   This property hold an array of AppScript object.
   They are mutually exclusive in execution.

\code

AppScript {
    id: script1
}

AppScript {
    id: script2
}

AppScriptGroup {
    scripts: [script1, script2]
}

\endcode
 */

QJSValue QFAppScriptGroup::scripts() const
{
    return m_scripts;
}

void QFAppScriptGroup::setScripts(const QJSValue &scripts)
{
    for (int i = 0 ; i < objects.count() ; i++) {
        if (objects.at(i).data()) {
            objects.at(i)->disconnect(this);
        }
    }

    objects.clear();
    m_scripts = scripts;

    if (!scripts.isArray()) {
        qWarning() << "AppScriptGroup: Invalid scripts property";
        return;
    }

    int count = scripts.property("length").toInt();

    for (int i = 0 ; i < count ; i++) {
        QJSValue item = scripts.property(i);

        if (!item.isQObject()) {
            qWarning() << "AppScriptGroup: Invalid scripts property";
            continue;
        }

        QFAppScript* object = qobject_cast<QFAppScript*>(item.toQObject());

        if (!object) {
            qWarning() << "AppScriptGroup: Invalid scripts property";
            continue;
        }
        objects << object;
        connect(object,SIGNAL(started()),
                this,SLOT(onStarted()));
    }

    emit scriptsChanged();
}

/*! \qmlmethod AppScriptGroup::exitAll()

  Terminate all AppScript objects

 */

void QFAppScriptGroup::exitAll()
{
    for (int i = 0 ; i < objects.count() ; i++) {
        if (objects.at(i).data()) {
            objects.at(i)->exit();
        }
    }
}

void QFAppScriptGroup::onStarted()
{
    QFAppScript* source = qobject_cast<QFAppScript*>(sender());

    for (int i = 0 ; i < objects.count() ; i++) {
        QPointer<QFAppScript> object = objects.at(i);
        if (object.isNull())
            continue;

        if (object.data() != source) {
            object->exit();
        }
    }

}

