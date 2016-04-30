#define protected public
#define private public

#include "designer.h"
#include "_cgo_export.h"

#include <QAbstractExtensionFactory>
#include <QAbstractExtensionManager>
#include <QAbstractFormBuilder>
#include <QAction>
#include <QActionEvent>
#include <QActionGroup>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDesignerActionEditorInterface>
#include <QDesignerContainerExtension>
#include <QDesignerCustomWidgetCollectionInterface>
#include <QDesignerCustomWidgetInterface>
#include <QDesignerDynamicPropertySheetExtension>
#include <QDesignerFormEditorInterface>
#include <QDesignerFormWindowCursorInterface>
#include <QDesignerFormWindowInterface>
#include <QDesignerFormWindowManagerInterface>
#include <QDesignerMemberSheetExtension>
#include <QDesignerObjectInspectorInterface>
#include <QDesignerPropertyEditorInterface>
#include <QDesignerPropertySheetExtension>
#include <QDesignerTaskMenuExtension>
#include <QDesignerWidgetBoxInterface>
#include <QDir>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QFormBuilder>
#include <QHideEvent>
#include <QIODevice>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPixmap>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QStringList>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

class MyQAbstractExtensionFactory: public QAbstractExtensionFactory
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QObject * extension(QObject * object, const QString & iid) const { return static_cast<QObject*>(callbackQAbstractExtensionFactory_Extension(const_cast<MyQAbstractExtensionFactory*>(this), this->objectNameAbs().toUtf8().data(), object, iid.toUtf8().data())); };
};

void* QAbstractExtensionFactory_Extension(void* ptr, void* object, char* iid)
{
	return static_cast<QAbstractExtensionFactory*>(ptr)->extension(static_cast<QObject*>(object), QString(iid));
}

void QAbstractExtensionFactory_DestroyQAbstractExtensionFactory(void* ptr)
{
	static_cast<QAbstractExtensionFactory*>(ptr)->~QAbstractExtensionFactory();
}

char* QAbstractExtensionFactory_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAbstractExtensionFactory*>(static_cast<QAbstractExtensionFactory*>(ptr))) {
		return static_cast<MyQAbstractExtensionFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractExtensionFactory_BASE").toUtf8().data();
}

void QAbstractExtensionFactory_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAbstractExtensionFactory*>(static_cast<QAbstractExtensionFactory*>(ptr))) {
		static_cast<MyQAbstractExtensionFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAbstractExtensionManager: public QAbstractExtensionManager
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QObject * extension(QObject * object, const QString & iid) const { return static_cast<QObject*>(callbackQAbstractExtensionManager_Extension(const_cast<MyQAbstractExtensionManager*>(this), this->objectNameAbs().toUtf8().data(), object, iid.toUtf8().data())); };
	void registerExtensions(QAbstractExtensionFactory * factory, const QString & iid) { callbackQAbstractExtensionManager_RegisterExtensions(this, this->objectNameAbs().toUtf8().data(), factory, iid.toUtf8().data()); };
	void unregisterExtensions(QAbstractExtensionFactory * factory, const QString & iid) { callbackQAbstractExtensionManager_UnregisterExtensions(this, this->objectNameAbs().toUtf8().data(), factory, iid.toUtf8().data()); };
};

void* QAbstractExtensionManager_Extension(void* ptr, void* object, char* iid)
{
	return static_cast<QAbstractExtensionManager*>(ptr)->extension(static_cast<QObject*>(object), QString(iid));
}

void QAbstractExtensionManager_RegisterExtensions(void* ptr, void* factory, char* iid)
{
	static_cast<QAbstractExtensionManager*>(ptr)->registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QAbstractExtensionManager_UnregisterExtensions(void* ptr, void* factory, char* iid)
{
	static_cast<QAbstractExtensionManager*>(ptr)->unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QAbstractExtensionManager_DestroyQAbstractExtensionManager(void* ptr)
{
	static_cast<QAbstractExtensionManager*>(ptr)->~QAbstractExtensionManager();
}

char* QAbstractExtensionManager_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAbstractExtensionManager*>(static_cast<QAbstractExtensionManager*>(ptr))) {
		return static_cast<MyQAbstractExtensionManager*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractExtensionManager_BASE").toUtf8().data();
}

void QAbstractExtensionManager_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAbstractExtensionManager*>(static_cast<QAbstractExtensionManager*>(ptr))) {
		static_cast<MyQAbstractExtensionManager*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAbstractFormBuilder: public QAbstractFormBuilder
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQAbstractFormBuilder() : QAbstractFormBuilder() {};
	QWidget * load(QIODevice * device, QWidget * parent) { return static_cast<QWidget*>(callbackQAbstractFormBuilder_Load(this, this->objectNameAbs().toUtf8().data(), device, parent)); };
	void save(QIODevice * device, QWidget * widget) { callbackQAbstractFormBuilder_Save(this, this->objectNameAbs().toUtf8().data(), device, widget); };
};

void* QAbstractFormBuilder_Load(void* ptr, void* device, void* parent)
{
	return static_cast<QAbstractFormBuilder*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
}

void* QAbstractFormBuilder_LoadDefault(void* ptr, void* device, void* parent)
{
	return static_cast<QAbstractFormBuilder*>(ptr)->QAbstractFormBuilder::load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
}

void QAbstractFormBuilder_Save(void* ptr, void* device, void* widget)
{
	static_cast<QAbstractFormBuilder*>(ptr)->save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
}

void QAbstractFormBuilder_SaveDefault(void* ptr, void* device, void* widget)
{
	static_cast<QAbstractFormBuilder*>(ptr)->QAbstractFormBuilder::save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
}

void* QAbstractFormBuilder_NewQAbstractFormBuilder()
{
	return new MyQAbstractFormBuilder();
}

char* QAbstractFormBuilder_ErrorString(void* ptr)
{
	return static_cast<QAbstractFormBuilder*>(ptr)->errorString().toUtf8().data();
}

void QAbstractFormBuilder_SetWorkingDirectory(void* ptr, void* directory)
{
	static_cast<QAbstractFormBuilder*>(ptr)->setWorkingDirectory(*static_cast<QDir*>(directory));
}

void* QAbstractFormBuilder_WorkingDirectory(void* ptr)
{
	return new QDir(static_cast<QAbstractFormBuilder*>(ptr)->workingDirectory());
}

void QAbstractFormBuilder_DestroyQAbstractFormBuilder(void* ptr)
{
	static_cast<QAbstractFormBuilder*>(ptr)->~QAbstractFormBuilder();
}

char* QAbstractFormBuilder_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAbstractFormBuilder*>(static_cast<QAbstractFormBuilder*>(ptr))) {
		return static_cast<MyQAbstractFormBuilder*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractFormBuilder_BASE").toUtf8().data();
}

void QAbstractFormBuilder_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAbstractFormBuilder*>(static_cast<QAbstractFormBuilder*>(ptr))) {
		static_cast<MyQAbstractFormBuilder*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerActionEditorInterface: public QDesignerActionEditorInterface
{
public:
	MyQDesignerActionEditorInterface(QWidget *parent, Qt::WindowFlags flags) : QDesignerActionEditorInterface(parent, flags) {};
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerActionEditorInterface_Core(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void manageAction(QAction * action) { callbackQDesignerActionEditorInterface_ManageAction(this, this->objectName().toUtf8().data(), action); };
	void setFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerActionEditorInterface_SetFormWindow(this, this->objectName().toUtf8().data(), formWindow); };
	void unmanageAction(QAction * action) { callbackQDesignerActionEditorInterface_UnmanageAction(this, this->objectName().toUtf8().data(), action); };
	void actionEvent(QActionEvent * event) { callbackQDesignerActionEditorInterface_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerActionEditorInterface_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerActionEditorInterface_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerActionEditorInterface_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerActionEditorInterface_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQDesignerActionEditorInterface_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerActionEditorInterface_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerActionEditorInterface_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQDesignerActionEditorInterface_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQDesignerActionEditorInterface_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerActionEditorInterface_Metric(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerActionEditorInterface_MinimumSizeHint(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerActionEditorInterface_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerActionEditorInterface_PaintEngine(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerActionEditorInterface_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQDesignerActionEditorInterface_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQDesignerActionEditorInterface_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQDesignerActionEditorInterface_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQDesignerActionEditorInterface_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQDesignerActionEditorInterface_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQDesignerActionEditorInterface_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerActionEditorInterface_SizeHint(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQDesignerActionEditorInterface_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQDesignerActionEditorInterface_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQDesignerActionEditorInterface_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerActionEditorInterface_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * event) { return callbackQDesignerActionEditorInterface_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerActionEditorInterface_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQDesignerActionEditorInterface_HasHeightForWidth(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerActionEditorInterface_HeightForWidth(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQDesignerActionEditorInterface_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQDesignerActionEditorInterface_InitPainter(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerActionEditorInterface_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerActionEditorInterface_InputMethodQuery(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerActionEditorInterface_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerActionEditorInterface_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQDesignerActionEditorInterface_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerActionEditorInterface_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQDesignerActionEditorInterface_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQDesignerActionEditorInterface_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerActionEditorInterface_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQDesignerActionEditorInterface_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQDesignerActionEditorInterface_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQDesignerActionEditorInterface_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQDesignerActionEditorInterface_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQDesignerActionEditorInterface_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQDesignerActionEditorInterface_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQDesignerActionEditorInterface_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQDesignerActionEditorInterface_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerActionEditorInterface_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQDesignerActionEditorInterface_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQDesignerActionEditorInterface_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerActionEditorInterface_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerActionEditorInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerActionEditorInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerActionEditorInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerActionEditorInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerActionEditorInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerActionEditorInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerActionEditorInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerActionEditorInterface_MetaObject(const_cast<MyQDesignerActionEditorInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerActionEditorInterface_NewQDesignerActionEditorInterface(void* parent, int flags)
{
	return new MyQDesignerActionEditorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDesignerActionEditorInterface_Core(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->core();
}

void* QDesignerActionEditorInterface_CoreDefault(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::core();
}

void QDesignerActionEditorInterface_ManageAction(void* ptr, void* action)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->manageAction(static_cast<QAction*>(action));
}

void QDesignerActionEditorInterface_SetFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerActionEditorInterface_UnmanageAction(void* ptr, void* action)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->unmanageAction(static_cast<QAction*>(action));
}

void QDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->~QDesignerActionEditorInterface();
}

void QDesignerActionEditorInterface_ActionEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerActionEditorInterface_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerActionEditorInterface_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerActionEditorInterface_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerActionEditorInterface_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerActionEditorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerActionEditorInterface_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerActionEditorInterface_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerActionEditorInterface_DropEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerActionEditorInterface_DropEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerActionEditorInterface_EnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_FocusInEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_HideEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerActionEditorInterface_HideEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerActionEditorInterface_LeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::leaveEvent(static_cast<QEvent*>(event));
}

int QDesignerActionEditorInterface_Metric(void* ptr, int m)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QDesignerActionEditorInterface_MetricDefault(void* ptr, int m)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QDesignerActionEditorInterface_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->minimumSizeHint()).height());
}

void* QDesignerActionEditorInterface_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::minimumSizeHint()).height());
}

void QDesignerActionEditorInterface_MoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerActionEditorInterface_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QDesignerActionEditorInterface_PaintEngine(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->paintEngine();
}

void* QDesignerActionEditorInterface_PaintEngineDefault(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::paintEngine();
}

void QDesignerActionEditorInterface_PaintEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerActionEditorInterface_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerActionEditorInterface_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QDesignerActionEditorInterface_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setEnabled(vbo != 0);
}

void QDesignerActionEditorInterface_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QDesignerActionEditorInterface_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setStyleSheet(QString(styleSheet));
}

void QDesignerActionEditorInterface_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QDesignerActionEditorInterface_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setVisible(visible != 0);
}

void QDesignerActionEditorInterface_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QDesignerActionEditorInterface_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setWindowModified(vbo != 0);
}

void QDesignerActionEditorInterface_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QDesignerActionEditorInterface_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setWindowTitle(QString(vqs));
}

void QDesignerActionEditorInterface_ShowEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerActionEditorInterface_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void* QDesignerActionEditorInterface_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->sizeHint()).height());
}

void* QDesignerActionEditorInterface_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::sizeHint()).height());
}

void QDesignerActionEditorInterface_ChangeEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::changeEvent(static_cast<QEvent*>(event));
}

int QDesignerActionEditorInterface_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDesignerActionEditorInterface_CloseDefault(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::close();
}

void QDesignerActionEditorInterface_CloseEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerActionEditorInterface_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerActionEditorInterface_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerActionEditorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QDesignerActionEditorInterface_Event(void* ptr, void* event)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDesignerActionEditorInterface_EventDefault(void* ptr, void* event)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::event(static_cast<QEvent*>(event));
}

int QDesignerActionEditorInterface_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->focusNextPrevChild(next != 0);
}

int QDesignerActionEditorInterface_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusNextPrevChild(next != 0);
}

int QDesignerActionEditorInterface_HasHeightForWidth(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->hasHeightForWidth();
}

int QDesignerActionEditorInterface_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hasHeightForWidth();
}

int QDesignerActionEditorInterface_HeightForWidth(void* ptr, int w)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->heightForWidth(w);
}

int QDesignerActionEditorInterface_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::heightForWidth(w);
}

void QDesignerActionEditorInterface_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "hide");
}

void QDesignerActionEditorInterface_HideDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hide();
}

void QDesignerActionEditorInterface_InitPainter(void* ptr, void* painter)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QDesignerActionEditorInterface_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::initPainter(static_cast<QPainter*>(painter));
}

void QDesignerActionEditorInterface_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerActionEditorInterface_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QDesignerActionEditorInterface_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerActionEditorInterface*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QDesignerActionEditorInterface_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QDesignerActionEditorInterface_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "lower");
}

void QDesignerActionEditorInterface_LowerDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::lower();
}

void QDesignerActionEditorInterface_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MousePressEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QDesignerActionEditorInterface_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QDesignerActionEditorInterface_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::nativeEvent(QByteArray(eventType), message, &result);
}

void QDesignerActionEditorInterface_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "raise");
}

void QDesignerActionEditorInterface_RaiseDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::raise();
}

void QDesignerActionEditorInterface_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "repaint");
}

void QDesignerActionEditorInterface_RepaintDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::repaint();
}

void QDesignerActionEditorInterface_ResizeEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerActionEditorInterface_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerActionEditorInterface_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QDesignerActionEditorInterface_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setDisabled(disable != 0);
}

void QDesignerActionEditorInterface_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setFocus");
}

void QDesignerActionEditorInterface_SetFocus2Default(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setFocus();
}

void QDesignerActionEditorInterface_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QDesignerActionEditorInterface_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setHidden(hidden != 0);
}

void QDesignerActionEditorInterface_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "show");
}

void QDesignerActionEditorInterface_ShowDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::show();
}

void QDesignerActionEditorInterface_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "showFullScreen");
}

