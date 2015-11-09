#include "qwidget.h"
#include <QUrl>
#include <QObject>
#include <QWindow>
#include <QPaintDevice>
#include <QBitmap>
#include <QAction>
#include <QSize>
#include <QMetaObject>
#include <QPoint>
#include <QPalette>
#include <QSizePolicy>
#include <QKeySequence>
#include <QIcon>
#include <QFont>
#include <QPainter>
#include <QRegion>
#include <QByteArray>
#include <QLayout>
#include <QStyle>
#include <QMargins>
#include <QVariant>
#include <QGraphicsEffect>
#include <QRect>
#include <QCursor>
#include <QString>
#include <QModelIndex>
#include <QLocale>
#include <QWidget>
#include "_cgo_export.h"

class MyQWidget: public QWidget {
public:
void Signal_WindowIconTextChanged(const QString & iconText){callbackQWidgetWindowIconTextChanged(this->objectName().toUtf8().data(), iconText.toUtf8().data());};
void Signal_WindowTitleChanged(const QString & title){callbackQWidgetWindowTitleChanged(this->objectName().toUtf8().data(), title.toUtf8().data());};
};

int QWidget_AcceptDrops(void* ptr){
	return static_cast<QWidget*>(ptr)->acceptDrops();
}

char* QWidget_AccessibleDescription(void* ptr){
	return static_cast<QWidget*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QWidget_AccessibleName(void* ptr){
	return static_cast<QWidget*>(ptr)->accessibleName().toUtf8().data();
}

void QWidget_ActivateWindow(void* ptr){
	static_cast<QWidget*>(ptr)->activateWindow();
}

int QWidget_AutoFillBackground(void* ptr){
	return static_cast<QWidget*>(ptr)->autoFillBackground();
}

void* QWidget_ChildrenRegion(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->childrenRegion());
}

void QWidget_ClearMask(void* ptr){
	static_cast<QWidget*>(ptr)->clearMask();
}

int QWidget_ContextMenuPolicy(void* ptr){
	return static_cast<QWidget*>(ptr)->contextMenuPolicy();
}

int QWidget_FocusPolicy(void* ptr){
	return static_cast<QWidget*>(ptr)->focusPolicy();
}

void QWidget_GrabKeyboard(void* ptr){
	static_cast<QWidget*>(ptr)->grabKeyboard();
}

void QWidget_GrabMouse(void* ptr){
	static_cast<QWidget*>(ptr)->grabMouse();
}

void QWidget_GrabMouse2(void* ptr, void* cursor){
	static_cast<QWidget*>(ptr)->grabMouse(*static_cast<QCursor*>(cursor));
}

int QWidget_HasFocus(void* ptr){
	return static_cast<QWidget*>(ptr)->hasFocus();
}

int QWidget_InputMethodHints(void* ptr){
	return static_cast<QWidget*>(ptr)->inputMethodHints();
}

int QWidget_IsActiveWindow(void* ptr){
	return static_cast<QWidget*>(ptr)->isActiveWindow();
}

int QWidget_IsFullScreen(void* ptr){
	return static_cast<QWidget*>(ptr)->isFullScreen();
}

int QWidget_IsMaximized(void* ptr){
	return static_cast<QWidget*>(ptr)->isMaximized();
}

int QWidget_IsMinimized(void* ptr){
	return static_cast<QWidget*>(ptr)->isMinimized();
}

int QWidget_IsWindowModified(void* ptr){
	return static_cast<QWidget*>(ptr)->isWindowModified();
}

void* QWidget_QWidget_KeyboardGrabber(){
	return QWidget::keyboardGrabber();
}

int QWidget_LayoutDirection(void* ptr){
	return static_cast<QWidget*>(ptr)->layoutDirection();
}

void* QWidget_QWidget_MouseGrabber(){
	return QWidget::mouseGrabber();
}

void QWidget_Move(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->move(*static_cast<QPoint*>(v));
}

