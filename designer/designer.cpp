// +build !minimal

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
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
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
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIODevice>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPdfWriter>
#include <QPixmap>
#include <QPoint>
#include <QQuickItem>
#include <QRadioData>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSignalSpy>
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
#include <QWindow>

class MyQAbstractExtensionManager: public QAbstractExtensionManager
{
public:
	void registerExtensions(QAbstractExtensionFactory * factory, const QString & iid) { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };callbackQAbstractExtensionManager_RegisterExtensions(this, factory, iidPacked); };
	void unregisterExtensions(QAbstractExtensionFactory * factory, const QString & iid) { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };callbackQAbstractExtensionManager_UnregisterExtensions(this, factory, iidPacked); };
	 ~MyQAbstractExtensionManager() { callbackQAbstractExtensionManager_DestroyQAbstractExtensionManager(this); };
	QObject * extension(QObject * object, const QString & iid) const { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };return static_cast<QObject*>(callbackQAbstractExtensionManager_Extension(const_cast<void*>(static_cast<const void*>(this)), object, iidPacked)); };
};

void QAbstractExtensionManager_RegisterExtensions(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
			static_cast<QExtensionManager*>(ptr)->registerExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		} else {
			static_cast<QExtensionManager*>(ptr)->registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		}
	} else {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
			static_cast<QAbstractExtensionManager*>(ptr)->registerExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		} else {
			static_cast<QAbstractExtensionManager*>(ptr)->registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		}
	}
}

void QAbstractExtensionManager_UnregisterExtensions(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
			static_cast<QExtensionManager*>(ptr)->unregisterExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		} else {
			static_cast<QExtensionManager*>(ptr)->unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		}
	} else {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
			static_cast<QAbstractExtensionManager*>(ptr)->unregisterExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		} else {
			static_cast<QAbstractExtensionManager*>(ptr)->unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
		}
	}
}

void QAbstractExtensionManager_DestroyQAbstractExtensionManager(void* ptr)
{
	static_cast<QAbstractExtensionManager*>(ptr)->~QAbstractExtensionManager();
}

void QAbstractExtensionManager_DestroyQAbstractExtensionManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractExtensionManager_Extension(void* ptr, void* object, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
			return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
		} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
			return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
		} else {
			return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
		}
	} else {
		if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
			return static_cast<QAbstractExtensionManager*>(ptr)->extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
		} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
			return static_cast<QAbstractExtensionManager*>(ptr)->extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
		} else {
			return static_cast<QAbstractExtensionManager*>(ptr)->extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
		}
	}
}

class MyQAbstractFormBuilder: public QAbstractFormBuilder
{
public:
	MyQAbstractFormBuilder() : QAbstractFormBuilder() {};
	QWidget * load(QIODevice * device, QWidget * parent) { return static_cast<QWidget*>(callbackQAbstractFormBuilder_Load(this, device, parent)); };
	void save(QIODevice * device, QWidget * widget) { callbackQAbstractFormBuilder_Save(this, device, widget); };
	 ~MyQAbstractFormBuilder() { callbackQAbstractFormBuilder_DestroyQAbstractFormBuilder(this); };
};

void* QAbstractFormBuilder_NewQAbstractFormBuilder()
{
	return new MyQAbstractFormBuilder();
}

void* QAbstractFormBuilder_Load(void* ptr, void* device, void* parent)
{
	return static_cast<QAbstractFormBuilder*>(ptr)->load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
}

void* QAbstractFormBuilder_LoadDefault(void* ptr, void* device, void* parent)
{
	if (dynamic_cast<QFormBuilder*>(static_cast<QAbstractFormBuilder*>(ptr))) {
		return static_cast<QFormBuilder*>(ptr)->QFormBuilder::load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
	} else {
		return static_cast<QAbstractFormBuilder*>(ptr)->QAbstractFormBuilder::load(static_cast<QIODevice*>(device), static_cast<QWidget*>(parent));
	}
}

void QAbstractFormBuilder_Save(void* ptr, void* device, void* widget)
{
	static_cast<QAbstractFormBuilder*>(ptr)->save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
}

void QAbstractFormBuilder_SaveDefault(void* ptr, void* device, void* widget)
{
	if (dynamic_cast<QFormBuilder*>(static_cast<QAbstractFormBuilder*>(ptr))) {
		static_cast<QFormBuilder*>(ptr)->QFormBuilder::save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
	} else {
		static_cast<QAbstractFormBuilder*>(ptr)->QAbstractFormBuilder::save(static_cast<QIODevice*>(device), static_cast<QWidget*>(widget));
	}
}

void QAbstractFormBuilder_SetWorkingDirectory(void* ptr, void* directory)
{
	static_cast<QAbstractFormBuilder*>(ptr)->setWorkingDirectory(*static_cast<QDir*>(directory));
}

void QAbstractFormBuilder_DestroyQAbstractFormBuilder(void* ptr)
{
	static_cast<QAbstractFormBuilder*>(ptr)->~QAbstractFormBuilder();
}

void QAbstractFormBuilder_DestroyQAbstractFormBuilderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractFormBuilder_WorkingDirectory(void* ptr)
{
	return new QDir(static_cast<QAbstractFormBuilder*>(ptr)->workingDirectory());
}

struct QtDesigner_PackedString QAbstractFormBuilder_ErrorString(void* ptr)
{
	return ({ QByteArray t590e87 = static_cast<QAbstractFormBuilder*>(ptr)->errorString().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t590e87.prepend("WHITESPACE").constData()+10), t590e87.size()-10 }; });
}

void* QAbstractFormBuilder___propertyMap_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DomProperty *>;
}

void* QAbstractFormBuilder___computeProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DomProperty *>;
}

void* QAbstractFormBuilder___applyProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<DomProperty *>;
}

class MyQDesignerActionEditorInterface: public QDesignerActionEditorInterface
{
public:
	MyQDesignerActionEditorInterface(QWidget *parent, Qt::WindowFlags flags = Qt::WindowFlags()) : QDesignerActionEditorInterface(parent, flags) {QDesignerActionEditorInterface_QDesignerActionEditorInterface_QRegisterMetaType();};
	void manageAction(QAction * action) { callbackQDesignerActionEditorInterface_ManageAction(this, action); };
	void setFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerActionEditorInterface_SetFormWindow(this, formWindow); };
	void unmanageAction(QAction * action) { callbackQDesignerActionEditorInterface_UnmanageAction(this, action); };
	 ~MyQDesignerActionEditorInterface() { callbackQDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface(this); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerActionEditorInterface_Core(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQDesignerActionEditorInterface_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQDesignerActionEditorInterface_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerActionEditorInterface_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerActionEditorInterface_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQDesignerActionEditorInterface_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQDesignerActionEditorInterface_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQDesignerActionEditorInterface_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerActionEditorInterface_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQDesignerActionEditorInterface_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerActionEditorInterface_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerActionEditorInterface_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerActionEditorInterface_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerActionEditorInterface_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQDesignerActionEditorInterface_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerActionEditorInterface_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerActionEditorInterface_FocusOutEvent(this, event); };
	void hide() { callbackQDesignerActionEditorInterface_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQDesignerActionEditorInterface_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerActionEditorInterface_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerActionEditorInterface_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerActionEditorInterface_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQDesignerActionEditorInterface_LeaveEvent(this, event); };
	void lower() { callbackQDesignerActionEditorInterface_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerActionEditorInterface_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerActionEditorInterface_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerActionEditorInterface_PaintEvent(this, event); };
	void raise() { callbackQDesignerActionEditorInterface_Raise(this); };
	void repaint() { callbackQDesignerActionEditorInterface_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerActionEditorInterface_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQDesignerActionEditorInterface_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQDesignerActionEditorInterface_SetEnabled(this, vbo); };
	void setFocus() { callbackQDesignerActionEditorInterface_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQDesignerActionEditorInterface_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtDesigner_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQDesignerActionEditorInterface_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQDesignerActionEditorInterface_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQDesignerActionEditorInterface_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtDesigner_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQDesignerActionEditorInterface_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQDesignerActionEditorInterface_Show(this); };
	void showEvent(QShowEvent * event) { callbackQDesignerActionEditorInterface_ShowEvent(this, event); };
	void showFullScreen() { callbackQDesignerActionEditorInterface_ShowFullScreen(this); };
	void showMaximized() { callbackQDesignerActionEditorInterface_ShowMaximized(this); };
	void showMinimized() { callbackQDesignerActionEditorInterface_ShowMinimized(this); };
	void showNormal() { callbackQDesignerActionEditorInterface_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerActionEditorInterface_TabletEvent(this, event); };
	void update() { callbackQDesignerActionEditorInterface_Update(this); };
	void updateMicroFocus() { callbackQDesignerActionEditorInterface_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerActionEditorInterface_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQDesignerActionEditorInterface_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtDesigner_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQDesignerActionEditorInterface_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerActionEditorInterface_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerActionEditorInterface_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerActionEditorInterface_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerActionEditorInterface_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQDesignerActionEditorInterface_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerActionEditorInterface_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerActionEditorInterface_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerActionEditorInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerActionEditorInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerActionEditorInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerActionEditorInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerActionEditorInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerActionEditorInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerActionEditorInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerActionEditorInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerActionEditorInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerActionEditorInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerActionEditorInterface*)

int QDesignerActionEditorInterface_QDesignerActionEditorInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerActionEditorInterface*>(); return qRegisterMetaType<MyQDesignerActionEditorInterface*>();}

void* QDesignerActionEditorInterface_NewQDesignerActionEditorInterface(void* parent, long long flags)
{
		return new MyQDesignerActionEditorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
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

void QDesignerActionEditorInterface_DestroyQDesignerActionEditorInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerActionEditorInterface_Core(void* ptr)
{
	return static_cast<QDesignerActionEditorInterface*>(ptr)->core();
}

void* QDesignerActionEditorInterface_CoreDefault(void* ptr)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::core();
}

void* QDesignerActionEditorInterface___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerActionEditorInterface___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerActionEditorInterface___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerActionEditorInterface___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerActionEditorInterface___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerActionEditorInterface___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerActionEditorInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerActionEditorInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerActionEditorInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerActionEditorInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerActionEditorInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerActionEditorInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerActionEditorInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerActionEditorInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerActionEditorInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerActionEditorInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerActionEditorInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerActionEditorInterface_CloseDefault(void* ptr)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::close();
}

char QDesignerActionEditorInterface_EventDefault(void* ptr, void* event)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::event(static_cast<QEvent*>(event));
}

char QDesignerActionEditorInterface_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusNextPrevChild(next != 0);
}

char QDesignerActionEditorInterface_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QDesignerActionEditorInterface_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerActionEditorInterface_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::changeEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerActionEditorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerActionEditorInterface_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerActionEditorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerActionEditorInterface_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerActionEditorInterface_DropEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerActionEditorInterface_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerActionEditorInterface_HideDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hide();
}

void QDesignerActionEditorInterface_HideEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerActionEditorInterface_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerActionEditorInterface_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerActionEditorInterface_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_LowerDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::lower();
}

void QDesignerActionEditorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerActionEditorInterface_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerActionEditorInterface_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerActionEditorInterface_RaiseDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::raise();
}

void QDesignerActionEditorInterface_RepaintDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::repaint();
}

void QDesignerActionEditorInterface_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerActionEditorInterface_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setDisabled(disable != 0);
}

void QDesignerActionEditorInterface_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setEnabled(vbo != 0);
}

void QDesignerActionEditorInterface_SetFocus2Default(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setFocus();
}

void QDesignerActionEditorInterface_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setHidden(hidden != 0);
}

void QDesignerActionEditorInterface_SetStyleSheetDefault(void* ptr, struct QtDesigner_PackedString styleSheet)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QDesignerActionEditorInterface_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setVisible(visible != 0);
}

void QDesignerActionEditorInterface_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setWindowModified(vbo != 0);
}

void QDesignerActionEditorInterface_SetWindowTitleDefault(void* ptr, struct QtDesigner_PackedString vqs)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QDesignerActionEditorInterface_ShowDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::show();
}

void QDesignerActionEditorInterface_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerActionEditorInterface_ShowFullScreenDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showFullScreen();
}

void QDesignerActionEditorInterface_ShowMaximizedDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showMaximized();
}

void QDesignerActionEditorInterface_ShowMinimizedDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showMinimized();
}

void QDesignerActionEditorInterface_ShowNormalDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::showNormal();
}

void QDesignerActionEditorInterface_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerActionEditorInterface_UpdateDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::update();
}

void QDesignerActionEditorInterface_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::updateMicroFocus();
}

void QDesignerActionEditorInterface_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QDesignerActionEditorInterface_PaintEngineDefault(void* ptr)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::paintEngine();
}

void* QDesignerActionEditorInterface_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerActionEditorInterface_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerActionEditorInterface_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QDesignerActionEditorInterface_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::hasHeightForWidth();
}

int QDesignerActionEditorInterface_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::heightForWidth(w);
}

int QDesignerActionEditorInterface_MetricDefault(void* ptr, long long m)
{
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QDesignerActionEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerActionEditorInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerActionEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerActionEditorInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerActionEditorInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::deleteLater();
}

void QDesignerActionEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerActionEditorInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerActionEditorInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerActionEditorInterface*>(ptr)->QDesignerActionEditorInterface::metaObject());
}