void QDesignerActionEditorInterface_ShowFullScreenDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showFullScreen();
}

void QDesignerActionEditorInterface_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "showMaximized");
}

void QDesignerActionEditorInterface_ShowMaximizedDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showMaximized();
}

void QDesignerActionEditorInterface_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "showMinimized");
}

void QDesignerActionEditorInterface_ShowMinimizedDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showMinimized();
}

void QDesignerActionEditorInterface_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "showNormal");
}

void QDesignerActionEditorInterface_ShowNormalDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showNormal();
}

void QDesignerActionEditorInterface_TabletEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerActionEditorInterface_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerActionEditorInterface_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "update");
}

void QDesignerActionEditorInterface_UpdateDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::update();
}

void QDesignerActionEditorInterface_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "updateMicroFocus");
}

void QDesignerActionEditorInterface_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::updateMicroFocus();
}

void QDesignerActionEditorInterface_WheelEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerActionEditorInterface_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerActionEditorInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerActionEditorInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerActionEditorInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerActionEditorInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerActionEditorInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerActionEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerActionEditorInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerActionEditorInterface*>(ptr), "deleteLater");
}

void QDesignerActionEditorInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::deleteLater();
}

void QDesignerActionEditorInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerActionEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerActionEditorInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerActionEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerActionEditorInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerActionEditorInterface*>(ptr)->metaObject());
}

void* QDesignerActionEditorInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::metaObject());
}

class MyQDesignerContainerExtension: public QDesignerContainerExtension
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void addWidget(QWidget * page) { callbackQDesignerContainerExtension_AddWidget(this, this->objectNameAbs().toUtf8().data(), page); };
	bool canAddWidget() const { return callbackQDesignerContainerExtension_CanAddWidget(const_cast<MyQDesignerContainerExtension*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	bool canRemove(int index) const { return callbackQDesignerContainerExtension_CanRemove(const_cast<MyQDesignerContainerExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	int count() const { return callbackQDesignerContainerExtension_Count(const_cast<MyQDesignerContainerExtension*>(this), this->objectNameAbs().toUtf8().data()); };
	int currentIndex() const { return callbackQDesignerContainerExtension_CurrentIndex(const_cast<MyQDesignerContainerExtension*>(this), this->objectNameAbs().toUtf8().data()); };
	void insertWidget(int index, QWidget * page) { callbackQDesignerContainerExtension_InsertWidget(this, this->objectNameAbs().toUtf8().data(), index, page); };
	void remove(int index) { callbackQDesignerContainerExtension_Remove(this, this->objectNameAbs().toUtf8().data(), index); };
	void setCurrentIndex(int index) { callbackQDesignerContainerExtension_SetCurrentIndex(this, this->objectNameAbs().toUtf8().data(), index); };
	QWidget * widget(int index) const { return static_cast<QWidget*>(callbackQDesignerContainerExtension_Widget(const_cast<MyQDesignerContainerExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
};

void QDesignerContainerExtension_AddWidget(void* ptr, void* page)
{
	static_cast<QDesignerContainerExtension*>(ptr)->addWidget(static_cast<QWidget*>(page));
}

int QDesignerContainerExtension_CanAddWidget(void* ptr)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->canAddWidget();
}

int QDesignerContainerExtension_CanAddWidgetDefault(void* ptr)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->QDesignerContainerExtension::canAddWidget();
}

int QDesignerContainerExtension_CanRemove(void* ptr, int index)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->canRemove(index);
}

int QDesignerContainerExtension_CanRemoveDefault(void* ptr, int index)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->QDesignerContainerExtension::canRemove(index);
}

int QDesignerContainerExtension_Count(void* ptr)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->count();
}

int QDesignerContainerExtension_CurrentIndex(void* ptr)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->currentIndex();
}

void QDesignerContainerExtension_InsertWidget(void* ptr, int index, void* page)
{
	static_cast<QDesignerContainerExtension*>(ptr)->insertWidget(index, static_cast<QWidget*>(page));
}

void QDesignerContainerExtension_Remove(void* ptr, int index)
{
	static_cast<QDesignerContainerExtension*>(ptr)->remove(index);
}

void QDesignerContainerExtension_SetCurrentIndex(void* ptr, int index)
{
	static_cast<QDesignerContainerExtension*>(ptr)->setCurrentIndex(index);
}

void* QDesignerContainerExtension_Widget(void* ptr, int index)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->widget(index);
}

void QDesignerContainerExtension_DestroyQDesignerContainerExtension(void* ptr)
{
	static_cast<QDesignerContainerExtension*>(ptr)->~QDesignerContainerExtension();
}

char* QDesignerContainerExtension_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerContainerExtension*>(static_cast<QDesignerContainerExtension*>(ptr))) {
		return static_cast<MyQDesignerContainerExtension*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerContainerExtension_BASE").toUtf8().data();
}

void QDesignerContainerExtension_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerContainerExtension*>(static_cast<QDesignerContainerExtension*>(ptr))) {
		static_cast<MyQDesignerContainerExtension*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerCustomWidgetCollectionInterface: public QDesignerCustomWidgetCollectionInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(void* ptr)
{
	static_cast<QDesignerCustomWidgetCollectionInterface*>(ptr)->~QDesignerCustomWidgetCollectionInterface();
}

char* QDesignerCustomWidgetCollectionInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerCustomWidgetCollectionInterface*>(static_cast<QDesignerCustomWidgetCollectionInterface*>(ptr))) {
		return static_cast<MyQDesignerCustomWidgetCollectionInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerCustomWidgetCollectionInterface_BASE").toUtf8().data();
}

void QDesignerCustomWidgetCollectionInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerCustomWidgetCollectionInterface*>(static_cast<QDesignerCustomWidgetCollectionInterface*>(ptr))) {
		static_cast<MyQDesignerCustomWidgetCollectionInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerCustomWidgetInterface: public QDesignerCustomWidgetInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QString codeTemplate() const { return QString(callbackQDesignerCustomWidgetInterface_CodeTemplate(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QWidget * createWidget(QWidget * parent) { return static_cast<QWidget*>(callbackQDesignerCustomWidgetInterface_CreateWidget(this, this->objectNameAbs().toUtf8().data(), parent)); };
	QString domXml() const { return QString(callbackQDesignerCustomWidgetInterface_DomXml(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QString group() const { return QString(callbackQDesignerCustomWidgetInterface_Group(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QIcon icon() const { return *static_cast<QIcon*>(callbackQDesignerCustomWidgetInterface_Icon(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QString includeFile() const { return QString(callbackQDesignerCustomWidgetInterface_IncludeFile(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	void initialize(QDesignerFormEditorInterface * formEditor) { callbackQDesignerCustomWidgetInterface_Initialize(this, this->objectNameAbs().toUtf8().data(), formEditor); };
	bool isContainer() const { return callbackQDesignerCustomWidgetInterface_IsContainer(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	bool isInitialized() const { return callbackQDesignerCustomWidgetInterface_IsInitialized(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	QString name() const { return QString(callbackQDesignerCustomWidgetInterface_Name(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QString toolTip() const { return QString(callbackQDesignerCustomWidgetInterface_ToolTip(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QString whatsThis() const { return QString(callbackQDesignerCustomWidgetInterface_WhatsThis(const_cast<MyQDesignerCustomWidgetInterface*>(this), this->objectNameAbs().toUtf8().data())); };
};

char* QDesignerCustomWidgetInterface_CodeTemplate(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->codeTemplate().toUtf8().data();
}

char* QDesignerCustomWidgetInterface_CodeTemplateDefault(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::codeTemplate().toUtf8().data();
}

void* QDesignerCustomWidgetInterface_CreateWidget(void* ptr, void* parent)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->createWidget(static_cast<QWidget*>(parent));
}

char* QDesignerCustomWidgetInterface_DomXml(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->domXml().toUtf8().data();
}

char* QDesignerCustomWidgetInterface_DomXmlDefault(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::domXml().toUtf8().data();
}

char* QDesignerCustomWidgetInterface_Group(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->group().toUtf8().data();
}

void* QDesignerCustomWidgetInterface_Icon(void* ptr)
{
	return new QIcon(static_cast<QDesignerCustomWidgetInterface*>(ptr)->icon());
}

char* QDesignerCustomWidgetInterface_IncludeFile(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->includeFile().toUtf8().data();
}

void QDesignerCustomWidgetInterface_Initialize(void* ptr, void* formEditor)
{
	static_cast<QDesignerCustomWidgetInterface*>(ptr)->initialize(static_cast<QDesignerFormEditorInterface*>(formEditor));
}

void QDesignerCustomWidgetInterface_InitializeDefault(void* ptr, void* formEditor)
{
	static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::initialize(static_cast<QDesignerFormEditorInterface*>(formEditor));
}

int QDesignerCustomWidgetInterface_IsContainer(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->isContainer();
}

int QDesignerCustomWidgetInterface_IsInitialized(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->isInitialized();
}

int QDesignerCustomWidgetInterface_IsInitializedDefault(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::isInitialized();
}

char* QDesignerCustomWidgetInterface_Name(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->name().toUtf8().data();
}

char* QDesignerCustomWidgetInterface_ToolTip(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->toolTip().toUtf8().data();
}

char* QDesignerCustomWidgetInterface_WhatsThis(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->whatsThis().toUtf8().data();
}

void QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(void* ptr)
{
	static_cast<QDesignerCustomWidgetInterface*>(ptr)->~QDesignerCustomWidgetInterface();
}

char* QDesignerCustomWidgetInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerCustomWidgetInterface*>(static_cast<QDesignerCustomWidgetInterface*>(ptr))) {
		return static_cast<MyQDesignerCustomWidgetInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerCustomWidgetInterface_BASE").toUtf8().data();
}

void QDesignerCustomWidgetInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerCustomWidgetInterface*>(static_cast<QDesignerCustomWidgetInterface*>(ptr))) {
		static_cast<MyQDesignerCustomWidgetInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerDynamicPropertySheetExtension: public QDesignerDynamicPropertySheetExtension
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	int addDynamicProperty(const QString & propertyName, const QVariant & value) { return callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty(this, this->objectNameAbs().toUtf8().data(), propertyName.toUtf8().data(), new QVariant(value)); };
	bool canAddDynamicProperty(const QString & propertyName) const { return callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(const_cast<MyQDesignerDynamicPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), propertyName.toUtf8().data()) != 0; };
	bool dynamicPropertiesAllowed() const { return callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(const_cast<MyQDesignerDynamicPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	bool isDynamicProperty(int index) const { return callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty(const_cast<MyQDesignerDynamicPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool removeDynamicProperty(int index) { return callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(this, this->objectNameAbs().toUtf8().data(), index) != 0; };
};

int QDesignerDynamicPropertySheetExtension_AddDynamicProperty(void* ptr, char* propertyName, void* value)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->addDynamicProperty(QString(propertyName), *static_cast<QVariant*>(value));
}

int QDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(void* ptr, char* propertyName)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->canAddDynamicProperty(QString(propertyName));
}

int QDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(void* ptr)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->dynamicPropertiesAllowed();
}

int QDesignerDynamicPropertySheetExtension_IsDynamicProperty(void* ptr, int index)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->isDynamicProperty(index);
}

int QDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(void* ptr, int index)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->removeDynamicProperty(index);
}

void QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(void* ptr)
{
	static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->~QDesignerDynamicPropertySheetExtension();
}

char* QDesignerDynamicPropertySheetExtension_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerDynamicPropertySheetExtension*>(static_cast<QDesignerDynamicPropertySheetExtension*>(ptr))) {
		return static_cast<MyQDesignerDynamicPropertySheetExtension*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerDynamicPropertySheetExtension_BASE").toUtf8().data();
}

void QDesignerDynamicPropertySheetExtension_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerDynamicPropertySheetExtension*>(static_cast<QDesignerDynamicPropertySheetExtension*>(ptr))) {
		static_cast<MyQDesignerDynamicPropertySheetExtension*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerFormEditorInterface: public QDesignerFormEditorInterface
{
public:
	MyQDesignerFormEditorInterface(QObject *parent) : QDesignerFormEditorInterface(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormEditorInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerFormEditorInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormEditorInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormEditorInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerFormEditorInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormEditorInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQDesignerFormEditorInterface_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormEditorInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormEditorInterface_MetaObject(const_cast<MyQDesignerFormEditorInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerFormEditorInterface_NewQDesignerFormEditorInterface(void* parent)
{
	return new MyQDesignerFormEditorInterface(static_cast<QObject*>(parent));
}

void* QDesignerFormEditorInterface_ActionEditor(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->actionEditor();
}

void* QDesignerFormEditorInterface_ExtensionManager(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->extensionManager();
}

void* QDesignerFormEditorInterface_FormWindowManager(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->formWindowManager();
}

void* QDesignerFormEditorInterface_ObjectInspector(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->objectInspector();
}

void* QDesignerFormEditorInterface_PropertyEditor(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->propertyEditor();
}

void QDesignerFormEditorInterface_SetActionEditor(void* ptr, void* actionEditor)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->setActionEditor(static_cast<QDesignerActionEditorInterface*>(actionEditor));
}

void QDesignerFormEditorInterface_SetObjectInspector(void* ptr, void* objectInspector)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->setObjectInspector(static_cast<QDesignerObjectInspectorInterface*>(objectInspector));
}

void QDesignerFormEditorInterface_SetPropertyEditor(void* ptr, void* propertyEditor)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->setPropertyEditor(static_cast<QDesignerPropertyEditorInterface*>(propertyEditor));
}

void QDesignerFormEditorInterface_SetWidgetBox(void* ptr, void* widgetBox)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->setWidgetBox(static_cast<QDesignerWidgetBoxInterface*>(widgetBox));
}

void* QDesignerFormEditorInterface_TopLevel(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->topLevel();
}

void* QDesignerFormEditorInterface_WidgetBox(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->widgetBox();
}

void QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(void* ptr)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->~QDesignerFormEditorInterface();
}

void QDesignerFormEditorInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormEditorInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormEditorInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormEditorInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormEditorInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormEditorInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormEditorInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormEditorInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormEditorInterface*>(ptr), "deleteLater");
}

void QDesignerFormEditorInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::deleteLater();
}

void QDesignerFormEditorInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerFormEditorInterface_Event(void* ptr, void* e)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDesignerFormEditorInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::event(static_cast<QEvent*>(e));
}

int QDesignerFormEditorInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerFormEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerFormEditorInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormEditorInterface*>(ptr)->metaObject());
}

void* QDesignerFormEditorInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::metaObject());
}

class MyQDesignerFormWindowCursorInterface: public QDesignerFormWindowCursorInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QWidget * current() const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_Current(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QDesignerFormWindowInterface * formWindow() const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowCursorInterface_FormWindow(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	bool hasSelection() const { return callbackQDesignerFormWindowCursorInterface_HasSelection(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	bool movePosition(QDesignerFormWindowCursorInterface::MoveOperation operation, QDesignerFormWindowCursorInterface::MoveMode mode) { return callbackQDesignerFormWindowCursorInterface_MovePosition(this, this->objectNameAbs().toUtf8().data(), operation, mode) != 0; };
	int position() const { return callbackQDesignerFormWindowCursorInterface_Position(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	void resetWidgetProperty(QWidget * widget, const QString & name) { callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty(this, this->objectNameAbs().toUtf8().data(), widget, name.toUtf8().data()); };
	QWidget * selectedWidget(int index) const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_SelectedWidget(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	int selectedWidgetCount() const { return callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	void setPosition(int position, QDesignerFormWindowCursorInterface::MoveMode mode) { callbackQDesignerFormWindowCursorInterface_SetPosition(this, this->objectNameAbs().toUtf8().data(), position, mode); };
	void setProperty(const QString & name, const QVariant & value) { callbackQDesignerFormWindowCursorInterface_SetProperty(this, this->objectNameAbs().toUtf8().data(), name.toUtf8().data(), new QVariant(value)); };
	void setWidgetProperty(QWidget * widget, const QString & name, const QVariant & value) { callbackQDesignerFormWindowCursorInterface_SetWidgetProperty(this, this->objectNameAbs().toUtf8().data(), widget, name.toUtf8().data(), new QVariant(value)); };
	QWidget * widget(int index) const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_Widget(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	int widgetCount() const { return callbackQDesignerFormWindowCursorInterface_WidgetCount(const_cast<MyQDesignerFormWindowCursorInterface*>(this), this->objectNameAbs().toUtf8().data()); };
};

void* QDesignerFormWindowCursorInterface_Current(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->current();
}

void* QDesignerFormWindowCursorInterface_FormWindow(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->formWindow();
}

int QDesignerFormWindowCursorInterface_HasSelection(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->hasSelection();
}

int QDesignerFormWindowCursorInterface_IsWidgetSelected(void* ptr, void* widget)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->isWidgetSelected(static_cast<QWidget*>(widget));
}

int QDesignerFormWindowCursorInterface_MovePosition(void* ptr, int operation, int mode)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->movePosition(static_cast<QDesignerFormWindowCursorInterface::MoveOperation>(operation), static_cast<QDesignerFormWindowCursorInterface::MoveMode>(mode));
}

int QDesignerFormWindowCursorInterface_Position(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->position();
}

void QDesignerFormWindowCursorInterface_ResetWidgetProperty(void* ptr, void* widget, char* name)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->resetWidgetProperty(static_cast<QWidget*>(widget), QString(name));
}

void* QDesignerFormWindowCursorInterface_SelectedWidget(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->selectedWidget(index);
}

int QDesignerFormWindowCursorInterface_SelectedWidgetCount(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->selectedWidgetCount();
}

void QDesignerFormWindowCursorInterface_SetPosition(void* ptr, int position, int mode)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setPosition(position, static_cast<QDesignerFormWindowCursorInterface::MoveMode>(mode));
}

void QDesignerFormWindowCursorInterface_SetProperty(void* ptr, char* name, void* value)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setProperty(QString(name), *static_cast<QVariant*>(value));
}

void QDesignerFormWindowCursorInterface_SetWidgetProperty(void* ptr, void* widget, char* name, void* value)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setWidgetProperty(static_cast<QWidget*>(widget), QString(name), *static_cast<QVariant*>(value));
}

void* QDesignerFormWindowCursorInterface_Widget(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->widget(index);
}

int QDesignerFormWindowCursorInterface_WidgetCount(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->widgetCount();
}

void QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(void* ptr)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->~QDesignerFormWindowCursorInterface();
}

char* QDesignerFormWindowCursorInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerFormWindowCursorInterface*>(static_cast<QDesignerFormWindowCursorInterface*>(ptr))) {
		return static_cast<MyQDesignerFormWindowCursorInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerFormWindowCursorInterface_BASE").toUtf8().data();
}

void QDesignerFormWindowCursorInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerFormWindowCursorInterface*>(static_cast<QDesignerFormWindowCursorInterface*>(ptr))) {
		static_cast<MyQDesignerFormWindowCursorInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerFormWindowInterface: public QDesignerFormWindowInterface
{
public:
	void Signal_AboutToUnmanageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_AboutToUnmanageWidget(this, this->objectName().toUtf8().data(), widget); };
	QDir absoluteDir() const { return *static_cast<QDir*>(callbackQDesignerFormWindowInterface_AbsoluteDir(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void activateResourceFilePaths(const QStringList & paths, int * errorCount, QString * errorMessages) { callbackQDesignerFormWindowInterface_ActivateResourceFilePaths(this, this->objectName().toUtf8().data(), paths.join("|").toUtf8().data(), *errorCount, errorMessages->toUtf8().data()); };
	void Signal_Activated(QWidget * widget) { callbackQDesignerFormWindowInterface_Activated(this, this->objectName().toUtf8().data(), widget); };
	void addResourceFile(const QString & path) { callbackQDesignerFormWindowInterface_AddResourceFile(this, this->objectName().toUtf8().data(), path.toUtf8().data()); };
	QString author() const { return QString(callbackQDesignerFormWindowInterface_Author(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_Changed() { callbackQDesignerFormWindowInterface_Changed(this, this->objectName().toUtf8().data()); };
	QStringList checkContents() const { return QString(callbackQDesignerFormWindowInterface_CheckContents(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	void clearSelection(bool update) { callbackQDesignerFormWindowInterface_ClearSelection(this, this->objectName().toUtf8().data(), update); };
	QString comment() const { return QString(callbackQDesignerFormWindowInterface_Comment(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	QString contents() const { return QString(callbackQDesignerFormWindowInterface_Contents(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerFormWindowInterface_Core(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	QDesignerFormWindowCursorInterface * cursor() const { return static_cast<QDesignerFormWindowCursorInterface*>(callbackQDesignerFormWindowInterface_Cursor(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void emitSelectionChanged() { callbackQDesignerFormWindowInterface_EmitSelectionChanged(this, this->objectName().toUtf8().data()); };
	QString exportMacro() const { return QString(callbackQDesignerFormWindowInterface_ExportMacro(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_FeatureChanged(QDesignerFormWindowInterface::Feature feature) { callbackQDesignerFormWindowInterface_FeatureChanged(this, this->objectName().toUtf8().data(), feature); };
	Feature features() const { return static_cast<QDesignerFormWindowInterface::FeatureFlag>(callbackQDesignerFormWindowInterface_Features(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	QString fileName() const { return QString(callbackQDesignerFormWindowInterface_FileName(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_FileNameChanged(const QString & fileName) { callbackQDesignerFormWindowInterface_FileNameChanged(this, this->objectName().toUtf8().data(), fileName.toUtf8().data()); };
	QWidget * formContainer() const { return static_cast<QWidget*>(callbackQDesignerFormWindowInterface_FormContainer(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_GeometryChanged() { callbackQDesignerFormWindowInterface_GeometryChanged(this, this->objectName().toUtf8().data()); };
	QPoint grid() const { return *static_cast<QPoint*>(callbackQDesignerFormWindowInterface_Grid(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	bool hasFeature(QDesignerFormWindowInterface::Feature feature) const { return callbackQDesignerFormWindowInterface_HasFeature(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), feature) != 0; };
	QStringList includeHints() const { return QString(callbackQDesignerFormWindowInterface_IncludeHints(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	bool isDirty() const { return callbackQDesignerFormWindowInterface_IsDirty(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	bool isManaged(QWidget * widget) const { return callbackQDesignerFormWindowInterface_IsManaged(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), widget) != 0; };
	void layoutDefault(int * margin, int * spacing) { callbackQDesignerFormWindowInterface_LayoutDefault(this, this->objectName().toUtf8().data(), *margin, *spacing); };
	void layoutFunction(QString * margin, QString * spacing) { callbackQDesignerFormWindowInterface_LayoutFunction(this, this->objectName().toUtf8().data(), margin->toUtf8().data(), spacing->toUtf8().data()); };
	void Signal_MainContainerChanged(QWidget * mainContainer) { callbackQDesignerFormWindowInterface_MainContainerChanged(this, this->objectName().toUtf8().data(), mainContainer); };
	void manageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_ManageWidget(this, this->objectName().toUtf8().data(), widget); };
	void Signal_ObjectRemoved(QObject * object) { callbackQDesignerFormWindowInterface_ObjectRemoved(this, this->objectName().toUtf8().data(), object); };
	QString pixmapFunction() const { return QString(callbackQDesignerFormWindowInterface_PixmapFunction(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void removeResourceFile(const QString & path) { callbackQDesignerFormWindowInterface_RemoveResourceFile(this, this->objectName().toUtf8().data(), path.toUtf8().data()); };
	ResourceFileSaveMode resourceFileSaveMode() const { return static_cast<QDesignerFormWindowInterface::ResourceFileSaveMode>(callbackQDesignerFormWindowInterface_ResourceFileSaveMode(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	QStringList resourceFiles() const { return QString(callbackQDesignerFormWindowInterface_ResourceFiles(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	void Signal_ResourceFilesChanged() { callbackQDesignerFormWindowInterface_ResourceFilesChanged(this, this->objectName().toUtf8().data()); };
	void selectWidget(QWidget * widget, bool sele) { callbackQDesignerFormWindowInterface_SelectWidget(this, this->objectName().toUtf8().data(), widget, sele); };
	void Signal_SelectionChanged() { callbackQDesignerFormWindowInterface_SelectionChanged(this, this->objectName().toUtf8().data()); };
	void setAuthor(const QString & author) { callbackQDesignerFormWindowInterface_SetAuthor(this, this->objectName().toUtf8().data(), author.toUtf8().data()); };
	void setComment(const QString & comment) { callbackQDesignerFormWindowInterface_SetComment(this, this->objectName().toUtf8().data(), comment.toUtf8().data()); };
	bool setContents(QIODevice * device, QString * errorMessage) { return callbackQDesignerFormWindowInterface_SetContents(this, this->objectName().toUtf8().data(), device, errorMessage->toUtf8().data()) != 0; };
	bool setContents(const QString & contents) { return callbackQDesignerFormWindowInterface_SetContents2(this, this->objectName().toUtf8().data(), contents.toUtf8().data()) != 0; };
	void setDirty(bool dirty) { callbackQDesignerFormWindowInterface_SetDirty(this, this->objectName().toUtf8().data(), dirty); };
	void setExportMacro(const QString & exportMacro) { callbackQDesignerFormWindowInterface_SetExportMacro(this, this->objectName().toUtf8().data(), exportMacro.toUtf8().data()); };
	void setFeatures(QDesignerFormWindowInterface::Feature features) { callbackQDesignerFormWindowInterface_SetFeatures(this, this->objectName().toUtf8().data(), features); };
	void setFileName(const QString & fileName) { callbackQDesignerFormWindowInterface_SetFileName(this, this->objectName().toUtf8().data(), fileName.toUtf8().data()); };
	void setGrid(const QPoint & grid) { callbackQDesignerFormWindowInterface_SetGrid(this, this->objectName().toUtf8().data(), new QPoint(static_cast<QPoint>(grid).x(), static_cast<QPoint>(grid).y())); };
	void setIncludeHints(const QStringList & includeHints) { callbackQDesignerFormWindowInterface_SetIncludeHints(this, this->objectName().toUtf8().data(), includeHints.join("|").toUtf8().data()); };
	void setLayoutDefault(int margin, int spacing) { callbackQDesignerFormWindowInterface_SetLayoutDefault(this, this->objectName().toUtf8().data(), margin, spacing); };
	void setLayoutFunction(const QString & margin, const QString & spacing) { callbackQDesignerFormWindowInterface_SetLayoutFunction(this, this->objectName().toUtf8().data(), margin.toUtf8().data(), spacing.toUtf8().data()); };
	void setMainContainer(QWidget * mainContainer) { callbackQDesignerFormWindowInterface_SetMainContainer(this, this->objectName().toUtf8().data(), mainContainer); };
	void setPixmapFunction(const QString & pixmapFunction) { callbackQDesignerFormWindowInterface_SetPixmapFunction(this, this->objectName().toUtf8().data(), pixmapFunction.toUtf8().data()); };
	void setResourceFileSaveMode(QDesignerFormWindowInterface::ResourceFileSaveMode behaviour) { callbackQDesignerFormWindowInterface_SetResourceFileSaveMode(this, this->objectName().toUtf8().data(), behaviour); };
	void unmanageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_UnmanageWidget(this, this->objectName().toUtf8().data(), widget); };
	void Signal_WidgetManaged(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetManaged(this, this->objectName().toUtf8().data(), widget); };
	void Signal_WidgetRemoved(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetRemoved(this, this->objectName().toUtf8().data(), widget); };
	void Signal_WidgetUnmanaged(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetUnmanaged(this, this->objectName().toUtf8().data(), widget); };
	void actionEvent(QActionEvent * event) { callbackQDesignerFormWindowInterface_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerFormWindowInterface_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerFormWindowInterface_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerFormWindowInterface_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerFormWindowInterface_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQDesignerFormWindowInterface_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerFormWindowInterface_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerFormWindowInterface_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQDesignerFormWindowInterface_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQDesignerFormWindowInterface_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerFormWindowInterface_Metric(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerFormWindowInterface_MinimumSizeHint(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerFormWindowInterface_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerFormWindowInterface_PaintEngine(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerFormWindowInterface_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQDesignerFormWindowInterface_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQDesignerFormWindowInterface_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQDesignerFormWindowInterface_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQDesignerFormWindowInterface_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQDesignerFormWindowInterface_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQDesignerFormWindowInterface_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerFormWindowInterface_SizeHint(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQDesignerFormWindowInterface_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQDesignerFormWindowInterface_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQDesignerFormWindowInterface_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerFormWindowInterface_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * event) { return callbackQDesignerFormWindowInterface_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerFormWindowInterface_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQDesignerFormWindowInterface_HasHeightForWidth(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerFormWindowInterface_HeightForWidth(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQDesignerFormWindowInterface_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQDesignerFormWindowInterface_InitPainter(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerFormWindowInterface_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerFormWindowInterface_InputMethodQuery(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerFormWindowInterface_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerFormWindowInterface_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQDesignerFormWindowInterface_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerFormWindowInterface_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQDesignerFormWindowInterface_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQDesignerFormWindowInterface_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerFormWindowInterface_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQDesignerFormWindowInterface_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQDesignerFormWindowInterface_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQDesignerFormWindowInterface_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQDesignerFormWindowInterface_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQDesignerFormWindowInterface_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQDesignerFormWindowInterface_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQDesignerFormWindowInterface_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQDesignerFormWindowInterface_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerFormWindowInterface_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQDesignerFormWindowInterface_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQDesignerFormWindowInterface_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerFormWindowInterface_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormWindowInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerFormWindowInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormWindowInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerFormWindowInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormWindowInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormWindowInterface_MetaObject(const_cast<MyQDesignerFormWindowInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(void* object)
{
	return QDesignerFormWindowInterface::findFormWindow(static_cast<QObject*>(object));
}

void* QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(void* widget)
{
	return QDesignerFormWindowInterface::findFormWindow(static_cast<QWidget*>(widget));
}

void QDesignerFormWindowInterface_ConnectAboutToUnmanageWidget(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::aboutToUnmanageWidget), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_AboutToUnmanageWidget));
}

void QDesignerFormWindowInterface_DisconnectAboutToUnmanageWidget(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::aboutToUnmanageWidget), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_AboutToUnmanageWidget));
}

void QDesignerFormWindowInterface_AboutToUnmanageWidget(void* ptr, void* widget)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->aboutToUnmanageWidget(static_cast<QWidget*>(widget));
}

void* QDesignerFormWindowInterface_AbsoluteDir(void* ptr)
{
	return new QDir(static_cast<QDesignerFormWindowInterface*>(ptr)->absoluteDir());
}

void QDesignerFormWindowInterface_ActivateResourceFilePaths(void* ptr, char* paths, int errorCount, char* errorMessages)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "activateResourceFilePaths", Q_ARG(QStringList, QString(paths).split("|", QString::SkipEmptyParts)), Q_ARG(int*, &errorCount), Q_ARG(QString*, new QString(errorMessages)));
}

void QDesignerFormWindowInterface_ConnectActivated(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::activated), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_Activated));
}

void QDesignerFormWindowInterface_DisconnectActivated(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::activated), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_Activated));
}

void QDesignerFormWindowInterface_Activated(void* ptr, void* widget)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->activated(static_cast<QWidget*>(widget));
}

char* QDesignerFormWindowInterface_ActiveResourceFilePaths(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->activeResourceFilePaths().join("|").toUtf8().data();
}

void QDesignerFormWindowInterface_AddResourceFile(void* ptr, char* path)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->addResourceFile(QString(path));
}

char* QDesignerFormWindowInterface_Author(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->author().toUtf8().data();
}

void QDesignerFormWindowInterface_ConnectChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::changed), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_Changed));
}

void QDesignerFormWindowInterface_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::changed), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_Changed));
}

void QDesignerFormWindowInterface_Changed(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->changed();
}

char* QDesignerFormWindowInterface_CheckContents(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->checkContents().join("|").toUtf8().data();
}

void QDesignerFormWindowInterface_ClearSelection(void* ptr, int update)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "clearSelection", Q_ARG(bool, update != 0));
}

char* QDesignerFormWindowInterface_Comment(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->comment().toUtf8().data();
}

char* QDesignerFormWindowInterface_Contents(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->contents().toUtf8().data();
}

void* QDesignerFormWindowInterface_Core(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->core();
}

void* QDesignerFormWindowInterface_CoreDefault(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::core();
}

void* QDesignerFormWindowInterface_Cursor(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->cursor();
}

void QDesignerFormWindowInterface_EmitSelectionChanged(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->emitSelectionChanged();
}

char* QDesignerFormWindowInterface_ExportMacro(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->exportMacro().toUtf8().data();
}

void QDesignerFormWindowInterface_ConnectFeatureChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&QDesignerFormWindowInterface::featureChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&MyQDesignerFormWindowInterface::Signal_FeatureChanged));
}

void QDesignerFormWindowInterface_DisconnectFeatureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&QDesignerFormWindowInterface::featureChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&MyQDesignerFormWindowInterface::Signal_FeatureChanged));
}

void QDesignerFormWindowInterface_FeatureChanged(void* ptr, int feature)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->featureChanged(static_cast<QDesignerFormWindowInterface::FeatureFlag>(feature));
}

int QDesignerFormWindowInterface_Features(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->features();
}

char* QDesignerFormWindowInterface_FileName(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->fileName().toUtf8().data();
}

void QDesignerFormWindowInterface_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(const QString &)>(&QDesignerFormWindowInterface::fileNameChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(const QString &)>(&MyQDesignerFormWindowInterface::Signal_FileNameChanged));
}

void QDesignerFormWindowInterface_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(const QString &)>(&QDesignerFormWindowInterface::fileNameChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(const QString &)>(&MyQDesignerFormWindowInterface::Signal_FileNameChanged));
}

void QDesignerFormWindowInterface_FileNameChanged(void* ptr, char* fileName)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->fileNameChanged(QString(fileName));
}

void* QDesignerFormWindowInterface_FormContainer(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->formContainer();
}

void QDesignerFormWindowInterface_ConnectGeometryChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::geometryChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_GeometryChanged));
}

void QDesignerFormWindowInterface_DisconnectGeometryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::geometryChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_GeometryChanged));
}

void QDesignerFormWindowInterface_GeometryChanged(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->geometryChanged();
}

void* QDesignerFormWindowInterface_Grid(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QDesignerFormWindowInterface*>(ptr)->grid()).x(), static_cast<QPoint>(static_cast<QDesignerFormWindowInterface*>(ptr)->grid()).y());
}

int QDesignerFormWindowInterface_HasFeature(void* ptr, int feature)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->hasFeature(static_cast<QDesignerFormWindowInterface::FeatureFlag>(feature));
}

char* QDesignerFormWindowInterface_IncludeHints(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->includeHints().join("|").toUtf8().data();
}

int QDesignerFormWindowInterface_IsDirty(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->isDirty();
}

int QDesignerFormWindowInterface_IsManaged(void* ptr, void* widget)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->isManaged(static_cast<QWidget*>(widget));
}

void QDesignerFormWindowInterface_LayoutDefault(void* ptr, int margin, int spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->layoutDefault(&margin, &spacing);
}

void QDesignerFormWindowInterface_LayoutFunction(void* ptr, char* margin, char* spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->layoutFunction(new QString(margin), new QString(spacing));
}

void QDesignerFormWindowInterface_ConnectMainContainerChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::mainContainerChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_MainContainerChanged));
}

void QDesignerFormWindowInterface_DisconnectMainContainerChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::mainContainerChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_MainContainerChanged));
}

void QDesignerFormWindowInterface_MainContainerChanged(void* ptr, void* mainContainer)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->mainContainerChanged(static_cast<QWidget*>(mainContainer));
}

void QDesignerFormWindowInterface_ManageWidget(void* ptr, void* widget)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "manageWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QDesignerFormWindowInterface_ConnectObjectRemoved(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QObject *)>(&QDesignerFormWindowInterface::objectRemoved), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QObject *)>(&MyQDesignerFormWindowInterface::Signal_ObjectRemoved));
}

void QDesignerFormWindowInterface_DisconnectObjectRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QObject *)>(&QDesignerFormWindowInterface::objectRemoved), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QObject *)>(&MyQDesignerFormWindowInterface::Signal_ObjectRemoved));
}

void QDesignerFormWindowInterface_ObjectRemoved(void* ptr, void* object)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->objectRemoved(static_cast<QObject*>(object));
}

char* QDesignerFormWindowInterface_PixmapFunction(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->pixmapFunction().toUtf8().data();
}

void QDesignerFormWindowInterface_RemoveResourceFile(void* ptr, char* path)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->removeResourceFile(QString(path));
}

int QDesignerFormWindowInterface_ResourceFileSaveMode(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->resourceFileSaveMode();
}

char* QDesignerFormWindowInterface_ResourceFiles(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->resourceFiles().join("|").toUtf8().data();
}

void QDesignerFormWindowInterface_ConnectResourceFilesChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::resourceFilesChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_ResourceFilesChanged));
}

