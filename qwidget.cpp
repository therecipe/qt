#include "qwidget.h"
#include <QWidget>
#include "cgoexport.h"

//MyQWidget
class MyQWidget: public QWidget {
Q_OBJECT
public:
void Signal_WindowIconTextChanged() { callbackQWidget(this, QString("windowIconTextChanged").toUtf8().data()); };
void Signal_WindowTitleChanged() { callbackQWidget(this, QString("windowTitleChanged").toUtf8().data()); };

signals:
void Slot_Close();
void Slot_Hide();
void Slot_Lower();
void Slot_Raise();
void Slot_SetDisabled(bool disable);
void Slot_SetEnabled(bool enabled);
void Slot_SetHidden(bool hidden);
void Slot_SetStyleSheet(QString styleSheet);
void Slot_SetWindowModified(bool modified);
void Slot_SetWindowTitle(QString windowTitle);
void Slot_Show();
void Slot_ShowFullScreen();
void Slot_ShowMaximized();
void Slot_ShowMinimized();
void Slot_ShowNormal();

};
#include "qwidget.moc"


//Public Functions
QtObjectPtr QWidget_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QWidget(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QWidget_Destroy(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->~QWidget();
}

int QWidget_AcceptDrops(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->acceptDrops();
}

char* QWidget_AccessibleDescription(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->accessibleDescription().toUtf8().data();
}

char* QWidget_AccessibleName(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->accessibleName().toUtf8().data();
}

void QWidget_ActivateWindow(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->activateWindow();
}

void QWidget_AdjustSize(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->adjustSize();
}

int QWidget_AutoFillBackground(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->autoFillBackground();
}

QtObjectPtr QWidget_ChildAt_Int_Int(QtObjectPtr ptr, int x, int y)
{
	return ((QWidget*)(ptr))->childAt(x, y);
}

void QWidget_ClearFocus(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->clearFocus();
}

void QWidget_ClearMask(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->clearMask();
}

int QWidget_ContextMenuPolicy(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->contextMenuPolicy();
}

void QWidget_EnsurePolished(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->ensurePolished();
}

int QWidget_FocusPolicy(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->focusPolicy();
}

QtObjectPtr QWidget_FocusProxy(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->focusProxy();
}

QtObjectPtr QWidget_FocusWidget(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->focusWidget();
}

void QWidget_GrabKeyboard(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->grabKeyboard();
}

void QWidget_GrabMouse(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->grabMouse();
}

int QWidget_HasFocus(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->hasFocus();
}

int QWidget_HasHeightForWidth(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->hasHeightForWidth();
}

int QWidget_HasMouseTracking(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->hasMouseTracking();
}

int QWidget_Height(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->height();
}

int QWidget_HeightForWidth_Int(QtObjectPtr ptr, int w)
{
	return ((QWidget*)(ptr))->heightForWidth(w);
}

int QWidget_InputMethodHints(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->inputMethodHints();
}

int QWidget_IsActiveWindow(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isActiveWindow();
}

int QWidget_IsAncestorOf_QWidget(QtObjectPtr ptr, QtObjectPtr child)
{
	return ((QWidget*)(ptr))->isAncestorOf(((QWidget*)(child)));
}

int QWidget_IsEnabled(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isEnabled();
}

int QWidget_IsEnabledTo_QWidget(QtObjectPtr ptr, QtObjectPtr ancestor)
{
	return ((QWidget*)(ptr))->isEnabledTo(((QWidget*)(ancestor)));
}

int QWidget_IsFullScreen(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isFullScreen();
}

int QWidget_IsHidden(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isHidden();
}

int QWidget_IsMaximized(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isMaximized();
}

int QWidget_IsMinimized(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isMinimized();
}

int QWidget_IsModal(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isModal();
}

int QWidget_IsVisible(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isVisible();
}

int QWidget_IsVisibleTo_QWidget(QtObjectPtr ptr, QtObjectPtr ancestor)
{
	return ((QWidget*)(ptr))->isVisibleTo(((QWidget*)(ancestor)));
}

int QWidget_IsWindow(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isWindow();
}

int QWidget_IsWindowModified(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->isWindowModified();
}

QtObjectPtr QWidget_Layout(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->layout();
}

int QWidget_LayoutDirection(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->layoutDirection();
}

int QWidget_MaximumHeight(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->maximumHeight();
}

int QWidget_MaximumWidth(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->maximumWidth();
}

int QWidget_MinimumHeight(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->minimumHeight();
}

int QWidget_MinimumWidth(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->minimumWidth();
}

void QWidget_Move_Int_Int(QtObjectPtr ptr, int x, int y)
{
	((QWidget*)(ptr))->move(x, y);
}

QtObjectPtr QWidget_NativeParentWidget(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->nativeParentWidget();
}

QtObjectPtr QWidget_NextInFocusChain(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->nextInFocusChain();
}

void QWidget_OverrideWindowFlags_WindowType(QtObjectPtr ptr, int flags)
{
	((QWidget*)(ptr))->overrideWindowFlags(((Qt::WindowType)(flags)));
}

QtObjectPtr QWidget_ParentWidget(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->parentWidget();
}

QtObjectPtr QWidget_PreviousInFocusChain(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->previousInFocusChain();
}

void QWidget_ReleaseKeyboard(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->releaseKeyboard();
}

void QWidget_ReleaseMouse(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->releaseMouse();
}

void QWidget_ReleaseShortcut_Int(QtObjectPtr ptr, int id)
{
	((QWidget*)(ptr))->releaseShortcut(id);
}

void QWidget_Repaint_Int_Int_Int_Int(QtObjectPtr ptr, int x, int y, int w, int h)
{
	((QWidget*)(ptr))->repaint(x, y, w, h);
}

void QWidget_Resize_Int_Int(QtObjectPtr ptr, int w, int h)
{
	((QWidget*)(ptr))->resize(w, h);
}

void QWidget_Scroll_Int_Int(QtObjectPtr ptr, int dx, int dy)
{
	((QWidget*)(ptr))->scroll(dx, dy);
}

void QWidget_SetAcceptDrops_Bool(QtObjectPtr ptr, int on)
{
	((QWidget*)(ptr))->setAcceptDrops(on != 0);
}

void QWidget_SetAccessibleDescription_String(QtObjectPtr ptr, char* description)
{
	((QWidget*)(ptr))->setAccessibleDescription(QString(description));
}

void QWidget_SetAccessibleName_String(QtObjectPtr ptr, char* name)
{
	((QWidget*)(ptr))->setAccessibleName(QString(name));
}

void QWidget_SetAttribute_WidgetAttribute_Bool(QtObjectPtr ptr, int attribute, int on)
{
	((QWidget*)(ptr))->setAttribute(((Qt::WidgetAttribute)(attribute)), on != 0);
}

void QWidget_SetAutoFillBackground_Bool(QtObjectPtr ptr, int enabled)
{
	((QWidget*)(ptr))->setAutoFillBackground(enabled != 0);
}

void QWidget_SetBaseSize_Int_Int(QtObjectPtr ptr, int basew, int baseh)
{
	((QWidget*)(ptr))->setBaseSize(basew, baseh);
}

void QWidget_SetContentsMargins_Int_Int_Int_Int(QtObjectPtr ptr, int left, int top, int right, int bottom)
{
	((QWidget*)(ptr))->setContentsMargins(left, top, right, bottom);
}

void QWidget_SetContextMenuPolicy_ContextMenuPolicy(QtObjectPtr ptr, int policy)
{
	((QWidget*)(ptr))->setContextMenuPolicy(((Qt::ContextMenuPolicy)(policy)));
}

void QWidget_SetFixedHeight_Int(QtObjectPtr ptr, int h)
{
	((QWidget*)(ptr))->setFixedHeight(h);
}

void QWidget_SetFixedSize_Int_Int(QtObjectPtr ptr, int w, int h)
{
	((QWidget*)(ptr))->setFixedSize(w, h);
}

void QWidget_SetFixedWidth_Int(QtObjectPtr ptr, int w)
{
	((QWidget*)(ptr))->setFixedWidth(w);
}

void QWidget_SetFocus_FocusReason(QtObjectPtr ptr, int reason)
{
	((QWidget*)(ptr))->setFocus(((Qt::FocusReason)(reason)));
}

void QWidget_SetFocusPolicy_FocusPolicy(QtObjectPtr ptr, int policy)
{
	((QWidget*)(ptr))->setFocusPolicy(((Qt::FocusPolicy)(policy)));
}

void QWidget_SetFocusProxy_QWidget(QtObjectPtr ptr, QtObjectPtr w)
{
	((QWidget*)(ptr))->setFocusProxy(((QWidget*)(w)));
}

void QWidget_SetGeometry_Int_Int_Int_Int(QtObjectPtr ptr, int x, int y, int w, int h)
{
	((QWidget*)(ptr))->setGeometry(x, y, w, h);
}

void QWidget_SetLayout_QLayout(QtObjectPtr ptr, QtObjectPtr layout)
{
	((QWidget*)(ptr))->setLayout(((QLayout*)(layout)));
}

void QWidget_SetLayoutDirection_LayoutDirection(QtObjectPtr ptr, int direction)
{
	((QWidget*)(ptr))->setLayoutDirection(((Qt::LayoutDirection)(direction)));
}

void QWidget_SetMaximumHeight_Int(QtObjectPtr ptr, int maxh)
{
	((QWidget*)(ptr))->setMaximumHeight(maxh);
}

void QWidget_SetMaximumSize_Int_Int(QtObjectPtr ptr, int maxw, int maxh)
{
	((QWidget*)(ptr))->setMaximumSize(maxw, maxh);
}

void QWidget_SetMaximumWidth_Int(QtObjectPtr ptr, int maxw)
{
	((QWidget*)(ptr))->setMaximumWidth(maxw);
}

void QWidget_SetMinimumHeight_Int(QtObjectPtr ptr, int minh)
{
	((QWidget*)(ptr))->setMinimumHeight(minh);
}

void QWidget_SetMinimumSize_Int_Int(QtObjectPtr ptr, int minw, int minh)
{
	((QWidget*)(ptr))->setMinimumSize(minw, minh);
}

void QWidget_SetMinimumWidth_Int(QtObjectPtr ptr, int minw)
{
	((QWidget*)(ptr))->setMinimumWidth(minw);
}

void QWidget_SetMouseTracking_Bool(QtObjectPtr ptr, int enable)
{
	((QWidget*)(ptr))->setMouseTracking(enable != 0);
}

void QWidget_SetShortcutAutoRepeat_Int_Bool(QtObjectPtr ptr, int id, int enable)
{
	((QWidget*)(ptr))->setShortcutAutoRepeat(id, enable != 0);
}

void QWidget_SetShortcutEnabled_Int_Bool(QtObjectPtr ptr, int id, int enable)
{
	((QWidget*)(ptr))->setShortcutEnabled(id, enable != 0);
}

void QWidget_SetSizeIncrement_Int_Int(QtObjectPtr ptr, int w, int h)
{
	((QWidget*)(ptr))->setSizeIncrement(w, h);
}

void QWidget_SetStatusTip_String(QtObjectPtr ptr, char* statusTip)
{
	((QWidget*)(ptr))->setStatusTip(QString(statusTip));
}

void QWidget_SetToolTip_String(QtObjectPtr ptr, char* toolTip)
{
	((QWidget*)(ptr))->setToolTip(QString(toolTip));
}

void QWidget_SetToolTipDuration_Int(QtObjectPtr ptr, int msec)
{
	((QWidget*)(ptr))->setToolTipDuration(msec);
}

void QWidget_SetUpdatesEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QWidget*)(ptr))->setUpdatesEnabled(enable != 0);
}