class MyQDesignerContainerExtension: public QDesignerContainerExtension
{
public:
	void addWidget(QWidget * page) { callbackQDesignerContainerExtension_AddWidget(this, page); };
	void insertWidget(int index, QWidget * page) { callbackQDesignerContainerExtension_InsertWidget(this, index, page); };
	void remove(int index) { callbackQDesignerContainerExtension_Remove(this, index); };
	void setCurrentIndex(int index) { callbackQDesignerContainerExtension_SetCurrentIndex(this, index); };
	 ~MyQDesignerContainerExtension() { callbackQDesignerContainerExtension_DestroyQDesignerContainerExtension(this); };
	QWidget * widget(int index) const { return static_cast<QWidget*>(callbackQDesignerContainerExtension_Widget(const_cast<void*>(static_cast<const void*>(this)), index)); };
	bool canAddWidget() const { return callbackQDesignerContainerExtension_CanAddWidget(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool canRemove(int index) const { return callbackQDesignerContainerExtension_CanRemove(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	int count() const { return callbackQDesignerContainerExtension_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int currentIndex() const { return callbackQDesignerContainerExtension_CurrentIndex(const_cast<void*>(static_cast<const void*>(this))); };
};

void QDesignerContainerExtension_AddWidget(void* ptr, void* page)
{
	static_cast<QDesignerContainerExtension*>(ptr)->addWidget(static_cast<QWidget*>(page));
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

void QDesignerContainerExtension_DestroyQDesignerContainerExtension(void* ptr)
{
	static_cast<QDesignerContainerExtension*>(ptr)->~QDesignerContainerExtension();
}

void QDesignerContainerExtension_DestroyQDesignerContainerExtensionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerContainerExtension_Widget(void* ptr, int index)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->widget(index);
}

char QDesignerContainerExtension_CanAddWidget(void* ptr)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->canAddWidget();
}

char QDesignerContainerExtension_CanAddWidgetDefault(void* ptr)
{
		return static_cast<QDesignerContainerExtension*>(ptr)->QDesignerContainerExtension::canAddWidget();
}

char QDesignerContainerExtension_CanRemove(void* ptr, int index)
{
	return static_cast<QDesignerContainerExtension*>(ptr)->canRemove(index);
}

char QDesignerContainerExtension_CanRemoveDefault(void* ptr, int index)
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

class MyQDesignerCustomWidgetCollectionInterface: public QDesignerCustomWidgetCollectionInterface
{
public:
	 ~MyQDesignerCustomWidgetCollectionInterface() { callbackQDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(this); };
	QList<QDesignerCustomWidgetInterface *> customWidgets() const { return *static_cast<QList<QDesignerCustomWidgetInterface *>*>(callbackQDesignerCustomWidgetCollectionInterface_CustomWidgets(const_cast<void*>(static_cast<const void*>(this)))); };
};

void QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(void* ptr)
{
	static_cast<QDesignerCustomWidgetCollectionInterface*>(ptr)->~QDesignerCustomWidgetCollectionInterface();
}

void QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDesigner_PackedList QDesignerCustomWidgetCollectionInterface_CustomWidgets(void* ptr)
{
	return ({ QList<QDesignerCustomWidgetInterface *>* tmpValue = new QList<QDesignerCustomWidgetInterface *>(static_cast<QDesignerCustomWidgetCollectionInterface*>(ptr)->customWidgets()); QtDesigner_PackedList { tmpValue, tmpValue->size() }; });
}

void* QDesignerCustomWidgetCollectionInterface___customWidgets_atList(void* ptr, int i)
{
	return const_cast<QDesignerCustomWidgetInterface*>(static_cast<QList<QDesignerCustomWidgetInterface *>*>(ptr)->at(i));
}

void QDesignerCustomWidgetCollectionInterface___customWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDesignerCustomWidgetInterface *>*>(ptr)->append(static_cast<QDesignerCustomWidgetInterface*>(i));
}

void* QDesignerCustomWidgetCollectionInterface___customWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerCustomWidgetInterface *>;
}

class MyQDesignerCustomWidgetInterface: public QDesignerCustomWidgetInterface
{
public:
	QWidget * createWidget(QWidget * parent) { return static_cast<QWidget*>(callbackQDesignerCustomWidgetInterface_CreateWidget(this, parent)); };
	void initialize(QDesignerFormEditorInterface * formEditor) { callbackQDesignerCustomWidgetInterface_Initialize(this, formEditor); };
	 ~MyQDesignerCustomWidgetInterface() { callbackQDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(this); };
	QIcon icon() const { return *static_cast<QIcon*>(callbackQDesignerCustomWidgetInterface_Icon(const_cast<void*>(static_cast<const void*>(this)))); };
	QString codeTemplate() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_CodeTemplate(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString domXml() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_DomXml(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString group() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_Group(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString includeFile() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_IncludeFile(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString name() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_Name(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString toolTip() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_ToolTip(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString whatsThis() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerCustomWidgetInterface_WhatsThis(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool isContainer() const { return callbackQDesignerCustomWidgetInterface_IsContainer(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isInitialized() const { return callbackQDesignerCustomWidgetInterface_IsInitialized(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QDesignerCustomWidgetInterface_CreateWidget(void* ptr, void* parent)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->createWidget(static_cast<QWidget*>(parent));
}

void QDesignerCustomWidgetInterface_Initialize(void* ptr, void* formEditor)
{
	static_cast<QDesignerCustomWidgetInterface*>(ptr)->initialize(static_cast<QDesignerFormEditorInterface*>(formEditor));
}

void QDesignerCustomWidgetInterface_InitializeDefault(void* ptr, void* formEditor)
{
		static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::initialize(static_cast<QDesignerFormEditorInterface*>(formEditor));
}

void QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(void* ptr)
{
	static_cast<QDesignerCustomWidgetInterface*>(ptr)->~QDesignerCustomWidgetInterface();
}

void QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerCustomWidgetInterface_Icon(void* ptr)
{
	return new QIcon(static_cast<QDesignerCustomWidgetInterface*>(ptr)->icon());
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_CodeTemplate(void* ptr)
{
	return ({ QByteArray tca2cf4 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->codeTemplate().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tca2cf4.prepend("WHITESPACE").constData()+10), tca2cf4.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_CodeTemplateDefault(void* ptr)
{
		return ({ QByteArray t381b53 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::codeTemplate().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t381b53.prepend("WHITESPACE").constData()+10), t381b53.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_DomXml(void* ptr)
{
	return ({ QByteArray tb6d4a2 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->domXml().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tb6d4a2.prepend("WHITESPACE").constData()+10), tb6d4a2.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_DomXmlDefault(void* ptr)
{
		return ({ QByteArray tc81641 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::domXml().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tc81641.prepend("WHITESPACE").constData()+10), tc81641.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_Group(void* ptr)
{
	return ({ QByteArray t8fa903 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->group().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t8fa903.prepend("WHITESPACE").constData()+10), t8fa903.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_IncludeFile(void* ptr)
{
	return ({ QByteArray tc020b8 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->includeFile().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tc020b8.prepend("WHITESPACE").constData()+10), tc020b8.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_Name(void* ptr)
{
	return ({ QByteArray t1f39d1 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->name().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t1f39d1.prepend("WHITESPACE").constData()+10), t1f39d1.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_ToolTip(void* ptr)
{
	return ({ QByteArray t2cddae = static_cast<QDesignerCustomWidgetInterface*>(ptr)->toolTip().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t2cddae.prepend("WHITESPACE").constData()+10), t2cddae.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerCustomWidgetInterface_WhatsThis(void* ptr)
{
	return ({ QByteArray t647573 = static_cast<QDesignerCustomWidgetInterface*>(ptr)->whatsThis().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t647573.prepend("WHITESPACE").constData()+10), t647573.size()-10 }; });
}

char QDesignerCustomWidgetInterface_IsContainer(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->isContainer();
}

char QDesignerCustomWidgetInterface_IsInitialized(void* ptr)
{
	return static_cast<QDesignerCustomWidgetInterface*>(ptr)->isInitialized();
}

char QDesignerCustomWidgetInterface_IsInitializedDefault(void* ptr)
{
		return static_cast<QDesignerCustomWidgetInterface*>(ptr)->QDesignerCustomWidgetInterface::isInitialized();
}

class MyQDesignerDynamicPropertySheetExtension: public QDesignerDynamicPropertySheetExtension
{
public:
	bool removeDynamicProperty(int index) { return callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(this, index) != 0; };
	int addDynamicProperty(const QString & propertyName, const QVariant & value) { QByteArray tdeaeb2 = propertyName.toUtf8(); QtDesigner_PackedString propertyNamePacked = { const_cast<char*>(tdeaeb2.prepend("WHITESPACE").constData()+10), tdeaeb2.size()-10 };return callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty(this, propertyNamePacked, const_cast<QVariant*>(&value)); };
	 ~MyQDesignerDynamicPropertySheetExtension() { callbackQDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(this); };
	bool canAddDynamicProperty(const QString & propertyName) const { QByteArray tdeaeb2 = propertyName.toUtf8(); QtDesigner_PackedString propertyNamePacked = { const_cast<char*>(tdeaeb2.prepend("WHITESPACE").constData()+10), tdeaeb2.size()-10 };return callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(const_cast<void*>(static_cast<const void*>(this)), propertyNamePacked) != 0; };
	bool dynamicPropertiesAllowed() const { return callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isDynamicProperty(int index) const { return callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
};

char QDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(void* ptr, int index)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->removeDynamicProperty(index);
}

int QDesignerDynamicPropertySheetExtension_AddDynamicProperty(void* ptr, struct QtDesigner_PackedString propertyName, void* value)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->addDynamicProperty(QString::fromUtf8(propertyName.data, propertyName.len), *static_cast<QVariant*>(value));
}

void QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(void* ptr)
{
	static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->~QDesignerDynamicPropertySheetExtension();
}

void QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtensionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(void* ptr, struct QtDesigner_PackedString propertyName)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->canAddDynamicProperty(QString::fromUtf8(propertyName.data, propertyName.len));
}

char QDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(void* ptr)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->dynamicPropertiesAllowed();
}

char QDesignerDynamicPropertySheetExtension_IsDynamicProperty(void* ptr, int index)
{
	return static_cast<QDesignerDynamicPropertySheetExtension*>(ptr)->isDynamicProperty(index);
}

class MyQDesignerFormEditorInterface: public QDesignerFormEditorInterface
{
public:
	MyQDesignerFormEditorInterface(QObject *parent = Q_NULLPTR) : QDesignerFormEditorInterface(parent) {QDesignerFormEditorInterface_QDesignerFormEditorInterface_QRegisterMetaType();};
	 ~MyQDesignerFormEditorInterface() { callbackQDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(this); };
	bool event(QEvent * e) { return callbackQDesignerFormEditorInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormEditorInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerFormEditorInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormEditorInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormEditorInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerFormEditorInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerFormEditorInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormEditorInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerFormEditorInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormEditorInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormEditorInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerFormEditorInterface*)

int QDesignerFormEditorInterface_QDesignerFormEditorInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerFormEditorInterface*>(); return qRegisterMetaType<MyQDesignerFormEditorInterface*>();}

void* QDesignerFormEditorInterface_NewQDesignerFormEditorInterface(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDesignerFormEditorInterface(static_cast<QWindow*>(parent));
	} else {
		return new MyQDesignerFormEditorInterface(static_cast<QObject*>(parent));
	}
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

void QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(void* ptr)
{
	static_cast<QDesignerFormEditorInterface*>(ptr)->~QDesignerFormEditorInterface();
}

void QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerFormEditorInterface_ActionEditor(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->actionEditor();
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

void* QDesignerFormEditorInterface_WidgetBox(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->widgetBox();
}

void* QDesignerFormEditorInterface_ExtensionManager(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->extensionManager();
}

void* QDesignerFormEditorInterface_TopLevel(void* ptr)
{
	return static_cast<QDesignerFormEditorInterface*>(ptr)->topLevel();
}

void* QDesignerFormEditorInterface___setOptionsPages_optionsPages_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerOptionsPageInterface *>;
}

void* QDesignerFormEditorInterface___optionsPages_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerOptionsPageInterface *>;
}

void* QDesignerFormEditorInterface___pluginInstances_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___pluginInstances_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormEditorInterface___pluginInstances_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QDesignerFormEditorInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerFormEditorInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerFormEditorInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormEditorInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormEditorInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormEditorInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormEditorInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormEditorInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormEditorInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerFormEditorInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormEditorInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerFormEditorInterface_EventDefault(void* ptr, void* e)
{
		return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::event(static_cast<QEvent*>(e));
}

char QDesignerFormEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerFormEditorInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormEditorInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormEditorInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::deleteLater();
}

void QDesignerFormEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormEditorInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerFormEditorInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerFormEditorInterface*>(ptr)->QDesignerFormEditorInterface::metaObject());
}

class MyQDesignerFormWindowCursorInterface: public QDesignerFormWindowCursorInterface
{
public:
	bool movePosition(QDesignerFormWindowCursorInterface::MoveOperation operation, QDesignerFormWindowCursorInterface::MoveMode mode) { return callbackQDesignerFormWindowCursorInterface_MovePosition(this, operation, mode) != 0; };
	void resetWidgetProperty(QWidget * widget, const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty(this, widget, namePacked); };
	void setPosition(int position, QDesignerFormWindowCursorInterface::MoveMode mode) { callbackQDesignerFormWindowCursorInterface_SetPosition(this, position, mode); };
	void setProperty(const QString & name, const QVariant & value) { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDesignerFormWindowCursorInterface_SetProperty(this, namePacked, const_cast<QVariant*>(&value)); };
	void setWidgetProperty(QWidget * widget, const QString & name, const QVariant & value) { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDesignerFormWindowCursorInterface_SetWidgetProperty(this, widget, namePacked, const_cast<QVariant*>(&value)); };
	 ~MyQDesignerFormWindowCursorInterface() { callbackQDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(this); };
	QDesignerFormWindowInterface * formWindow() const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowCursorInterface_FormWindow(const_cast<void*>(static_cast<const void*>(this)))); };
	QWidget * current() const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_Current(const_cast<void*>(static_cast<const void*>(this)))); };
	QWidget * selectedWidget(int index) const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_SelectedWidget(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QWidget * widget(int index) const { return static_cast<QWidget*>(callbackQDesignerFormWindowCursorInterface_Widget(const_cast<void*>(static_cast<const void*>(this)), index)); };
	bool hasSelection() const { return callbackQDesignerFormWindowCursorInterface_HasSelection(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int position() const { return callbackQDesignerFormWindowCursorInterface_Position(const_cast<void*>(static_cast<const void*>(this))); };
	int selectedWidgetCount() const { return callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount(const_cast<void*>(static_cast<const void*>(this))); };
	int widgetCount() const { return callbackQDesignerFormWindowCursorInterface_WidgetCount(const_cast<void*>(static_cast<const void*>(this))); };
};

char QDesignerFormWindowCursorInterface_MovePosition(void* ptr, long long operation, long long mode)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->movePosition(static_cast<QDesignerFormWindowCursorInterface::MoveOperation>(operation), static_cast<QDesignerFormWindowCursorInterface::MoveMode>(mode));
}

void QDesignerFormWindowCursorInterface_ResetWidgetProperty(void* ptr, void* widget, struct QtDesigner_PackedString name)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->resetWidgetProperty(static_cast<QWidget*>(widget), QString::fromUtf8(name.data, name.len));
}

void QDesignerFormWindowCursorInterface_SetPosition(void* ptr, int position, long long mode)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setPosition(position, static_cast<QDesignerFormWindowCursorInterface::MoveMode>(mode));
}

void QDesignerFormWindowCursorInterface_SetProperty(void* ptr, struct QtDesigner_PackedString name, void* value)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value));
}

void QDesignerFormWindowCursorInterface_SetWidgetProperty(void* ptr, void* widget, struct QtDesigner_PackedString name, void* value)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->setWidgetProperty(static_cast<QWidget*>(widget), QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value));
}

void QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(void* ptr)
{
	static_cast<QDesignerFormWindowCursorInterface*>(ptr)->~QDesignerFormWindowCursorInterface();
}

void QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerFormWindowCursorInterface_FormWindow(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->formWindow();
}

void* QDesignerFormWindowCursorInterface_Current(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->current();
}

void* QDesignerFormWindowCursorInterface_SelectedWidget(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->selectedWidget(index);
}

void* QDesignerFormWindowCursorInterface_Widget(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->widget(index);
}

char QDesignerFormWindowCursorInterface_HasSelection(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->hasSelection();
}

char QDesignerFormWindowCursorInterface_IsWidgetSelected(void* ptr, void* widget)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->isWidgetSelected(static_cast<QWidget*>(widget));
}

int QDesignerFormWindowCursorInterface_Position(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->position();
}

int QDesignerFormWindowCursorInterface_SelectedWidgetCount(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->selectedWidgetCount();
}

int QDesignerFormWindowCursorInterface_WidgetCount(void* ptr)
{
	return static_cast<QDesignerFormWindowCursorInterface*>(ptr)->widgetCount();
}

class MyQDesignerFormWindowInterface: public QDesignerFormWindowInterface
{
public:
	bool setContents(QIODevice * device, QString * errorMessage) { QByteArray t3f2abc = errorMessage->toUtf8(); QtDesigner_PackedString errorMessagePacked = { const_cast<char*>(t3f2abc.prepend("WHITESPACE").constData()+10), t3f2abc.size()-10 };return callbackQDesignerFormWindowInterface_SetContents(this, device, errorMessagePacked) != 0; };
	bool setContents(const QString & contents) { QByteArray t4a756c = contents.toUtf8(); QtDesigner_PackedString contentsPacked = { const_cast<char*>(t4a756c.prepend("WHITESPACE").constData()+10), t4a756c.size()-10 };return callbackQDesignerFormWindowInterface_SetContents2(this, contentsPacked) != 0; };
	void Signal_AboutToUnmanageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_AboutToUnmanageWidget(this, widget); };
	void activateResourceFilePaths(const QStringList & paths, int * errorCount, QString * errorMessages) { QByteArray t0c6eec = paths.join("|").toUtf8(); QtDesigner_PackedString pathsPacked = { const_cast<char*>(t0c6eec.prepend("WHITESPACE").constData()+10), t0c6eec.size()-10 };QByteArray t99ab00 = errorMessages->toUtf8(); QtDesigner_PackedString errorMessagesPacked = { const_cast<char*>(t99ab00.prepend("WHITESPACE").constData()+10), t99ab00.size()-10 };callbackQDesignerFormWindowInterface_ActivateResourceFilePaths(this, pathsPacked, *errorCount, errorMessagesPacked); };
	void Signal_Activated(QWidget * widget) { callbackQDesignerFormWindowInterface_Activated(this, widget); };
	void addResourceFile(const QString & path) { QByteArray t3150ec = path.toUtf8(); QtDesigner_PackedString pathPacked = { const_cast<char*>(t3150ec.prepend("WHITESPACE").constData()+10), t3150ec.size()-10 };callbackQDesignerFormWindowInterface_AddResourceFile(this, pathPacked); };
	void Signal_Changed() { callbackQDesignerFormWindowInterface_Changed(this); };
	void clearSelection(bool update) { callbackQDesignerFormWindowInterface_ClearSelection(this, update); };
	void emitSelectionChanged() { callbackQDesignerFormWindowInterface_EmitSelectionChanged(this); };
	void Signal_FeatureChanged(QDesignerFormWindowInterface::Feature feature) { callbackQDesignerFormWindowInterface_FeatureChanged(this, feature); };
	void Signal_FileNameChanged(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtDesigner_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQDesignerFormWindowInterface_FileNameChanged(this, fileNamePacked); };
	void Signal_GeometryChanged() { callbackQDesignerFormWindowInterface_GeometryChanged(this); };
	void layoutDefault(int * margin, int * spacing) { callbackQDesignerFormWindowInterface_LayoutDefault(this, *margin, *spacing); };
	void layoutFunction(QString * margin, QString * spacing) { QByteArray t5033fb = margin->toUtf8(); QtDesigner_PackedString marginPacked = { const_cast<char*>(t5033fb.prepend("WHITESPACE").constData()+10), t5033fb.size()-10 };QByteArray t7561d7 = spacing->toUtf8(); QtDesigner_PackedString spacingPacked = { const_cast<char*>(t7561d7.prepend("WHITESPACE").constData()+10), t7561d7.size()-10 };callbackQDesignerFormWindowInterface_LayoutFunction(this, marginPacked, spacingPacked); };
	void Signal_MainContainerChanged(QWidget * mainContainer) { callbackQDesignerFormWindowInterface_MainContainerChanged(this, mainContainer); };
	void manageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_ManageWidget(this, widget); };
	void Signal_ObjectRemoved(QObject * object) { callbackQDesignerFormWindowInterface_ObjectRemoved(this, object); };
	void removeResourceFile(const QString & path) { QByteArray t3150ec = path.toUtf8(); QtDesigner_PackedString pathPacked = { const_cast<char*>(t3150ec.prepend("WHITESPACE").constData()+10), t3150ec.size()-10 };callbackQDesignerFormWindowInterface_RemoveResourceFile(this, pathPacked); };
	void Signal_ResourceFilesChanged() { callbackQDesignerFormWindowInterface_ResourceFilesChanged(this); };
	void selectWidget(QWidget * widget, bool sele) { callbackQDesignerFormWindowInterface_SelectWidget(this, widget, sele); };
	void Signal_SelectionChanged() { callbackQDesignerFormWindowInterface_SelectionChanged(this); };
	void setAuthor(const QString & author) { QByteArray tf64cd8 = author.toUtf8(); QtDesigner_PackedString authorPacked = { const_cast<char*>(tf64cd8.prepend("WHITESPACE").constData()+10), tf64cd8.size()-10 };callbackQDesignerFormWindowInterface_SetAuthor(this, authorPacked); };
	void setComment(const QString & comment) { QByteArray t118a99 = comment.toUtf8(); QtDesigner_PackedString commentPacked = { const_cast<char*>(t118a99.prepend("WHITESPACE").constData()+10), t118a99.size()-10 };callbackQDesignerFormWindowInterface_SetComment(this, commentPacked); };
	void setDirty(bool dirty) { callbackQDesignerFormWindowInterface_SetDirty(this, dirty); };
	void setExportMacro(const QString & exportMacro) { QByteArray td651f6 = exportMacro.toUtf8(); QtDesigner_PackedString exportMacroPacked = { const_cast<char*>(td651f6.prepend("WHITESPACE").constData()+10), td651f6.size()-10 };callbackQDesignerFormWindowInterface_SetExportMacro(this, exportMacroPacked); };
	void setFeatures(QDesignerFormWindowInterface::Feature features) { callbackQDesignerFormWindowInterface_SetFeatures(this, features); };
	void setFileName(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtDesigner_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQDesignerFormWindowInterface_SetFileName(this, fileNamePacked); };
	void setGrid(const QPoint & grid) { callbackQDesignerFormWindowInterface_SetGrid(this, const_cast<QPoint*>(&grid)); };
	void setIncludeHints(const QStringList & includeHints) { QByteArray t14dd6a = includeHints.join("|").toUtf8(); QtDesigner_PackedString includeHintsPacked = { const_cast<char*>(t14dd6a.prepend("WHITESPACE").constData()+10), t14dd6a.size()-10 };callbackQDesignerFormWindowInterface_SetIncludeHints(this, includeHintsPacked); };
	void setLayoutDefault(int margin, int spacing) { callbackQDesignerFormWindowInterface_SetLayoutDefault(this, margin, spacing); };
	void setLayoutFunction(const QString & margin, const QString & spacing) { QByteArray t5033fb = margin.toUtf8(); QtDesigner_PackedString marginPacked = { const_cast<char*>(t5033fb.prepend("WHITESPACE").constData()+10), t5033fb.size()-10 };QByteArray t7561d7 = spacing.toUtf8(); QtDesigner_PackedString spacingPacked = { const_cast<char*>(t7561d7.prepend("WHITESPACE").constData()+10), t7561d7.size()-10 };callbackQDesignerFormWindowInterface_SetLayoutFunction(this, marginPacked, spacingPacked); };
	void setMainContainer(QWidget * mainContainer) { callbackQDesignerFormWindowInterface_SetMainContainer(this, mainContainer); };
	void setPixmapFunction(const QString & pixmapFunction) { QByteArray t24ca9b = pixmapFunction.toUtf8(); QtDesigner_PackedString pixmapFunctionPacked = { const_cast<char*>(t24ca9b.prepend("WHITESPACE").constData()+10), t24ca9b.size()-10 };callbackQDesignerFormWindowInterface_SetPixmapFunction(this, pixmapFunctionPacked); };
	void setResourceFileSaveMode(QDesignerFormWindowInterface::ResourceFileSaveMode behaviour) { callbackQDesignerFormWindowInterface_SetResourceFileSaveMode(this, behaviour); };
	void unmanageWidget(QWidget * widget) { callbackQDesignerFormWindowInterface_UnmanageWidget(this, widget); };
	void Signal_WidgetManaged(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetManaged(this, widget); };
	void Signal_WidgetRemoved(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetRemoved(this, widget); };
	void Signal_WidgetUnmanaged(QWidget * widget) { callbackQDesignerFormWindowInterface_WidgetUnmanaged(this, widget); };
	 ~MyQDesignerFormWindowInterface() { callbackQDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface(this); };
	Feature features() const { return static_cast<QDesignerFormWindowInterface::FeatureFlag>(callbackQDesignerFormWindowInterface_Features(const_cast<void*>(static_cast<const void*>(this)))); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerFormWindowInterface_Core(const_cast<void*>(static_cast<const void*>(this)))); };
	QDesignerFormWindowCursorInterface * cursor() const { return static_cast<QDesignerFormWindowCursorInterface*>(callbackQDesignerFormWindowInterface_Cursor(const_cast<void*>(static_cast<const void*>(this)))); };
	QDir absoluteDir() const { return *static_cast<QDir*>(callbackQDesignerFormWindowInterface_AbsoluteDir(const_cast<void*>(static_cast<const void*>(this)))); };
	QPoint grid() const { return *static_cast<QPoint*>(callbackQDesignerFormWindowInterface_Grid(const_cast<void*>(static_cast<const void*>(this)))); };
	QString author() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_Author(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString comment() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_Comment(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString contents() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_Contents(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString exportMacro() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_ExportMacro(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString fileName() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString pixmapFunction() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_PixmapFunction(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QStringList checkContents() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_CheckContents(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList includeHints() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_IncludeHints(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList resourceFiles() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerFormWindowInterface_ResourceFiles(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QWidget * formContainer() const { return static_cast<QWidget*>(callbackQDesignerFormWindowInterface_FormContainer(const_cast<void*>(static_cast<const void*>(this)))); };
	ResourceFileSaveMode resourceFileSaveMode() const { return static_cast<QDesignerFormWindowInterface::ResourceFileSaveMode>(callbackQDesignerFormWindowInterface_ResourceFileSaveMode(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasFeature(QDesignerFormWindowInterface::Feature feature) const { return callbackQDesignerFormWindowInterface_HasFeature(const_cast<void*>(static_cast<const void*>(this)), feature) != 0; };
	bool isDirty() const { return callbackQDesignerFormWindowInterface_IsDirty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isManaged(QWidget * widget) const { return callbackQDesignerFormWindowInterface_IsManaged(const_cast<void*>(static_cast<const void*>(this)), widget) != 0; };
	bool close() { return callbackQDesignerFormWindowInterface_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQDesignerFormWindowInterface_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerFormWindowInterface_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerFormWindowInterface_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQDesignerFormWindowInterface_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQDesignerFormWindowInterface_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQDesignerFormWindowInterface_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerFormWindowInterface_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQDesignerFormWindowInterface_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerFormWindowInterface_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerFormWindowInterface_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerFormWindowInterface_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerFormWindowInterface_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQDesignerFormWindowInterface_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerFormWindowInterface_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerFormWindowInterface_FocusOutEvent(this, event); };
	void hide() { callbackQDesignerFormWindowInterface_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQDesignerFormWindowInterface_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerFormWindowInterface_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerFormWindowInterface_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerFormWindowInterface_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQDesignerFormWindowInterface_LeaveEvent(this, event); };
	void lower() { callbackQDesignerFormWindowInterface_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerFormWindowInterface_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerFormWindowInterface_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerFormWindowInterface_PaintEvent(this, event); };
	void raise() { callbackQDesignerFormWindowInterface_Raise(this); };
	void repaint() { callbackQDesignerFormWindowInterface_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerFormWindowInterface_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQDesignerFormWindowInterface_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQDesignerFormWindowInterface_SetEnabled(this, vbo); };
	void setFocus() { callbackQDesignerFormWindowInterface_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQDesignerFormWindowInterface_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtDesigner_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQDesignerFormWindowInterface_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQDesignerFormWindowInterface_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQDesignerFormWindowInterface_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtDesigner_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQDesignerFormWindowInterface_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQDesignerFormWindowInterface_Show(this); };
	void showEvent(QShowEvent * event) { callbackQDesignerFormWindowInterface_ShowEvent(this, event); };
	void showFullScreen() { callbackQDesignerFormWindowInterface_ShowFullScreen(this); };
	void showMaximized() { callbackQDesignerFormWindowInterface_ShowMaximized(this); };
	void showMinimized() { callbackQDesignerFormWindowInterface_ShowMinimized(this); };
	void showNormal() { callbackQDesignerFormWindowInterface_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerFormWindowInterface_TabletEvent(this, event); };
	void update() { callbackQDesignerFormWindowInterface_Update(this); };
	void updateMicroFocus() { callbackQDesignerFormWindowInterface_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerFormWindowInterface_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQDesignerFormWindowInterface_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtDesigner_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQDesignerFormWindowInterface_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerFormWindowInterface_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerFormWindowInterface_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerFormWindowInterface_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerFormWindowInterface_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQDesignerFormWindowInterface_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerFormWindowInterface_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerFormWindowInterface_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormWindowInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerFormWindowInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormWindowInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerFormWindowInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerFormWindowInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerFormWindowInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormWindowInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormWindowInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerFormWindowInterface*)

int QDesignerFormWindowInterface_QDesignerFormWindowInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerFormWindowInterface*>(); return qRegisterMetaType<MyQDesignerFormWindowInterface*>();}

void* QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(void* object)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return QDesignerFormWindowInterface::findFormWindow(static_cast<QExtensionFactory*>(object));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return QDesignerFormWindowInterface::findFormWindow(static_cast<QExtensionManager*>(object));
	} else {
		return QDesignerFormWindowInterface::findFormWindow(static_cast<QObject*>(object));
	}
}

void* QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(void* widget)
{
	return QDesignerFormWindowInterface::findFormWindow(static_cast<QWidget*>(widget));
}

char QDesignerFormWindowInterface_SetContents(void* ptr, void* device, struct QtDesigner_PackedString errorMessage)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->setContents(static_cast<QIODevice*>(device), new QString(QString::fromUtf8(errorMessage.data, errorMessage.len)));
}

char QDesignerFormWindowInterface_SetContents2(void* ptr, struct QtDesigner_PackedString contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setContents", Q_RETURN_ARG(bool, returnArg), Q_ARG(const QString, QString::fromUtf8(contents.data, contents.len)));
	return returnArg;
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

void QDesignerFormWindowInterface_ActivateResourceFilePaths(void* ptr, struct QtDesigner_PackedString paths, int errorCount, struct QtDesigner_PackedString errorMessages)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "activateResourceFilePaths", Q_ARG(const QStringList, QString::fromUtf8(paths.data, paths.len).split("|", QString::SkipEmptyParts)), Q_ARG(int*, &errorCount), Q_ARG(QString*, new QString(QString::fromUtf8(errorMessages.data, errorMessages.len))));
}

void QDesignerFormWindowInterface_ActivateResourceFilePathsDefault(void* ptr, struct QtDesigner_PackedString paths, int errorCount, struct QtDesigner_PackedString errorMessages)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::activateResourceFilePaths(QString::fromUtf8(paths.data, paths.len).split("|", QString::SkipEmptyParts), &errorCount, new QString(QString::fromUtf8(errorMessages.data, errorMessages.len)));
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

void QDesignerFormWindowInterface_AddResourceFile(void* ptr, struct QtDesigner_PackedString path)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->addResourceFile(QString::fromUtf8(path.data, path.len));
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

void QDesignerFormWindowInterface_ClearSelection(void* ptr, char update)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "clearSelection", Q_ARG(bool, update != 0));
}

void QDesignerFormWindowInterface_EmitSelectionChanged(void* ptr)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->emitSelectionChanged();
}

void QDesignerFormWindowInterface_ConnectFeatureChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&QDesignerFormWindowInterface::featureChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&MyQDesignerFormWindowInterface::Signal_FeatureChanged));
}

void QDesignerFormWindowInterface_DisconnectFeatureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&QDesignerFormWindowInterface::featureChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(QDesignerFormWindowInterface::Feature)>(&MyQDesignerFormWindowInterface::Signal_FeatureChanged));
}

void QDesignerFormWindowInterface_FeatureChanged(void* ptr, long long feature)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->featureChanged(static_cast<QDesignerFormWindowInterface::FeatureFlag>(feature));
}

void QDesignerFormWindowInterface_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(const QString &)>(&QDesignerFormWindowInterface::fileNameChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(const QString &)>(&MyQDesignerFormWindowInterface::Signal_FileNameChanged));
}

void QDesignerFormWindowInterface_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerFormWindowInterface*>(ptr), static_cast<void (QDesignerFormWindowInterface::*)(const QString &)>(&QDesignerFormWindowInterface::fileNameChanged), static_cast<MyQDesignerFormWindowInterface*>(ptr), static_cast<void (MyQDesignerFormWindowInterface::*)(const QString &)>(&MyQDesignerFormWindowInterface::Signal_FileNameChanged));
}

void QDesignerFormWindowInterface_FileNameChanged(void* ptr, struct QtDesigner_PackedString fileName)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
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

void QDesignerFormWindowInterface_LayoutDefault(void* ptr, int margin, int spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->layoutDefault(&margin, &spacing);
}

void QDesignerFormWindowInterface_LayoutFunction(void* ptr, struct QtDesigner_PackedString margin, struct QtDesigner_PackedString spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->layoutFunction(new QString(QString::fromUtf8(margin.data, margin.len)), new QString(QString::fromUtf8(spacing.data, spacing.len)));
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
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		static_cast<QDesignerFormWindowInterface*>(ptr)->objectRemoved(static_cast<QExtensionFactory*>(object));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		static_cast<QDesignerFormWindowInterface*>(ptr)->objectRemoved(static_cast<QExtensionManager*>(object));
	} else {
		static_cast<QDesignerFormWindowInterface*>(ptr)->objectRemoved(static_cast<QObject*>(object));
	}
}

void QDesignerFormWindowInterface_RemoveResourceFile(void* ptr, struct QtDesigner_PackedString path)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->removeResourceFile(QString::fromUtf8(path.data, path.len));
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

void QDesignerFormWindowInterface_SelectWidget(void* ptr, void* widget, char sele)
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

void QDesignerFormWindowInterface_SetAuthor(void* ptr, struct QtDesigner_PackedString author)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setAuthor(QString::fromUtf8(author.data, author.len));
}

void QDesignerFormWindowInterface_SetComment(void* ptr, struct QtDesigner_PackedString comment)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setComment(QString::fromUtf8(comment.data, comment.len));
}

void QDesignerFormWindowInterface_SetDirty(void* ptr, char dirty)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setDirty", Q_ARG(bool, dirty != 0));
}

void QDesignerFormWindowInterface_SetExportMacro(void* ptr, struct QtDesigner_PackedString exportMacro)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setExportMacro(QString::fromUtf8(exportMacro.data, exportMacro.len));
}