void QDesignerFormWindowInterface_DisconnectResourceFilesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::resourceFilesChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_ResourceFilesChanged));
}

void QDesignerFormWindowInterface_ResourceFilesChanged(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->resourceFilesChanged();
}

void QDesignerFormWindowInterface_SelectWidget(void* ptr, void* widget, int sele)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "selectWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)), Q_ARG(bool, sele != 0));
}

void QDesignerFormWindowInterface_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::selectionChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_SelectionChanged));
}

void QDesignerFormWindowInterface_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)()>(&QDesignerFormWindowInterface::selectionChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)()>(&MyQDesignerFormWindowInterface::Signal_SelectionChanged));
}

void QDesignerFormWindowInterface_SelectionChanged(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->selectionChanged();
}

void QDesignerFormWindowInterface_SetAuthor(void* ptr, char* author)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setAuthor(QString(author));
}

void QDesignerFormWindowInterface_SetComment(void* ptr, char* comment)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setComment(QString(comment));
}

int QDesignerFormWindowInterface_SetContents(void* ptr, void* device, char* errorMessage)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->setContents(static_cast<QIODevice*>(device), new QString(errorMessage));
}

int QDesignerFormWindowInterface_SetContents2(void* ptr, char* contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setContents", Q_RETURN_ARG(bool, returnArg), Q_ARG(QString, QString(contents)));
	return returnArg;
}

void QDesignerFormWindowInterface_SetDirty(void* ptr, int dirty)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setDirty", Q_ARG(bool, dirty != 0));
}

void QDesignerFormWindowInterface_SetExportMacro(void* ptr, char* exportMacro)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setExportMacro(QString(exportMacro));
}

void QDesignerFormWindowInterface_SetFeatures(void* ptr, int features)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setFeatures", Q_ARG(QDesignerFormWindowInterface::FeatureFlag, static_cast<QDesignerFormWindowInterface::FeatureFlag>(features)));
}

void QDesignerFormWindowInterface_SetFileName(void* ptr, char* fileName)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setFileName", Q_ARG(QString, QString(fileName)));
}

void QDesignerFormWindowInterface_SetGrid(void* ptr, void* grid)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setGrid", Q_ARG(QPoint, *static_cast<QPoint*>(grid)));
}

void QDesignerFormWindowInterface_SetIncludeHints(void* ptr, char* includeHints)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setIncludeHints(QString(includeHints).split("|", QString::SkipEmptyParts));
}

void QDesignerFormWindowInterface_SetLayoutDefault(void* ptr, int margin, int spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setLayoutDefault(margin, spacing);
}

void QDesignerFormWindowInterface_SetLayoutFunction(void* ptr, char* margin, char* spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setLayoutFunction(QString(margin), QString(spacing));
}

void QDesignerFormWindowInterface_SetMainContainer(void* ptr, void* mainContainer)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setMainContainer(static_cast<QWidget*>(mainContainer));
}

void QDesignerFormWindowInterface_SetPixmapFunction(void* ptr, char* pixmapFunction)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setPixmapFunction(QString(pixmapFunction));
}

void QDesignerFormWindowInterface_SetResourceFileSaveMode(void* ptr, int behaviour)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setResourceFileSaveMode(static_cast<QDesignerFormWindowInterface::ResourceFileSaveMode>(behaviour));
}

void QDesignerFormWindowInterface_UnmanageWidget(void* ptr, void* widget)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "unmanageWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QDesignerFormWindowInterface_ConnectWidgetManaged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetManaged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetManaged));
}

void QDesignerFormWindowInterface_DisconnectWidgetManaged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetManaged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetManaged));
}

void QDesignerFormWindowInterface_WidgetManaged(void* ptr, void* widget)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->widgetManaged(static_cast<QWidget*>(widget));
}

void QDesignerFormWindowInterface_ConnectWidgetRemoved(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetRemoved), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetRemoved));
}

void QDesignerFormWindowInterface_DisconnectWidgetRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetRemoved), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetRemoved));
}

void QDesignerFormWindowInterface_WidgetRemoved(void* ptr, void* widget)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->widgetRemoved(static_cast<QWidget*>(widget));
}

void QDesignerFormWindowInterface_ConnectWidgetUnmanaged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetUnmanaged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetUnmanaged));
}

void QDesignerFormWindowInterface_DisconnectWidgetUnmanaged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QWidget *)>(&QDesignerFormWindowInterface::widgetUnmanaged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QWidget *)>(&MyQDesignerFormWindowInterface::Signal_WidgetUnmanaged));
}

void QDesignerFormWindowInterface_WidgetUnmanaged(void* ptr, void* widget)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->widgetUnmanaged(static_cast<QWidget*>(widget));
}

void QDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->~QDesignerFormWindowInterface();
}

void QDesignerFormWindowInterface_ActionEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerFormWindowInterface_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerFormWindowInterface_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerFormWindowInterface_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerFormWindowInterface_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerFormWindowInterface_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerFormWindowInterface_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerFormWindowInterface_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerFormWindowInterface_DropEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerFormWindowInterface_DropEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerFormWindowInterface_EnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_FocusInEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_HideEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerFormWindowInterface_HideEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerFormWindowInterface_LeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::leaveEvent(static_cast<QEvent*>(event));
}

int QDesignerFormWindowInterface_Metric(void* ptr, int m)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QDesignerFormWindowInterface_MetricDefault(void* ptr, int m)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QDesignerFormWindowInterface_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->minimumSizeHint()).height());
}

void* QDesignerFormWindowInterface_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::minimumSizeHint()).height());
}

void QDesignerFormWindowInterface_MoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerFormWindowInterface_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QDesignerFormWindowInterface_PaintEngine(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->paintEngine();
}

void* QDesignerFormWindowInterface_PaintEngineDefault(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::paintEngine();
}

void QDesignerFormWindowInterface_PaintEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerFormWindowInterface_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerFormWindowInterface_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QDesignerFormWindowInterface_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setEnabled(vbo != 0);
}

void QDesignerFormWindowInterface_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QDesignerFormWindowInterface_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setStyleSheet(QString(styleSheet));
}

void QDesignerFormWindowInterface_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QDesignerFormWindowInterface_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setVisible(visible != 0);
}

void QDesignerFormWindowInterface_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QDesignerFormWindowInterface_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setWindowModified(vbo != 0);
}

void QDesignerFormWindowInterface_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QDesignerFormWindowInterface_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setWindowTitle(QString(vqs));
}

void QDesignerFormWindowInterface_ShowEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerFormWindowInterface_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showEvent(static_cast<QShowEvent*>(event));
}

void* QDesignerFormWindowInterface_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->sizeHint()).height());
}

void* QDesignerFormWindowInterface_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::sizeHint()).height());
}

void QDesignerFormWindowInterface_ChangeEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::changeEvent(static_cast<QEvent*>(event));
}

int QDesignerFormWindowInterface_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDesignerFormWindowInterface_CloseDefault(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::close();
}

void QDesignerFormWindowInterface_CloseEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerFormWindowInterface_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerFormWindowInterface_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerFormWindowInterface_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QDesignerFormWindowInterface_Event(void* ptr, void* event)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDesignerFormWindowInterface_EventDefault(void* ptr, void* event)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::event(static_cast<QEvent*>(event));
}

int QDesignerFormWindowInterface_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->focusNextPrevChild(next != 0);
}

int QDesignerFormWindowInterface_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusNextPrevChild(next != 0);
}

int QDesignerFormWindowInterface_HasHeightForWidth(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->hasHeightForWidth();
}

int QDesignerFormWindowInterface_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hasHeightForWidth();
}

int QDesignerFormWindowInterface_HeightForWidth(void* ptr, int w)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->heightForWidth(w);
}