void* QWidget_PaintEngine(void* ptr){
	return static_cast<QWidget*>(ptr)->paintEngine();
}

void QWidget_ReleaseKeyboard(void* ptr){
	static_cast<QWidget*>(ptr)->releaseKeyboard();
}

void QWidget_ReleaseMouse(void* ptr){
	static_cast<QWidget*>(ptr)->releaseMouse();
}

void QWidget_Resize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->resize(*static_cast<QSize*>(v));
}

void QWidget_SetAcceptDrops(void* ptr, int on){
	static_cast<QWidget*>(ptr)->setAcceptDrops(on != 0);
}

void QWidget_SetAccessibleDescription(void* ptr, char* description){
	static_cast<QWidget*>(ptr)->setAccessibleDescription(QString(description));
}

void QWidget_SetAccessibleName(void* ptr, char* name){
	static_cast<QWidget*>(ptr)->setAccessibleName(QString(name));
}

void QWidget_SetAutoFillBackground(void* ptr, int enabled){
	static_cast<QWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QWidget_SetContextMenuPolicy(void* ptr, int policy){
	static_cast<QWidget*>(ptr)->setContextMenuPolicy(static_cast<Qt::ContextMenuPolicy>(policy));
}

void QWidget_SetCursor(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setCursor(*static_cast<QCursor*>(v));
}

void QWidget_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QWidget_SetFixedSize2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setFixedSize(w, h);
}

