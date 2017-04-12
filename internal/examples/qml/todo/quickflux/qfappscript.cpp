#include <QtQml>
#include <QVariant>
#include <QPointer>
#include <QQmlExpression>
#include <QtCore>
#include "qfappscript.h"
#include "qfapplistener.h"

/*! \qmltype AppScript
    \inqmlmodule QuickFlux

AppScript is a helper component to handle asynchronous sequential workflow.
The immediate benefit of using AppScript  is the centralisation of code in a place.
Instead of placing them within onXXXX code block in several components in several places.

AppScript also manage the life cycle of registered callbacks. You could remove them by a single function. Nothing will leave in memory.


\b {Remarks: AppScript is not a solution for making Store component}

\l {https://github.com/benlau/quickflux/blob/master/examples/photoalbum/scripts/ImagePickerScript.qml} {Example Code}

\code

    AppScript {
        // Run this script if "Pick Image" button is clicked.
        runWhen: ActionTypes.askToPickPhoto

        script: {
            // Step 1. Open file dialog
            dialog.open();

            // Register a set of callbacks as workflow
            // Registered callbacks will be executed once only per script execution
            once(dialog.onAccepted, function() {

                // Step 2. Once user picked an image, launch preview window and ask for confirmation.
                AppActions.navigateTo(imagePreview,
                                      {source: dialog.fileUrl});

            }).then(ActionTypes.pickPhoto, function(message) {
                // The function of then() is same as once() but it won't
                // trigger the callback until once() is triggered.

                // Step 3. Add picked image to store and go back to previous page.

                PhotoStore.add(String(message.url));

                AppActions.navigateBack();

            }); // <-- You may chain then() function again.

            // Condition to terminate the workflow:
            // Force to terminate if dialog is rejected / or navigateBack is dispatched
            // That will remove all the registered callbacks

            once(dialog.onRejected,exit.bind(this,0));

            once(ActionTypes.navigateBack,exit.bind(this,0));
        }
    }
\endcode

\b {Benefit of using AppScript}

\list
\li Centralize your workflow code in one place
\li Highly integrated with AppDispatcher. The order of callback execution is guaranteed in sequence order.
\li Only one script can be executed at a time. Registered callbacks by previous script will be removed before starting.
\li exit() will remove all the registered callbacks. No callback leave after termination.
\endlist

\b {Why not just use Promise?}

\list
\li 1. You need another library. (e.g QuickPromise)
\li 2. You need to set reject condition per callback/promise
\endlist

\b Explanation:

Coding in a promise way requires you to handle every reject condition correctly. Otherwise,
promise will leave in memory and their behaviour will be unexpected.
AppScript::run() / AppScript::exit() clear all the registered callback completely. You can write less code.

 */

/*! \qmlsignal AppScript::started()
 This signal is emitted when the script is started.
*/

/*! \qmlsignal AppScript::finished(int returnCode)
 This signal is emitted when the script is finished.
 */

QFAppScript::QFAppScript(QQuickItem *parent) : QQuickItem(parent)
{
    m_running = false;
    m_processing = false;
    m_listenerId = 0;
    m_listener = 0;
    m_autoExit = true;
}

/*! \qmlmethod AppScript::exit(int returnCode)

  Terminate current executing script by removing all the registered callbacks.

 */

void QFAppScript::exit(int returnCode)
{
    clear();
    setRunning(false);
    emit finished(returnCode);
}

/*! \qmlmethod AppScript::run()

  Call this function to execute script.
  If the previous script is still running.
  AppScript will terminate previous script by removing all the registered callbacks.

 */

void QFAppScript::run(QJSValue message)
{
    if (m_processing) {
        qWarning() << "AppScript::run(): Don't call run() within script / wait callback";
        return;
    }

    m_processing = true;
    clear();
    setMessage(message);

    if (m_dispatcher.isNull()) {
        qWarning() << "AppScript::run() - Missing AppDispatcher. Aborted.";
        m_processing = false;
        return;
    }

    setRunning(true);

    emit started();

    QQmlExpression expr(m_script);

    if (!m_script.isEmpty()) {
        expr.evaluate();
    }

    if (expr.hasError()) {
        qWarning() << expr.error();
    }

    if (m_runnables.size() == 0) {
        exit(0);
    }

    m_processing = false;
}

/*! \qmlmethod chian AppScript::once(var type, func callback)

  Register a callback to be triggered when a matched message type is dispatched or a signal is emitted.
  It will be triggered once only.
  User should call this function within the script code block.

  The callback will be removed on script termination.

  Moreover, this function is chainable:

  Example

  \code

AppScript {

    script: {

        once(ActionTypes.askToRemoveSelectedItem, funciton() {

            removeConfirmationDialog.open();

        }).then(removeConfirmationDialog.onAccepted, function() {

            AppActions.removeSelectedItem();

        }); // <-- You may chain then() function again.
    }
}

  \endcode

  The callback in then() will not be registrated immediately.
  It is deferred until the previouew callback triggered.

 */