int QDesignerFormWindowInterface_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::heightForWidth(w);
}

void QDesignerFormWindowInterface_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "hide");
}

void QDesignerFormWindowInterface_HideDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hide();
}

void QDesignerFormWindowInterface_InitPainter(void* ptr, void* painter)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QDesignerFormWindowInterface_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::initPainter(static_cast<QPainter*>(painter));
}

void QDesignerFormWindowInterface_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerFormWindowInterface_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QDesignerFormWindowInterface_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerFormWindowInterface*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QDesignerFormWindowInterface_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QDesignerFormWindowInterface_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "lower");
}

void QDesignerFormWindowInterface_LowerDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::lower();
}

void QDesignerFormWindowInterface_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MousePressEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QDesignerFormWindowInterface_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QDesignerFormWindowInterface_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::nativeEvent(QByteArray(eventType), message, &result);
}

void QDesignerFormWindowInterface_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "raise");
}

void QDesignerFormWindowInterface_RaiseDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::raise();
}

void QDesignerFormWindowInterface_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "repaint");
}

void QDesignerFormWindowInterface_RepaintDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::repaint();
}

void QDesignerFormWindowInterface_ResizeEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerFormWindowInterface_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerFormWindowInterface_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QDesignerFormWindowInterface_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setDisabled(disable != 0);
}

void QDesignerFormWindowInterface_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setFocus");
}

void QDesignerFormWindowInterface_SetFocus2Default(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setFocus();
}

void QDesignerFormWindowInterface_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QDesignerFormWindowInterface_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setHidden(hidden != 0);
}

void QDesignerFormWindowInterface_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "show");
}

void QDesignerFormWindowInterface_ShowDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::show();
}

void QDesignerFormWindowInterface_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "showFullScreen");
}

void QDesignerFormWindowInterface_ShowFullScreenDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showFullScreen();
}

void QDesignerFormWindowInterface_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "showMaximized");
}

void QDesignerFormWindowInterface_ShowMaximizedDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showMaximized();
}

void QDesignerFormWindowInterface_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "showMinimized");
}

void QDesignerFormWindowInterface_ShowMinimizedDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showMinimized();
}

void QDesignerFormWindowInterface_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "showNormal");
}

void QDesignerFormWindowInterface_ShowNormalDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showNormal();
}

void QDesignerFormWindowInterface_TabletEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerFormWindowInterface_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerFormWindowInterface_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "update");
}

void QDesignerFormWindowInterface_UpdateDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::update();
}

void QDesignerFormWindowInterface_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "updateMicroFocus");
}

void QDesignerFormWindowInterface_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::updateMicroFocus();
}

void QDesignerFormWindowInterface_WheelEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerFormWindowInterface_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerFormWindowInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormWindowInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormWindowInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "deleteLater");
}

void QDesignerFormWindowInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::deleteLater();
}

void QDesignerFormWindowInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerFormWindowInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerFormWindowInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerFormWindowInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowInterface*>(ptr)->metaObject());
}

void* QDesignerFormWindowInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::metaObject());
}

class MyQDesignerFormWindowManagerInterface: public QDesignerFormWindowManagerInterface
{
public:
	QAction * action(QDesignerFormWindowManagerInterface::Action action) const { return static_cast<QAction*>(callbackQDesignerFormWindowManagerInterface_Action(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data(), action)); };
	QActionGroup * actionGroup(QDesignerFormWindowManagerInterface::ActionGroup actionGroup) const { return static_cast<QActionGroup*>(callbackQDesignerFormWindowManagerInterface_ActionGroup(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data(), actionGroup)); };
	QDesignerFormWindowInterface * activeFormWindow() const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_ActiveFormWindow(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_ActiveFormWindowChanged(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged(this, this->objectName().toUtf8().data(), formWindow); };
	void addFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_AddFormWindow(this, this->objectName().toUtf8().data(), formWindow); };
	void closeAllPreviews() { callbackQDesignerFormWindowManagerInterface_CloseAllPreviews(this, this->objectName().toUtf8().data()); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerFormWindowManagerInterface_Core(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data())); };
	QDesignerFormWindowInterface * createFormWindow(QWidget * parent, Qt::WindowFlags flags) { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_CreateFormWindow(this, this->objectName().toUtf8().data(), parent, flags)); };
	QPixmap createPreviewPixmap() const { return *static_cast<QPixmap*>(callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data())); };
	QDesignerFormWindowInterface * formWindow(int index) const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_FormWindow(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data(), index)); };
	void Signal_FormWindowAdded(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowAdded(this, this->objectName().toUtf8().data(), formWindow); };
	int formWindowCount() const { return callbackQDesignerFormWindowManagerInterface_FormWindowCount(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data()); };
	void Signal_FormWindowRemoved(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowRemoved(this, this->objectName().toUtf8().data(), formWindow); };
	void Signal_FormWindowSettingsChanged(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged(this, this->objectName().toUtf8().data(), formWindow); };
	void removeFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_RemoveFormWindow(this, this->objectName().toUtf8().data(), formWindow); };
	void setActiveFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow(this, this->objectName().toUtf8().data(), formWindow); };
	void showPluginDialog() { callbackQDesignerFormWindowManagerInterface_ShowPluginDialog(this, this->objectName().toUtf8().data()); };
	void showPreview() { callbackQDesignerFormWindowManagerInterface_ShowPreview(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormWindowManagerInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerFormWindowManagerInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowManagerInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormWindowManagerInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerFormWindowManagerInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowManagerInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQDesignerFormWindowManagerInterface_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormWindowManagerInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormWindowManagerInterface_MetaObject(const_cast<MyQDesignerFormWindowManagerInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerFormWindowManagerInterface_Action(void* ptr, int action)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->action(static_cast<QDesignerFormWindowManagerInterface::Action>(action));
}

void* QDesignerFormWindowManagerInterface_ActionGroup(void* ptr, int actionGroup)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->actionGroup(static_cast<QDesignerFormWindowManagerInterface::ActionGroup>(actionGroup));
}

void* QDesignerFormWindowManagerInterface_ActiveFormWindow(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->activeFormWindow();
}

void QDesignerFormWindowManagerInterface_ConnectActiveFormWindowChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::activeFormWindowChanged), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_ActiveFormWindowChanged));
}

void QDesignerFormWindowManagerInterface_DisconnectActiveFormWindowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::activeFormWindowChanged), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_ActiveFormWindowChanged));
}

void QDesignerFormWindowManagerInterface_ActiveFormWindowChanged(void* ptr, void* formWindow)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->activeFormWindowChanged(static_cast<QDesignerFormWindowInterface*>(formWindow));
}

void QDesignerFormWindowManagerInterface_AddFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "addFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerFormWindowManagerInterface_CloseAllPreviews(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "closeAllPreviews");
}

void* QDesignerFormWindowManagerInterface_Core(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->core();
}

void* QDesignerFormWindowManagerInterface_CreateFormWindow(void* ptr, void* parent, int flags)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->createFormWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDesignerFormWindowManagerInterface_CreatePreviewPixmap(void* ptr)
{
	return new QPixmap(static_cast<QDesignerFormWindowManagerInterface*>(ptr)->createPreviewPixmap());
}

void* QDesignerFormWindowManagerInterface_FormWindow(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindow(index);
}

void QDesignerFormWindowManagerInterface_ConnectFormWindowAdded(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowAdded), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowAdded));
}

void QDesignerFormWindowManagerInterface_DisconnectFormWindowAdded(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowAdded), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowAdded));
}

void QDesignerFormWindowManagerInterface_FormWindowAdded(void* ptr, void* formWindow)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindowAdded(static_cast<QDesignerFormWindowInterface*>(formWindow));
}

int QDesignerFormWindowManagerInterface_FormWindowCount(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindowCount();
}

void QDesignerFormWindowManagerInterface_ConnectFormWindowRemoved(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowRemoved), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowRemoved));
}

void QDesignerFormWindowManagerInterface_DisconnectFormWindowRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowRemoved), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowRemoved));
}

void QDesignerFormWindowManagerInterface_FormWindowRemoved(void* ptr, void* formWindow)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindowRemoved(static_cast<QDesignerFormWindowInterface*>(formWindow));
}

void QDesignerFormWindowManagerInterface_ConnectFormWindowSettingsChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowSettingsChanged), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowSettingsChanged));
}

void QDesignerFormWindowManagerInterface_DisconnectFormWindowSettingsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowManagerInterface*>(ptr), static_cast<void (QDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&QDesignerFormWindowManagerInterface::formWindowSettingsChanged), static_cast<MyQDesignerFormWindowManagerInterface*>(ptr), static_cast<void (MyQDesignerFormWindowManagerInterface::*)(QDesignerFormWindowInterface *)>(&MyQDesignerFormWindowManagerInterface::Signal_FormWindowSettingsChanged));
}

void QDesignerFormWindowManagerInterface_FormWindowSettingsChanged(void* ptr, void* formWindow)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindowSettingsChanged(static_cast<QDesignerFormWindowInterface*>(formWindow));
}

void QDesignerFormWindowManagerInterface_RemoveFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "removeFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerFormWindowManagerInterface_SetActiveFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "setActiveFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerFormWindowManagerInterface_ShowPluginDialog(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "showPluginDialog");
}

void QDesignerFormWindowManagerInterface_ShowPreview(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "showPreview");
}

void QDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface(void* ptr)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->~QDesignerFormWindowManagerInterface();
}

void QDesignerFormWindowManagerInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormWindowManagerInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerFormWindowManagerInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowManagerInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowManagerInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowManagerInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowManagerInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowManagerInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowManagerInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowManagerInterface*>(ptr), "deleteLater");
}

void QDesignerFormWindowManagerInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::deleteLater();
}

void QDesignerFormWindowManagerInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowManagerInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerFormWindowManagerInterface_Event(void* ptr, void* e)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDesignerFormWindowManagerInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::event(static_cast<QEvent*>(e));
}

int QDesignerFormWindowManagerInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerFormWindowManagerInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerFormWindowManagerInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowManagerInterface*>(ptr)->metaObject());
}

void* QDesignerFormWindowManagerInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::metaObject());
}

class MyQDesignerMemberSheetExtension: public QDesignerMemberSheetExtension
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	int count() const { return callbackQDesignerMemberSheetExtension_Count(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data()); };
	QString declaredInClass(int index) const { return QString(callbackQDesignerMemberSheetExtension_DeclaredInClass(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	int indexOf(const QString & name) const { return callbackQDesignerMemberSheetExtension_IndexOf(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), name.toUtf8().data()); };
	bool inheritedFromWidget(int index) const { return callbackQDesignerMemberSheetExtension_InheritedFromWidget(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isSignal(int index) const { return callbackQDesignerMemberSheetExtension_IsSignal(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isSlot(int index) const { return callbackQDesignerMemberSheetExtension_IsSlot(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isVisible(int index) const { return callbackQDesignerMemberSheetExtension_IsVisible(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	QString memberGroup(int index) const { return QString(callbackQDesignerMemberSheetExtension_MemberGroup(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	QString memberName(int index) const { return QString(callbackQDesignerMemberSheetExtension_MemberName(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	void setMemberGroup(int index, const QString & group) { callbackQDesignerMemberSheetExtension_SetMemberGroup(this, this->objectNameAbs().toUtf8().data(), index, group.toUtf8().data()); };
	void setVisible(int index, bool visible) { callbackQDesignerMemberSheetExtension_SetVisible(this, this->objectNameAbs().toUtf8().data(), index, visible); };
	QString signature(int index) const { return QString(callbackQDesignerMemberSheetExtension_Signature(const_cast<MyQDesignerMemberSheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
};

int QDesignerMemberSheetExtension_Count(void* ptr)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->count();
}

char* QDesignerMemberSheetExtension_DeclaredInClass(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->declaredInClass(index).toUtf8().data();
}

int QDesignerMemberSheetExtension_IndexOf(void* ptr, char* name)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->indexOf(QString(name));
}

int QDesignerMemberSheetExtension_InheritedFromWidget(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->inheritedFromWidget(index);
}

int QDesignerMemberSheetExtension_IsSignal(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isSignal(index);
}

int QDesignerMemberSheetExtension_IsSlot(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isSlot(index);
}

int QDesignerMemberSheetExtension_IsVisible(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isVisible(index);
}

char* QDesignerMemberSheetExtension_MemberGroup(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->memberGroup(index).toUtf8().data();
}

char* QDesignerMemberSheetExtension_MemberName(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->memberName(index).toUtf8().data();
}

void QDesignerMemberSheetExtension_SetMemberGroup(void* ptr, int index, char* group)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->setMemberGroup(index, QString(group));
}

void QDesignerMemberSheetExtension_SetVisible(void* ptr, int index, int visible)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->setVisible(index, visible != 0);
}

char* QDesignerMemberSheetExtension_Signature(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->signature(index).toUtf8().data();
}

void QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(void* ptr)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->~QDesignerMemberSheetExtension();
}

char* QDesignerMemberSheetExtension_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerMemberSheetExtension*>(static_cast<QDesignerMemberSheetExtension*>(ptr))) {
		return static_cast<MyQDesignerMemberSheetExtension*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerMemberSheetExtension_BASE").toUtf8().data();
}

void QDesignerMemberSheetExtension_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerMemberSheetExtension*>(static_cast<QDesignerMemberSheetExtension*>(ptr))) {
		static_cast<MyQDesignerMemberSheetExtension*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerObjectInspectorInterface: public QDesignerObjectInspectorInterface
{
public:
	MyQDesignerObjectInspectorInterface(QWidget *parent, Qt::WindowFlags flags) : QDesignerObjectInspectorInterface(parent, flags) {};
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerObjectInspectorInterface_Core(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data())); };
	void setFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerObjectInspectorInterface_SetFormWindow(this, this->objectName().toUtf8().data(), formWindow); };
	void actionEvent(QActionEvent * event) { callbackQDesignerObjectInspectorInterface_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerObjectInspectorInterface_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerObjectInspectorInterface_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerObjectInspectorInterface_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerObjectInspectorInterface_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerObjectInspectorInterface_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerObjectInspectorInterface_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQDesignerObjectInspectorInterface_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerObjectInspectorInterface_Metric(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerObjectInspectorInterface_MinimumSizeHint(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerObjectInspectorInterface_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerObjectInspectorInterface_PaintEngine(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerObjectInspectorInterface_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQDesignerObjectInspectorInterface_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQDesignerObjectInspectorInterface_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQDesignerObjectInspectorInterface_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQDesignerObjectInspectorInterface_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQDesignerObjectInspectorInterface_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQDesignerObjectInspectorInterface_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerObjectInspectorInterface_SizeHint(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQDesignerObjectInspectorInterface_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQDesignerObjectInspectorInterface_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerObjectInspectorInterface_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * event) { return callbackQDesignerObjectInspectorInterface_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerObjectInspectorInterface_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQDesignerObjectInspectorInterface_HasHeightForWidth(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerObjectInspectorInterface_HeightForWidth(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQDesignerObjectInspectorInterface_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQDesignerObjectInspectorInterface_InitPainter(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerObjectInspectorInterface_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerObjectInspectorInterface_InputMethodQuery(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerObjectInspectorInterface_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerObjectInspectorInterface_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQDesignerObjectInspectorInterface_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerObjectInspectorInterface_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQDesignerObjectInspectorInterface_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQDesignerObjectInspectorInterface_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerObjectInspectorInterface_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQDesignerObjectInspectorInterface_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQDesignerObjectInspectorInterface_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQDesignerObjectInspectorInterface_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQDesignerObjectInspectorInterface_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQDesignerObjectInspectorInterface_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQDesignerObjectInspectorInterface_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQDesignerObjectInspectorInterface_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQDesignerObjectInspectorInterface_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerObjectInspectorInterface_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQDesignerObjectInspectorInterface_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQDesignerObjectInspectorInterface_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerObjectInspectorInterface_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerObjectInspectorInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerObjectInspectorInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerObjectInspectorInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerObjectInspectorInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerObjectInspectorInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerObjectInspectorInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerObjectInspectorInterface_MetaObject(const_cast<MyQDesignerObjectInspectorInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerObjectInspectorInterface_NewQDesignerObjectInspectorInterface(void* parent, int flags)
{
	return new MyQDesignerObjectInspectorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDesignerObjectInspectorInterface_Core(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->core();
}

void* QDesignerObjectInspectorInterface_CoreDefault(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::core();
}

void QDesignerObjectInspectorInterface_SetFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->~QDesignerObjectInspectorInterface();
}

void QDesignerObjectInspectorInterface_ActionEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerObjectInspectorInterface_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DropEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerObjectInspectorInterface_DropEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerObjectInspectorInterface_EnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusInEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_HideEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerObjectInspectorInterface_HideEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerObjectInspectorInterface_LeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::leaveEvent(static_cast<QEvent*>(event));
}

int QDesignerObjectInspectorInterface_Metric(void* ptr, int m)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QDesignerObjectInspectorInterface_MetricDefault(void* ptr, int m)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QDesignerObjectInspectorInterface_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->minimumSizeHint()).height());
}

void* QDesignerObjectInspectorInterface_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::minimumSizeHint()).height());
}

void QDesignerObjectInspectorInterface_MoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerObjectInspectorInterface_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QDesignerObjectInspectorInterface_PaintEngine(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->paintEngine();
}

void* QDesignerObjectInspectorInterface_PaintEngineDefault(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::paintEngine();
}

void QDesignerObjectInspectorInterface_PaintEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerObjectInspectorInterface_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerObjectInspectorInterface_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QDesignerObjectInspectorInterface_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setEnabled(vbo != 0);
}

void QDesignerObjectInspectorInterface_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QDesignerObjectInspectorInterface_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setStyleSheet(QString(styleSheet));
}

void QDesignerObjectInspectorInterface_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QDesignerObjectInspectorInterface_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setVisible(visible != 0);
}

void QDesignerObjectInspectorInterface_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QDesignerObjectInspectorInterface_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setWindowModified(vbo != 0);
}

void QDesignerObjectInspectorInterface_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QDesignerObjectInspectorInterface_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setWindowTitle(QString(vqs));
}

void QDesignerObjectInspectorInterface_ShowEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerObjectInspectorInterface_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void* QDesignerObjectInspectorInterface_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->sizeHint()).height());
}

void* QDesignerObjectInspectorInterface_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::sizeHint()).height());
}

void QDesignerObjectInspectorInterface_ChangeEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::changeEvent(static_cast<QEvent*>(event));
}

int QDesignerObjectInspectorInterface_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDesignerObjectInspectorInterface_CloseDefault(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::close();
}

void QDesignerObjectInspectorInterface_CloseEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerObjectInspectorInterface_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerObjectInspectorInterface_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerObjectInspectorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QDesignerObjectInspectorInterface_Event(void* ptr, void* event)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDesignerObjectInspectorInterface_EventDefault(void* ptr, void* event)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::event(static_cast<QEvent*>(event));
}