void QDesignerFormWindowInterface_SetFeatures(void* ptr, long long features)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setFeatures", Q_ARG(QDesignerFormWindowInterface::Feature, static_cast<QDesignerFormWindowInterface::FeatureFlag>(features)));
}

void QDesignerFormWindowInterface_SetFileName(void* ptr, struct QtDesigner_PackedString fileName)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setFileName", Q_ARG(const QString, QString::fromUtf8(fileName.data, fileName.len)));
}

void QDesignerFormWindowInterface_SetGrid(void* ptr, void* grid)
{
	QMetaObject::invokeMethod(static_cast<QDesignerFormWindowInterface*>(ptr), "setGrid", Q_ARG(const QPoint, *static_cast<QPoint*>(grid)));
}

void QDesignerFormWindowInterface_SetIncludeHints(void* ptr, struct QtDesigner_PackedString includeHints)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setIncludeHints(QString::fromUtf8(includeHints.data, includeHints.len).split("|", QString::SkipEmptyParts));
}

void QDesignerFormWindowInterface_SetLayoutDefault(void* ptr, int margin, int spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setLayoutDefault(margin, spacing);
}

void QDesignerFormWindowInterface_SetLayoutFunction(void* ptr, struct QtDesigner_PackedString margin, struct QtDesigner_PackedString spacing)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setLayoutFunction(QString::fromUtf8(margin.data, margin.len), QString::fromUtf8(spacing.data, spacing.len));
}