void QWidget_SetWhatsThis_String(QtObjectPtr ptr, char* whatsThis)
{
	((QWidget*)(ptr))->setWhatsThis(QString(whatsThis));
}

void QWidget_SetWindowFilePath_String(QtObjectPtr ptr, char* filePath)
{
	((QWidget*)(ptr))->setWindowFilePath(QString(filePath));
}

void QWidget_SetWindowFlags_WindowType(QtObjectPtr ptr, int typ)
{
	((QWidget*)(ptr))->setWindowFlags(((Qt::WindowType)(typ)));
}

void QWidget_SetWindowIconText_String(QtObjectPtr ptr, char* iconText)
{
	((QWidget*)(ptr))->setWindowIconText(QString(iconText));
}

void QWidget_SetWindowModality_WindowModality(QtObjectPtr ptr, int windowModality)
{
	((QWidget*)(ptr))->setWindowModality(((Qt::WindowModality)(windowModality)));
}

void QWidget_SetWindowRole_String(QtObjectPtr ptr, char* role)
{
	((QWidget*)(ptr))->setWindowRole(QString(role));
}

void QWidget_SetWindowState_WindowState(QtObjectPtr ptr, int windowState)
{
	((QWidget*)(ptr))->setWindowState(((Qt::WindowState)(windowState)));
}