int QDesignerObjectInspectorInterface_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->focusNextPrevChild(next != 0);
}

int QDesignerObjectInspectorInterface_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusNextPrevChild(next != 0);
}

int QDesignerObjectInspectorInterface_HasHeightForWidth(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->hasHeightForWidth();
}

int QDesignerObjectInspectorInterface_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hasHeightForWidth();
}

int QDesignerObjectInspectorInterface_HeightForWidth(void* ptr, int w)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->heightForWidth(w);
}

int QDesignerObjectInspectorInterface_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::heightForWidth(w);
}

void QDesignerObjectInspectorInterface_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "hide");
}

void QDesignerObjectInspectorInterface_HideDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hide();
}

void QDesignerObjectInspectorInterface_InitPainter(void* ptr, void* painter)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QDesignerObjectInspectorInterface_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::initPainter(static_cast<QPainter*>(painter));
}

void QDesignerObjectInspectorInterface_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerObjectInspectorInterface_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QDesignerObjectInspectorInterface_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerObjectInspectorInterface*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QDesignerObjectInspectorInterface_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QDesignerObjectInspectorInterface_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "lower");
}

void QDesignerObjectInspectorInterface_LowerDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::lower();
}

void QDesignerObjectInspectorInterface_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MousePressEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QDesignerObjectInspectorInterface_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QDesignerObjectInspectorInterface_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::nativeEvent(QByteArray(eventType), message, &result);
}

void QDesignerObjectInspectorInterface_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "raise");
}

void QDesignerObjectInspectorInterface_RaiseDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::raise();
}

void QDesignerObjectInspectorInterface_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "repaint");
}

void QDesignerObjectInspectorInterface_RepaintDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::repaint();
}

void QDesignerObjectInspectorInterface_ResizeEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerObjectInspectorInterface_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerObjectInspectorInterface_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QDesignerObjectInspectorInterface_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setDisabled(disable != 0);
}

void QDesignerObjectInspectorInterface_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setFocus");
}

void QDesignerObjectInspectorInterface_SetFocus2Default(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setFocus();
}

void QDesignerObjectInspectorInterface_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QDesignerObjectInspectorInterface_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setHidden(hidden != 0);
}

void QDesignerObjectInspectorInterface_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "show");
}

void QDesignerObjectInspectorInterface_ShowDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::show();
}

void QDesignerObjectInspectorInterface_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "showFullScreen");
}

void QDesignerObjectInspectorInterface_ShowFullScreenDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showFullScreen();
}

void QDesignerObjectInspectorInterface_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "showMaximized");
}

void QDesignerObjectInspectorInterface_ShowMaximizedDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showMaximized();
}

void QDesignerObjectInspectorInterface_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "showMinimized");
}

void QDesignerObjectInspectorInterface_ShowMinimizedDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showMinimized();
}

void QDesignerObjectInspectorInterface_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "showNormal");
}

void QDesignerObjectInspectorInterface_ShowNormalDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showNormal();
}

void QDesignerObjectInspectorInterface_TabletEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerObjectInspectorInterface_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerObjectInspectorInterface_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "update");
}

void QDesignerObjectInspectorInterface_UpdateDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::update();
}

void QDesignerObjectInspectorInterface_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "updateMicroFocus");
}

void QDesignerObjectInspectorInterface_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::updateMicroFocus();
}

void QDesignerObjectInspectorInterface_WheelEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerObjectInspectorInterface_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerObjectInspectorInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerObjectInspectorInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerObjectInspectorInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerObjectInspectorInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerObjectInspectorInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerObjectInspectorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerObjectInspectorInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "deleteLater");
}

void QDesignerObjectInspectorInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::deleteLater();
}

void QDesignerObjectInspectorInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerObjectInspectorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerObjectInspectorInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerObjectInspectorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerObjectInspectorInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->metaObject());
}

void* QDesignerObjectInspectorInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::metaObject());
}

class MyQDesignerPropertyEditorInterface: public QDesignerPropertyEditorInterface
{
public:
	MyQDesignerPropertyEditorInterface(QWidget *parent, Qt::WindowFlags flags) : QDesignerPropertyEditorInterface(parent, flags) {};
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerPropertyEditorInterface_Core(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	QString currentPropertyName() const { return QString(callbackQDesignerPropertyEditorInterface_CurrentPropertyName(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	bool isReadOnly() const { return callbackQDesignerPropertyEditorInterface_IsReadOnly(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	QObject * object() const { return static_cast<QObject*>(callbackQDesignerPropertyEditorInterface_Object(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void Signal_PropertyChanged(const QString & name, const QVariant & value) { callbackQDesignerPropertyEditorInterface_PropertyChanged(this, this->objectName().toUtf8().data(), name.toUtf8().data(), new QVariant(value)); };
	void setObject(QObject * object) { callbackQDesignerPropertyEditorInterface_SetObject(this, this->objectName().toUtf8().data(), object); };
	void setPropertyValue(const QString & name, const QVariant & value, bool changed) { callbackQDesignerPropertyEditorInterface_SetPropertyValue(this, this->objectName().toUtf8().data(), name.toUtf8().data(), new QVariant(value), changed); };
	void setReadOnly(bool readOnly) { callbackQDesignerPropertyEditorInterface_SetReadOnly(this, this->objectName().toUtf8().data(), readOnly); };
	void actionEvent(QActionEvent * event) { callbackQDesignerPropertyEditorInterface_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerPropertyEditorInterface_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerPropertyEditorInterface_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerPropertyEditorInterface_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerPropertyEditorInterface_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerPropertyEditorInterface_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerPropertyEditorInterface_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQDesignerPropertyEditorInterface_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerPropertyEditorInterface_Metric(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerPropertyEditorInterface_MinimumSizeHint(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerPropertyEditorInterface_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerPropertyEditorInterface_PaintEngine(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerPropertyEditorInterface_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQDesignerPropertyEditorInterface_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQDesignerPropertyEditorInterface_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQDesignerPropertyEditorInterface_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQDesignerPropertyEditorInterface_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQDesignerPropertyEditorInterface_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQDesignerPropertyEditorInterface_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerPropertyEditorInterface_SizeHint(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQDesignerPropertyEditorInterface_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQDesignerPropertyEditorInterface_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerPropertyEditorInterface_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * event) { return callbackQDesignerPropertyEditorInterface_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerPropertyEditorInterface_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQDesignerPropertyEditorInterface_HasHeightForWidth(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerPropertyEditorInterface_HeightForWidth(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQDesignerPropertyEditorInterface_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQDesignerPropertyEditorInterface_InitPainter(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerPropertyEditorInterface_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerPropertyEditorInterface_InputMethodQuery(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerPropertyEditorInterface_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerPropertyEditorInterface_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQDesignerPropertyEditorInterface_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerPropertyEditorInterface_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQDesignerPropertyEditorInterface_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQDesignerPropertyEditorInterface_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerPropertyEditorInterface_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQDesignerPropertyEditorInterface_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQDesignerPropertyEditorInterface_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQDesignerPropertyEditorInterface_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQDesignerPropertyEditorInterface_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQDesignerPropertyEditorInterface_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQDesignerPropertyEditorInterface_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQDesignerPropertyEditorInterface_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQDesignerPropertyEditorInterface_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerPropertyEditorInterface_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQDesignerPropertyEditorInterface_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQDesignerPropertyEditorInterface_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerPropertyEditorInterface_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerPropertyEditorInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerPropertyEditorInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerPropertyEditorInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerPropertyEditorInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerPropertyEditorInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerPropertyEditorInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerPropertyEditorInterface_MetaObject(const_cast<MyQDesignerPropertyEditorInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QDesignerPropertyEditorInterface_NewQDesignerPropertyEditorInterface(void* parent, int flags)
{
	return new MyQDesignerPropertyEditorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDesignerPropertyEditorInterface_Core(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->core();
}

void* QDesignerPropertyEditorInterface_CoreDefault(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::core();
}

char* QDesignerPropertyEditorInterface_CurrentPropertyName(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->currentPropertyName().toUtf8().data();
}

int QDesignerPropertyEditorInterface_IsReadOnly(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->isReadOnly();
}

void* QDesignerPropertyEditorInterface_Object(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->object();
}

void QDesignerPropertyEditorInterface_ConnectPropertyChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerPropertyEditorInterface*>(ptr), static_cast<void (QDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&QDesignerPropertyEditorInterface::propertyChanged), static_cast<MyQDesignerPropertyEditorInterface*>(ptr), static_cast<void (MyQDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&MyQDesignerPropertyEditorInterface::Signal_PropertyChanged));
}

void QDesignerPropertyEditorInterface_DisconnectPropertyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerPropertyEditorInterface*>(ptr), static_cast<void (QDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&QDesignerPropertyEditorInterface::propertyChanged), static_cast<MyQDesignerPropertyEditorInterface*>(ptr), static_cast<void (MyQDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&MyQDesignerPropertyEditorInterface::Signal_PropertyChanged));
}

void QDesignerPropertyEditorInterface_PropertyChanged(void* ptr, char* name, void* value)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->propertyChanged(QString(name), *static_cast<QVariant*>(value));
}

void QDesignerPropertyEditorInterface_SetObject(void* ptr, void* object)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setObject", Q_ARG(QObject*, static_cast<QObject*>(object)));
}

void QDesignerPropertyEditorInterface_SetPropertyValue(void* ptr, char* name, void* value, int changed)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setPropertyValue", Q_ARG(QString, QString(name)), Q_ARG(QVariant, *static_cast<QVariant*>(value)), Q_ARG(bool, changed != 0));
}

void QDesignerPropertyEditorInterface_SetReadOnly(void* ptr, int readOnly)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setReadOnly", Q_ARG(bool, readOnly != 0));
}

void QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->~QDesignerPropertyEditorInterface();
}

void QDesignerPropertyEditorInterface_ActionEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerPropertyEditorInterface_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DropEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerPropertyEditorInterface_DropEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerPropertyEditorInterface_EnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusInEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_HideEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerPropertyEditorInterface_HideEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerPropertyEditorInterface_LeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::leaveEvent(static_cast<QEvent*>(event));
}

int QDesignerPropertyEditorInterface_Metric(void* ptr, int m)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QDesignerPropertyEditorInterface_MetricDefault(void* ptr, int m)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QDesignerPropertyEditorInterface_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->minimumSizeHint()).height());
}

void* QDesignerPropertyEditorInterface_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::minimumSizeHint()).height());
}

void QDesignerPropertyEditorInterface_MoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerPropertyEditorInterface_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QDesignerPropertyEditorInterface_PaintEngine(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->paintEngine();
}

void* QDesignerPropertyEditorInterface_PaintEngineDefault(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::paintEngine();
}

void QDesignerPropertyEditorInterface_PaintEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerPropertyEditorInterface_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerPropertyEditorInterface_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QDesignerPropertyEditorInterface_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setEnabled(vbo != 0);
}

void QDesignerPropertyEditorInterface_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QDesignerPropertyEditorInterface_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setStyleSheet(QString(styleSheet));
}

void QDesignerPropertyEditorInterface_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QDesignerPropertyEditorInterface_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setVisible(visible != 0);
}

void QDesignerPropertyEditorInterface_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QDesignerPropertyEditorInterface_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setWindowModified(vbo != 0);
}

void QDesignerPropertyEditorInterface_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QDesignerPropertyEditorInterface_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setWindowTitle(QString(vqs));
}

