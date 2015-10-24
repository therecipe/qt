#include "qwidget.h"
#include <QByteArray>
#include <QObject>
#include <QPalette>
#include <QCursor>
#include <QModelIndex>
#include <QKeySequence>
#include <QLayout>
#include <QPainter>
#include <QMargins>
#include <QAction>
#include <QFont>
#include <QMetaObject>
#include <QString>
#include <QRegion>
#include <QRect>
#include <QSize>
#include <QBitmap>
#include <QSizePolicy>
#include <QPoint>
#include <QGraphicsEffect>
#include <QPaintDevice>
#include <QUrl>
#include <QStyle>
#include <QVariant>
#include <QLocale>
#include <QIcon>
#include <QWindow>
#include <QWidget>
#include "_cgo_export.h"

class MyQWidget: public QWidget {
public:
void Signal_WindowIconTextChanged(const QString & iconText){callbackQWidgetWindowIconTextChanged(this->objectName().toUtf8().data(), iconText.toUtf8().data());};
void Signal_WindowTitleChanged(const QString & title){callbackQWidgetWindowTitleChanged(this->objectName().toUtf8().data(), title.toUtf8().data());};
};

int QWidget_AcceptDrops(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->acceptDrops();
}

char* QWidget_AccessibleDescription(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QWidget_AccessibleName(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->accessibleName().toUtf8().data();
}

void QWidget_ActivateWindow(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->activateWindow();
}

int QWidget_AutoFillBackground(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->autoFillBackground();
}

void QWidget_ClearMask(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->clearMask();
}

int QWidget_ContextMenuPolicy(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->contextMenuPolicy();
}

int QWidget_FocusPolicy(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->focusPolicy();
}

void QWidget_GrabKeyboard(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->grabKeyboard();
}

void QWidget_GrabMouse(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->grabMouse();
}

void QWidget_GrabMouse2(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QWidget*>(ptr)->grabMouse(*static_cast<QCursor*>(cursor));
}

int QWidget_HasFocus(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->hasFocus();
}

int QWidget_InputMethodHints(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->inputMethodHints();
}

int QWidget_IsActiveWindow(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isActiveWindow();
}

int QWidget_IsFullScreen(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isFullScreen();
}

int QWidget_IsMaximized(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isMaximized();
}

int QWidget_IsMinimized(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isMinimized();
}

int QWidget_IsWindowModified(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isWindowModified();
}

QtObjectPtr QWidget_QWidget_KeyboardGrabber(){
	return QWidget::keyboardGrabber();
}

int QWidget_LayoutDirection(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->layoutDirection();
}

QtObjectPtr QWidget_QWidget_MouseGrabber(){
	return QWidget::mouseGrabber();
}

void QWidget_Move(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->move(*static_cast<QPoint*>(v));
}

QtObjectPtr QWidget_PaintEngine(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->paintEngine();
}

void QWidget_ReleaseKeyboard(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->releaseKeyboard();
}

void QWidget_ReleaseMouse(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->releaseMouse();
}

void QWidget_Resize(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->resize(*static_cast<QSize*>(v));
}

void QWidget_SetAcceptDrops(QtObjectPtr ptr, int on){
	static_cast<QWidget*>(ptr)->setAcceptDrops(on != 0);
}

void QWidget_SetAccessibleDescription(QtObjectPtr ptr, char* description){
	static_cast<QWidget*>(ptr)->setAccessibleDescription(QString(description));
}

void QWidget_SetAccessibleName(QtObjectPtr ptr, char* name){
	static_cast<QWidget*>(ptr)->setAccessibleName(QString(name));
}

void QWidget_SetAutoFillBackground(QtObjectPtr ptr, int enabled){
	static_cast<QWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QWidget_SetContextMenuPolicy(QtObjectPtr ptr, int policy){
	static_cast<QWidget*>(ptr)->setContextMenuPolicy(static_cast<Qt::ContextMenuPolicy>(policy));
}

void QWidget_SetCursor(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setCursor(*static_cast<QCursor*>(v));
}

void QWidget_SetEnabled(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QWidget_SetFixedSize2(QtObjectPtr ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setFixedSize(w, h);
}

void QWidget_SetFocusPolicy(QtObjectPtr ptr, int policy){
	static_cast<QWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QWidget_SetFont(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setFont(*static_cast<QFont*>(v));
}

void QWidget_SetGeometry(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setGeometry(*static_cast<QRect*>(v));
}

void QWidget_SetInputMethodHints(QtObjectPtr ptr, int hints){
	static_cast<QWidget*>(ptr)->setInputMethodHints(static_cast<Qt::InputMethodHint>(hints));
}

void QWidget_SetLayout(QtObjectPtr ptr, QtObjectPtr layout){
	static_cast<QWidget*>(ptr)->setLayout(static_cast<QLayout*>(layout));
}

void QWidget_SetLayoutDirection(QtObjectPtr ptr, int direction){
	static_cast<QWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QWidget_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QWidget*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QWidget_SetMask(QtObjectPtr ptr, QtObjectPtr bitmap){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QBitmap*>(bitmap));
}

void QWidget_SetMask2(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWidget_SetMaximumHeight(QtObjectPtr ptr, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumHeight(maxh);
}

void QWidget_SetMaximumWidth(QtObjectPtr ptr, int maxw){
	static_cast<QWidget*>(ptr)->setMaximumWidth(maxw);
}

void QWidget_SetMinimumHeight(QtObjectPtr ptr, int minh){
	static_cast<QWidget*>(ptr)->setMinimumHeight(minh);
}

void QWidget_SetMinimumWidth(QtObjectPtr ptr, int minw){
	static_cast<QWidget*>(ptr)->setMinimumWidth(minw);
}

void QWidget_SetPalette(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setPalette(*static_cast<QPalette*>(v));
}

void QWidget_SetSizePolicy(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(v));
}

void QWidget_SetStatusTip(QtObjectPtr ptr, char* v){
	static_cast<QWidget*>(ptr)->setStatusTip(QString(v));
}

void QWidget_SetStyleSheet(QtObjectPtr ptr, char* styleSheet){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QWidget_SetToolTip(QtObjectPtr ptr, char* v){
	static_cast<QWidget*>(ptr)->setToolTip(QString(v));
}

void QWidget_SetToolTipDuration(QtObjectPtr ptr, int msec){
	static_cast<QWidget*>(ptr)->setToolTipDuration(msec);
}

void QWidget_SetUpdatesEnabled(QtObjectPtr ptr, int enable){
	static_cast<QWidget*>(ptr)->setUpdatesEnabled(enable != 0);
}

void QWidget_SetVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWidget_SetWhatsThis(QtObjectPtr ptr, char* v){
	static_cast<QWidget*>(ptr)->setWhatsThis(QString(v));
}

void QWidget_SetWindowFilePath(QtObjectPtr ptr, char* filePath){
	static_cast<QWidget*>(ptr)->setWindowFilePath(QString(filePath));
}

void QWidget_SetWindowFlags(QtObjectPtr ptr, int ty){
	static_cast<QWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(ty));
}

void QWidget_SetWindowIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QWidget*>(ptr)->setWindowIcon(*static_cast<QIcon*>(icon));
}

void QWidget_SetWindowIconText(QtObjectPtr ptr, char* v){
	static_cast<QWidget*>(ptr)->setWindowIconText(QString(v));
}

void QWidget_SetWindowModality(QtObjectPtr ptr, int windowModality){
	static_cast<QWidget*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QWidget_SetWindowModified(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowModified", Q_ARG(bool, v != 0));
}

void QWidget_SetWindowState(QtObjectPtr ptr, int windowState){
	static_cast<QWidget*>(ptr)->setWindowState(static_cast<Qt::WindowState>(windowState));
}

void QWidget_SetWindowTitle(QtObjectPtr ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(v)));
}

char* QWidget_StatusTip(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->statusTip().toUtf8().data();
}

char* QWidget_StyleSheet(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->styleSheet().toUtf8().data();
}

char* QWidget_ToolTip(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->toolTip().toUtf8().data();
}

int QWidget_ToolTipDuration(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->toolTipDuration();
}

void QWidget_UnsetCursor(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->unsetCursor();
}

void QWidget_UnsetLayoutDirection(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->unsetLayoutDirection();
}

void QWidget_UnsetLocale(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->unsetLocale();
}

char* QWidget_WhatsThis(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->whatsThis().toUtf8().data();
}

char* QWidget_WindowFilePath(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowFilePath().toUtf8().data();
}

char* QWidget_WindowIconText(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowIconText().toUtf8().data();
}

int QWidget_WindowModality(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowModality();
}

char* QWidget_WindowTitle(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowTitle().toUtf8().data();
}

int QWidget_X(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->x();
}

int QWidget_Y(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->y();
}

QtObjectPtr QWidget_NewQWidget(QtObjectPtr parent, int f){
	return new QWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_AddAction(QtObjectPtr ptr, QtObjectPtr action){
	static_cast<QWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QWidget_AdjustSize(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->adjustSize();
}

int QWidget_BackgroundRole(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->backgroundRole();
}

QtObjectPtr QWidget_BackingStore(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->backingStore();
}

QtObjectPtr QWidget_ChildAt2(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QWidget*>(ptr)->childAt(*static_cast<QPoint*>(p));
}

QtObjectPtr QWidget_ChildAt(QtObjectPtr ptr, int x, int y){
	return static_cast<QWidget*>(ptr)->childAt(x, y);
}

void QWidget_ClearFocus(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->clearFocus();
}

int QWidget_Close(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "close");
}

void QWidget_EnsurePolished(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->ensurePolished();
}

QtObjectPtr QWidget_FocusProxy(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->focusProxy();
}

QtObjectPtr QWidget_FocusWidget(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->focusWidget();
}

int QWidget_ForegroundRole(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->foregroundRole();
}

void QWidget_GetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

void QWidget_GrabGesture(QtObjectPtr ptr, int gesture, int flags){
	static_cast<QWidget*>(ptr)->grabGesture(static_cast<Qt::GestureType>(gesture), static_cast<Qt::GestureFlag>(flags));
}

int QWidget_GrabShortcut(QtObjectPtr ptr, QtObjectPtr key, int context){
	return static_cast<QWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(key), static_cast<Qt::ShortcutContext>(context));
}

QtObjectPtr QWidget_GraphicsEffect(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->graphicsEffect();
}

QtObjectPtr QWidget_GraphicsProxyWidget(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->graphicsProxyWidget();
}

int QWidget_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->hasHeightForWidth();
}

int QWidget_HasMouseTracking(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->hasMouseTracking();
}

int QWidget_Height(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->height();
}

int QWidget_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QWidget*>(ptr)->heightForWidth(w);
}