void QWidget_SetFocusPolicy(void* ptr, int policy){
	static_cast<QWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QWidget_SetFont(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setFont(*static_cast<QFont*>(v));
}

void QWidget_SetGeometry(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setGeometry(*static_cast<QRect*>(v));
}

void QWidget_SetInputMethodHints(void* ptr, int hints){
	static_cast<QWidget*>(ptr)->setInputMethodHints(static_cast<Qt::InputMethodHint>(hints));
}

void QWidget_SetLayout(void* ptr, void* layout){
	static_cast<QWidget*>(ptr)->setLayout(static_cast<QLayout*>(layout));
}

void QWidget_SetLayoutDirection(void* ptr, int direction){
	static_cast<QWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QWidget_SetLocale(void* ptr, void* locale){
	static_cast<QWidget*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QWidget_SetMask(void* ptr, void* bitmap){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QBitmap*>(bitmap));
}

void QWidget_SetMask2(void* ptr, void* region){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWidget_SetMaximumHeight(void* ptr, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumHeight(maxh);
}

void QWidget_SetMaximumWidth(void* ptr, int maxw){
	static_cast<QWidget*>(ptr)->setMaximumWidth(maxw);
}

void QWidget_SetMinimumHeight(void* ptr, int minh){
	static_cast<QWidget*>(ptr)->setMinimumHeight(minh);
}

void QWidget_SetMinimumWidth(void* ptr, int minw){
	static_cast<QWidget*>(ptr)->setMinimumWidth(minw);
}

void QWidget_SetPalette(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setPalette(*static_cast<QPalette*>(v));
}

void QWidget_SetSizePolicy(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(v));
}

void QWidget_SetStatusTip(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setStatusTip(QString(v));
}

void QWidget_SetStyleSheet(void* ptr, char* styleSheet){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QWidget_SetToolTip(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setToolTip(QString(v));
}

void QWidget_SetToolTipDuration(void* ptr, int msec){
	static_cast<QWidget*>(ptr)->setToolTipDuration(msec);
}

void QWidget_SetUpdatesEnabled(void* ptr, int enable){
	static_cast<QWidget*>(ptr)->setUpdatesEnabled(enable != 0);
}

void QWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWidget_SetWhatsThis(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setWhatsThis(QString(v));
}

void QWidget_SetWindowFilePath(void* ptr, char* filePath){
	static_cast<QWidget*>(ptr)->setWindowFilePath(QString(filePath));
}

void QWidget_SetWindowFlags(void* ptr, int ty){
	static_cast<QWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(ty));
}

void QWidget_SetWindowIcon(void* ptr, void* icon){
	static_cast<QWidget*>(ptr)->setWindowIcon(*static_cast<QIcon*>(icon));
}

void QWidget_SetWindowIconText(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setWindowIconText(QString(v));
}

void QWidget_SetWindowModality(void* ptr, int windowModality){
	static_cast<QWidget*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QWidget_SetWindowModified(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowModified", Q_ARG(bool, v != 0));
}

void QWidget_SetWindowOpacity(void* ptr, double level){
	static_cast<QWidget*>(ptr)->setWindowOpacity(static_cast<qreal>(level));
}

void QWidget_SetWindowState(void* ptr, int windowState){
	static_cast<QWidget*>(ptr)->setWindowState(static_cast<Qt::WindowState>(windowState));
}

void QWidget_SetWindowTitle(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(v)));
}

char* QWidget_StatusTip(void* ptr){
	return static_cast<QWidget*>(ptr)->statusTip().toUtf8().data();
}

char* QWidget_StyleSheet(void* ptr){
	return static_cast<QWidget*>(ptr)->styleSheet().toUtf8().data();
}

char* QWidget_ToolTip(void* ptr){
	return static_cast<QWidget*>(ptr)->toolTip().toUtf8().data();
}

int QWidget_ToolTipDuration(void* ptr){
	return static_cast<QWidget*>(ptr)->toolTipDuration();
}

void QWidget_UnsetCursor(void* ptr){
	static_cast<QWidget*>(ptr)->unsetCursor();
}

void QWidget_UnsetLayoutDirection(void* ptr){
	static_cast<QWidget*>(ptr)->unsetLayoutDirection();
}

void QWidget_UnsetLocale(void* ptr){
	static_cast<QWidget*>(ptr)->unsetLocale();
}

char* QWidget_WhatsThis(void* ptr){
	return static_cast<QWidget*>(ptr)->whatsThis().toUtf8().data();
}

char* QWidget_WindowFilePath(void* ptr){
	return static_cast<QWidget*>(ptr)->windowFilePath().toUtf8().data();
}

char* QWidget_WindowIconText(void* ptr){
	return static_cast<QWidget*>(ptr)->windowIconText().toUtf8().data();
}

int QWidget_WindowModality(void* ptr){
	return static_cast<QWidget*>(ptr)->windowModality();
}

double QWidget_WindowOpacity(void* ptr){
	return static_cast<double>(static_cast<QWidget*>(ptr)->windowOpacity());
}

char* QWidget_WindowTitle(void* ptr){
	return static_cast<QWidget*>(ptr)->windowTitle().toUtf8().data();
}

int QWidget_X(void* ptr){
	return static_cast<QWidget*>(ptr)->x();
}

int QWidget_Y(void* ptr){
	return static_cast<QWidget*>(ptr)->y();
}

void* QWidget_NewQWidget(void* parent, int f){
	return new QWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_AddAction(void* ptr, void* action){
	static_cast<QWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QWidget_AdjustSize(void* ptr){
	static_cast<QWidget*>(ptr)->adjustSize();
}

int QWidget_BackgroundRole(void* ptr){
	return static_cast<QWidget*>(ptr)->backgroundRole();
}

void* QWidget_BackingStore(void* ptr){
	return static_cast<QWidget*>(ptr)->backingStore();
}

void* QWidget_ChildAt2(void* ptr, void* p){
	return static_cast<QWidget*>(ptr)->childAt(*static_cast<QPoint*>(p));
}

void* QWidget_ChildAt(void* ptr, int x, int y){
	return static_cast<QWidget*>(ptr)->childAt(x, y);
}

void QWidget_ClearFocus(void* ptr){
	static_cast<QWidget*>(ptr)->clearFocus();
}

int QWidget_Close(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "close");
}

void QWidget_EnsurePolished(void* ptr){
	static_cast<QWidget*>(ptr)->ensurePolished();
}

void* QWidget_FocusProxy(void* ptr){
	return static_cast<QWidget*>(ptr)->focusProxy();
}

void* QWidget_FocusWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->focusWidget();
}

int QWidget_ForegroundRole(void* ptr){
	return static_cast<QWidget*>(ptr)->foregroundRole();
}

void QWidget_GetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

void QWidget_GrabGesture(void* ptr, int gesture, int flags){
	static_cast<QWidget*>(ptr)->grabGesture(static_cast<Qt::GestureType>(gesture), static_cast<Qt::GestureFlag>(flags));
}

int QWidget_GrabShortcut(void* ptr, void* key, int context){
	return static_cast<QWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(key), static_cast<Qt::ShortcutContext>(context));
}