void QWidget_StackUnder_QWidget(QtObjectPtr ptr, QtObjectPtr w)
{
	((QWidget*)(ptr))->stackUnder(((QWidget*)(w)));
}

char* QWidget_StatusTip(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->statusTip().toUtf8().data();
}

char* QWidget_StyleSheet(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->styleSheet().toUtf8().data();
}

int QWidget_TestAttribute_WidgetAttribute(QtObjectPtr ptr, int attribute)
{
	return ((QWidget*)(ptr))->testAttribute(((Qt::WidgetAttribute)(attribute)));
}

char* QWidget_ToolTip(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->toolTip().toUtf8().data();
}

int QWidget_ToolTipDuration(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->toolTipDuration();
}

int QWidget_UnderMouse(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->underMouse();
}

void QWidget_UngrabGesture_GestureType(QtObjectPtr ptr, int gesture)
{
	((QWidget*)(ptr))->ungrabGesture(((Qt::GestureType)(gesture)));
}

void QWidget_UnsetCursor(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->unsetCursor();
}

void QWidget_UnsetLayoutDirection(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->unsetLayoutDirection();
}

void QWidget_UnsetLocale(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->unsetLocale();
}

void QWidget_Update_Int_Int_Int_Int(QtObjectPtr ptr, int x, int y, int w, int h)
{
	((QWidget*)(ptr))->update(x, y, w, h);
}