void QWidget_Hide(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "hide");
}

char* QWidget_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

void QWidget_InsertAction(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr action){
	static_cast<QWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QWidget_IsAncestorOf(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QWidget*>(ptr)->isAncestorOf(static_cast<QWidget*>(child));
}

int QWidget_IsEnabled(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isEnabled();
}

int QWidget_IsEnabledTo(QtObjectPtr ptr, QtObjectPtr ancestor){
	return static_cast<QWidget*>(ptr)->isEnabledTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsHidden(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isHidden();
}

int QWidget_IsModal(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isModal();
}

int QWidget_IsVisible(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isVisible();
}

int QWidget_IsVisibleTo(QtObjectPtr ptr, QtObjectPtr ancestor){
	return static_cast<QWidget*>(ptr)->isVisibleTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsWindow(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->isWindow();
}

QtObjectPtr QWidget_Layout(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->layout();
}

void QWidget_Lower(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "lower");
}

int QWidget_MaximumHeight(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->maximumHeight();
}

int QWidget_MaximumWidth(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->maximumWidth();
}

int QWidget_MinimumHeight(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->minimumHeight();
}

int QWidget_MinimumWidth(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->minimumWidth();
}

void QWidget_Move2(QtObjectPtr ptr, int x, int y){
	static_cast<QWidget*>(ptr)->move(x, y);
}

QtObjectPtr QWidget_NativeParentWidget(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->nativeParentWidget();
}

QtObjectPtr QWidget_NextInFocusChain(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->nextInFocusChain();
}

void QWidget_OverrideWindowFlags(QtObjectPtr ptr, int flags){
	static_cast<QWidget*>(ptr)->overrideWindowFlags(static_cast<Qt::WindowType>(flags));
}

QtObjectPtr QWidget_ParentWidget(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->parentWidget();
}

QtObjectPtr QWidget_PreviousInFocusChain(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->previousInFocusChain();
}

void QWidget_Raise(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "raise");
}

void QWidget_ReleaseShortcut(QtObjectPtr ptr, int id){
	static_cast<QWidget*>(ptr)->releaseShortcut(id);
}

void QWidget_RemoveAction(QtObjectPtr ptr, QtObjectPtr action){
	static_cast<QWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QWidget_Render(QtObjectPtr ptr, QtObjectPtr target, QtObjectPtr targetOffset, QtObjectPtr sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPaintDevice*>(target), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Render2(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr targetOffset, QtObjectPtr sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Repaint(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "repaint");
}

void QWidget_Repaint3(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRect*>(rect));
}

void QWidget_Repaint4(QtObjectPtr ptr, QtObjectPtr rgn){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRegion*>(rgn));
}

void QWidget_Repaint2(QtObjectPtr ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->repaint(x, y, w, h);
}

void QWidget_Resize2(QtObjectPtr ptr, int w, int h){
	static_cast<QWidget*>(ptr)->resize(w, h);
}

int QWidget_RestoreGeometry(QtObjectPtr ptr, QtObjectPtr geometry){
	return static_cast<QWidget*>(ptr)->restoreGeometry(*static_cast<QByteArray*>(geometry));
}

void QWidget_Scroll(QtObjectPtr ptr, int dx, int dy){
	static_cast<QWidget*>(ptr)->scroll(dx, dy);
}

void QWidget_Scroll2(QtObjectPtr ptr, int dx, int dy, QtObjectPtr r){
	static_cast<QWidget*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(r));
}