void* QWidget_GraphicsEffect(void* ptr){
	return static_cast<QWidget*>(ptr)->graphicsEffect();
}

void* QWidget_GraphicsProxyWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->graphicsProxyWidget();
}

int QWidget_HasHeightForWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->hasHeightForWidth();
}

int QWidget_HasMouseTracking(void* ptr){
	return static_cast<QWidget*>(ptr)->hasMouseTracking();
}

int QWidget_Height(void* ptr){
	return static_cast<QWidget*>(ptr)->height();
}

int QWidget_HeightForWidth(void* ptr, int w){
	return static_cast<QWidget*>(ptr)->heightForWidth(w);
}

void QWidget_Hide(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "hide");
}

void* QWidget_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QWidget_InsertAction(void* ptr, void* before, void* action){
	static_cast<QWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QWidget_IsAncestorOf(void* ptr, void* child){
	return static_cast<QWidget*>(ptr)->isAncestorOf(static_cast<QWidget*>(child));
}

int QWidget_IsEnabled(void* ptr){
	return static_cast<QWidget*>(ptr)->isEnabled();
}

int QWidget_IsEnabledTo(void* ptr, void* ancestor){
	return static_cast<QWidget*>(ptr)->isEnabledTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsHidden(void* ptr){
	return static_cast<QWidget*>(ptr)->isHidden();
}

int QWidget_IsModal(void* ptr){
	return static_cast<QWidget*>(ptr)->isModal();
}

int QWidget_IsVisible(void* ptr){
	return static_cast<QWidget*>(ptr)->isVisible();
}

int QWidget_IsVisibleTo(void* ptr, void* ancestor){
	return static_cast<QWidget*>(ptr)->isVisibleTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsWindow(void* ptr){
	return static_cast<QWidget*>(ptr)->isWindow();
}

void* QWidget_Layout(void* ptr){
	return static_cast<QWidget*>(ptr)->layout();
}

void QWidget_Lower(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "lower");
}

void* QWidget_Mask(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->mask());
}

int QWidget_MaximumHeight(void* ptr){
	return static_cast<QWidget*>(ptr)->maximumHeight();
}

int QWidget_MaximumWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->maximumWidth();
}

int QWidget_MinimumHeight(void* ptr){
	return static_cast<QWidget*>(ptr)->minimumHeight();
}

int QWidget_MinimumWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->minimumWidth();
}

void QWidget_Move2(void* ptr, int x, int y){
	static_cast<QWidget*>(ptr)->move(x, y);
}

void* QWidget_NativeParentWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->nativeParentWidget();
}

void* QWidget_NextInFocusChain(void* ptr){
	return static_cast<QWidget*>(ptr)->nextInFocusChain();
}

void QWidget_OverrideWindowFlags(void* ptr, int flags){
	static_cast<QWidget*>(ptr)->overrideWindowFlags(static_cast<Qt::WindowType>(flags));
}

void* QWidget_ParentWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->parentWidget();
}

void* QWidget_PreviousInFocusChain(void* ptr){
	return static_cast<QWidget*>(ptr)->previousInFocusChain();
}