void QDesignerFormWindowInterface_SetMainContainer(void* ptr, void* mainContainer)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setMainContainer(static_cast<QWidget*>(mainContainer));
}

void QDesignerFormWindowInterface_SetPixmapFunction(void* ptr, struct QtDesigner_PackedString pixmapFunction)
{
	static_cast<QDesignerFormWindowInterface*>(ptr)->setPixmapFunction(QString::fromUtf8(pixmapFunction.data, pixmapFunction.len));
}

void QDesignerFormWindowInterface_SetResourceFileSaveMode(void* ptr, long long behaviour)
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

void QDesignerFormWindowInterface_DestroyQDesignerFormWindowInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QDesignerFormWindowInterface_Features(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->features();
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

void* QDesignerFormWindowInterface_AbsoluteDir(void* ptr)
{
	return new QDir(static_cast<QDesignerFormWindowInterface*>(ptr)->absoluteDir());
}

void* QDesignerFormWindowInterface_Grid(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QDesignerFormWindowInterface*>(ptr)->grid(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_Author(void* ptr)
{
	return ({ QByteArray tbc2d85 = static_cast<QDesignerFormWindowInterface*>(ptr)->author().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tbc2d85.prepend("WHITESPACE").constData()+10), tbc2d85.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_Comment(void* ptr)
{
	return ({ QByteArray t9cffdd = static_cast<QDesignerFormWindowInterface*>(ptr)->comment().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t9cffdd.prepend("WHITESPACE").constData()+10), t9cffdd.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_Contents(void* ptr)
{
	return ({ QByteArray taa8971 = static_cast<QDesignerFormWindowInterface*>(ptr)->contents().toUtf8(); QtDesigner_PackedString { const_cast<char*>(taa8971.prepend("WHITESPACE").constData()+10), taa8971.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_ExportMacro(void* ptr)
{
	return ({ QByteArray t94104e = static_cast<QDesignerFormWindowInterface*>(ptr)->exportMacro().toUtf8(); QtDesigner_PackedString { const_cast<char*>(t94104e.prepend("WHITESPACE").constData()+10), t94104e.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_FileName(void* ptr)
{
	return ({ QByteArray tf7a1cc = static_cast<QDesignerFormWindowInterface*>(ptr)->fileName().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tf7a1cc.prepend("WHITESPACE").constData()+10), tf7a1cc.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_PixmapFunction(void* ptr)
{
	return ({ QByteArray taa7947 = static_cast<QDesignerFormWindowInterface*>(ptr)->pixmapFunction().toUtf8(); QtDesigner_PackedString { const_cast<char*>(taa7947.prepend("WHITESPACE").constData()+10), taa7947.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_ActiveResourceFilePaths(void* ptr)
{
	return ({ QByteArray t8dca06 = static_cast<QDesignerFormWindowInterface*>(ptr)->activeResourceFilePaths().join("|").toUtf8(); QtDesigner_PackedString { const_cast<char*>(t8dca06.prepend("WHITESPACE").constData()+10), t8dca06.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_CheckContents(void* ptr)
{
	return ({ QByteArray t13d91e = static_cast<QDesignerFormWindowInterface*>(ptr)->checkContents().join("|").toUtf8(); QtDesigner_PackedString { const_cast<char*>(t13d91e.prepend("WHITESPACE").constData()+10), t13d91e.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_IncludeHints(void* ptr)
{
	return ({ QByteArray ta8d26f = static_cast<QDesignerFormWindowInterface*>(ptr)->includeHints().join("|").toUtf8(); QtDesigner_PackedString { const_cast<char*>(ta8d26f.prepend("WHITESPACE").constData()+10), ta8d26f.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerFormWindowInterface_ResourceFiles(void* ptr)
{
	return ({ QByteArray t55124a = static_cast<QDesignerFormWindowInterface*>(ptr)->resourceFiles().join("|").toUtf8(); QtDesigner_PackedString { const_cast<char*>(t55124a.prepend("WHITESPACE").constData()+10), t55124a.size()-10 }; });
}

void* QDesignerFormWindowInterface_FormContainer(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->formContainer();
}

long long QDesignerFormWindowInterface_ResourceFileSaveMode(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->resourceFileSaveMode();
}

char QDesignerFormWindowInterface_HasFeature(void* ptr, long long feature)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->hasFeature(static_cast<QDesignerFormWindowInterface::FeatureFlag>(feature));
}

char QDesignerFormWindowInterface_IsDirty(void* ptr)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->isDirty();
}

char QDesignerFormWindowInterface_IsManaged(void* ptr, void* widget)
{
	return static_cast<QDesignerFormWindowInterface*>(ptr)->isManaged(static_cast<QWidget*>(widget));
}

void* QDesignerFormWindowInterface___simplifySelection_widgets_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___simplifySelection_widgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QDesignerFormWindowInterface___simplifySelection_widgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>;
}

void* QDesignerFormWindowInterface___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerFormWindowInterface___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerFormWindowInterface___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerFormWindowInterface___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerFormWindowInterface___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerFormWindowInterface___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerFormWindowInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerFormWindowInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerFormWindowInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerFormWindowInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerFormWindowInterface_CloseDefault(void* ptr)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::close();
}

char QDesignerFormWindowInterface_EventDefault(void* ptr, void* event)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::event(static_cast<QEvent*>(event));
}

char QDesignerFormWindowInterface_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusNextPrevChild(next != 0);
}

char QDesignerFormWindowInterface_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QDesignerFormWindowInterface_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerFormWindowInterface_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::changeEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerFormWindowInterface_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerFormWindowInterface_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerFormWindowInterface_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerFormWindowInterface_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerFormWindowInterface_DropEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerFormWindowInterface_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerFormWindowInterface_HideDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hide();
}

void QDesignerFormWindowInterface_HideEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerFormWindowInterface_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerFormWindowInterface_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerFormWindowInterface_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_LowerDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::lower();
}

void QDesignerFormWindowInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerFormWindowInterface_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerFormWindowInterface_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerFormWindowInterface_RaiseDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::raise();
}

void QDesignerFormWindowInterface_RepaintDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::repaint();
}

void QDesignerFormWindowInterface_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerFormWindowInterface_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setDisabled(disable != 0);
}

void QDesignerFormWindowInterface_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setEnabled(vbo != 0);
}

void QDesignerFormWindowInterface_SetFocus2Default(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setFocus();
}

void QDesignerFormWindowInterface_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setHidden(hidden != 0);
}

void QDesignerFormWindowInterface_SetStyleSheetDefault(void* ptr, struct QtDesigner_PackedString styleSheet)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QDesignerFormWindowInterface_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setVisible(visible != 0);
}

void QDesignerFormWindowInterface_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setWindowModified(vbo != 0);
}

void QDesignerFormWindowInterface_SetWindowTitleDefault(void* ptr, struct QtDesigner_PackedString vqs)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QDesignerFormWindowInterface_ShowDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::show();
}

void QDesignerFormWindowInterface_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerFormWindowInterface_ShowFullScreenDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showFullScreen();
}

void QDesignerFormWindowInterface_ShowMaximizedDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showMaximized();
}

void QDesignerFormWindowInterface_ShowMinimizedDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showMinimized();
}

void QDesignerFormWindowInterface_ShowNormalDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::showNormal();
}

void QDesignerFormWindowInterface_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerFormWindowInterface_UpdateDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::update();
}

void QDesignerFormWindowInterface_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::updateMicroFocus();
}

void QDesignerFormWindowInterface_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QDesignerFormWindowInterface_PaintEngineDefault(void* ptr)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::paintEngine();
}

void* QDesignerFormWindowInterface_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerFormWindowInterface_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerFormWindowInterface_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QDesignerFormWindowInterface_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::hasHeightForWidth();
}

int QDesignerFormWindowInterface_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::heightForWidth(w);
}

int QDesignerFormWindowInterface_MetricDefault(void* ptr, long long m)
{
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QDesignerFormWindowInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerFormWindowInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::deleteLater();
}

void QDesignerFormWindowInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerFormWindowInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowInterface*>(ptr)->QDesignerFormWindowInterface::metaObject());
}

class MyQDesignerFormWindowManagerInterface: public QDesignerFormWindowManagerInterface
{
public:
	QDesignerFormWindowInterface * createFormWindow(QWidget * parent, Qt::WindowFlags flags) { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_CreateFormWindow(this, parent, flags)); };
	void Signal_ActiveFormWindowChanged(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged(this, formWindow); };
	void addFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_AddFormWindow(this, formWindow); };
	void closeAllPreviews() { callbackQDesignerFormWindowManagerInterface_CloseAllPreviews(this); };
	void Signal_FormWindowAdded(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowAdded(this, formWindow); };
	void Signal_FormWindowRemoved(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowRemoved(this, formWindow); };
	void Signal_FormWindowSettingsChanged(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged(this, formWindow); };
	void removeFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_RemoveFormWindow(this, formWindow); };
	void setActiveFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow(this, formWindow); };
	void showPluginDialog() { callbackQDesignerFormWindowManagerInterface_ShowPluginDialog(this); };
	void showPreview() { callbackQDesignerFormWindowManagerInterface_ShowPreview(this); };
	 ~MyQDesignerFormWindowManagerInterface() { callbackQDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface(this); };
	QAction * action(QDesignerFormWindowManagerInterface::Action action) const { return static_cast<QAction*>(callbackQDesignerFormWindowManagerInterface_Action(const_cast<void*>(static_cast<const void*>(this)), action)); };
	QActionGroup * actionGroup(QDesignerFormWindowManagerInterface::ActionGroup actionGroup) const { return static_cast<QActionGroup*>(callbackQDesignerFormWindowManagerInterface_ActionGroup(const_cast<void*>(static_cast<const void*>(this)), actionGroup)); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerFormWindowManagerInterface_Core(const_cast<void*>(static_cast<const void*>(this)))); };
	QDesignerFormWindowInterface * activeFormWindow() const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_ActiveFormWindow(const_cast<void*>(static_cast<const void*>(this)))); };
	QDesignerFormWindowInterface * formWindow(int index) const { return static_cast<QDesignerFormWindowInterface*>(callbackQDesignerFormWindowManagerInterface_FormWindow(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QPixmap createPreviewPixmap() const { return *static_cast<QPixmap*>(callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap(const_cast<void*>(static_cast<const void*>(this)))); };
	int formWindowCount() const { return callbackQDesignerFormWindowManagerInterface_FormWindowCount(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQDesignerFormWindowManagerInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerFormWindowManagerInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerFormWindowManagerInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowManagerInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerFormWindowManagerInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerFormWindowManagerInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerFormWindowManagerInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerFormWindowManagerInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerFormWindowManagerInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerFormWindowManagerInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerFormWindowManagerInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerFormWindowManagerInterface*)

int QDesignerFormWindowManagerInterface_QDesignerFormWindowManagerInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerFormWindowManagerInterface*>(); return qRegisterMetaType<MyQDesignerFormWindowManagerInterface*>();}

void* QDesignerFormWindowManagerInterface_CreateFormWindow(void* ptr, void* parent, long long flags)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->createFormWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
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

void QDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerFormWindowManagerInterface_Action(void* ptr, long long action)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->action(static_cast<QDesignerFormWindowManagerInterface::Action>(action));
}

void* QDesignerFormWindowManagerInterface_ActionGroup(void* ptr, long long actionGroup)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->actionGroup(static_cast<QDesignerFormWindowManagerInterface::ActionGroup>(actionGroup));
}

void* QDesignerFormWindowManagerInterface_Core(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->core();
}

void* QDesignerFormWindowManagerInterface_ActiveFormWindow(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->activeFormWindow();
}

void* QDesignerFormWindowManagerInterface_FormWindow(void* ptr, int index)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindow(index);
}