void QDesignerPropertyEditorInterface_ShowEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerPropertyEditorInterface_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void* QDesignerPropertyEditorInterface_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->sizeHint()).height());
}

void* QDesignerPropertyEditorInterface_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::sizeHint()).height());
}

void QDesignerPropertyEditorInterface_ChangeEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::changeEvent(static_cast<QEvent*>(event));
}

int QDesignerPropertyEditorInterface_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDesignerPropertyEditorInterface_CloseDefault(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::close();
}

void QDesignerPropertyEditorInterface_CloseEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerPropertyEditorInterface_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerPropertyEditorInterface_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerPropertyEditorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QDesignerPropertyEditorInterface_Event(void* ptr, void* event)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDesignerPropertyEditorInterface_EventDefault(void* ptr, void* event)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::event(static_cast<QEvent*>(event));
}

int QDesignerPropertyEditorInterface_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->focusNextPrevChild(next != 0);
}

int QDesignerPropertyEditorInterface_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusNextPrevChild(next != 0);
}

int QDesignerPropertyEditorInterface_HasHeightForWidth(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->hasHeightForWidth();
}

int QDesignerPropertyEditorInterface_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hasHeightForWidth();
}

int QDesignerPropertyEditorInterface_HeightForWidth(void* ptr, int w)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->heightForWidth(w);
}

int QDesignerPropertyEditorInterface_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::heightForWidth(w);
}

void QDesignerPropertyEditorInterface_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "hide");
}

void QDesignerPropertyEditorInterface_HideDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hide();
}

void QDesignerPropertyEditorInterface_InitPainter(void* ptr, void* painter)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QDesignerPropertyEditorInterface_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::initPainter(static_cast<QPainter*>(painter));
}

void QDesignerPropertyEditorInterface_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerPropertyEditorInterface_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QDesignerPropertyEditorInterface_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerPropertyEditorInterface*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QDesignerPropertyEditorInterface_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QDesignerPropertyEditorInterface_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "lower");
}

void QDesignerPropertyEditorInterface_LowerDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::lower();
}

void QDesignerPropertyEditorInterface_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MousePressEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QDesignerPropertyEditorInterface_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QDesignerPropertyEditorInterface_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::nativeEvent(QByteArray(eventType), message, &result);
}

void QDesignerPropertyEditorInterface_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "raise");
}

void QDesignerPropertyEditorInterface_RaiseDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::raise();
}

void QDesignerPropertyEditorInterface_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "repaint");
}

void QDesignerPropertyEditorInterface_RepaintDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::repaint();
}

void QDesignerPropertyEditorInterface_ResizeEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerPropertyEditorInterface_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerPropertyEditorInterface_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QDesignerPropertyEditorInterface_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setDisabled(disable != 0);
}

void QDesignerPropertyEditorInterface_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setFocus");
}

void QDesignerPropertyEditorInterface_SetFocus2Default(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setFocus();
}

void QDesignerPropertyEditorInterface_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QDesignerPropertyEditorInterface_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setHidden(hidden != 0);
}

void QDesignerPropertyEditorInterface_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "show");
}

void QDesignerPropertyEditorInterface_ShowDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::show();
}

void QDesignerPropertyEditorInterface_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "showFullScreen");
}

void QDesignerPropertyEditorInterface_ShowFullScreenDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showFullScreen();
}

void QDesignerPropertyEditorInterface_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "showMaximized");
}

void QDesignerPropertyEditorInterface_ShowMaximizedDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showMaximized();
}

void QDesignerPropertyEditorInterface_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "showMinimized");
}

void QDesignerPropertyEditorInterface_ShowMinimizedDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showMinimized();
}

void QDesignerPropertyEditorInterface_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "showNormal");
}

void QDesignerPropertyEditorInterface_ShowNormalDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showNormal();
}

void QDesignerPropertyEditorInterface_TabletEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerPropertyEditorInterface_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerPropertyEditorInterface_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "update");
}

void QDesignerPropertyEditorInterface_UpdateDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::update();
}

void QDesignerPropertyEditorInterface_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "updateMicroFocus");
}

void QDesignerPropertyEditorInterface_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::updateMicroFocus();
}

void QDesignerPropertyEditorInterface_WheelEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerPropertyEditorInterface_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerPropertyEditorInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerPropertyEditorInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerPropertyEditorInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerPropertyEditorInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerPropertyEditorInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerPropertyEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerPropertyEditorInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "deleteLater");
}

void QDesignerPropertyEditorInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::deleteLater();
}

void QDesignerPropertyEditorInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerPropertyEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerPropertyEditorInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerPropertyEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerPropertyEditorInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->metaObject());
}

void* QDesignerPropertyEditorInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::metaObject());
}

class MyQDesignerPropertySheetExtension: public QDesignerPropertySheetExtension
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	int count() const { return callbackQDesignerPropertySheetExtension_Count(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data()); };
	bool hasReset(int index) const { return callbackQDesignerPropertySheetExtension_HasReset(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	int indexOf(const QString & name) const { return callbackQDesignerPropertySheetExtension_IndexOf(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), name.toUtf8().data()); };
	bool isAttribute(int index) const { return callbackQDesignerPropertySheetExtension_IsAttribute(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isChanged(int index) const { return callbackQDesignerPropertySheetExtension_IsChanged(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isEnabled(int index) const { return callbackQDesignerPropertySheetExtension_IsEnabled(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	bool isVisible(int index) const { return callbackQDesignerPropertySheetExtension_IsVisible(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index) != 0; };
	QVariant property(int index) const { return *static_cast<QVariant*>(callbackQDesignerPropertySheetExtension_Property(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	QString propertyGroup(int index) const { return QString(callbackQDesignerPropertySheetExtension_PropertyGroup(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	QString propertyName(int index) const { return QString(callbackQDesignerPropertySheetExtension_PropertyName(const_cast<MyQDesignerPropertySheetExtension*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	bool reset(int index) { return callbackQDesignerPropertySheetExtension_Reset(this, this->objectNameAbs().toUtf8().data(), index) != 0; };
	void setAttribute(int index, bool attribute) { callbackQDesignerPropertySheetExtension_SetAttribute(this, this->objectNameAbs().toUtf8().data(), index, attribute); };
	void setChanged(int index, bool changed) { callbackQDesignerPropertySheetExtension_SetChanged(this, this->objectNameAbs().toUtf8().data(), index, changed); };
	void setProperty(int index, const QVariant & value) { callbackQDesignerPropertySheetExtension_SetProperty(this, this->objectNameAbs().toUtf8().data(), index, new QVariant(value)); };
	void setPropertyGroup(int index, const QString & group) { callbackQDesignerPropertySheetExtension_SetPropertyGroup(this, this->objectNameAbs().toUtf8().data(), index, group.toUtf8().data()); };
	void setVisible(int index, bool visible) { callbackQDesignerPropertySheetExtension_SetVisible(this, this->objectNameAbs().toUtf8().data(), index, visible); };
};

int QDesignerPropertySheetExtension_Count(void* ptr)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->count();
}

int QDesignerPropertySheetExtension_HasReset(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->hasReset(index);
}

int QDesignerPropertySheetExtension_IndexOf(void* ptr, char* name)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->indexOf(QString(name));
}

int QDesignerPropertySheetExtension_IsAttribute(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isAttribute(index);
}

int QDesignerPropertySheetExtension_IsChanged(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isChanged(index);
}

int QDesignerPropertySheetExtension_IsEnabled(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isEnabled(index);
}

int QDesignerPropertySheetExtension_IsEnabledDefault(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->QDesignerPropertySheetExtension::isEnabled(index);
}

int QDesignerPropertySheetExtension_IsVisible(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isVisible(index);
}

void* QDesignerPropertySheetExtension_Property(void* ptr, int index)
{
	return new QVariant(static_cast<QDesignerPropertySheetExtension*>(ptr)->property(index));
}

char* QDesignerPropertySheetExtension_PropertyGroup(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->propertyGroup(index).toUtf8().data();
}

char* QDesignerPropertySheetExtension_PropertyName(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->propertyName(index).toUtf8().data();
}

int QDesignerPropertySheetExtension_Reset(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->reset(index);
}

void QDesignerPropertySheetExtension_SetAttribute(void* ptr, int index, int attribute)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setAttribute(index, attribute != 0);
}

void QDesignerPropertySheetExtension_SetChanged(void* ptr, int index, int changed)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setChanged(index, changed != 0);
}

void QDesignerPropertySheetExtension_SetProperty(void* ptr, int index, void* value)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setProperty(index, *static_cast<QVariant*>(value));
}

void QDesignerPropertySheetExtension_SetPropertyGroup(void* ptr, int index, char* group)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setPropertyGroup(index, QString(group));
}

void QDesignerPropertySheetExtension_SetVisible(void* ptr, int index, int visible)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setVisible(index, visible != 0);
}

void QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(void* ptr)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->~QDesignerPropertySheetExtension();
}

char* QDesignerPropertySheetExtension_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerPropertySheetExtension*>(static_cast<QDesignerPropertySheetExtension*>(ptr))) {
		return static_cast<MyQDesignerPropertySheetExtension*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerPropertySheetExtension_BASE").toUtf8().data();
}

void QDesignerPropertySheetExtension_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerPropertySheetExtension*>(static_cast<QDesignerPropertySheetExtension*>(ptr))) {
		static_cast<MyQDesignerPropertySheetExtension*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerTaskMenuExtension: public QDesignerTaskMenuExtension
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QAction * preferredEditAction() const { return static_cast<QAction*>(callbackQDesignerTaskMenuExtension_PreferredEditAction(const_cast<MyQDesignerTaskMenuExtension*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QDesignerTaskMenuExtension_PreferredEditAction(void* ptr)
{
	return static_cast<QDesignerTaskMenuExtension*>(ptr)->preferredEditAction();
}

void* QDesignerTaskMenuExtension_PreferredEditActionDefault(void* ptr)
{
	return static_cast<QDesignerTaskMenuExtension*>(ptr)->QDesignerTaskMenuExtension::preferredEditAction();
}

void QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(void* ptr)
{
	static_cast<QDesignerTaskMenuExtension*>(ptr)->~QDesignerTaskMenuExtension();
}

char* QDesignerTaskMenuExtension_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDesignerTaskMenuExtension*>(static_cast<QDesignerTaskMenuExtension*>(ptr))) {
		return static_cast<MyQDesignerTaskMenuExtension*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDesignerTaskMenuExtension_BASE").toUtf8().data();
}

void QDesignerTaskMenuExtension_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDesignerTaskMenuExtension*>(static_cast<QDesignerTaskMenuExtension*>(ptr))) {
		static_cast<MyQDesignerTaskMenuExtension*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQDesignerWidgetBoxInterface: public QDesignerWidgetBoxInterface
{
public:
	QString fileName() const { return QString(callbackQDesignerWidgetBoxInterface_FileName(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data())); };
	bool load() { return callbackQDesignerWidgetBoxInterface_Load(this, this->objectName().toUtf8().data()) != 0; };
	bool save() { return callbackQDesignerWidgetBoxInterface_Save(this, this->objectName().toUtf8().data()) != 0; };
	void setFileName(const QString & fileName) { callbackQDesignerWidgetBoxInterface_SetFileName(this, this->objectName().toUtf8().data(), fileName.toUtf8().data()); };
	void actionEvent(QActionEvent * event) { callbackQDesignerWidgetBoxInterface_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerWidgetBoxInterface_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerWidgetBoxInterface_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerWidgetBoxInterface_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerWidgetBoxInterface_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerWidgetBoxInterface_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerWidgetBoxInterface_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQDesignerWidgetBoxInterface_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerWidgetBoxInterface_Metric(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerWidgetBoxInterface_MinimumSizeHint(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerWidgetBoxInterface_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerWidgetBoxInterface_PaintEngine(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerWidgetBoxInterface_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQDesignerWidgetBoxInterface_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQDesignerWidgetBoxInterface_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQDesignerWidgetBoxInterface_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQDesignerWidgetBoxInterface_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQDesignerWidgetBoxInterface_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQDesignerWidgetBoxInterface_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerWidgetBoxInterface_SizeHint(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQDesignerWidgetBoxInterface_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQDesignerWidgetBoxInterface_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerWidgetBoxInterface_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * event) { return callbackQDesignerWidgetBoxInterface_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerWidgetBoxInterface_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQDesignerWidgetBoxInterface_HasHeightForWidth(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerWidgetBoxInterface_HeightForWidth(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQDesignerWidgetBoxInterface_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQDesignerWidgetBoxInterface_InitPainter(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerWidgetBoxInterface_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerWidgetBoxInterface_InputMethodQuery(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerWidgetBoxInterface_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerWidgetBoxInterface_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQDesignerWidgetBoxInterface_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerWidgetBoxInterface_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQDesignerWidgetBoxInterface_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQDesignerWidgetBoxInterface_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerWidgetBoxInterface_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQDesignerWidgetBoxInterface_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQDesignerWidgetBoxInterface_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQDesignerWidgetBoxInterface_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQDesignerWidgetBoxInterface_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQDesignerWidgetBoxInterface_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQDesignerWidgetBoxInterface_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQDesignerWidgetBoxInterface_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQDesignerWidgetBoxInterface_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerWidgetBoxInterface_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQDesignerWidgetBoxInterface_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQDesignerWidgetBoxInterface_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerWidgetBoxInterface_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerWidgetBoxInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDesignerWidgetBoxInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerWidgetBoxInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDesignerWidgetBoxInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerWidgetBoxInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerWidgetBoxInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerWidgetBoxInterface_MetaObject(const_cast<MyQDesignerWidgetBoxInterface*>(this), this->objectName().toUtf8().data())); };
};

char* QDesignerWidgetBoxInterface_FileName(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->fileName().toUtf8().data();
}

int QDesignerWidgetBoxInterface_Load(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->load();
}

int QDesignerWidgetBoxInterface_Save(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->save();
}

void QDesignerWidgetBoxInterface_SetFileName(void* ptr, char* fileName)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->setFileName(QString(fileName));
}

void QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->~QDesignerWidgetBoxInterface();
}

void QDesignerWidgetBoxInterface_ActionEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerWidgetBoxInterface_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DropEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerWidgetBoxInterface_DropEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerWidgetBoxInterface_EnterEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusInEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_HideEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerWidgetBoxInterface_HideEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerWidgetBoxInterface_LeaveEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::leaveEvent(static_cast<QEvent*>(event));
}

int QDesignerWidgetBoxInterface_Metric(void* ptr, int m)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QDesignerWidgetBoxInterface_MetricDefault(void* ptr, int m)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QDesignerWidgetBoxInterface_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->minimumSizeHint()).height());
}

void* QDesignerWidgetBoxInterface_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::minimumSizeHint()).height());
}

void QDesignerWidgetBoxInterface_MoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerWidgetBoxInterface_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QDesignerWidgetBoxInterface_PaintEngine(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->paintEngine();
}

void* QDesignerWidgetBoxInterface_PaintEngineDefault(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::paintEngine();
}

void QDesignerWidgetBoxInterface_PaintEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerWidgetBoxInterface_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerWidgetBoxInterface_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QDesignerWidgetBoxInterface_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setEnabled(vbo != 0);
}

void QDesignerWidgetBoxInterface_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QDesignerWidgetBoxInterface_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setStyleSheet(QString(styleSheet));
}

void QDesignerWidgetBoxInterface_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QDesignerWidgetBoxInterface_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setVisible(visible != 0);
}

void QDesignerWidgetBoxInterface_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QDesignerWidgetBoxInterface_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setWindowModified(vbo != 0);
}

void QDesignerWidgetBoxInterface_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QDesignerWidgetBoxInterface_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setWindowTitle(QString(vqs));
}

void QDesignerWidgetBoxInterface_ShowEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerWidgetBoxInterface_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showEvent(static_cast<QShowEvent*>(event));
}