void QWidget_Raise(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "raise");
}

void QWidget_ReleaseShortcut(void* ptr, int id){
	static_cast<QWidget*>(ptr)->releaseShortcut(id);
}

void QWidget_RemoveAction(void* ptr, void* action){
	static_cast<QWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QWidget_Render(void* ptr, void* target, void* targetOffset, void* sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPaintDevice*>(target), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Render2(void* ptr, void* painter, void* targetOffset, void* sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Repaint(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "repaint");
}

void QWidget_Repaint3(void* ptr, void* rect){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRect*>(rect));
}

void QWidget_Repaint4(void* ptr, void* rgn){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRegion*>(rgn));
}

void QWidget_Repaint2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->repaint(x, y, w, h);
}

void QWidget_Resize2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->resize(w, h);
}

int QWidget_RestoreGeometry(void* ptr, void* geometry){
	return static_cast<QWidget*>(ptr)->restoreGeometry(*static_cast<QByteArray*>(geometry));
}

void* QWidget_SaveGeometry(void* ptr){
	return new QByteArray(static_cast<QWidget*>(ptr)->saveGeometry());
}

void QWidget_Scroll(void* ptr, int dx, int dy){
	static_cast<QWidget*>(ptr)->scroll(dx, dy);
}

void QWidget_Scroll2(void* ptr, int dx, int dy, void* r){
	static_cast<QWidget*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(r));
}

void QWidget_SetAttribute(void* ptr, int attribute, int on){
	static_cast<QWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QWidget_SetBackgroundRole(void* ptr, int role){
	static_cast<QWidget*>(ptr)->setBackgroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetBaseSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setBaseSize(*static_cast<QSize*>(v));
}

void QWidget_SetBaseSize2(void* ptr, int basew, int baseh){
	static_cast<QWidget*>(ptr)->setBaseSize(basew, baseh);
}

void QWidget_SetContentsMargins2(void* ptr, void* margins){
	static_cast<QWidget*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QWidget_SetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QWidget_SetDisabled(void* ptr, int disable){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWidget_SetFixedHeight(void* ptr, int h){
	static_cast<QWidget*>(ptr)->setFixedHeight(h);
}

void QWidget_SetFixedSize(void* ptr, void* s){
	static_cast<QWidget*>(ptr)->setFixedSize(*static_cast<QSize*>(s));
}

void QWidget_SetFixedWidth(void* ptr, int w){
	static_cast<QWidget*>(ptr)->setFixedWidth(w);
}

void QWidget_SetFocus2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setFocus");
}

void QWidget_SetFocus(void* ptr, int reason){
	static_cast<QWidget*>(ptr)->setFocus(static_cast<Qt::FocusReason>(reason));
}

void QWidget_SetFocusProxy(void* ptr, void* w){
	static_cast<QWidget*>(ptr)->setFocusProxy(static_cast<QWidget*>(w));
}

void QWidget_SetForegroundRole(void* ptr, int role){
	static_cast<QWidget*>(ptr)->setForegroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetGeometry2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->setGeometry(x, y, w, h);
}

void QWidget_SetGraphicsEffect(void* ptr, void* effect){
	static_cast<QWidget*>(ptr)->setGraphicsEffect(static_cast<QGraphicsEffect*>(effect));
}

void QWidget_SetHidden(void* ptr, int hidden){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QWidget_SetMaximumSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setMaximumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMaximumSize2(void* ptr, int maxw, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumSize(maxw, maxh);
}

void QWidget_SetMinimumSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setMinimumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMinimumSize2(void* ptr, int minw, int minh){
	static_cast<QWidget*>(ptr)->setMinimumSize(minw, minh);
}

void QWidget_SetMouseTracking(void* ptr, int enable){
	static_cast<QWidget*>(ptr)->setMouseTracking(enable != 0);
}

void QWidget_SetParent(void* ptr, void* parent){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent));
}

void QWidget_SetParent2(void* ptr, void* parent, int f){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_SetShortcutAutoRepeat(void* ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutAutoRepeat(id, enable != 0);
}

void QWidget_SetShortcutEnabled(void* ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutEnabled(id, enable != 0);
}

void QWidget_SetSizeIncrement(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setSizeIncrement(*static_cast<QSize*>(v));
}

void QWidget_SetSizeIncrement2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setSizeIncrement(w, h);
}

void QWidget_SetSizePolicy2(void* ptr, int horizontal, int vertical){
	static_cast<QWidget*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical));
}

