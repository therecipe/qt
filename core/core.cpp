#define protected public

#include "core.h"
#include "_cgo_export.h"

#include <QAbstractAnimation>
#include <QAbstractEventDispatcher>
#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QAbstractNativeEventFilter>
#include <QAbstractProxyModel>
#include <QAbstractState>
#include <QAbstractTableModel>
#include <QAbstractTransition>
#include <QAssociativeIterable>
#include <QBasicTimer>
#include <QBitArray>
#include <QBuffer>
#include <QByteArray>
#include <QByteArrayList>
#include <QByteArrayMatcher>
#include <QChar>
#include <QChildEvent>
#include <QCollator>
#include <QCollatorSortKey>
#include <QCommandLineOption>
#include <QCommandLineParser>
#include <QCoreApplication>
#include <QCryptographicHash>
#include <QDataStream>
#include <QDate>
#include <QDateTime>
#include <QDebug>
#include <QDebugStateSaver>
#include <QDir>
#include <QDynamicPropertyChangeEvent>
#include <QEasingCurve>
#include <QElapsedTimer>
#include <QEvent>
#include <QEventLoop>
#include <QEventLoopLocker>
#include <QEventTransition>
#include <QFile>
#include <QFileDevice>
#include <QFileInfo>
#include <QFileSelector>
#include <QFileSystemWatcher>
#include <QFinalState>
#include <QFlag>
#include <QGenericArgument>
#include <QGenericReturnArgument>
#include <QHistoryState>
#include <QIODevice>
#include <QIdentityProxyModel>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QItemSelectionRange>
#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonParseError>
#include <QJsonValue>
#include <QLatin1Char>
#include <QLatin1String>
#include <QLibrary>
#include <QLibraryInfo>
#include <QLine>
#include <QLineF>
#include <QList>
#include <QLocale>
#include <QLockFile>
#include <QLoggingCategory>
#include <QMargins>
#include <QMarginsF>
#include <QMessageAuthenticationCode>
#include <QMessageLogger>
#include <QMetaEnum>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMetaProperty>
#include <QMetaType>
#include <QMimeData>
#include <QMimeDatabase>
#include <QMimeType>
#include <QModelIndex>
#include <QMutex>
#include <QMutexLocker>
#include <QObject>
#include <QObjectCleanupHandler>
#include <QParallelAnimationGroup>
#include <QPauseAnimation>
#include <QPersistentModelIndex>
#include <QPluginLoader>
#include <QPoint>
#include <QPointF>
#include <QProcess>
#include <QProcessEnvironment>
#include <QPropertyAnimation>
#include <QReadLocker>
#include <QReadWriteLock>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include <QRegularExpressionMatchIterator>
#include <QResource>
#include <QRunnable>
#include <QSaveFile>
#include <QSemaphore>
#include <QSequentialAnimationGroup>
#include <QSequentialIterable>
#include <QSet>
#include <QSettings>
#include <QSharedData>
#include <QSharedMemory>
#include <QSignalBlocker>
#include <QSignalMapper>
#include <QSignalTransition>
#include <QSize>
#include <QSizeF>
#include <QSocketNotifier>
#include <QSortFilterProxyModel>
#include <QStandardPaths>
#include <QState>
#include <QStateMachine>
#include <QStaticPlugin>
#include <QStorageInfo>
#include <QString>
#include <QStringList>
#include <QStringListModel>
#include <QStringMatcher>
#include <QStringRef>
#include <QSysInfo>
#include <QSystemSemaphore>
#include <QTemporaryDir>
#include <QTemporaryFile>
#include <QTextBoundaryFinder>
#include <QTextCodec>
#include <QTextDecoder>
#include <QTextEncoder>
#include <QTextStream>
#include <QThread>
#include <QThreadPool>
#include <QTime>
#include <QTimeLine>
#include <QTimeZone>
#include <QTimer>
#include <QTimerEvent>
#include <QTranslator>
#include <QUrl>
#include <QUrlQuery>
#include <QUuid>
#include <QVariant>
#include <QVariantAnimation>
#include <QWaitCondition>
#include <QWidget>
#include <QWriteLocker>
#include <QXmlStreamAttribute>
#include <QXmlStreamAttributes>
#include <QXmlStreamEntityDeclaration>
#include <QXmlStreamEntityResolver>
#include <QXmlStreamNamespaceDeclaration>
#include <QXmlStreamNotationDeclaration>
#include <QXmlStreamReader>
#include <QXmlStreamWriter>

class MyQAbstractAnimation: public QAbstractAnimation {
public:
	void Signal_CurrentLoopChanged(int currentLoop) { callbackQAbstractAnimationCurrentLoopChanged(this, this->objectName().toUtf8().data(), currentLoop); };
	void Signal_DirectionChanged(QAbstractAnimation::Direction newDirection) { callbackQAbstractAnimationDirectionChanged(this, this->objectName().toUtf8().data(), newDirection); };
	void Signal_Finished() { callbackQAbstractAnimationFinished(this, this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQAbstractAnimationStateChanged(this, this->objectName().toUtf8().data(), newState, oldState); };
	void updateDirection(QAbstractAnimation::Direction direction) { callbackQAbstractAnimationUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQAbstractAnimationUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAnimationTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractAnimationChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractAnimationCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QAbstractAnimation_CurrentLoop(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoop();
}

int QAbstractAnimation_CurrentTime(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentTime();
}

int QAbstractAnimation_Direction(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->direction();
}

int QAbstractAnimation_LoopCount(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->loopCount();
}

void QAbstractAnimation_SetCurrentTime(void* ptr, int msecs){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setCurrentTime", Q_ARG(int, msecs));
}

void QAbstractAnimation_SetDirection(void* ptr, int direction){
	static_cast<QAbstractAnimation*>(ptr)->setDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QAbstractAnimation_SetLoopCount(void* ptr, int loopCount){
	static_cast<QAbstractAnimation*>(ptr)->setLoopCount(loopCount);
}

int QAbstractAnimation_State(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->state();
}

void QAbstractAnimation_ConnectCurrentLoopChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

void QAbstractAnimation_DisconnectCurrentLoopChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

void QAbstractAnimation_CurrentLoopChanged(void* ptr, int currentLoop){
	static_cast<QAbstractAnimation*>(ptr)->currentLoopChanged(currentLoop);
}

int QAbstractAnimation_CurrentLoopTime(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoopTime();
}

void QAbstractAnimation_ConnectDirectionChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

void QAbstractAnimation_DisconnectDirectionChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

void QAbstractAnimation_DirectionChanged(void* ptr, int newDirection){
	static_cast<QAbstractAnimation*>(ptr)->directionChanged(static_cast<QAbstractAnimation::Direction>(newDirection));
}

int QAbstractAnimation_Duration(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->duration();
}

int QAbstractAnimation_Event(void* ptr, void* event){
	return static_cast<QAbstractAnimation*>(ptr)->event(static_cast<QEvent*>(event));
}

void QAbstractAnimation_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

void QAbstractAnimation_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

void QAbstractAnimation_Finished(void* ptr){
	static_cast<QAbstractAnimation*>(ptr)->finished();
}

void* QAbstractAnimation_Group(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->group();
}

void QAbstractAnimation_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "pause");
}

void QAbstractAnimation_Resume(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "resume");
}

void QAbstractAnimation_SetPaused(void* ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QAbstractAnimation_Start(void* ptr, int policy){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "start", Q_ARG(QAbstractAnimation::DeletionPolicy, static_cast<QAbstractAnimation::DeletionPolicy>(policy)));
}

void QAbstractAnimation_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_StateChanged(void* ptr, int newState, int oldState){
	static_cast<QAbstractAnimation*>(ptr)->stateChanged(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QAbstractAnimation_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "stop");
}

int QAbstractAnimation_TotalDuration(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->totalDuration();
}

void QAbstractAnimation_UpdateDirection(void* ptr, int direction){
	static_cast<MyQAbstractAnimation*>(ptr)->updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QAbstractAnimation_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QAbstractAnimation*>(ptr)->QAbstractAnimation::updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QAbstractAnimation_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQAbstractAnimation*>(ptr)->updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QAbstractAnimation_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QAbstractAnimation*>(ptr)->QAbstractAnimation::updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QAbstractAnimation_DestroyQAbstractAnimation(void* ptr){
	static_cast<QAbstractAnimation*>(ptr)->~QAbstractAnimation();
}

void QAbstractAnimation_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractAnimation*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractAnimation_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractAnimation*>(ptr)->QAbstractAnimation::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractAnimation_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractAnimation*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractAnimation_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractAnimation*>(ptr)->QAbstractAnimation::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractAnimation_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractAnimation*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractAnimation_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractAnimation*>(ptr)->QAbstractAnimation::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractEventDispatcher: public QAbstractEventDispatcher {
public:
	void Signal_AboutToBlock() { callbackQAbstractEventDispatcherAboutToBlock(this, this->objectName().toUtf8().data()); };
	void Signal_Awake() { callbackQAbstractEventDispatcherAwake(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractEventDispatcherTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractEventDispatcherChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractEventDispatcherCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QAbstractEventDispatcher_ConnectAboutToBlock(void* ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_DisconnectAboutToBlock(void* ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_AboutToBlock(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->aboutToBlock();
}

void QAbstractEventDispatcher_ConnectAwake(void* ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_DisconnectAwake(void* ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_Awake(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->awake();
}

void QAbstractEventDispatcher_Flush(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->flush();
}

void QAbstractEventDispatcher_InstallNativeEventFilter(void* ptr, void* filterObj){
	static_cast<QAbstractEventDispatcher*>(ptr)->installNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObj));
}

void* QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(void* thread){
	return QAbstractEventDispatcher::instance(static_cast<QThread*>(thread));
}

void QAbstractEventDispatcher_Interrupt(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->interrupt();
}

int QAbstractEventDispatcher_ProcessEvents(void* ptr, int flags){
	return static_cast<QAbstractEventDispatcher*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QAbstractEventDispatcher_RegisterSocketNotifier(void* ptr, void* notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->registerSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_RemainingTime(void* ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->remainingTime(timerId);
}

void QAbstractEventDispatcher_RemoveNativeEventFilter(void* ptr, void* filter){
	static_cast<QAbstractEventDispatcher*>(ptr)->removeNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filter));
}

void QAbstractEventDispatcher_UnregisterSocketNotifier(void* ptr, void* notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->unregisterSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_UnregisterTimer(void* ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimer(timerId);
}

int QAbstractEventDispatcher_UnregisterTimers(void* ptr, void* object){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimers(static_cast<QObject*>(object));
}

void QAbstractEventDispatcher_WakeUp(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->wakeUp();
}

void QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->~QAbstractEventDispatcher();
}

void QAbstractEventDispatcher_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractEventDispatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractEventDispatcher_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractEventDispatcher*>(ptr)->QAbstractEventDispatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractEventDispatcher_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractEventDispatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractEventDispatcher_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractEventDispatcher*>(ptr)->QAbstractEventDispatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractEventDispatcher_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractEventDispatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractEventDispatcher_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractEventDispatcher*>(ptr)->QAbstractEventDispatcher::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractItemModel: public QAbstractItemModel {
public:
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelColumnsAboutToBeInserted(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModelColumnsAboutToBeMoved(this, this->objectName().toUtf8().data(), new QModelIndex(sourceParent), sourceStart, sourceEnd, new QModelIndex(destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelColumnsAboutToBeRemoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelColumnsInserted(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModelColumnsMoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), start, end, new QModelIndex(destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelColumnsRemoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModelHeaderDataChanged(this, this->objectName().toUtf8().data(), orientation, first, last); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModelModelAboutToBeReset(this, this->objectName().toUtf8().data()); };
	void Signal_ModelReset() { callbackQAbstractItemModelModelReset(this, this->objectName().toUtf8().data()); };
	void revert() { if (!callbackQAbstractItemModelRevert(this, this->objectName().toUtf8().data())) { QAbstractItemModel::revert(); }; };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModelRowsAboutToBeInserted(this, this->objectName().toUtf8().data(), new QModelIndex(parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModelRowsAboutToBeMoved(this, this->objectName().toUtf8().data(), new QModelIndex(sourceParent), sourceStart, sourceEnd, new QModelIndex(destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelRowsAboutToBeRemoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelRowsInserted(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModelRowsMoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), start, end, new QModelIndex(destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelRowsRemoved(this, this->objectName().toUtf8().data(), new QModelIndex(parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractItemModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractItemModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractItemModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractItemModel_Sibling(void* ptr, int row, int column, void* index){
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(index)));
}

void* QAbstractItemModel_Buddy(void* ptr, void* index){
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

int QAbstractItemModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectColumnsAboutToBeInserted(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeInserted(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted));;
}

void QAbstractItemModel_ConnectColumnsAboutToBeMoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeMoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved));;
}

void QAbstractItemModel_ConnectColumnsAboutToBeRemoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved));;
}

void QAbstractItemModel_ConnectColumnsInserted(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted));;
}

void QAbstractItemModel_DisconnectColumnsInserted(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted));;
}

void QAbstractItemModel_ConnectColumnsMoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved));;
}

void QAbstractItemModel_DisconnectColumnsMoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved));;
}

void QAbstractItemModel_ConnectColumnsRemoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved));;
}

void QAbstractItemModel_DisconnectColumnsRemoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved));;
}

void* QAbstractItemModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QAbstractItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QAbstractItemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQAbstractItemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QAbstractItemModel_HasChildren(void* ptr, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_HasIndex(void* ptr, int row, int column, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->hasIndex(row, column, *static_cast<QModelIndex*>(parent));
}

void* QAbstractItemModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QAbstractItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void QAbstractItemModel_ConnectHeaderDataChanged(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged));;
}

void QAbstractItemModel_DisconnectHeaderDataChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged));;
}

void QAbstractItemModel_HeaderDataChanged(void* ptr, int orientation, int first, int last){
	static_cast<QAbstractItemModel*>(ptr)->headerDataChanged(static_cast<Qt::Orientation>(orientation), first, last);
}

void* QAbstractItemModel_Index(void* ptr, int row, int column, void* parent){
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QAbstractItemModel_InsertColumn(void* ptr, int column, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertRow(void* ptr, int row, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char* QAbstractItemModel_MimeTypes(void* ptr){
	return static_cast<QAbstractItemModel*>(ptr)->mimeTypes().join(",,,").toUtf8().data();
}

void QAbstractItemModel_ConnectModelAboutToBeReset(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset));;
}

void QAbstractItemModel_DisconnectModelAboutToBeReset(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset));;
}

void QAbstractItemModel_ConnectModelReset(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset));;
}

void QAbstractItemModel_DisconnectModelReset(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset));;
}

int QAbstractItemModel_MoveColumn(void* ptr, void* sourceParent, int sourceColumn, void* destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveColumn(*static_cast<QModelIndex*>(sourceParent), sourceColumn, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveRow(void* ptr, void* sourceParent, int sourceRow, void* destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveRow(*static_cast<QModelIndex*>(sourceParent), sourceRow, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QAbstractItemModel_Parent(void* ptr, void* index){
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

int QAbstractItemModel_RemoveColumn(void* ptr, int column, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeColumn(column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveRow(void* ptr, int row, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeRow(row, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQAbstractItemModel*>(ptr), "revert");
}

void QAbstractItemModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "revert");
}

int QAbstractItemModel_RowCount(void* ptr, void* parent){
	return static_cast<QAbstractItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectRowsAboutToBeInserted(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeInserted(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted));;
}

void QAbstractItemModel_ConnectRowsAboutToBeMoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeMoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved));;
}

void QAbstractItemModel_ConnectRowsAboutToBeRemoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeRemoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved));;
}

void QAbstractItemModel_ConnectRowsInserted(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted));;
}

void QAbstractItemModel_DisconnectRowsInserted(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted));;
}

void QAbstractItemModel_ConnectRowsMoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved));;
}

void QAbstractItemModel_DisconnectRowsMoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved));;
}

void QAbstractItemModel_ConnectRowsRemoved(void* ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved));;
}

void QAbstractItemModel_DisconnectRowsRemoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved));;
}

int QAbstractItemModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QAbstractItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QAbstractItemModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QAbstractItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QAbstractItemModel_Sort(void* ptr, int column, int order){
	static_cast<MyQAbstractItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractItemModel_SortDefault(void* ptr, int column, int order){
	static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QAbstractItemModel_Span(void* ptr, void* index){
	return new QSize(static_cast<QSize>(static_cast<QAbstractItemModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).width(), static_cast<QSize>(static_cast<QAbstractItemModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).height());
}

int QAbstractItemModel_Submit(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "submit");
}

int QAbstractItemModel_SupportedDragActions(void* ptr){
	return static_cast<QAbstractItemModel*>(ptr)->supportedDragActions();
}

int QAbstractItemModel_SupportedDropActions(void* ptr){
	return static_cast<QAbstractItemModel*>(ptr)->supportedDropActions();
}

void QAbstractItemModel_DestroyQAbstractItemModel(void* ptr){
	static_cast<QAbstractItemModel*>(ptr)->~QAbstractItemModel();
}

void QAbstractItemModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractItemModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractItemModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractItemModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractItemModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractItemModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractItemModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractItemModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractItemModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractListModel: public QAbstractListModel {
public:
	void fetchMore(const QModelIndex & parent) { callbackQAbstractListModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void revert() { if (!callbackQAbstractListModelRevert(this, this->objectName().toUtf8().data())) { QAbstractListModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractListModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractListModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractListModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractListModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractListModel_Index(void* ptr, int row, int column, void* parent){
	return new QModelIndex(static_cast<QAbstractListModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QAbstractListModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractListModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractListModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QAbstractListModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QAbstractListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QAbstractListModel_DestroyQAbstractListModel(void* ptr){
	static_cast<QAbstractListModel*>(ptr)->~QAbstractListModel();
}

void QAbstractListModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQAbstractListModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractListModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QAbstractListModel*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractListModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQAbstractListModel*>(ptr), "revert");
}

void QAbstractListModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractListModel*>(ptr), "revert");
}

void QAbstractListModel_Sort(void* ptr, int column, int order){
	static_cast<MyQAbstractListModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractListModel_SortDefault(void* ptr, int column, int order){
	static_cast<QAbstractListModel*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractListModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractListModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractListModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractListModel*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractListModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractListModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractListModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractListModel*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractListModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractListModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractListModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractListModel*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractNativeEventFilter: public QAbstractNativeEventFilter {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QAbstractNativeEventFilter_DestroyQAbstractNativeEventFilter(void* ptr){
	static_cast<QAbstractNativeEventFilter*>(ptr)->~QAbstractNativeEventFilter();
}

char* QAbstractNativeEventFilter_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQAbstractNativeEventFilter*>(static_cast<QAbstractNativeEventFilter*>(ptr))) {
		return static_cast<MyQAbstractNativeEventFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractNativeEventFilter_BASE").toUtf8().data();
}

void QAbstractNativeEventFilter_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQAbstractNativeEventFilter*>(static_cast<QAbstractNativeEventFilter*>(ptr))) {
		static_cast<MyQAbstractNativeEventFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAbstractProxyModel: public QAbstractProxyModel {
public:
	void fetchMore(const QModelIndex & parent) { callbackQAbstractProxyModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void revert() { if (!callbackQAbstractProxyModelRevert(this, this->objectName().toUtf8().data())) { QAbstractProxyModel::revert(); }; };
	void setSourceModel(QAbstractItemModel * sourceModel) { callbackQAbstractProxyModelSetSourceModel(this, this->objectName().toUtf8().data(), sourceModel); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractProxyModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void Signal_SourceModelChanged() { callbackQAbstractProxyModelSourceModelChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractProxyModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractProxyModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractProxyModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractProxyModel_Buddy(void* ptr, void* index){
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

int QAbstractProxyModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

void* QAbstractProxyModel_Data(void* ptr, void* proxyIndex, int role){
	return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(proxyIndex), role));
}

int QAbstractProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractProxyModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQAbstractProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractProxyModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QAbstractProxyModel_HasChildren(void* ptr, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QAbstractProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QAbstractProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QAbstractProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

char* QAbstractProxyModel_MimeTypes(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->mimeTypes().join(",,,").toUtf8().data();
}

void QAbstractProxyModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQAbstractProxyModel*>(ptr), "revert");
}

void QAbstractProxyModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractProxyModel*>(ptr), "revert");
}

int QAbstractProxyModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QAbstractProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QAbstractProxyModel_SetSourceModel(void* ptr, void* sourceModel){
	static_cast<MyQAbstractProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void QAbstractProxyModel_SetSourceModelDefault(void* ptr, void* sourceModel){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void* QAbstractProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QAbstractProxyModel_Sort(void* ptr, int column, int order){
	static_cast<MyQAbstractProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractProxyModel_SortDefault(void* ptr, int column, int order){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QAbstractProxyModel_SourceModel(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->sourceModel();
}

void QAbstractProxyModel_ConnectSourceModelChanged(void* ptr){
	QObject::connect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

void QAbstractProxyModel_DisconnectSourceModelChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

void* QAbstractProxyModel_Span(void* ptr, void* index){
	return new QSize(static_cast<QSize>(static_cast<QAbstractProxyModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).width(), static_cast<QSize>(static_cast<QAbstractProxyModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).height());
}

int QAbstractProxyModel_Submit(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->submit();
}

int QAbstractProxyModel_SupportedDragActions(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDragActions();
}

int QAbstractProxyModel_SupportedDropActions(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDropActions();
}

void QAbstractProxyModel_DestroyQAbstractProxyModel(void* ptr){
	static_cast<QAbstractProxyModel*>(ptr)->~QAbstractProxyModel();
}

void QAbstractProxyModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractProxyModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractProxyModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractProxyModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractProxyModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractProxyModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractProxyModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractProxyModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractProxyModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractState: public QAbstractState {
public:
	void Signal_ActiveChanged(bool active) { callbackQAbstractStateActiveChanged(this, this->objectName().toUtf8().data(), active); };
	void Signal_Entered() { callbackQAbstractStateEntered(this, this->objectName().toUtf8().data()); };
	void Signal_Exited() { callbackQAbstractStateExited(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractStateTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractStateChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractStateCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QAbstractState_Active(void* ptr){
	return static_cast<QAbstractState*>(ptr)->active();
}

void QAbstractState_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), static_cast<void (QAbstractState::*)(bool)>(&QAbstractState::activeChanged), static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)(bool)>(&MyQAbstractState::Signal_ActiveChanged));;
}

void QAbstractState_ActiveChanged(void* ptr, int active){
	static_cast<QAbstractState*>(ptr)->activeChanged(active != 0);
}

void QAbstractState_ConnectEntered(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

void QAbstractState_DisconnectEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::entered, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Entered));;
}

int QAbstractState_Event(void* ptr, void* e){
	return static_cast<QAbstractState*>(ptr)->event(static_cast<QEvent*>(e));
}

void QAbstractState_ConnectExited(void* ptr){
	QObject::connect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

void QAbstractState_DisconnectExited(void* ptr){
	QObject::disconnect(static_cast<QAbstractState*>(ptr), &QAbstractState::exited, static_cast<MyQAbstractState*>(ptr), static_cast<void (MyQAbstractState::*)()>(&MyQAbstractState::Signal_Exited));;
}

void* QAbstractState_Machine(void* ptr){
	return static_cast<QAbstractState*>(ptr)->machine();
}

void* QAbstractState_ParentState(void* ptr){
	return static_cast<QAbstractState*>(ptr)->parentState();
}

void QAbstractState_DestroyQAbstractState(void* ptr){
	static_cast<QAbstractState*>(ptr)->~QAbstractState();
}

void QAbstractState_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractState*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractState_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractState*>(ptr)->QAbstractState::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractState_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractState*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractState_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractState*>(ptr)->QAbstractState::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractState_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractState*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractState_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractState*>(ptr)->QAbstractState::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractTableModel: public QAbstractTableModel {
public:
	void fetchMore(const QModelIndex & parent) { callbackQAbstractTableModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void revert() { if (!callbackQAbstractTableModelRevert(this, this->objectName().toUtf8().data())) { QAbstractTableModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractTableModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractTableModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractTableModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractTableModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractTableModel_Index(void* ptr, int row, int column, void* parent){
	return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QAbstractTableModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractTableModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractTableModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QAbstractTableModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QAbstractTableModel_DestroyQAbstractTableModel(void* ptr){
	static_cast<QAbstractTableModel*>(ptr)->~QAbstractTableModel();
}

void QAbstractTableModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQAbstractTableModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractTableModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractTableModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQAbstractTableModel*>(ptr), "revert");
}

void QAbstractTableModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractTableModel*>(ptr), "revert");
}

void QAbstractTableModel_Sort(void* ptr, int column, int order){
	static_cast<MyQAbstractTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractTableModel_SortDefault(void* ptr, int column, int order){
	static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractTableModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractTableModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTableModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTableModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractTableModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTableModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTableModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractTableModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractTableModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractTransition: public QAbstractTransition {
public:
	void Signal_TargetStateChanged() { callbackQAbstractTransitionTargetStateChanged(this, this->objectName().toUtf8().data()); };
	void Signal_TargetStatesChanged() { callbackQAbstractTransitionTargetStatesChanged(this, this->objectName().toUtf8().data()); };
	void Signal_Triggered() { callbackQAbstractTransitionTriggered(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractTransitionTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractTransitionChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractTransitionCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QAbstractTransition_AddAnimation(void* ptr, void* animation){
	static_cast<QAbstractTransition*>(ptr)->addAnimation(static_cast<QAbstractAnimation*>(animation));
}

int QAbstractTransition_Event(void* ptr, void* e){
	return static_cast<QAbstractTransition*>(ptr)->event(static_cast<QEvent*>(e));
}

void* QAbstractTransition_Machine(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->machine();
}

void QAbstractTransition_RemoveAnimation(void* ptr, void* animation){
	static_cast<QAbstractTransition*>(ptr)->removeAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QAbstractTransition_SetTargetState(void* ptr, void* target){
	static_cast<QAbstractTransition*>(ptr)->setTargetState(static_cast<QAbstractState*>(target));
}

void QAbstractTransition_SetTransitionType(void* ptr, int ty){
	static_cast<QAbstractTransition*>(ptr)->setTransitionType(static_cast<QAbstractTransition::TransitionType>(ty));
}

void* QAbstractTransition_SourceState(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->sourceState();
}

void* QAbstractTransition_TargetState(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->targetState();
}

void QAbstractTransition_ConnectTargetStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_DisconnectTargetStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStateChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStateChanged));;
}

void QAbstractTransition_ConnectTargetStatesChanged(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

void QAbstractTransition_DisconnectTargetStatesChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::targetStatesChanged, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_TargetStatesChanged));;
}

int QAbstractTransition_TransitionType(void* ptr){
	return static_cast<QAbstractTransition*>(ptr)->transitionType();
}

void QAbstractTransition_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QAbstractTransition*>(ptr), &QAbstractTransition::triggered, static_cast<MyQAbstractTransition*>(ptr), static_cast<void (MyQAbstractTransition::*)()>(&MyQAbstractTransition::Signal_Triggered));;
}

void QAbstractTransition_DestroyQAbstractTransition(void* ptr){
	static_cast<QAbstractTransition*>(ptr)->~QAbstractTransition();
}

void QAbstractTransition_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractTransition*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTransition_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractTransition*>(ptr)->QAbstractTransition::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTransition_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractTransition*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTransition_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractTransition*>(ptr)->QAbstractTransition::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTransition_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractTransition*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractTransition_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractTransition*>(ptr)->QAbstractTransition::customEvent(static_cast<QEvent*>(event));
}

int QAssociativeIterable_Size(void* ptr){
	return static_cast<QAssociativeIterable*>(ptr)->size();
}

void* QAssociativeIterable_Value(void* ptr, void* key){
	return new QVariant(static_cast<QAssociativeIterable*>(ptr)->value(*static_cast<QVariant*>(key)));
}

void QBasicTimer_Start(void* ptr, int msec, void* object){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<QObject*>(object));
}

void* QBasicTimer_NewQBasicTimer(){
	return new QBasicTimer();
}

int QBasicTimer_IsActive(void* ptr){
	return static_cast<QBasicTimer*>(ptr)->isActive();
}

void QBasicTimer_Start2(void* ptr, int msec, int timerType, void* obj){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(obj));
}

void QBasicTimer_Stop(void* ptr){
	static_cast<QBasicTimer*>(ptr)->stop();
}

int QBasicTimer_TimerId(void* ptr){
	return static_cast<QBasicTimer*>(ptr)->timerId();
}

void QBasicTimer_DestroyQBasicTimer(void* ptr){
	static_cast<QBasicTimer*>(ptr)->~QBasicTimer();
}

void* QBitArray_NewQBitArray(){
	return new QBitArray();
}

void* QBitArray_NewQBitArray4(void* other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray3(void* other){
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray2(int size, int value){
	return new QBitArray(size, value != 0);
}

int QBitArray_At(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->at(i);
}

void QBitArray_Clear(void* ptr){
	static_cast<QBitArray*>(ptr)->clear();
}

void QBitArray_ClearBit(void* ptr, int i){
	static_cast<QBitArray*>(ptr)->clearBit(i);
}

int QBitArray_Count(void* ptr){
	return static_cast<QBitArray*>(ptr)->count();
}

int QBitArray_Count2(void* ptr, int on){
	return static_cast<QBitArray*>(ptr)->count(on != 0);
}

int QBitArray_Fill(void* ptr, int value, int size){
	return static_cast<QBitArray*>(ptr)->fill(value != 0, size);
}

void QBitArray_Fill2(void* ptr, int value, int begin, int end){
	static_cast<QBitArray*>(ptr)->fill(value != 0, begin, end);
}

int QBitArray_IsEmpty(void* ptr){
	return static_cast<QBitArray*>(ptr)->isEmpty();
}

int QBitArray_IsNull(void* ptr){
	return static_cast<QBitArray*>(ptr)->isNull();
}

void QBitArray_Resize(void* ptr, int size){
	static_cast<QBitArray*>(ptr)->resize(size);
}

void QBitArray_SetBit(void* ptr, int i){
	static_cast<QBitArray*>(ptr)->setBit(i);
}

void QBitArray_SetBit2(void* ptr, int i, int value){
	static_cast<QBitArray*>(ptr)->setBit(i, value != 0);
}

int QBitArray_Size(void* ptr){
	return static_cast<QBitArray*>(ptr)->size();
}

void QBitArray_Swap(void* ptr, void* other){
	static_cast<QBitArray*>(ptr)->swap(*static_cast<QBitArray*>(other));
}

int QBitArray_TestBit(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->testBit(i);
}

int QBitArray_ToggleBit(void* ptr, int i){
	return static_cast<QBitArray*>(ptr)->toggleBit(i);
}

void QBitArray_Truncate(void* ptr, int pos){
	static_cast<QBitArray*>(ptr)->truncate(pos);
}

class MyQBuffer: public QBuffer {
public:
	MyQBuffer(QByteArray *byteArray, QObject *parent) : QBuffer(byteArray, parent) {};
	MyQBuffer(QObject *parent) : QBuffer(parent) {};
	void close() { callbackQBufferClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQBufferTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQBufferChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQBufferCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QBuffer_NewQBuffer2(void* byteArray, void* parent){
	return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QObject*>(parent));
}

void* QBuffer_NewQBuffer(void* parent){
	return new MyQBuffer(static_cast<QObject*>(parent));
}

int QBuffer_AtEnd(void* ptr){
	return static_cast<QBuffer*>(ptr)->atEnd();
}

void* QBuffer_Buffer(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->buffer());
}

void* QBuffer_Buffer2(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->buffer());
}

int QBuffer_CanReadLine(void* ptr){
	return static_cast<QBuffer*>(ptr)->canReadLine();
}

void QBuffer_Close(void* ptr){
	static_cast<MyQBuffer*>(ptr)->close();
}

void QBuffer_CloseDefault(void* ptr){
	static_cast<QBuffer*>(ptr)->QBuffer::close();
}

void* QBuffer_Data(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->data());
}

int QBuffer_Open(void* ptr, int flags){
	return static_cast<QBuffer*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(flags));
}

long long QBuffer_Pos(void* ptr){
	return static_cast<long long>(static_cast<QBuffer*>(ptr)->pos());
}

long long QBuffer_ReadData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QBuffer*>(ptr)->readData(data, static_cast<long long>(len)));
}

int QBuffer_Seek(void* ptr, long long pos){
	return static_cast<QBuffer*>(ptr)->seek(static_cast<long long>(pos));
}

void QBuffer_SetBuffer(void* ptr, void* byteArray){
	static_cast<QBuffer*>(ptr)->setBuffer(static_cast<QByteArray*>(byteArray));
}

void QBuffer_SetData(void* ptr, void* data){
	static_cast<QBuffer*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QBuffer_SetData2(void* ptr, char* data, int size){
	static_cast<QBuffer*>(ptr)->setData(const_cast<const char*>(data), size);
}

long long QBuffer_Size(void* ptr){
	return static_cast<long long>(static_cast<QBuffer*>(ptr)->size());
}

long long QBuffer_WriteData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QBuffer*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QBuffer_DestroyQBuffer(void* ptr){
	static_cast<QBuffer*>(ptr)->~QBuffer();
}

void QBuffer_TimerEvent(void* ptr, void* event){
	static_cast<MyQBuffer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBuffer_TimerEventDefault(void* ptr, void* event){
	static_cast<QBuffer*>(ptr)->QBuffer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBuffer_ChildEvent(void* ptr, void* event){
	static_cast<MyQBuffer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBuffer_ChildEventDefault(void* ptr, void* event){
	static_cast<QBuffer*>(ptr)->QBuffer::childEvent(static_cast<QChildEvent*>(event));
}

void QBuffer_CustomEvent(void* ptr, void* event){
	static_cast<MyQBuffer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBuffer_CustomEventDefault(void* ptr, void* event){
	static_cast<QBuffer*>(ptr)->QBuffer::customEvent(static_cast<QEvent*>(event));
}

void QByteArray_Clear(void* ptr){
	static_cast<QByteArray*>(ptr)->clear();
}

int QByteArray_IndexOf3(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(const_cast<const char*>(str), from);
}

int QByteArray_IsNull(void* ptr){
	return static_cast<QByteArray*>(ptr)->isNull();
}

int QByteArray_LastIndexOf(void* ptr, void* ba, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_LastIndexOf3(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(const_cast<const char*>(str), from);
}

void* QByteArray_NewQByteArray(){
	return new QByteArray();
}

void* QByteArray_NewQByteArray6(void* other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray5(void* other){
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray2(char* data, int size){
	return new QByteArray(const_cast<const char*>(data), size);
}

void* QByteArray_NewQByteArray3(int size, char* ch){
	return new QByteArray(size, *ch);
}

void* QByteArray_Append5(void* ptr, char* ch){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*ch));
}

void* QByteArray_Append(void* ptr, void* ba){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*static_cast<QByteArray*>(ba)));
}

void* QByteArray_Append2(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(QString(str)));
}

void* QByteArray_Append3(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str)));
}

void* QByteArray_Append4(void* ptr, char* str, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str), len));
}

int QByteArray_Capacity(void* ptr){
	return static_cast<QByteArray*>(ptr)->capacity();
}

void QByteArray_Chop(void* ptr, int n){
	static_cast<QByteArray*>(ptr)->chop(n);
}

int QByteArray_Contains3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->contains(*ch);
}

int QByteArray_Contains(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->contains(*static_cast<QByteArray*>(ba));
}

int QByteArray_Contains2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->contains(const_cast<const char*>(str));
}

int QByteArray_Count4(void* ptr){
	return static_cast<QByteArray*>(ptr)->count();
}

int QByteArray_Count3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->count(*ch);
}

int QByteArray_Count(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->count(*static_cast<QByteArray*>(ba));
}

int QByteArray_Count2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->count(const_cast<const char*>(str));
}

int QByteArray_EndsWith3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->endsWith(*ch);
}

int QByteArray_EndsWith(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->endsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_EndsWith2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->endsWith(const_cast<const char*>(str));
}

void* QByteArray_Fill(void* ptr, char* ch, int size){
	return new QByteArray(static_cast<QByteArray*>(ptr)->fill(*ch, size));
}

void* QByteArray_QByteArray_FromBase64(void* base64){
	return new QByteArray(QByteArray::fromBase64(*static_cast<QByteArray*>(base64)));
}

void* QByteArray_QByteArray_FromBase642(void* base64, int options){
	return new QByteArray(QByteArray::fromBase64(*static_cast<QByteArray*>(base64), static_cast<QByteArray::Base64Option>(options)));
}

void* QByteArray_QByteArray_FromHex(void* hexEncoded){
	return new QByteArray(QByteArray::fromHex(*static_cast<QByteArray*>(hexEncoded)));
}

void* QByteArray_QByteArray_FromPercentEncoding(void* input, char* percent){
	return new QByteArray(QByteArray::fromPercentEncoding(*static_cast<QByteArray*>(input), *percent));
}

void* QByteArray_QByteArray_FromRawData(char* data, int size){
	return new QByteArray(QByteArray::fromRawData(const_cast<const char*>(data), size));
}

int QByteArray_IndexOf4(void* ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*ch, from);
}

int QByteArray_IndexOf(void* ptr, void* ba, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_IndexOf2(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->indexOf(QString(str), from);
}

int QByteArray_IsEmpty(void* ptr){
	return static_cast<QByteArray*>(ptr)->isEmpty();
}

int QByteArray_LastIndexOf4(void* ptr, char* ch, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(*ch, from);
}

int QByteArray_LastIndexOf2(void* ptr, char* str, int from){
	return static_cast<QByteArray*>(ptr)->lastIndexOf(QString(str), from);
}

void* QByteArray_Left(void* ptr, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->left(len));
}

void* QByteArray_LeftJustified(void* ptr, int width, char* fill, int truncate){
	return new QByteArray(static_cast<QByteArray*>(ptr)->leftJustified(width, *fill, truncate != 0));
}

int QByteArray_Length(void* ptr){
	return static_cast<QByteArray*>(ptr)->length();
}

void* QByteArray_Mid(void* ptr, int pos, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->mid(pos, len));
}

void* QByteArray_QByteArray_Number(int n, int base){
	return new QByteArray(QByteArray::number(n, base));
}

void* QByteArray_Prepend4(void* ptr, char* ch){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(*ch));
}

void* QByteArray_Prepend(void* ptr, void* ba){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(*static_cast<QByteArray*>(ba)));
}

void* QByteArray_Prepend2(void* ptr, char* str){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(const_cast<const char*>(str)));
}

void* QByteArray_Prepend3(void* ptr, char* str, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->prepend(const_cast<const char*>(str), len));
}

void QByteArray_Push_back3(void* ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_back(*ch);
}

void QByteArray_Push_back(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->push_back(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_back2(void* ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_back(const_cast<const char*>(str));
}

void QByteArray_Push_front3(void* ptr, char* ch){
	static_cast<QByteArray*>(ptr)->push_front(*ch);
}

void QByteArray_Push_front(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->push_front(*static_cast<QByteArray*>(other));
}

void QByteArray_Push_front2(void* ptr, char* str){
	static_cast<QByteArray*>(ptr)->push_front(const_cast<const char*>(str));
}

void* QByteArray_Repeated(void* ptr, int times){
	return new QByteArray(static_cast<QByteArray*>(ptr)->repeated(times));
}

void QByteArray_Reserve(void* ptr, int size){
	static_cast<QByteArray*>(ptr)->reserve(size);
}

void QByteArray_Resize(void* ptr, int size){
	static_cast<QByteArray*>(ptr)->resize(size);
}

void* QByteArray_Right(void* ptr, int len){
	return new QByteArray(static_cast<QByteArray*>(ptr)->right(len));
}

void* QByteArray_RightJustified(void* ptr, int width, char* fill, int truncate){
	return new QByteArray(static_cast<QByteArray*>(ptr)->rightJustified(width, *fill, truncate != 0));
}

void* QByteArray_SetNum(void* ptr, int n, int base){
	return new QByteArray(static_cast<QByteArray*>(ptr)->setNum(n, base));
}

int QByteArray_Size(void* ptr){
	return static_cast<QByteArray*>(ptr)->size();
}

void QByteArray_Squeeze(void* ptr){
	static_cast<QByteArray*>(ptr)->squeeze();
}

int QByteArray_StartsWith3(void* ptr, char* ch){
	return static_cast<QByteArray*>(ptr)->startsWith(*ch);
}

int QByteArray_StartsWith(void* ptr, void* ba){
	return static_cast<QByteArray*>(ptr)->startsWith(*static_cast<QByteArray*>(ba));
}

int QByteArray_StartsWith2(void* ptr, char* str){
	return static_cast<QByteArray*>(ptr)->startsWith(const_cast<const char*>(str));
}

void QByteArray_Swap(void* ptr, void* other){
	static_cast<QByteArray*>(ptr)->swap(*static_cast<QByteArray*>(other));
}

void* QByteArray_ToBase64(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toBase64());
}

void* QByteArray_ToBase642(void* ptr, int options){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toBase64(static_cast<QByteArray::Base64Option>(options)));
}

void* QByteArray_ToHex(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toHex());
}

int QByteArray_ToInt(void* ptr, int ok, int base){
	return static_cast<QByteArray*>(ptr)->toInt(NULL, base);
}

void* QByteArray_ToPercentEncoding(void* ptr, void* exclude, void* include, char* percent){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toPercentEncoding(*static_cast<QByteArray*>(exclude), *static_cast<QByteArray*>(include), *percent));
}

void QByteArray_Truncate(void* ptr, int pos){
	static_cast<QByteArray*>(ptr)->truncate(pos);
}

void QByteArray_DestroyQByteArray(void* ptr){
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

void* QByteArray_Simplified(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->simplified());
}

void* QByteArray_ToLower(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toLower());
}

void* QByteArray_ToUpper(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->toUpper());
}

void* QByteArray_Trimmed(void* ptr){
	return new QByteArray(static_cast<QByteArray*>(ptr)->trimmed());
}

void* QByteArrayList_Join(void* ptr){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join());
}

void* QByteArrayList_Join3(void* ptr, char* separator){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join(*separator));
}

void* QByteArrayList_Join2(void* ptr, void* separator){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join(*static_cast<QByteArray*>(separator)));
}

void* QByteArrayMatcher_NewQByteArrayMatcher(){
	return new QByteArrayMatcher();
}

void* QByteArrayMatcher_NewQByteArrayMatcher2(void* pattern){
	return new QByteArrayMatcher(*static_cast<QByteArray*>(pattern));
}

void* QByteArrayMatcher_NewQByteArrayMatcher4(void* other){
	return new QByteArrayMatcher(*static_cast<QByteArrayMatcher*>(other));
}

void* QByteArrayMatcher_NewQByteArrayMatcher3(char* pattern, int length){
	return new QByteArrayMatcher(const_cast<const char*>(pattern), length);
}

int QByteArrayMatcher_IndexIn(void* ptr, void* ba, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(*static_cast<QByteArray*>(ba), from);
}

int QByteArrayMatcher_IndexIn2(void* ptr, char* str, int len, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(const_cast<const char*>(str), len, from);
}

void* QByteArrayMatcher_Pattern(void* ptr){
	return new QByteArray(static_cast<QByteArrayMatcher*>(ptr)->pattern());
}

void QByteArrayMatcher_SetPattern(void* ptr, void* pattern){
	static_cast<QByteArrayMatcher*>(ptr)->setPattern(*static_cast<QByteArray*>(pattern));
}

void QByteArrayMatcher_DestroyQByteArrayMatcher(void* ptr){
	static_cast<QByteArrayMatcher*>(ptr)->~QByteArrayMatcher();
}

void* QChar_NewQChar(){
	return new QChar();
}

void* QChar_NewQChar8(void* ch){
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

void* QChar_NewQChar7(int ch){
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

void* QChar_NewQChar9(char* ch){
	return new QChar(*ch);
}

void* QChar_NewQChar6(int code){
	return new QChar(code);
}

int QChar_Category(void* ptr){
	return static_cast<QChar*>(ptr)->category();
}

int QChar_QChar_CurrentUnicodeVersion(){
	return QChar::currentUnicodeVersion();
}

char* QChar_Decomposition(void* ptr){
	return static_cast<QChar*>(ptr)->decomposition().toUtf8().data();
}

int QChar_DecompositionTag(void* ptr){
	return static_cast<QChar*>(ptr)->decompositionTag();
}

int QChar_DigitValue(void* ptr){
	return static_cast<QChar*>(ptr)->digitValue();
}

int QChar_Direction(void* ptr){
	return static_cast<QChar*>(ptr)->direction();
}

int QChar_HasMirrored(void* ptr){
	return static_cast<QChar*>(ptr)->hasMirrored();
}

int QChar_IsDigit(void* ptr){
	return static_cast<QChar*>(ptr)->isDigit();
}

int QChar_IsHighSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isHighSurrogate();
}

int QChar_IsLetter(void* ptr){
	return static_cast<QChar*>(ptr)->isLetter();
}

int QChar_IsLetterOrNumber(void* ptr){
	return static_cast<QChar*>(ptr)->isLetterOrNumber();
}

int QChar_IsLower(void* ptr){
	return static_cast<QChar*>(ptr)->isLower();
}

int QChar_IsLowSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isLowSurrogate();
}

int QChar_IsMark(void* ptr){
	return static_cast<QChar*>(ptr)->isMark();
}

int QChar_IsNonCharacter(void* ptr){
	return static_cast<QChar*>(ptr)->isNonCharacter();
}

int QChar_IsNull(void* ptr){
	return static_cast<QChar*>(ptr)->isNull();
}

int QChar_IsNumber(void* ptr){
	return static_cast<QChar*>(ptr)->isNumber();
}

int QChar_IsPrint(void* ptr){
	return static_cast<QChar*>(ptr)->isPrint();
}

int QChar_IsPunct(void* ptr){
	return static_cast<QChar*>(ptr)->isPunct();
}

int QChar_IsSpace(void* ptr){
	return static_cast<QChar*>(ptr)->isSpace();
}

int QChar_IsSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isSurrogate();
}

int QChar_IsSymbol(void* ptr){
	return static_cast<QChar*>(ptr)->isSymbol();
}

int QChar_IsTitleCase(void* ptr){
	return static_cast<QChar*>(ptr)->isTitleCase();
}

int QChar_IsUpper(void* ptr){
	return static_cast<QChar*>(ptr)->isUpper();
}

int QChar_JoiningType(void* ptr){
	return static_cast<QChar*>(ptr)->joiningType();
}

int QChar_Script(void* ptr){
	return static_cast<QChar*>(ptr)->script();
}

int QChar_UnicodeVersion(void* ptr){
	return static_cast<QChar*>(ptr)->unicodeVersion();
}

void* QChildEvent_NewQChildEvent(int ty, void* child){
	return new QChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
}

int QChildEvent_Added(void* ptr){
	return static_cast<QChildEvent*>(ptr)->added();
}

void* QChildEvent_Child(void* ptr){
	return static_cast<QChildEvent*>(ptr)->child();
}

int QChildEvent_Polished(void* ptr){
	return static_cast<QChildEvent*>(ptr)->polished();
}

int QChildEvent_Removed(void* ptr){
	return static_cast<QChildEvent*>(ptr)->removed();
}

int QCollator_CaseSensitivity(void* ptr){
	return static_cast<QCollator*>(ptr)->caseSensitivity();
}

int QCollator_IgnorePunctuation(void* ptr){
	return static_cast<QCollator*>(ptr)->ignorePunctuation();
}

int QCollator_NumericMode(void* ptr){
	return static_cast<QCollator*>(ptr)->numericMode();
}

void QCollator_SetCaseSensitivity(void* ptr, int sensitivity){
	static_cast<QCollator*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(sensitivity));
}

void QCollator_SetIgnorePunctuation(void* ptr, int on){
	static_cast<QCollator*>(ptr)->setIgnorePunctuation(on != 0);
}

void QCollator_SetNumericMode(void* ptr, int on){
	static_cast<QCollator*>(ptr)->setNumericMode(on != 0);
}

void* QCollator_NewQCollator3(void* other){
	return new QCollator(*static_cast<QCollator*>(other));
}

void* QCollator_NewQCollator2(void* other){
	return new QCollator(*static_cast<QCollator*>(other));
}

void* QCollator_NewQCollator(void* locale){
	return new QCollator(*static_cast<QLocale*>(locale));
}

void QCollator_SetLocale(void* ptr, void* locale){
	static_cast<QCollator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QCollator_Swap(void* ptr, void* other){
	static_cast<QCollator*>(ptr)->swap(*static_cast<QCollator*>(other));
}

void QCollator_DestroyQCollator(void* ptr){
	static_cast<QCollator*>(ptr)->~QCollator();
}

int QCollator_Compare3(void* ptr, void* s1, int len1, void* s2, int len2){
	return static_cast<QCollator*>(ptr)->compare(static_cast<QChar*>(s1), len1, static_cast<QChar*>(s2), len2);
}

int QCollator_Compare(void* ptr, char* s1, char* s2){
	return static_cast<QCollator*>(ptr)->compare(QString(s1), QString(s2));
}

int QCollator_Compare2(void* ptr, void* s1, void* s2){
	return static_cast<QCollator*>(ptr)->compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

void* QCollatorSortKey_NewQCollatorSortKey(void* other){
	return new QCollatorSortKey(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_Swap(void* ptr, void* other){
	static_cast<QCollatorSortKey*>(ptr)->swap(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_DestroyQCollatorSortKey(void* ptr){
	static_cast<QCollatorSortKey*>(ptr)->~QCollatorSortKey();
}

int QCollatorSortKey_Compare(void* ptr, void* otherKey){
	return static_cast<QCollatorSortKey*>(ptr)->compare(*static_cast<QCollatorSortKey*>(otherKey));
}

void* QCommandLineOption_NewQCommandLineOption5(void* other){
	return new QCommandLineOption(*static_cast<QCommandLineOption*>(other));
}

void* QCommandLineOption_NewQCommandLineOption(char* name){
	return new QCommandLineOption(QString(name));
}

void* QCommandLineOption_NewQCommandLineOption3(char* name, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(name), QString(description), QString(valueName), QString(defaultValue));
}

void* QCommandLineOption_NewQCommandLineOption2(char* names){
	return new QCommandLineOption(QString(names).split(",,,", QString::SkipEmptyParts));
}

void* QCommandLineOption_NewQCommandLineOption4(char* names, char* description, char* valueName, char* defaultValue){
	return new QCommandLineOption(QString(names).split(",,,", QString::SkipEmptyParts), QString(description), QString(valueName), QString(defaultValue));
}

char* QCommandLineOption_DefaultValues(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->defaultValues().join(",,,").toUtf8().data();
}

char* QCommandLineOption_Description(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->description().toUtf8().data();
}

char* QCommandLineOption_Names(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->names().join(",,,").toUtf8().data();
}

void QCommandLineOption_SetDefaultValue(void* ptr, char* defaultValue){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValue(QString(defaultValue));
}

void QCommandLineOption_SetDefaultValues(void* ptr, char* defaultValues){
	static_cast<QCommandLineOption*>(ptr)->setDefaultValues(QString(defaultValues).split(",,,", QString::SkipEmptyParts));
}

void QCommandLineOption_SetDescription(void* ptr, char* description){
	static_cast<QCommandLineOption*>(ptr)->setDescription(QString(description));
}

void QCommandLineOption_SetValueName(void* ptr, char* valueName){
	static_cast<QCommandLineOption*>(ptr)->setValueName(QString(valueName));
}

void QCommandLineOption_Swap(void* ptr, void* other){
	static_cast<QCommandLineOption*>(ptr)->swap(*static_cast<QCommandLineOption*>(other));
}

char* QCommandLineOption_ValueName(void* ptr){
	return static_cast<QCommandLineOption*>(ptr)->valueName().toUtf8().data();
}

void QCommandLineOption_DestroyQCommandLineOption(void* ptr){
	static_cast<QCommandLineOption*>(ptr)->~QCommandLineOption();
}

void* QCommandLineParser_NewQCommandLineParser(){
	return new QCommandLineParser();
}

void* QCommandLineParser_AddHelpOption(void* ptr){
	return new QCommandLineOption(static_cast<QCommandLineParser*>(ptr)->addHelpOption());
}

int QCommandLineParser_AddOption(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->addOption(*static_cast<QCommandLineOption*>(option));
}

void QCommandLineParser_AddPositionalArgument(void* ptr, char* name, char* description, char* syntax){
	static_cast<QCommandLineParser*>(ptr)->addPositionalArgument(QString(name), QString(description), QString(syntax));
}

void* QCommandLineParser_AddVersionOption(void* ptr){
	return new QCommandLineOption(static_cast<QCommandLineParser*>(ptr)->addVersionOption());
}

char* QCommandLineParser_ApplicationDescription(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->applicationDescription().toUtf8().data();
}

void QCommandLineParser_ClearPositionalArguments(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->clearPositionalArguments();
}

char* QCommandLineParser_ErrorText(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->errorText().toUtf8().data();
}

char* QCommandLineParser_HelpText(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->helpText().toUtf8().data();
}

int QCommandLineParser_IsSet2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->isSet(*static_cast<QCommandLineOption*>(option));
}

int QCommandLineParser_IsSet(void* ptr, char* name){
	return static_cast<QCommandLineParser*>(ptr)->isSet(QString(name));
}

char* QCommandLineParser_OptionNames(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->optionNames().join(",,,").toUtf8().data();
}

int QCommandLineParser_Parse(void* ptr, char* arguments){
	return static_cast<QCommandLineParser*>(ptr)->parse(QString(arguments).split(",,,", QString::SkipEmptyParts));
}

char* QCommandLineParser_PositionalArguments(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->positionalArguments().join(",,,").toUtf8().data();
}

void QCommandLineParser_Process2(void* ptr, void* app){
	static_cast<QCommandLineParser*>(ptr)->process(*static_cast<QCoreApplication*>(app));
}

void QCommandLineParser_Process(void* ptr, char* arguments){
	static_cast<QCommandLineParser*>(ptr)->process(QString(arguments).split(",,,", QString::SkipEmptyParts));
}

void QCommandLineParser_SetApplicationDescription(void* ptr, char* description){
	static_cast<QCommandLineParser*>(ptr)->setApplicationDescription(QString(description));
}

void QCommandLineParser_SetSingleDashWordOptionMode(void* ptr, int singleDashWordOptionMode){
	static_cast<QCommandLineParser*>(ptr)->setSingleDashWordOptionMode(static_cast<QCommandLineParser::SingleDashWordOptionMode>(singleDashWordOptionMode));
}

void QCommandLineParser_ShowHelp(void* ptr, int exitCode){
	static_cast<QCommandLineParser*>(ptr)->showHelp(exitCode);
}

void QCommandLineParser_ShowVersion(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->showVersion();
}

char* QCommandLineParser_UnknownOptionNames(void* ptr){
	return static_cast<QCommandLineParser*>(ptr)->unknownOptionNames().join(",,,").toUtf8().data();
}

char* QCommandLineParser_Value2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->value(*static_cast<QCommandLineOption*>(option)).toUtf8().data();
}

char* QCommandLineParser_Value(void* ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->value(QString(optionName)).toUtf8().data();
}

char* QCommandLineParser_Values2(void* ptr, void* option){
	return static_cast<QCommandLineParser*>(ptr)->values(*static_cast<QCommandLineOption*>(option)).join(",,,").toUtf8().data();
}

char* QCommandLineParser_Values(void* ptr, char* optionName){
	return static_cast<QCommandLineParser*>(ptr)->values(QString(optionName)).join(",,,").toUtf8().data();
}

void QCommandLineParser_DestroyQCommandLineParser(void* ptr){
	static_cast<QCommandLineParser*>(ptr)->~QCommandLineParser();
}

class MyQCoreApplication: public QCoreApplication {
public:
	MyQCoreApplication(int &argc, char **argv) : QCoreApplication(argc, argv) {};
	void Signal_AboutToQuit() { callbackQCoreApplicationAboutToQuit(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQCoreApplicationTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCoreApplicationChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQCoreApplicationCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

char* QCoreApplication_QCoreApplication_ApplicationName(){
	return QCoreApplication::applicationName().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_ApplicationVersion(){
	return QCoreApplication::applicationVersion().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_OrganizationDomain(){
	return QCoreApplication::organizationDomain().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_OrganizationName(){
	return QCoreApplication::organizationName().toUtf8().data();
}

void QCoreApplication_QCoreApplication_SetApplicationName(char* application){
	QCoreApplication::setApplicationName(QString(application));
}

void QCoreApplication_QCoreApplication_SetApplicationVersion(char* version){
	QCoreApplication::setApplicationVersion(QString(version));
}

void QCoreApplication_QCoreApplication_SetOrganizationDomain(char* orgDomain){
	QCoreApplication::setOrganizationDomain(QString(orgDomain));
}

void QCoreApplication_QCoreApplication_SetOrganizationName(char* orgName){
	QCoreApplication::setOrganizationName(QString(orgName));
}

void* QCoreApplication_NewQCoreApplication(int argc, char* argv){
	QList<QByteArray> aList = QByteArray(argv).split(',,,');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = aList[i].data();

	return new MyQCoreApplication(argcs, argvs);
}

void QCoreApplication_ConnectAboutToQuit(void* ptr){
	QObject::connect(static_cast<QCoreApplication*>(ptr), &QCoreApplication::aboutToQuit, static_cast<MyQCoreApplication*>(ptr), static_cast<void (MyQCoreApplication::*)()>(&MyQCoreApplication::Signal_AboutToQuit));;
}

void QCoreApplication_DisconnectAboutToQuit(void* ptr){
	QObject::disconnect(static_cast<QCoreApplication*>(ptr), &QCoreApplication::aboutToQuit, static_cast<MyQCoreApplication*>(ptr), static_cast<void (MyQCoreApplication::*)()>(&MyQCoreApplication::Signal_AboutToQuit));;
}

void QCoreApplication_QCoreApplication_AddLibraryPath(char* path){
	QCoreApplication::addLibraryPath(QString(path));
}

char* QCoreApplication_QCoreApplication_ApplicationDirPath(){
	return QCoreApplication::applicationDirPath().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_ApplicationFilePath(){
	return QCoreApplication::applicationFilePath().toUtf8().data();
}

long long QCoreApplication_QCoreApplication_ApplicationPid(){
	return static_cast<long long>(QCoreApplication::applicationPid());
}

char* QCoreApplication_QCoreApplication_Arguments(){
	return QCoreApplication::arguments().join(",,,").toUtf8().data();
}

int QCoreApplication_QCoreApplication_ClosingDown(){
	return QCoreApplication::closingDown();
}

int QCoreApplication_Event(void* ptr, void* e){
	return static_cast<QCoreApplication*>(ptr)->event(static_cast<QEvent*>(e));
}

void* QCoreApplication_QCoreApplication_EventDispatcher(){
	return QCoreApplication::eventDispatcher();
}

int QCoreApplication_QCoreApplication_Exec(){
	return QCoreApplication::exec();
}

void QCoreApplication_QCoreApplication_Exit(int returnCode){
	QCoreApplication::exit(returnCode);
}

void QCoreApplication_QCoreApplication_Flush(){
	QCoreApplication::flush();
}

void QCoreApplication_InstallNativeEventFilter(void* ptr, void* filterObj){
	static_cast<QCoreApplication*>(ptr)->installNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObj));
}

int QCoreApplication_QCoreApplication_InstallTranslator(void* translationFile){
	return QCoreApplication::installTranslator(static_cast<QTranslator*>(translationFile));
}

void* QCoreApplication_QCoreApplication_Instance(){
	return QCoreApplication::instance();
}

int QCoreApplication_QCoreApplication_IsQuitLockEnabled(){
	return QCoreApplication::isQuitLockEnabled();
}

int QCoreApplication_QCoreApplication_IsSetuidAllowed(){
	return QCoreApplication::isSetuidAllowed();
}

char* QCoreApplication_QCoreApplication_LibraryPaths(){
	return QCoreApplication::libraryPaths().join(",,,").toUtf8().data();
}

int QCoreApplication_Notify(void* ptr, void* receiver, void* event){
	return static_cast<QCoreApplication*>(ptr)->notify(static_cast<QObject*>(receiver), static_cast<QEvent*>(event));
}

void QCoreApplication_QCoreApplication_PostEvent(void* receiver, void* event, int priority){
	QCoreApplication::postEvent(static_cast<QObject*>(receiver), static_cast<QEvent*>(event), priority);
}

void QCoreApplication_QCoreApplication_ProcessEvents(int flags){
	QCoreApplication::processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QCoreApplication_QCoreApplication_ProcessEvents2(int flags, int maxtime){
	QCoreApplication::processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags), maxtime);
}

void QCoreApplication_QCoreApplication_Quit(){
	QMetaObject::invokeMethod(QCoreApplication::instance(), "quit");
}

void QCoreApplication_QCoreApplication_RemoveLibraryPath(char* path){
	QCoreApplication::removeLibraryPath(QString(path));
}

void QCoreApplication_RemoveNativeEventFilter(void* ptr, void* filterObject){
	static_cast<QCoreApplication*>(ptr)->removeNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObject));
}

void QCoreApplication_QCoreApplication_RemovePostedEvents(void* receiver, int eventType){
	QCoreApplication::removePostedEvents(static_cast<QObject*>(receiver), eventType);
}

int QCoreApplication_QCoreApplication_RemoveTranslator(void* translationFile){
	return QCoreApplication::removeTranslator(static_cast<QTranslator*>(translationFile));
}

int QCoreApplication_QCoreApplication_SendEvent(void* receiver, void* event){
	return QCoreApplication::sendEvent(static_cast<QObject*>(receiver), static_cast<QEvent*>(event));
}

void QCoreApplication_QCoreApplication_SendPostedEvents(void* receiver, int event_type){
	QCoreApplication::sendPostedEvents(static_cast<QObject*>(receiver), event_type);
}

void QCoreApplication_QCoreApplication_SetAttribute(int attribute, int on){
	QCoreApplication::setAttribute(static_cast<Qt::ApplicationAttribute>(attribute), on != 0);
}

void QCoreApplication_QCoreApplication_SetEventDispatcher(void* eventDispatcher){
	QCoreApplication::setEventDispatcher(static_cast<QAbstractEventDispatcher*>(eventDispatcher));
}

void QCoreApplication_QCoreApplication_SetLibraryPaths(char* paths){
	QCoreApplication::setLibraryPaths(QString(paths).split(",,,", QString::SkipEmptyParts));
}

void QCoreApplication_QCoreApplication_SetQuitLockEnabled(int enabled){
	QCoreApplication::setQuitLockEnabled(enabled != 0);
}

void QCoreApplication_QCoreApplication_SetSetuidAllowed(int allow){
	QCoreApplication::setSetuidAllowed(allow != 0);
}

int QCoreApplication_QCoreApplication_StartingUp(){
	return QCoreApplication::startingUp();
}

int QCoreApplication_QCoreApplication_TestAttribute(int attribute){
	return QCoreApplication::testAttribute(static_cast<Qt::ApplicationAttribute>(attribute));
}

char* QCoreApplication_QCoreApplication_Translate(char* context, char* sourceText, char* disambiguation, int n){
	return QCoreApplication::translate(const_cast<const char*>(context), const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QCoreApplication_DestroyQCoreApplication(void* ptr){
	static_cast<QCoreApplication*>(ptr)->~QCoreApplication();
}

void QCoreApplication_TimerEvent(void* ptr, void* event){
	static_cast<MyQCoreApplication*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCoreApplication_TimerEventDefault(void* ptr, void* event){
	static_cast<QCoreApplication*>(ptr)->QCoreApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCoreApplication_ChildEvent(void* ptr, void* event){
	static_cast<MyQCoreApplication*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCoreApplication_ChildEventDefault(void* ptr, void* event){
	static_cast<QCoreApplication*>(ptr)->QCoreApplication::childEvent(static_cast<QChildEvent*>(event));
}

void QCoreApplication_CustomEvent(void* ptr, void* event){
	static_cast<MyQCoreApplication*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCoreApplication_CustomEventDefault(void* ptr, void* event){
	static_cast<QCoreApplication*>(ptr)->QCoreApplication::customEvent(static_cast<QEvent*>(event));
}

void* QCryptographicHash_NewQCryptographicHash(int method){
	return new QCryptographicHash(static_cast<QCryptographicHash::Algorithm>(method));
}

int QCryptographicHash_AddData2(void* ptr, void* device){
	return static_cast<QCryptographicHash*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QCryptographicHash_AddData3(void* ptr, void* data){
	static_cast<QCryptographicHash*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QCryptographicHash_AddData(void* ptr, char* data, int length){
	static_cast<QCryptographicHash*>(ptr)->addData(const_cast<const char*>(data), length);
}

void* QCryptographicHash_QCryptographicHash_Hash(void* data, int method){
	return new QByteArray(QCryptographicHash::hash(*static_cast<QByteArray*>(data), static_cast<QCryptographicHash::Algorithm>(method)));
}

void QCryptographicHash_Reset(void* ptr){
	static_cast<QCryptographicHash*>(ptr)->reset();
}

void* QCryptographicHash_Result(void* ptr){
	return new QByteArray(static_cast<QCryptographicHash*>(ptr)->result());
}

void QCryptographicHash_DestroyQCryptographicHash(void* ptr){
	static_cast<QCryptographicHash*>(ptr)->~QCryptographicHash();
}

void* QDataStream_NewQDataStream3(void* a, int mode){
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QDataStream_AtEnd(void* ptr){
	return static_cast<QDataStream*>(ptr)->atEnd();
}

void* QDataStream_NewQDataStream(){
	return new QDataStream();
}

void* QDataStream_NewQDataStream2(void* d){
	return new QDataStream(static_cast<QIODevice*>(d));
}

void* QDataStream_NewQDataStream4(void* a){
	return new QDataStream(*static_cast<QByteArray*>(a));
}

int QDataStream_ByteOrder(void* ptr){
	return static_cast<QDataStream*>(ptr)->byteOrder();
}

void* QDataStream_Device(void* ptr){
	return static_cast<QDataStream*>(ptr)->device();
}

int QDataStream_FloatingPointPrecision(void* ptr){
	return static_cast<QDataStream*>(ptr)->floatingPointPrecision();
}

int QDataStream_ReadRawData(void* ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->readRawData(s, len);
}

void QDataStream_ResetStatus(void* ptr){
	static_cast<QDataStream*>(ptr)->resetStatus();
}

void QDataStream_SetByteOrder(void* ptr, int bo){
	static_cast<QDataStream*>(ptr)->setByteOrder(static_cast<QDataStream::ByteOrder>(bo));
}

void QDataStream_SetDevice(void* ptr, void* d){
	static_cast<QDataStream*>(ptr)->setDevice(static_cast<QIODevice*>(d));
}

void QDataStream_SetFloatingPointPrecision(void* ptr, int precision){
	static_cast<QDataStream*>(ptr)->setFloatingPointPrecision(static_cast<QDataStream::FloatingPointPrecision>(precision));
}

void QDataStream_SetStatus(void* ptr, int status){
	static_cast<QDataStream*>(ptr)->setStatus(static_cast<QDataStream::Status>(status));
}

void QDataStream_SetVersion(void* ptr, int v){
	static_cast<QDataStream*>(ptr)->setVersion(v);
}

int QDataStream_SkipRawData(void* ptr, int len){
	return static_cast<QDataStream*>(ptr)->skipRawData(len);
}

int QDataStream_Status(void* ptr){
	return static_cast<QDataStream*>(ptr)->status();
}

int QDataStream_Version(void* ptr){
	return static_cast<QDataStream*>(ptr)->version();
}

int QDataStream_WriteRawData(void* ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->writeRawData(const_cast<const char*>(s), len);
}

void QDataStream_DestroyQDataStream(void* ptr){
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

int QDate_QDate_IsLeapYear(int year){
	return QDate::isLeapYear(year);
}

char* QDate_ToString2(void* ptr, int format){
	return static_cast<QDate*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

void* QDate_NewQDate(){
	return new QDate();
}

void* QDate_NewQDate3(int y, int m, int d){
	return new QDate(y, m, d);
}

int QDate_Day(void* ptr){
	return static_cast<QDate*>(ptr)->day();
}

int QDate_DayOfWeek(void* ptr){
	return static_cast<QDate*>(ptr)->dayOfWeek();
}

int QDate_DayOfYear(void* ptr){
	return static_cast<QDate*>(ptr)->dayOfYear();
}

int QDate_DaysInMonth(void* ptr){
	return static_cast<QDate*>(ptr)->daysInMonth();
}

int QDate_DaysInYear(void* ptr){
	return static_cast<QDate*>(ptr)->daysInYear();
}

long long QDate_DaysTo(void* ptr, void* d){
	return static_cast<long long>(static_cast<QDate*>(ptr)->daysTo(*static_cast<QDate*>(d)));
}

void QDate_GetDate(void* ptr, int year, int month, int day){
	static_cast<QDate*>(ptr)->getDate(&year, &month, &day);
}

int QDate_IsNull(void* ptr){
	return static_cast<QDate*>(ptr)->isNull();
}

int QDate_QDate_IsValid2(int year, int month, int day){
	return QDate::isValid(year, month, day);
}

int QDate_IsValid(void* ptr){
	return static_cast<QDate*>(ptr)->isValid();
}

char* QDate_QDate_LongDayName(int weekday, int ty){
	return QDate::longDayName(weekday, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

char* QDate_QDate_LongMonthName(int month, int ty){
	return QDate::longMonthName(month, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

int QDate_Month(void* ptr){
	return static_cast<QDate*>(ptr)->month();
}

int QDate_SetDate(void* ptr, int year, int month, int day){
	return static_cast<QDate*>(ptr)->setDate(year, month, day);
}

char* QDate_QDate_ShortDayName(int weekday, int ty){
	return QDate::shortDayName(weekday, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

char* QDate_QDate_ShortMonthName(int month, int ty){
	return QDate::shortMonthName(month, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

long long QDate_ToJulianDay(void* ptr){
	return static_cast<long long>(static_cast<QDate*>(ptr)->toJulianDay());
}

char* QDate_ToString(void* ptr, char* format){
	return static_cast<QDate*>(ptr)->toString(QString(format)).toUtf8().data();
}

int QDate_WeekNumber(void* ptr, int yearNumber){
	return static_cast<QDate*>(ptr)->weekNumber(&yearNumber);
}

int QDate_Year(void* ptr){
	return static_cast<QDate*>(ptr)->year();
}

void* QDateTime_QDateTime_CurrentDateTime(){
	return new QDateTime(QDateTime::currentDateTime());
}

void* QDateTime_QDateTime_CurrentDateTimeUtc(){
	return new QDateTime(QDateTime::currentDateTimeUtc());
}

long long QDateTime_QDateTime_CurrentMSecsSinceEpoch(){
	return static_cast<long long>(QDateTime::currentMSecsSinceEpoch());
}

void* QDateTime_QDateTime_FromString(char* stri, int format){
	return new QDateTime(QDateTime::fromString(QString(stri), static_cast<Qt::DateFormat>(format)));
}

void* QDateTime_QDateTime_FromString2(char* stri, char* format){
	return new QDateTime(QDateTime::fromString(QString(stri), QString(format)));
}

void* QDateTime_ToOffsetFromUtc(void* ptr, int offsetSeconds){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toOffsetFromUtc(offsetSeconds));
}

char* QDateTime_ToString2(void* ptr, int format){
	return static_cast<QDateTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

void* QDateTime_ToTimeSpec(void* ptr, int spec){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toTimeSpec(static_cast<Qt::TimeSpec>(spec)));
}

void* QDateTime_NewQDateTime(){
	return new QDateTime();
}

void* QDateTime_NewQDateTime2(void* date){
	return new QDateTime(*static_cast<QDate*>(date));
}

void* QDateTime_NewQDateTime3(void* date, void* time, int spec){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec));
}

void* QDateTime_NewQDateTime4(void* date, void* time, int spec, int offsetSeconds){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

void* QDateTime_NewQDateTime5(void* date, void* time, void* timeZone){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), *static_cast<QTimeZone*>(timeZone));
}

void* QDateTime_NewQDateTime6(void* other){
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_AddDays(void* ptr, long long ndays){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addDays(static_cast<long long>(ndays)));
}

void* QDateTime_AddMSecs(void* ptr, long long msecs){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addMSecs(static_cast<long long>(msecs)));
}

void* QDateTime_AddMonths(void* ptr, int nmonths){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addMonths(nmonths));
}

void* QDateTime_AddSecs(void* ptr, long long s){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addSecs(static_cast<long long>(s)));
}

void* QDateTime_AddYears(void* ptr, int nyears){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addYears(nyears));
}

long long QDateTime_DaysTo(void* ptr, void* other){
	return static_cast<long long>(static_cast<QDateTime*>(ptr)->daysTo(*static_cast<QDateTime*>(other)));
}

void* QDateTime_QDateTime_FromMSecsSinceEpoch(long long msecs){
	return new QDateTime(QDateTime::fromMSecsSinceEpoch(static_cast<long long>(msecs)));
}

void* QDateTime_QDateTime_FromMSecsSinceEpoch2(long long msecs, int spec, int offsetSeconds){
	return new QDateTime(QDateTime::fromMSecsSinceEpoch(static_cast<long long>(msecs), static_cast<Qt::TimeSpec>(spec), offsetSeconds));
}

void* QDateTime_QDateTime_FromMSecsSinceEpoch3(long long msecs, void* timeZone){
	return new QDateTime(QDateTime::fromMSecsSinceEpoch(static_cast<long long>(msecs), *static_cast<QTimeZone*>(timeZone)));
}

int QDateTime_IsDaylightTime(void* ptr){
	return static_cast<QDateTime*>(ptr)->isDaylightTime();
}

int QDateTime_IsNull(void* ptr){
	return static_cast<QDateTime*>(ptr)->isNull();
}

int QDateTime_IsValid(void* ptr){
	return static_cast<QDateTime*>(ptr)->isValid();
}

long long QDateTime_MsecsTo(void* ptr, void* other){
	return static_cast<long long>(static_cast<QDateTime*>(ptr)->msecsTo(*static_cast<QDateTime*>(other)));
}

int QDateTime_OffsetFromUtc(void* ptr){
	return static_cast<QDateTime*>(ptr)->offsetFromUtc();
}

long long QDateTime_SecsTo(void* ptr, void* other){
	return static_cast<long long>(static_cast<QDateTime*>(ptr)->secsTo(*static_cast<QDateTime*>(other)));
}

void QDateTime_SetDate(void* ptr, void* date){
	static_cast<QDateTime*>(ptr)->setDate(*static_cast<QDate*>(date));
}

void QDateTime_SetMSecsSinceEpoch(void* ptr, long long msecs){
	static_cast<QDateTime*>(ptr)->setMSecsSinceEpoch(static_cast<long long>(msecs));
}

void QDateTime_SetOffsetFromUtc(void* ptr, int offsetSeconds){
	static_cast<QDateTime*>(ptr)->setOffsetFromUtc(offsetSeconds);
}

void QDateTime_SetTime(void* ptr, void* time){
	static_cast<QDateTime*>(ptr)->setTime(*static_cast<QTime*>(time));
}

void QDateTime_SetTimeSpec(void* ptr, int spec){
	static_cast<QDateTime*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

void QDateTime_SetTimeZone(void* ptr, void* toZone){
	static_cast<QDateTime*>(ptr)->setTimeZone(*static_cast<QTimeZone*>(toZone));
}

void QDateTime_Swap(void* ptr, void* other){
	static_cast<QDateTime*>(ptr)->swap(*static_cast<QDateTime*>(other));
}

int QDateTime_TimeSpec(void* ptr){
	return static_cast<QDateTime*>(ptr)->timeSpec();
}

void* QDateTime_TimeZone(void* ptr){
	return new QTimeZone(static_cast<QDateTime*>(ptr)->timeZone());
}

char* QDateTime_TimeZoneAbbreviation(void* ptr){
	return static_cast<QDateTime*>(ptr)->timeZoneAbbreviation().toUtf8().data();
}

void* QDateTime_ToLocalTime(void* ptr){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toLocalTime());
}

long long QDateTime_ToMSecsSinceEpoch(void* ptr){
	return static_cast<long long>(static_cast<QDateTime*>(ptr)->toMSecsSinceEpoch());
}

char* QDateTime_ToString(void* ptr, char* format){
	return static_cast<QDateTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

void* QDateTime_ToTimeZone(void* ptr, void* timeZone){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toTimeZone(*static_cast<QTimeZone*>(timeZone)));
}

void* QDateTime_ToUTC(void* ptr){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toUTC());
}

void QDateTime_DestroyQDateTime(void* ptr){
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

void* QDebugStateSaver_NewQDebugStateSaver(void* dbg){
	return new QDebugStateSaver(*static_cast<QDebug*>(dbg));
}

void QDebugStateSaver_DestroyQDebugStateSaver(void* ptr){
	static_cast<QDebugStateSaver*>(ptr)->~QDebugStateSaver();
}

void* QDir_NewQDir(void* dir){
	return new QDir(*static_cast<QDir*>(dir));
}

void* QDir_NewQDir2(char* path){
	return new QDir(QString(path));
}

void* QDir_NewQDir3(char* path, char* nameFilter, int sort, int filters){
	return new QDir(QString(path), QString(nameFilter), static_cast<QDir::SortFlag>(sort), static_cast<QDir::Filter>(filters));
}

char* QDir_AbsoluteFilePath(void* ptr, char* fileName){
	return static_cast<QDir*>(ptr)->absoluteFilePath(QString(fileName)).toUtf8().data();
}

char* QDir_AbsolutePath(void* ptr){
	return static_cast<QDir*>(ptr)->absolutePath().toUtf8().data();
}

void QDir_QDir_AddSearchPath(char* prefix, char* path){
	QDir::addSearchPath(QString(prefix), QString(path));
}

char* QDir_CanonicalPath(void* ptr){
	return static_cast<QDir*>(ptr)->canonicalPath().toUtf8().data();
}

int QDir_Cd(void* ptr, char* dirName){
	return static_cast<QDir*>(ptr)->cd(QString(dirName));
}

int QDir_CdUp(void* ptr){
	return static_cast<QDir*>(ptr)->cdUp();
}

char* QDir_QDir_CleanPath(char* path){
	return QDir::cleanPath(QString(path)).toUtf8().data();
}

void* QDir_QDir_Current(){
	return new QDir(QDir::current());
}

char* QDir_QDir_CurrentPath(){
	return QDir::currentPath().toUtf8().data();
}

char* QDir_DirName(void* ptr){
	return static_cast<QDir*>(ptr)->dirName().toUtf8().data();
}

char* QDir_EntryList2(void* ptr, int filters, int sort){
	return static_cast<QDir*>(ptr)->entryList(static_cast<QDir::Filter>(filters), static_cast<QDir::SortFlag>(sort)).join(",,,").toUtf8().data();
}

char* QDir_EntryList(void* ptr, char* nameFilters, int filters, int sort){
	return static_cast<QDir*>(ptr)->entryList(QString(nameFilters).split(",,,", QString::SkipEmptyParts), static_cast<QDir::Filter>(filters), static_cast<QDir::SortFlag>(sort)).join(",,,").toUtf8().data();
}

int QDir_Exists2(void* ptr){
	return static_cast<QDir*>(ptr)->exists();
}

int QDir_Exists(void* ptr, char* name){
	return static_cast<QDir*>(ptr)->exists(QString(name));
}

char* QDir_FilePath(void* ptr, char* fileName){
	return static_cast<QDir*>(ptr)->filePath(QString(fileName)).toUtf8().data();
}

int QDir_Filter(void* ptr){
	return static_cast<QDir*>(ptr)->filter();
}

char* QDir_QDir_FromNativeSeparators(char* pathName){
	return QDir::fromNativeSeparators(QString(pathName)).toUtf8().data();
}

void* QDir_QDir_Home(){
	return new QDir(QDir::home());
}

char* QDir_QDir_HomePath(){
	return QDir::homePath().toUtf8().data();
}

int QDir_IsAbsolute(void* ptr){
	return static_cast<QDir*>(ptr)->isAbsolute();
}

int QDir_QDir_IsAbsolutePath(char* path){
	return QDir::isAbsolutePath(QString(path));
}

int QDir_IsReadable(void* ptr){
	return static_cast<QDir*>(ptr)->isReadable();
}

int QDir_IsRelative(void* ptr){
	return static_cast<QDir*>(ptr)->isRelative();
}

int QDir_QDir_IsRelativePath(char* path){
	return QDir::isRelativePath(QString(path));
}

int QDir_IsRoot(void* ptr){
	return static_cast<QDir*>(ptr)->isRoot();
}

int QDir_MakeAbsolute(void* ptr){
	return static_cast<QDir*>(ptr)->makeAbsolute();
}

int QDir_QDir_Match(char* filter, char* fileName){
	return QDir::match(QString(filter), QString(fileName));
}

int QDir_QDir_Match2(char* filters, char* fileName){
	return QDir::match(QString(filters).split(",,,", QString::SkipEmptyParts), QString(fileName));
}

int QDir_Mkdir(void* ptr, char* dirName){
	return static_cast<QDir*>(ptr)->mkdir(QString(dirName));
}

int QDir_Mkpath(void* ptr, char* dirPath){
	return static_cast<QDir*>(ptr)->mkpath(QString(dirPath));
}

char* QDir_NameFilters(void* ptr){
	return static_cast<QDir*>(ptr)->nameFilters().join(",,,").toUtf8().data();
}

char* QDir_Path(void* ptr){
	return static_cast<QDir*>(ptr)->path().toUtf8().data();
}

void QDir_Refresh(void* ptr){
	static_cast<QDir*>(ptr)->refresh();
}

char* QDir_RelativeFilePath(void* ptr, char* fileName){
	return static_cast<QDir*>(ptr)->relativeFilePath(QString(fileName)).toUtf8().data();
}

int QDir_RemoveRecursively(void* ptr){
	return static_cast<QDir*>(ptr)->removeRecursively();
}

int QDir_Rename(void* ptr, char* oldName, char* newName){
	return static_cast<QDir*>(ptr)->rename(QString(oldName), QString(newName));
}

int QDir_Rmdir(void* ptr, char* dirName){
	return static_cast<QDir*>(ptr)->rmdir(QString(dirName));
}

int QDir_Rmpath(void* ptr, char* dirPath){
	return static_cast<QDir*>(ptr)->rmpath(QString(dirPath));
}

void* QDir_QDir_Root(){
	return new QDir(QDir::root());
}

char* QDir_QDir_RootPath(){
	return QDir::rootPath().toUtf8().data();
}

char* QDir_QDir_SearchPaths(char* prefix){
	return QDir::searchPaths(QString(prefix)).join(",,,").toUtf8().data();
}

int QDir_QDir_SetCurrent(char* path){
	return QDir::setCurrent(QString(path));
}

void QDir_SetFilter(void* ptr, int filters){
	static_cast<QDir*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QDir_SetNameFilters(void* ptr, char* nameFilters){
	static_cast<QDir*>(ptr)->setNameFilters(QString(nameFilters).split(",,,", QString::SkipEmptyParts));
}

void QDir_SetPath(void* ptr, char* path){
	static_cast<QDir*>(ptr)->setPath(QString(path));
}

void QDir_QDir_SetSearchPaths(char* prefix, char* searchPaths){
	QDir::setSearchPaths(QString(prefix), QString(searchPaths).split(",,,", QString::SkipEmptyParts));
}

void QDir_SetSorting(void* ptr, int sort){
	static_cast<QDir*>(ptr)->setSorting(static_cast<QDir::SortFlag>(sort));
}

int QDir_Sorting(void* ptr){
	return static_cast<QDir*>(ptr)->sorting();
}

void QDir_Swap(void* ptr, void* other){
	static_cast<QDir*>(ptr)->swap(*static_cast<QDir*>(other));
}

void* QDir_QDir_Temp(){
	return new QDir(QDir::temp());
}

char* QDir_QDir_TempPath(){
	return QDir::tempPath().toUtf8().data();
}

char* QDir_QDir_ToNativeSeparators(char* pathName){
	return QDir::toNativeSeparators(QString(pathName)).toUtf8().data();
}

void QDir_DestroyQDir(void* ptr){
	static_cast<QDir*>(ptr)->~QDir();
}

void* QDynamicPropertyChangeEvent_NewQDynamicPropertyChangeEvent(void* name){
	return new QDynamicPropertyChangeEvent(*static_cast<QByteArray*>(name));
}

void* QDynamicPropertyChangeEvent_PropertyName(void* ptr){
	return new QByteArray(static_cast<QDynamicPropertyChangeEvent*>(ptr)->propertyName());
}

void* QEasingCurve_NewQEasingCurve3(void* other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void* QEasingCurve_NewQEasingCurve(int ty){
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

void* QEasingCurve_NewQEasingCurve2(void* other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void QEasingCurve_AddCubicBezierSegment(void* ptr, void* c1, void* c2, void* endPoint){
	static_cast<QEasingCurve*>(ptr)->addCubicBezierSegment(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

void QEasingCurve_AddTCBSegment(void* ptr, void* nextPoint, double t, double c, double b){
	static_cast<QEasingCurve*>(ptr)->addTCBSegment(*static_cast<QPointF*>(nextPoint), static_cast<double>(t), static_cast<double>(c), static_cast<double>(b));
}

double QEasingCurve_Amplitude(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->amplitude());
}

double QEasingCurve_Overshoot(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->overshoot());
}

double QEasingCurve_Period(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->period());
}

void QEasingCurve_SetAmplitude(void* ptr, double amplitude){
	static_cast<QEasingCurve*>(ptr)->setAmplitude(static_cast<double>(amplitude));
}

void QEasingCurve_SetOvershoot(void* ptr, double overshoot){
	static_cast<QEasingCurve*>(ptr)->setOvershoot(static_cast<double>(overshoot));
}

void QEasingCurve_SetPeriod(void* ptr, double period){
	static_cast<QEasingCurve*>(ptr)->setPeriod(static_cast<double>(period));
}

void QEasingCurve_SetType(void* ptr, int ty){
	static_cast<QEasingCurve*>(ptr)->setType(static_cast<QEasingCurve::Type>(ty));
}

void QEasingCurve_Swap(void* ptr, void* other){
	static_cast<QEasingCurve*>(ptr)->swap(*static_cast<QEasingCurve*>(other));
}

int QEasingCurve_Type(void* ptr){
	return static_cast<QEasingCurve*>(ptr)->type();
}

double QEasingCurve_ValueForProgress(void* ptr, double progress){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->valueForProgress(static_cast<double>(progress)));
}

void QEasingCurve_DestroyQEasingCurve(void* ptr){
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

void* QElapsedTimer_NewQElapsedTimer(){
	return new QElapsedTimer();
}

int QElapsedTimer_HasExpired(void* ptr, long long timeout){
	return static_cast<QElapsedTimer*>(ptr)->hasExpired(static_cast<long long>(timeout));
}

void QElapsedTimer_Invalidate(void* ptr){
	static_cast<QElapsedTimer*>(ptr)->invalidate();
}

int QElapsedTimer_IsValid(void* ptr){
	return static_cast<QElapsedTimer*>(ptr)->isValid();
}

int QElapsedTimer_QElapsedTimer_ClockType(){
	return QElapsedTimer::clockType();
}

long long QElapsedTimer_Elapsed(void* ptr){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->elapsed());
}

int QElapsedTimer_QElapsedTimer_IsMonotonic(){
	return QElapsedTimer::isMonotonic();
}

long long QElapsedTimer_MsecsSinceReference(void* ptr){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->msecsSinceReference());
}

long long QElapsedTimer_MsecsTo(void* ptr, void* other){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->msecsTo(*static_cast<QElapsedTimer*>(other)));
}

long long QElapsedTimer_NsecsElapsed(void* ptr){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->nsecsElapsed());
}

long long QElapsedTimer_Restart(void* ptr){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->restart());
}

long long QElapsedTimer_SecsTo(void* ptr, void* other){
	return static_cast<long long>(static_cast<QElapsedTimer*>(ptr)->secsTo(*static_cast<QElapsedTimer*>(other)));
}

void QElapsedTimer_Start(void* ptr){
	static_cast<QElapsedTimer*>(ptr)->start();
}

class MyQEvent: public QEvent {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQEvent(Type type) : QEvent(type) {};
};

void* QEvent_NewQEvent(int ty){
	return new MyQEvent(static_cast<QEvent::Type>(ty));
}

void QEvent_Accept(void* ptr){
	static_cast<QEvent*>(ptr)->accept();
}

void QEvent_Ignore(void* ptr){
	static_cast<QEvent*>(ptr)->ignore();
}

int QEvent_IsAccepted(void* ptr){
	return static_cast<QEvent*>(ptr)->isAccepted();
}

int QEvent_QEvent_RegisterEventType(int hint){
	return QEvent::registerEventType(hint);
}

void QEvent_SetAccepted(void* ptr, int accepted){
	static_cast<QEvent*>(ptr)->setAccepted(accepted != 0);
}

int QEvent_Spontaneous(void* ptr){
	return static_cast<QEvent*>(ptr)->spontaneous();
}

int QEvent_Type(void* ptr){
	return static_cast<QEvent*>(ptr)->type();
}

void QEvent_DestroyQEvent(void* ptr){
	static_cast<QEvent*>(ptr)->~QEvent();
}

char* QEvent_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQEvent*>(static_cast<QEvent*>(ptr))) {
		return static_cast<MyQEvent*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QEvent_BASE").toUtf8().data();
}

void QEvent_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQEvent*>(static_cast<QEvent*>(ptr))) {
		static_cast<MyQEvent*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQEventLoop: public QEventLoop {
public:
	MyQEventLoop(QObject *parent) : QEventLoop(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQEventLoopTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQEventLoopChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQEventLoopCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QEventLoop_NewQEventLoop(void* parent){
	return new MyQEventLoop(static_cast<QObject*>(parent));
}

int QEventLoop_Event(void* ptr, void* event){
	return static_cast<QEventLoop*>(ptr)->event(static_cast<QEvent*>(event));
}

int QEventLoop_Exec(void* ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->exec(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_Exit(void* ptr, int returnCode){
	static_cast<QEventLoop*>(ptr)->exit(returnCode);
}

int QEventLoop_IsRunning(void* ptr){
	return static_cast<QEventLoop*>(ptr)->isRunning();
}

int QEventLoop_ProcessEvents(void* ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_ProcessEvents2(void* ptr, int flags, int maxTime){
	static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags), maxTime);
}

void QEventLoop_Quit(void* ptr){
	QMetaObject::invokeMethod(static_cast<QEventLoop*>(ptr), "quit");
}

void QEventLoop_WakeUp(void* ptr){
	static_cast<QEventLoop*>(ptr)->wakeUp();
}

void QEventLoop_DestroyQEventLoop(void* ptr){
	static_cast<QEventLoop*>(ptr)->~QEventLoop();
}

void QEventLoop_TimerEvent(void* ptr, void* event){
	static_cast<MyQEventLoop*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QEventLoop_TimerEventDefault(void* ptr, void* event){
	static_cast<QEventLoop*>(ptr)->QEventLoop::timerEvent(static_cast<QTimerEvent*>(event));
}

void QEventLoop_ChildEvent(void* ptr, void* event){
	static_cast<MyQEventLoop*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QEventLoop_ChildEventDefault(void* ptr, void* event){
	static_cast<QEventLoop*>(ptr)->QEventLoop::childEvent(static_cast<QChildEvent*>(event));
}

void QEventLoop_CustomEvent(void* ptr, void* event){
	static_cast<MyQEventLoop*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QEventLoop_CustomEventDefault(void* ptr, void* event){
	static_cast<QEventLoop*>(ptr)->QEventLoop::customEvent(static_cast<QEvent*>(event));
}

void* QEventLoopLocker_NewQEventLoopLocker(){
	return new QEventLoopLocker();
}

void* QEventLoopLocker_NewQEventLoopLocker2(void* loop){
	return new QEventLoopLocker(static_cast<QEventLoop*>(loop));
}

void* QEventLoopLocker_NewQEventLoopLocker3(void* thread){
	return new QEventLoopLocker(static_cast<QThread*>(thread));
}

void QEventLoopLocker_DestroyQEventLoopLocker(void* ptr){
	static_cast<QEventLoopLocker*>(ptr)->~QEventLoopLocker();
}

class MyQEventTransition: public QEventTransition {
public:
	MyQEventTransition(QObject *object, QEvent::Type type, QState *sourceState) : QEventTransition(object, type, sourceState) {};
	MyQEventTransition(QState *sourceState) : QEventTransition(sourceState) {};
	void onTransition(QEvent * event) { callbackQEventTransitionOnTransition(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQEventTransitionTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQEventTransitionChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQEventTransitionCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QEventTransition_NewQEventTransition2(void* object, int ty, void* sourceState){
	return new MyQEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<QState*>(sourceState));
}

void* QEventTransition_NewQEventTransition(void* sourceState){
	return new MyQEventTransition(static_cast<QState*>(sourceState));
}

int QEventTransition_Event(void* ptr, void* e){
	return static_cast<QEventTransition*>(ptr)->event(static_cast<QEvent*>(e));
}

void* QEventTransition_EventSource(void* ptr){
	return static_cast<QEventTransition*>(ptr)->eventSource();
}

int QEventTransition_EventTest(void* ptr, void* event){
	return static_cast<QEventTransition*>(ptr)->eventTest(static_cast<QEvent*>(event));
}

int QEventTransition_EventType(void* ptr){
	return static_cast<QEventTransition*>(ptr)->eventType();
}

void QEventTransition_OnTransition(void* ptr, void* event){
	static_cast<MyQEventTransition*>(ptr)->onTransition(static_cast<QEvent*>(event));
}

void QEventTransition_OnTransitionDefault(void* ptr, void* event){
	static_cast<QEventTransition*>(ptr)->QEventTransition::onTransition(static_cast<QEvent*>(event));
}

void QEventTransition_SetEventSource(void* ptr, void* object){
	static_cast<QEventTransition*>(ptr)->setEventSource(static_cast<QObject*>(object));
}

void QEventTransition_SetEventType(void* ptr, int ty){
	static_cast<QEventTransition*>(ptr)->setEventType(static_cast<QEvent::Type>(ty));
}

void QEventTransition_DestroyQEventTransition(void* ptr){
	static_cast<QEventTransition*>(ptr)->~QEventTransition();
}

void QEventTransition_TimerEvent(void* ptr, void* event){
	static_cast<MyQEventTransition*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QEventTransition_TimerEventDefault(void* ptr, void* event){
	static_cast<QEventTransition*>(ptr)->QEventTransition::timerEvent(static_cast<QTimerEvent*>(event));
}

void QEventTransition_ChildEvent(void* ptr, void* event){
	static_cast<MyQEventTransition*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QEventTransition_ChildEventDefault(void* ptr, void* event){
	static_cast<QEventTransition*>(ptr)->QEventTransition::childEvent(static_cast<QChildEvent*>(event));
}

void QEventTransition_CustomEvent(void* ptr, void* event){
	static_cast<MyQEventTransition*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QEventTransition_CustomEventDefault(void* ptr, void* event){
	static_cast<QEventTransition*>(ptr)->QEventTransition::customEvent(static_cast<QEvent*>(event));
}

class MyQFile: public QFile {
public:
	MyQFile(QObject *parent) : QFile(parent) {};
	MyQFile(const QString &name) : QFile(name) {};
	MyQFile(const QString &name, QObject *parent) : QFile(name, parent) {};
	void close() { callbackQFileClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQFileTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQFileChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQFileCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QFile_NewQFile3(void* parent){
	return new MyQFile(static_cast<QObject*>(parent));
}

void* QFile_NewQFile(char* name){
	return new MyQFile(QString(name));
}

void* QFile_NewQFile4(char* name, void* parent){
	return new MyQFile(QString(name), static_cast<QObject*>(parent));
}

int QFile_QFile_Copy2(char* fileName, char* newName){
	return QFile::copy(QString(fileName), QString(newName));
}

int QFile_Copy(void* ptr, char* newName){
	return static_cast<QFile*>(ptr)->copy(QString(newName));
}

char* QFile_QFile_DecodeName(void* localFileName){
	return QFile::decodeName(*static_cast<QByteArray*>(localFileName)).toUtf8().data();
}

char* QFile_QFile_DecodeName2(char* localFileName){
	return QFile::decodeName(const_cast<const char*>(localFileName)).toUtf8().data();
}

void* QFile_QFile_EncodeName(char* fileName){
	return new QByteArray(QFile::encodeName(QString(fileName)));
}

int QFile_QFile_Exists(char* fileName){
	return QFile::exists(QString(fileName));
}

int QFile_Exists2(void* ptr){
	return static_cast<QFile*>(ptr)->exists();
}

char* QFile_FileName(void* ptr){
	return static_cast<QFile*>(ptr)->fileName().toUtf8().data();
}

int QFile_QFile_Link2(char* fileName, char* linkName){
	return QFile::link(QString(fileName), QString(linkName));
}

int QFile_Link(void* ptr, char* linkName){
	return static_cast<QFile*>(ptr)->link(QString(linkName));
}

int QFile_Open(void* ptr, int mode){
	return static_cast<QFile*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QFile_Open3(void* ptr, int fd, int mode, int handleFlags){
	return static_cast<QFile*>(ptr)->open(fd, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QFileDevice::FileHandleFlag>(handleFlags));
}

int QFile_QFile_Permissions2(char* fileName){
	return QFile::permissions(QString(fileName));
}

int QFile_Permissions(void* ptr){
	return static_cast<QFile*>(ptr)->permissions();
}

int QFile_Rename(void* ptr, char* newName){
	return static_cast<QFile*>(ptr)->rename(QString(newName));
}

int QFile_QFile_Rename2(char* oldName, char* newName){
	return QFile::rename(QString(oldName), QString(newName));
}

int QFile_QFile_Resize2(char* fileName, long long sz){
	return QFile::resize(QString(fileName), static_cast<long long>(sz));
}

int QFile_Resize(void* ptr, long long sz){
	return static_cast<QFile*>(ptr)->resize(static_cast<long long>(sz));
}

void QFile_SetFileName(void* ptr, char* name){
	static_cast<QFile*>(ptr)->setFileName(QString(name));
}

int QFile_SetPermissions(void* ptr, int permissions){
	return static_cast<QFile*>(ptr)->setPermissions(static_cast<QFileDevice::Permission>(permissions));
}

int QFile_QFile_SetPermissions2(char* fileName, int permissions){
	return QFile::setPermissions(QString(fileName), static_cast<QFileDevice::Permission>(permissions));
}

long long QFile_Size(void* ptr){
	return static_cast<long long>(static_cast<QFile*>(ptr)->size());
}

char* QFile_QFile_SymLinkTarget(char* fileName){
	return QFile::symLinkTarget(QString(fileName)).toUtf8().data();
}

char* QFile_SymLinkTarget2(void* ptr){
	return static_cast<QFile*>(ptr)->symLinkTarget().toUtf8().data();
}

void QFile_DestroyQFile(void* ptr){
	static_cast<QFile*>(ptr)->~QFile();
}

void QFile_Close(void* ptr){
	static_cast<MyQFile*>(ptr)->close();
}

void QFile_CloseDefault(void* ptr){
	static_cast<QFile*>(ptr)->QFile::close();
}

void QFile_TimerEvent(void* ptr, void* event){
	static_cast<MyQFile*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QFile_TimerEventDefault(void* ptr, void* event){
	static_cast<QFile*>(ptr)->QFile::timerEvent(static_cast<QTimerEvent*>(event));
}

void QFile_ChildEvent(void* ptr, void* event){
	static_cast<MyQFile*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QFile_ChildEventDefault(void* ptr, void* event){
	static_cast<QFile*>(ptr)->QFile::childEvent(static_cast<QChildEvent*>(event));
}

void QFile_CustomEvent(void* ptr, void* event){
	static_cast<MyQFile*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QFile_CustomEventDefault(void* ptr, void* event){
	static_cast<QFile*>(ptr)->QFile::customEvent(static_cast<QEvent*>(event));
}

class MyQFileDevice: public QFileDevice {
public:
	void close() { callbackQFileDeviceClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQFileDeviceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQFileDeviceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQFileDeviceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QFileDevice_Seek(void* ptr, long long pos){
	return static_cast<QFileDevice*>(ptr)->seek(static_cast<long long>(pos));
}

int QFileDevice_AtEnd(void* ptr){
	return static_cast<QFileDevice*>(ptr)->atEnd();
}

void QFileDevice_Close(void* ptr){
	static_cast<MyQFileDevice*>(ptr)->close();
}

void QFileDevice_CloseDefault(void* ptr){
	static_cast<QFileDevice*>(ptr)->QFileDevice::close();
}

int QFileDevice_Error(void* ptr){
	return static_cast<QFileDevice*>(ptr)->error();
}

char* QFileDevice_FileName(void* ptr){
	return static_cast<QFileDevice*>(ptr)->fileName().toUtf8().data();
}

int QFileDevice_Flush(void* ptr){
	return static_cast<QFileDevice*>(ptr)->flush();
}

int QFileDevice_Handle(void* ptr){
	return static_cast<QFileDevice*>(ptr)->handle();
}

int QFileDevice_IsSequential(void* ptr){
	return static_cast<QFileDevice*>(ptr)->isSequential();
}

int QFileDevice_Permissions(void* ptr){
	return static_cast<QFileDevice*>(ptr)->permissions();
}

long long QFileDevice_Pos(void* ptr){
	return static_cast<long long>(static_cast<QFileDevice*>(ptr)->pos());
}

long long QFileDevice_ReadData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QFileDevice*>(ptr)->readData(data, static_cast<long long>(len)));
}

long long QFileDevice_ReadLineData(void* ptr, char* data, long long maxlen){
	return static_cast<long long>(static_cast<QFileDevice*>(ptr)->readLineData(data, static_cast<long long>(maxlen)));
}

int QFileDevice_Resize(void* ptr, long long sz){
	return static_cast<QFileDevice*>(ptr)->resize(static_cast<long long>(sz));
}

int QFileDevice_SetPermissions(void* ptr, int permissions){
	return static_cast<QFileDevice*>(ptr)->setPermissions(static_cast<QFileDevice::Permission>(permissions));
}

long long QFileDevice_Size(void* ptr){
	return static_cast<long long>(static_cast<QFileDevice*>(ptr)->size());
}

void QFileDevice_UnsetError(void* ptr){
	static_cast<QFileDevice*>(ptr)->unsetError();
}

long long QFileDevice_WriteData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QFileDevice*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QFileDevice_DestroyQFileDevice(void* ptr){
	static_cast<QFileDevice*>(ptr)->~QFileDevice();
}

void QFileDevice_TimerEvent(void* ptr, void* event){
	static_cast<MyQFileDevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileDevice_TimerEventDefault(void* ptr, void* event){
	static_cast<QFileDevice*>(ptr)->QFileDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileDevice_ChildEvent(void* ptr, void* event){
	static_cast<MyQFileDevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QFileDevice_ChildEventDefault(void* ptr, void* event){
	static_cast<QFileDevice*>(ptr)->QFileDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QFileDevice_CustomEvent(void* ptr, void* event){
	static_cast<MyQFileDevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QFileDevice_CustomEventDefault(void* ptr, void* event){
	static_cast<QFileDevice*>(ptr)->QFileDevice::customEvent(static_cast<QEvent*>(event));
}

void* QFileInfo_NewQFileInfo(){
	return new QFileInfo();
}

void* QFileInfo_NewQFileInfo5(void* dir, char* file){
	return new QFileInfo(*static_cast<QDir*>(dir), QString(file));
}

void* QFileInfo_NewQFileInfo4(void* file){
	return new QFileInfo(*static_cast<QFile*>(file));
}

void* QFileInfo_NewQFileInfo6(void* fileinfo){
	return new QFileInfo(*static_cast<QFileInfo*>(fileinfo));
}

void* QFileInfo_NewQFileInfo3(char* file){
	return new QFileInfo(QString(file));
}

void* QFileInfo_AbsoluteDir(void* ptr){
	return new QDir(static_cast<QFileInfo*>(ptr)->absoluteDir());
}

char* QFileInfo_AbsoluteFilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QFileInfo_AbsolutePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->absolutePath().toUtf8().data();
}

char* QFileInfo_BaseName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->baseName().toUtf8().data();
}

char* QFileInfo_BundleName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->bundleName().toUtf8().data();
}

int QFileInfo_Caching(void* ptr){
	return static_cast<QFileInfo*>(ptr)->caching();
}

char* QFileInfo_CanonicalFilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalFilePath().toUtf8().data();
}

char* QFileInfo_CanonicalPath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalPath().toUtf8().data();
}

char* QFileInfo_CompleteBaseName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->completeBaseName().toUtf8().data();
}

char* QFileInfo_CompleteSuffix(void* ptr){
	return static_cast<QFileInfo*>(ptr)->completeSuffix().toUtf8().data();
}

void* QFileInfo_Created(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->created());
}

void* QFileInfo_Dir(void* ptr){
	return new QDir(static_cast<QFileInfo*>(ptr)->dir());
}

int QFileInfo_QFileInfo_Exists2(char* file){
	return QFileInfo::exists(QString(file));
}

int QFileInfo_Exists(void* ptr){
	return static_cast<QFileInfo*>(ptr)->exists();
}

char* QFileInfo_FileName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->fileName().toUtf8().data();
}

char* QFileInfo_FilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->filePath().toUtf8().data();
}

char* QFileInfo_Group(void* ptr){
	return static_cast<QFileInfo*>(ptr)->group().toUtf8().data();
}

int QFileInfo_IsAbsolute(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isAbsolute();
}

int QFileInfo_IsBundle(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isBundle();
}

int QFileInfo_IsDir(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isDir();
}

int QFileInfo_IsExecutable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isExecutable();
}

int QFileInfo_IsFile(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isFile();
}

int QFileInfo_IsHidden(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isHidden();
}

int QFileInfo_IsNativePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isNativePath();
}

int QFileInfo_IsReadable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isReadable();
}

int QFileInfo_IsRelative(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isRelative();
}

int QFileInfo_IsRoot(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isRoot();
}

int QFileInfo_IsSymLink(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isSymLink();
}

int QFileInfo_IsWritable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isWritable();
}

void* QFileInfo_LastModified(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->lastModified());
}

void* QFileInfo_LastRead(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->lastRead());
}

int QFileInfo_MakeAbsolute(void* ptr){
	return static_cast<QFileInfo*>(ptr)->makeAbsolute();
}

char* QFileInfo_Owner(void* ptr){
	return static_cast<QFileInfo*>(ptr)->owner().toUtf8().data();
}

char* QFileInfo_Path(void* ptr){
	return static_cast<QFileInfo*>(ptr)->path().toUtf8().data();
}

void QFileInfo_Refresh(void* ptr){
	static_cast<QFileInfo*>(ptr)->refresh();
}

void QFileInfo_SetCaching(void* ptr, int enable){
	static_cast<QFileInfo*>(ptr)->setCaching(enable != 0);
}

void QFileInfo_SetFile3(void* ptr, void* dir, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QDir*>(dir), QString(file));
}

void QFileInfo_SetFile2(void* ptr, void* file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QFile*>(file));
}

void QFileInfo_SetFile(void* ptr, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(QString(file));
}

long long QFileInfo_Size(void* ptr){
	return static_cast<long long>(static_cast<QFileInfo*>(ptr)->size());
}

char* QFileInfo_Suffix(void* ptr){
	return static_cast<QFileInfo*>(ptr)->suffix().toUtf8().data();
}

void QFileInfo_Swap(void* ptr, void* other){
	static_cast<QFileInfo*>(ptr)->swap(*static_cast<QFileInfo*>(other));
}

char* QFileInfo_SymLinkTarget(void* ptr){
	return static_cast<QFileInfo*>(ptr)->symLinkTarget().toUtf8().data();
}

void QFileInfo_DestroyQFileInfo(void* ptr){
	static_cast<QFileInfo*>(ptr)->~QFileInfo();
}

void* QFileSelector_NewQFileSelector(void* parent){
	return new QFileSelector(static_cast<QObject*>(parent));
}

char* QFileSelector_AllSelectors(void* ptr){
	return static_cast<QFileSelector*>(ptr)->allSelectors().join(",,,").toUtf8().data();
}

char* QFileSelector_ExtraSelectors(void* ptr){
	return static_cast<QFileSelector*>(ptr)->extraSelectors().join(",,,").toUtf8().data();
}

char* QFileSelector_Select(void* ptr, char* filePath){
	return static_cast<QFileSelector*>(ptr)->select(QString(filePath)).toUtf8().data();
}

void* QFileSelector_Select2(void* ptr, void* filePath){
	return new QUrl(static_cast<QFileSelector*>(ptr)->select(*static_cast<QUrl*>(filePath)));
}

void QFileSelector_SetExtraSelectors(void* ptr, char* list){
	static_cast<QFileSelector*>(ptr)->setExtraSelectors(QString(list).split(",,,", QString::SkipEmptyParts));
}

void QFileSelector_DestroyQFileSelector(void* ptr){
	static_cast<QFileSelector*>(ptr)->~QFileSelector();
}

void QFileSelector_TimerEvent(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileSelector_TimerEventDefault(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->QFileSelector::timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileSelector_ChildEvent(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QFileSelector_ChildEventDefault(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->QFileSelector::childEvent(static_cast<QChildEvent*>(event));
}

void QFileSelector_CustomEvent(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QFileSelector_CustomEventDefault(void* ptr, void* event){
	static_cast<QFileSelector*>(ptr)->QFileSelector::customEvent(static_cast<QEvent*>(event));
}

class MyQFileSystemWatcher: public QFileSystemWatcher {
public:
	void Signal_DirectoryChanged(const QString & path) { callbackQFileSystemWatcherDirectoryChanged(this, this->objectName().toUtf8().data(), path.toUtf8().data()); };
	void Signal_FileChanged(const QString & path) { callbackQFileSystemWatcherFileChanged(this, this->objectName().toUtf8().data(), path.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQFileSystemWatcherTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQFileSystemWatcherChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQFileSystemWatcherCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

char* QFileSystemWatcher_Directories(void* ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->directories().join(",,,").toUtf8().data();
}

char* QFileSystemWatcher_Files(void* ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->files().join(",,,").toUtf8().data();
}

void* QFileSystemWatcher_NewQFileSystemWatcher(void* parent){
	return new QFileSystemWatcher(static_cast<QObject*>(parent));
}

void* QFileSystemWatcher_NewQFileSystemWatcher2(char* paths, void* parent){
	return new QFileSystemWatcher(QString(paths).split(",,,", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

int QFileSystemWatcher_AddPath(void* ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->addPath(QString(path));
}

char* QFileSystemWatcher_AddPaths(void* ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->addPaths(QString(paths).split(",,,", QString::SkipEmptyParts)).join(",,,").toUtf8().data();
}

void QFileSystemWatcher_ConnectDirectoryChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_DisconnectDirectoryChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_ConnectFileChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

void QFileSystemWatcher_DisconnectFileChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

int QFileSystemWatcher_RemovePath(void* ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->removePath(QString(path));
}

char* QFileSystemWatcher_RemovePaths(void* ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->removePaths(QString(paths).split(",,,", QString::SkipEmptyParts)).join(",,,").toUtf8().data();
}

void QFileSystemWatcher_DestroyQFileSystemWatcher(void* ptr){
	static_cast<QFileSystemWatcher*>(ptr)->~QFileSystemWatcher();
}

void QFileSystemWatcher_TimerEvent(void* ptr, void* event){
	static_cast<MyQFileSystemWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileSystemWatcher_TimerEventDefault(void* ptr, void* event){
	static_cast<QFileSystemWatcher*>(ptr)->QFileSystemWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QFileSystemWatcher_ChildEvent(void* ptr, void* event){
	static_cast<MyQFileSystemWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QFileSystemWatcher_ChildEventDefault(void* ptr, void* event){
	static_cast<QFileSystemWatcher*>(ptr)->QFileSystemWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QFileSystemWatcher_CustomEvent(void* ptr, void* event){
	static_cast<MyQFileSystemWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QFileSystemWatcher_CustomEventDefault(void* ptr, void* event){
	static_cast<QFileSystemWatcher*>(ptr)->QFileSystemWatcher::customEvent(static_cast<QEvent*>(event));
}

class MyQFinalState: public QFinalState {
public:
	MyQFinalState(QState *parent) : QFinalState(parent) {};
	void onEntry(QEvent * event) { callbackQFinalStateOnEntry(this, this->objectName().toUtf8().data(), event); };
	void onExit(QEvent * event) { callbackQFinalStateOnExit(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQFinalStateTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQFinalStateChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQFinalStateCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QFinalState_NewQFinalState(void* parent){
	return new MyQFinalState(static_cast<QState*>(parent));
}

int QFinalState_Event(void* ptr, void* e){
	return static_cast<QFinalState*>(ptr)->event(static_cast<QEvent*>(e));
}

void QFinalState_OnEntry(void* ptr, void* event){
	static_cast<MyQFinalState*>(ptr)->onEntry(static_cast<QEvent*>(event));
}

void QFinalState_OnEntryDefault(void* ptr, void* event){
	static_cast<QFinalState*>(ptr)->QFinalState::onEntry(static_cast<QEvent*>(event));
}

void QFinalState_OnExit(void* ptr, void* event){
	static_cast<MyQFinalState*>(ptr)->onExit(static_cast<QEvent*>(event));
}

void QFinalState_OnExitDefault(void* ptr, void* event){
	static_cast<QFinalState*>(ptr)->QFinalState::onExit(static_cast<QEvent*>(event));
}

void QFinalState_DestroyQFinalState(void* ptr){
	static_cast<QFinalState*>(ptr)->~QFinalState();
}

void QFinalState_TimerEvent(void* ptr, void* event){
	static_cast<MyQFinalState*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QFinalState_TimerEventDefault(void* ptr, void* event){
	static_cast<QFinalState*>(ptr)->QFinalState::timerEvent(static_cast<QTimerEvent*>(event));
}

void QFinalState_ChildEvent(void* ptr, void* event){
	static_cast<MyQFinalState*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QFinalState_ChildEventDefault(void* ptr, void* event){
	static_cast<QFinalState*>(ptr)->QFinalState::childEvent(static_cast<QChildEvent*>(event));
}

void QFinalState_CustomEvent(void* ptr, void* event){
	static_cast<MyQFinalState*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QFinalState_CustomEventDefault(void* ptr, void* event){
	static_cast<QFinalState*>(ptr)->QFinalState::customEvent(static_cast<QEvent*>(event));
}

void* QFlag_NewQFlag(int value){
	return new QFlag(value);
}

void* QGenericArgument_Data(void* ptr){
	return static_cast<QGenericArgument*>(ptr)->data();
}

class MyQHistoryState: public QHistoryState {
public:
	MyQHistoryState(HistoryType type, QState *parent) : QHistoryState(type, parent) {};
	MyQHistoryState(QState *parent) : QHistoryState(parent) {};
	void Signal_DefaultStateChanged() { callbackQHistoryStateDefaultStateChanged(this, this->objectName().toUtf8().data()); };
	void Signal_HistoryTypeChanged() { callbackQHistoryStateHistoryTypeChanged(this, this->objectName().toUtf8().data()); };
	void onEntry(QEvent * event) { callbackQHistoryStateOnEntry(this, this->objectName().toUtf8().data(), event); };
	void onExit(QEvent * event) { callbackQHistoryStateOnExit(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQHistoryStateTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHistoryStateChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHistoryStateCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QHistoryState_NewQHistoryState2(int ty, void* parent){
	return new MyQHistoryState(static_cast<QHistoryState::HistoryType>(ty), static_cast<QState*>(parent));
}

void* QHistoryState_NewQHistoryState(void* parent){
	return new MyQHistoryState(static_cast<QState*>(parent));
}

void* QHistoryState_DefaultState(void* ptr){
	return static_cast<QHistoryState*>(ptr)->defaultState();
}

void QHistoryState_ConnectDefaultStateChanged(void* ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

void QHistoryState_DisconnectDefaultStateChanged(void* ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::defaultStateChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_DefaultStateChanged));;
}

int QHistoryState_Event(void* ptr, void* e){
	return static_cast<QHistoryState*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHistoryState_HistoryType(void* ptr){
	return static_cast<QHistoryState*>(ptr)->historyType();
}

void QHistoryState_ConnectHistoryTypeChanged(void* ptr){
	QObject::connect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_DisconnectHistoryTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QHistoryState*>(ptr), &QHistoryState::historyTypeChanged, static_cast<MyQHistoryState*>(ptr), static_cast<void (MyQHistoryState::*)()>(&MyQHistoryState::Signal_HistoryTypeChanged));;
}

void QHistoryState_OnEntry(void* ptr, void* event){
	static_cast<MyQHistoryState*>(ptr)->onEntry(static_cast<QEvent*>(event));
}

void QHistoryState_OnEntryDefault(void* ptr, void* event){
	static_cast<QHistoryState*>(ptr)->QHistoryState::onEntry(static_cast<QEvent*>(event));
}

void QHistoryState_OnExit(void* ptr, void* event){
	static_cast<MyQHistoryState*>(ptr)->onExit(static_cast<QEvent*>(event));
}

void QHistoryState_OnExitDefault(void* ptr, void* event){
	static_cast<QHistoryState*>(ptr)->QHistoryState::onExit(static_cast<QEvent*>(event));
}

void QHistoryState_SetDefaultState(void* ptr, void* state){
	static_cast<QHistoryState*>(ptr)->setDefaultState(static_cast<QAbstractState*>(state));
}

void QHistoryState_SetHistoryType(void* ptr, int ty){
	static_cast<QHistoryState*>(ptr)->setHistoryType(static_cast<QHistoryState::HistoryType>(ty));
}

void QHistoryState_DestroyQHistoryState(void* ptr){
	static_cast<QHistoryState*>(ptr)->~QHistoryState();
}

void QHistoryState_TimerEvent(void* ptr, void* event){
	static_cast<MyQHistoryState*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHistoryState_TimerEventDefault(void* ptr, void* event){
	static_cast<QHistoryState*>(ptr)->QHistoryState::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHistoryState_ChildEvent(void* ptr, void* event){
	static_cast<MyQHistoryState*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHistoryState_ChildEventDefault(void* ptr, void* event){
	static_cast<QHistoryState*>(ptr)->QHistoryState::childEvent(static_cast<QChildEvent*>(event));
}

void QHistoryState_CustomEvent(void* ptr, void* event){
	static_cast<MyQHistoryState*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHistoryState_CustomEventDefault(void* ptr, void* event){
	static_cast<QHistoryState*>(ptr)->QHistoryState::customEvent(static_cast<QEvent*>(event));
}

class MyQIODevice: public QIODevice {
public:
	void Signal_AboutToClose() { callbackQIODeviceAboutToClose(this, this->objectName().toUtf8().data()); };
	void Signal_BytesWritten(qint64 bytes) { callbackQIODeviceBytesWritten(this, this->objectName().toUtf8().data(), static_cast<long long>(bytes)); };
	void close() { callbackQIODeviceClose(this, this->objectName().toUtf8().data()); };
	void Signal_ReadChannelFinished() { callbackQIODeviceReadChannelFinished(this, this->objectName().toUtf8().data()); };
	void Signal_ReadyRead() { callbackQIODeviceReadyRead(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQIODeviceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIODeviceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQIODeviceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QIODevice_GetChar(void* ptr, char* c){
	return static_cast<QIODevice*>(ptr)->getChar(c);
}

int QIODevice_PutChar(void* ptr, char* c){
	return static_cast<QIODevice*>(ptr)->putChar(*c);
}

void QIODevice_ConnectAboutToClose(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::aboutToClose), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_AboutToClose));;
}

void QIODevice_DisconnectAboutToClose(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::aboutToClose), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_AboutToClose));;
}

void QIODevice_AboutToClose(void* ptr){
	static_cast<QIODevice*>(ptr)->aboutToClose();
}

int QIODevice_AtEnd(void* ptr){
	return static_cast<QIODevice*>(ptr)->atEnd();
}

long long QIODevice_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->bytesAvailable());
}

long long QIODevice_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->bytesToWrite());
}

void QIODevice_ConnectBytesWritten(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)(qint64)>(&QIODevice::bytesWritten), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)(qint64)>(&MyQIODevice::Signal_BytesWritten));;
}

void QIODevice_DisconnectBytesWritten(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)(qint64)>(&QIODevice::bytesWritten), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)(qint64)>(&MyQIODevice::Signal_BytesWritten));;
}

void QIODevice_BytesWritten(void* ptr, long long bytes){
	static_cast<QIODevice*>(ptr)->bytesWritten(static_cast<long long>(bytes));
}

int QIODevice_CanReadLine(void* ptr){
	return static_cast<QIODevice*>(ptr)->canReadLine();
}

void QIODevice_Close(void* ptr){
	static_cast<MyQIODevice*>(ptr)->close();
}

void QIODevice_CloseDefault(void* ptr){
	static_cast<QIODevice*>(ptr)->QIODevice::close();
}

char* QIODevice_ErrorString(void* ptr){
	return static_cast<QIODevice*>(ptr)->errorString().toUtf8().data();
}

int QIODevice_IsOpen(void* ptr){
	return static_cast<QIODevice*>(ptr)->isOpen();
}

int QIODevice_IsReadable(void* ptr){
	return static_cast<QIODevice*>(ptr)->isReadable();
}

int QIODevice_IsSequential(void* ptr){
	return static_cast<QIODevice*>(ptr)->isSequential();
}

int QIODevice_IsTextModeEnabled(void* ptr){
	return static_cast<QIODevice*>(ptr)->isTextModeEnabled();
}

int QIODevice_IsWritable(void* ptr){
	return static_cast<QIODevice*>(ptr)->isWritable();
}

int QIODevice_Open(void* ptr, int mode){
	return static_cast<QIODevice*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QIODevice_OpenMode(void* ptr){
	return static_cast<QIODevice*>(ptr)->openMode();
}

void* QIODevice_Peek2(void* ptr, long long maxSize){
	return new QByteArray(static_cast<QIODevice*>(ptr)->peek(static_cast<long long>(maxSize)));
}

long long QIODevice_Peek(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->peek(data, static_cast<long long>(maxSize)));
}

long long QIODevice_Pos(void* ptr){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->pos());
}

void* QIODevice_Read2(void* ptr, long long maxSize){
	return new QByteArray(static_cast<QIODevice*>(ptr)->read(static_cast<long long>(maxSize)));
}

long long QIODevice_Read(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->read(data, static_cast<long long>(maxSize)));
}

void* QIODevice_ReadAll(void* ptr){
	return new QByteArray(static_cast<QIODevice*>(ptr)->readAll());
}

void QIODevice_ConnectReadChannelFinished(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readChannelFinished), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadChannelFinished));;
}

void QIODevice_DisconnectReadChannelFinished(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readChannelFinished), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadChannelFinished));;
}

void QIODevice_ReadChannelFinished(void* ptr){
	static_cast<QIODevice*>(ptr)->readChannelFinished();
}

void* QIODevice_ReadLine2(void* ptr, long long maxSize){
	return new QByteArray(static_cast<QIODevice*>(ptr)->readLine(static_cast<long long>(maxSize)));
}

long long QIODevice_ReadLine(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->readLine(data, static_cast<long long>(maxSize)));
}

long long QIODevice_ReadLineData(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->readLineData(data, static_cast<long long>(maxSize)));
}

void QIODevice_ConnectReadyRead(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readyRead), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadyRead));;
}

void QIODevice_DisconnectReadyRead(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readyRead), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadyRead));;
}

void QIODevice_ReadyRead(void* ptr){
	static_cast<QIODevice*>(ptr)->readyRead();
}

int QIODevice_Reset(void* ptr){
	return static_cast<QIODevice*>(ptr)->reset();
}

int QIODevice_Seek(void* ptr, long long pos){
	return static_cast<QIODevice*>(ptr)->seek(static_cast<long long>(pos));
}

void QIODevice_SetTextModeEnabled(void* ptr, int enabled){
	static_cast<QIODevice*>(ptr)->setTextModeEnabled(enabled != 0);
}

long long QIODevice_Size(void* ptr){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->size());
}

void QIODevice_UngetChar(void* ptr, char* c){
	static_cast<QIODevice*>(ptr)->ungetChar(*c);
}

int QIODevice_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QIODevice*>(ptr)->waitForBytesWritten(msecs);
}

int QIODevice_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QIODevice*>(ptr)->waitForReadyRead(msecs);
}

long long QIODevice_Write3(void* ptr, void* byteArray){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->write(*static_cast<QByteArray*>(byteArray)));
}

long long QIODevice_Write2(void* ptr, char* data){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data)));
}

long long QIODevice_Write(void* ptr, char* data, long long maxSize){
	return static_cast<long long>(static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data), static_cast<long long>(maxSize)));
}

void QIODevice_DestroyQIODevice(void* ptr){
	static_cast<QIODevice*>(ptr)->~QIODevice();
}

void QIODevice_TimerEvent(void* ptr, void* event){
	static_cast<MyQIODevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIODevice_TimerEventDefault(void* ptr, void* event){
	static_cast<QIODevice*>(ptr)->QIODevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIODevice_ChildEvent(void* ptr, void* event){
	static_cast<MyQIODevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIODevice_ChildEventDefault(void* ptr, void* event){
	static_cast<QIODevice*>(ptr)->QIODevice::childEvent(static_cast<QChildEvent*>(event));
}

void QIODevice_CustomEvent(void* ptr, void* event){
	static_cast<MyQIODevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIODevice_CustomEventDefault(void* ptr, void* event){
	static_cast<QIODevice*>(ptr)->QIODevice::customEvent(static_cast<QEvent*>(event));
}

class MyQIdentityProxyModel: public QIdentityProxyModel {
public:
	MyQIdentityProxyModel(QObject *parent) : QIdentityProxyModel(parent) {};
	void setSourceModel(QAbstractItemModel * newSourceModel) { callbackQIdentityProxyModelSetSourceModel(this, this->objectName().toUtf8().data(), newSourceModel); };
	void fetchMore(const QModelIndex & parent) { callbackQIdentityProxyModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void revert() { if (!callbackQIdentityProxyModelRevert(this, this->objectName().toUtf8().data())) { QIdentityProxyModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { callbackQIdentityProxyModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQIdentityProxyModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIdentityProxyModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQIdentityProxyModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QIdentityProxyModel_NewQIdentityProxyModel(void* parent){
	return new MyQIdentityProxyModel(static_cast<QObject*>(parent));
}

int QIdentityProxyModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QIdentityProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QIdentityProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QIdentityProxyModel_Index(void* ptr, int row, int column, void* parent){
	return new QModelIndex(static_cast<QIdentityProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QIdentityProxyModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void* QIdentityProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return new QModelIndex(static_cast<QIdentityProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QIdentityProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return new QModelIndex(static_cast<QIdentityProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

void* QIdentityProxyModel_Parent(void* ptr, void* child){
	return new QModelIndex(static_cast<QIdentityProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)));
}

int QIdentityProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RowCount(void* ptr, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QIdentityProxyModel_SetSourceModel(void* ptr, void* newSourceModel){
	static_cast<MyQIdentityProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(newSourceModel));
}

void QIdentityProxyModel_SetSourceModelDefault(void* ptr, void* newSourceModel){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::setSourceModel(static_cast<QAbstractItemModel*>(newSourceModel));
}

void* QIdentityProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QIdentityProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QIdentityProxyModel_DestroyQIdentityProxyModel(void* ptr){
	static_cast<QIdentityProxyModel*>(ptr)->~QIdentityProxyModel();
}

void QIdentityProxyModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQIdentityProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QIdentityProxyModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QIdentityProxyModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQIdentityProxyModel*>(ptr), "revert");
}

void QIdentityProxyModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QIdentityProxyModel*>(ptr), "revert");
}

void QIdentityProxyModel_Sort(void* ptr, int column, int order){
	static_cast<MyQIdentityProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QIdentityProxyModel_SortDefault(void* ptr, int column, int order){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QIdentityProxyModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQIdentityProxyModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIdentityProxyModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIdentityProxyModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQIdentityProxyModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIdentityProxyModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::childEvent(static_cast<QChildEvent*>(event));
}

void QIdentityProxyModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQIdentityProxyModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIdentityProxyModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QIdentityProxyModel*>(ptr)->QIdentityProxyModel::customEvent(static_cast<QEvent*>(event));
}

void* QItemSelection_NewQItemSelection(){
	return new QItemSelection();
}

void* QItemSelection_NewQItemSelection2(void* topLeft, void* bottomRight){
	return new QItemSelection(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelection_Contains(void* ptr, void* index){
	return static_cast<QItemSelection*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

void QItemSelection_Merge(void* ptr, void* other, int command){
	static_cast<QItemSelection*>(ptr)->merge(*static_cast<QItemSelection*>(other), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QItemSelection_Select(void* ptr, void* topLeft, void* bottomRight){
	static_cast<QItemSelection*>(ptr)->select(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

void QItemSelection_QItemSelection_Split(void* ran, void* other, void* result){
	QItemSelection::split(*static_cast<QItemSelectionRange*>(ran), *static_cast<QItemSelectionRange*>(other), static_cast<QItemSelection*>(result));
}

class MyQItemSelectionModel: public QItemSelectionModel {
public:
	MyQItemSelectionModel(QAbstractItemModel *model) : QItemSelectionModel(model) {};
	MyQItemSelectionModel(QAbstractItemModel *model, QObject *parent) : QItemSelectionModel(model, parent) {};
	void clear() { if (!callbackQItemSelectionModelClear(this, this->objectName().toUtf8().data())) { QItemSelectionModel::clear(); }; };
	void clearCurrentIndex() { if (!callbackQItemSelectionModelClearCurrentIndex(this, this->objectName().toUtf8().data())) { QItemSelectionModel::clearCurrentIndex(); }; };
	void Signal_CurrentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQItemSelectionModelCurrentChanged(this, this->objectName().toUtf8().data(), new QModelIndex(current), new QModelIndex(previous)); };
	void Signal_CurrentColumnChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQItemSelectionModelCurrentColumnChanged(this, this->objectName().toUtf8().data(), new QModelIndex(current), new QModelIndex(previous)); };
	void Signal_CurrentRowChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQItemSelectionModelCurrentRowChanged(this, this->objectName().toUtf8().data(), new QModelIndex(current), new QModelIndex(previous)); };
	void Signal_ModelChanged(QAbstractItemModel * model) { callbackQItemSelectionModelModelChanged(this, this->objectName().toUtf8().data(), model); };
	void reset() { if (!callbackQItemSelectionModelReset(this, this->objectName().toUtf8().data())) { QItemSelectionModel::reset(); }; };
	void select(const QModelIndex & index, QItemSelectionModel::SelectionFlags command) { if (!callbackQItemSelectionModelSelect(this, this->objectName().toUtf8().data(), new QModelIndex(index), command)) { QItemSelectionModel::select(index, command); }; };
	void setCurrentIndex(const QModelIndex & index, QItemSelectionModel::SelectionFlags command) { if (!callbackQItemSelectionModelSetCurrentIndex(this, this->objectName().toUtf8().data(), new QModelIndex(index), command)) { QItemSelectionModel::setCurrentIndex(index, command); }; };
	void timerEvent(QTimerEvent * event) { callbackQItemSelectionModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQItemSelectionModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQItemSelectionModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QItemSelectionModel_NewQItemSelectionModel(void* model){
	return new MyQItemSelectionModel(static_cast<QAbstractItemModel*>(model));
}

void* QItemSelectionModel_NewQItemSelectionModel2(void* model, void* parent){
	return new MyQItemSelectionModel(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

void QItemSelectionModel_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQItemSelectionModel*>(ptr), "clear");
}

void QItemSelectionModel_ClearDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clear");
}

void QItemSelectionModel_ClearCurrentIndex(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQItemSelectionModel*>(ptr), "clearCurrentIndex");
}

void QItemSelectionModel_ClearCurrentIndexDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearCurrentIndex");
}

void QItemSelectionModel_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearSelection");
}

int QItemSelectionModel_ColumnIntersectsSelection(void* ptr, int column, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->columnIntersectsSelection(column, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_CurrentChanged(void* ptr, void* current, void* previous){
	static_cast<QItemSelectionModel*>(ptr)->currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QItemSelectionModel_ConnectCurrentColumnChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

void QItemSelectionModel_DisconnectCurrentColumnChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

void QItemSelectionModel_CurrentColumnChanged(void* ptr, void* current, void* previous){
	static_cast<QItemSelectionModel*>(ptr)->currentColumnChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void* QItemSelectionModel_CurrentIndex(void* ptr){
	return new QModelIndex(static_cast<QItemSelectionModel*>(ptr)->currentIndex());
}

void QItemSelectionModel_ConnectCurrentRowChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

void QItemSelectionModel_DisconnectCurrentRowChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

void QItemSelectionModel_CurrentRowChanged(void* ptr, void* current, void* previous){
	static_cast<QItemSelectionModel*>(ptr)->currentRowChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

int QItemSelectionModel_HasSelection(void* ptr){
	return static_cast<QItemSelectionModel*>(ptr)->hasSelection();
}

int QItemSelectionModel_IsColumnSelected(void* ptr, int column, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->isColumnSelected(column, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsRowSelected(void* ptr, int row, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->isRowSelected(row, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsSelected(void* ptr, void* index){
	return static_cast<QItemSelectionModel*>(ptr)->isSelected(*static_cast<QModelIndex*>(index));
}

void* QItemSelectionModel_Model2(void* ptr){
	return static_cast<QItemSelectionModel*>(ptr)->model();
}

void* QItemSelectionModel_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionModel*>(ptr)->model());
}

void QItemSelectionModel_ConnectModelChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_DisconnectModelChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_ModelChanged(void* ptr, void* model){
	static_cast<QItemSelectionModel*>(ptr)->modelChanged(static_cast<QAbstractItemModel*>(model));
}

void QItemSelectionModel_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQItemSelectionModel*>(ptr), "reset");
}

void QItemSelectionModel_ResetDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "reset");
}

int QItemSelectionModel_RowIntersectsSelection(void* ptr, int row, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->rowIntersectsSelection(row, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_Select(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<MyQItemSelectionModel*>(ptr), "select", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SelectDefault(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "select", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetCurrentIndex(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<MyQItemSelectionModel*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetCurrentIndexDefault(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetModel(void* ptr, void* model){
	static_cast<QItemSelectionModel*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QItemSelectionModel_DestroyQItemSelectionModel(void* ptr){
	static_cast<QItemSelectionModel*>(ptr)->~QItemSelectionModel();
}

void QItemSelectionModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQItemSelectionModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QItemSelectionModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QItemSelectionModel*>(ptr)->QItemSelectionModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QItemSelectionModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQItemSelectionModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QItemSelectionModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QItemSelectionModel*>(ptr)->QItemSelectionModel::childEvent(static_cast<QChildEvent*>(event));
}

void QItemSelectionModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQItemSelectionModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QItemSelectionModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QItemSelectionModel*>(ptr)->QItemSelectionModel::customEvent(static_cast<QEvent*>(event));
}

int QItemSelectionRange_Intersects(void* ptr, void* other){
	return static_cast<QItemSelectionRange*>(ptr)->intersects(*static_cast<QItemSelectionRange*>(other));
}

void* QItemSelectionRange_NewQItemSelectionRange(){
	return new QItemSelectionRange();
}

void* QItemSelectionRange_NewQItemSelectionRange2(void* other){
	return new QItemSelectionRange(*static_cast<QItemSelectionRange*>(other));
}

void* QItemSelectionRange_NewQItemSelectionRange4(void* index){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(index));
}

void* QItemSelectionRange_NewQItemSelectionRange3(void* topLeft, void* bottomRight){
	return new QItemSelectionRange(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

int QItemSelectionRange_Bottom(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->bottom();
}

int QItemSelectionRange_Contains(void* ptr, void* index){
	return static_cast<QItemSelectionRange*>(ptr)->contains(*static_cast<QModelIndex*>(index));
}

int QItemSelectionRange_Contains2(void* ptr, int row, int column, void* parentIndex){
	return static_cast<QItemSelectionRange*>(ptr)->contains(row, column, *static_cast<QModelIndex*>(parentIndex));
}

int QItemSelectionRange_Height(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->height();
}

int QItemSelectionRange_IsEmpty(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isEmpty();
}

int QItemSelectionRange_IsValid(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->isValid();
}

int QItemSelectionRange_Left(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->left();
}

void* QItemSelectionRange_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionRange*>(ptr)->model());
}

void* QItemSelectionRange_Parent(void* ptr){
	return new QModelIndex(static_cast<QItemSelectionRange*>(ptr)->parent());
}

int QItemSelectionRange_Right(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->right();
}

int QItemSelectionRange_Top(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->top();
}

int QItemSelectionRange_Width(void* ptr){
	return static_cast<QItemSelectionRange*>(ptr)->width();
}

void QJsonArray_Append(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->append(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Contains(void* ptr, void* value){
	return static_cast<QJsonArray*>(ptr)->contains(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Count(void* ptr){
	return static_cast<QJsonArray*>(ptr)->count();
}

int QJsonArray_Empty(void* ptr){
	return static_cast<QJsonArray*>(ptr)->empty();
}

void* QJsonArray_QJsonArray_FromStringList(char* list){
	return new QJsonArray(QJsonArray::fromStringList(QString(list).split(",,,", QString::SkipEmptyParts)));
}

int QJsonArray_IsEmpty(void* ptr){
	return static_cast<QJsonArray*>(ptr)->isEmpty();
}

void QJsonArray_Pop_back(void* ptr){
	static_cast<QJsonArray*>(ptr)->pop_back();
}

void QJsonArray_Pop_front(void* ptr){
	static_cast<QJsonArray*>(ptr)->pop_front();
}

void QJsonArray_Prepend(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->prepend(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_back(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->push_back(*static_cast<QJsonValue*>(value));
}

void QJsonArray_Push_front(void* ptr, void* value){
	static_cast<QJsonArray*>(ptr)->push_front(*static_cast<QJsonValue*>(value));
}

void QJsonArray_RemoveAt(void* ptr, int i){
	static_cast<QJsonArray*>(ptr)->removeAt(i);
}

void QJsonArray_RemoveFirst(void* ptr){
	static_cast<QJsonArray*>(ptr)->removeFirst();
}

void QJsonArray_RemoveLast(void* ptr){
	static_cast<QJsonArray*>(ptr)->removeLast();
}

int QJsonArray_Size(void* ptr){
	return static_cast<QJsonArray*>(ptr)->size();
}

void QJsonArray_DestroyQJsonArray(void* ptr){
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

void* QJsonDocument_NewQJsonDocument(){
	return new QJsonDocument();
}

void* QJsonDocument_NewQJsonDocument3(void* array){
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

void* QJsonDocument_NewQJsonDocument4(void* other){
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_NewQJsonDocument2(void* object){
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_Array(void* ptr){
	return new QJsonArray(static_cast<QJsonDocument*>(ptr)->array());
}

void* QJsonDocument_QJsonDocument_FromBinaryData(void* data, int validation){
	return new QJsonDocument(QJsonDocument::fromBinaryData(*static_cast<QByteArray*>(data), static_cast<QJsonDocument::DataValidation>(validation)));
}

void* QJsonDocument_QJsonDocument_FromJson(void* json, void* error){
	return new QJsonDocument(QJsonDocument::fromJson(*static_cast<QByteArray*>(json), static_cast<QJsonParseError*>(error)));
}

void* QJsonDocument_QJsonDocument_FromRawData(char* data, int size, int validation){
	return new QJsonDocument(QJsonDocument::fromRawData(const_cast<const char*>(data), size, static_cast<QJsonDocument::DataValidation>(validation)));
}

void* QJsonDocument_QJsonDocument_FromVariant(void* variant){
	return new QJsonDocument(QJsonDocument::fromVariant(*static_cast<QVariant*>(variant)));
}

int QJsonDocument_IsArray(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isArray();
}

int QJsonDocument_IsEmpty(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isEmpty();
}

int QJsonDocument_IsNull(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isNull();
}

int QJsonDocument_IsObject(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isObject();
}

void* QJsonDocument_Object(void* ptr){
	return new QJsonObject(static_cast<QJsonDocument*>(ptr)->object());
}

void QJsonDocument_SetArray(void* ptr, void* array){
	static_cast<QJsonDocument*>(ptr)->setArray(*static_cast<QJsonArray*>(array));
}

void QJsonDocument_SetObject(void* ptr, void* object){
	static_cast<QJsonDocument*>(ptr)->setObject(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_ToBinaryData(void* ptr){
	return new QByteArray(static_cast<QJsonDocument*>(ptr)->toBinaryData());
}

void* QJsonDocument_ToJson(void* ptr, int format){
	return new QByteArray(static_cast<QJsonDocument*>(ptr)->toJson(static_cast<QJsonDocument::JsonFormat>(format)));
}

void* QJsonDocument_ToVariant(void* ptr){
	return new QVariant(static_cast<QJsonDocument*>(ptr)->toVariant());
}

void QJsonDocument_DestroyQJsonDocument(void* ptr){
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

int QJsonObject_Contains(void* ptr, char* key){
	return static_cast<QJsonObject*>(ptr)->contains(QString(key));
}

int QJsonObject_Count(void* ptr){
	return static_cast<QJsonObject*>(ptr)->count();
}

int QJsonObject_Empty(void* ptr){
	return static_cast<QJsonObject*>(ptr)->empty();
}

int QJsonObject_IsEmpty(void* ptr){
	return static_cast<QJsonObject*>(ptr)->isEmpty();
}

char* QJsonObject_Keys(void* ptr){
	return static_cast<QJsonObject*>(ptr)->keys().join(",,,").toUtf8().data();
}

int QJsonObject_Length(void* ptr){
	return static_cast<QJsonObject*>(ptr)->length();
}

int QJsonObject_Size(void* ptr){
	return static_cast<QJsonObject*>(ptr)->size();
}

void QJsonObject_DestroyQJsonObject(void* ptr){
	static_cast<QJsonObject*>(ptr)->~QJsonObject();
}

char* QJsonParseError_ErrorString(void* ptr){
	return static_cast<QJsonParseError*>(ptr)->errorString().toUtf8().data();
}

void* QJsonValue_NewQJsonValue5(void* s){
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

void* QJsonValue_NewQJsonValue(int ty){
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

void* QJsonValue_NewQJsonValue2(int b){
	return new QJsonValue(b != 0);
}

void* QJsonValue_NewQJsonValue7(void* a){
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

void* QJsonValue_NewQJsonValue8(void* o){
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

void* QJsonValue_NewQJsonValue9(void* other){
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

void* QJsonValue_NewQJsonValue4(char* s){
	return new QJsonValue(QString(s));
}

void* QJsonValue_NewQJsonValue6(char* s){
	return new QJsonValue(const_cast<const char*>(s));
}

void* QJsonValue_NewQJsonValue12(int n){
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue13(long long n){
	return new QJsonValue(static_cast<long long>(n));
}

int QJsonValue_IsArray(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isArray();
}

int QJsonValue_IsBool(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isBool();
}

int QJsonValue_IsDouble(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isDouble();
}

int QJsonValue_IsNull(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isNull();
}

int QJsonValue_IsObject(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isObject();
}

int QJsonValue_IsString(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isString();
}

int QJsonValue_IsUndefined(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isUndefined();
}

void* QJsonValue_ToArray2(void* ptr){
	return new QJsonArray(static_cast<QJsonValue*>(ptr)->toArray());
}

void* QJsonValue_ToArray(void* ptr, void* defaultValue){
	return new QJsonArray(static_cast<QJsonValue*>(ptr)->toArray(*static_cast<QJsonArray*>(defaultValue)));
}

int QJsonValue_ToBool(void* ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toBool(defaultValue != 0);
}

int QJsonValue_ToInt(void* ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toInt(defaultValue);
}

void* QJsonValue_ToObject2(void* ptr){
	return new QJsonObject(static_cast<QJsonValue*>(ptr)->toObject());
}

void* QJsonValue_ToObject(void* ptr, void* defaultValue){
	return new QJsonObject(static_cast<QJsonValue*>(ptr)->toObject(*static_cast<QJsonObject*>(defaultValue)));
}

char* QJsonValue_ToString(void* ptr, char* defaultValue){
	return static_cast<QJsonValue*>(ptr)->toString(QString(defaultValue)).toUtf8().data();
}

void* QJsonValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QJsonValue*>(ptr)->toVariant());
}

int QJsonValue_Type(void* ptr){
	return static_cast<QJsonValue*>(ptr)->type();
}

void QJsonValue_DestroyQJsonValue(void* ptr){
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

void* QLatin1Char_NewQLatin1Char(char* c){
	return new QLatin1Char(*c);
}

void* QLatin1String_NewQLatin1String3(void* str){
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

void* QLatin1String_NewQLatin1String(char* str){
	return new QLatin1String(const_cast<const char*>(str));
}

void* QLatin1String_NewQLatin1String2(char* str, int size){
	return new QLatin1String(const_cast<const char*>(str), size);
}

int QLatin1String_Size(void* ptr){
	return static_cast<QLatin1String*>(ptr)->size();
}

char* QLibrary_FileName(void* ptr){
	return static_cast<QLibrary*>(ptr)->fileName().toUtf8().data();
}

int QLibrary_LoadHints(void* ptr){
	return static_cast<QLibrary*>(ptr)->loadHints();
}

void QLibrary_SetFileName(void* ptr, char* fileName){
	static_cast<QLibrary*>(ptr)->setFileName(QString(fileName));
}

void QLibrary_SetFileNameAndVersion(void* ptr, char* fileName, int versionNumber){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), versionNumber);
}

void QLibrary_SetLoadHints(void* ptr, int hints){
	static_cast<QLibrary*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(hints));
}

void* QLibrary_NewQLibrary(void* parent){
	return new QLibrary(static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary2(char* fileName, void* parent){
	return new QLibrary(QString(fileName), static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary4(char* fileName, char* version, void* parent){
	return new QLibrary(QString(fileName), QString(version), static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary3(char* fileName, int verNum, void* parent){
	return new QLibrary(QString(fileName), verNum, static_cast<QObject*>(parent));
}

char* QLibrary_ErrorString(void* ptr){
	return static_cast<QLibrary*>(ptr)->errorString().toUtf8().data();
}

int QLibrary_QLibrary_IsLibrary(char* fileName){
	return QLibrary::isLibrary(QString(fileName));
}

int QLibrary_IsLoaded(void* ptr){
	return static_cast<QLibrary*>(ptr)->isLoaded();
}

int QLibrary_Load(void* ptr){
	return static_cast<QLibrary*>(ptr)->load();
}

void QLibrary_SetFileNameAndVersion2(void* ptr, char* fileName, char* version){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), QString(version));
}

int QLibrary_Unload(void* ptr){
	return static_cast<QLibrary*>(ptr)->unload();
}

void QLibrary_DestroyQLibrary(void* ptr){
	static_cast<QLibrary*>(ptr)->~QLibrary();
}

void QLibrary_TimerEvent(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLibrary_TimerEventDefault(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->QLibrary::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLibrary_ChildEvent(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLibrary_ChildEventDefault(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->QLibrary::childEvent(static_cast<QChildEvent*>(event));
}

void QLibrary_CustomEvent(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLibrary_CustomEventDefault(void* ptr, void* event){
	static_cast<QLibrary*>(ptr)->QLibrary::customEvent(static_cast<QEvent*>(event));
}

int QLibraryInfo_QLibraryInfo_IsDebugBuild(){
	return QLibraryInfo::isDebugBuild();
}

char* QLibraryInfo_QLibraryInfo_LicensedProducts(){
	return QLibraryInfo::licensedProducts().toUtf8().data();
}

char* QLibraryInfo_QLibraryInfo_Licensee(){
	return QLibraryInfo::licensee().toUtf8().data();
}

char* QLibraryInfo_QLibraryInfo_Location(int loc){
	return QLibraryInfo::location(static_cast<QLibraryInfo::LibraryLocation>(loc)).toUtf8().data();
}

void* QLine_NewQLine(){
	return new QLine();
}

void* QLine_NewQLine2(void* p1, void* p2){
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void* QLine_NewQLine3(int x1, int y1, int x2, int y2){
	return new QLine(x1, y1, x2, y2);
}

int QLine_Dx(void* ptr){
	return static_cast<QLine*>(ptr)->dx();
}

int QLine_Dy(void* ptr){
	return static_cast<QLine*>(ptr)->dy();
}

int QLine_IsNull(void* ptr){
	return static_cast<QLine*>(ptr)->isNull();
}

void* QLine_P1(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QLine*>(ptr)->p1()).x(), static_cast<QPoint>(static_cast<QLine*>(ptr)->p1()).y());
}

void* QLine_P2(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QLine*>(ptr)->p2()).x(), static_cast<QPoint>(static_cast<QLine*>(ptr)->p2()).y());
}

void QLine_SetLine(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QLine*>(ptr)->setLine(x1, y1, x2, y2);
}

void QLine_SetP1(void* ptr, void* p1){
	static_cast<QLine*>(ptr)->setP1(*static_cast<QPoint*>(p1));
}

void QLine_SetP2(void* ptr, void* p2){
	static_cast<QLine*>(ptr)->setP2(*static_cast<QPoint*>(p2));
}

void QLine_SetPoints(void* ptr, void* p1, void* p2){
	static_cast<QLine*>(ptr)->setPoints(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QLine_Translate(void* ptr, void* offset){
	static_cast<QLine*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QLine_Translate2(void* ptr, int dx, int dy){
	static_cast<QLine*>(ptr)->translate(dx, dy);
}

int QLine_X1(void* ptr){
	return static_cast<QLine*>(ptr)->x1();
}

int QLine_X2(void* ptr){
	return static_cast<QLine*>(ptr)->x2();
}

int QLine_Y1(void* ptr){
	return static_cast<QLine*>(ptr)->y1();
}

int QLine_Y2(void* ptr){
	return static_cast<QLine*>(ptr)->y2();
}

double QLineF_AngleTo(void* ptr, void* line){
	return static_cast<double>(static_cast<QLineF*>(ptr)->angleTo(*static_cast<QLineF*>(line)));
}

int QLineF_Intersect(void* ptr, void* line, void* intersectionPoint){
	return static_cast<QLineF*>(ptr)->intersect(*static_cast<QLineF*>(line), static_cast<QPointF*>(intersectionPoint));
}

void* QLineF_NewQLineF(){
	return new QLineF();
}

void* QLineF_NewQLineF4(void* line){
	return new QLineF(*static_cast<QLine*>(line));
}

void* QLineF_NewQLineF2(void* p1, void* p2){
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2){
	return new QLineF(static_cast<double>(x1), static_cast<double>(y1), static_cast<double>(x2), static_cast<double>(y2));
}

double QLineF_Angle(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->angle());
}

double QLineF_Dx(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->dx());
}

double QLineF_Dy(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->dy());
}

int QLineF_IsNull(void* ptr){
	return static_cast<QLineF*>(ptr)->isNull();
}

double QLineF_Length(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->length());
}

void QLineF_SetAngle(void* ptr, double angle){
	static_cast<QLineF*>(ptr)->setAngle(static_cast<double>(angle));
}

void QLineF_SetLength(void* ptr, double length){
	static_cast<QLineF*>(ptr)->setLength(static_cast<double>(length));
}

void QLineF_SetLine(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QLineF*>(ptr)->setLine(static_cast<double>(x1), static_cast<double>(y1), static_cast<double>(x2), static_cast<double>(y2));
}

void QLineF_SetP1(void* ptr, void* p1){
	static_cast<QLineF*>(ptr)->setP1(*static_cast<QPointF*>(p1));
}

void QLineF_SetP2(void* ptr, void* p2){
	static_cast<QLineF*>(ptr)->setP2(*static_cast<QPointF*>(p2));
}

void QLineF_SetPoints(void* ptr, void* p1, void* p2){
	static_cast<QLineF*>(ptr)->setPoints(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QLineF_Translate(void* ptr, void* offset){
	static_cast<QLineF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QLineF_Translate2(void* ptr, double dx, double dy){
	static_cast<QLineF*>(ptr)->translate(static_cast<double>(dx), static_cast<double>(dy));
}

double QLineF_X1(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->x1());
}

double QLineF_X2(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->x2());
}

double QLineF_Y1(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->y1());
}

double QLineF_Y2(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->y2());
}

void* QLocale_NewQLocale(){
	return new QLocale();
}

void* QLocale_NewQLocale3(int language, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale4(int language, int script, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(script), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale5(void* other){
	return new QLocale(*static_cast<QLocale*>(other));
}

void* QLocale_NewQLocale2(char* name){
	return new QLocale(QString(name));
}

char* QLocale_AmText(void* ptr){
	return static_cast<QLocale*>(ptr)->amText().toUtf8().data();
}

char* QLocale_Bcp47Name(void* ptr){
	return static_cast<QLocale*>(ptr)->bcp47Name().toUtf8().data();
}

int QLocale_Country(void* ptr){
	return static_cast<QLocale*>(ptr)->country();
}

char* QLocale_QLocale_CountryToString(int country){
	return QLocale::countryToString(static_cast<QLocale::Country>(country)).toUtf8().data();
}

char* QLocale_CreateSeparatedList(void* ptr, char* list){
	return static_cast<QLocale*>(ptr)->createSeparatedList(QString(list).split(",,,", QString::SkipEmptyParts)).toUtf8().data();
}

char* QLocale_CurrencySymbol(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->currencySymbol(static_cast<QLocale::CurrencySymbolFormat>(format)).toUtf8().data();
}

char* QLocale_DateFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->dateFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DateTimeFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->dateTimeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DayName(void* ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->dayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_FirstDayOfWeek(void* ptr){
	return static_cast<QLocale*>(ptr)->firstDayOfWeek();
}

int QLocale_Language(void* ptr){
	return static_cast<QLocale*>(ptr)->language();
}

char* QLocale_QLocale_LanguageToString(int language){
	return QLocale::languageToString(static_cast<QLocale::Language>(language)).toUtf8().data();
}

int QLocale_MeasurementSystem(void* ptr){
	return static_cast<QLocale*>(ptr)->measurementSystem();
}

char* QLocale_MonthName(void* ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->monthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_Name(void* ptr){
	return static_cast<QLocale*>(ptr)->name().toUtf8().data();
}

char* QLocale_NativeCountryName(void* ptr){
	return static_cast<QLocale*>(ptr)->nativeCountryName().toUtf8().data();
}

char* QLocale_NativeLanguageName(void* ptr){
	return static_cast<QLocale*>(ptr)->nativeLanguageName().toUtf8().data();
}

int QLocale_NumberOptions(void* ptr){
	return static_cast<QLocale*>(ptr)->numberOptions();
}

char* QLocale_PmText(void* ptr){
	return static_cast<QLocale*>(ptr)->pmText().toUtf8().data();
}

char* QLocale_QuoteString(void* ptr, char* str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(QString(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

char* QLocale_QuoteString2(void* ptr, void* str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(*static_cast<QStringRef*>(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

int QLocale_Script(void* ptr){
	return static_cast<QLocale*>(ptr)->script();
}

char* QLocale_QLocale_ScriptToString(int script){
	return QLocale::scriptToString(static_cast<QLocale::Script>(script)).toUtf8().data();
}

void QLocale_QLocale_SetDefault(void* locale){
	QLocale::setDefault(*static_cast<QLocale*>(locale));
}

void QLocale_SetNumberOptions(void* ptr, int options){
	static_cast<QLocale*>(ptr)->setNumberOptions(static_cast<QLocale::NumberOption>(options));
}

char* QLocale_StandaloneDayName(void* ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->standaloneDayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_StandaloneMonthName(void* ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->standaloneMonthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_TextDirection(void* ptr){
	return static_cast<QLocale*>(ptr)->textDirection();
}

char* QLocale_TimeFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->timeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToCurrencyString6(void* ptr, int value, char* symbol){
	return static_cast<QLocale*>(ptr)->toCurrencyString(value, QString(symbol)).toUtf8().data();
}

void* QLocale_ToDateTime(void* ptr, char* stri, int format){
	return new QDateTime(static_cast<QLocale*>(ptr)->toDateTime(QString(stri), static_cast<QLocale::FormatType>(format)));
}

void* QLocale_ToDateTime2(void* ptr, char* stri, char* format){
	return new QDateTime(static_cast<QLocale*>(ptr)->toDateTime(QString(stri), QString(format)));
}

int QLocale_ToInt(void* ptr, char* s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(QString(s), NULL);
}

int QLocale_ToInt2(void* ptr, void* s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringRef*>(s), NULL);
}

char* QLocale_ToLower(void* ptr, char* str){
	return static_cast<QLocale*>(ptr)->toLower(QString(str)).toUtf8().data();
}

char* QLocale_ToString3(void* ptr, void* date, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString2(void* ptr, void* date, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), QString(format)).toUtf8().data();
}

char* QLocale_ToString6(void* ptr, void* dateTime, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString7(void* ptr, void* dateTime, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), QString(format)).toUtf8().data();
}

char* QLocale_ToString5(void* ptr, void* time, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString4(void* ptr, void* time, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), QString(format)).toUtf8().data();
}

char* QLocale_ToString12(void* ptr, int i){
	return static_cast<QLocale*>(ptr)->toString(i).toUtf8().data();
}

char* QLocale_ToUpper(void* ptr, char* str){
	return static_cast<QLocale*>(ptr)->toUpper(QString(str)).toUtf8().data();
}

char* QLocale_UiLanguages(void* ptr){
	return static_cast<QLocale*>(ptr)->uiLanguages().join(",,,").toUtf8().data();
}

void QLocale_DestroyQLocale(void* ptr){
	static_cast<QLocale*>(ptr)->~QLocale();
}

void* QLockFile_NewQLockFile(char* fileName){
	return new QLockFile(QString(fileName));
}

int QLockFile_Error(void* ptr){
	return static_cast<QLockFile*>(ptr)->error();
}

int QLockFile_IsLocked(void* ptr){
	return static_cast<QLockFile*>(ptr)->isLocked();
}

int QLockFile_Lock(void* ptr){
	return static_cast<QLockFile*>(ptr)->lock();
}

int QLockFile_RemoveStaleLockFile(void* ptr){
	return static_cast<QLockFile*>(ptr)->removeStaleLockFile();
}

void QLockFile_SetStaleLockTime(void* ptr, int staleLockTime){
	static_cast<QLockFile*>(ptr)->setStaleLockTime(staleLockTime);
}

int QLockFile_StaleLockTime(void* ptr){
	return static_cast<QLockFile*>(ptr)->staleLockTime();
}

int QLockFile_TryLock(void* ptr, int timeout){
	return static_cast<QLockFile*>(ptr)->tryLock(timeout);
}

void QLockFile_DestroyQLockFile(void* ptr){
	static_cast<QLockFile*>(ptr)->~QLockFile();
}

void QLockFile_Unlock(void* ptr){
	static_cast<QLockFile*>(ptr)->unlock();
}

void* QLoggingCategory_NewQLoggingCategory(char* category){
	return new QLoggingCategory(const_cast<const char*>(category));
}

void* QLoggingCategory_QLoggingCategory_DefaultCategory(){
	return QLoggingCategory::defaultCategory();
}

int QLoggingCategory_IsCriticalEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isCriticalEnabled();
}

int QLoggingCategory_IsDebugEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isDebugEnabled();
}

int QLoggingCategory_IsInfoEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isInfoEnabled();
}

int QLoggingCategory_IsWarningEnabled(void* ptr){
	return static_cast<QLoggingCategory*>(ptr)->isWarningEnabled();
}

void QLoggingCategory_QLoggingCategory_SetFilterRules(char* rules){
	QLoggingCategory::setFilterRules(QString(rules));
}

void QLoggingCategory_DestroyQLoggingCategory(void* ptr){
	static_cast<QLoggingCategory*>(ptr)->~QLoggingCategory();
}

void* QMargins_NewQMargins(){
	return new QMargins();
}

void* QMargins_NewQMargins2(int left, int top, int right, int bottom){
	return new QMargins(left, top, right, bottom);
}

int QMargins_Bottom(void* ptr){
	return static_cast<QMargins*>(ptr)->bottom();
}

int QMargins_IsNull(void* ptr){
	return static_cast<QMargins*>(ptr)->isNull();
}

int QMargins_Left(void* ptr){
	return static_cast<QMargins*>(ptr)->left();
}

int QMargins_Right(void* ptr){
	return static_cast<QMargins*>(ptr)->right();
}

void QMargins_SetBottom(void* ptr, int bottom){
	static_cast<QMargins*>(ptr)->setBottom(bottom);
}

void QMargins_SetLeft(void* ptr, int left){
	static_cast<QMargins*>(ptr)->setLeft(left);
}

void QMargins_SetRight(void* ptr, int right){
	static_cast<QMargins*>(ptr)->setRight(right);
}

void QMargins_SetTop(void* ptr, int Top){
	static_cast<QMargins*>(ptr)->setTop(Top);
}

int QMargins_Top(void* ptr){
	return static_cast<QMargins*>(ptr)->top();
}

void* QMarginsF_NewQMarginsF(){
	return new QMarginsF();
}

void* QMarginsF_NewQMarginsF3(void* margins){
	return new QMarginsF(*static_cast<QMargins*>(margins));
}

void* QMarginsF_NewQMarginsF2(double left, double top, double right, double bottom){
	return new QMarginsF(static_cast<double>(left), static_cast<double>(top), static_cast<double>(right), static_cast<double>(bottom));
}

double QMarginsF_Bottom(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->bottom());
}

int QMarginsF_IsNull(void* ptr){
	return static_cast<QMarginsF*>(ptr)->isNull();
}

double QMarginsF_Left(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->left());
}

double QMarginsF_Right(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->right());
}

void QMarginsF_SetBottom(void* ptr, double bottom){
	static_cast<QMarginsF*>(ptr)->setBottom(static_cast<double>(bottom));
}

void QMarginsF_SetLeft(void* ptr, double left){
	static_cast<QMarginsF*>(ptr)->setLeft(static_cast<double>(left));
}

void QMarginsF_SetRight(void* ptr, double right){
	static_cast<QMarginsF*>(ptr)->setRight(static_cast<double>(right));
}

void QMarginsF_SetTop(void* ptr, double Top){
	static_cast<QMarginsF*>(ptr)->setTop(static_cast<double>(Top));
}

double QMarginsF_Top(void* ptr){
	return static_cast<double>(static_cast<QMarginsF*>(ptr)->top());
}

void* QMessageAuthenticationCode_NewQMessageAuthenticationCode(int method, void* key){
	return new QMessageAuthenticationCode(static_cast<QCryptographicHash::Algorithm>(method), *static_cast<QByteArray*>(key));
}

int QMessageAuthenticationCode_AddData2(void* ptr, void* device){
	return static_cast<QMessageAuthenticationCode*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QMessageAuthenticationCode_AddData3(void* ptr, void* data){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QMessageAuthenticationCode_AddData(void* ptr, char* data, int length){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(const_cast<const char*>(data), length);
}

void* QMessageAuthenticationCode_QMessageAuthenticationCode_Hash(void* message, void* key, int method){
	return new QByteArray(QMessageAuthenticationCode::hash(*static_cast<QByteArray*>(message), *static_cast<QByteArray*>(key), static_cast<QCryptographicHash::Algorithm>(method)));
}

void QMessageAuthenticationCode_Reset(void* ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->reset();
}

void* QMessageAuthenticationCode_Result(void* ptr){
	return new QByteArray(static_cast<QMessageAuthenticationCode*>(ptr)->result());
}

void QMessageAuthenticationCode_SetKey(void* ptr, void* key){
	static_cast<QMessageAuthenticationCode*>(ptr)->setKey(*static_cast<QByteArray*>(key));
}

void QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(void* ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->~QMessageAuthenticationCode();
}

void* QMessageLogger_NewQMessageLogger(){
	return new QMessageLogger();
}

void* QMessageLogger_NewQMessageLogger2(char* file, int line, char* function){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function));
}

void* QMessageLogger_NewQMessageLogger3(char* file, int line, char* function, char* category){
	return new QMessageLogger(const_cast<const char*>(file), line, const_cast<const char*>(function), const_cast<const char*>(category));
}

int QMetaEnum_IsFlag(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->isFlag();
}

int QMetaEnum_IsValid(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->isValid();
}

int QMetaEnum_KeyCount(void* ptr){
	return static_cast<QMetaEnum*>(ptr)->keyCount();
}

int QMetaEnum_KeyToValue(void* ptr, char* key, int ok){
	return static_cast<QMetaEnum*>(ptr)->keyToValue(const_cast<const char*>(key), NULL);
}

int QMetaEnum_KeysToValue(void* ptr, char* keys, int ok){
	return static_cast<QMetaEnum*>(ptr)->keysToValue(const_cast<const char*>(keys), NULL);
}

int QMetaEnum_Value(void* ptr, int index){
	return static_cast<QMetaEnum*>(ptr)->value(index);
}

void* QMetaEnum_ValueToKeys(void* ptr, int value){
	return new QByteArray(static_cast<QMetaEnum*>(ptr)->valueToKeys(value));
}

int QMetaMethod_Access(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->access();
}

int QMetaMethod_Invoke4(void* ptr, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke2(void* ptr, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke3(void* ptr, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_Invoke(void* ptr, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaMethod*>(ptr)->invoke(static_cast<QObject*>(object), static_cast<Qt::ConnectionType>(connectionType), *static_cast<QGenericReturnArgument*>(returnValue), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaMethod_IsValid(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->isValid();
}

int QMetaMethod_MethodIndex(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->methodIndex();
}

void* QMetaMethod_MethodSignature(void* ptr){
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->methodSignature());
}

int QMetaMethod_MethodType(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->methodType();
}

void* QMetaMethod_Name(void* ptr){
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->name());
}

int QMetaMethod_ParameterCount(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->parameterCount();
}

int QMetaMethod_ParameterType(void* ptr, int index){
	return static_cast<QMetaMethod*>(ptr)->parameterType(index);
}

int QMetaMethod_ReturnType(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->returnType();
}

int QMetaMethod_Revision(void* ptr){
	return static_cast<QMetaMethod*>(ptr)->revision();
}

void QMetaObject_QMetaObject_ConnectSlotsByName(void* object){
	QMetaObject::connectSlotsByName(static_cast<QObject*>(object));
}

int QMetaObject_QMetaObject_CheckConnectArgs2(void* signal, void* method){
	return QMetaObject::checkConnectArgs(*static_cast<QMetaMethod*>(signal), *static_cast<QMetaMethod*>(method));
}

int QMetaObject_QMetaObject_CheckConnectArgs(char* signal, char* method){
	return QMetaObject::checkConnectArgs(const_cast<const char*>(signal), const_cast<const char*>(method));
}

int QMetaObject_ClassInfoCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoCount();
}

int QMetaObject_ClassInfoOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->classInfoOffset();
}

int QMetaObject_ConstructorCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->constructorCount();
}

int QMetaObject_EnumeratorCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorCount();
}

int QMetaObject_EnumeratorOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->enumeratorOffset();
}

int QMetaObject_IndexOfClassInfo(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfClassInfo(const_cast<const char*>(name));
}

int QMetaObject_IndexOfConstructor(void* ptr, char* constructor){
	return static_cast<QMetaObject*>(ptr)->indexOfConstructor(const_cast<const char*>(constructor));
}

int QMetaObject_IndexOfEnumerator(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfEnumerator(const_cast<const char*>(name));
}

int QMetaObject_IndexOfMethod(void* ptr, char* method){
	return static_cast<QMetaObject*>(ptr)->indexOfMethod(const_cast<const char*>(method));
}

int QMetaObject_IndexOfProperty(void* ptr, char* name){
	return static_cast<QMetaObject*>(ptr)->indexOfProperty(const_cast<const char*>(name));
}

int QMetaObject_IndexOfSignal(void* ptr, char* signal){
	return static_cast<QMetaObject*>(ptr)->indexOfSignal(const_cast<const char*>(signal));
}

int QMetaObject_IndexOfSlot(void* ptr, char* slot){
	return static_cast<QMetaObject*>(ptr)->indexOfSlot(const_cast<const char*>(slot));
}

int QMetaObject_QMetaObject_InvokeMethod4(void* obj, char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod2(void* obj, char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod3(void* obj, char* member, int ty, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_QMetaObject_InvokeMethod(void* obj, char* member, int ty, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return QMetaObject::invokeMethod(static_cast<QObject*>(obj), const_cast<const char*>(member), static_cast<Qt::ConnectionType>(ty), *static_cast<QGenericReturnArgument*>(ret), *static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

int QMetaObject_MethodCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->methodCount();
}

int QMetaObject_MethodOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->methodOffset();
}

void* QMetaObject_NewInstance(void* ptr, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9){
	return static_cast<QMetaObject*>(ptr)->newInstance(*static_cast<QGenericArgument*>(val0), *static_cast<QGenericArgument*>(val1), *static_cast<QGenericArgument*>(val2), *static_cast<QGenericArgument*>(val3), *static_cast<QGenericArgument*>(val4), *static_cast<QGenericArgument*>(val5), *static_cast<QGenericArgument*>(val6), *static_cast<QGenericArgument*>(val7), *static_cast<QGenericArgument*>(val8), *static_cast<QGenericArgument*>(val9));
}

void* QMetaObject_QMetaObject_NormalizedSignature(char* method){
	return new QByteArray(QMetaObject::normalizedSignature(const_cast<const char*>(method)));
}

void* QMetaObject_QMetaObject_NormalizedType(char* ty){
	return new QByteArray(QMetaObject::normalizedType(const_cast<const char*>(ty)));
}

int QMetaObject_PropertyCount(void* ptr){
	return static_cast<QMetaObject*>(ptr)->propertyCount();
}

int QMetaObject_PropertyOffset(void* ptr){
	return static_cast<QMetaObject*>(ptr)->propertyOffset();
}

void* QMetaObject_SuperClass(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QMetaObject*>(ptr)->superClass());
}

int QMetaProperty_HasNotifySignal(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->hasNotifySignal();
}

int QMetaProperty_IsConstant(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isConstant();
}

int QMetaProperty_IsDesignable(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isDesignable(static_cast<QObject*>(object));
}

int QMetaProperty_IsEnumType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isEnumType();
}

int QMetaProperty_IsFinal(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isFinal();
}

int QMetaProperty_IsFlagType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isFlagType();
}

int QMetaProperty_IsReadable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isReadable();
}

int QMetaProperty_IsResettable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isResettable();
}

int QMetaProperty_IsScriptable(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isScriptable(static_cast<QObject*>(object));
}

int QMetaProperty_IsStored(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isStored(static_cast<QObject*>(object));
}

int QMetaProperty_IsUser(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->isUser(static_cast<QObject*>(object));
}

int QMetaProperty_IsValid(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isValid();
}

int QMetaProperty_IsWritable(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->isWritable();
}

int QMetaProperty_NotifySignalIndex(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->notifySignalIndex();
}

int QMetaProperty_PropertyIndex(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->propertyIndex();
}

void* QMetaProperty_Read(void* ptr, void* object){
	return new QVariant(static_cast<QMetaProperty*>(ptr)->read(static_cast<QObject*>(object)));
}

int QMetaProperty_Reset(void* ptr, void* object){
	return static_cast<QMetaProperty*>(ptr)->reset(static_cast<QObject*>(object));
}

int QMetaProperty_Revision(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->revision();
}

int QMetaProperty_UserType(void* ptr){
	return static_cast<QMetaProperty*>(ptr)->userType();
}

int QMetaProperty_Write(void* ptr, void* object, void* value){
	return static_cast<QMetaProperty*>(ptr)->write(static_cast<QObject*>(object), *static_cast<QVariant*>(value));
}

void* QMetaType_NewQMetaType(int typeId){
	return new QMetaType(typeId);
}

int QMetaType_Flags(void* ptr){
	return static_cast<QMetaType*>(ptr)->flags();
}

int QMetaType_QMetaType_IsRegistered(int ty){
	return QMetaType::isRegistered(ty);
}

int QMetaType_IsRegistered2(void* ptr){
	return static_cast<QMetaType*>(ptr)->isRegistered();
}

int QMetaType_IsValid(void* ptr){
	return static_cast<QMetaType*>(ptr)->isValid();
}

void* QMetaType_MetaObject(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QMetaType*>(ptr)->metaObject());
}

void* QMetaType_QMetaType_MetaObjectForType(int ty){
	return const_cast<QMetaObject*>(QMetaType::metaObjectForType(ty));
}

int QMetaType_QMetaType_SizeOf(int ty){
	return QMetaType::sizeOf(ty);
}

int QMetaType_SizeOf2(void* ptr){
	return static_cast<QMetaType*>(ptr)->sizeOf();
}

int QMetaType_QMetaType_Type2(void* typeName){
	return QMetaType::type(*static_cast<QByteArray*>(typeName));
}

int QMetaType_QMetaType_Type(char* typeName){
	return QMetaType::type(const_cast<const char*>(typeName));
}

int QMetaType_QMetaType_TypeFlags(int ty){
	return QMetaType::typeFlags(ty);
}

void QMetaType_DestroyQMetaType(void* ptr){
	static_cast<QMetaType*>(ptr)->~QMetaType();
}

class MyQMimeData: public QMimeData {
public:
	MyQMimeData() : QMimeData() {};
	void timerEvent(QTimerEvent * event) { callbackQMimeDataTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQMimeDataChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQMimeDataCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QMimeData_NewQMimeData(){
	return new MyQMimeData();
}

void QMimeData_Clear(void* ptr){
	static_cast<QMimeData*>(ptr)->clear();
}

void* QMimeData_ColorData(void* ptr){
	return new QVariant(static_cast<QMimeData*>(ptr)->colorData());
}

void* QMimeData_Data(void* ptr, char* mimeType){
	return new QByteArray(static_cast<QMimeData*>(ptr)->data(QString(mimeType)));
}

char* QMimeData_Formats(void* ptr){
	return static_cast<QMimeData*>(ptr)->formats().join(",,,").toUtf8().data();
}

int QMimeData_HasColor(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasColor();
}

int QMimeData_HasFormat(void* ptr, char* mimeType){
	return static_cast<QMimeData*>(ptr)->hasFormat(QString(mimeType));
}

int QMimeData_HasHtml(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasHtml();
}

int QMimeData_HasImage(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasImage();
}

int QMimeData_HasText(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasText();
}

int QMimeData_HasUrls(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasUrls();
}

char* QMimeData_Html(void* ptr){
	return static_cast<QMimeData*>(ptr)->html().toUtf8().data();
}

void* QMimeData_ImageData(void* ptr){
	return new QVariant(static_cast<QMimeData*>(ptr)->imageData());
}

void QMimeData_RemoveFormat(void* ptr, char* mimeType){
	static_cast<QMimeData*>(ptr)->removeFormat(QString(mimeType));
}

void QMimeData_SetColorData(void* ptr, void* color){
	static_cast<QMimeData*>(ptr)->setColorData(*static_cast<QVariant*>(color));
}

void QMimeData_SetData(void* ptr, char* mimeType, void* data){
	static_cast<QMimeData*>(ptr)->setData(QString(mimeType), *static_cast<QByteArray*>(data));
}

void QMimeData_SetHtml(void* ptr, char* html){
	static_cast<QMimeData*>(ptr)->setHtml(QString(html));
}

void QMimeData_SetImageData(void* ptr, void* image){
	static_cast<QMimeData*>(ptr)->setImageData(*static_cast<QVariant*>(image));
}

void QMimeData_SetText(void* ptr, char* text){
	static_cast<QMimeData*>(ptr)->setText(QString(text));
}

char* QMimeData_Text(void* ptr){
	return static_cast<QMimeData*>(ptr)->text().toUtf8().data();
}

void QMimeData_DestroyQMimeData(void* ptr){
	static_cast<QMimeData*>(ptr)->~QMimeData();
}

void QMimeData_TimerEvent(void* ptr, void* event){
	static_cast<MyQMimeData*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMimeData_TimerEventDefault(void* ptr, void* event){
	static_cast<QMimeData*>(ptr)->QMimeData::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMimeData_ChildEvent(void* ptr, void* event){
	static_cast<MyQMimeData*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMimeData_ChildEventDefault(void* ptr, void* event){
	static_cast<QMimeData*>(ptr)->QMimeData::childEvent(static_cast<QChildEvent*>(event));
}

void QMimeData_CustomEvent(void* ptr, void* event){
	static_cast<MyQMimeData*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMimeData_CustomEventDefault(void* ptr, void* event){
	static_cast<QMimeData*>(ptr)->QMimeData::customEvent(static_cast<QEvent*>(event));
}

void* QMimeDatabase_NewQMimeDatabase(){
	return new QMimeDatabase();
}

void QMimeDatabase_DestroyQMimeDatabase(void* ptr){
	static_cast<QMimeDatabase*>(ptr)->~QMimeDatabase();
}

char* QMimeDatabase_SuffixForFileName(void* ptr, char* fileName){
	return static_cast<QMimeDatabase*>(ptr)->suffixForFileName(QString(fileName)).toUtf8().data();
}

void* QMimeType_NewQMimeType(){
	return new QMimeType();
}

void* QMimeType_NewQMimeType2(void* other){
	return new QMimeType(*static_cast<QMimeType*>(other));
}

char* QMimeType_FilterString(void* ptr){
	return static_cast<QMimeType*>(ptr)->filterString().toUtf8().data();
}

char* QMimeType_GenericIconName(void* ptr){
	return static_cast<QMimeType*>(ptr)->genericIconName().toUtf8().data();
}

char* QMimeType_GlobPatterns(void* ptr){
	return static_cast<QMimeType*>(ptr)->globPatterns().join(",,,").toUtf8().data();
}

char* QMimeType_IconName(void* ptr){
	return static_cast<QMimeType*>(ptr)->iconName().toUtf8().data();
}

int QMimeType_Inherits(void* ptr, char* mimeTypeName){
	return static_cast<QMimeType*>(ptr)->inherits(QString(mimeTypeName));
}

int QMimeType_IsDefault(void* ptr){
	return static_cast<QMimeType*>(ptr)->isDefault();
}

int QMimeType_IsValid(void* ptr){
	return static_cast<QMimeType*>(ptr)->isValid();
}

char* QMimeType_Name(void* ptr){
	return static_cast<QMimeType*>(ptr)->name().toUtf8().data();
}

void QMimeType_DestroyQMimeType(void* ptr){
	static_cast<QMimeType*>(ptr)->~QMimeType();
}

char* QMimeType_Aliases(void* ptr){
	return static_cast<QMimeType*>(ptr)->aliases().join(",,,").toUtf8().data();
}

char* QMimeType_AllAncestors(void* ptr){
	return static_cast<QMimeType*>(ptr)->allAncestors().join(",,,").toUtf8().data();
}

char* QMimeType_Comment(void* ptr){
	return static_cast<QMimeType*>(ptr)->comment().toUtf8().data();
}

char* QMimeType_ParentMimeTypes(void* ptr){
	return static_cast<QMimeType*>(ptr)->parentMimeTypes().join(",,,").toUtf8().data();
}

char* QMimeType_PreferredSuffix(void* ptr){
	return static_cast<QMimeType*>(ptr)->preferredSuffix().toUtf8().data();
}

char* QMimeType_Suffixes(void* ptr){
	return static_cast<QMimeType*>(ptr)->suffixes().join(",,,").toUtf8().data();
}

void QMimeType_Swap(void* ptr, void* other){
	static_cast<QMimeType*>(ptr)->swap(*static_cast<QMimeType*>(other));
}

void* QModelIndex_NewQModelIndex(){
	return new QModelIndex();
}

void* QModelIndex_Child(void* ptr, int row, int column){
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->child(row, column));
}

int QModelIndex_Column(void* ptr){
	return static_cast<QModelIndex*>(ptr)->column();
}

void* QModelIndex_Data(void* ptr, int role){
	return new QVariant(static_cast<QModelIndex*>(ptr)->data(role));
}

int QModelIndex_Flags(void* ptr){
	return static_cast<QModelIndex*>(ptr)->flags();
}

void* QModelIndex_InternalPointer(void* ptr){
	return static_cast<QModelIndex*>(ptr)->internalPointer();
}

int QModelIndex_IsValid(void* ptr){
	return static_cast<QModelIndex*>(ptr)->isValid();
}

void* QModelIndex_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QModelIndex*>(ptr)->model());
}

void* QModelIndex_Parent(void* ptr){
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->parent());
}

int QModelIndex_Row(void* ptr){
	return static_cast<QModelIndex*>(ptr)->row();
}

void* QModelIndex_Sibling(void* ptr, int row, int column){
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->sibling(row, column));
}

void QMutex_Lock(void* ptr){
	static_cast<QMutex*>(ptr)->lock();
}

int QMutex_TryLock(void* ptr, int timeout){
	return static_cast<QMutex*>(ptr)->tryLock(timeout);
}

void QMutex_Unlock(void* ptr){
	static_cast<QMutex*>(ptr)->unlock();
}

void* QMutex_NewQMutex(int mode){
	return new QMutex(static_cast<QMutex::RecursionMode>(mode));
}

int QMutex_IsRecursive(void* ptr){
	return static_cast<QMutex*>(ptr)->isRecursive();
}

void* QMutexLocker_NewQMutexLocker(void* mutex){
	return new QMutexLocker(static_cast<QMutex*>(mutex));
}

void* QMutexLocker_Mutex(void* ptr){
	return static_cast<QMutexLocker*>(ptr)->mutex();
}

void QMutexLocker_Relock(void* ptr){
	static_cast<QMutexLocker*>(ptr)->relock();
}

void QMutexLocker_Unlock(void* ptr){
	static_cast<QMutexLocker*>(ptr)->unlock();
}

void QMutexLocker_DestroyQMutexLocker(void* ptr){
	static_cast<QMutexLocker*>(ptr)->~QMutexLocker();
}

class MyQObject: public QObject {
public:
	MyQObject(QObject *parent) : QObject(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQObjectTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQObjectChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQObjectCustomEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_Destroyed(QObject * obj) { callbackQObjectDestroyed(this, this->objectName().toUtf8().data(), obj); };
	void Signal_ObjectNameChanged(const QString & objectName) { callbackQObjectObjectNameChanged(this, this->objectName().toUtf8().data(), objectName.toUtf8().data()); };
};

void QObject_InstallEventFilter(void* ptr, void* filterObj){
	static_cast<QObject*>(ptr)->installEventFilter(static_cast<QObject*>(filterObj));
}

char* QObject_ObjectName(void* ptr){
	return static_cast<QObject*>(ptr)->objectName().toUtf8().data();
}

void QObject_SetObjectName(void* ptr, char* name){
	static_cast<QObject*>(ptr)->setObjectName(QString(name));
}

void QObject_TimerEvent(void* ptr, void* event){
	static_cast<MyQObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QObject_TimerEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QObject_NewQObject(void* parent){
	return new MyQObject(static_cast<QObject*>(parent));
}

int QObject_BlockSignals(void* ptr, int block){
	return static_cast<QObject*>(ptr)->blockSignals(block != 0);
}

void QObject_ChildEvent(void* ptr, void* event){
	static_cast<MyQObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QObject_ChildEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QObject_CustomEvent(void* ptr, void* event){
	static_cast<MyQObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QObject_CustomEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QObject_DeleteLater(void* ptr){
	QMetaObject::invokeMethod(static_cast<QObject*>(ptr), "deleteLater");
}

void QObject_ConnectDestroyed(void* ptr){
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_DisconnectDestroyed(void* ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_Destroyed(void* ptr, void* obj){
	static_cast<QObject*>(ptr)->destroyed(static_cast<QObject*>(obj));
}

void QObject_DumpObjectInfo(void* ptr){
	static_cast<QObject*>(ptr)->dumpObjectInfo();
}

void QObject_DumpObjectTree(void* ptr){
	static_cast<QObject*>(ptr)->dumpObjectTree();
}

int QObject_Event(void* ptr, void* e){
	return static_cast<QObject*>(ptr)->event(static_cast<QEvent*>(e));
}

int QObject_EventFilter(void* ptr, void* watched, void* event){
	return static_cast<QObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QObject_FindChild(void* ptr, char* name, int options){
	return static_cast<QObject*>(ptr)->findChild<QObject*>(QString(name), static_cast<Qt::FindChildOption>(options));
}

int QObject_Inherits(void* ptr, char* className){
	return static_cast<QObject*>(ptr)->inherits(const_cast<const char*>(className));
}

int QObject_IsWidgetType(void* ptr){
	return static_cast<QObject*>(ptr)->isWidgetType();
}

int QObject_IsWindowType(void* ptr){
	return static_cast<QObject*>(ptr)->isWindowType();
}

void QObject_KillTimer(void* ptr, int id){
	static_cast<QObject*>(ptr)->killTimer(id);
}

void* QObject_MetaObject(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QObject*>(ptr)->metaObject());
}

void QObject_MoveToThread(void* ptr, void* targetThread){
	static_cast<QObject*>(ptr)->moveToThread(static_cast<QThread*>(targetThread));
}

void QObject_ConnectObjectNameChanged(void* ptr){
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

void QObject_DisconnectObjectNameChanged(void* ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

void* QObject_Parent(void* ptr){
	return static_cast<QObject*>(ptr)->parent();
}

void* QObject_Property(void* ptr, char* name){
	return new QVariant(static_cast<QObject*>(ptr)->property(const_cast<const char*>(name)));
}

void QObject_RemoveEventFilter(void* ptr, void* obj){
	static_cast<QObject*>(ptr)->removeEventFilter(static_cast<QObject*>(obj));
}

void QObject_SetParent(void* ptr, void* parent){
	static_cast<QObject*>(ptr)->setParent(static_cast<QObject*>(parent));
}

int QObject_SetProperty(void* ptr, char* name, void* value){
	return static_cast<QObject*>(ptr)->setProperty(const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

int QObject_StartTimer(void* ptr, int interval, int timerType){
	return static_cast<QObject*>(ptr)->startTimer(interval, static_cast<Qt::TimerType>(timerType));
}

int QObject_SignalsBlocked(void* ptr){
	return static_cast<QObject*>(ptr)->signalsBlocked();
}

void* QObject_Thread(void* ptr){
	return static_cast<QObject*>(ptr)->thread();
}

char* QObject_QObject_Tr(char* sourceText, char* disambiguation, int n){
	return QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QObject_DestroyQObject(void* ptr){
	static_cast<QObject*>(ptr)->~QObject();
}

void* QObjectCleanupHandler_NewQObjectCleanupHandler(){
	return new QObjectCleanupHandler();
}

void* QObjectCleanupHandler_Add(void* ptr, void* object){
	return static_cast<QObjectCleanupHandler*>(ptr)->add(static_cast<QObject*>(object));
}

void QObjectCleanupHandler_Clear(void* ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->clear();
}

int QObjectCleanupHandler_IsEmpty(void* ptr){
	return static_cast<QObjectCleanupHandler*>(ptr)->isEmpty();
}

void QObjectCleanupHandler_DestroyQObjectCleanupHandler(void* ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->~QObjectCleanupHandler();
}

void QObjectCleanupHandler_TimerEvent(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QObjectCleanupHandler_TimerEventDefault(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->QObjectCleanupHandler::timerEvent(static_cast<QTimerEvent*>(event));
}

void QObjectCleanupHandler_ChildEvent(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QObjectCleanupHandler_ChildEventDefault(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->QObjectCleanupHandler::childEvent(static_cast<QChildEvent*>(event));
}

void QObjectCleanupHandler_CustomEvent(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QObjectCleanupHandler_CustomEventDefault(void* ptr, void* event){
	static_cast<QObjectCleanupHandler*>(ptr)->QObjectCleanupHandler::customEvent(static_cast<QEvent*>(event));
}

class MyQParallelAnimationGroup: public QParallelAnimationGroup {
public:
	void updateCurrentTime(int currentTime) { callbackQParallelAnimationGroupUpdateCurrentTime(this, this->objectName().toUtf8().data(), currentTime); };
	void updateDirection(QAbstractAnimation::Direction direction) { callbackQParallelAnimationGroupUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQParallelAnimationGroupUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void timerEvent(QTimerEvent * event) { callbackQParallelAnimationGroupTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQParallelAnimationGroupChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQParallelAnimationGroupCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QParallelAnimationGroup_Duration(void* ptr){
	return static_cast<QParallelAnimationGroup*>(ptr)->duration();
}

int QParallelAnimationGroup_Event(void* ptr, void* event){
	return static_cast<QParallelAnimationGroup*>(ptr)->event(static_cast<QEvent*>(event));
}

void QParallelAnimationGroup_UpdateCurrentTime(void* ptr, int currentTime){
	static_cast<MyQParallelAnimationGroup*>(ptr)->updateCurrentTime(currentTime);
}

void QParallelAnimationGroup_UpdateCurrentTimeDefault(void* ptr, int currentTime){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::updateCurrentTime(currentTime);
}

void QParallelAnimationGroup_UpdateDirection(void* ptr, int direction){
	static_cast<MyQParallelAnimationGroup*>(ptr)->updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QParallelAnimationGroup_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QParallelAnimationGroup_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQParallelAnimationGroup*>(ptr)->updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QParallelAnimationGroup_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QParallelAnimationGroup_DestroyQParallelAnimationGroup(void* ptr){
	static_cast<QParallelAnimationGroup*>(ptr)->~QParallelAnimationGroup();
}

void QParallelAnimationGroup_TimerEvent(void* ptr, void* event){
	static_cast<MyQParallelAnimationGroup*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QParallelAnimationGroup_TimerEventDefault(void* ptr, void* event){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::timerEvent(static_cast<QTimerEvent*>(event));
}

void QParallelAnimationGroup_ChildEvent(void* ptr, void* event){
	static_cast<MyQParallelAnimationGroup*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QParallelAnimationGroup_ChildEventDefault(void* ptr, void* event){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::childEvent(static_cast<QChildEvent*>(event));
}

void QParallelAnimationGroup_CustomEvent(void* ptr, void* event){
	static_cast<MyQParallelAnimationGroup*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QParallelAnimationGroup_CustomEventDefault(void* ptr, void* event){
	static_cast<QParallelAnimationGroup*>(ptr)->QParallelAnimationGroup::customEvent(static_cast<QEvent*>(event));
}

class MyQPauseAnimation: public QPauseAnimation {
public:
	MyQPauseAnimation(QObject *parent) : QPauseAnimation(parent) {};
	MyQPauseAnimation(int msecs, QObject *parent) : QPauseAnimation(msecs, parent) {};
	void updateCurrentTime(int v) { callbackQPauseAnimationUpdateCurrentTime(this, this->objectName().toUtf8().data(), v); };
	void updateDirection(QPauseAnimation::Direction direction) { callbackQPauseAnimationUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void updateState(QPauseAnimation::State newState, QPauseAnimation::State oldState) { callbackQPauseAnimationUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void timerEvent(QTimerEvent * event) { callbackQPauseAnimationTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQPauseAnimationChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQPauseAnimationCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QPauseAnimation_Duration(void* ptr){
	return static_cast<QPauseAnimation*>(ptr)->duration();
}

void QPauseAnimation_SetDuration(void* ptr, int msecs){
	static_cast<QPauseAnimation*>(ptr)->setDuration(msecs);
}

void* QPauseAnimation_NewQPauseAnimation(void* parent){
	return new MyQPauseAnimation(static_cast<QObject*>(parent));
}

void* QPauseAnimation_NewQPauseAnimation2(int msecs, void* parent){
	return new MyQPauseAnimation(msecs, static_cast<QObject*>(parent));
}

int QPauseAnimation_Event(void* ptr, void* e){
	return static_cast<QPauseAnimation*>(ptr)->event(static_cast<QEvent*>(e));
}

void QPauseAnimation_UpdateCurrentTime(void* ptr, int v){
	static_cast<MyQPauseAnimation*>(ptr)->updateCurrentTime(v);
}

void QPauseAnimation_UpdateCurrentTimeDefault(void* ptr, int v){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::updateCurrentTime(v);
}

void QPauseAnimation_DestroyQPauseAnimation(void* ptr){
	static_cast<QPauseAnimation*>(ptr)->~QPauseAnimation();
}

void QPauseAnimation_UpdateDirection(void* ptr, int direction){
	static_cast<MyQPauseAnimation*>(ptr)->updateDirection(static_cast<QPauseAnimation::Direction>(direction));
}

void QPauseAnimation_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::updateDirection(static_cast<QPauseAnimation::Direction>(direction));
}

void QPauseAnimation_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQPauseAnimation*>(ptr)->updateState(static_cast<QPauseAnimation::State>(newState), static_cast<QPauseAnimation::State>(oldState));
}

void QPauseAnimation_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::updateState(static_cast<QPauseAnimation::State>(newState), static_cast<QPauseAnimation::State>(oldState));
}

void QPauseAnimation_TimerEvent(void* ptr, void* event){
	static_cast<MyQPauseAnimation*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPauseAnimation_TimerEventDefault(void* ptr, void* event){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPauseAnimation_ChildEvent(void* ptr, void* event){
	static_cast<MyQPauseAnimation*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPauseAnimation_ChildEventDefault(void* ptr, void* event){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::childEvent(static_cast<QChildEvent*>(event));
}

void QPauseAnimation_CustomEvent(void* ptr, void* event){
	static_cast<MyQPauseAnimation*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPauseAnimation_CustomEventDefault(void* ptr, void* event){
	static_cast<QPauseAnimation*>(ptr)->QPauseAnimation::customEvent(static_cast<QEvent*>(event));
}

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

int QPersistentModelIndex_Column(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->column();
}

int QPersistentModelIndex_IsValid(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->isValid();
}

int QPersistentModelIndex_Row(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->row();
}

void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_NewQPersistentModelIndex(void* index){
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

void* QPersistentModelIndex_Child(void* ptr, int row, int column){
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->child(row, column));
}

void* QPersistentModelIndex_Data(void* ptr, int role){
	return new QVariant(static_cast<QPersistentModelIndex*>(ptr)->data(role));
}

int QPersistentModelIndex_Flags(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->flags();
}

void* QPersistentModelIndex_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QPersistentModelIndex*>(ptr)->model());
}

void* QPersistentModelIndex_Parent(void* ptr){
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->parent());
}

void* QPersistentModelIndex_Sibling(void* ptr, int row, int column){
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->sibling(row, column));
}

void QPersistentModelIndex_Swap(void* ptr, void* other){
	static_cast<QPersistentModelIndex*>(ptr)->swap(*static_cast<QPersistentModelIndex*>(other));
}

char* QPluginLoader_FileName(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->fileName().toUtf8().data();
}

int QPluginLoader_LoadHints(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->loadHints();
}

void QPluginLoader_SetFileName(void* ptr, char* fileName){
	static_cast<QPluginLoader*>(ptr)->setFileName(QString(fileName));
}

void QPluginLoader_SetLoadHints(void* ptr, int loadHints){
	static_cast<QPluginLoader*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(loadHints));
}

void* QPluginLoader_NewQPluginLoader(void* parent){
	return new QPluginLoader(static_cast<QObject*>(parent));
}

void* QPluginLoader_NewQPluginLoader2(char* fileName, void* parent){
	return new QPluginLoader(QString(fileName), static_cast<QObject*>(parent));
}

char* QPluginLoader_ErrorString(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->errorString().toUtf8().data();
}

void* QPluginLoader_Instance(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->instance();
}

int QPluginLoader_IsLoaded(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->isLoaded();
}

int QPluginLoader_Load(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->load();
}

void* QPluginLoader_MetaData(void* ptr){
	return new QJsonObject(static_cast<QPluginLoader*>(ptr)->metaData());
}

int QPluginLoader_Unload(void* ptr){
	return static_cast<QPluginLoader*>(ptr)->unload();
}

void QPluginLoader_DestroyQPluginLoader(void* ptr){
	static_cast<QPluginLoader*>(ptr)->~QPluginLoader();
}

void QPluginLoader_TimerEvent(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPluginLoader_TimerEventDefault(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->QPluginLoader::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPluginLoader_ChildEvent(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPluginLoader_ChildEventDefault(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->QPluginLoader::childEvent(static_cast<QChildEvent*>(event));
}

void QPluginLoader_CustomEvent(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPluginLoader_CustomEventDefault(void* ptr, void* event){
	static_cast<QPluginLoader*>(ptr)->QPluginLoader::customEvent(static_cast<QEvent*>(event));
}

void* QPoint_NewQPoint(){
	return new QPoint();
}

void* QPoint_NewQPoint2(int xpos, int ypos){
	return new QPoint(xpos, ypos);
}

int QPoint_QPoint_DotProduct(void* p1, void* p2){
	return QPoint::dotProduct(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

int QPoint_IsNull(void* ptr){
	return static_cast<QPoint*>(ptr)->isNull();
}

int QPoint_ManhattanLength(void* ptr){
	return static_cast<QPoint*>(ptr)->manhattanLength();
}

int QPoint_Rx(void* ptr){
	return static_cast<QPoint*>(ptr)->rx();
}

int QPoint_Ry(void* ptr){
	return static_cast<QPoint*>(ptr)->ry();
}

void QPoint_SetX(void* ptr, int x){
	static_cast<QPoint*>(ptr)->setX(x);
}

void QPoint_SetY(void* ptr, int y){
	static_cast<QPoint*>(ptr)->setY(y);
}

int QPoint_X(void* ptr){
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(void* ptr){
	return static_cast<QPoint*>(ptr)->y();
}

void* QPointF_NewQPointF(){
	return new QPointF();
}

void* QPointF_NewQPointF2(void* point){
	return new QPointF(*static_cast<QPoint*>(point));
}

void* QPointF_NewQPointF3(double xpos, double ypos){
	return new QPointF(static_cast<double>(xpos), static_cast<double>(ypos));
}

double QPointF_QPointF_DotProduct(void* p1, void* p2){
	return static_cast<double>(QPointF::dotProduct(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2)));
}

int QPointF_IsNull(void* ptr){
	return static_cast<QPointF*>(ptr)->isNull();
}

double QPointF_ManhattanLength(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->manhattanLength());
}

double QPointF_Rx(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->rx());
}

double QPointF_Ry(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->ry());
}

void QPointF_SetX(void* ptr, double x){
	static_cast<QPointF*>(ptr)->setX(static_cast<double>(x));
}

void QPointF_SetY(void* ptr, double y){
	static_cast<QPointF*>(ptr)->setY(static_cast<double>(y));
}

void* QPointF_ToPoint(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QPointF*>(ptr)->toPoint()).x(), static_cast<QPoint>(static_cast<QPointF*>(ptr)->toPoint()).y());
}

double QPointF_X(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->x());
}

double QPointF_Y(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->y());
}

class MyQProcess: public QProcess {
public:
	MyQProcess(QObject *parent) : QProcess(parent) {};
	void close() { callbackQProcessClose(this, this->objectName().toUtf8().data()); };
	void Signal_Error2(QProcess::ProcessError error) { callbackQProcessError2(this, this->objectName().toUtf8().data(), error); };
	void Signal_Finished(int exitCode, QProcess::ExitStatus exitStatus) { callbackQProcessFinished(this, this->objectName().toUtf8().data(), exitCode, exitStatus); };
	void Signal_ReadyReadStandardError() { callbackQProcessReadyReadStandardError(this, this->objectName().toUtf8().data()); };
	void Signal_ReadyReadStandardOutput() { callbackQProcessReadyReadStandardOutput(this, this->objectName().toUtf8().data()); };
	void setupChildProcess() { callbackQProcessSetupChildProcess(this, this->objectName().toUtf8().data()); };
	void Signal_Started() { callbackQProcessStarted(this, this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QProcess::ProcessState newState) { callbackQProcessStateChanged(this, this->objectName().toUtf8().data(), newState); };
	void timerEvent(QTimerEvent * event) { callbackQProcessTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQProcessChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQProcessCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QProcess_NewQProcess(void* parent){
	return new MyQProcess(static_cast<QObject*>(parent));
}

char* QProcess_Arguments(void* ptr){
	return static_cast<QProcess*>(ptr)->arguments().join(",,,").toUtf8().data();
}

int QProcess_AtEnd(void* ptr){
	return static_cast<QProcess*>(ptr)->atEnd();
}

long long QProcess_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QProcess*>(ptr)->bytesAvailable());
}

long long QProcess_BytesToWrite(void* ptr){
	return static_cast<long long>(static_cast<QProcess*>(ptr)->bytesToWrite());
}

int QProcess_CanReadLine(void* ptr){
	return static_cast<QProcess*>(ptr)->canReadLine();
}

void QProcess_Close(void* ptr){
	static_cast<MyQProcess*>(ptr)->close();
}

void QProcess_CloseDefault(void* ptr){
	static_cast<QProcess*>(ptr)->QProcess::close();
}

void QProcess_CloseReadChannel(void* ptr, int channel){
	static_cast<QProcess*>(ptr)->closeReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_CloseWriteChannel(void* ptr){
	static_cast<QProcess*>(ptr)->closeWriteChannel();
}

void QProcess_ConnectError2(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(QProcess::ProcessError)>(&QProcess::error), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessError)>(&MyQProcess::Signal_Error2));;
}

void QProcess_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(QProcess::ProcessError)>(&QProcess::error), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessError)>(&MyQProcess::Signal_Error2));;
}

void QProcess_Error2(void* ptr, int error){
	static_cast<QProcess*>(ptr)->error(static_cast<QProcess::ProcessError>(error));
}

int QProcess_Error(void* ptr){
	return static_cast<QProcess*>(ptr)->error();
}

int QProcess_QProcess_Execute2(char* command){
	return QProcess::execute(QString(command));
}

int QProcess_QProcess_Execute(char* program, char* arguments){
	return QProcess::execute(QString(program), QString(arguments).split(",,,", QString::SkipEmptyParts));
}

int QProcess_ExitCode(void* ptr){
	return static_cast<QProcess*>(ptr)->exitCode();
}

int QProcess_ExitStatus(void* ptr){
	return static_cast<QProcess*>(ptr)->exitStatus();
}

void QProcess_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

void QProcess_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

void QProcess_Finished(void* ptr, int exitCode, int exitStatus){
	static_cast<QProcess*>(ptr)->finished(exitCode, static_cast<QProcess::ExitStatus>(exitStatus));
}

int QProcess_InputChannelMode(void* ptr){
	return static_cast<QProcess*>(ptr)->inputChannelMode();
}

int QProcess_IsSequential(void* ptr){
	return static_cast<QProcess*>(ptr)->isSequential();
}

void QProcess_Kill(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "kill");
}

char* QProcess_QProcess_NullDevice(){
	return QProcess::nullDevice().toUtf8().data();
}

int QProcess_Open(void* ptr, int mode){
	return static_cast<QProcess*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_ProcessChannelMode(void* ptr){
	return static_cast<QProcess*>(ptr)->processChannelMode();
}

long long QProcess_ProcessId(void* ptr){
	return static_cast<long long>(static_cast<QProcess*>(ptr)->processId());
}

char* QProcess_Program(void* ptr){
	return static_cast<QProcess*>(ptr)->program().toUtf8().data();
}

void* QProcess_ReadAllStandardError(void* ptr){
	return new QByteArray(static_cast<QProcess*>(ptr)->readAllStandardError());
}

void* QProcess_ReadAllStandardOutput(void* ptr){
	return new QByteArray(static_cast<QProcess*>(ptr)->readAllStandardOutput());
}

int QProcess_ReadChannel(void* ptr){
	return static_cast<QProcess*>(ptr)->readChannel();
}

long long QProcess_ReadData(void* ptr, char* data, long long maxlen){
	return static_cast<long long>(static_cast<QProcess*>(ptr)->readData(data, static_cast<long long>(maxlen)));
}

void QProcess_ConnectReadyReadStandardError(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_DisconnectReadyReadStandardError(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_ConnectReadyReadStandardOutput(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_DisconnectReadyReadStandardOutput(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_SetArguments(void* ptr, char* arguments){
	static_cast<QProcess*>(ptr)->setArguments(QString(arguments).split(",,,", QString::SkipEmptyParts));
}

void QProcess_SetInputChannelMode(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->setInputChannelMode(static_cast<QProcess::InputChannelMode>(mode));
}

void QProcess_SetProcessChannelMode(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->setProcessChannelMode(static_cast<QProcess::ProcessChannelMode>(mode));
}

void QProcess_SetProcessEnvironment(void* ptr, void* environment){
	static_cast<QProcess*>(ptr)->setProcessEnvironment(*static_cast<QProcessEnvironment*>(environment));
}

void QProcess_SetProgram(void* ptr, char* program){
	static_cast<QProcess*>(ptr)->setProgram(QString(program));
}

void QProcess_SetReadChannel(void* ptr, int channel){
	static_cast<QProcess*>(ptr)->setReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_SetStandardErrorFile(void* ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardErrorFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardInputFile(void* ptr, char* fileName){
	static_cast<QProcess*>(ptr)->setStandardInputFile(QString(fileName));
}

void QProcess_SetStandardOutputFile(void* ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardOutputFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardOutputProcess(void* ptr, void* destination){
	static_cast<QProcess*>(ptr)->setStandardOutputProcess(static_cast<QProcess*>(destination));
}

void QProcess_SetWorkingDirectory(void* ptr, char* dir){
	static_cast<QProcess*>(ptr)->setWorkingDirectory(QString(dir));
}

void QProcess_SetupChildProcess(void* ptr){
	static_cast<MyQProcess*>(ptr)->setupChildProcess();
}

void QProcess_SetupChildProcessDefault(void* ptr){
	static_cast<QProcess*>(ptr)->QProcess::setupChildProcess();
}

void QProcess_Start2(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->start(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start3(void* ptr, char* command, int mode){
	static_cast<QProcess*>(ptr)->start(QString(command), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start(void* ptr, char* program, char* arguments, int mode){
	static_cast<QProcess*>(ptr)->start(QString(program), QString(arguments).split(",,,", QString::SkipEmptyParts), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_QProcess_StartDetached2(char* command){
	return QProcess::startDetached(QString(command));
}

void QProcess_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

void QProcess_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

int QProcess_State(void* ptr){
	return static_cast<QProcess*>(ptr)->state();
}

void QProcess_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

void QProcess_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

char* QProcess_QProcess_SystemEnvironment(){
	return QProcess::systemEnvironment().join(",,,").toUtf8().data();
}

void QProcess_Terminate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "terminate");
}

int QProcess_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForBytesWritten(msecs);
}

int QProcess_WaitForFinished(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForFinished(msecs);
}

int QProcess_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForReadyRead(msecs);
}

int QProcess_WaitForStarted(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForStarted(msecs);
}

char* QProcess_WorkingDirectory(void* ptr){
	return static_cast<QProcess*>(ptr)->workingDirectory().toUtf8().data();
}

long long QProcess_WriteData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QProcess*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QProcess_DestroyQProcess(void* ptr){
	static_cast<QProcess*>(ptr)->~QProcess();
}

void QProcess_TimerEvent(void* ptr, void* event){
	static_cast<MyQProcess*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QProcess_TimerEventDefault(void* ptr, void* event){
	static_cast<QProcess*>(ptr)->QProcess::timerEvent(static_cast<QTimerEvent*>(event));
}

void QProcess_ChildEvent(void* ptr, void* event){
	static_cast<MyQProcess*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QProcess_ChildEventDefault(void* ptr, void* event){
	static_cast<QProcess*>(ptr)->QProcess::childEvent(static_cast<QChildEvent*>(event));
}

void QProcess_CustomEvent(void* ptr, void* event){
	static_cast<MyQProcess*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QProcess_CustomEventDefault(void* ptr, void* event){
	static_cast<QProcess*>(ptr)->QProcess::customEvent(static_cast<QEvent*>(event));
}

void* QProcessEnvironment_NewQProcessEnvironment(){
	return new QProcessEnvironment();
}

void* QProcessEnvironment_NewQProcessEnvironment2(void* other){
	return new QProcessEnvironment(*static_cast<QProcessEnvironment*>(other));
}

void QProcessEnvironment_Clear(void* ptr){
	static_cast<QProcessEnvironment*>(ptr)->clear();
}

int QProcessEnvironment_Contains(void* ptr, char* name){
	return static_cast<QProcessEnvironment*>(ptr)->contains(QString(name));
}

int QProcessEnvironment_IsEmpty(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->isEmpty();
}

char* QProcessEnvironment_Keys(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->keys().join(",,,").toUtf8().data();
}

void QProcessEnvironment_Swap(void* ptr, void* other){
	static_cast<QProcessEnvironment*>(ptr)->swap(*static_cast<QProcessEnvironment*>(other));
}

char* QProcessEnvironment_ToStringList(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->toStringList().join(",,,").toUtf8().data();
}

char* QProcessEnvironment_Value(void* ptr, char* name, char* defaultValue){
	return static_cast<QProcessEnvironment*>(ptr)->value(QString(name), QString(defaultValue)).toUtf8().data();
}

void QProcessEnvironment_DestroyQProcessEnvironment(void* ptr){
	static_cast<QProcessEnvironment*>(ptr)->~QProcessEnvironment();
}

class MyQPropertyAnimation: public QPropertyAnimation {
public:
	MyQPropertyAnimation(QObject *parent) : QPropertyAnimation(parent) {};
	MyQPropertyAnimation(QObject *target, const QByteArray &propertyName, QObject *parent) : QPropertyAnimation(target, propertyName, parent) {};
	void updateCurrentValue(const QVariant & value) { callbackQPropertyAnimationUpdateCurrentValue(this, this->objectName().toUtf8().data(), new QVariant(value)); };
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQPropertyAnimationUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void updateCurrentTime(int v) { callbackQPropertyAnimationUpdateCurrentTime(this, this->objectName().toUtf8().data(), v); };
	void updateDirection(QPropertyAnimation::Direction direction) { callbackQPropertyAnimationUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void timerEvent(QTimerEvent * event) { callbackQPropertyAnimationTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQPropertyAnimationChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQPropertyAnimationCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QPropertyAnimation_PropertyName(void* ptr){
	return new QByteArray(static_cast<QPropertyAnimation*>(ptr)->propertyName());
}

void QPropertyAnimation_SetPropertyName(void* ptr, void* propertyName){
	static_cast<QPropertyAnimation*>(ptr)->setPropertyName(*static_cast<QByteArray*>(propertyName));
}

void QPropertyAnimation_SetTargetObject(void* ptr, void* target){
	static_cast<QPropertyAnimation*>(ptr)->setTargetObject(static_cast<QObject*>(target));
}

void* QPropertyAnimation_TargetObject(void* ptr){
	return static_cast<QPropertyAnimation*>(ptr)->targetObject();
}

void* QPropertyAnimation_NewQPropertyAnimation(void* parent){
	return new MyQPropertyAnimation(static_cast<QObject*>(parent));
}

void* QPropertyAnimation_NewQPropertyAnimation2(void* target, void* propertyName, void* parent){
	return new MyQPropertyAnimation(static_cast<QObject*>(target), *static_cast<QByteArray*>(propertyName), static_cast<QObject*>(parent));
}

int QPropertyAnimation_Event(void* ptr, void* event){
	return static_cast<QPropertyAnimation*>(ptr)->event(static_cast<QEvent*>(event));
}

void QPropertyAnimation_UpdateCurrentValue(void* ptr, void* value){
	static_cast<MyQPropertyAnimation*>(ptr)->updateCurrentValue(*static_cast<QVariant*>(value));
}

void QPropertyAnimation_UpdateCurrentValueDefault(void* ptr, void* value){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::updateCurrentValue(*static_cast<QVariant*>(value));
}

void QPropertyAnimation_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQPropertyAnimation*>(ptr)->updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QPropertyAnimation_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QPropertyAnimation_DestroyQPropertyAnimation(void* ptr){
	static_cast<QPropertyAnimation*>(ptr)->~QPropertyAnimation();
}

void QPropertyAnimation_UpdateCurrentTime(void* ptr, int v){
	static_cast<MyQPropertyAnimation*>(ptr)->updateCurrentTime(v);
}

void QPropertyAnimation_UpdateCurrentTimeDefault(void* ptr, int v){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::updateCurrentTime(v);
}

void QPropertyAnimation_UpdateDirection(void* ptr, int direction){
	static_cast<MyQPropertyAnimation*>(ptr)->updateDirection(static_cast<QPropertyAnimation::Direction>(direction));
}

void QPropertyAnimation_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::updateDirection(static_cast<QPropertyAnimation::Direction>(direction));
}

void QPropertyAnimation_TimerEvent(void* ptr, void* event){
	static_cast<MyQPropertyAnimation*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPropertyAnimation_TimerEventDefault(void* ptr, void* event){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPropertyAnimation_ChildEvent(void* ptr, void* event){
	static_cast<MyQPropertyAnimation*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPropertyAnimation_ChildEventDefault(void* ptr, void* event){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::childEvent(static_cast<QChildEvent*>(event));
}

void QPropertyAnimation_CustomEvent(void* ptr, void* event){
	static_cast<MyQPropertyAnimation*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPropertyAnimation_CustomEventDefault(void* ptr, void* event){
	static_cast<QPropertyAnimation*>(ptr)->QPropertyAnimation::customEvent(static_cast<QEvent*>(event));
}

void* QReadLocker_NewQReadLocker(void* lock){
	return new QReadLocker(static_cast<QReadWriteLock*>(lock));
}

void* QReadLocker_ReadWriteLock(void* ptr){
	return static_cast<QReadLocker*>(ptr)->readWriteLock();
}

void QReadLocker_Relock(void* ptr){
	static_cast<QReadLocker*>(ptr)->relock();
}

void QReadLocker_Unlock(void* ptr){
	static_cast<QReadLocker*>(ptr)->unlock();
}

void QReadLocker_DestroyQReadLocker(void* ptr){
	static_cast<QReadLocker*>(ptr)->~QReadLocker();
}

void* QReadWriteLock_NewQReadWriteLock(int recursionMode){
	return new QReadWriteLock(static_cast<QReadWriteLock::RecursionMode>(recursionMode));
}

void QReadWriteLock_LockForRead(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForRead();
}

void QReadWriteLock_LockForWrite(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForWrite();
}

int QReadWriteLock_TryLockForRead(void* ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead();
}

int QReadWriteLock_TryLockForRead2(void* ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead(timeout);
}

int QReadWriteLock_TryLockForWrite(void* ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite();
}

int QReadWriteLock_TryLockForWrite2(void* ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite(timeout);
}

void QReadWriteLock_Unlock(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->unlock();
}

void QReadWriteLock_DestroyQReadWriteLock(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->~QReadWriteLock();
}

int QRect_Contains(void* ptr, void* point, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QPoint*>(point), proper != 0);
}

int QRect_Contains4(void* ptr, void* rectangle, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QRect*>(rectangle), proper != 0);
}

int QRect_Intersects(void* ptr, void* rectangle){
	return static_cast<QRect*>(ptr)->intersects(*static_cast<QRect*>(rectangle));
}

void* QRect_NewQRect(){
	return new QRect();
}

void* QRect_NewQRect2(void* topLeft, void* bottomRight){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

void* QRect_NewQRect3(void* topLeft, void* size){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

void* QRect_NewQRect4(int x, int y, int width, int height){
	return new QRect(x, y, width, height);
}

void QRect_Adjust(void* ptr, int dx1, int dy1, int dx2, int dy2){
	static_cast<QRect*>(ptr)->adjust(dx1, dy1, dx2, dy2);
}

void* QRect_Adjusted(void* ptr, int dx1, int dy1, int dx2, int dy2){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->adjusted(dx1, dy1, dx2, dy2)).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->adjusted(dx1, dy1, dx2, dy2)).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->adjusted(dx1, dy1, dx2, dy2)).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->adjusted(dx1, dy1, dx2, dy2)).height());
}

int QRect_Bottom(void* ptr){
	return static_cast<QRect*>(ptr)->bottom();
}

void* QRect_BottomLeft(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QRect*>(ptr)->bottomLeft()).x(), static_cast<QPoint>(static_cast<QRect*>(ptr)->bottomLeft()).y());
}

void* QRect_BottomRight(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QRect*>(ptr)->bottomRight()).x(), static_cast<QPoint>(static_cast<QRect*>(ptr)->bottomRight()).y());
}

void* QRect_Center(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QRect*>(ptr)->center()).x(), static_cast<QPoint>(static_cast<QRect*>(ptr)->center()).y());
}

int QRect_Contains3(void* ptr, int x, int y){
	return static_cast<QRect*>(ptr)->contains(x, y);
}

int QRect_Contains2(void* ptr, int x, int y, int proper){
	return static_cast<QRect*>(ptr)->contains(x, y, proper != 0);
}

void QRect_GetCoords(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->getCoords(&x1, &y1, &x2, &y2);
}

void QRect_GetRect(void* ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->getRect(&x, &y, &width, &height);
}

int QRect_Height(void* ptr){
	return static_cast<QRect*>(ptr)->height();
}

void* QRect_Intersected(void* ptr, void* rectangle){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->intersected(*static_cast<QRect*>(rectangle))).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->intersected(*static_cast<QRect*>(rectangle))).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->intersected(*static_cast<QRect*>(rectangle))).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->intersected(*static_cast<QRect*>(rectangle))).height());
}

int QRect_IsEmpty(void* ptr){
	return static_cast<QRect*>(ptr)->isEmpty();
}

int QRect_IsNull(void* ptr){
	return static_cast<QRect*>(ptr)->isNull();
}

int QRect_IsValid(void* ptr){
	return static_cast<QRect*>(ptr)->isValid();
}

int QRect_Left(void* ptr){
	return static_cast<QRect*>(ptr)->left();
}

void* QRect_MarginsAdded(void* ptr, void* margins){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->marginsAdded(*static_cast<QMargins*>(margins))).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsAdded(*static_cast<QMargins*>(margins))).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsAdded(*static_cast<QMargins*>(margins))).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsAdded(*static_cast<QMargins*>(margins))).height());
}

void* QRect_MarginsRemoved(void* ptr, void* margins){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->marginsRemoved(*static_cast<QMargins*>(margins))).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsRemoved(*static_cast<QMargins*>(margins))).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsRemoved(*static_cast<QMargins*>(margins))).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->marginsRemoved(*static_cast<QMargins*>(margins))).height());
}

void QRect_MoveBottom(void* ptr, int y){
	static_cast<QRect*>(ptr)->moveBottom(y);
}

void QRect_MoveBottomLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveBottomRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveBottomRight(*static_cast<QPoint*>(position));
}

void QRect_MoveCenter(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveCenter(*static_cast<QPoint*>(position));
}

void QRect_MoveLeft(void* ptr, int x){
	static_cast<QRect*>(ptr)->moveLeft(x);
}

void QRect_MoveRight(void* ptr, int x){
	static_cast<QRect*>(ptr)->moveRight(x);
}

void QRect_MoveTo2(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTo(*static_cast<QPoint*>(position));
}

void QRect_MoveTo(void* ptr, int x, int y){
	static_cast<QRect*>(ptr)->moveTo(x, y);
}

void QRect_MoveTop(void* ptr, int y){
	static_cast<QRect*>(ptr)->moveTop(y);
}

void QRect_MoveTopLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTopLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveTopRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTopRight(*static_cast<QPoint*>(position));
}

void* QRect_Normalized(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->normalized()).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->normalized()).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->normalized()).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->normalized()).height());
}

int QRect_Right(void* ptr){
	return static_cast<QRect*>(ptr)->right();
}

void QRect_SetBottom(void* ptr, int y){
	static_cast<QRect*>(ptr)->setBottom(y);
}

void QRect_SetBottomLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_SetBottomRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setBottomRight(*static_cast<QPoint*>(position));
}

void QRect_SetCoords(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->setCoords(x1, y1, x2, y2);
}

void QRect_SetHeight(void* ptr, int height){
	static_cast<QRect*>(ptr)->setHeight(height);
}

void QRect_SetLeft(void* ptr, int x){
	static_cast<QRect*>(ptr)->setLeft(x);
}

void QRect_SetRect(void* ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->setRect(x, y, width, height);
}

void QRect_SetRight(void* ptr, int x){
	static_cast<QRect*>(ptr)->setRight(x);
}

void QRect_SetSize(void* ptr, void* size){
	static_cast<QRect*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QRect_SetTop(void* ptr, int y){
	static_cast<QRect*>(ptr)->setTop(y);
}

void QRect_SetTopLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setTopLeft(*static_cast<QPoint*>(position));
}

void QRect_SetTopRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setTopRight(*static_cast<QPoint*>(position));
}

void QRect_SetWidth(void* ptr, int width){
	static_cast<QRect*>(ptr)->setWidth(width);
}

void QRect_SetX(void* ptr, int x){
	static_cast<QRect*>(ptr)->setX(x);
}

void QRect_SetY(void* ptr, int y){
	static_cast<QRect*>(ptr)->setY(y);
}

void* QRect_Size(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QRect*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QRect*>(ptr)->size()).height());
}

int QRect_Top(void* ptr){
	return static_cast<QRect*>(ptr)->top();
}

void* QRect_TopLeft(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QRect*>(ptr)->topLeft()).x(), static_cast<QPoint>(static_cast<QRect*>(ptr)->topLeft()).y());
}

void* QRect_TopRight(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QRect*>(ptr)->topRight()).x(), static_cast<QPoint>(static_cast<QRect*>(ptr)->topRight()).y());
}

void QRect_Translate2(void* ptr, void* offset){
	static_cast<QRect*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QRect_Translate(void* ptr, int dx, int dy){
	static_cast<QRect*>(ptr)->translate(dx, dy);
}

void* QRect_Translated2(void* ptr, void* offset){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->translated(*static_cast<QPoint*>(offset))).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(*static_cast<QPoint*>(offset))).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(*static_cast<QPoint*>(offset))).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(*static_cast<QPoint*>(offset))).height());
}

void* QRect_Translated(void* ptr, int dx, int dy){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->translated(dx, dy)).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(dx, dy)).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(dx, dy)).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->translated(dx, dy)).height());
}

void* QRect_United(void* ptr, void* rectangle){
	return new QRect(static_cast<QRect>(static_cast<QRect*>(ptr)->united(*static_cast<QRect*>(rectangle))).x(), static_cast<QRect>(static_cast<QRect*>(ptr)->united(*static_cast<QRect*>(rectangle))).y(), static_cast<QRect>(static_cast<QRect*>(ptr)->united(*static_cast<QRect*>(rectangle))).width(), static_cast<QRect>(static_cast<QRect*>(ptr)->united(*static_cast<QRect*>(rectangle))).height());
}

int QRect_Width(void* ptr){
	return static_cast<QRect*>(ptr)->width();
}

int QRect_X(void* ptr){
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(void* ptr){
	return static_cast<QRect*>(ptr)->y();
}

int QRectF_Contains(void* ptr, void* point){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QRectF_Contains3(void* ptr, void* rectangle){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

int QRectF_Intersects(void* ptr, void* rectangle){
	return static_cast<QRectF*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

void* QRectF_ToAlignedRect(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QRectF*>(ptr)->toAlignedRect()).x(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toAlignedRect()).y(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toAlignedRect()).width(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toAlignedRect()).height());
}

void* QRectF_NewQRectF(){
	return new QRectF();
}

void* QRectF_NewQRectF3(void* topLeft, void* bottomRight){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

void* QRectF_NewQRectF2(void* topLeft, void* size){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

void* QRectF_NewQRectF5(void* rectangle){
	return new QRectF(*static_cast<QRect*>(rectangle));
}

void* QRectF_NewQRectF4(double x, double y, double width, double height){
	return new QRectF(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height));
}

void QRectF_Adjust(void* ptr, double dx1, double dy1, double dx2, double dy2){
	static_cast<QRectF*>(ptr)->adjust(static_cast<double>(dx1), static_cast<double>(dy1), static_cast<double>(dx2), static_cast<double>(dy2));
}

double QRectF_Bottom(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->bottom());
}

int QRectF_Contains2(void* ptr, double x, double y){
	return static_cast<QRectF*>(ptr)->contains(static_cast<double>(x), static_cast<double>(y));
}

double QRectF_Height(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->height());
}

int QRectF_IsEmpty(void* ptr){
	return static_cast<QRectF*>(ptr)->isEmpty();
}

int QRectF_IsNull(void* ptr){
	return static_cast<QRectF*>(ptr)->isNull();
}

int QRectF_IsValid(void* ptr){
	return static_cast<QRectF*>(ptr)->isValid();
}

double QRectF_Left(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->left());
}

void QRectF_MoveBottom(void* ptr, double y){
	static_cast<QRectF*>(ptr)->moveBottom(static_cast<double>(y));
}

void QRectF_MoveBottomLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveBottomRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_MoveCenter(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveCenter(*static_cast<QPointF*>(position));
}

void QRectF_MoveLeft(void* ptr, double x){
	static_cast<QRectF*>(ptr)->moveLeft(static_cast<double>(x));
}

void QRectF_MoveRight(void* ptr, double x){
	static_cast<QRectF*>(ptr)->moveRight(static_cast<double>(x));
}

void QRectF_MoveTo2(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTo(*static_cast<QPointF*>(position));
}

void QRectF_MoveTo(void* ptr, double x, double y){
	static_cast<QRectF*>(ptr)->moveTo(static_cast<double>(x), static_cast<double>(y));
}

void QRectF_MoveTop(void* ptr, double y){
	static_cast<QRectF*>(ptr)->moveTop(static_cast<double>(y));
}

void QRectF_MoveTopLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveTopRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTopRight(*static_cast<QPointF*>(position));
}

double QRectF_Right(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->right());
}

void QRectF_SetBottom(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setBottom(static_cast<double>(y));
}

void QRectF_SetBottomLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetBottomRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_SetCoords(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QRectF*>(ptr)->setCoords(static_cast<double>(x1), static_cast<double>(y1), static_cast<double>(x2), static_cast<double>(y2));
}

void QRectF_SetHeight(void* ptr, double height){
	static_cast<QRectF*>(ptr)->setHeight(static_cast<double>(height));
}

void QRectF_SetLeft(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setLeft(static_cast<double>(x));
}

void QRectF_SetRect(void* ptr, double x, double y, double width, double height){
	static_cast<QRectF*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height));
}

void QRectF_SetRight(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setRight(static_cast<double>(x));
}

void QRectF_SetSize(void* ptr, void* size){
	static_cast<QRectF*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QRectF_SetTop(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setTop(static_cast<double>(y));
}

void QRectF_SetTopLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetTopRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setTopRight(*static_cast<QPointF*>(position));
}

void QRectF_SetWidth(void* ptr, double width){
	static_cast<QRectF*>(ptr)->setWidth(static_cast<double>(width));
}

void QRectF_SetX(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setX(static_cast<double>(x));
}

void QRectF_SetY(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setY(static_cast<double>(y));
}

void* QRectF_ToRect(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QRectF*>(ptr)->toRect()).x(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toRect()).y(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toRect()).width(), static_cast<QRect>(static_cast<QRectF*>(ptr)->toRect()).height());
}

double QRectF_Top(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->top());
}

void QRectF_Translate2(void* ptr, void* offset){
	static_cast<QRectF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QRectF_Translate(void* ptr, double dx, double dy){
	static_cast<QRectF*>(ptr)->translate(static_cast<double>(dx), static_cast<double>(dy));
}

double QRectF_Width(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->width());
}

double QRectF_X(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->x());
}

double QRectF_Y(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->y());
}

void* QRegExp_NewQRegExp(){
	return new QRegExp();
}

void* QRegExp_NewQRegExp3(void* rx){
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

void* QRegExp_NewQRegExp2(char* pattern, int cs, int syntax){
	return new QRegExp(QString(pattern), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

char* QRegExp_Cap(void* ptr, int nth){
	return static_cast<QRegExp*>(ptr)->cap(nth).toUtf8().data();
}

char* QRegExp_ErrorString(void* ptr){
	return static_cast<QRegExp*>(ptr)->errorString().toUtf8().data();
}

int QRegExp_CaptureCount(void* ptr){
	return static_cast<QRegExp*>(ptr)->captureCount();
}

char* QRegExp_CapturedTexts(void* ptr){
	return static_cast<QRegExp*>(ptr)->capturedTexts().join(",,,").toUtf8().data();
}

int QRegExp_CaseSensitivity(void* ptr){
	return static_cast<QRegExp*>(ptr)->caseSensitivity();
}

char* QRegExp_QRegExp_Escape(char* str){
	return QRegExp::escape(QString(str)).toUtf8().data();
}

int QRegExp_ExactMatch(void* ptr, char* str){
	return static_cast<QRegExp*>(ptr)->exactMatch(QString(str));
}

int QRegExp_IndexIn(void* ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->indexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_IsEmpty(void* ptr){
	return static_cast<QRegExp*>(ptr)->isEmpty();
}

int QRegExp_IsMinimal(void* ptr){
	return static_cast<QRegExp*>(ptr)->isMinimal();
}

int QRegExp_IsValid(void* ptr){
	return static_cast<QRegExp*>(ptr)->isValid();
}

int QRegExp_LastIndexIn(void* ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->lastIndexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_MatchedLength(void* ptr){
	return static_cast<QRegExp*>(ptr)->matchedLength();
}

char* QRegExp_Pattern(void* ptr){
	return static_cast<QRegExp*>(ptr)->pattern().toUtf8().data();
}

int QRegExp_PatternSyntax(void* ptr){
	return static_cast<QRegExp*>(ptr)->patternSyntax();
}

int QRegExp_Pos(void* ptr, int nth){
	return static_cast<QRegExp*>(ptr)->pos(nth);
}

void QRegExp_SetCaseSensitivity(void* ptr, int cs){
	static_cast<QRegExp*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QRegExp_SetMinimal(void* ptr, int minimal){
	static_cast<QRegExp*>(ptr)->setMinimal(minimal != 0);
}

void QRegExp_SetPattern(void* ptr, char* pattern){
	static_cast<QRegExp*>(ptr)->setPattern(QString(pattern));
}

void QRegExp_SetPatternSyntax(void* ptr, int syntax){
	static_cast<QRegExp*>(ptr)->setPatternSyntax(static_cast<QRegExp::PatternSyntax>(syntax));
}

void QRegExp_Swap(void* ptr, void* other){
	static_cast<QRegExp*>(ptr)->swap(*static_cast<QRegExp*>(other));
}

void QRegExp_DestroyQRegExp(void* ptr){
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

void* QRegularExpression_NewQRegularExpression(){
	return new QRegularExpression();
}

void* QRegularExpression_NewQRegularExpression3(void* re){
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

void* QRegularExpression_NewQRegularExpression2(char* pattern, int options){
	return new QRegularExpression(QString(pattern), static_cast<QRegularExpression::PatternOption>(options));
}

int QRegularExpression_CaptureCount(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->captureCount();
}

char* QRegularExpression_ErrorString(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->errorString().toUtf8().data();
}

char* QRegularExpression_QRegularExpression_Escape(char* str){
	return QRegularExpression::escape(QString(str)).toUtf8().data();
}

void* QRegularExpression_GlobalMatch(void* ptr, char* subject, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatchIterator(static_cast<QRegularExpression*>(ptr)->globalMatch(QString(subject), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_GlobalMatch2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatchIterator(static_cast<QRegularExpression*>(ptr)->globalMatch(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

int QRegularExpression_IsValid(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->isValid();
}

void* QRegularExpression_Match(void* ptr, char* subject, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(QString(subject), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_Match2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

char* QRegularExpression_NamedCaptureGroups(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->namedCaptureGroups().join(",,,").toUtf8().data();
}

void QRegularExpression_Optimize(void* ptr){
	static_cast<QRegularExpression*>(ptr)->optimize();
}

char* QRegularExpression_Pattern(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->pattern().toUtf8().data();
}

int QRegularExpression_PatternErrorOffset(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->patternErrorOffset();
}

int QRegularExpression_PatternOptions(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->patternOptions();
}

void QRegularExpression_SetPattern(void* ptr, char* pattern){
	static_cast<QRegularExpression*>(ptr)->setPattern(QString(pattern));
}

void QRegularExpression_SetPatternOptions(void* ptr, int options){
	static_cast<QRegularExpression*>(ptr)->setPatternOptions(static_cast<QRegularExpression::PatternOption>(options));
}

void QRegularExpression_Swap(void* ptr, void* other){
	static_cast<QRegularExpression*>(ptr)->swap(*static_cast<QRegularExpression*>(other));
}

void QRegularExpression_DestroyQRegularExpression(void* ptr){
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch(){
	return new QRegularExpressionMatch();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match){
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

char* QRegularExpressionMatch_Captured2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(QString(name)).toUtf8().data();
}

char* QRegularExpressionMatch_Captured(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(nth).toUtf8().data();
}

int QRegularExpressionMatch_CapturedEnd2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(QString(name));
}

int QRegularExpressionMatch_CapturedEnd(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(nth);
}

int QRegularExpressionMatch_CapturedLength2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(QString(name));
}

int QRegularExpressionMatch_CapturedLength(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(nth);
}

void* QRegularExpressionMatch_CapturedRef2(void* ptr, char* name){
	return new QStringRef(static_cast<QRegularExpressionMatch*>(ptr)->capturedRef(QString(name)));
}

void* QRegularExpressionMatch_CapturedRef(void* ptr, int nth){
	return new QStringRef(static_cast<QRegularExpressionMatch*>(ptr)->capturedRef(nth));
}

int QRegularExpressionMatch_CapturedStart2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(QString(name));
}

int QRegularExpressionMatch_CapturedStart(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(nth);
}

char* QRegularExpressionMatch_CapturedTexts(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedTexts().join(",,,").toUtf8().data();
}

int QRegularExpressionMatch_HasMatch(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasMatch();
}

int QRegularExpressionMatch_HasPartialMatch(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasPartialMatch();
}

int QRegularExpressionMatch_IsValid(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->isValid();
}

int QRegularExpressionMatch_LastCapturedIndex(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->lastCapturedIndex();
}

int QRegularExpressionMatch_MatchOptions(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchOptions();
}

int QRegularExpressionMatch_MatchType(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchType();
}

void* QRegularExpressionMatch_RegularExpression(void* ptr){
	return new QRegularExpression(static_cast<QRegularExpressionMatch*>(ptr)->regularExpression());
}

void QRegularExpressionMatch_Swap(void* ptr, void* other){
	static_cast<QRegularExpressionMatch*>(ptr)->swap(*static_cast<QRegularExpressionMatch*>(other));
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr){
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

int QResource_QResource_RegisterResource(char* rccFileName, char* mapRoot){
	return QResource::registerResource(QString(rccFileName), QString(mapRoot));
}

int QResource_QResource_UnregisterResource(char* rccFileName, char* mapRoot){
	return QResource::unregisterResource(QString(rccFileName), QString(mapRoot));
}

void* QResource_NewQResource(char* file, void* locale){
	return new QResource(QString(file), *static_cast<QLocale*>(locale));
}

char* QResource_AbsoluteFilePath(void* ptr){
	return static_cast<QResource*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QResource_FileName(void* ptr){
	return static_cast<QResource*>(ptr)->fileName().toUtf8().data();
}

int QResource_IsCompressed(void* ptr){
	return static_cast<QResource*>(ptr)->isCompressed();
}

int QResource_IsValid(void* ptr){
	return static_cast<QResource*>(ptr)->isValid();
}

void QResource_SetFileName(void* ptr, char* file){
	static_cast<QResource*>(ptr)->setFileName(QString(file));
}

void QResource_SetLocale(void* ptr, void* locale){
	static_cast<QResource*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

long long QResource_Size(void* ptr){
	return static_cast<long long>(static_cast<QResource*>(ptr)->size());
}

void QResource_DestroyQResource(void* ptr){
	static_cast<QResource*>(ptr)->~QResource();
}

class MyQRunnable: public QRunnable {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QRunnable_AutoDelete(void* ptr){
	return static_cast<QRunnable*>(ptr)->autoDelete();
}

void QRunnable_Run(void* ptr){
	static_cast<QRunnable*>(ptr)->run();
}

void QRunnable_SetAutoDelete(void* ptr, int autoDelete){
	static_cast<QRunnable*>(ptr)->setAutoDelete(autoDelete != 0);
}

void QRunnable_DestroyQRunnable(void* ptr){
	static_cast<QRunnable*>(ptr)->~QRunnable();
}

char* QRunnable_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQRunnable*>(static_cast<QRunnable*>(ptr))) {
		return static_cast<MyQRunnable*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QRunnable_BASE").toUtf8().data();
}

void QRunnable_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQRunnable*>(static_cast<QRunnable*>(ptr))) {
		static_cast<MyQRunnable*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSaveFile: public QSaveFile {
public:
	MyQSaveFile(QObject *parent) : QSaveFile(parent) {};
	MyQSaveFile(const QString &name) : QSaveFile(name) {};
	MyQSaveFile(const QString &name, QObject *parent) : QSaveFile(name, parent) {};
	void timerEvent(QTimerEvent * event) { callbackQSaveFileTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSaveFileChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSaveFileCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSaveFile_NewQSaveFile2(void* parent){
	return new MyQSaveFile(static_cast<QObject*>(parent));
}

void* QSaveFile_NewQSaveFile(char* name){
	return new MyQSaveFile(QString(name));
}

void* QSaveFile_NewQSaveFile3(char* name, void* parent){
	return new MyQSaveFile(QString(name), static_cast<QObject*>(parent));
}

void QSaveFile_CancelWriting(void* ptr){
	static_cast<QSaveFile*>(ptr)->cancelWriting();
}

int QSaveFile_Commit(void* ptr){
	return static_cast<QSaveFile*>(ptr)->commit();
}

int QSaveFile_DirectWriteFallback(void* ptr){
	return static_cast<QSaveFile*>(ptr)->directWriteFallback();
}

char* QSaveFile_FileName(void* ptr){
	return static_cast<QSaveFile*>(ptr)->fileName().toUtf8().data();
}

int QSaveFile_Open(void* ptr, int mode){
	return static_cast<QSaveFile*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QSaveFile_SetDirectWriteFallback(void* ptr, int enabled){
	static_cast<QSaveFile*>(ptr)->setDirectWriteFallback(enabled != 0);
}

void QSaveFile_SetFileName(void* ptr, char* name){
	static_cast<QSaveFile*>(ptr)->setFileName(QString(name));
}

long long QSaveFile_WriteData(void* ptr, char* data, long long len){
	return static_cast<long long>(static_cast<QSaveFile*>(ptr)->writeData(const_cast<const char*>(data), static_cast<long long>(len)));
}

void QSaveFile_DestroyQSaveFile(void* ptr){
	static_cast<QSaveFile*>(ptr)->~QSaveFile();
}

void QSaveFile_TimerEvent(void* ptr, void* event){
	static_cast<MyQSaveFile*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSaveFile_TimerEventDefault(void* ptr, void* event){
	static_cast<QSaveFile*>(ptr)->QSaveFile::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSaveFile_ChildEvent(void* ptr, void* event){
	static_cast<MyQSaveFile*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSaveFile_ChildEventDefault(void* ptr, void* event){
	static_cast<QSaveFile*>(ptr)->QSaveFile::childEvent(static_cast<QChildEvent*>(event));
}

void QSaveFile_CustomEvent(void* ptr, void* event){
	static_cast<MyQSaveFile*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSaveFile_CustomEventDefault(void* ptr, void* event){
	static_cast<QSaveFile*>(ptr)->QSaveFile::customEvent(static_cast<QEvent*>(event));
}

void* QSemaphore_NewQSemaphore(int n){
	return new QSemaphore(n);
}

void QSemaphore_Acquire(void* ptr, int n){
	static_cast<QSemaphore*>(ptr)->acquire(n);
}

int QSemaphore_Available(void* ptr){
	return static_cast<QSemaphore*>(ptr)->available();
}

void QSemaphore_Release(void* ptr, int n){
	static_cast<QSemaphore*>(ptr)->release(n);
}

int QSemaphore_TryAcquire(void* ptr, int n){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n);
}

int QSemaphore_TryAcquire2(void* ptr, int n, int timeout){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n, timeout);
}

void QSemaphore_DestroyQSemaphore(void* ptr){
	static_cast<QSemaphore*>(ptr)->~QSemaphore();
}

class MyQSequentialAnimationGroup: public QSequentialAnimationGroup {
public:
	void Signal_CurrentAnimationChanged(QAbstractAnimation * current) { callbackQSequentialAnimationGroupCurrentAnimationChanged(this, this->objectName().toUtf8().data(), current); };
	void updateCurrentTime(int currentTime) { callbackQSequentialAnimationGroupUpdateCurrentTime(this, this->objectName().toUtf8().data(), currentTime); };
	void updateDirection(QAbstractAnimation::Direction direction) { callbackQSequentialAnimationGroupUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQSequentialAnimationGroupUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void timerEvent(QTimerEvent * event) { callbackQSequentialAnimationGroupTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSequentialAnimationGroupChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSequentialAnimationGroupCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSequentialAnimationGroup_CurrentAnimation(void* ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->currentAnimation();
}

void* QSequentialAnimationGroup_AddPause(void* ptr, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->addPause(msecs);
}

void QSequentialAnimationGroup_ConnectCurrentAnimationChanged(void* ptr){
	QObject::connect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

void QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(void* ptr){
	QObject::disconnect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

void QSequentialAnimationGroup_CurrentAnimationChanged(void* ptr, void* current){
	static_cast<QSequentialAnimationGroup*>(ptr)->currentAnimationChanged(static_cast<QAbstractAnimation*>(current));
}

int QSequentialAnimationGroup_Duration(void* ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->duration();
}

int QSequentialAnimationGroup_Event(void* ptr, void* event){
	return static_cast<QSequentialAnimationGroup*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QSequentialAnimationGroup_InsertPause(void* ptr, int index, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->insertPause(index, msecs);
}

void QSequentialAnimationGroup_UpdateCurrentTime(void* ptr, int currentTime){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->updateCurrentTime(currentTime);
}

void QSequentialAnimationGroup_UpdateCurrentTimeDefault(void* ptr, int currentTime){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::updateCurrentTime(currentTime);
}

void QSequentialAnimationGroup_UpdateDirection(void* ptr, int direction){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QSequentialAnimationGroup_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::updateDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QSequentialAnimationGroup_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QSequentialAnimationGroup_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(void* ptr){
	static_cast<QSequentialAnimationGroup*>(ptr)->~QSequentialAnimationGroup();
}

void QSequentialAnimationGroup_TimerEvent(void* ptr, void* event){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSequentialAnimationGroup_TimerEventDefault(void* ptr, void* event){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSequentialAnimationGroup_ChildEvent(void* ptr, void* event){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSequentialAnimationGroup_ChildEventDefault(void* ptr, void* event){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::childEvent(static_cast<QChildEvent*>(event));
}

void QSequentialAnimationGroup_CustomEvent(void* ptr, void* event){
	static_cast<MyQSequentialAnimationGroup*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSequentialAnimationGroup_CustomEventDefault(void* ptr, void* event){
	static_cast<QSequentialAnimationGroup*>(ptr)->QSequentialAnimationGroup::customEvent(static_cast<QEvent*>(event));
}

void* QSequentialIterable_At(void* ptr, int idx){
	return new QVariant(static_cast<QSequentialIterable*>(ptr)->at(idx));
}

int QSequentialIterable_CanReverseIterate(void* ptr){
	return static_cast<QSequentialIterable*>(ptr)->canReverseIterate();
}

int QSequentialIterable_Size(void* ptr){
	return static_cast<QSequentialIterable*>(ptr)->size();
}

class MyQSettings: public QSettings {
public:
	MyQSettings(Format format, Scope scope, const QString &organization, const QString &application, QObject *parent) : QSettings(format, scope, organization, application, parent) {};
	MyQSettings(QObject *parent) : QSettings(parent) {};
	MyQSettings(Scope scope, const QString &organization, const QString &application, QObject *parent) : QSettings(scope, organization, application, parent) {};
	MyQSettings(const QString &fileName, Format format, QObject *parent) : QSettings(fileName, format, parent) {};
	MyQSettings(const QString &organization, const QString &application, QObject *parent) : QSettings(organization, application, parent) {};
	void timerEvent(QTimerEvent * event) { callbackQSettingsTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSettingsChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSettingsCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSettings_NewQSettings3(int format, int scope, char* organization, char* application, void* parent){
	return new MyQSettings(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings5(void* parent){
	return new MyQSettings(static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings2(int scope, char* organization, char* application, void* parent){
	return new MyQSettings(static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings4(char* fileName, int format, void* parent){
	return new MyQSettings(QString(fileName), static_cast<QSettings::Format>(format), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings(char* organization, char* application, void* parent){
	return new MyQSettings(QString(organization), QString(application), static_cast<QObject*>(parent));
}

char* QSettings_AllKeys(void* ptr){
	return static_cast<QSettings*>(ptr)->allKeys().join(",,,").toUtf8().data();
}

char* QSettings_ApplicationName(void* ptr){
	return static_cast<QSettings*>(ptr)->applicationName().toUtf8().data();
}

void QSettings_BeginGroup(void* ptr, char* prefix){
	static_cast<QSettings*>(ptr)->beginGroup(QString(prefix));
}

int QSettings_BeginReadArray(void* ptr, char* prefix){
	return static_cast<QSettings*>(ptr)->beginReadArray(QString(prefix));
}

void QSettings_BeginWriteArray(void* ptr, char* prefix, int size){
	static_cast<QSettings*>(ptr)->beginWriteArray(QString(prefix), size);
}

char* QSettings_ChildGroups(void* ptr){
	return static_cast<QSettings*>(ptr)->childGroups().join(",,,").toUtf8().data();
}

char* QSettings_ChildKeys(void* ptr){
	return static_cast<QSettings*>(ptr)->childKeys().join(",,,").toUtf8().data();
}

void QSettings_Clear(void* ptr){
	static_cast<QSettings*>(ptr)->clear();
}

int QSettings_Contains(void* ptr, char* key){
	return static_cast<QSettings*>(ptr)->contains(QString(key));
}

int QSettings_QSettings_DefaultFormat(){
	return QSettings::defaultFormat();
}

void QSettings_EndArray(void* ptr){
	static_cast<QSettings*>(ptr)->endArray();
}

void QSettings_EndGroup(void* ptr){
	static_cast<QSettings*>(ptr)->endGroup();
}

int QSettings_Event(void* ptr, void* event){
	return static_cast<QSettings*>(ptr)->event(static_cast<QEvent*>(event));
}

int QSettings_FallbacksEnabled(void* ptr){
	return static_cast<QSettings*>(ptr)->fallbacksEnabled();
}

char* QSettings_FileName(void* ptr){
	return static_cast<QSettings*>(ptr)->fileName().toUtf8().data();
}

int QSettings_Format(void* ptr){
	return static_cast<QSettings*>(ptr)->format();
}

char* QSettings_Group(void* ptr){
	return static_cast<QSettings*>(ptr)->group().toUtf8().data();
}

void* QSettings_IniCodec(void* ptr){
	return static_cast<QSettings*>(ptr)->iniCodec();
}

int QSettings_IsWritable(void* ptr){
	return static_cast<QSettings*>(ptr)->isWritable();
}

char* QSettings_OrganizationName(void* ptr){
	return static_cast<QSettings*>(ptr)->organizationName().toUtf8().data();
}

int QSettings_Scope(void* ptr){
	return static_cast<QSettings*>(ptr)->scope();
}

void QSettings_SetArrayIndex(void* ptr, int i){
	static_cast<QSettings*>(ptr)->setArrayIndex(i);
}

void QSettings_QSettings_SetDefaultFormat(int format){
	QSettings::setDefaultFormat(static_cast<QSettings::Format>(format));
}

void QSettings_SetFallbacksEnabled(void* ptr, int b){
	static_cast<QSettings*>(ptr)->setFallbacksEnabled(b != 0);
}

void QSettings_SetIniCodec(void* ptr, void* codec){
	static_cast<QSettings*>(ptr)->setIniCodec(static_cast<QTextCodec*>(codec));
}

void QSettings_SetIniCodec2(void* ptr, char* codecName){
	static_cast<QSettings*>(ptr)->setIniCodec(const_cast<const char*>(codecName));
}

void QSettings_QSettings_SetPath(int format, int scope, char* path){
	QSettings::setPath(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(path));
}

void QSettings_SetValue(void* ptr, char* key, void* value){
	static_cast<QSettings*>(ptr)->setValue(QString(key), *static_cast<QVariant*>(value));
}

int QSettings_Status(void* ptr){
	return static_cast<QSettings*>(ptr)->status();
}

void QSettings_Sync(void* ptr){
	static_cast<QSettings*>(ptr)->sync();
}

void* QSettings_Value(void* ptr, char* key, void* defaultValue){
	return new QVariant(static_cast<QSettings*>(ptr)->value(QString(key), *static_cast<QVariant*>(defaultValue)));
}

void QSettings_DestroyQSettings(void* ptr){
	static_cast<QSettings*>(ptr)->~QSettings();
}

void QSettings_TimerEvent(void* ptr, void* event){
	static_cast<MyQSettings*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSettings_TimerEventDefault(void* ptr, void* event){
	static_cast<QSettings*>(ptr)->QSettings::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSettings_ChildEvent(void* ptr, void* event){
	static_cast<MyQSettings*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSettings_ChildEventDefault(void* ptr, void* event){
	static_cast<QSettings*>(ptr)->QSettings::childEvent(static_cast<QChildEvent*>(event));
}

void QSettings_CustomEvent(void* ptr, void* event){
	static_cast<MyQSettings*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSettings_CustomEventDefault(void* ptr, void* event){
	static_cast<QSettings*>(ptr)->QSettings::customEvent(static_cast<QEvent*>(event));
}

void* QSharedData_NewQSharedData(){
	return new QSharedData();
}

void* QSharedData_NewQSharedData2(void* other){
	return new QSharedData(*static_cast<QSharedData*>(other));
}

void* QSharedMemory_NewQSharedMemory2(void* parent){
	return new QSharedMemory(static_cast<QObject*>(parent));
}

void* QSharedMemory_NewQSharedMemory(char* key, void* parent){
	return new QSharedMemory(QString(key), static_cast<QObject*>(parent));
}

int QSharedMemory_Attach(void* ptr, int mode){
	return static_cast<QSharedMemory*>(ptr)->attach(static_cast<QSharedMemory::AccessMode>(mode));
}

void* QSharedMemory_ConstData(void* ptr){
	return const_cast<void*>(static_cast<QSharedMemory*>(ptr)->constData());
}

int QSharedMemory_Create(void* ptr, int size, int mode){
	return static_cast<QSharedMemory*>(ptr)->create(size, static_cast<QSharedMemory::AccessMode>(mode));
}

void* QSharedMemory_Data(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->data();
}

void* QSharedMemory_Data2(void* ptr){
	return const_cast<void*>(static_cast<QSharedMemory*>(ptr)->data());
}

int QSharedMemory_Detach(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->detach();
}

int QSharedMemory_Error(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->error();
}

char* QSharedMemory_ErrorString(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->errorString().toUtf8().data();
}

int QSharedMemory_IsAttached(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->isAttached();
}

char* QSharedMemory_Key(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->key().toUtf8().data();
}

int QSharedMemory_Lock(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->lock();
}

char* QSharedMemory_NativeKey(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->nativeKey().toUtf8().data();
}

void QSharedMemory_SetKey(void* ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setKey(QString(key));
}

void QSharedMemory_SetNativeKey(void* ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setNativeKey(QString(key));
}

int QSharedMemory_Size(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->size();
}

int QSharedMemory_Unlock(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->unlock();
}

void QSharedMemory_DestroyQSharedMemory(void* ptr){
	static_cast<QSharedMemory*>(ptr)->~QSharedMemory();
}

void QSharedMemory_TimerEvent(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSharedMemory_TimerEventDefault(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->QSharedMemory::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSharedMemory_ChildEvent(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSharedMemory_ChildEventDefault(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->QSharedMemory::childEvent(static_cast<QChildEvent*>(event));
}

void QSharedMemory_CustomEvent(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSharedMemory_CustomEventDefault(void* ptr, void* event){
	static_cast<QSharedMemory*>(ptr)->QSharedMemory::customEvent(static_cast<QEvent*>(event));
}

void QSignalBlocker_Reblock(void* ptr){
	static_cast<QSignalBlocker*>(ptr)->reblock();
}

void QSignalBlocker_Unblock(void* ptr){
	static_cast<QSignalBlocker*>(ptr)->unblock();
}

void QSignalBlocker_DestroyQSignalBlocker(void* ptr){
	static_cast<QSignalBlocker*>(ptr)->~QSignalBlocker();
}

class MyQSignalMapper: public QSignalMapper {
public:
	void Signal_Mapped4(QObject * object) { callbackQSignalMapperMapped4(this, this->objectName().toUtf8().data(), object); };
	void Signal_Mapped3(QWidget * widget) { callbackQSignalMapperMapped3(this, this->objectName().toUtf8().data(), widget); };
	void Signal_Mapped2(const QString & text) { callbackQSignalMapperMapped2(this, this->objectName().toUtf8().data(), text.toUtf8().data()); };
	void Signal_Mapped(int i) { callbackQSignalMapperMapped(this, this->objectName().toUtf8().data(), i); };
	void timerEvent(QTimerEvent * event) { callbackQSignalMapperTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSignalMapperChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSignalMapperCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSignalMapper_NewQSignalMapper(void* parent){
	return new QSignalMapper(static_cast<QObject*>(parent));
}

void QSignalMapper_Map(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map");
}

void QSignalMapper_Map2(void* ptr, void* sender){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map", Q_ARG(QObject*, static_cast<QObject*>(sender)));
}

void QSignalMapper_ConnectMapped4(void* ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(QObject *)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(QObject *)>(&MyQSignalMapper::Signal_Mapped4));;
}

void QSignalMapper_DisconnectMapped4(void* ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(QObject *)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(QObject *)>(&MyQSignalMapper::Signal_Mapped4));;
}

void QSignalMapper_Mapped4(void* ptr, void* object){
	static_cast<QSignalMapper*>(ptr)->mapped(static_cast<QObject*>(object));
}

void QSignalMapper_ConnectMapped3(void* ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(QWidget *)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(QWidget *)>(&MyQSignalMapper::Signal_Mapped3));;
}

void QSignalMapper_DisconnectMapped3(void* ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(QWidget *)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(QWidget *)>(&MyQSignalMapper::Signal_Mapped3));;
}

void QSignalMapper_Mapped3(void* ptr, void* widget){
	static_cast<QSignalMapper*>(ptr)->mapped(static_cast<QWidget*>(widget));
}

void QSignalMapper_ConnectMapped2(void* ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(const QString &)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(const QString &)>(&MyQSignalMapper::Signal_Mapped2));;
}

void QSignalMapper_DisconnectMapped2(void* ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(const QString &)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(const QString &)>(&MyQSignalMapper::Signal_Mapped2));;
}

void QSignalMapper_Mapped2(void* ptr, char* text){
	static_cast<QSignalMapper*>(ptr)->mapped(QString(text));
}

void QSignalMapper_ConnectMapped(void* ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

void QSignalMapper_DisconnectMapped(void* ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

void QSignalMapper_Mapped(void* ptr, int i){
	static_cast<QSignalMapper*>(ptr)->mapped(i);
}

void* QSignalMapper_Mapping4(void* ptr, void* object){
	return static_cast<QSignalMapper*>(ptr)->mapping(static_cast<QObject*>(object));
}

void* QSignalMapper_Mapping3(void* ptr, void* widget){
	return static_cast<QSignalMapper*>(ptr)->mapping(static_cast<QWidget*>(widget));
}

void* QSignalMapper_Mapping2(void* ptr, char* id){
	return static_cast<QSignalMapper*>(ptr)->mapping(QString(id));
}

void* QSignalMapper_Mapping(void* ptr, int id){
	return static_cast<QSignalMapper*>(ptr)->mapping(id);
}

void QSignalMapper_RemoveMappings(void* ptr, void* sender){
	static_cast<QSignalMapper*>(ptr)->removeMappings(static_cast<QObject*>(sender));
}

void QSignalMapper_SetMapping4(void* ptr, void* sender, void* object){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), static_cast<QObject*>(object));
}

void QSignalMapper_SetMapping3(void* ptr, void* sender, void* widget){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), static_cast<QWidget*>(widget));
}

void QSignalMapper_SetMapping2(void* ptr, void* sender, char* text){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), QString(text));
}

void QSignalMapper_SetMapping(void* ptr, void* sender, int id){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), id);
}

void QSignalMapper_DestroyQSignalMapper(void* ptr){
	static_cast<QSignalMapper*>(ptr)->~QSignalMapper();
}

void QSignalMapper_TimerEvent(void* ptr, void* event){
	static_cast<MyQSignalMapper*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalMapper_TimerEventDefault(void* ptr, void* event){
	static_cast<QSignalMapper*>(ptr)->QSignalMapper::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalMapper_ChildEvent(void* ptr, void* event){
	static_cast<MyQSignalMapper*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSignalMapper_ChildEventDefault(void* ptr, void* event){
	static_cast<QSignalMapper*>(ptr)->QSignalMapper::childEvent(static_cast<QChildEvent*>(event));
}

void QSignalMapper_CustomEvent(void* ptr, void* event){
	static_cast<MyQSignalMapper*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSignalMapper_CustomEventDefault(void* ptr, void* event){
	static_cast<QSignalMapper*>(ptr)->QSignalMapper::customEvent(static_cast<QEvent*>(event));
}

class MyQSignalTransition: public QSignalTransition {
public:
	MyQSignalTransition(QState *sourceState) : QSignalTransition(sourceState) {};
	MyQSignalTransition(const QObject *sender, const char *signal, QState *sourceState) : QSignalTransition(sender, signal, sourceState) {};
	void onTransition(QEvent * event) { callbackQSignalTransitionOnTransition(this, this->objectName().toUtf8().data(), event); };
	void Signal_SenderObjectChanged() { callbackQSignalTransitionSenderObjectChanged(this, this->objectName().toUtf8().data()); };
	void Signal_SignalChanged() { callbackQSignalTransitionSignalChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSignalTransitionTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSignalTransitionChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSignalTransitionCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSignalTransition_NewQSignalTransition(void* sourceState){
	return new MyQSignalTransition(static_cast<QState*>(sourceState));
}

void* QSignalTransition_NewQSignalTransition2(void* sender, char* signal, void* sourceState){
	return new MyQSignalTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QState*>(sourceState));
}

int QSignalTransition_Event(void* ptr, void* e){
	return static_cast<QSignalTransition*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSignalTransition_EventTest(void* ptr, void* event){
	return static_cast<QSignalTransition*>(ptr)->eventTest(static_cast<QEvent*>(event));
}

void QSignalTransition_OnTransition(void* ptr, void* event){
	static_cast<MyQSignalTransition*>(ptr)->onTransition(static_cast<QEvent*>(event));
}

void QSignalTransition_OnTransitionDefault(void* ptr, void* event){
	static_cast<QSignalTransition*>(ptr)->QSignalTransition::onTransition(static_cast<QEvent*>(event));
}

void* QSignalTransition_SenderObject(void* ptr){
	return static_cast<QSignalTransition*>(ptr)->senderObject();
}

void QSignalTransition_ConnectSenderObjectChanged(void* ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_DisconnectSenderObjectChanged(void* ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_SetSenderObject(void* ptr, void* sender){
	static_cast<QSignalTransition*>(ptr)->setSenderObject(static_cast<QObject*>(sender));
}

void QSignalTransition_SetSignal(void* ptr, void* signal){
	static_cast<QSignalTransition*>(ptr)->setSignal(*static_cast<QByteArray*>(signal));
}

void* QSignalTransition_Signal(void* ptr){
	return new QByteArray(static_cast<QSignalTransition*>(ptr)->signal());
}

void QSignalTransition_ConnectSignalChanged(void* ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DisconnectSignalChanged(void* ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DestroyQSignalTransition(void* ptr){
	static_cast<QSignalTransition*>(ptr)->~QSignalTransition();
}

void QSignalTransition_TimerEvent(void* ptr, void* event){
	static_cast<MyQSignalTransition*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalTransition_TimerEventDefault(void* ptr, void* event){
	static_cast<QSignalTransition*>(ptr)->QSignalTransition::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalTransition_ChildEvent(void* ptr, void* event){
	static_cast<MyQSignalTransition*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSignalTransition_ChildEventDefault(void* ptr, void* event){
	static_cast<QSignalTransition*>(ptr)->QSignalTransition::childEvent(static_cast<QChildEvent*>(event));
}

void QSignalTransition_CustomEvent(void* ptr, void* event){
	static_cast<MyQSignalTransition*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSignalTransition_CustomEventDefault(void* ptr, void* event){
	static_cast<QSignalTransition*>(ptr)->QSignalTransition::customEvent(static_cast<QEvent*>(event));
}

void* QSize_NewQSize(){
	return new QSize();
}

void* QSize_NewQSize2(int width, int height){
	return new QSize(width, height);
}

void* QSize_BoundedTo(void* ptr, void* otherSize){
	return new QSize(static_cast<QSize>(static_cast<QSize*>(ptr)->boundedTo(*static_cast<QSize*>(otherSize))).width(), static_cast<QSize>(static_cast<QSize*>(ptr)->boundedTo(*static_cast<QSize*>(otherSize))).height());
}

void* QSize_ExpandedTo(void* ptr, void* otherSize){
	return new QSize(static_cast<QSize>(static_cast<QSize*>(ptr)->expandedTo(*static_cast<QSize*>(otherSize))).width(), static_cast<QSize>(static_cast<QSize*>(ptr)->expandedTo(*static_cast<QSize*>(otherSize))).height());
}

int QSize_Height(void* ptr){
	return static_cast<QSize*>(ptr)->height();
}

int QSize_IsEmpty(void* ptr){
	return static_cast<QSize*>(ptr)->isEmpty();
}

int QSize_IsNull(void* ptr){
	return static_cast<QSize*>(ptr)->isNull();
}

int QSize_IsValid(void* ptr){
	return static_cast<QSize*>(ptr)->isValid();
}

int QSize_Rheight(void* ptr){
	return static_cast<QSize*>(ptr)->rheight();
}

int QSize_Rwidth(void* ptr){
	return static_cast<QSize*>(ptr)->rwidth();
}

void QSize_Scale2(void* ptr, void* size, int mode){
	static_cast<QSize*>(ptr)->scale(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_Scale(void* ptr, int width, int height, int mode){
	static_cast<QSize*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void* QSize_Scaled2(void* ptr, void* s, int mode){
	return new QSize(static_cast<QSize>(static_cast<QSize*>(ptr)->scaled(*static_cast<QSize*>(s), static_cast<Qt::AspectRatioMode>(mode))).width(), static_cast<QSize>(static_cast<QSize*>(ptr)->scaled(*static_cast<QSize*>(s), static_cast<Qt::AspectRatioMode>(mode))).height());
}

void* QSize_Scaled(void* ptr, int width, int height, int mode){
	return new QSize(static_cast<QSize>(static_cast<QSize*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(mode))).width(), static_cast<QSize>(static_cast<QSize*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(mode))).height());
}

void QSize_SetHeight(void* ptr, int height){
	static_cast<QSize*>(ptr)->setHeight(height);
}

void QSize_SetWidth(void* ptr, int width){
	static_cast<QSize*>(ptr)->setWidth(width);
}

void QSize_Transpose(void* ptr){
	static_cast<QSize*>(ptr)->transpose();
}

void* QSize_Transposed(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSize*>(ptr)->transposed()).width(), static_cast<QSize>(static_cast<QSize*>(ptr)->transposed()).height());
}

int QSize_Width(void* ptr){
	return static_cast<QSize*>(ptr)->width();
}

void* QSizeF_NewQSizeF(){
	return new QSizeF();
}

void* QSizeF_NewQSizeF2(void* size){
	return new QSizeF(*static_cast<QSize*>(size));
}

void* QSizeF_NewQSizeF3(double width, double height){
	return new QSizeF(static_cast<double>(width), static_cast<double>(height));
}

double QSizeF_Height(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->height());
}

int QSizeF_IsEmpty(void* ptr){
	return static_cast<QSizeF*>(ptr)->isEmpty();
}

int QSizeF_IsNull(void* ptr){
	return static_cast<QSizeF*>(ptr)->isNull();
}

int QSizeF_IsValid(void* ptr){
	return static_cast<QSizeF*>(ptr)->isValid();
}

double QSizeF_Rheight(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->rheight());
}

double QSizeF_Rwidth(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->rwidth());
}

void QSizeF_Scale2(void* ptr, void* size, int mode){
	static_cast<QSizeF*>(ptr)->scale(*static_cast<QSizeF*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_Scale(void* ptr, double width, double height, int mode){
	static_cast<QSizeF*>(ptr)->scale(static_cast<double>(width), static_cast<double>(height), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_SetHeight(void* ptr, double height){
	static_cast<QSizeF*>(ptr)->setHeight(static_cast<double>(height));
}

void QSizeF_SetWidth(void* ptr, double width){
	static_cast<QSizeF*>(ptr)->setWidth(static_cast<double>(width));
}

void* QSizeF_ToSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSizeF*>(ptr)->toSize()).width(), static_cast<QSize>(static_cast<QSizeF*>(ptr)->toSize()).height());
}

void QSizeF_Transpose(void* ptr){
	static_cast<QSizeF*>(ptr)->transpose();
}

double QSizeF_Width(void* ptr){
	return static_cast<double>(static_cast<QSizeF*>(ptr)->width());
}

class MyQSocketNotifier: public QSocketNotifier {
public:
	void Signal_Activated(int socket) { callbackQSocketNotifierActivated(this, this->objectName().toUtf8().data(), socket); };
	void timerEvent(QTimerEvent * event) { callbackQSocketNotifierTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSocketNotifierChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSocketNotifierCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSocketNotifier_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

void QSocketNotifier_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

int QSocketNotifier_Event(void* ptr, void* e){
	return static_cast<QSocketNotifier*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSocketNotifier_IsEnabled(void* ptr){
	return static_cast<QSocketNotifier*>(ptr)->isEnabled();
}

void QSocketNotifier_SetEnabled(void* ptr, int enable){
	QMetaObject::invokeMethod(static_cast<QSocketNotifier*>(ptr), "setEnabled", Q_ARG(bool, enable != 0));
}

int QSocketNotifier_Type(void* ptr){
	return static_cast<QSocketNotifier*>(ptr)->type();
}

void QSocketNotifier_DestroyQSocketNotifier(void* ptr){
	static_cast<QSocketNotifier*>(ptr)->~QSocketNotifier();
}

void QSocketNotifier_TimerEvent(void* ptr, void* event){
	static_cast<MyQSocketNotifier*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSocketNotifier_TimerEventDefault(void* ptr, void* event){
	static_cast<QSocketNotifier*>(ptr)->QSocketNotifier::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSocketNotifier_ChildEvent(void* ptr, void* event){
	static_cast<MyQSocketNotifier*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSocketNotifier_ChildEventDefault(void* ptr, void* event){
	static_cast<QSocketNotifier*>(ptr)->QSocketNotifier::childEvent(static_cast<QChildEvent*>(event));
}

void QSocketNotifier_CustomEvent(void* ptr, void* event){
	static_cast<MyQSocketNotifier*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSocketNotifier_CustomEventDefault(void* ptr, void* event){
	static_cast<QSocketNotifier*>(ptr)->QSocketNotifier::customEvent(static_cast<QEvent*>(event));
}

class MyQSortFilterProxyModel: public QSortFilterProxyModel {
public:
	MyQSortFilterProxyModel(QObject *parent) : QSortFilterProxyModel(parent) {};
	void fetchMore(const QModelIndex & parent) { callbackQSortFilterProxyModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void setSourceModel(QAbstractItemModel * sourceModel) { callbackQSortFilterProxyModelSetSourceModel(this, this->objectName().toUtf8().data(), sourceModel); };
	void sort(int column, Qt::SortOrder order) { callbackQSortFilterProxyModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void revert() { if (!callbackQSortFilterProxyModelRevert(this, this->objectName().toUtf8().data())) { QSortFilterProxyModel::revert(); }; };
	void timerEvent(QTimerEvent * event) { callbackQSortFilterProxyModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSortFilterProxyModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSortFilterProxyModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QSortFilterProxyModel_DynamicSortFilter(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->dynamicSortFilter();
}

int QSortFilterProxyModel_FilterCaseSensitivity(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterCaseSensitivity();
}

int QSortFilterProxyModel_FilterKeyColumn(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterKeyColumn();
}

void* QSortFilterProxyModel_FilterRegExp(void* ptr){
	return new QRegExp(static_cast<QSortFilterProxyModel*>(ptr)->filterRegExp());
}

int QSortFilterProxyModel_FilterRole(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterRole();
}

int QSortFilterProxyModel_IsSortLocaleAware(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->isSortLocaleAware();
}

void QSortFilterProxyModel_SetDynamicSortFilter(void* ptr, int enable){
	static_cast<QSortFilterProxyModel*>(ptr)->setDynamicSortFilter(enable != 0);
}

void QSortFilterProxyModel_SetFilterCaseSensitivity(void* ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetFilterKeyColumn(void* ptr, int column){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterKeyColumn(column);
}

void QSortFilterProxyModel_SetFilterRegExp(void* ptr, void* regExp){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRegExp(*static_cast<QRegExp*>(regExp));
}

void QSortFilterProxyModel_SetFilterRole(void* ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRole(role);
}

void QSortFilterProxyModel_SetSortCaseSensitivity(void* ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetSortLocaleAware(void* ptr, int on){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortLocaleAware(on != 0);
}

void QSortFilterProxyModel_SetSortRole(void* ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortRole(role);
}

int QSortFilterProxyModel_SortCaseSensitivity(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortCaseSensitivity();
}

int QSortFilterProxyModel_SortRole(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortRole();
}

void* QSortFilterProxyModel_NewQSortFilterProxyModel(void* parent){
	return new MyQSortFilterProxyModel(static_cast<QObject*>(parent));
}

void* QSortFilterProxyModel_Buddy(void* ptr, void* index){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

int QSortFilterProxyModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QSortFilterProxyModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QSortFilterProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQSortFilterProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_FilterAcceptsColumn(void* ptr, int source_column, void* source_parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterAcceptsColumn(source_column, *static_cast<QModelIndex*>(source_parent));
}

int QSortFilterProxyModel_FilterAcceptsRow(void* ptr, int source_row, void* source_parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterAcceptsRow(source_row, *static_cast<QModelIndex*>(source_parent));
}

int QSortFilterProxyModel_Flags(void* ptr, void* index){
	return static_cast<QSortFilterProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QSortFilterProxyModel_HasChildren(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QSortFilterProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSortFilterProxyModel_Index(void* ptr, int row, int column, void* parent){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QSortFilterProxyModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_Invalidate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "invalidate");
}

int QSortFilterProxyModel_LessThan(void* ptr, void* source_left, void* source_right){
	return static_cast<QSortFilterProxyModel*>(ptr)->lessThan(*static_cast<QModelIndex*>(source_left), *static_cast<QModelIndex*>(source_right));
}

void* QSortFilterProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QSortFilterProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

char* QSortFilterProxyModel_MimeTypes(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->mimeTypes().join(",,,").toUtf8().data();
}

void* QSortFilterProxyModel_Parent(void* ptr, void* child){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)));
}

int QSortFilterProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RowCount(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSortFilterProxyModel_SetFilterFixedString(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterFixedString", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterRegExp2(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterRegExp", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterWildcard(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterWildcard", Q_ARG(QString, QString(pattern)));
}

int QSortFilterProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QSortFilterProxyModel_SetSourceModel(void* ptr, void* sourceModel){
	static_cast<MyQSortFilterProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void QSortFilterProxyModel_SetSourceModelDefault(void* ptr, void* sourceModel){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void* QSortFilterProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QSortFilterProxyModel_Sort(void* ptr, int column, int order){
	static_cast<MyQSortFilterProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSortFilterProxyModel_SortDefault(void* ptr, int column, int order){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::sort(column, static_cast<Qt::SortOrder>(order));
}

int QSortFilterProxyModel_SortColumn(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortColumn();
}

int QSortFilterProxyModel_SortOrder(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortOrder();
}

void* QSortFilterProxyModel_Span(void* ptr, void* index){
	return new QSize(static_cast<QSize>(static_cast<QSortFilterProxyModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).width(), static_cast<QSize>(static_cast<QSortFilterProxyModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).height());
}

int QSortFilterProxyModel_SupportedDropActions(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->supportedDropActions();
}

void QSortFilterProxyModel_DestroyQSortFilterProxyModel(void* ptr){
	static_cast<QSortFilterProxyModel*>(ptr)->~QSortFilterProxyModel();
}

void QSortFilterProxyModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQSortFilterProxyModel*>(ptr), "revert");
}

void QSortFilterProxyModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "revert");
}

void QSortFilterProxyModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQSortFilterProxyModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSortFilterProxyModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSortFilterProxyModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQSortFilterProxyModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSortFilterProxyModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSortFilterProxyModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQSortFilterProxyModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSortFilterProxyModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::customEvent(static_cast<QEvent*>(event));
}

void QStandardPaths_QStandardPaths_SetTestModeEnabled(int testMode){
	QStandardPaths::setTestModeEnabled(testMode != 0);
}

char* QStandardPaths_QStandardPaths_FindExecutable(char* executableName, char* paths){
	return QStandardPaths::findExecutable(QString(executableName), QString(paths).split(",,,", QString::SkipEmptyParts)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_Locate(int ty, char* fileName, int options){
	return QStandardPaths::locate(static_cast<QStandardPaths::StandardLocation>(ty), QString(fileName), static_cast<QStandardPaths::LocateOption>(options)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_LocateAll(int ty, char* fileName, int options){
	return QStandardPaths::locateAll(static_cast<QStandardPaths::StandardLocation>(ty), QString(fileName), static_cast<QStandardPaths::LocateOption>(options)).join(",,,").toUtf8().data();
}

char* QStandardPaths_QStandardPaths_DisplayName(int ty){
	return QStandardPaths::displayName(static_cast<QStandardPaths::StandardLocation>(ty)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_StandardLocations(int ty){
	return QStandardPaths::standardLocations(static_cast<QStandardPaths::StandardLocation>(ty)).join(",,,").toUtf8().data();
}

char* QStandardPaths_QStandardPaths_WritableLocation(int ty){
	return QStandardPaths::writableLocation(static_cast<QStandardPaths::StandardLocation>(ty)).toUtf8().data();
}

class MyQState: public QState {
public:
	MyQState(ChildMode childMode, QState *parent) : QState(childMode, parent) {};
	MyQState(QState *parent) : QState(parent) {};
	void Signal_ChildModeChanged() { callbackQStateChildModeChanged(this, this->objectName().toUtf8().data()); };
	void Signal_ErrorStateChanged() { callbackQStateErrorStateChanged(this, this->objectName().toUtf8().data()); };
	void Signal_Finished() { callbackQStateFinished(this, this->objectName().toUtf8().data()); };
	void Signal_InitialStateChanged() { callbackQStateInitialStateChanged(this, this->objectName().toUtf8().data()); };
	void onEntry(QEvent * event) { callbackQStateOnEntry(this, this->objectName().toUtf8().data(), event); };
	void onExit(QEvent * event) { callbackQStateOnExit(this, this->objectName().toUtf8().data(), event); };
	void Signal_PropertiesAssigned() { callbackQStatePropertiesAssigned(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQStateTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQStateChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQStateCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QState_NewQState2(int childMode, void* parent){
	return new MyQState(static_cast<QState::ChildMode>(childMode), static_cast<QState*>(parent));
}

void* QState_NewQState(void* parent){
	return new MyQState(static_cast<QState*>(parent));
}

void* QState_AddTransition3(void* ptr, void* target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractState*>(target));
}

void* QState_AddTransition2(void* ptr, void* sender, char* signal, void* target){
	return static_cast<QState*>(ptr)->addTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QAbstractState*>(target));
}

void QState_AddTransition(void* ptr, void* transition){
	static_cast<QState*>(ptr)->addTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_AssignProperty(void* ptr, void* object, char* name, void* value){
	static_cast<QState*>(ptr)->assignProperty(static_cast<QObject*>(object), const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

int QState_ChildMode(void* ptr){
	return static_cast<QState*>(ptr)->childMode();
}

void QState_ConnectChildModeChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

void QState_DisconnectChildModeChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::childModeChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ChildModeChanged));;
}

void* QState_ErrorState(void* ptr){
	return static_cast<QState*>(ptr)->errorState();
}

void QState_ConnectErrorStateChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

void QState_DisconnectErrorStateChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::errorStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_ErrorStateChanged));;
}

int QState_Event(void* ptr, void* e){
	return static_cast<QState*>(ptr)->event(static_cast<QEvent*>(e));
}

void QState_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

void QState_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::finished, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_Finished));;
}

void* QState_InitialState(void* ptr){
	return static_cast<QState*>(ptr)->initialState();
}

void QState_ConnectInitialStateChanged(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_DisconnectInitialStateChanged(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::initialStateChanged, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_InitialStateChanged));;
}

void QState_OnEntry(void* ptr, void* event){
	static_cast<MyQState*>(ptr)->onEntry(static_cast<QEvent*>(event));
}

void QState_OnEntryDefault(void* ptr, void* event){
	static_cast<QState*>(ptr)->QState::onEntry(static_cast<QEvent*>(event));
}

void QState_OnExit(void* ptr, void* event){
	static_cast<MyQState*>(ptr)->onExit(static_cast<QEvent*>(event));
}

void QState_OnExitDefault(void* ptr, void* event){
	static_cast<QState*>(ptr)->QState::onExit(static_cast<QEvent*>(event));
}

void QState_ConnectPropertiesAssigned(void* ptr){
	QObject::connect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_DisconnectPropertiesAssigned(void* ptr){
	QObject::disconnect(static_cast<QState*>(ptr), &QState::propertiesAssigned, static_cast<MyQState*>(ptr), static_cast<void (MyQState::*)()>(&MyQState::Signal_PropertiesAssigned));;
}

void QState_RemoveTransition(void* ptr, void* transition){
	static_cast<QState*>(ptr)->removeTransition(static_cast<QAbstractTransition*>(transition));
}

void QState_SetChildMode(void* ptr, int mode){
	static_cast<QState*>(ptr)->setChildMode(static_cast<QState::ChildMode>(mode));
}

void QState_SetErrorState(void* ptr, void* state){
	static_cast<QState*>(ptr)->setErrorState(static_cast<QAbstractState*>(state));
}

void QState_SetInitialState(void* ptr, void* state){
	static_cast<QState*>(ptr)->setInitialState(static_cast<QAbstractState*>(state));
}

void QState_DestroyQState(void* ptr){
	static_cast<QState*>(ptr)->~QState();
}

void QState_TimerEvent(void* ptr, void* event){
	static_cast<MyQState*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QState_TimerEventDefault(void* ptr, void* event){
	static_cast<QState*>(ptr)->QState::timerEvent(static_cast<QTimerEvent*>(event));
}

void QState_ChildEvent(void* ptr, void* event){
	static_cast<MyQState*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QState_ChildEventDefault(void* ptr, void* event){
	static_cast<QState*>(ptr)->QState::childEvent(static_cast<QChildEvent*>(event));
}

void QState_CustomEvent(void* ptr, void* event){
	static_cast<MyQState*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QState_CustomEventDefault(void* ptr, void* event){
	static_cast<QState*>(ptr)->QState::customEvent(static_cast<QEvent*>(event));
}

class MyQStateMachine: public QStateMachine {
public:
	MyQStateMachine(QObject *parent) : QStateMachine(parent) {};
	MyQStateMachine(QState::ChildMode childMode, QObject *parent) : QStateMachine(childMode, parent) {};
	void onEntry(QEvent * event) { callbackQStateMachineOnEntry(this, this->objectName().toUtf8().data(), event); };
	void onExit(QEvent * event) { callbackQStateMachineOnExit(this, this->objectName().toUtf8().data(), event); };
	void Signal_RunningChanged(bool running) { callbackQStateMachineRunningChanged(this, this->objectName().toUtf8().data(), running); };
	void Signal_Started() { callbackQStateMachineStarted(this, this->objectName().toUtf8().data()); };
	void Signal_Stopped() { callbackQStateMachineStopped(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQStateMachineTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQStateMachineChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQStateMachineCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QStateMachine_NewQStateMachine(void* parent){
	return new MyQStateMachine(static_cast<QObject*>(parent));
}

void* QStateMachine_NewQStateMachine2(int childMode, void* parent){
	return new MyQStateMachine(static_cast<QState::ChildMode>(childMode), static_cast<QObject*>(parent));
}

void QStateMachine_AddDefaultAnimation(void* ptr, void* animation){
	static_cast<QStateMachine*>(ptr)->addDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_AddState(void* ptr, void* state){
	static_cast<QStateMachine*>(ptr)->addState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ClearError(void* ptr){
	static_cast<QStateMachine*>(ptr)->clearError();
}

int QStateMachine_CancelDelayedEvent(void* ptr, int id){
	return static_cast<QStateMachine*>(ptr)->cancelDelayedEvent(id);
}

int QStateMachine_Error(void* ptr){
	return static_cast<QStateMachine*>(ptr)->error();
}

char* QStateMachine_ErrorString(void* ptr){
	return static_cast<QStateMachine*>(ptr)->errorString().toUtf8().data();
}

int QStateMachine_Event(void* ptr, void* e){
	return static_cast<QStateMachine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QStateMachine_EventFilter(void* ptr, void* watched, void* event){
	return static_cast<QStateMachine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QStateMachine_GlobalRestorePolicy(void* ptr){
	return static_cast<QStateMachine*>(ptr)->globalRestorePolicy();
}

int QStateMachine_IsAnimated(void* ptr){
	return static_cast<QStateMachine*>(ptr)->isAnimated();
}

int QStateMachine_IsRunning(void* ptr){
	return static_cast<QStateMachine*>(ptr)->isRunning();
}

void QStateMachine_OnEntry(void* ptr, void* event){
	static_cast<MyQStateMachine*>(ptr)->onEntry(static_cast<QEvent*>(event));
}

void QStateMachine_OnEntryDefault(void* ptr, void* event){
	static_cast<QStateMachine*>(ptr)->QStateMachine::onEntry(static_cast<QEvent*>(event));
}

void QStateMachine_OnExit(void* ptr, void* event){
	static_cast<MyQStateMachine*>(ptr)->onExit(static_cast<QEvent*>(event));
}

void QStateMachine_OnExitDefault(void* ptr, void* event){
	static_cast<QStateMachine*>(ptr)->QStateMachine::onExit(static_cast<QEvent*>(event));
}

int QStateMachine_PostDelayedEvent(void* ptr, void* event, int delay){
	return static_cast<QStateMachine*>(ptr)->postDelayedEvent(static_cast<QEvent*>(event), delay);
}

void QStateMachine_PostEvent(void* ptr, void* event, int priority){
	static_cast<QStateMachine*>(ptr)->postEvent(static_cast<QEvent*>(event), static_cast<QStateMachine::EventPriority>(priority));
}

void QStateMachine_RemoveDefaultAnimation(void* ptr, void* animation){
	static_cast<QStateMachine*>(ptr)->removeDefaultAnimation(static_cast<QAbstractAnimation*>(animation));
}

void QStateMachine_RemoveState(void* ptr, void* state){
	static_cast<QStateMachine*>(ptr)->removeState(static_cast<QAbstractState*>(state));
}

void QStateMachine_ConnectRunningChanged(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_DisconnectRunningChanged(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), static_cast<void (QStateMachine::*)(bool)>(&QStateMachine::runningChanged), static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)(bool)>(&MyQStateMachine::Signal_RunningChanged));;
}

void QStateMachine_RunningChanged(void* ptr, int running){
	static_cast<QStateMachine*>(ptr)->runningChanged(running != 0);
}

void QStateMachine_SetAnimated(void* ptr, int enabled){
	static_cast<QStateMachine*>(ptr)->setAnimated(enabled != 0);
}

void QStateMachine_SetGlobalRestorePolicy(void* ptr, int restorePolicy){
	static_cast<QStateMachine*>(ptr)->setGlobalRestorePolicy(static_cast<QState::RestorePolicy>(restorePolicy));
}

void QStateMachine_SetRunning(void* ptr, int running){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "setRunning", Q_ARG(bool, running != 0));
}

void QStateMachine_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "start");
}

void QStateMachine_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::started, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Started));;
}

void QStateMachine_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStateMachine*>(ptr), "stop");
}

void QStateMachine_ConnectStopped(void* ptr){
	QObject::connect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DisconnectStopped(void* ptr){
	QObject::disconnect(static_cast<QStateMachine*>(ptr), &QStateMachine::stopped, static_cast<MyQStateMachine*>(ptr), static_cast<void (MyQStateMachine::*)()>(&MyQStateMachine::Signal_Stopped));;
}

void QStateMachine_DestroyQStateMachine(void* ptr){
	static_cast<QStateMachine*>(ptr)->~QStateMachine();
}

void QStateMachine_TimerEvent(void* ptr, void* event){
	static_cast<MyQStateMachine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QStateMachine_TimerEventDefault(void* ptr, void* event){
	static_cast<QStateMachine*>(ptr)->QStateMachine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QStateMachine_ChildEvent(void* ptr, void* event){
	static_cast<MyQStateMachine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QStateMachine_ChildEventDefault(void* ptr, void* event){
	static_cast<QStateMachine*>(ptr)->QStateMachine::childEvent(static_cast<QChildEvent*>(event));
}

void QStateMachine_CustomEvent(void* ptr, void* event){
	static_cast<MyQStateMachine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QStateMachine_CustomEventDefault(void* ptr, void* event){
	static_cast<QStateMachine*>(ptr)->QStateMachine::customEvent(static_cast<QEvent*>(event));
}

void* QStaticPlugin_Instance(void* ptr){
	return static_cast<QStaticPlugin*>(ptr)->instance();
}

void* QStaticPlugin_MetaData(void* ptr){
	return new QJsonObject(static_cast<QStaticPlugin*>(ptr)->metaData());
}

void* QStorageInfo_NewQStorageInfo(){
	return new QStorageInfo();
}

void* QStorageInfo_NewQStorageInfo3(void* dir){
	return new QStorageInfo(*static_cast<QDir*>(dir));
}

void* QStorageInfo_NewQStorageInfo4(void* other){
	return new QStorageInfo(*static_cast<QStorageInfo*>(other));
}

void* QStorageInfo_NewQStorageInfo2(char* path){
	return new QStorageInfo(QString(path));
}

long long QStorageInfo_BytesAvailable(void* ptr){
	return static_cast<long long>(static_cast<QStorageInfo*>(ptr)->bytesAvailable());
}

long long QStorageInfo_BytesFree(void* ptr){
	return static_cast<long long>(static_cast<QStorageInfo*>(ptr)->bytesFree());
}

long long QStorageInfo_BytesTotal(void* ptr){
	return static_cast<long long>(static_cast<QStorageInfo*>(ptr)->bytesTotal());
}

void* QStorageInfo_Device(void* ptr){
	return new QByteArray(static_cast<QStorageInfo*>(ptr)->device());
}

char* QStorageInfo_DisplayName(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->displayName().toUtf8().data();
}

void* QStorageInfo_FileSystemType(void* ptr){
	return new QByteArray(static_cast<QStorageInfo*>(ptr)->fileSystemType());
}

int QStorageInfo_IsReadOnly(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isReadOnly();
}

int QStorageInfo_IsReady(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isReady();
}

int QStorageInfo_IsRoot(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isRoot();
}

int QStorageInfo_IsValid(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isValid();
}

char* QStorageInfo_Name(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->name().toUtf8().data();
}

void QStorageInfo_Refresh(void* ptr){
	static_cast<QStorageInfo*>(ptr)->refresh();
}

char* QStorageInfo_RootPath(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->rootPath().toUtf8().data();
}

void QStorageInfo_SetPath(void* ptr, char* path){
	static_cast<QStorageInfo*>(ptr)->setPath(QString(path));
}

void QStorageInfo_Swap(void* ptr, void* other){
	static_cast<QStorageInfo*>(ptr)->swap(*static_cast<QStorageInfo*>(other));
}

void QStorageInfo_DestroyQStorageInfo(void* ptr){
	static_cast<QStorageInfo*>(ptr)->~QStorageInfo();
}

class MyQStringListModel: public QStringListModel {
public:
	void sort(int column, Qt::SortOrder order) { callbackQStringListModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void fetchMore(const QModelIndex & parent) { callbackQStringListModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void revert() { if (!callbackQStringListModelRevert(this, this->objectName().toUtf8().data())) { QStringListModel::revert(); }; };
	void timerEvent(QTimerEvent * event) { callbackQStringListModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQStringListModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQStringListModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QStringListModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QStringListModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QStringListModel_Flags(void* ptr, void* index){
	return static_cast<QStringListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStringListModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStringListModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStringListModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RowCount(void* ptr, void* parent){
	return static_cast<QStringListModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QStringListModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QStringListModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QStringListModel_SetStringList(void* ptr, char* strin){
	static_cast<QStringListModel*>(ptr)->setStringList(QString(strin).split(",,,", QString::SkipEmptyParts));
}

void* QStringListModel_Sibling(void* ptr, int row, int column, void* idx){
	return new QModelIndex(static_cast<QStringListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QStringListModel_Sort(void* ptr, int column, int order){
	static_cast<MyQStringListModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QStringListModel_SortDefault(void* ptr, int column, int order){
	static_cast<QStringListModel*>(ptr)->QStringListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

char* QStringListModel_StringList(void* ptr){
	return static_cast<QStringListModel*>(ptr)->stringList().join(",,,").toUtf8().data();
}

int QStringListModel_SupportedDropActions(void* ptr){
	return static_cast<QStringListModel*>(ptr)->supportedDropActions();
}

void QStringListModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQStringListModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QStringListModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QStringListModel*>(ptr)->QStringListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QStringListModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQStringListModel*>(ptr), "revert");
}

void QStringListModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStringListModel*>(ptr), "revert");
}

void QStringListModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQStringListModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QStringListModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QStringListModel*>(ptr)->QStringListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QStringListModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQStringListModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QStringListModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QStringListModel*>(ptr)->QStringListModel::childEvent(static_cast<QChildEvent*>(event));
}

void QStringListModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQStringListModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QStringListModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QStringListModel*>(ptr)->QStringListModel::customEvent(static_cast<QEvent*>(event));
}

void* QStringMatcher_NewQStringMatcher3(void* uc, int length, int cs){
	return new QStringMatcher(static_cast<QChar*>(uc), length, static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringMatcher_Pattern(void* ptr){
	return static_cast<QStringMatcher*>(ptr)->pattern().toUtf8().data();
}

void* QStringMatcher_NewQStringMatcher(){
	return new QStringMatcher();
}

void* QStringMatcher_NewQStringMatcher2(char* pattern, int cs){
	return new QStringMatcher(QString(pattern), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringMatcher_NewQStringMatcher4(void* other){
	return new QStringMatcher(*static_cast<QStringMatcher*>(other));
}

int QStringMatcher_CaseSensitivity(void* ptr){
	return static_cast<QStringMatcher*>(ptr)->caseSensitivity();
}

int QStringMatcher_IndexIn2(void* ptr, void* str, int length, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(static_cast<QChar*>(str), length, from);
}

int QStringMatcher_IndexIn(void* ptr, char* str, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(QString(str), from);
}

void QStringMatcher_SetCaseSensitivity(void* ptr, int cs){
	static_cast<QStringMatcher*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QStringMatcher_SetPattern(void* ptr, char* pattern){
	static_cast<QStringMatcher*>(ptr)->setPattern(QString(pattern));
}

void QStringMatcher_DestroyQStringMatcher(void* ptr){
	static_cast<QStringMatcher*>(ptr)->~QStringMatcher();
}

void* QStringRef_Left(void* ptr, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->left(n));
}

void* QStringRef_Mid(void* ptr, int position, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->mid(position, n));
}

void* QStringRef_Right(void* ptr, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->right(n));
}

void* QStringRef_AppendTo(void* ptr, char* stri){
	return new QStringRef(static_cast<QStringRef*>(ptr)->appendTo(new QString(stri)));
}

void* QStringRef_Begin(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->begin());
}

void* QStringRef_Cbegin(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cbegin());
}

void* QStringRef_Cend(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cend());
}

void QStringRef_Clear(void* ptr){
	static_cast<QStringRef*>(ptr)->clear();
}

int QStringRef_QStringRef_Compare3(void* s1, void* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QLatin1String*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare(void* s1, char* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), QString(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare2(void* s1, void* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare6(void* ptr, void* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QLatin1String*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare4(void* ptr, char* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(QString(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare5(void* ptr, void* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QStringRef*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_ConstData(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->constData());
}

int QStringRef_Contains2(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count(void* ptr){
	return static_cast<QStringRef*>(ptr)->count();
}

int QStringRef_Count3(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count2(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->count(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_Data(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

void* QStringRef_End(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->end());
}

int QStringRef_EndsWith2(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf3(void* ptr, void* ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf2(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf(void* ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf4(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IsEmpty(void* ptr){
	return static_cast<QStringRef*>(ptr)->isEmpty();
}

int QStringRef_IsNull(void* ptr){
	return static_cast<QStringRef*>(ptr)->isNull();
}

int QStringRef_LastIndexOf2(void* ptr, void* ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf3(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf(void* ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf4(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Length(void* ptr){
	return static_cast<QStringRef*>(ptr)->length();
}

int QStringRef_QStringRef_LocaleAwareCompare(void* s1, char* s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), QString(s2));
}

int QStringRef_QStringRef_LocaleAwareCompare2(void* s1, void* s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

int QStringRef_LocaleAwareCompare3(void* ptr, char* other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(QString(other));
}

int QStringRef_LocaleAwareCompare4(void* ptr, void* other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(*static_cast<QStringRef*>(other));
}

int QStringRef_Position(void* ptr){
	return static_cast<QStringRef*>(ptr)->position();
}

int QStringRef_Size(void* ptr){
	return static_cast<QStringRef*>(ptr)->size();
}

int QStringRef_StartsWith4(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith2(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringRef_String(void* ptr){
	return static_cast<QStringRef*>(ptr)->string()->toUtf8().data();
}

int QStringRef_ToInt(void* ptr, int ok, int base){
	return static_cast<QStringRef*>(ptr)->toInt(NULL, base);
}

void* QStringRef_ToLatin1(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toLatin1());
}

void* QStringRef_ToLocal8Bit(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toLocal8Bit());
}

char* QStringRef_ToString(void* ptr){
	return static_cast<QStringRef*>(ptr)->toString().toUtf8().data();
}

void* QStringRef_ToUtf8(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toUtf8());
}

void* QStringRef_Trimmed(void* ptr){
	return new QStringRef(static_cast<QStringRef*>(ptr)->trimmed());
}

void* QStringRef_Unicode(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->unicode());
}

void QStringRef_DestroyQStringRef(void* ptr){
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

int QSysInfo_ByteOrder_Type(){
	return QSysInfo::ByteOrder;
}

int QSysInfo_MV_IOS_Type(){
	return QSysInfo::MV_IOS;
}

int QSysInfo_MV_IOS_4_3_Type(){
	return QSysInfo::MV_IOS_4_3;
}

int QSysInfo_MV_IOS_5_0_Type(){
	return QSysInfo::MV_IOS_5_0;
}

int QSysInfo_MV_IOS_5_1_Type(){
	return QSysInfo::MV_IOS_5_1;
}

int QSysInfo_MV_IOS_6_0_Type(){
	return QSysInfo::MV_IOS_6_0;
}

int QSysInfo_MV_IOS_6_1_Type(){
	return QSysInfo::MV_IOS_6_1;
}

int QSysInfo_MV_IOS_7_0_Type(){
	return QSysInfo::MV_IOS_7_0;
}

int QSysInfo_MV_IOS_7_1_Type(){
	return QSysInfo::MV_IOS_7_1;
}

int QSysInfo_MV_IOS_8_0_Type(){
	return QSysInfo::MV_IOS_8_0;
}

int QSysInfo_MV_IOS_8_1_Type(){
	return QSysInfo::MV_IOS_8_1;
}

int QSysInfo_MV_IOS_8_2_Type(){
	return QSysInfo::MV_IOS_8_2;
}

int QSysInfo_MV_IOS_8_3_Type(){
	return QSysInfo::MV_IOS_8_3;
}

int QSysInfo_MV_IOS_8_4_Type(){
	return QSysInfo::MV_IOS_8_4;
}

int QSysInfo_MV_IOS_9_0_Type(){
	return QSysInfo::MV_IOS_9_0;
}

int QSysInfo_WordSize_Type(){
	return QSysInfo::WordSize;
}

int QSysInfo_QSysInfo_MacVersion(){
	return QSysInfo::macVersion();
}

char* QSysInfo_QSysInfo_BuildAbi(){
	return QSysInfo::buildAbi().toUtf8().data();
}

char* QSysInfo_QSysInfo_BuildCpuArchitecture(){
	return QSysInfo::buildCpuArchitecture().toUtf8().data();
}

char* QSysInfo_QSysInfo_CurrentCpuArchitecture(){
	return QSysInfo::currentCpuArchitecture().toUtf8().data();
}

char* QSysInfo_QSysInfo_KernelType(){
	return QSysInfo::kernelType().toUtf8().data();
}

char* QSysInfo_QSysInfo_KernelVersion(){
	return QSysInfo::kernelVersion().toUtf8().data();
}

char* QSysInfo_QSysInfo_PrettyProductName(){
	return QSysInfo::prettyProductName().toUtf8().data();
}

char* QSysInfo_QSysInfo_ProductType(){
	return QSysInfo::productType().toUtf8().data();
}

char* QSysInfo_QSysInfo_ProductVersion(){
	return QSysInfo::productVersion().toUtf8().data();
}

int QSysInfo_QSysInfo_WindowsVersion(){
	return QSysInfo::windowsVersion();
}

void* QSystemSemaphore_NewQSystemSemaphore(char* key, int initialValue, int mode){
	return new QSystemSemaphore(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

int QSystemSemaphore_Acquire(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->acquire();
}

int QSystemSemaphore_Error(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->error();
}

char* QSystemSemaphore_ErrorString(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->errorString().toUtf8().data();
}

char* QSystemSemaphore_Key(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->key().toUtf8().data();
}

int QSystemSemaphore_Release(void* ptr, int n){
	return static_cast<QSystemSemaphore*>(ptr)->release(n);
}

void QSystemSemaphore_SetKey(void* ptr, char* key, int initialValue, int mode){
	static_cast<QSystemSemaphore*>(ptr)->setKey(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

void QSystemSemaphore_DestroyQSystemSemaphore(void* ptr){
	static_cast<QSystemSemaphore*>(ptr)->~QSystemSemaphore();
}

void* QTemporaryDir_NewQTemporaryDir(){
	return new QTemporaryDir();
}

void* QTemporaryDir_NewQTemporaryDir2(char* templatePath){
	return new QTemporaryDir(QString(templatePath));
}

int QTemporaryDir_AutoRemove(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->autoRemove();
}

int QTemporaryDir_IsValid(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->isValid();
}

char* QTemporaryDir_Path(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->path().toUtf8().data();
}

void QTemporaryDir_SetAutoRemove(void* ptr, int b){
	static_cast<QTemporaryDir*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryDir_DestroyQTemporaryDir(void* ptr){
	static_cast<QTemporaryDir*>(ptr)->~QTemporaryDir();
}

class MyQTemporaryFile: public QTemporaryFile {
public:
	MyQTemporaryFile() : QTemporaryFile() {};
	MyQTemporaryFile(QObject *parent) : QTemporaryFile(parent) {};
	MyQTemporaryFile(const QString &templateName) : QTemporaryFile(templateName) {};
	MyQTemporaryFile(const QString &templateName, QObject *parent) : QTemporaryFile(templateName, parent) {};
	void close() { callbackQTemporaryFileClose(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQTemporaryFileTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTemporaryFileChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTemporaryFileCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QTemporaryFile_NewQTemporaryFile(){
	return new MyQTemporaryFile();
}

void* QTemporaryFile_NewQTemporaryFile3(void* parent){
	return new MyQTemporaryFile(static_cast<QObject*>(parent));
}

void* QTemporaryFile_NewQTemporaryFile2(char* templateName){
	return new MyQTemporaryFile(QString(templateName));
}

void* QTemporaryFile_NewQTemporaryFile4(char* templateName, void* parent){
	return new MyQTemporaryFile(QString(templateName), static_cast<QObject*>(parent));
}

int QTemporaryFile_AutoRemove(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->autoRemove();
}

void* QTemporaryFile_QTemporaryFile_CreateNativeFile(void* file){
	return QTemporaryFile::createNativeFile(*static_cast<QFile*>(file));
}

void* QTemporaryFile_QTemporaryFile_CreateNativeFile2(char* fileName){
	return QTemporaryFile::createNativeFile(QString(fileName));
}

char* QTemporaryFile_FileName(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileName().toUtf8().data();
}

char* QTemporaryFile_FileTemplate(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileTemplate().toUtf8().data();
}

int QTemporaryFile_Open(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->open();
}

void QTemporaryFile_SetAutoRemove(void* ptr, int b){
	static_cast<QTemporaryFile*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryFile_SetFileTemplate(void* ptr, char* name){
	static_cast<QTemporaryFile*>(ptr)->setFileTemplate(QString(name));
}

void QTemporaryFile_DestroyQTemporaryFile(void* ptr){
	static_cast<QTemporaryFile*>(ptr)->~QTemporaryFile();
}

void QTemporaryFile_Close(void* ptr){
	static_cast<MyQTemporaryFile*>(ptr)->close();
}

void QTemporaryFile_CloseDefault(void* ptr){
	static_cast<QTemporaryFile*>(ptr)->QTemporaryFile::close();
}

void QTemporaryFile_TimerEvent(void* ptr, void* event){
	static_cast<MyQTemporaryFile*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTemporaryFile_TimerEventDefault(void* ptr, void* event){
	static_cast<QTemporaryFile*>(ptr)->QTemporaryFile::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTemporaryFile_ChildEvent(void* ptr, void* event){
	static_cast<MyQTemporaryFile*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTemporaryFile_ChildEventDefault(void* ptr, void* event){
	static_cast<QTemporaryFile*>(ptr)->QTemporaryFile::childEvent(static_cast<QChildEvent*>(event));
}

void QTemporaryFile_CustomEvent(void* ptr, void* event){
	static_cast<MyQTemporaryFile*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTemporaryFile_CustomEventDefault(void* ptr, void* event){
	static_cast<QTemporaryFile*>(ptr)->QTemporaryFile::customEvent(static_cast<QEvent*>(event));
}

void* QTextBoundaryFinder_NewQTextBoundaryFinder(){
	return new QTextBoundaryFinder();
}

void* QTextBoundaryFinder_NewQTextBoundaryFinder3(int ty, char* stri){
	return new QTextBoundaryFinder(static_cast<QTextBoundaryFinder::BoundaryType>(ty), QString(stri));
}

void* QTextBoundaryFinder_NewQTextBoundaryFinder2(void* other){
	return new QTextBoundaryFinder(*static_cast<QTextBoundaryFinder*>(other));
}

int QTextBoundaryFinder_BoundaryReasons(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->boundaryReasons();
}

int QTextBoundaryFinder_IsAtBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isAtBoundary();
}

int QTextBoundaryFinder_IsValid(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isValid();
}

int QTextBoundaryFinder_Position(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->position();
}

void QTextBoundaryFinder_SetPosition(void* ptr, int position){
	static_cast<QTextBoundaryFinder*>(ptr)->setPosition(position);
}

char* QTextBoundaryFinder_String(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->string().toUtf8().data();
}

void QTextBoundaryFinder_ToEnd(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toEnd();
}

int QTextBoundaryFinder_ToNextBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toNextBoundary();
}

int QTextBoundaryFinder_ToPreviousBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toPreviousBoundary();
}

void QTextBoundaryFinder_ToStart(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toStart();
}

int QTextBoundaryFinder_Type(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->type();
}

void QTextBoundaryFinder_DestroyQTextBoundaryFinder(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->~QTextBoundaryFinder();
}

class MyQTextCodec: public QTextCodec {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QTextCodec_CanEncode(void* ptr, void* ch){
	return static_cast<QTextCodec*>(ptr)->canEncode(*static_cast<QChar*>(ch));
}

int QTextCodec_CanEncode2(void* ptr, char* s){
	return static_cast<QTextCodec*>(ptr)->canEncode(QString(s));
}

void* QTextCodec_QTextCodec_CodecForHtml2(void* ba){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba));
}

void* QTextCodec_QTextCodec_CodecForHtml(void* ba, void* defaultCodec){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

void* QTextCodec_QTextCodec_CodecForLocale(){
	return QTextCodec::codecForLocale();
}

void* QTextCodec_QTextCodec_CodecForMib(int mib){
	return QTextCodec::codecForMib(mib);
}

void* QTextCodec_QTextCodec_CodecForName(void* name){
	return QTextCodec::codecForName(*static_cast<QByteArray*>(name));
}

void* QTextCodec_QTextCodec_CodecForName2(char* name){
	return QTextCodec::codecForName(const_cast<const char*>(name));
}

void* QTextCodec_QTextCodec_CodecForUtfText2(void* ba){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba));
}

void* QTextCodec_QTextCodec_CodecForUtfText(void* ba, void* defaultCodec){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

void* QTextCodec_FromUnicode(void* ptr, char* str){
	return new QByteArray(static_cast<QTextCodec*>(ptr)->fromUnicode(QString(str)));
}

void* QTextCodec_MakeDecoder(void* ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeDecoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

void* QTextCodec_MakeEncoder(void* ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeEncoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

int QTextCodec_MibEnum(void* ptr){
	return static_cast<QTextCodec*>(ptr)->mibEnum();
}

void* QTextCodec_Name(void* ptr){
	return new QByteArray(static_cast<QTextCodec*>(ptr)->name());
}

void QTextCodec_QTextCodec_SetCodecForLocale(void* c){
	QTextCodec::setCodecForLocale(static_cast<QTextCodec*>(c));
}

void QTextCodec_DestroyQTextCodec(void* ptr){
	static_cast<QTextCodec*>(ptr)->~QTextCodec();
}

char* QTextCodec_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQTextCodec*>(static_cast<QTextCodec*>(ptr))) {
		return static_cast<MyQTextCodec*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTextCodec_BASE").toUtf8().data();
}

void QTextCodec_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQTextCodec*>(static_cast<QTextCodec*>(ptr))) {
		static_cast<MyQTextCodec*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QTextDecoder_NewQTextDecoder(void* codec){
	return new QTextDecoder(static_cast<QTextCodec*>(codec));
}

void* QTextDecoder_NewQTextDecoder2(void* codec, int flags){
	return new QTextDecoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void QTextDecoder_DestroyQTextDecoder(void* ptr){
	static_cast<QTextDecoder*>(ptr)->~QTextDecoder();
}

void* QTextEncoder_NewQTextEncoder(void* codec){
	return new QTextEncoder(static_cast<QTextCodec*>(codec));
}

void* QTextEncoder_NewQTextEncoder2(void* codec, int flags){
	return new QTextEncoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void* QTextEncoder_FromUnicode2(void* ptr, void* uc, int len){
	return new QByteArray(static_cast<QTextEncoder*>(ptr)->fromUnicode(static_cast<QChar*>(uc), len));
}

void* QTextEncoder_FromUnicode(void* ptr, char* str){
	return new QByteArray(static_cast<QTextEncoder*>(ptr)->fromUnicode(QString(str)));
}

void QTextEncoder_DestroyQTextEncoder(void* ptr){
	static_cast<QTextEncoder*>(ptr)->~QTextEncoder();
}

class MyQTextStream: public QTextStream {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QTextStream_AtEnd(void* ptr){
	return static_cast<QTextStream*>(ptr)->atEnd();
}

int QTextStream_AutoDetectUnicode(void* ptr){
	return static_cast<QTextStream*>(ptr)->autoDetectUnicode();
}

void* QTextStream_Codec(void* ptr){
	return static_cast<QTextStream*>(ptr)->codec();
}

void* QTextStream_Device(void* ptr){
	return static_cast<QTextStream*>(ptr)->device();
}

int QTextStream_FieldAlignment(void* ptr){
	return static_cast<QTextStream*>(ptr)->fieldAlignment();
}

int QTextStream_FieldWidth(void* ptr){
	return static_cast<QTextStream*>(ptr)->fieldWidth();
}

void QTextStream_Flush(void* ptr){
	static_cast<QTextStream*>(ptr)->flush();
}

int QTextStream_GenerateByteOrderMark(void* ptr){
	return static_cast<QTextStream*>(ptr)->generateByteOrderMark();
}

int QTextStream_IntegerBase(void* ptr){
	return static_cast<QTextStream*>(ptr)->integerBase();
}

int QTextStream_NumberFlags(void* ptr){
	return static_cast<QTextStream*>(ptr)->numberFlags();
}

long long QTextStream_Pos(void* ptr){
	return static_cast<long long>(static_cast<QTextStream*>(ptr)->pos());
}

char* QTextStream_Read(void* ptr, long long maxlen){
	return static_cast<QTextStream*>(ptr)->read(static_cast<long long>(maxlen)).toUtf8().data();
}

char* QTextStream_ReadAll(void* ptr){
	return static_cast<QTextStream*>(ptr)->readAll().toUtf8().data();
}

char* QTextStream_ReadLine(void* ptr, long long maxlen){
	return static_cast<QTextStream*>(ptr)->readLine(static_cast<long long>(maxlen)).toUtf8().data();
}

int QTextStream_ReadLineInto(void* ptr, char* line, long long maxlen){
	return static_cast<QTextStream*>(ptr)->readLineInto(new QString(line), static_cast<long long>(maxlen));
}

int QTextStream_RealNumberNotation(void* ptr){
	return static_cast<QTextStream*>(ptr)->realNumberNotation();
}

int QTextStream_RealNumberPrecision(void* ptr){
	return static_cast<QTextStream*>(ptr)->realNumberPrecision();
}

void QTextStream_Reset(void* ptr){
	static_cast<QTextStream*>(ptr)->reset();
}

void QTextStream_ResetStatus(void* ptr){
	static_cast<QTextStream*>(ptr)->resetStatus();
}

int QTextStream_Seek(void* ptr, long long pos){
	return static_cast<QTextStream*>(ptr)->seek(static_cast<long long>(pos));
}

void QTextStream_SetAutoDetectUnicode(void* ptr, int enabled){
	static_cast<QTextStream*>(ptr)->setAutoDetectUnicode(enabled != 0);
}

void QTextStream_SetCodec(void* ptr, void* codec){
	static_cast<QTextStream*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextStream_SetCodec2(void* ptr, char* codecName){
	static_cast<QTextStream*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QTextStream_SetDevice(void* ptr, void* device){
	static_cast<QTextStream*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextStream_SetFieldAlignment(void* ptr, int mode){
	static_cast<QTextStream*>(ptr)->setFieldAlignment(static_cast<QTextStream::FieldAlignment>(mode));
}

void QTextStream_SetFieldWidth(void* ptr, int width){
	static_cast<QTextStream*>(ptr)->setFieldWidth(width);
}

void QTextStream_SetGenerateByteOrderMark(void* ptr, int generate){
	static_cast<QTextStream*>(ptr)->setGenerateByteOrderMark(generate != 0);
}

void QTextStream_SetIntegerBase(void* ptr, int base){
	static_cast<QTextStream*>(ptr)->setIntegerBase(base);
}

void QTextStream_SetLocale(void* ptr, void* locale){
	static_cast<QTextStream*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QTextStream_SetNumberFlags(void* ptr, int flags){
	static_cast<QTextStream*>(ptr)->setNumberFlags(static_cast<QTextStream::NumberFlag>(flags));
}

void QTextStream_SetPadChar(void* ptr, void* ch){
	static_cast<QTextStream*>(ptr)->setPadChar(*static_cast<QChar*>(ch));
}

void QTextStream_SetRealNumberNotation(void* ptr, int notation){
	static_cast<QTextStream*>(ptr)->setRealNumberNotation(static_cast<QTextStream::RealNumberNotation>(notation));
}

void QTextStream_SetRealNumberPrecision(void* ptr, int precision){
	static_cast<QTextStream*>(ptr)->setRealNumberPrecision(precision);
}

void QTextStream_SetStatus(void* ptr, int status){
	static_cast<QTextStream*>(ptr)->setStatus(static_cast<QTextStream::Status>(status));
}

void QTextStream_SetString(void* ptr, char* stri, int openMode){
	static_cast<QTextStream*>(ptr)->setString(new QString(stri), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QTextStream_SkipWhiteSpace(void* ptr){
	static_cast<QTextStream*>(ptr)->skipWhiteSpace();
}

int QTextStream_Status(void* ptr){
	return static_cast<QTextStream*>(ptr)->status();
}

char* QTextStream_String(void* ptr){
	return static_cast<QTextStream*>(ptr)->string()->toUtf8().data();
}

void QTextStream_DestroyQTextStream(void* ptr){
	static_cast<QTextStream*>(ptr)->~QTextStream();
}

char* QTextStream_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQTextStream*>(static_cast<QTextStream*>(ptr))) {
		return static_cast<MyQTextStream*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTextStream_BASE").toUtf8().data();
}

void QTextStream_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQTextStream*>(static_cast<QTextStream*>(ptr))) {
		static_cast<MyQTextStream*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQThread: public QThread {
public:
	MyQThread(QObject *parent) : QThread(parent) {};
	void Signal_Finished() { callbackQThreadFinished(this, this->objectName().toUtf8().data()); };
	void run() { callbackQThreadRun(this, this->objectName().toUtf8().data()); };
	void Signal_Started() { callbackQThreadStarted(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQThreadTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQThreadChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQThreadCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QThread_SetPriority(void* ptr, int priority){
	static_cast<QThread*>(ptr)->setPriority(static_cast<QThread::Priority>(priority));
}

void* QThread_NewQThread(void* parent){
	return new MyQThread(static_cast<QObject*>(parent));
}

void* QThread_QThread_CurrentThread(){
	return QThread::currentThread();
}

int QThread_Event(void* ptr, void* event){
	return static_cast<QThread*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QThread_EventDispatcher(void* ptr){
	return static_cast<QThread*>(ptr)->eventDispatcher();
}

void QThread_Exit(void* ptr, int returnCode){
	static_cast<QThread*>(ptr)->exit(returnCode);
}

void QThread_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

void QThread_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

int QThread_IsFinished(void* ptr){
	return static_cast<QThread*>(ptr)->isFinished();
}

int QThread_IsInterruptionRequested(void* ptr){
	return static_cast<QThread*>(ptr)->isInterruptionRequested();
}

int QThread_IsRunning(void* ptr){
	return static_cast<QThread*>(ptr)->isRunning();
}

int QThread_LoopLevel(void* ptr){
	return static_cast<QThread*>(ptr)->loopLevel();
}

int QThread_Priority(void* ptr){
	return static_cast<QThread*>(ptr)->priority();
}

void QThread_Quit(void* ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "quit");
}

void QThread_RequestInterruption(void* ptr){
	static_cast<QThread*>(ptr)->requestInterruption();
}

void QThread_Run(void* ptr){
	static_cast<MyQThread*>(ptr)->run();
}

void QThread_RunDefault(void* ptr){
	static_cast<QThread*>(ptr)->QThread::run();
}

void QThread_SetEventDispatcher(void* ptr, void* eventDispatcher){
	static_cast<QThread*>(ptr)->setEventDispatcher(static_cast<QAbstractEventDispatcher*>(eventDispatcher));
}

void QThread_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DestroyQThread(void* ptr){
	static_cast<QThread*>(ptr)->~QThread();
}

int QThread_QThread_IdealThreadCount(){
	return QThread::idealThreadCount();
}

void QThread_Start(void* ptr, int priority){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "start", Q_ARG(QThread::Priority, static_cast<QThread::Priority>(priority)));
}

void QThread_Terminate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "terminate");
}

void QThread_QThread_YieldCurrentThread(){
	QThread::yieldCurrentThread();
}

void QThread_TimerEvent(void* ptr, void* event){
	static_cast<MyQThread*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QThread_TimerEventDefault(void* ptr, void* event){
	static_cast<QThread*>(ptr)->QThread::timerEvent(static_cast<QTimerEvent*>(event));
}

void QThread_ChildEvent(void* ptr, void* event){
	static_cast<MyQThread*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QThread_ChildEventDefault(void* ptr, void* event){
	static_cast<QThread*>(ptr)->QThread::childEvent(static_cast<QChildEvent*>(event));
}

void QThread_CustomEvent(void* ptr, void* event){
	static_cast<MyQThread*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QThread_CustomEventDefault(void* ptr, void* event){
	static_cast<QThread*>(ptr)->QThread::customEvent(static_cast<QEvent*>(event));
}

int QThreadPool_ActiveThreadCount(void* ptr){
	return static_cast<QThreadPool*>(ptr)->activeThreadCount();
}

int QThreadPool_ExpiryTimeout(void* ptr){
	return static_cast<QThreadPool*>(ptr)->expiryTimeout();
}

int QThreadPool_MaxThreadCount(void* ptr){
	return static_cast<QThreadPool*>(ptr)->maxThreadCount();
}

void QThreadPool_SetExpiryTimeout(void* ptr, int expiryTimeout){
	static_cast<QThreadPool*>(ptr)->setExpiryTimeout(expiryTimeout);
}

void QThreadPool_SetMaxThreadCount(void* ptr, int maxThreadCount){
	static_cast<QThreadPool*>(ptr)->setMaxThreadCount(maxThreadCount);
}

void* QThreadPool_NewQThreadPool(void* parent){
	return new QThreadPool(static_cast<QObject*>(parent));
}

void QThreadPool_Cancel(void* ptr, void* runnable){
	static_cast<QThreadPool*>(ptr)->cancel(static_cast<QRunnable*>(runnable));
}

void QThreadPool_Clear(void* ptr){
	static_cast<QThreadPool*>(ptr)->clear();
}

void* QThreadPool_QThreadPool_GlobalInstance(){
	return QThreadPool::globalInstance();
}

void QThreadPool_ReleaseThread(void* ptr){
	static_cast<QThreadPool*>(ptr)->releaseThread();
}

void QThreadPool_ReserveThread(void* ptr){
	static_cast<QThreadPool*>(ptr)->reserveThread();
}

void QThreadPool_Start(void* ptr, void* runnable, int priority){
	static_cast<QThreadPool*>(ptr)->start(static_cast<QRunnable*>(runnable), priority);
}

int QThreadPool_TryStart(void* ptr, void* runnable){
	return static_cast<QThreadPool*>(ptr)->tryStart(static_cast<QRunnable*>(runnable));
}

int QThreadPool_WaitForDone(void* ptr, int msecs){
	return static_cast<QThreadPool*>(ptr)->waitForDone(msecs);
}

void QThreadPool_DestroyQThreadPool(void* ptr){
	static_cast<QThreadPool*>(ptr)->~QThreadPool();
}

void QThreadPool_TimerEvent(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QThreadPool_TimerEventDefault(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->QThreadPool::timerEvent(static_cast<QTimerEvent*>(event));
}

void QThreadPool_ChildEvent(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QThreadPool_ChildEventDefault(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->QThreadPool::childEvent(static_cast<QChildEvent*>(event));
}

void QThreadPool_CustomEvent(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QThreadPool_CustomEventDefault(void* ptr, void* event){
	static_cast<QThreadPool*>(ptr)->QThreadPool::customEvent(static_cast<QEvent*>(event));
}

void* QTime_NewQTime(){
	return new QTime();
}

void* QTime_NewQTime3(int h, int m, int s, int ms){
	return new QTime(h, m, s, ms);
}

int QTime_Elapsed(void* ptr){
	return static_cast<QTime*>(ptr)->elapsed();
}

int QTime_Hour(void* ptr){
	return static_cast<QTime*>(ptr)->hour();
}

int QTime_IsNull(void* ptr){
	return static_cast<QTime*>(ptr)->isNull();
}

int QTime_QTime_IsValid2(int h, int m, int s, int ms){
	return QTime::isValid(h, m, s, ms);
}

int QTime_IsValid(void* ptr){
	return static_cast<QTime*>(ptr)->isValid();
}

int QTime_Minute(void* ptr){
	return static_cast<QTime*>(ptr)->minute();
}

int QTime_Msec(void* ptr){
	return static_cast<QTime*>(ptr)->msec();
}

int QTime_MsecsSinceStartOfDay(void* ptr){
	return static_cast<QTime*>(ptr)->msecsSinceStartOfDay();
}

int QTime_MsecsTo(void* ptr, void* t){
	return static_cast<QTime*>(ptr)->msecsTo(*static_cast<QTime*>(t));
}

int QTime_Restart(void* ptr){
	return static_cast<QTime*>(ptr)->restart();
}

int QTime_Second(void* ptr){
	return static_cast<QTime*>(ptr)->second();
}

int QTime_SecsTo(void* ptr, void* t){
	return static_cast<QTime*>(ptr)->secsTo(*static_cast<QTime*>(t));
}

int QTime_SetHMS(void* ptr, int h, int m, int s, int ms){
	return static_cast<QTime*>(ptr)->setHMS(h, m, s, ms);
}

void QTime_Start(void* ptr){
	static_cast<QTime*>(ptr)->start();
}

char* QTime_ToString2(void* ptr, int format){
	return static_cast<QTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

char* QTime_ToString(void* ptr, char* format){
	return static_cast<QTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

class MyQTimeLine: public QTimeLine {
public:
	MyQTimeLine(int duration, QObject *parent) : QTimeLine(duration, parent) {};
	void Signal_Finished() { callbackQTimeLineFinished(this, this->objectName().toUtf8().data()); };
	void Signal_FrameChanged(int frame) { callbackQTimeLineFrameChanged(this, this->objectName().toUtf8().data(), frame); };
	void Signal_StateChanged(QTimeLine::State newState) { callbackQTimeLineStateChanged(this, this->objectName().toUtf8().data(), newState); };
	void timerEvent(QTimerEvent * event) { callbackQTimeLineTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_ValueChanged(qreal value) { callbackQTimeLineValueChanged(this, this->objectName().toUtf8().data(), static_cast<double>(value)); };
	void childEvent(QChildEvent * event) { callbackQTimeLineChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTimeLineCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QTimeLine_CurrentTime(void* ptr){
	return static_cast<QTimeLine*>(ptr)->currentTime();
}

int QTimeLine_CurveShape(void* ptr){
	return static_cast<QTimeLine*>(ptr)->curveShape();
}

int QTimeLine_Direction(void* ptr){
	return static_cast<QTimeLine*>(ptr)->direction();
}

int QTimeLine_Duration(void* ptr){
	return static_cast<QTimeLine*>(ptr)->duration();
}

void* QTimeLine_EasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QTimeLine*>(ptr)->easingCurve());
}

int QTimeLine_LoopCount(void* ptr){
	return static_cast<QTimeLine*>(ptr)->loopCount();
}

void QTimeLine_SetCurrentTime(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setCurrentTime", Q_ARG(int, msec));
}

void QTimeLine_SetCurveShape(void* ptr, int shape){
	static_cast<QTimeLine*>(ptr)->setCurveShape(static_cast<QTimeLine::CurveShape>(shape));
}

void QTimeLine_SetDirection(void* ptr, int direction){
	static_cast<QTimeLine*>(ptr)->setDirection(static_cast<QTimeLine::Direction>(direction));
}

void QTimeLine_SetDuration(void* ptr, int duration){
	static_cast<QTimeLine*>(ptr)->setDuration(duration);
}

void QTimeLine_SetEasingCurve(void* ptr, void* curve){
	static_cast<QTimeLine*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(curve));
}

void QTimeLine_SetLoopCount(void* ptr, int count){
	static_cast<QTimeLine*>(ptr)->setLoopCount(count);
}

void QTimeLine_SetUpdateInterval(void* ptr, int interval){
	static_cast<QTimeLine*>(ptr)->setUpdateInterval(interval);
}

int QTimeLine_UpdateInterval(void* ptr){
	return static_cast<QTimeLine*>(ptr)->updateInterval();
}

void* QTimeLine_NewQTimeLine(int duration, void* parent){
	return new MyQTimeLine(duration, static_cast<QObject*>(parent));
}

int QTimeLine_CurrentFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->currentFrame();
}

double QTimeLine_CurrentValue(void* ptr){
	return static_cast<double>(static_cast<QTimeLine*>(ptr)->currentValue());
}

int QTimeLine_EndFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->endFrame();
}

void QTimeLine_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_ConnectFrameChanged(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

void QTimeLine_DisconnectFrameChanged(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

int QTimeLine_FrameForTime(void* ptr, int msec){
	return static_cast<QTimeLine*>(ptr)->frameForTime(msec);
}

void QTimeLine_Resume(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "resume");
}

void QTimeLine_SetEndFrame(void* ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setEndFrame(frame);
}

void QTimeLine_SetFrameRange(void* ptr, int startFrame, int endFrame){
	static_cast<QTimeLine*>(ptr)->setFrameRange(startFrame, endFrame);
}

void QTimeLine_SetPaused(void* ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QTimeLine_SetStartFrame(void* ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setStartFrame(frame);
}

void QTimeLine_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "start");
}

int QTimeLine_StartFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->startFrame();
}

int QTimeLine_State(void* ptr){
	return static_cast<QTimeLine*>(ptr)->state();
}

void QTimeLine_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "stop");
}

void QTimeLine_TimerEvent(void* ptr, void* event){
	static_cast<MyQTimeLine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTimeLine_TimerEventDefault(void* ptr, void* event){
	static_cast<QTimeLine*>(ptr)->QTimeLine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTimeLine_ToggleDirection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "toggleDirection");
}

void QTimeLine_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::valueChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(qreal)>(&MyQTimeLine::Signal_ValueChanged));;
}

void QTimeLine_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::valueChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(qreal)>(&MyQTimeLine::Signal_ValueChanged));;
}

double QTimeLine_ValueForTime(void* ptr, int msec){
	return static_cast<double>(static_cast<QTimeLine*>(ptr)->valueForTime(msec));
}

void QTimeLine_DestroyQTimeLine(void* ptr){
	static_cast<QTimeLine*>(ptr)->~QTimeLine();
}

void QTimeLine_ChildEvent(void* ptr, void* event){
	static_cast<MyQTimeLine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTimeLine_ChildEventDefault(void* ptr, void* event){
	static_cast<QTimeLine*>(ptr)->QTimeLine::childEvent(static_cast<QChildEvent*>(event));
}

void QTimeLine_CustomEvent(void* ptr, void* event){
	static_cast<MyQTimeLine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTimeLine_CustomEventDefault(void* ptr, void* event){
	static_cast<QTimeLine*>(ptr)->QTimeLine::customEvent(static_cast<QEvent*>(event));
}

void* QTimeZone_NewQTimeZone(){
	return new QTimeZone();
}

void* QTimeZone_NewQTimeZone2(void* ianaId){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, char* name, char* abbreviation, int country, char* comment){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString(name), QString(abbreviation), static_cast<QLocale::Country>(country), QString(comment));
}

void* QTimeZone_NewQTimeZone5(void* other){
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

void* QTimeZone_NewQTimeZone3(int offsetSeconds){
	return new QTimeZone(offsetSeconds);
}

char* QTimeZone_Abbreviation(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->abbreviation(*static_cast<QDateTime*>(atDateTime)).toUtf8().data();
}

char* QTimeZone_Comment(void* ptr){
	return static_cast<QTimeZone*>(ptr)->comment().toUtf8().data();
}

int QTimeZone_Country(void* ptr){
	return static_cast<QTimeZone*>(ptr)->country();
}

int QTimeZone_DaylightTimeOffset(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->daylightTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

char* QTimeZone_DisplayName2(void* ptr, int timeType, int nameType, void* locale){
	return static_cast<QTimeZone*>(ptr)->displayName(static_cast<QTimeZone::TimeType>(timeType), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

char* QTimeZone_DisplayName(void* ptr, void* atDateTime, int nameType, void* locale){
	return static_cast<QTimeZone*>(ptr)->displayName(*static_cast<QDateTime*>(atDateTime), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

int QTimeZone_HasDaylightTime(void* ptr){
	return static_cast<QTimeZone*>(ptr)->hasDaylightTime();
}

int QTimeZone_HasTransitions(void* ptr){
	return static_cast<QTimeZone*>(ptr)->hasTransitions();
}

void* QTimeZone_QTimeZone_IanaIdToWindowsId(void* ianaId){
	return new QByteArray(QTimeZone::ianaIdToWindowsId(*static_cast<QByteArray*>(ianaId)));
}

void* QTimeZone_Id(void* ptr){
	return new QByteArray(static_cast<QTimeZone*>(ptr)->id());
}

int QTimeZone_IsDaylightTime(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->isDaylightTime(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_QTimeZone_IsTimeZoneIdAvailable(void* ianaId){
	return QTimeZone::isTimeZoneIdAvailable(*static_cast<QByteArray*>(ianaId));
}

int QTimeZone_IsValid(void* ptr){
	return static_cast<QTimeZone*>(ptr)->isValid();
}

int QTimeZone_OffsetFromUtc(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->offsetFromUtc(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_StandardTimeOffset(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->standardTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

void QTimeZone_Swap(void* ptr, void* other){
	static_cast<QTimeZone*>(ptr)->swap(*static_cast<QTimeZone*>(other));
}

void* QTimeZone_QTimeZone_SystemTimeZone(){
	return new QTimeZone(QTimeZone::systemTimeZone());
}

void* QTimeZone_QTimeZone_SystemTimeZoneId(){
	return new QByteArray(QTimeZone::systemTimeZoneId());
}

void* QTimeZone_QTimeZone_Utc(){
	return new QTimeZone(QTimeZone::utc());
}

void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId(void* windowsId){
	return new QByteArray(QTimeZone::windowsIdToDefaultIanaId(*static_cast<QByteArray*>(windowsId)));
}

void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId2(void* windowsId, int country){
	return new QByteArray(QTimeZone::windowsIdToDefaultIanaId(*static_cast<QByteArray*>(windowsId), static_cast<QLocale::Country>(country)));
}

void QTimeZone_DestroyQTimeZone(void* ptr){
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

class MyQTimer: public QTimer {
public:
	MyQTimer(QObject *parent) : QTimer(parent) {};
	void Signal_Timeout() { callbackQTimerTimeout(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * e) { callbackQTimerTimerEvent(this, this->objectName().toUtf8().data(), e); };
	void childEvent(QChildEvent * event) { callbackQTimerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTimerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QTimer_RemainingTime(void* ptr){
	return static_cast<QTimer*>(ptr)->remainingTime();
}

void QTimer_SetInterval(void* ptr, int msec){
	static_cast<QTimer*>(ptr)->setInterval(msec);
}

void* QTimer_NewQTimer(void* parent){
	return new MyQTimer(static_cast<QObject*>(parent));
}

int QTimer_Interval(void* ptr){
	return static_cast<QTimer*>(ptr)->interval();
}

int QTimer_IsActive(void* ptr){
	return static_cast<QTimer*>(ptr)->isActive();
}

int QTimer_IsSingleShot(void* ptr){
	return static_cast<QTimer*>(ptr)->isSingleShot();
}

void QTimer_SetSingleShot(void* ptr, int singleShot){
	static_cast<QTimer*>(ptr)->setSingleShot(singleShot != 0);
}

void QTimer_SetTimerType(void* ptr, int atype){
	static_cast<QTimer*>(ptr)->setTimerType(static_cast<Qt::TimerType>(atype));
}

void QTimer_QTimer_SingleShot2(int msec, int timerType, void* receiver, char* member){
	QTimer::singleShot(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_QTimer_SingleShot(int msec, void* receiver, char* member){
	QTimer::singleShot(msec, static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_Start2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start");
}

void QTimer_Start(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start", Q_ARG(int, msec));
}

void QTimer_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "stop");
}

void QTimer_ConnectTimeout(void* ptr){
	QObject::connect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

void QTimer_DisconnectTimeout(void* ptr){
	QObject::disconnect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

void QTimer_TimerEvent(void* ptr, void* e){
	static_cast<MyQTimer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(e));
}

void QTimer_TimerEventDefault(void* ptr, void* e){
	static_cast<QTimer*>(ptr)->QTimer::timerEvent(static_cast<QTimerEvent*>(e));
}

int QTimer_TimerId(void* ptr){
	return static_cast<QTimer*>(ptr)->timerId();
}

int QTimer_TimerType(void* ptr){
	return static_cast<QTimer*>(ptr)->timerType();
}

void QTimer_DestroyQTimer(void* ptr){
	static_cast<QTimer*>(ptr)->~QTimer();
}

void QTimer_ChildEvent(void* ptr, void* event){
	static_cast<MyQTimer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTimer_ChildEventDefault(void* ptr, void* event){
	static_cast<QTimer*>(ptr)->QTimer::childEvent(static_cast<QChildEvent*>(event));
}

void QTimer_CustomEvent(void* ptr, void* event){
	static_cast<MyQTimer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTimer_CustomEventDefault(void* ptr, void* event){
	static_cast<QTimer*>(ptr)->QTimer::customEvent(static_cast<QEvent*>(event));
}

void* QTimerEvent_NewQTimerEvent(int timerId){
	return new QTimerEvent(timerId);
}

int QTimerEvent_TimerId(void* ptr){
	return static_cast<QTimerEvent*>(ptr)->timerId();
}

class MyQTranslator: public QTranslator {
public:
	MyQTranslator(QObject *parent) : QTranslator(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQTranslatorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTranslatorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTranslatorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QTranslator_NewQTranslator(void* parent){
	return new MyQTranslator(static_cast<QObject*>(parent));
}

int QTranslator_IsEmpty(void* ptr){
	return static_cast<QTranslator*>(ptr)->isEmpty();
}

int QTranslator_Load2(void* ptr, void* locale, char* filename, char* prefix, char* directory, char* suffix){
	return static_cast<QTranslator*>(ptr)->load(*static_cast<QLocale*>(locale), QString(filename), QString(prefix), QString(directory), QString(suffix));
}

int QTranslator_Load(void* ptr, char* filename, char* directory, char* search_delimiters, char* suffix){
	return static_cast<QTranslator*>(ptr)->load(QString(filename), QString(directory), QString(search_delimiters), QString(suffix));
}

char* QTranslator_Translate(void* ptr, char* context, char* sourceText, char* disambiguation, int n){
	return static_cast<QTranslator*>(ptr)->translate(const_cast<const char*>(context), const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QTranslator_DestroyQTranslator(void* ptr){
	static_cast<QTranslator*>(ptr)->~QTranslator();
}

void QTranslator_TimerEvent(void* ptr, void* event){
	static_cast<MyQTranslator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTranslator_TimerEventDefault(void* ptr, void* event){
	static_cast<QTranslator*>(ptr)->QTranslator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTranslator_ChildEvent(void* ptr, void* event){
	static_cast<MyQTranslator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTranslator_ChildEventDefault(void* ptr, void* event){
	static_cast<QTranslator*>(ptr)->QTranslator::childEvent(static_cast<QChildEvent*>(event));
}

void QTranslator_CustomEvent(void* ptr, void* event){
	static_cast<MyQTranslator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTranslator_CustomEventDefault(void* ptr, void* event){
	static_cast<QTranslator*>(ptr)->QTranslator::customEvent(static_cast<QEvent*>(event));
}

void* QUrl_QUrl_FromEncoded(void* input, int parsingMode){
	return new QUrl(QUrl::fromEncoded(*static_cast<QByteArray*>(input), static_cast<QUrl::ParsingMode>(parsingMode)));
}

void* QUrl_NewQUrl(){
	return new QUrl();
}

void* QUrl_NewQUrl4(void* other){
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_NewQUrl3(char* url, int parsingMode){
	return new QUrl(QString(url), static_cast<QUrl::ParsingMode>(parsingMode));
}

void* QUrl_NewQUrl2(void* other){
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_Adjusted(void* ptr, int options){
	return new QUrl(static_cast<QUrl*>(ptr)->adjusted(static_cast<QUrl::UrlFormattingOption>(options)));
}

char* QUrl_Authority(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->authority(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

void QUrl_Clear(void* ptr){
	static_cast<QUrl*>(ptr)->clear();
}

char* QUrl_ErrorString(void* ptr){
	return static_cast<QUrl*>(ptr)->errorString().toUtf8().data();
}

char* QUrl_FileName(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->fileName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Fragment(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->fragment(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_QUrl_FromAce(void* domain){
	return QUrl::fromAce(*static_cast<QByteArray*>(domain)).toUtf8().data();
}

void* QUrl_QUrl_FromLocalFile(char* localFile){
	return new QUrl(QUrl::fromLocalFile(QString(localFile)));
}

char* QUrl_QUrl_FromPercentEncoding(void* input){
	return QUrl::fromPercentEncoding(*static_cast<QByteArray*>(input)).toUtf8().data();
}

void* QUrl_QUrl_FromUserInput(char* userInput){
	return new QUrl(QUrl::fromUserInput(QString(userInput)));
}

void* QUrl_QUrl_FromUserInput2(char* userInput, char* workingDirectory, int options){
	return new QUrl(QUrl::fromUserInput(QString(userInput), QString(workingDirectory), static_cast<QUrl::UserInputResolutionOption>(options)));
}

int QUrl_HasFragment(void* ptr){
	return static_cast<QUrl*>(ptr)->hasFragment();
}

int QUrl_HasQuery(void* ptr){
	return static_cast<QUrl*>(ptr)->hasQuery();
}

char* QUrl_Host(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->host(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_QUrl_IdnWhitelist(){
	return QUrl::idnWhitelist().join(",,,").toUtf8().data();
}

int QUrl_IsEmpty(void* ptr){
	return static_cast<QUrl*>(ptr)->isEmpty();
}

int QUrl_IsLocalFile(void* ptr){
	return static_cast<QUrl*>(ptr)->isLocalFile();
}

int QUrl_IsParentOf(void* ptr, void* childUrl){
	return static_cast<QUrl*>(ptr)->isParentOf(*static_cast<QUrl*>(childUrl));
}

int QUrl_IsRelative(void* ptr){
	return static_cast<QUrl*>(ptr)->isRelative();
}

int QUrl_IsValid(void* ptr){
	return static_cast<QUrl*>(ptr)->isValid();
}

int QUrl_Matches(void* ptr, void* url, int options){
	return static_cast<QUrl*>(ptr)->matches(*static_cast<QUrl*>(url), static_cast<QUrl::UrlFormattingOption>(options));
}

char* QUrl_Password(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->password(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Path(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->path(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

int QUrl_Port(void* ptr, int defaultPort){
	return static_cast<QUrl*>(ptr)->port(defaultPort);
}

char* QUrl_Query(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

void* QUrl_Resolved(void* ptr, void* relative){
	return new QUrl(static_cast<QUrl*>(ptr)->resolved(*static_cast<QUrl*>(relative)));
}

char* QUrl_Scheme(void* ptr){
	return static_cast<QUrl*>(ptr)->scheme().toUtf8().data();
}

void QUrl_SetAuthority(void* ptr, char* authority, int mode){
	static_cast<QUrl*>(ptr)->setAuthority(QString(authority), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetFragment(void* ptr, char* fragment, int mode){
	static_cast<QUrl*>(ptr)->setFragment(QString(fragment), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetHost(void* ptr, char* host, int mode){
	static_cast<QUrl*>(ptr)->setHost(QString(host), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_QUrl_SetIdnWhitelist(char* list){
	QUrl::setIdnWhitelist(QString(list).split(",,,", QString::SkipEmptyParts));
}

void QUrl_SetPassword(void* ptr, char* password, int mode){
	static_cast<QUrl*>(ptr)->setPassword(QString(password), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetPath(void* ptr, char* path, int mode){
	static_cast<QUrl*>(ptr)->setPath(QString(path), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetPort(void* ptr, int port){
	static_cast<QUrl*>(ptr)->setPort(port);
}

void QUrl_SetQuery(void* ptr, char* query, int mode){
	static_cast<QUrl*>(ptr)->setQuery(QString(query), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetQuery2(void* ptr, void* query){
	static_cast<QUrl*>(ptr)->setQuery(*static_cast<QUrlQuery*>(query));
}

void QUrl_SetScheme(void* ptr, char* scheme){
	static_cast<QUrl*>(ptr)->setScheme(QString(scheme));
}

void QUrl_SetUrl(void* ptr, char* url, int parsingMode){
	static_cast<QUrl*>(ptr)->setUrl(QString(url), static_cast<QUrl::ParsingMode>(parsingMode));
}

void QUrl_SetUserInfo(void* ptr, char* userInfo, int mode){
	static_cast<QUrl*>(ptr)->setUserInfo(QString(userInfo), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetUserName(void* ptr, char* userName, int mode){
	static_cast<QUrl*>(ptr)->setUserName(QString(userName), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_Swap(void* ptr, void* other){
	static_cast<QUrl*>(ptr)->swap(*static_cast<QUrl*>(other));
}

void* QUrl_QUrl_ToAce(char* domain){
	return new QByteArray(QUrl::toAce(QString(domain)));
}

char* QUrl_ToDisplayString(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->toDisplayString(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

void* QUrl_ToEncoded(void* ptr, int options){
	return new QByteArray(static_cast<QUrl*>(ptr)->toEncoded(static_cast<QUrl::UrlFormattingOption>(options)));
}

char* QUrl_ToLocalFile(void* ptr){
	return static_cast<QUrl*>(ptr)->toLocalFile().toUtf8().data();
}

void* QUrl_QUrl_ToPercentEncoding(char* input, void* exclude, void* include){
	return new QByteArray(QUrl::toPercentEncoding(QString(input), *static_cast<QByteArray*>(exclude), *static_cast<QByteArray*>(include)));
}

char* QUrl_ToString(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->toString(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

char* QUrl_TopLevelDomain(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->topLevelDomain(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Url(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->url(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

char* QUrl_UserInfo(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->userInfo(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_UserName(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->userName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

void QUrl_DestroyQUrl(void* ptr){
	static_cast<QUrl*>(ptr)->~QUrl();
}

void* QUrlQuery_NewQUrlQuery(){
	return new QUrlQuery();
}

void* QUrlQuery_NewQUrlQuery3(char* queryString){
	return new QUrlQuery(QString(queryString));
}

void* QUrlQuery_NewQUrlQuery2(void* url){
	return new QUrlQuery(*static_cast<QUrl*>(url));
}

void* QUrlQuery_NewQUrlQuery4(void* other){
	return new QUrlQuery(*static_cast<QUrlQuery*>(other));
}

void QUrlQuery_AddQueryItem(void* ptr, char* key, char* value){
	static_cast<QUrlQuery*>(ptr)->addQueryItem(QString(key), QString(value));
}

char* QUrlQuery_AllQueryItemValues(void* ptr, char* key, int encoding){
	return static_cast<QUrlQuery*>(ptr)->allQueryItemValues(QString(key), static_cast<QUrl::ComponentFormattingOption>(encoding)).join(",,,").toUtf8().data();
}

void QUrlQuery_Clear(void* ptr){
	static_cast<QUrlQuery*>(ptr)->clear();
}

int QUrlQuery_IsEmpty(void* ptr){
	return static_cast<QUrlQuery*>(ptr)->isEmpty();
}

char* QUrlQuery_Query(void* ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_RemoveAllQueryItems(void* ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeAllQueryItems(QString(key));
}

void QUrlQuery_RemoveQueryItem(void* ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeQueryItem(QString(key));
}

void QUrlQuery_SetQuery(void* ptr, char* queryString){
	static_cast<QUrlQuery*>(ptr)->setQuery(QString(queryString));
}

void QUrlQuery_SetQueryDelimiters(void* ptr, void* valueDelimiter, void* pairDelimiter){
	static_cast<QUrlQuery*>(ptr)->setQueryDelimiters(*static_cast<QChar*>(valueDelimiter), *static_cast<QChar*>(pairDelimiter));
}

void QUrlQuery_Swap(void* ptr, void* other){
	static_cast<QUrlQuery*>(ptr)->swap(*static_cast<QUrlQuery*>(other));
}

char* QUrlQuery_ToString(void* ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->toString(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_DestroyQUrlQuery(void* ptr){
	static_cast<QUrlQuery*>(ptr)->~QUrlQuery();
}

int QUuid_Variant(void* ptr){
	return static_cast<QUuid*>(ptr)->variant();
}

int QUuid_Version(void* ptr){
	return static_cast<QUuid*>(ptr)->version();
}

void* QUuid_NewQUuid(){
	return new QUuid();
}

void* QUuid_NewQUuid5(void* text){
	return new QUuid(*static_cast<QByteArray*>(text));
}

void* QUuid_NewQUuid3(char* text){
	return new QUuid(QString(text));
}

int QUuid_IsNull(void* ptr){
	return static_cast<QUuid*>(ptr)->isNull();
}

void* QUuid_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QUuid*>(ptr)->toByteArray());
}

void* QUuid_ToRfc4122(void* ptr){
	return new QByteArray(static_cast<QUuid*>(ptr)->toRfc4122());
}

char* QUuid_ToString(void* ptr){
	return static_cast<QUuid*>(ptr)->toString().toUtf8().data();
}

void* QVariant_NewQVariant20(void* c){
	return new QVariant(*static_cast<QChar*>(c));
}

void* QVariant_NewQVariant18(void* val){
	return new QVariant(*static_cast<QLatin1String*>(val));
}

void* QVariant_NewQVariant11(int val){
	return new QVariant(val != 0);
}

void* QVariant_NewQVariant16(void* val){
	return new QVariant(*static_cast<QBitArray*>(val));
}

void* QVariant_NewQVariant15(void* val){
	return new QVariant(*static_cast<QByteArray*>(val));
}

void* QVariant_NewQVariant21(void* val){
	return new QVariant(*static_cast<QDate*>(val));
}

void* QVariant_NewQVariant23(void* val){
	return new QVariant(*static_cast<QDateTime*>(val));
}

void* QVariant_NewQVariant39(void* val){
	return new QVariant(*static_cast<QEasingCurve*>(val));
}

void* QVariant_NewQVariant45(void* val){
	return new QVariant(*static_cast<QJsonArray*>(val));
}

void* QVariant_NewQVariant46(void* val){
	return new QVariant(*static_cast<QJsonDocument*>(val));
}

void* QVariant_NewQVariant44(void* val){
	return new QVariant(*static_cast<QJsonObject*>(val));
}

void* QVariant_NewQVariant43(void* val){
	return new QVariant(*static_cast<QJsonValue*>(val));
}

void* QVariant_NewQVariant31(void* val){
	return new QVariant(*static_cast<QLine*>(val));
}

void* QVariant_NewQVariant32(void* val){
	return new QVariant(*static_cast<QLineF*>(val));
}

void* QVariant_NewQVariant35(void* l){
	return new QVariant(*static_cast<QLocale*>(l));
}

void* QVariant_NewQVariant41(void* val){
	return new QVariant(*static_cast<QModelIndex*>(val));
}

void* QVariant_NewQVariant42(void* val){
	return new QVariant(*static_cast<QPersistentModelIndex*>(val));
}

void* QVariant_NewQVariant29(void* val){
	return new QVariant(*static_cast<QPoint*>(val));
}

void* QVariant_NewQVariant30(void* val){
	return new QVariant(*static_cast<QPointF*>(val));
}

void* QVariant_NewQVariant33(void* val){
	return new QVariant(*static_cast<QRect*>(val));
}

void* QVariant_NewQVariant34(void* val){
	return new QVariant(*static_cast<QRectF*>(val));
}

void* QVariant_NewQVariant36(void* regExp){
	return new QVariant(*static_cast<QRegExp*>(regExp));
}

void* QVariant_NewQVariant37(void* re){
	return new QVariant(*static_cast<QRegularExpression*>(re));
}

void* QVariant_NewQVariant27(void* val){
	return new QVariant(*static_cast<QSize*>(val));
}

void* QVariant_NewQVariant28(void* val){
	return new QVariant(*static_cast<QSizeF*>(val));
}

void* QVariant_NewQVariant17(char* val){
	return new QVariant(QString(val));
}

void* QVariant_NewQVariant19(char* val){
	return new QVariant(QString(val).split(",,,", QString::SkipEmptyParts));
}

void* QVariant_NewQVariant22(void* val){
	return new QVariant(*static_cast<QTime*>(val));
}

void* QVariant_NewQVariant38(void* val){
	return new QVariant(*static_cast<QUrl*>(val));
}

void* QVariant_NewQVariant40(void* val){
	return new QVariant(*static_cast<QUuid*>(val));
}

void* QVariant_NewQVariant5(void* p){
	return new QVariant(*static_cast<QVariant*>(p));
}

void* QVariant_NewQVariant14(char* val){
	return new QVariant(const_cast<const char*>(val));
}

void* QVariant_NewQVariant7(int val){
	return new QVariant(val);
}

void* QVariant_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QVariant*>(ptr)->toByteArray());
}

void* QVariant_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QVariant*>(ptr)->toDateTime());
}

void* QVariant_ToEasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QVariant*>(ptr)->toEasingCurve());
}

void* QVariant_ToPoint(void* ptr){
	return new QPoint(static_cast<QPoint>(static_cast<QVariant*>(ptr)->toPoint()).x(), static_cast<QPoint>(static_cast<QVariant*>(ptr)->toPoint()).y());
}

void* QVariant_ToRect(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QVariant*>(ptr)->toRect()).x(), static_cast<QRect>(static_cast<QVariant*>(ptr)->toRect()).y(), static_cast<QRect>(static_cast<QVariant*>(ptr)->toRect()).width(), static_cast<QRect>(static_cast<QVariant*>(ptr)->toRect()).height());
}

void* QVariant_ToRegExp(void* ptr){
	return new QRegExp(static_cast<QVariant*>(ptr)->toRegExp());
}

void* QVariant_ToRegularExpression(void* ptr){
	return new QRegularExpression(static_cast<QVariant*>(ptr)->toRegularExpression());
}

void* QVariant_ToSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QVariant*>(ptr)->toSize()).width(), static_cast<QSize>(static_cast<QVariant*>(ptr)->toSize()).height());
}

char* QVariant_ToStringList(void* ptr){
	return static_cast<QVariant*>(ptr)->toStringList().join(",,,").toUtf8().data();
}

void* QVariant_ToUrl(void* ptr){
	return new QUrl(static_cast<QVariant*>(ptr)->toUrl());
}

void QVariant_DestroyQVariant(void* ptr){
	static_cast<QVariant*>(ptr)->~QVariant();
}

void* QVariant_NewQVariant(){
	return new QVariant();
}

void* QVariant_NewQVariant6(void* s){
	return new QVariant(*static_cast<QDataStream*>(s));
}

void* QVariant_NewQVariant47(void* other){
	return new QVariant(*static_cast<QVariant*>(other));
}

void QVariant_Clear(void* ptr){
	static_cast<QVariant*>(ptr)->clear();
}

int QVariant_Convert(void* ptr, int targetTypeId){
	return static_cast<QVariant*>(ptr)->convert(targetTypeId);
}

int QVariant_IsNull(void* ptr){
	return static_cast<QVariant*>(ptr)->isNull();
}

int QVariant_IsValid(void* ptr){
	return static_cast<QVariant*>(ptr)->isValid();
}

void QVariant_Swap(void* ptr, void* other){
	static_cast<QVariant*>(ptr)->swap(*static_cast<QVariant*>(other));
}

int QVariant_ToBool(void* ptr){
	return static_cast<QVariant*>(ptr)->toBool();
}

int QVariant_ToInt(void* ptr, int ok){
	return static_cast<QVariant*>(ptr)->toInt(NULL);
}

void* QVariant_ToJsonArray(void* ptr){
	return new QJsonArray(static_cast<QVariant*>(ptr)->toJsonArray());
}

void* QVariant_ToJsonDocument(void* ptr){
	return new QJsonDocument(static_cast<QVariant*>(ptr)->toJsonDocument());
}

void* QVariant_ToJsonObject(void* ptr){
	return new QJsonObject(static_cast<QVariant*>(ptr)->toJsonObject());
}

void* QVariant_ToModelIndex(void* ptr){
	return new QModelIndex(static_cast<QVariant*>(ptr)->toModelIndex());
}

double QVariant_ToReal(void* ptr, int ok){
	return static_cast<double>(static_cast<QVariant*>(ptr)->toReal(NULL));
}

char* QVariant_ToString(void* ptr){
	return static_cast<QVariant*>(ptr)->toString().toUtf8().data();
}

int QVariant_UserType(void* ptr){
	return static_cast<QVariant*>(ptr)->userType();
}

class MyQVariantAnimation: public QVariantAnimation {
public:
	MyQVariantAnimation(QObject *parent) : QVariantAnimation(parent) {};
	void updateCurrentTime(int v) { callbackQVariantAnimationUpdateCurrentTime(this, this->objectName().toUtf8().data(), v); };
	void updateCurrentValue(const QVariant & value) { callbackQVariantAnimationUpdateCurrentValue(this, this->objectName().toUtf8().data(), new QVariant(value)); };
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) { callbackQVariantAnimationUpdateState(this, this->objectName().toUtf8().data(), newState, oldState); };
	void Signal_ValueChanged(const QVariant & value) { callbackQVariantAnimationValueChanged(this, this->objectName().toUtf8().data(), new QVariant(value)); };
	void updateDirection(QVariantAnimation::Direction direction) { callbackQVariantAnimationUpdateDirection(this, this->objectName().toUtf8().data(), direction); };
	void timerEvent(QTimerEvent * event) { callbackQVariantAnimationTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQVariantAnimationChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQVariantAnimationCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QVariantAnimation_CurrentValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->currentValue());
}

int QVariantAnimation_Duration(void* ptr){
	return static_cast<QVariantAnimation*>(ptr)->duration();
}

void* QVariantAnimation_EasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QVariantAnimation*>(ptr)->easingCurve());
}

void* QVariantAnimation_EndValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->endValue());
}

void QVariantAnimation_SetDuration(void* ptr, int msecs){
	static_cast<QVariantAnimation*>(ptr)->setDuration(msecs);
}

void QVariantAnimation_SetEasingCurve(void* ptr, void* easing){
	static_cast<QVariantAnimation*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(easing));
}

void QVariantAnimation_SetEndValue(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->setEndValue(*static_cast<QVariant*>(value));
}

void QVariantAnimation_SetStartValue(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->setStartValue(*static_cast<QVariant*>(value));
}

void* QVariantAnimation_StartValue(void* ptr){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->startValue());
}

void* QVariantAnimation_NewQVariantAnimation(void* parent){
	return new MyQVariantAnimation(static_cast<QObject*>(parent));
}

void* QVariantAnimation_KeyValueAt(void* ptr, double step){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->keyValueAt(static_cast<double>(step)));
}

int QVariantAnimation_Event(void* ptr, void* event){
	return static_cast<QVariantAnimation*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QVariantAnimation_Interpolated(void* ptr, void* from, void* to, double progress){
	return new QVariant(static_cast<QVariantAnimation*>(ptr)->interpolated(*static_cast<QVariant*>(from), *static_cast<QVariant*>(to), static_cast<double>(progress)));
}

void QVariantAnimation_SetKeyValueAt(void* ptr, double step, void* value){
	static_cast<QVariantAnimation*>(ptr)->setKeyValueAt(static_cast<double>(step), *static_cast<QVariant*>(value));
}

void QVariantAnimation_UpdateCurrentTime(void* ptr, int v){
	static_cast<MyQVariantAnimation*>(ptr)->updateCurrentTime(v);
}

void QVariantAnimation_UpdateCurrentTimeDefault(void* ptr, int v){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::updateCurrentTime(v);
}

void QVariantAnimation_UpdateCurrentValue(void* ptr, void* value){
	static_cast<MyQVariantAnimation*>(ptr)->updateCurrentValue(*static_cast<QVariant*>(value));
}

void QVariantAnimation_UpdateCurrentValueDefault(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::updateCurrentValue(*static_cast<QVariant*>(value));
}

void QVariantAnimation_UpdateState(void* ptr, int newState, int oldState){
	static_cast<MyQVariantAnimation*>(ptr)->updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QVariantAnimation_UpdateStateDefault(void* ptr, int newState, int oldState){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::updateState(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}

void QVariantAnimation_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QVariantAnimation*>(ptr), static_cast<void (QVariantAnimation::*)(const QVariant &)>(&QVariantAnimation::valueChanged), static_cast<MyQVariantAnimation*>(ptr), static_cast<void (MyQVariantAnimation::*)(const QVariant &)>(&MyQVariantAnimation::Signal_ValueChanged));;
}

void QVariantAnimation_ValueChanged(void* ptr, void* value){
	static_cast<QVariantAnimation*>(ptr)->valueChanged(*static_cast<QVariant*>(value));
}

void QVariantAnimation_DestroyQVariantAnimation(void* ptr){
	static_cast<QVariantAnimation*>(ptr)->~QVariantAnimation();
}

void QVariantAnimation_UpdateDirection(void* ptr, int direction){
	static_cast<MyQVariantAnimation*>(ptr)->updateDirection(static_cast<QVariantAnimation::Direction>(direction));
}

void QVariantAnimation_UpdateDirectionDefault(void* ptr, int direction){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::updateDirection(static_cast<QVariantAnimation::Direction>(direction));
}

void QVariantAnimation_TimerEvent(void* ptr, void* event){
	static_cast<MyQVariantAnimation*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVariantAnimation_TimerEventDefault(void* ptr, void* event){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVariantAnimation_ChildEvent(void* ptr, void* event){
	static_cast<MyQVariantAnimation*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVariantAnimation_ChildEventDefault(void* ptr, void* event){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::childEvent(static_cast<QChildEvent*>(event));
}

void QVariantAnimation_CustomEvent(void* ptr, void* event){
	static_cast<MyQVariantAnimation*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVariantAnimation_CustomEventDefault(void* ptr, void* event){
	static_cast<QVariantAnimation*>(ptr)->QVariantAnimation::customEvent(static_cast<QEvent*>(event));
}

void* QWaitCondition_NewQWaitCondition(){
	return new QWaitCondition();
}

void QWaitCondition_WakeAll(void* ptr){
	static_cast<QWaitCondition*>(ptr)->wakeAll();
}

void QWaitCondition_WakeOne(void* ptr){
	static_cast<QWaitCondition*>(ptr)->wakeOne();
}

void QWaitCondition_DestroyQWaitCondition(void* ptr){
	static_cast<QWaitCondition*>(ptr)->~QWaitCondition();
}

void* QWriteLocker_NewQWriteLocker(void* lock){
	return new QWriteLocker(static_cast<QReadWriteLock*>(lock));
}

void* QWriteLocker_ReadWriteLock(void* ptr){
	return static_cast<QWriteLocker*>(ptr)->readWriteLock();
}

void QWriteLocker_Relock(void* ptr){
	static_cast<QWriteLocker*>(ptr)->relock();
}

void QWriteLocker_Unlock(void* ptr){
	static_cast<QWriteLocker*>(ptr)->unlock();
}

void QWriteLocker_DestroyQWriteLocker(void* ptr){
	static_cast<QWriteLocker*>(ptr)->~QWriteLocker();
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute(){
	return new QXmlStreamAttribute();
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute3(char* namespaceUri, char* name, char* value){
	return new QXmlStreamAttribute(QString(namespaceUri), QString(name), QString(value));
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute2(char* qualifiedName, char* value){
	return new QXmlStreamAttribute(QString(qualifiedName), QString(value));
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute4(void* other){
	return new QXmlStreamAttribute(*static_cast<QXmlStreamAttribute*>(other));
}

int QXmlStreamAttribute_IsDefault(void* ptr){
	return static_cast<QXmlStreamAttribute*>(ptr)->isDefault();
}

void* QXmlStreamAttribute_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->name());
}

void* QXmlStreamAttribute_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->namespaceUri());
}

void* QXmlStreamAttribute_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->prefix());
}

void* QXmlStreamAttribute_QualifiedName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->qualifiedName());
}

void* QXmlStreamAttribute_Value(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->value());
}

void QXmlStreamAttribute_DestroyQXmlStreamAttribute(void* ptr){
	static_cast<QXmlStreamAttribute*>(ptr)->~QXmlStreamAttribute();
}

void* QXmlStreamAttributes_NewQXmlStreamAttributes(){
	return new QXmlStreamAttributes();
}

void QXmlStreamAttributes_Append(void* ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamAttributes_Append2(void* ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(qualifiedName), QString(value));
}

int QXmlStreamAttributes_HasAttribute2(void* ptr, void* qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(*static_cast<QLatin1String*>(qualifiedName));
}

int QXmlStreamAttributes_HasAttribute3(void* ptr, char* namespaceUri, char* name){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(namespaceUri), QString(name));
}

int QXmlStreamAttributes_HasAttribute(void* ptr, char* qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(qualifiedName));
}

void* QXmlStreamAttributes_Value3(void* ptr, void* namespaceUri, void* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(*static_cast<QLatin1String*>(namespaceUri), *static_cast<QLatin1String*>(name)));
}

void* QXmlStreamAttributes_Value5(void* ptr, void* qualifiedName){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qualifiedName)));
}

void* QXmlStreamAttributes_Value2(void* ptr, char* namespaceUri, void* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(namespaceUri), *static_cast<QLatin1String*>(name)));
}

void* QXmlStreamAttributes_Value(void* ptr, char* namespaceUri, char* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(namespaceUri), QString(name)));
}

void* QXmlStreamAttributes_Value4(void* ptr, char* qualifiedName){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(qualifiedName)));
}

void* QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration(){
	return new QXmlStreamEntityDeclaration();
}

void* QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration2(void* other){
	return new QXmlStreamEntityDeclaration(*static_cast<QXmlStreamEntityDeclaration*>(other));
}

void* QXmlStreamEntityDeclaration_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->name());
}

void* QXmlStreamEntityDeclaration_NotationName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->notationName());
}

void* QXmlStreamEntityDeclaration_PublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->publicId());
}

void* QXmlStreamEntityDeclaration_SystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->systemId());
}

void* QXmlStreamEntityDeclaration_Value(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->value());
}

void QXmlStreamEntityDeclaration_DestroyQXmlStreamEntityDeclaration(void* ptr){
	static_cast<QXmlStreamEntityDeclaration*>(ptr)->~QXmlStreamEntityDeclaration();
}

class MyQXmlStreamEntityResolver: public QXmlStreamEntityResolver {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

char* QXmlStreamEntityResolver_ResolveUndeclaredEntity(void* ptr, char* name){
	return static_cast<QXmlStreamEntityResolver*>(ptr)->resolveUndeclaredEntity(QString(name)).toUtf8().data();
}

void QXmlStreamEntityResolver_DestroyQXmlStreamEntityResolver(void* ptr){
	static_cast<QXmlStreamEntityResolver*>(ptr)->~QXmlStreamEntityResolver();
}

char* QXmlStreamEntityResolver_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlStreamEntityResolver*>(static_cast<QXmlStreamEntityResolver*>(ptr))) {
		return static_cast<MyQXmlStreamEntityResolver*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlStreamEntityResolver_BASE").toUtf8().data();
}

void QXmlStreamEntityResolver_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlStreamEntityResolver*>(static_cast<QXmlStreamEntityResolver*>(ptr))) {
		static_cast<MyQXmlStreamEntityResolver*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration(){
	return new QXmlStreamNamespaceDeclaration();
}

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration3(char* prefix, char* namespaceUri){
	return new QXmlStreamNamespaceDeclaration(QString(prefix), QString(namespaceUri));
}

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration2(void* other){
	return new QXmlStreamNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(other));
}

void* QXmlStreamNamespaceDeclaration_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->namespaceUri());
}

void* QXmlStreamNamespaceDeclaration_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->prefix());
}

void QXmlStreamNamespaceDeclaration_DestroyQXmlStreamNamespaceDeclaration(void* ptr){
	static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->~QXmlStreamNamespaceDeclaration();
}

void* QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration(){
	return new QXmlStreamNotationDeclaration();
}

void* QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration2(void* other){
	return new QXmlStreamNotationDeclaration(*static_cast<QXmlStreamNotationDeclaration*>(other));
}

void* QXmlStreamNotationDeclaration_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->name());
}

void* QXmlStreamNotationDeclaration_PublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->publicId());
}

void* QXmlStreamNotationDeclaration_SystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->systemId());
}

void QXmlStreamNotationDeclaration_DestroyQXmlStreamNotationDeclaration(void* ptr){
	static_cast<QXmlStreamNotationDeclaration*>(ptr)->~QXmlStreamNotationDeclaration();
}

int QXmlStreamReader_NamespaceProcessing(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->namespaceProcessing();
}

void QXmlStreamReader_SetNamespaceProcessing(void* ptr, int v){
	static_cast<QXmlStreamReader*>(ptr)->setNamespaceProcessing(v != 0);
}

void* QXmlStreamReader_NewQXmlStreamReader(){
	return new QXmlStreamReader();
}

void* QXmlStreamReader_NewQXmlStreamReader2(void* device){
	return new QXmlStreamReader(static_cast<QIODevice*>(device));
}

void* QXmlStreamReader_NewQXmlStreamReader3(void* data){
	return new QXmlStreamReader(*static_cast<QByteArray*>(data));
}

void* QXmlStreamReader_NewQXmlStreamReader4(char* data){
	return new QXmlStreamReader(QString(data));
}

void* QXmlStreamReader_NewQXmlStreamReader5(char* data){
	return new QXmlStreamReader(const_cast<const char*>(data));
}

void QXmlStreamReader_AddData(void* ptr, void* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QXmlStreamReader_AddData2(void* ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(QString(data));
}

void QXmlStreamReader_AddData3(void* ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(const_cast<const char*>(data));
}

void QXmlStreamReader_AddExtraNamespaceDeclaration(void* ptr, void* extraNamespaceDeclaration){
	static_cast<QXmlStreamReader*>(ptr)->addExtraNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(extraNamespaceDeclaration));
}

int QXmlStreamReader_AtEnd(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->atEnd();
}

long long QXmlStreamReader_CharacterOffset(void* ptr){
	return static_cast<long long>(static_cast<QXmlStreamReader*>(ptr)->characterOffset());
}

void QXmlStreamReader_Clear(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->clear();
}

long long QXmlStreamReader_ColumnNumber(void* ptr){
	return static_cast<long long>(static_cast<QXmlStreamReader*>(ptr)->columnNumber());
}

void* QXmlStreamReader_Device(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->device();
}

void* QXmlStreamReader_DocumentEncoding(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->documentEncoding());
}

void* QXmlStreamReader_DocumentVersion(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->documentVersion());
}

void* QXmlStreamReader_DtdName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdName());
}

void* QXmlStreamReader_DtdPublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdPublicId());
}

void* QXmlStreamReader_DtdSystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdSystemId());
}

void* QXmlStreamReader_EntityResolver(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->entityResolver();
}

int QXmlStreamReader_Error(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->error();
}

char* QXmlStreamReader_ErrorString(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->errorString().toUtf8().data();
}

int QXmlStreamReader_HasError(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->hasError();
}

int QXmlStreamReader_IsCDATA(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCDATA();
}

int QXmlStreamReader_IsCharacters(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCharacters();
}

int QXmlStreamReader_IsComment(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isComment();
}

int QXmlStreamReader_IsDTD(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isDTD();
}

int QXmlStreamReader_IsEndDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndDocument();
}

int QXmlStreamReader_IsEndElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndElement();
}

int QXmlStreamReader_IsEntityReference(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEntityReference();
}

int QXmlStreamReader_IsProcessingInstruction(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isProcessingInstruction();
}

int QXmlStreamReader_IsStandaloneDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStandaloneDocument();
}

int QXmlStreamReader_IsStartDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartDocument();
}

int QXmlStreamReader_IsStartElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartElement();
}

int QXmlStreamReader_IsWhitespace(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isWhitespace();
}

long long QXmlStreamReader_LineNumber(void* ptr){
	return static_cast<long long>(static_cast<QXmlStreamReader*>(ptr)->lineNumber());
}

void* QXmlStreamReader_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->name());
}

void* QXmlStreamReader_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->namespaceUri());
}

void* QXmlStreamReader_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->prefix());
}

void* QXmlStreamReader_ProcessingInstructionData(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->processingInstructionData());
}

void* QXmlStreamReader_ProcessingInstructionTarget(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->processingInstructionTarget());
}

void* QXmlStreamReader_QualifiedName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->qualifiedName());
}

void QXmlStreamReader_RaiseError(void* ptr, char* message){
	static_cast<QXmlStreamReader*>(ptr)->raiseError(QString(message));
}

char* QXmlStreamReader_ReadElementText(void* ptr, int behaviour){
	return static_cast<QXmlStreamReader*>(ptr)->readElementText(static_cast<QXmlStreamReader::ReadElementTextBehaviour>(behaviour)).toUtf8().data();
}

int QXmlStreamReader_ReadNext(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNext();
}

int QXmlStreamReader_ReadNextStartElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNextStartElement();
}

void QXmlStreamReader_SetDevice(void* ptr, void* device){
	static_cast<QXmlStreamReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamReader_SetEntityResolver(void* ptr, void* resolver){
	static_cast<QXmlStreamReader*>(ptr)->setEntityResolver(static_cast<QXmlStreamEntityResolver*>(resolver));
}

void QXmlStreamReader_SkipCurrentElement(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->skipCurrentElement();
}

void* QXmlStreamReader_Text(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->text());
}

char* QXmlStreamReader_TokenString(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenString().toUtf8().data();
}

int QXmlStreamReader_TokenType(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenType();
}

void QXmlStreamReader_DestroyQXmlStreamReader(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->~QXmlStreamReader();
}

int QXmlStreamWriter_AutoFormattingIndent(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormattingIndent();
}

void QXmlStreamWriter_SetAutoFormattingIndent(void* ptr, int spacesOrTabs){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormattingIndent(spacesOrTabs);
}

int QXmlStreamWriter_AutoFormatting(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormatting();
}

void* QXmlStreamWriter_Codec(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->codec();
}

void* QXmlStreamWriter_Device(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->device();
}

int QXmlStreamWriter_HasError(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->hasError();
}

void QXmlStreamWriter_SetAutoFormatting(void* ptr, int enable){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormatting(enable != 0);
}

void QXmlStreamWriter_SetCodec(void* ptr, void* codec){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QXmlStreamWriter_SetCodec2(void* ptr, char* codecName){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QXmlStreamWriter_SetDevice(void* ptr, void* device){
	static_cast<QXmlStreamWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamWriter_WriteAttribute(void* ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamWriter_WriteAttribute2(void* ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(qualifiedName), QString(value));
}

void QXmlStreamWriter_WriteAttribute3(void* ptr, void* attribute){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(*static_cast<QXmlStreamAttribute*>(attribute));
}

void QXmlStreamWriter_WriteAttributes(void* ptr, void* attributes){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttributes(*static_cast<QXmlStreamAttributes*>(attributes));
}

void QXmlStreamWriter_WriteCDATA(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCDATA(QString(text));
}

void QXmlStreamWriter_WriteCharacters(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCharacters(QString(text));
}

void QXmlStreamWriter_WriteComment(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeComment(QString(text));
}

void QXmlStreamWriter_WriteCurrentToken(void* ptr, void* reader){
	static_cast<QXmlStreamWriter*>(ptr)->writeCurrentToken(*static_cast<QXmlStreamReader*>(reader));
}

void QXmlStreamWriter_WriteDTD(void* ptr, char* dtd){
	static_cast<QXmlStreamWriter*>(ptr)->writeDTD(QString(dtd));
}

void QXmlStreamWriter_WriteDefaultNamespace(void* ptr, char* namespaceUri){
	static_cast<QXmlStreamWriter*>(ptr)->writeDefaultNamespace(QString(namespaceUri));
}

void QXmlStreamWriter_WriteEmptyElement(void* ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteEmptyElement2(void* ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteEndDocument(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndDocument();
}

void QXmlStreamWriter_WriteEndElement(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndElement();
}

void QXmlStreamWriter_WriteEntityReference(void* ptr, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEntityReference(QString(name));
}

void QXmlStreamWriter_WriteNamespace(void* ptr, char* namespaceUri, char* prefix){
	static_cast<QXmlStreamWriter*>(ptr)->writeNamespace(QString(namespaceUri), QString(prefix));
}

void QXmlStreamWriter_WriteProcessingInstruction(void* ptr, char* target, char* data){
	static_cast<QXmlStreamWriter*>(ptr)->writeProcessingInstruction(QString(target), QString(data));
}

void QXmlStreamWriter_WriteStartDocument3(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument();
}

void QXmlStreamWriter_WriteStartDocument(void* ptr, char* version){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version));
}

void QXmlStreamWriter_WriteStartDocument2(void* ptr, char* version, int standalone){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version), standalone != 0);
}

void QXmlStreamWriter_WriteStartElement(void* ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteStartElement2(void* ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteTextElement(void* ptr, char* namespaceUri, char* name, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(namespaceUri), QString(name), QString(text));
}

void QXmlStreamWriter_WriteTextElement2(void* ptr, char* qualifiedName, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(qualifiedName), QString(text));
}

void QXmlStreamWriter_DestroyQXmlStreamWriter(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->~QXmlStreamWriter();
}

int Qt_LocaleDate_Type(){
	return Qt::LocaleDate;
}

int Qt_SystemLocaleShortDate_Type(){
	return Qt::SystemLocaleShortDate;
}

int Qt_SystemLocaleLongDate_Type(){
	return Qt::SystemLocaleLongDate;
}

int Qt_DefaultLocaleShortDate_Type(){
	return Qt::DefaultLocaleShortDate;
}

int Qt_DefaultLocaleLongDate_Type(){
	return Qt::DefaultLocaleLongDate;
}

int Qt_RFC2822Date_Type(){
	return Qt::RFC2822Date;
}

int Qt_LastGestureType_Type(){
	return Qt::LastGestureType;
}