void QWidget_SetAttribute(QtObjectPtr ptr, int attribute, int on){
	static_cast<QWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QWidget_SetBackgroundRole(QtObjectPtr ptr, int role){
	static_cast<QWidget*>(ptr)->setBackgroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetBaseSize(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setBaseSize(*static_cast<QSize*>(v));
}

void QWidget_SetBaseSize2(QtObjectPtr ptr, int basew, int baseh){
	static_cast<QWidget*>(ptr)->setBaseSize(basew, baseh);
}

void QWidget_SetContentsMargins2(QtObjectPtr ptr, QtObjectPtr margins){
	static_cast<QWidget*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QWidget_SetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QWidget_SetDisabled(QtObjectPtr ptr, int disable){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWidget_SetFixedHeight(QtObjectPtr ptr, int h){
	static_cast<QWidget*>(ptr)->setFixedHeight(h);
}

void QWidget_SetFixedSize(QtObjectPtr ptr, QtObjectPtr s){
	static_cast<QWidget*>(ptr)->setFixedSize(*static_cast<QSize*>(s));
}

void QWidget_SetFixedWidth(QtObjectPtr ptr, int w){
	static_cast<QWidget*>(ptr)->setFixedWidth(w);
}

void QWidget_SetFocus2(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setFocus");
}

void QWidget_SetFocus(QtObjectPtr ptr, int reason){
	static_cast<QWidget*>(ptr)->setFocus(static_cast<Qt::FocusReason>(reason));
}

void QWidget_SetFocusProxy(QtObjectPtr ptr, QtObjectPtr w){
	static_cast<QWidget*>(ptr)->setFocusProxy(static_cast<QWidget*>(w));
}

void QWidget_SetForegroundRole(QtObjectPtr ptr, int role){
	static_cast<QWidget*>(ptr)->setForegroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetGeometry2(QtObjectPtr ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->setGeometry(x, y, w, h);
}

void QWidget_SetGraphicsEffect(QtObjectPtr ptr, QtObjectPtr effect){
	static_cast<QWidget*>(ptr)->setGraphicsEffect(static_cast<QGraphicsEffect*>(effect));
}

void QWidget_SetHidden(QtObjectPtr ptr, int hidden){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QWidget_SetMaximumSize(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setMaximumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMaximumSize2(QtObjectPtr ptr, int maxw, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumSize(maxw, maxh);
}

void QWidget_SetMinimumSize(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setMinimumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMinimumSize2(QtObjectPtr ptr, int minw, int minh){
	static_cast<QWidget*>(ptr)->setMinimumSize(minw, minh);
}

void QWidget_SetMouseTracking(QtObjectPtr ptr, int enable){
	static_cast<QWidget*>(ptr)->setMouseTracking(enable != 0);
}

void QWidget_SetParent(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent));
}

void QWidget_SetParent2(QtObjectPtr ptr, QtObjectPtr parent, int f){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_SetShortcutAutoRepeat(QtObjectPtr ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutAutoRepeat(id, enable != 0);
}

void QWidget_SetShortcutEnabled(QtObjectPtr ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutEnabled(id, enable != 0);
}

void QWidget_SetSizeIncrement(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QWidget*>(ptr)->setSizeIncrement(*static_cast<QSize*>(v));
}

void QWidget_SetSizeIncrement2(QtObjectPtr ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setSizeIncrement(w, h);
}

void QWidget_SetSizePolicy2(QtObjectPtr ptr, int horizontal, int vertical){
	static_cast<QWidget*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical));
}

void QWidget_SetStyle(QtObjectPtr ptr, QtObjectPtr style){
	static_cast<QWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QWidget_QWidget_SetTabOrder(QtObjectPtr first, QtObjectPtr second){
	QWidget::setTabOrder(static_cast<QWidget*>(first), static_cast<QWidget*>(second));
}

void QWidget_SetWindowRole(QtObjectPtr ptr, char* role){
	static_cast<QWidget*>(ptr)->setWindowRole(QString(role));
}

void QWidget_Show(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "show");
}

void QWidget_ShowFullScreen(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showFullScreen");
}

void QWidget_ShowMaximized(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMaximized");
}

void QWidget_ShowMinimized(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMinimized");
}

void QWidget_ShowNormal(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showNormal");
}

void QWidget_StackUnder(QtObjectPtr ptr, QtObjectPtr w){
	static_cast<QWidget*>(ptr)->stackUnder(static_cast<QWidget*>(w));
}

QtObjectPtr QWidget_Style(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->style();
}

int QWidget_TestAttribute(QtObjectPtr ptr, int attribute){
	return static_cast<QWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QWidget_UnderMouse(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->underMouse();
}

void QWidget_UngrabGesture(QtObjectPtr ptr, int gesture){
	static_cast<QWidget*>(ptr)->ungrabGesture(static_cast<Qt::GestureType>(gesture));
}

void QWidget_Update(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "update");
}

void QWidget_Update3(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QWidget_Update4(QtObjectPtr ptr, QtObjectPtr rgn){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRegion*>(rgn));
}

void QWidget_Update2(QtObjectPtr ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->update(x, y, w, h);
}

void QWidget_UpdateGeometry(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->updateGeometry();
}

int QWidget_UpdatesEnabled(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->updatesEnabled();
}

int QWidget_Width(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->width();
}

QtObjectPtr QWidget_Window(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->window();
}

int QWidget_WindowFlags(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowFlags();
}

QtObjectPtr QWidget_WindowHandle(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowHandle();
}

void QWidget_ConnectWindowIconTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

void QWidget_DisconnectWindowIconTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

char* QWidget_WindowRole(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowRole().toUtf8().data();
}

int QWidget_WindowState(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowState();
}

void QWidget_ConnectWindowTitleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

void QWidget_DisconnectWindowTitleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

int QWidget_WindowType(QtObjectPtr ptr){
	return static_cast<QWidget*>(ptr)->windowType();
}

void QWidget_DestroyQWidget(QtObjectPtr ptr){
	static_cast<QWidget*>(ptr)->~QWidget();
}

QtObjectPtr QWidget_QWidget_CreateWindowContainer(QtObjectPtr window, QtObjectPtr parent, int flags){
	return QWidget::createWindowContainer(static_cast<QWindow*>(window), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