void QWidget_UpdateGeometry(QtObjectPtr ptr)
{
	((QWidget*)(ptr))->updateGeometry();
}

int QWidget_UpdatesEnabled(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->updatesEnabled();
}

int QWidget_Width(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->width();
}

QtObjectPtr QWidget_Window(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->window();
}

char* QWidget_WindowFilePath(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowFilePath().toUtf8().data();
}

int QWidget_WindowFlags(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowFlags();
}

char* QWidget_WindowIconText(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowIconText().toUtf8().data();
}

int QWidget_WindowModality(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowModality();
}

char* QWidget_WindowRole(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowRole().toUtf8().data();
}

int QWidget_WindowState(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowState();
}

char* QWidget_WindowTitle(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowTitle().toUtf8().data();
}

int QWidget_WindowType(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->windowType();
}

int QWidget_X(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->x();
}

int QWidget_Y(QtObjectPtr ptr)
{
	return ((QWidget*)(ptr))->y();
}

//Public Slots
void QWidget_ConnectSlotClose(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Close, ((QWidget*)(ptr)), &QWidget::close, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotClose(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Close, ((QWidget*)(ptr)), &QWidget::close);
}

void QWidget_Close(QtObjectPtr ptr)
{
	return ((MyQWidget*)(ptr))->Slot_Close();
}

void QWidget_ConnectSlotHide(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Hide, ((QWidget*)(ptr)), &QWidget::hide, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotHide(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Hide, ((QWidget*)(ptr)), &QWidget::hide);
}

void QWidget_Hide(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_Hide();
}

void QWidget_ConnectSlotLower(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Lower, ((QWidget*)(ptr)), &QWidget::lower, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotLower(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Lower, ((QWidget*)(ptr)), &QWidget::lower);
}

void QWidget_Lower(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_Lower();
}

void QWidget_ConnectSlotRaise(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Raise, ((QWidget*)(ptr)), &QWidget::raise, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotRaise(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Raise, ((QWidget*)(ptr)), &QWidget::raise);
}

void QWidget_Raise(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_Raise();
}

void QWidget_ConnectSlotSetDisabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetDisabled, ((QWidget*)(ptr)), &QWidget::setDisabled, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetDisabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetDisabled, ((QWidget*)(ptr)), &QWidget::setDisabled);
}

void QWidget_SetDisabled_Bool(QtObjectPtr ptr, int disable)
{
	((MyQWidget*)(ptr))->Slot_SetDisabled(disable != 0);
}

void QWidget_ConnectSlotSetEnabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetEnabled, ((QWidget*)(ptr)), &QWidget::setEnabled, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetEnabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetEnabled, ((QWidget*)(ptr)), &QWidget::setEnabled);
}

void QWidget_SetEnabled_Bool(QtObjectPtr ptr, int enabled)
{
	((MyQWidget*)(ptr))->Slot_SetEnabled(enabled != 0);
}

void QWidget_ConnectSlotSetHidden(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetHidden, ((QWidget*)(ptr)), &QWidget::setHidden, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetHidden(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetHidden, ((QWidget*)(ptr)), &QWidget::setHidden);
}

void QWidget_SetHidden_Bool(QtObjectPtr ptr, int hidden)
{
	((MyQWidget*)(ptr))->Slot_SetHidden(hidden != 0);
}

void QWidget_ConnectSlotSetStyleSheet(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetStyleSheet, ((QWidget*)(ptr)), &QWidget::setStyleSheet, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetStyleSheet(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetStyleSheet, ((QWidget*)(ptr)), &QWidget::setStyleSheet);
}