void* QDesignerFormWindowManagerInterface_CreatePreviewPixmap(void* ptr)
{
	return new QPixmap(static_cast<QDesignerFormWindowManagerInterface*>(ptr)->createPreviewPixmap());
}

int QDesignerFormWindowManagerInterface_FormWindowCount(void* ptr)
{
	return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->formWindowCount();
}

void* QDesignerFormWindowManagerInterface___dragItems_item_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerDnDItemInterface *>;
}

void* QDesignerFormWindowManagerInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerFormWindowManagerInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerFormWindowManagerInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerFormWindowManagerInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowManagerInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowManagerInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowManagerInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowManagerInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowManagerInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowManagerInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerFormWindowManagerInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowManagerInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerFormWindowManagerInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerFormWindowManagerInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerFormWindowManagerInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerFormWindowManagerInterface_EventDefault(void* ptr, void* e)
{
		return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::event(static_cast<QEvent*>(e));
}

char QDesignerFormWindowManagerInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerFormWindowManagerInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerFormWindowManagerInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowManagerInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerFormWindowManagerInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::deleteLater();
}

void QDesignerFormWindowManagerInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerFormWindowManagerInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerFormWindowManagerInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerFormWindowManagerInterface*>(ptr)->QDesignerFormWindowManagerInterface::metaObject());
}