QFAppScriptRunnable *QFAppScript::once(QJSValue condition, QJSValue script)
{
    QFAppScriptRunnable* runnable = new QFAppScriptRunnable(this);
    runnable->setEngine(qmlEngine(this));
    runnable->setCondition(condition);
    runnable->setScript(script);
    m_runnables.append(runnable);
    return runnable;
}

/*! \qmlmethod AppScript::on(var type, func callback)

  Register a callback to be triggered when a matched message type is dispatched or a signal is emitted.
  User should call this function within the script code block.

  The callback will be removed on script termination.
 */

void QFAppScript::on(QJSValue condition, QJSValue script)
{
    QFAppScriptRunnable* runnable = once(condition,script);
    runnable->setIsOnceOnly(false);
}

void QFAppScript::onDispatched(QString type, QJSValue message)
{
    if (!m_runWhen.isEmpty() &&
        type == m_runWhen &&
        !m_processing) {

        if (m_running) {
            abort();
        }
        run(message);
        return;
    }

    if (!m_running) {
        return;
    }

    m_processing = true;

    // Mark for removeal
    QList<int> marked;

    for (int i = 0 ; i < m_runnables.size() ; i++) {
        if (m_runnables[i]->type() == type) {
            m_runnables[i]->run(message);

            if (!m_running) {
                // If exit() is called in runnable. It shoud not process any more.
                break;
            }

            if (m_runnables[i]->isOnceOnly()) {
                marked << i;
            }
        }
    }

    if (!m_running) {
        // Terminate if exit() is called in runnable
        m_processing = false;
        return;
    }

    for (int i = marked.size() - 1 ; i >= 0 ; i--) {
        int idx = marked[i];
        QFAppScriptRunnable* runnable = m_runnables.takeAt(idx);

        QFAppScriptRunnable* next = runnable->next();
        if (next) {
            next->setParent(this);
            m_runnables.append(next);
        }
        runnable->release();
        runnable->deleteLater();
    }

    m_processing = false;

    // All the tasks are finished
    if (m_runnables.size() == 0 && m_autoExit) {
        exit(0);
    }
}

void QFAppScript::componentComplete()
{
    QQuickItem::componentComplete();

    QQmlEngine *engine = qmlEngine(this);
    Q_ASSERT(engine);


    m_dispatcher = QFAppDispatcher::instance(engine);

    m_listener = new QFListener(this);

    setListenerId(m_dispatcher->addListener(m_listener));

    setListenerWaitFor();

    connect(m_listener,SIGNAL(dispatched(QString,QJSValue)),
            this,SLOT(onDispatched(QString,QJSValue)));
}

void QFAppScript::abort()
{
    exit(-1);
}

void QFAppScript::clear()
{
    for (int i = 0 ; i < m_runnables.size(); i++) {
        m_runnables[i]->deleteLater();
    }
    m_runnables.clear();
}

/*! \qmlproperty bool AppScript::running

  This property hold a value to indicate is the script still running.
  If there has any registered callback leave, it will be considered as running.

 */

bool QFAppScript::running() const
{
    return m_running;
}

void QFAppScript::setRunning(bool running)
{
    if (m_running == running) {
        return;
    }
    m_running = running;
    emit runningChanged();
}

QList<int> QFAppScript::waitFor() const
{
    return m_waitFor;
}

void QFAppScript::setWaitFor(const QList<int> &waitFor)
{
    m_waitFor = waitFor;
    setListenerWaitFor();
    emit waitForChanged();
}

void QFAppScript::setListenerWaitFor()
{
    if (!m_listener) {
        return;
    }

    m_listener->setWaitFor(m_waitFor);
}

bool QFAppScript::autoExit() const
{
    return m_autoExit;
}

void QFAppScript::setAutoExit(bool autoExit)
{
    m_autoExit = autoExit;
    emit autoExitChanged();
}

int QFAppScript::listenerId() const
{
    return m_listenerId;
}

void QFAppScript::setListenerId(int listenerId)
{
    m_listenerId = listenerId;
    emit listenerIdChanged();
}

QJSValue QFAppScript::message() const
{
    return m_message;
}

void QFAppScript::setMessage(const QJSValue &message)
{
    m_message = message;
    emit messageChanged();
}

/*! \qmlproperty string AppScript::runWhen

   This property hold a string of message type.
   Whatever a dispatched message matched, it will trigger to call run() immediately.

 */

QString QFAppScript::runWhen() const
{
    return m_runWhen;
}

void QFAppScript::setRunWhen(const QString &runWhen)
{
    m_runWhen = runWhen;
    emit runWhenChanged();
}

/*! \qmlproperty script AppScript::script
  This property holds the script to run.
 */

QQmlScriptString QFAppScript::script() const
{
    return m_script;
}

void QFAppScript::setScript(const QQmlScriptString &script)
{
    m_script = script;
    emit scriptChanged();
}