void QWidget_SetStyleSheet_String(QtObjectPtr ptr, char* styleSheet)
{
	((MyQWidget*)(ptr))->Slot_SetStyleSheet(QString(styleSheet));
}

void QWidget_ConnectSlotSetWindowModified(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetWindowModified, ((QWidget*)(ptr)), &QWidget::setWindowModified, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetWindowModified(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetWindowModified, ((QWidget*)(ptr)), &QWidget::setWindowModified);
}

void QWidget_SetWindowModified_Bool(QtObjectPtr ptr, int modified)
{
	((MyQWidget*)(ptr))->Slot_SetWindowModified(modified != 0);
}

void QWidget_ConnectSlotSetWindowTitle(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetWindowTitle, ((QWidget*)(ptr)), &QWidget::setWindowTitle, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotSetWindowTitle(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_SetWindowTitle, ((QWidget*)(ptr)), &QWidget::setWindowTitle);
}

void QWidget_SetWindowTitle_String(QtObjectPtr ptr, char* windowTitle)
{
	((MyQWidget*)(ptr))->Slot_SetWindowTitle(QString(windowTitle));
}

void QWidget_ConnectSlotShow(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Show, ((QWidget*)(ptr)), &QWidget::show, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotShow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_Show, ((QWidget*)(ptr)), &QWidget::show);
}

void QWidget_Show(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_Show();
}

void QWidget_ConnectSlotShowFullScreen(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowFullScreen, ((QWidget*)(ptr)), &QWidget::showFullScreen, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotShowFullScreen(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowFullScreen, ((QWidget*)(ptr)), &QWidget::showFullScreen);
}

void QWidget_ShowFullScreen(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_ShowFullScreen();
}

void QWidget_ConnectSlotShowMaximized(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowMaximized, ((QWidget*)(ptr)), &QWidget::showMaximized, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotShowMaximized(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowMaximized, ((QWidget*)(ptr)), &QWidget::showMaximized);
}

void QWidget_ShowMaximized(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_ShowMaximized();
}

void QWidget_ConnectSlotShowMinimized(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowMinimized, ((QWidget*)(ptr)), &QWidget::showMinimized, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotShowMinimized(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowMinimized, ((QWidget*)(ptr)), &QWidget::showMinimized);
}

void QWidget_ShowMinimized(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_ShowMinimized();
}

void QWidget_ConnectSlotShowNormal(QtObjectPtr ptr)
{
	QObject::connect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowNormal, ((QWidget*)(ptr)), &QWidget::showNormal, Qt::QueuedConnection);
}

void QWidget_DisconnectSlotShowNormal(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQWidget*)(ptr)), &MyQWidget::Slot_ShowNormal, ((QWidget*)(ptr)), &QWidget::showNormal);
}

void QWidget_ShowNormal(QtObjectPtr ptr)
{
	((MyQWidget*)(ptr))->Slot_ShowNormal();
}

//Signals
void QWidget_ConnectSignalWindowIconTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QWidget*)(ptr)), &QWidget::windowIconTextChanged, ((MyQWidget*)(ptr)), &MyQWidget::Signal_WindowIconTextChanged, Qt::QueuedConnection);
}

void QWidget_DisconnectSignalWindowIconTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QWidget*)(ptr)), &QWidget::windowIconTextChanged, ((MyQWidget*)(ptr)), &MyQWidget::Signal_WindowIconTextChanged);
}

void QWidget_ConnectSignalWindowTitleChanged(QtObjectPtr ptr)
{
	QObject::connect(((QWidget*)(ptr)), &QWidget::windowTitleChanged, ((MyQWidget*)(ptr)), &MyQWidget::Signal_WindowTitleChanged, Qt::QueuedConnection);
}

void QWidget_DisconnectSignalWindowTitleChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QWidget*)(ptr)), &QWidget::windowTitleChanged, ((MyQWidget*)(ptr)), &MyQWidget::Signal_WindowTitleChanged);
}

//Static Public Members
QtObjectPtr QWidget_KeyboardGrabber()
{
	return QWidget::keyboardGrabber();
}

QtObjectPtr QWidget_MouseGrabber()
{
	return QWidget::mouseGrabber();
}

void QWidget_SetTabOrder_QWidget_QWidget(QtObjectPtr first, QtObjectPtr second)
{
	QWidget::setTabOrder(((QWidget*)(first)), ((QWidget*)(second)));
}