class MyQDesignerMemberSheetExtension: public QDesignerMemberSheetExtension
{
public:
	void setMemberGroup(int index, const QString & group) { QByteArray t64292b = group.toUtf8(); QtDesigner_PackedString groupPacked = { const_cast<char*>(t64292b.prepend("WHITESPACE").constData()+10), t64292b.size()-10 };callbackQDesignerMemberSheetExtension_SetMemberGroup(this, index, groupPacked); };
	void setVisible(int index, bool visible) { callbackQDesignerMemberSheetExtension_SetVisible(this, index, visible); };
	 ~MyQDesignerMemberSheetExtension() { callbackQDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(this); };
	QList<QByteArray> parameterNames(int index) const { return *static_cast<QList<QByteArray>*>(callbackQDesignerMemberSheetExtension_ParameterNames(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QList<QByteArray> parameterTypes(int index) const { return *static_cast<QList<QByteArray>*>(callbackQDesignerMemberSheetExtension_ParameterTypes(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QString declaredInClass(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerMemberSheetExtension_DeclaredInClass(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString memberGroup(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerMemberSheetExtension_MemberGroup(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString memberName(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerMemberSheetExtension_MemberName(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString signature(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerMemberSheetExtension_Signature(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool inheritedFromWidget(int index) const { return callbackQDesignerMemberSheetExtension_InheritedFromWidget(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isSignal(int index) const { return callbackQDesignerMemberSheetExtension_IsSignal(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isSlot(int index) const { return callbackQDesignerMemberSheetExtension_IsSlot(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isVisible(int index) const { return callbackQDesignerMemberSheetExtension_IsVisible(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	int count() const { return callbackQDesignerMemberSheetExtension_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int indexOf(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQDesignerMemberSheetExtension_IndexOf(const_cast<void*>(static_cast<const void*>(this)), namePacked); };
};

void QDesignerMemberSheetExtension_SetMemberGroup(void* ptr, int index, struct QtDesigner_PackedString group)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->setMemberGroup(index, QString::fromUtf8(group.data, group.len));
}

void QDesignerMemberSheetExtension_SetVisible(void* ptr, int index, char visible)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->setVisible(index, visible != 0);
}

void QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(void* ptr)
{
	static_cast<QDesignerMemberSheetExtension*>(ptr)->~QDesignerMemberSheetExtension();
}

void QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtensionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDesigner_PackedList QDesignerMemberSheetExtension_ParameterNames(void* ptr, int index)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QDesignerMemberSheetExtension*>(ptr)->parameterNames(index)); QtDesigner_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDesigner_PackedList QDesignerMemberSheetExtension_ParameterTypes(void* ptr, int index)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(static_cast<QDesignerMemberSheetExtension*>(ptr)->parameterTypes(index)); QtDesigner_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDesigner_PackedString QDesignerMemberSheetExtension_DeclaredInClass(void* ptr, int index)
{
	return ({ QByteArray te5097e = static_cast<QDesignerMemberSheetExtension*>(ptr)->declaredInClass(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(te5097e.prepend("WHITESPACE").constData()+10), te5097e.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerMemberSheetExtension_MemberGroup(void* ptr, int index)
{
	return ({ QByteArray t391e72 = static_cast<QDesignerMemberSheetExtension*>(ptr)->memberGroup(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(t391e72.prepend("WHITESPACE").constData()+10), t391e72.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerMemberSheetExtension_MemberName(void* ptr, int index)
{
	return ({ QByteArray ta1b0eb = static_cast<QDesignerMemberSheetExtension*>(ptr)->memberName(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(ta1b0eb.prepend("WHITESPACE").constData()+10), ta1b0eb.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerMemberSheetExtension_Signature(void* ptr, int index)
{
	return ({ QByteArray t6e65c6 = static_cast<QDesignerMemberSheetExtension*>(ptr)->signature(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(t6e65c6.prepend("WHITESPACE").constData()+10), t6e65c6.size()-10 }; });
}

char QDesignerMemberSheetExtension_InheritedFromWidget(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->inheritedFromWidget(index);
}

char QDesignerMemberSheetExtension_IsSignal(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isSignal(index);
}

char QDesignerMemberSheetExtension_IsSlot(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isSlot(index);
}

char QDesignerMemberSheetExtension_IsVisible(void* ptr, int index)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->isVisible(index);
}

int QDesignerMemberSheetExtension_Count(void* ptr)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->count();
}

int QDesignerMemberSheetExtension_IndexOf(void* ptr, struct QtDesigner_PackedString name)
{
	return static_cast<QDesignerMemberSheetExtension*>(ptr)->indexOf(QString::fromUtf8(name.data, name.len));
}

void* QDesignerMemberSheetExtension___parameterNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerMemberSheetExtension___parameterNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerMemberSheetExtension___parameterNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerMemberSheetExtension___parameterTypes_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerMemberSheetExtension___parameterTypes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerMemberSheetExtension___parameterTypes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

class MyQDesignerObjectInspectorInterface: public QDesignerObjectInspectorInterface
{
public:
	MyQDesignerObjectInspectorInterface(QWidget *parent, Qt::WindowFlags flags = Qt::WindowFlags()) : QDesignerObjectInspectorInterface(parent, flags) {QDesignerObjectInspectorInterface_QDesignerObjectInspectorInterface_QRegisterMetaType();};
	void setFormWindow(QDesignerFormWindowInterface * formWindow) { callbackQDesignerObjectInspectorInterface_SetFormWindow(this, formWindow); };
	 ~MyQDesignerObjectInspectorInterface() { callbackQDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(this); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerObjectInspectorInterface_Core(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQDesignerObjectInspectorInterface_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQDesignerObjectInspectorInterface_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerObjectInspectorInterface_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerObjectInspectorInterface_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQDesignerObjectInspectorInterface_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQDesignerObjectInspectorInterface_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerObjectInspectorInterface_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQDesignerObjectInspectorInterface_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerObjectInspectorInterface_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerObjectInspectorInterface_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerObjectInspectorInterface_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerObjectInspectorInterface_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerObjectInspectorInterface_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerObjectInspectorInterface_FocusOutEvent(this, event); };
	void hide() { callbackQDesignerObjectInspectorInterface_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQDesignerObjectInspectorInterface_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerObjectInspectorInterface_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerObjectInspectorInterface_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerObjectInspectorInterface_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_LeaveEvent(this, event); };
	void lower() { callbackQDesignerObjectInspectorInterface_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerObjectInspectorInterface_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerObjectInspectorInterface_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerObjectInspectorInterface_PaintEvent(this, event); };
	void raise() { callbackQDesignerObjectInspectorInterface_Raise(this); };
	void repaint() { callbackQDesignerObjectInspectorInterface_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerObjectInspectorInterface_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQDesignerObjectInspectorInterface_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQDesignerObjectInspectorInterface_SetEnabled(this, vbo); };
	void setFocus() { callbackQDesignerObjectInspectorInterface_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQDesignerObjectInspectorInterface_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtDesigner_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQDesignerObjectInspectorInterface_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQDesignerObjectInspectorInterface_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQDesignerObjectInspectorInterface_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtDesigner_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQDesignerObjectInspectorInterface_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQDesignerObjectInspectorInterface_Show(this); };
	void showEvent(QShowEvent * event) { callbackQDesignerObjectInspectorInterface_ShowEvent(this, event); };
	void showFullScreen() { callbackQDesignerObjectInspectorInterface_ShowFullScreen(this); };
	void showMaximized() { callbackQDesignerObjectInspectorInterface_ShowMaximized(this); };
	void showMinimized() { callbackQDesignerObjectInspectorInterface_ShowMinimized(this); };
	void showNormal() { callbackQDesignerObjectInspectorInterface_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerObjectInspectorInterface_TabletEvent(this, event); };
	void update() { callbackQDesignerObjectInspectorInterface_Update(this); };
	void updateMicroFocus() { callbackQDesignerObjectInspectorInterface_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerObjectInspectorInterface_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQDesignerObjectInspectorInterface_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtDesigner_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQDesignerObjectInspectorInterface_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerObjectInspectorInterface_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerObjectInspectorInterface_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerObjectInspectorInterface_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerObjectInspectorInterface_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQDesignerObjectInspectorInterface_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerObjectInspectorInterface_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerObjectInspectorInterface_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerObjectInspectorInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerObjectInspectorInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerObjectInspectorInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerObjectInspectorInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerObjectInspectorInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerObjectInspectorInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerObjectInspectorInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerObjectInspectorInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerObjectInspectorInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerObjectInspectorInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerObjectInspectorInterface*)

int QDesignerObjectInspectorInterface_QDesignerObjectInspectorInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerObjectInspectorInterface*>(); return qRegisterMetaType<MyQDesignerObjectInspectorInterface*>();}

void* QDesignerObjectInspectorInterface_NewQDesignerObjectInspectorInterface(void* parent, long long flags)
{
		return new MyQDesignerObjectInspectorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QDesignerObjectInspectorInterface_SetFormWindow(void* ptr, void* formWindow)
{
	QMetaObject::invokeMethod(static_cast<QDesignerObjectInspectorInterface*>(ptr), "setFormWindow", Q_ARG(QDesignerFormWindowInterface*, static_cast<QDesignerFormWindowInterface*>(formWindow)));
}

void QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(void* ptr)
{
	static_cast<QDesignerObjectInspectorInterface*>(ptr)->~QDesignerObjectInspectorInterface();
}

void QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerObjectInspectorInterface_Core(void* ptr)
{
	return static_cast<QDesignerObjectInspectorInterface*>(ptr)->core();
}

void* QDesignerObjectInspectorInterface_CoreDefault(void* ptr)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::core();
}

void* QDesignerObjectInspectorInterface___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerObjectInspectorInterface___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerObjectInspectorInterface___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerObjectInspectorInterface___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerObjectInspectorInterface___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerObjectInspectorInterface___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerObjectInspectorInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerObjectInspectorInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerObjectInspectorInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerObjectInspectorInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerObjectInspectorInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerObjectInspectorInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerObjectInspectorInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerObjectInspectorInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerObjectInspectorInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerObjectInspectorInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerObjectInspectorInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerObjectInspectorInterface_CloseDefault(void* ptr)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::close();
}

char QDesignerObjectInspectorInterface_EventDefault(void* ptr, void* event)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::event(static_cast<QEvent*>(event));
}

char QDesignerObjectInspectorInterface_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusNextPrevChild(next != 0);
}

char QDesignerObjectInspectorInterface_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QDesignerObjectInspectorInterface_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerObjectInspectorInterface_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::changeEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerObjectInspectorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerObjectInspectorInterface_DropEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerObjectInspectorInterface_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerObjectInspectorInterface_HideDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hide();
}

void QDesignerObjectInspectorInterface_HideEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerObjectInspectorInterface_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerObjectInspectorInterface_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerObjectInspectorInterface_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_LowerDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::lower();
}

void QDesignerObjectInspectorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerObjectInspectorInterface_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerObjectInspectorInterface_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerObjectInspectorInterface_RaiseDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::raise();
}

void QDesignerObjectInspectorInterface_RepaintDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::repaint();
}

void QDesignerObjectInspectorInterface_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerObjectInspectorInterface_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setDisabled(disable != 0);
}

void QDesignerObjectInspectorInterface_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setEnabled(vbo != 0);
}

void QDesignerObjectInspectorInterface_SetFocus2Default(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setFocus();
}

void QDesignerObjectInspectorInterface_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setHidden(hidden != 0);
}

void QDesignerObjectInspectorInterface_SetStyleSheetDefault(void* ptr, struct QtDesigner_PackedString styleSheet)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QDesignerObjectInspectorInterface_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setVisible(visible != 0);
}

void QDesignerObjectInspectorInterface_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setWindowModified(vbo != 0);
}

void QDesignerObjectInspectorInterface_SetWindowTitleDefault(void* ptr, struct QtDesigner_PackedString vqs)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QDesignerObjectInspectorInterface_ShowDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::show();
}

void QDesignerObjectInspectorInterface_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerObjectInspectorInterface_ShowFullScreenDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showFullScreen();
}

void QDesignerObjectInspectorInterface_ShowMaximizedDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showMaximized();
}

void QDesignerObjectInspectorInterface_ShowMinimizedDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showMinimized();
}

void QDesignerObjectInspectorInterface_ShowNormalDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::showNormal();
}

void QDesignerObjectInspectorInterface_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerObjectInspectorInterface_UpdateDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::update();
}

void QDesignerObjectInspectorInterface_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::updateMicroFocus();
}

void QDesignerObjectInspectorInterface_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QDesignerObjectInspectorInterface_PaintEngineDefault(void* ptr)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::paintEngine();
}

void* QDesignerObjectInspectorInterface_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerObjectInspectorInterface_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerObjectInspectorInterface_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QDesignerObjectInspectorInterface_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::hasHeightForWidth();
}

int QDesignerObjectInspectorInterface_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::heightForWidth(w);
}

int QDesignerObjectInspectorInterface_MetricDefault(void* ptr, long long m)
{
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QDesignerObjectInspectorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerObjectInspectorInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerObjectInspectorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerObjectInspectorInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerObjectInspectorInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::deleteLater();
}

void QDesignerObjectInspectorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerObjectInspectorInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerObjectInspectorInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerObjectInspectorInterface*>(ptr)->QDesignerObjectInspectorInterface::metaObject());
}

class MyQDesignerPropertyEditorInterface: public QDesignerPropertyEditorInterface
{
public:
	MyQDesignerPropertyEditorInterface(QWidget *parent, Qt::WindowFlags flags = Qt::WindowFlags()) : QDesignerPropertyEditorInterface(parent, flags) {QDesignerPropertyEditorInterface_QDesignerPropertyEditorInterface_QRegisterMetaType();};
	void Signal_PropertyChanged(const QString & name, const QVariant & value) { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDesignerPropertyEditorInterface_PropertyChanged(this, namePacked, const_cast<QVariant*>(&value)); };
	void setObject(QObject * object) { callbackQDesignerPropertyEditorInterface_SetObject(this, object); };
	void setPropertyValue(const QString & name, const QVariant & value, bool changed) { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQDesignerPropertyEditorInterface_SetPropertyValue(this, namePacked, const_cast<QVariant*>(&value), changed); };
	void setReadOnly(bool readOnly) { callbackQDesignerPropertyEditorInterface_SetReadOnly(this, readOnly); };
	 ~MyQDesignerPropertyEditorInterface() { callbackQDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(this); };
	QDesignerFormEditorInterface * core() const { return static_cast<QDesignerFormEditorInterface*>(callbackQDesignerPropertyEditorInterface_Core(const_cast<void*>(static_cast<const void*>(this)))); };
	QObject * object() const { return static_cast<QObject*>(callbackQDesignerPropertyEditorInterface_Object(const_cast<void*>(static_cast<const void*>(this)))); };
	QString currentPropertyName() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerPropertyEditorInterface_CurrentPropertyName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool isReadOnly() const { return callbackQDesignerPropertyEditorInterface_IsReadOnly(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool close() { return callbackQDesignerPropertyEditorInterface_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQDesignerPropertyEditorInterface_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerPropertyEditorInterface_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerPropertyEditorInterface_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQDesignerPropertyEditorInterface_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQDesignerPropertyEditorInterface_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerPropertyEditorInterface_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQDesignerPropertyEditorInterface_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerPropertyEditorInterface_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerPropertyEditorInterface_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerPropertyEditorInterface_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerPropertyEditorInterface_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerPropertyEditorInterface_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerPropertyEditorInterface_FocusOutEvent(this, event); };
	void hide() { callbackQDesignerPropertyEditorInterface_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQDesignerPropertyEditorInterface_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerPropertyEditorInterface_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerPropertyEditorInterface_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerPropertyEditorInterface_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_LeaveEvent(this, event); };
	void lower() { callbackQDesignerPropertyEditorInterface_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerPropertyEditorInterface_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerPropertyEditorInterface_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerPropertyEditorInterface_PaintEvent(this, event); };
	void raise() { callbackQDesignerPropertyEditorInterface_Raise(this); };
	void repaint() { callbackQDesignerPropertyEditorInterface_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerPropertyEditorInterface_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQDesignerPropertyEditorInterface_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQDesignerPropertyEditorInterface_SetEnabled(this, vbo); };
	void setFocus() { callbackQDesignerPropertyEditorInterface_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQDesignerPropertyEditorInterface_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtDesigner_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQDesignerPropertyEditorInterface_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQDesignerPropertyEditorInterface_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQDesignerPropertyEditorInterface_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtDesigner_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQDesignerPropertyEditorInterface_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQDesignerPropertyEditorInterface_Show(this); };
	void showEvent(QShowEvent * event) { callbackQDesignerPropertyEditorInterface_ShowEvent(this, event); };
	void showFullScreen() { callbackQDesignerPropertyEditorInterface_ShowFullScreen(this); };
	void showMaximized() { callbackQDesignerPropertyEditorInterface_ShowMaximized(this); };
	void showMinimized() { callbackQDesignerPropertyEditorInterface_ShowMinimized(this); };
	void showNormal() { callbackQDesignerPropertyEditorInterface_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerPropertyEditorInterface_TabletEvent(this, event); };
	void update() { callbackQDesignerPropertyEditorInterface_Update(this); };
	void updateMicroFocus() { callbackQDesignerPropertyEditorInterface_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerPropertyEditorInterface_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQDesignerPropertyEditorInterface_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtDesigner_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQDesignerPropertyEditorInterface_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerPropertyEditorInterface_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerPropertyEditorInterface_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerPropertyEditorInterface_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerPropertyEditorInterface_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQDesignerPropertyEditorInterface_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerPropertyEditorInterface_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerPropertyEditorInterface_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerPropertyEditorInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerPropertyEditorInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerPropertyEditorInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerPropertyEditorInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerPropertyEditorInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerPropertyEditorInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerPropertyEditorInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerPropertyEditorInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerPropertyEditorInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerPropertyEditorInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerPropertyEditorInterface*)

int QDesignerPropertyEditorInterface_QDesignerPropertyEditorInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerPropertyEditorInterface*>(); return qRegisterMetaType<MyQDesignerPropertyEditorInterface*>();}

void* QDesignerPropertyEditorInterface_NewQDesignerPropertyEditorInterface(void* parent, long long flags)
{
		return new MyQDesignerPropertyEditorInterface(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QDesignerPropertyEditorInterface_ConnectPropertyChanged(void* ptr)
{
	QObject::connect(static_cast<QDesignerPropertyEditorInterface*>(ptr), static_cast<void (QDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&QDesignerPropertyEditorInterface::propertyChanged), static_cast<MyQDesignerPropertyEditorInterface*>(ptr), static_cast<void (MyQDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&MyQDesignerPropertyEditorInterface::Signal_PropertyChanged));
}

void QDesignerPropertyEditorInterface_DisconnectPropertyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDesignerPropertyEditorInterface*>(ptr), static_cast<void (QDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&QDesignerPropertyEditorInterface::propertyChanged), static_cast<MyQDesignerPropertyEditorInterface*>(ptr), static_cast<void (MyQDesignerPropertyEditorInterface::*)(const QString &, const QVariant &)>(&MyQDesignerPropertyEditorInterface::Signal_PropertyChanged));
}

void QDesignerPropertyEditorInterface_PropertyChanged(void* ptr, struct QtDesigner_PackedString name, void* value)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->propertyChanged(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value));
}

void QDesignerPropertyEditorInterface_SetObject(void* ptr, void* object)
{
		QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setObject", Q_ARG(QObject*, static_cast<QObject*>(object)));
}

void QDesignerPropertyEditorInterface_SetPropertyValue(void* ptr, struct QtDesigner_PackedString name, void* value, char changed)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setPropertyValue", Q_ARG(const QString, QString::fromUtf8(name.data, name.len)), Q_ARG(const QVariant, *static_cast<QVariant*>(value)), Q_ARG(bool, changed != 0));
}

void QDesignerPropertyEditorInterface_SetReadOnly(void* ptr, char readOnly)
{
	QMetaObject::invokeMethod(static_cast<QDesignerPropertyEditorInterface*>(ptr), "setReadOnly", Q_ARG(bool, readOnly != 0));
}

void QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(void* ptr)
{
	static_cast<QDesignerPropertyEditorInterface*>(ptr)->~QDesignerPropertyEditorInterface();
}

void QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerPropertyEditorInterface_Core(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->core();
}

void* QDesignerPropertyEditorInterface_CoreDefault(void* ptr)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::core();
}

void* QDesignerPropertyEditorInterface_Object(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->object();
}

struct QtDesigner_PackedString QDesignerPropertyEditorInterface_CurrentPropertyName(void* ptr)
{
	return ({ QByteArray tc6cc46 = static_cast<QDesignerPropertyEditorInterface*>(ptr)->currentPropertyName().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tc6cc46.prepend("WHITESPACE").constData()+10), tc6cc46.size()-10 }; });
}

char QDesignerPropertyEditorInterface_IsReadOnly(void* ptr)
{
	return static_cast<QDesignerPropertyEditorInterface*>(ptr)->isReadOnly();
}

void* QDesignerPropertyEditorInterface___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerPropertyEditorInterface___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerPropertyEditorInterface___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerPropertyEditorInterface___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerPropertyEditorInterface___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerPropertyEditorInterface___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerPropertyEditorInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerPropertyEditorInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerPropertyEditorInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerPropertyEditorInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerPropertyEditorInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerPropertyEditorInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerPropertyEditorInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerPropertyEditorInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerPropertyEditorInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerPropertyEditorInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerPropertyEditorInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerPropertyEditorInterface_CloseDefault(void* ptr)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::close();
}

char QDesignerPropertyEditorInterface_EventDefault(void* ptr, void* event)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::event(static_cast<QEvent*>(event));
}

char QDesignerPropertyEditorInterface_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusNextPrevChild(next != 0);
}

char QDesignerPropertyEditorInterface_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QDesignerPropertyEditorInterface_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerPropertyEditorInterface_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::changeEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerPropertyEditorInterface_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerPropertyEditorInterface_DropEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerPropertyEditorInterface_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerPropertyEditorInterface_HideDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hide();
}

void QDesignerPropertyEditorInterface_HideEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerPropertyEditorInterface_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerPropertyEditorInterface_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerPropertyEditorInterface_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_LowerDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::lower();
}

void QDesignerPropertyEditorInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerPropertyEditorInterface_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerPropertyEditorInterface_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerPropertyEditorInterface_RaiseDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::raise();
}

void QDesignerPropertyEditorInterface_RepaintDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::repaint();
}

void QDesignerPropertyEditorInterface_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerPropertyEditorInterface_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setDisabled(disable != 0);
}