void* QDesignerWidgetBoxInterface_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->sizeHint()).height());
}

void* QDesignerWidgetBoxInterface_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::sizeHint()).width(), static_cast<QSize>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::sizeHint()).height());
}

void QDesignerWidgetBoxInterface_ChangeEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::changeEvent(static_cast<QEvent*>(event));
}

int QDesignerWidgetBoxInterface_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDesignerWidgetBoxInterface_CloseDefault(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::close();
}

void QDesignerWidgetBoxInterface_CloseEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerWidgetBoxInterface_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerWidgetBoxInterface_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerWidgetBoxInterface_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QDesignerWidgetBoxInterface_Event(void* ptr, void* event)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDesignerWidgetBoxInterface_EventDefault(void* ptr, void* event)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::event(static_cast<QEvent*>(event));
}

int QDesignerWidgetBoxInterface_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->focusNextPrevChild(next != 0);
}

int QDesignerWidgetBoxInterface_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusNextPrevChild(next != 0);
}

int QDesignerWidgetBoxInterface_HasHeightForWidth(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->hasHeightForWidth();
}

int QDesignerWidgetBoxInterface_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hasHeightForWidth();
}

int QDesignerWidgetBoxInterface_HeightForWidth(void* ptr, int w)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->heightForWidth(w);
}

int QDesignerWidgetBoxInterface_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::heightForWidth(w);
}

void QDesignerWidgetBoxInterface_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "hide");
}

void QDesignerWidgetBoxInterface_HideDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hide();
}

void QDesignerWidgetBoxInterface_InitPainter(void* ptr, void* painter)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QDesignerWidgetBoxInterface_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::initPainter(static_cast<QPainter*>(painter));
}

void QDesignerWidgetBoxInterface_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerWidgetBoxInterface_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QDesignerWidgetBoxInterface_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerWidgetBoxInterface*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QDesignerWidgetBoxInterface_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QDesignerWidgetBoxInterface_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "lower");
}

void QDesignerWidgetBoxInterface_LowerDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::lower();
}

void QDesignerWidgetBoxInterface_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MousePressEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QDesignerWidgetBoxInterface_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QDesignerWidgetBoxInterface_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::nativeEvent(QByteArray(eventType), message, &result);
}

void QDesignerWidgetBoxInterface_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "raise");
}

void QDesignerWidgetBoxInterface_RaiseDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::raise();
}

void QDesignerWidgetBoxInterface_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "repaint");
}

void QDesignerWidgetBoxInterface_RepaintDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::repaint();
}

void QDesignerWidgetBoxInterface_ResizeEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerWidgetBoxInterface_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerWidgetBoxInterface_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QDesignerWidgetBoxInterface_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setDisabled(disable != 0);
}

void QDesignerWidgetBoxInterface_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setFocus");
}

void QDesignerWidgetBoxInterface_SetFocus2Default(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setFocus();
}

void QDesignerWidgetBoxInterface_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QDesignerWidgetBoxInterface_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setHidden(hidden != 0);
}

void QDesignerWidgetBoxInterface_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "show");
}

void QDesignerWidgetBoxInterface_ShowDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::show();
}

void QDesignerWidgetBoxInterface_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "showFullScreen");
}

void QDesignerWidgetBoxInterface_ShowFullScreenDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showFullScreen();
}

void QDesignerWidgetBoxInterface_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "showMaximized");
}

void QDesignerWidgetBoxInterface_ShowMaximizedDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showMaximized();
}

void QDesignerWidgetBoxInterface_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "showMinimized");
}

void QDesignerWidgetBoxInterface_ShowMinimizedDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showMinimized();
}

void QDesignerWidgetBoxInterface_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "showNormal");
}

void QDesignerWidgetBoxInterface_ShowNormalDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showNormal();
}

void QDesignerWidgetBoxInterface_TabletEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerWidgetBoxInterface_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerWidgetBoxInterface_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "update");
}

void QDesignerWidgetBoxInterface_UpdateDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::update();
}

void QDesignerWidgetBoxInterface_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "updateMicroFocus");
}

void QDesignerWidgetBoxInterface_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::updateMicroFocus();
}

void QDesignerWidgetBoxInterface_WheelEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerWidgetBoxInterface_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QDesignerWidgetBoxInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerWidgetBoxInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDesignerWidgetBoxInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerWidgetBoxInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerWidgetBoxInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerWidgetBoxInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerWidgetBoxInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDesignerWidgetBoxInterface*>(ptr), "deleteLater");
}

void QDesignerWidgetBoxInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::deleteLater();
}

void QDesignerWidgetBoxInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerWidgetBoxInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDesignerWidgetBoxInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDesignerWidgetBoxInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDesignerWidgetBoxInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->metaObject());
}

void* QDesignerWidgetBoxInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::metaObject());
}

class MyQExtensionFactory: public QExtensionFactory
{
public:
	MyQExtensionFactory(QExtensionManager *parent) : QExtensionFactory(parent) {};
	QObject * createExtension(QObject * object, const QString & iid, QObject * parent) const { return static_cast<QObject*>(callbackQExtensionFactory_CreateExtension(const_cast<MyQExtensionFactory*>(this), this->objectName().toUtf8().data(), object, iid.toUtf8().data(), parent)); };
	QObject * extension(QObject * object, const QString & iid) const { return static_cast<QObject*>(callbackQExtensionFactory_Extension(const_cast<MyQExtensionFactory*>(this), this->objectName().toUtf8().data(), object, iid.toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQExtensionFactory_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQExtensionFactory_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQExtensionFactory_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQExtensionFactory_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQExtensionFactory_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQExtensionFactory_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQExtensionFactory_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQExtensionFactory_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQExtensionFactory_MetaObject(const_cast<MyQExtensionFactory*>(this), this->objectName().toUtf8().data())); };
};

void* QExtensionFactory_NewQExtensionFactory(void* parent)
{
	return new MyQExtensionFactory(static_cast<QExtensionManager*>(parent));
}

void* QExtensionFactory_CreateExtension(void* ptr, void* object, char* iid, void* parent)
{
	return static_cast<QExtensionFactory*>(ptr)->createExtension(static_cast<QObject*>(object), QString(iid), static_cast<QObject*>(parent));
}

void* QExtensionFactory_CreateExtensionDefault(void* ptr, void* object, char* iid, void* parent)
{
	return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::createExtension(static_cast<QObject*>(object), QString(iid), static_cast<QObject*>(parent));
}

void* QExtensionFactory_Extension(void* ptr, void* object, char* iid)
{
	return static_cast<QExtensionFactory*>(ptr)->extension(static_cast<QObject*>(object), QString(iid));
}

void* QExtensionFactory_ExtensionDefault(void* ptr, void* object, char* iid)
{
	return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::extension(static_cast<QObject*>(object), QString(iid));
}

void* QExtensionFactory_ExtensionManager(void* ptr)
{
	return static_cast<QExtensionFactory*>(ptr)->extensionManager();
}

void QExtensionFactory_TimerEvent(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionFactory_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionFactory_ChildEvent(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QExtensionFactory_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QExtensionFactory_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QExtensionFactory*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionFactory_CustomEvent(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QExtensionFactory_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::customEvent(static_cast<QEvent*>(event));
}

void QExtensionFactory_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QExtensionFactory*>(ptr), "deleteLater");
}

void QExtensionFactory_DeleteLaterDefault(void* ptr)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::deleteLater();
}

void QExtensionFactory_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QExtensionFactory*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QExtensionFactory_Event(void* ptr, void* e)
{
	return static_cast<QExtensionFactory*>(ptr)->event(static_cast<QEvent*>(e));
}

int QExtensionFactory_EventDefault(void* ptr, void* e)
{
	return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::event(static_cast<QEvent*>(e));
}

int QExtensionFactory_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QExtensionFactory*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QExtensionFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QExtensionFactory_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QExtensionFactory*>(ptr)->metaObject());
}

void* QExtensionFactory_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::metaObject());
}

class MyQExtensionManager: public QExtensionManager
{
public:
	MyQExtensionManager(QObject *parent) : QExtensionManager(parent) {};
	QObject * extension(QObject * object, const QString & iid) const { return static_cast<QObject*>(callbackQExtensionManager_Extension(const_cast<MyQExtensionManager*>(this), this->objectName().toUtf8().data(), object, iid.toUtf8().data())); };
	void registerExtensions(QAbstractExtensionFactory * factory, const QString & iid) { callbackQExtensionManager_RegisterExtensions(this, this->objectName().toUtf8().data(), factory, iid.toUtf8().data()); };
	void unregisterExtensions(QAbstractExtensionFactory * factory, const QString & iid) { callbackQExtensionManager_UnregisterExtensions(this, this->objectName().toUtf8().data(), factory, iid.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQExtensionManager_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQExtensionManager_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQExtensionManager_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQExtensionManager_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQExtensionManager_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQExtensionManager_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQExtensionManager_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQExtensionManager_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQExtensionManager_MetaObject(const_cast<MyQExtensionManager*>(this), this->objectName().toUtf8().data())); };
};

void* QExtensionManager_NewQExtensionManager(void* parent)
{
	return new MyQExtensionManager(static_cast<QObject*>(parent));
}

void* QExtensionManager_Extension(void* ptr, void* object, char* iid)
{
	return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QObject*>(object), QString(iid));
}

void* QExtensionManager_ExtensionDefault(void* ptr, void* object, char* iid)
{
	return static_cast<QExtensionManager*>(ptr)->QExtensionManager::extension(static_cast<QObject*>(object), QString(iid));
}

void QExtensionManager_RegisterExtensions(void* ptr, void* factory, char* iid)
{
	static_cast<QExtensionManager*>(ptr)->registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QExtensionManager_RegisterExtensionsDefault(void* ptr, void* factory, char* iid)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QExtensionManager_UnregisterExtensions(void* ptr, void* factory, char* iid)
{
	static_cast<QExtensionManager*>(ptr)->unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QExtensionManager_UnregisterExtensionsDefault(void* ptr, void* factory, char* iid)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString(iid));
}

void QExtensionManager_DestroyQExtensionManager(void* ptr)
{
	static_cast<QExtensionManager*>(ptr)->~QExtensionManager();
}

void QExtensionManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QExtensionManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::childEvent(static_cast<QChildEvent*>(event));
}

void QExtensionManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QExtensionManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QExtensionManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::customEvent(static_cast<QEvent*>(event));
}

void QExtensionManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QExtensionManager*>(ptr), "deleteLater");
}

void QExtensionManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::deleteLater();
}

void QExtensionManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QExtensionManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QExtensionManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QExtensionManager*>(ptr)->QExtensionManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QExtensionManager_Event(void* ptr, void* e)
{
	return static_cast<QExtensionManager*>(ptr)->event(static_cast<QEvent*>(e));
}

int QExtensionManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QExtensionManager*>(ptr)->QExtensionManager::event(static_cast<QEvent*>(e));
}

int QExtensionManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QExtensionManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QExtensionManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QExtensionManager*>(ptr)->QExtensionManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QExtensionManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QExtensionManager*>(ptr)->metaObject());
}

void* QExtensionManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QExtensionManager*>(ptr)->QExtensionManager::metaObject());
}

class MyQFormBuilder: public QFormBuilder
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQFormBuilder() : QFormBuilder() {};
	QWidget * load(QIODevice * device, QWidget * parent) { return static_cast<QWidget*>(callbackQFormBuilder_Load(this, this->objectNameAbs().toUtf8().data(), device, parent)); };
	void save(QIODevice * device, QWidget * widget) { callbackQFormBuilder_Save(this, this->objectNameAbs().toUtf8().data(), device, widget); };
};

void* QFormBuilder_NewQFormBuilder()
{
	return new MyQFormBuilder();
}

void QFormBuilder_AddPluginPath(void* ptr, char* pluginPath)
{
	static_cast<QFormBuilder*>(ptr)->addPluginPath(QString(pluginPath));
}

void QFormBuilder_ClearPluginPaths(void* ptr)
{
	static_cast<QFormBuilder*>(ptr)->clearPluginPaths();
}

char* QFormBuilder_PluginPaths(void* ptr)
{
	return static_cast<QFormBuilder*>(ptr)->pluginPaths().join("|").toUtf8().data();
}

void QFormBuilder_SetPluginPath(void* ptr, char* pluginPaths)
{
	static_cast<QFormBuilder*>(ptr)->setPluginPath(QString(pluginPaths).split("|", QString::SkipEmptyParts));
}

void QFormBuilder_DestroyQFormBuilder(void* ptr)
{
	static_cast<QFormBuilder*>(ptr)->~QFormBuilder();
}

char* QFormBuilder_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQFormBuilder*>(static_cast<QFormBuilder*>(ptr))) {
		return static_cast<MyQFormBuilder*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QFormBuilder_BASE").toUtf8().data();
}

void QFormBuilder_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQFormBuilder*>(static_cast<QFormBuilder*>(ptr))) {
		static_cast<MyQFormBuilder*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QFormBuilder_Load(void* ptr, void* device, void* parent)
{
	return static_cast<QFormBuilder*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
}

void* QFormBuilder_LoadDefault(void* ptr, void* device, void* parent)
{
	return static_cast<QFormBuilder*>(ptr)->QFormBuilder::load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
}

void QFormBuilder_Save(void* ptr, void* device, void* widget)
{
	static_cast<QFormBuilder*>(ptr)->save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
}

void QFormBuilder_SaveDefault(void* ptr, void* device, void* widget)
{
	static_cast<QFormBuilder*>(ptr)->QFormBuilder::save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
}