void QWidget_SetStyle(void* ptr, void* style){
	static_cast<QWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QWidget_QWidget_SetTabOrder(void* first, void* second){
	QWidget::setTabOrder(static_cast<QWidget*>(first), static_cast<QWidget*>(second));
}

void QWidget_SetWindowRole(void* ptr, char* role){
	static_cast<QWidget*>(ptr)->setWindowRole(QString(role));
}

void QWidget_Show(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "show");
}

void QWidget_ShowFullScreen(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showFullScreen");
}

void QWidget_ShowMaximized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMaximized");
}

void QWidget_ShowMinimized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMinimized");
}

void QWidget_ShowNormal(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showNormal");
}

void QWidget_StackUnder(void* ptr, void* w){
	static_cast<QWidget*>(ptr)->stackUnder(static_cast<QWidget*>(w));
}

void* QWidget_Style(void* ptr){
	return static_cast<QWidget*>(ptr)->style();
}

int QWidget_TestAttribute(void* ptr, int attribute){
	return static_cast<QWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QWidget_UnderMouse(void* ptr){
	return static_cast<QWidget*>(ptr)->underMouse();
}

void QWidget_UngrabGesture(void* ptr, int gesture){
	static_cast<QWidget*>(ptr)->ungrabGesture(static_cast<Qt::GestureType>(gesture));
}

void QWidget_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "update");
}

void QWidget_Update3(void* ptr, void* rect){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QWidget_Update4(void* ptr, void* rgn){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRegion*>(rgn));
}

void QWidget_Update2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->update(x, y, w, h);
}

void QWidget_UpdateGeometry(void* ptr){
	static_cast<QWidget*>(ptr)->updateGeometry();
}

int QWidget_UpdatesEnabled(void* ptr){
	return static_cast<QWidget*>(ptr)->updatesEnabled();
}

void* QWidget_VisibleRegion(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->visibleRegion());
}

int QWidget_Width(void* ptr){
	return static_cast<QWidget*>(ptr)->width();
}

void* QWidget_Window(void* ptr){
	return static_cast<QWidget*>(ptr)->window();
}

int QWidget_WindowFlags(void* ptr){
	return static_cast<QWidget*>(ptr)->windowFlags();
}

void* QWidget_WindowHandle(void* ptr){
	return static_cast<QWidget*>(ptr)->windowHandle();
}

void QWidget_ConnectWindowIconTextChanged(void* ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

void QWidget_DisconnectWindowIconTextChanged(void* ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

char* QWidget_WindowRole(void* ptr){
	return static_cast<QWidget*>(ptr)->windowRole().toUtf8().data();
}

int QWidget_WindowState(void* ptr){
	return static_cast<QWidget*>(ptr)->windowState();
}

void QWidget_ConnectWindowTitleChanged(void* ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

void QWidget_DisconnectWindowTitleChanged(void* ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

int QWidget_WindowType(void* ptr){
	return static_cast<QWidget*>(ptr)->windowType();
}

void QWidget_DestroyQWidget(void* ptr){
	static_cast<QWidget*>(ptr)->~QWidget();
}

void* QWidget_QWidget_CreateWindowContainer(void* window, void* parent, int flags){
	return QWidget::createWindowContainer(static_cast<QWindow*>(window), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