void QDesignerPropertyEditorInterface_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setEnabled(vbo != 0);
}

void QDesignerPropertyEditorInterface_SetFocus2Default(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setFocus();
}

void QDesignerPropertyEditorInterface_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setHidden(hidden != 0);
}

void QDesignerPropertyEditorInterface_SetStyleSheetDefault(void* ptr, struct QtDesigner_PackedString styleSheet)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QDesignerPropertyEditorInterface_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setVisible(visible != 0);
}

void QDesignerPropertyEditorInterface_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setWindowModified(vbo != 0);
}

void QDesignerPropertyEditorInterface_SetWindowTitleDefault(void* ptr, struct QtDesigner_PackedString vqs)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QDesignerPropertyEditorInterface_ShowDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::show();
}

void QDesignerPropertyEditorInterface_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerPropertyEditorInterface_ShowFullScreenDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showFullScreen();
}

void QDesignerPropertyEditorInterface_ShowMaximizedDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showMaximized();
}

void QDesignerPropertyEditorInterface_ShowMinimizedDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showMinimized();
}

void QDesignerPropertyEditorInterface_ShowNormalDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::showNormal();
}

void QDesignerPropertyEditorInterface_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerPropertyEditorInterface_UpdateDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::update();
}

void QDesignerPropertyEditorInterface_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::updateMicroFocus();
}

void QDesignerPropertyEditorInterface_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QDesignerPropertyEditorInterface_PaintEngineDefault(void* ptr)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::paintEngine();
}

void* QDesignerPropertyEditorInterface_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerPropertyEditorInterface_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerPropertyEditorInterface_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QDesignerPropertyEditorInterface_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::hasHeightForWidth();
}

int QDesignerPropertyEditorInterface_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::heightForWidth(w);
}

int QDesignerPropertyEditorInterface_MetricDefault(void* ptr, long long m)
{
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QDesignerPropertyEditorInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerPropertyEditorInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerPropertyEditorInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerPropertyEditorInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerPropertyEditorInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::deleteLater();
}

void QDesignerPropertyEditorInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerPropertyEditorInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerPropertyEditorInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerPropertyEditorInterface*>(ptr)->QDesignerPropertyEditorInterface::metaObject());
}

class MyQDesignerPropertySheetExtension: public QDesignerPropertySheetExtension
{
public:
	bool reset(int index) { return callbackQDesignerPropertySheetExtension_Reset(this, index) != 0; };
	void setAttribute(int index, bool attribute) { callbackQDesignerPropertySheetExtension_SetAttribute(this, index, attribute); };
	void setChanged(int index, bool changed) { callbackQDesignerPropertySheetExtension_SetChanged(this, index, changed); };
	void setProperty(int index, const QVariant & value) { callbackQDesignerPropertySheetExtension_SetProperty(this, index, const_cast<QVariant*>(&value)); };
	void setPropertyGroup(int index, const QString & group) { QByteArray t64292b = group.toUtf8(); QtDesigner_PackedString groupPacked = { const_cast<char*>(t64292b.prepend("WHITESPACE").constData()+10), t64292b.size()-10 };callbackQDesignerPropertySheetExtension_SetPropertyGroup(this, index, groupPacked); };
	void setVisible(int index, bool visible) { callbackQDesignerPropertySheetExtension_SetVisible(this, index, visible); };
	 ~MyQDesignerPropertySheetExtension() { callbackQDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(this); };
	QString propertyGroup(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerPropertySheetExtension_PropertyGroup(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString propertyName(int index) const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerPropertySheetExtension_PropertyName(const_cast<void*>(static_cast<const void*>(this)), index); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QVariant property(int index) const { return *static_cast<QVariant*>(callbackQDesignerPropertySheetExtension_Property(const_cast<void*>(static_cast<const void*>(this)), index)); };
	bool hasReset(int index) const { return callbackQDesignerPropertySheetExtension_HasReset(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isAttribute(int index) const { return callbackQDesignerPropertySheetExtension_IsAttribute(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isChanged(int index) const { return callbackQDesignerPropertySheetExtension_IsChanged(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isEnabled(int index) const { return callbackQDesignerPropertySheetExtension_IsEnabled(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	bool isVisible(int index) const { return callbackQDesignerPropertySheetExtension_IsVisible(const_cast<void*>(static_cast<const void*>(this)), index) != 0; };
	int count() const { return callbackQDesignerPropertySheetExtension_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int indexOf(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtDesigner_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQDesignerPropertySheetExtension_IndexOf(const_cast<void*>(static_cast<const void*>(this)), namePacked); };
};

char QDesignerPropertySheetExtension_Reset(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->reset(index);
}

void QDesignerPropertySheetExtension_SetAttribute(void* ptr, int index, char attribute)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setAttribute(index, attribute != 0);
}

void QDesignerPropertySheetExtension_SetChanged(void* ptr, int index, char changed)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setChanged(index, changed != 0);
}

void QDesignerPropertySheetExtension_SetProperty(void* ptr, int index, void* value)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setProperty(index, *static_cast<QVariant*>(value));
}

void QDesignerPropertySheetExtension_SetPropertyGroup(void* ptr, int index, struct QtDesigner_PackedString group)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setPropertyGroup(index, QString::fromUtf8(group.data, group.len));
}

void QDesignerPropertySheetExtension_SetVisible(void* ptr, int index, char visible)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->setVisible(index, visible != 0);
}

void QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(void* ptr)
{
	static_cast<QDesignerPropertySheetExtension*>(ptr)->~QDesignerPropertySheetExtension();
}

void QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtensionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDesigner_PackedString QDesignerPropertySheetExtension_PropertyGroup(void* ptr, int index)
{
	return ({ QByteArray t9e9d19 = static_cast<QDesignerPropertySheetExtension*>(ptr)->propertyGroup(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(t9e9d19.prepend("WHITESPACE").constData()+10), t9e9d19.size()-10 }; });
}

struct QtDesigner_PackedString QDesignerPropertySheetExtension_PropertyName(void* ptr, int index)
{
	return ({ QByteArray t6cba99 = static_cast<QDesignerPropertySheetExtension*>(ptr)->propertyName(index).toUtf8(); QtDesigner_PackedString { const_cast<char*>(t6cba99.prepend("WHITESPACE").constData()+10), t6cba99.size()-10 }; });
}

void* QDesignerPropertySheetExtension_Property(void* ptr, int index)
{
	return new QVariant(static_cast<QDesignerPropertySheetExtension*>(ptr)->property(index));
}

char QDesignerPropertySheetExtension_HasReset(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->hasReset(index);
}

char QDesignerPropertySheetExtension_IsAttribute(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isAttribute(index);
}

char QDesignerPropertySheetExtension_IsChanged(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isChanged(index);
}

char QDesignerPropertySheetExtension_IsEnabled(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isEnabled(index);
}

char QDesignerPropertySheetExtension_IsEnabledDefault(void* ptr, int index)
{
		return static_cast<QDesignerPropertySheetExtension*>(ptr)->QDesignerPropertySheetExtension::isEnabled(index);
}

char QDesignerPropertySheetExtension_IsVisible(void* ptr, int index)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->isVisible(index);
}

int QDesignerPropertySheetExtension_Count(void* ptr)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->count();
}

int QDesignerPropertySheetExtension_IndexOf(void* ptr, struct QtDesigner_PackedString name)
{
	return static_cast<QDesignerPropertySheetExtension*>(ptr)->indexOf(QString::fromUtf8(name.data, name.len));
}

class MyQDesignerTaskMenuExtension: public QDesignerTaskMenuExtension
{
public:
	 ~MyQDesignerTaskMenuExtension() { callbackQDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(this); };
	QAction * preferredEditAction() const { return static_cast<QAction*>(callbackQDesignerTaskMenuExtension_PreferredEditAction(const_cast<void*>(static_cast<const void*>(this)))); };
	QList<QAction *> taskActions() const { return *static_cast<QList<QAction *>*>(callbackQDesignerTaskMenuExtension_TaskActions(const_cast<void*>(static_cast<const void*>(this)))); };
};

void QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(void* ptr)
{
	static_cast<QDesignerTaskMenuExtension*>(ptr)->~QDesignerTaskMenuExtension();
}

void QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtensionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDesignerTaskMenuExtension_PreferredEditAction(void* ptr)
{
	return static_cast<QDesignerTaskMenuExtension*>(ptr)->preferredEditAction();
}

void* QDesignerTaskMenuExtension_PreferredEditActionDefault(void* ptr)
{
		return static_cast<QDesignerTaskMenuExtension*>(ptr)->QDesignerTaskMenuExtension::preferredEditAction();
}

struct QtDesigner_PackedList QDesignerTaskMenuExtension_TaskActions(void* ptr)
{
	return ({ QList<QAction *>* tmpValue = new QList<QAction *>(static_cast<QDesignerTaskMenuExtension*>(ptr)->taskActions()); QtDesigner_PackedList { tmpValue, tmpValue->size() }; });
}

void* QDesignerTaskMenuExtension___taskActions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerTaskMenuExtension___taskActions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerTaskMenuExtension___taskActions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

class MyQDesignerWidgetBoxInterface: public QDesignerWidgetBoxInterface
{
public:
	bool load() { return callbackQDesignerWidgetBoxInterface_Load(this) != 0; };
	bool save() { return callbackQDesignerWidgetBoxInterface_Save(this) != 0; };
	void setFileName(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtDesigner_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQDesignerWidgetBoxInterface_SetFileName(this, fileNamePacked); };
	 ~MyQDesignerWidgetBoxInterface() { callbackQDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(this); };
	QString fileName() const { return ({ QtDesigner_PackedString tempVal = callbackQDesignerWidgetBoxInterface_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool close() { return callbackQDesignerWidgetBoxInterface_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQDesignerWidgetBoxInterface_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQDesignerWidgetBoxInterface_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQDesignerWidgetBoxInterface_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQDesignerWidgetBoxInterface_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQDesignerWidgetBoxInterface_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQDesignerWidgetBoxInterface_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQDesignerWidgetBoxInterface_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQDesignerWidgetBoxInterface_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQDesignerWidgetBoxInterface_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQDesignerWidgetBoxInterface_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQDesignerWidgetBoxInterface_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQDesignerWidgetBoxInterface_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQDesignerWidgetBoxInterface_FocusOutEvent(this, event); };
	void hide() { callbackQDesignerWidgetBoxInterface_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQDesignerWidgetBoxInterface_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQDesignerWidgetBoxInterface_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQDesignerWidgetBoxInterface_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQDesignerWidgetBoxInterface_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_LeaveEvent(this, event); };
	void lower() { callbackQDesignerWidgetBoxInterface_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQDesignerWidgetBoxInterface_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQDesignerWidgetBoxInterface_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQDesignerWidgetBoxInterface_PaintEvent(this, event); };
	void raise() { callbackQDesignerWidgetBoxInterface_Raise(this); };
	void repaint() { callbackQDesignerWidgetBoxInterface_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQDesignerWidgetBoxInterface_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQDesignerWidgetBoxInterface_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQDesignerWidgetBoxInterface_SetEnabled(this, vbo); };
	void setFocus() { callbackQDesignerWidgetBoxInterface_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQDesignerWidgetBoxInterface_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtDesigner_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQDesignerWidgetBoxInterface_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQDesignerWidgetBoxInterface_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQDesignerWidgetBoxInterface_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtDesigner_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQDesignerWidgetBoxInterface_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQDesignerWidgetBoxInterface_Show(this); };
	void showEvent(QShowEvent * event) { callbackQDesignerWidgetBoxInterface_ShowEvent(this, event); };
	void showFullScreen() { callbackQDesignerWidgetBoxInterface_ShowFullScreen(this); };
	void showMaximized() { callbackQDesignerWidgetBoxInterface_ShowMaximized(this); };
	void showMinimized() { callbackQDesignerWidgetBoxInterface_ShowMinimized(this); };
	void showNormal() { callbackQDesignerWidgetBoxInterface_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQDesignerWidgetBoxInterface_TabletEvent(this, event); };
	void update() { callbackQDesignerWidgetBoxInterface_Update(this); };
	void updateMicroFocus() { callbackQDesignerWidgetBoxInterface_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQDesignerWidgetBoxInterface_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQDesignerWidgetBoxInterface_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtDesigner_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQDesignerWidgetBoxInterface_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQDesignerWidgetBoxInterface_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQDesignerWidgetBoxInterface_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQDesignerWidgetBoxInterface_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQDesignerWidgetBoxInterface_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQDesignerWidgetBoxInterface_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQDesignerWidgetBoxInterface_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQDesignerWidgetBoxInterface_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDesignerWidgetBoxInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDesignerWidgetBoxInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDesignerWidgetBoxInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDesignerWidgetBoxInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDesignerWidgetBoxInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDesignerWidgetBoxInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDesignerWidgetBoxInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDesignerWidgetBoxInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDesignerWidgetBoxInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDesignerWidgetBoxInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDesignerWidgetBoxInterface*)

int QDesignerWidgetBoxInterface_QDesignerWidgetBoxInterface_QRegisterMetaType(){qRegisterMetaType<QDesignerWidgetBoxInterface*>(); return qRegisterMetaType<MyQDesignerWidgetBoxInterface*>();}

char QDesignerWidgetBoxInterface_Load(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->load();
}

char QDesignerWidgetBoxInterface_Save(void* ptr)
{
	return static_cast<QDesignerWidgetBoxInterface*>(ptr)->save();
}

void QDesignerWidgetBoxInterface_SetFileName(void* ptr, struct QtDesigner_PackedString fileName)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(void* ptr)
{
	static_cast<QDesignerWidgetBoxInterface*>(ptr)->~QDesignerWidgetBoxInterface();
}

void QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDesigner_PackedString QDesignerWidgetBoxInterface_FileName(void* ptr)
{
	return ({ QByteArray tda04ff = static_cast<QDesignerWidgetBoxInterface*>(ptr)->fileName().toUtf8(); QtDesigner_PackedString { const_cast<char*>(tda04ff.prepend("WHITESPACE").constData()+10), tda04ff.size()-10 }; });
}

void* QDesignerWidgetBoxInterface___dropWidgets_item_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerDnDItemInterface *>;
}

void* QDesignerWidgetBoxInterface___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerWidgetBoxInterface___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerWidgetBoxInterface___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerWidgetBoxInterface___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerWidgetBoxInterface___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QDesignerWidgetBoxInterface___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QDesignerWidgetBoxInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDesignerWidgetBoxInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QDesignerWidgetBoxInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerWidgetBoxInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerWidgetBoxInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerWidgetBoxInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerWidgetBoxInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerWidgetBoxInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QDesignerWidgetBoxInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDesignerWidgetBoxInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDesignerWidgetBoxInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QDesignerWidgetBoxInterface_CloseDefault(void* ptr)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::close();
}

char QDesignerWidgetBoxInterface_EventDefault(void* ptr, void* event)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::event(static_cast<QEvent*>(event));
}

char QDesignerWidgetBoxInterface_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusNextPrevChild(next != 0);
}

char QDesignerWidgetBoxInterface_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QDesignerWidgetBoxInterface_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::actionEvent(static_cast<QActionEvent*>(event));
}

void QDesignerWidgetBoxInterface_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::changeEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::closeEvent(static_cast<QCloseEvent*>(event));
}

void QDesignerWidgetBoxInterface_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QDesignerWidgetBoxInterface_DropEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::dropEvent(static_cast<QDropEvent*>(event));
}

void QDesignerWidgetBoxInterface_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::enterEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QDesignerWidgetBoxInterface_HideDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hide();
}

void QDesignerWidgetBoxInterface_HideEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hideEvent(static_cast<QHideEvent*>(event));
}

void QDesignerWidgetBoxInterface_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QDesignerWidgetBoxInterface_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QDesignerWidgetBoxInterface_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::leaveEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_LowerDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::lower();
}

void QDesignerWidgetBoxInterface_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QDesignerWidgetBoxInterface_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::moveEvent(static_cast<QMoveEvent*>(event));
}

void QDesignerWidgetBoxInterface_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::paintEvent(static_cast<QPaintEvent*>(event));
}

void QDesignerWidgetBoxInterface_RaiseDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::raise();
}

void QDesignerWidgetBoxInterface_RepaintDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::repaint();
}

void QDesignerWidgetBoxInterface_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QDesignerWidgetBoxInterface_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setDisabled(disable != 0);
}

void QDesignerWidgetBoxInterface_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setEnabled(vbo != 0);
}

void QDesignerWidgetBoxInterface_SetFocus2Default(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setFocus();
}

void QDesignerWidgetBoxInterface_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setHidden(hidden != 0);
}

void QDesignerWidgetBoxInterface_SetStyleSheetDefault(void* ptr, struct QtDesigner_PackedString styleSheet)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QDesignerWidgetBoxInterface_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setVisible(visible != 0);
}

void QDesignerWidgetBoxInterface_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setWindowModified(vbo != 0);
}

void QDesignerWidgetBoxInterface_SetWindowTitleDefault(void* ptr, struct QtDesigner_PackedString vqs)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QDesignerWidgetBoxInterface_ShowDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::show();
}

void QDesignerWidgetBoxInterface_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showEvent(static_cast<QShowEvent*>(event));
}

void QDesignerWidgetBoxInterface_ShowFullScreenDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showFullScreen();
}

void QDesignerWidgetBoxInterface_ShowMaximizedDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showMaximized();
}

void QDesignerWidgetBoxInterface_ShowMinimizedDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showMinimized();
}

void QDesignerWidgetBoxInterface_ShowNormalDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::showNormal();
}

void QDesignerWidgetBoxInterface_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QDesignerWidgetBoxInterface_UpdateDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::update();
}

void QDesignerWidgetBoxInterface_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::updateMicroFocus();
}

void QDesignerWidgetBoxInterface_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QDesignerWidgetBoxInterface_PaintEngineDefault(void* ptr)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::paintEngine();
}

void* QDesignerWidgetBoxInterface_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerWidgetBoxInterface_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QDesignerWidgetBoxInterface_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QDesignerWidgetBoxInterface_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::hasHeightForWidth();
}

int QDesignerWidgetBoxInterface_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::heightForWidth(w);
}

int QDesignerWidgetBoxInterface_MetricDefault(void* ptr, long long m)
{
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QDesignerWidgetBoxInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDesignerWidgetBoxInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDesignerWidgetBoxInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerWidgetBoxInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::customEvent(static_cast<QEvent*>(event));
}

void QDesignerWidgetBoxInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::deleteLater();
}

void QDesignerWidgetBoxInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDesignerWidgetBoxInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDesignerWidgetBoxInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDesignerWidgetBoxInterface*>(ptr)->QDesignerWidgetBoxInterface::metaObject());
}

class MyQExtensionFactory: public QExtensionFactory
{
public:
	MyQExtensionFactory(QExtensionManager *parent = Q_NULLPTR) : QExtensionFactory(parent) {QExtensionFactory_QExtensionFactory_QRegisterMetaType();};
	QObject * createExtension(QObject * object, const QString & iid, QObject * parent) const { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };return static_cast<QObject*>(callbackQExtensionFactory_CreateExtension(const_cast<void*>(static_cast<const void*>(this)), object, iidPacked, parent)); };
	QObject * extension(QObject * object, const QString & iid) const { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };return static_cast<QObject*>(callbackQExtensionFactory_Extension(const_cast<void*>(static_cast<const void*>(this)), object, iidPacked)); };
	bool event(QEvent * e) { return callbackQExtensionFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQExtensionFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQExtensionFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQExtensionFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQExtensionFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQExtensionFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQExtensionFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQExtensionFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQExtensionFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQExtensionFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQExtensionFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQExtensionFactory*)

int QExtensionFactory_QExtensionFactory_QRegisterMetaType(){qRegisterMetaType<QExtensionFactory*>(); return qRegisterMetaType<MyQExtensionFactory*>();}

void* QExtensionFactory_NewQExtensionFactory(void* parent)
{
	return new MyQExtensionFactory(static_cast<QExtensionManager*>(parent));
}

void* QExtensionFactory_ExtensionManager(void* ptr)
{
		return static_cast<QExtensionFactory*>(ptr)->extensionManager();
}

void* QExtensionFactory_CreateExtension(void* ptr, void* object, struct QtDesigner_PackedString iid, void* parent)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->createExtension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->createExtension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->createExtension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	}
}

void* QExtensionFactory_CreateExtensionDefault(void* ptr, void* object, struct QtDesigner_PackedString iid, void* parent)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::createExtension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::createExtension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::createExtension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len), static_cast<QObject*>(parent));
	}
}

void* QExtensionFactory_Extension(void* ptr, void* object, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
	}
}

void* QExtensionFactory_ExtensionDefault(void* ptr, void* object, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
	}
}

void* QExtensionFactory___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QExtensionFactory___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QExtensionFactory___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>;
}

void* QExtensionFactory___findChildren_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionFactory___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionFactory___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionFactory___findChildren_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionFactory___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionFactory___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionFactory___findChildren_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionFactory___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionFactory___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionFactory___children_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QExtensionFactory___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionFactory___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>;
}

char QExtensionFactory_Event(void* ptr, void* e)
{
		return static_cast<QExtensionFactory*>(ptr)->event(static_cast<QEvent*>(e));
}

char QExtensionFactory_EventDefault(void* ptr, void* e)
{
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::event(static_cast<QEvent*>(e));
}

char QExtensionFactory_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionFactory*>(ptr)->eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionFactory*>(ptr)->eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QExtensionFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
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

void QExtensionFactory_TimerEvent(void* ptr, void* event)
{
		static_cast<QExtensionFactory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionFactory_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QExtensionFactory*>(ptr)->QExtensionFactory::timerEvent(static_cast<QTimerEvent*>(event));
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
	MyQExtensionManager(QObject *parent = Q_NULLPTR) : QExtensionManager(parent) {QExtensionManager_QExtensionManager_QRegisterMetaType();};
	void registerExtensions(QAbstractExtensionFactory * factory, const QString & iid) { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };callbackQExtensionManager_RegisterExtensions(this, factory, iidPacked); };
	void unregisterExtensions(QAbstractExtensionFactory * factory, const QString & iid) { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };callbackQExtensionManager_UnregisterExtensions(this, factory, iidPacked); };
	QObject * extension(QObject * object, const QString & iid) const { QByteArray t52efa9 = iid.toUtf8(); QtDesigner_PackedString iidPacked = { const_cast<char*>(t52efa9.prepend("WHITESPACE").constData()+10), t52efa9.size()-10 };return static_cast<QObject*>(callbackQExtensionManager_Extension(const_cast<void*>(static_cast<const void*>(this)), object, iidPacked)); };
	bool event(QEvent * e) { return callbackQExtensionManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQExtensionManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQExtensionManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQExtensionManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQExtensionManager_CustomEvent(this, event); };
	void deleteLater() { callbackQExtensionManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQExtensionManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQExtensionManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDesigner_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQExtensionManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQExtensionManager_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQExtensionManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQExtensionManager*)

int QExtensionManager_QExtensionManager_QRegisterMetaType(){qRegisterMetaType<QExtensionManager*>(); return qRegisterMetaType<MyQExtensionManager*>();}

void* QExtensionManager_NewQExtensionManager(void* parent)
{
	return new MyQExtensionManager(static_cast<QObject*>(parent));
}

void QExtensionManager_RegisterExtensions(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
		static_cast<QExtensionManager*>(ptr)->registerExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	} else {
		static_cast<QExtensionManager*>(ptr)->registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	}
}

void QExtensionManager_RegisterExtensionsDefault(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
		static_cast<QExtensionManager*>(ptr)->QExtensionManager::registerExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	} else {
		static_cast<QExtensionManager*>(ptr)->QExtensionManager::registerExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	}
}

void QExtensionManager_UnregisterExtensions(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
		static_cast<QExtensionManager*>(ptr)->unregisterExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	} else {
		static_cast<QExtensionManager*>(ptr)->unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	}
}

void QExtensionManager_UnregisterExtensionsDefault(void* ptr, void* factory, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(factory))) {
		static_cast<QExtensionManager*>(ptr)->QExtensionManager::unregisterExtensions(static_cast<QExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	} else {
		static_cast<QExtensionManager*>(ptr)->QExtensionManager::unregisterExtensions(static_cast<QAbstractExtensionFactory*>(factory), QString::fromUtf8(iid.data, iid.len));
	}
}

void QExtensionManager_DestroyQExtensionManager(void* ptr)
{
	static_cast<QExtensionManager*>(ptr)->~QExtensionManager();
}

void* QExtensionManager_Extension(void* ptr, void* object, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
	} else {
		return static_cast<QExtensionManager*>(ptr)->extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
	}
}

void* QExtensionManager_ExtensionDefault(void* ptr, void* object, struct QtDesigner_PackedString iid)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::extension(static_cast<QExtensionFactory*>(object), QString::fromUtf8(iid.data, iid.len));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::extension(static_cast<QExtensionManager*>(object), QString::fromUtf8(iid.data, iid.len));
	} else {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::extension(static_cast<QObject*>(object), QString::fromUtf8(iid.data, iid.len));
	}
}

void* QExtensionManager___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QExtensionManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QExtensionManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>;
}

void* QExtensionManager___findChildren_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionManager___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionManager___findChildren_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionManager___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionManager___findChildren_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QExtensionManager___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QExtensionManager___children_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QExtensionManager___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionFactory*>(i));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QExtensionManager*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QExtensionManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>;
}

char QExtensionManager_Event(void* ptr, void* e)
{
		return static_cast<QExtensionManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QExtensionManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::event(static_cast<QEvent*>(e));
}

char QExtensionManager_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionManager*>(ptr)->eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionManager*>(ptr)->eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QExtensionManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QExtensionManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::eventFilter(static_cast<QExtensionFactory*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(watched))) {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::eventFilter(static_cast<QExtensionManager*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QExtensionManager*>(ptr)->QExtensionManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
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

void QExtensionManager_TimerEvent(void* ptr, void* event)
{
		static_cast<QExtensionManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QExtensionManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QExtensionManager*>(ptr)->QExtensionManager::timerEvent(static_cast<QTimerEvent*>(event));
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
	MyQFormBuilder() : QFormBuilder() {};
	 ~MyQFormBuilder() { callbackQFormBuilder_DestroyQFormBuilder(this); };
	QWidget * load(QIODevice * device, QWidget * parent) { return static_cast<QWidget*>(callbackQAbstractFormBuilder_Load(this, device, parent)); };
	void save(QIODevice * device, QWidget * widget) { callbackQAbstractFormBuilder_Save(this, device, widget); };
};

void* QFormBuilder_NewQFormBuilder()
{
	return new MyQFormBuilder();
}

void QFormBuilder_AddPluginPath(void* ptr, struct QtDesigner_PackedString pluginPath)
{
	static_cast<QFormBuilder*>(ptr)->addPluginPath(QString::fromUtf8(pluginPath.data, pluginPath.len));
}

void QFormBuilder_ClearPluginPaths(void* ptr)
{
	static_cast<QFormBuilder*>(ptr)->clearPluginPaths();
}

void QFormBuilder_SetPluginPath(void* ptr, struct QtDesigner_PackedString pluginPaths)
{
	static_cast<QFormBuilder*>(ptr)->setPluginPath(QString::fromUtf8(pluginPaths.data, pluginPaths.len).split("|", QString::SkipEmptyParts));
}

void QFormBuilder_DestroyQFormBuilder(void* ptr)
{
	static_cast<QFormBuilder*>(ptr)->~QFormBuilder();
}

void QFormBuilder_DestroyQFormBuilderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDesigner_PackedList QFormBuilder_CustomWidgets(void* ptr)
{
	return ({ QList<QDesignerCustomWidgetInterface *>* tmpValue = new QList<QDesignerCustomWidgetInterface *>(static_cast<QFormBuilder*>(ptr)->customWidgets()); QtDesigner_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDesigner_PackedString QFormBuilder_PluginPaths(void* ptr)
{
	return ({ QByteArray t18141d = static_cast<QFormBuilder*>(ptr)->pluginPaths().join("|").toUtf8(); QtDesigner_PackedString { const_cast<char*>(t18141d.prepend("WHITESPACE").constData()+10), t18141d.size()-10 }; });
}

void* QFormBuilder___customWidgets_atList(void* ptr, int i)
{
	return const_cast<QDesignerCustomWidgetInterface*>(static_cast<QList<QDesignerCustomWidgetInterface *>*>(ptr)->at(i));
}

void QFormBuilder___customWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDesignerCustomWidgetInterface *>*>(ptr)->append(static_cast<QDesignerCustomWidgetInterface*>(i));
}

void* QFormBuilder___customWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDesignerCustomWidgetInterface *>;
}

