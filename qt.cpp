#include "qt.h"
#include <Qt>

// AlignmentFlag
int Qt_AlignLeft() { return Qt::AlignLeft; }
int Qt_AlignRight() { return Qt::AlignRight; }
int Qt_AlignHCenter() { return Qt::AlignHCenter; }
int Qt_AlignJustify() { return Qt::AlignJustify; }
int Qt_AlignTop() { return Qt::AlignTop; }
int Qt_AlignBottom() { return Qt::AlignBottom; }
int Qt_AlignVCenter() { return Qt::AlignVCenter; }
int Qt_AlignBaseline() { return Qt::AlignBaseline; }
int Qt_AlignCenter() { return Qt::AlignCenter; }
int Qt_AlignAbsolute() { return Qt::AlignAbsolute; }
int Qt_AlignLeading() { return Qt::AlignLeading; }
int Qt_AlignTrailing() { return Qt::AlignTrailing; }
int Qt_AlignHorizontal_Mask() { return Qt::AlignHorizontal_Mask; }
int Qt_AlignVertical_Mask() { return Qt::AlignVertical_Mask; }

// AnchorPoint
int Qt_AnchorLeft() { return Qt::AnchorLeft; }
int Qt_AnchorHorizontalCenter() { return Qt::AnchorHorizontalCenter; }
int Qt_AnchorRight() { return Qt::AnchorRight; }
int Qt_AnchorTop() { return Qt::AnchorTop; }
int Qt_AnchorVerticalCenter() { return Qt::AnchorVerticalCenter; }
int Qt_AnchorBottom() { return Qt::AnchorBottom; }

// ApplicationAttribute
int Qt_AA_DontShowIconsInMenus() { return Qt::AA_DontShowIconsInMenus; }
int Qt_AA_NativeWindows() { return Qt::AA_NativeWindows; }
int Qt_AA_DontCreateNativeWidgetSiblings() { return Qt::AA_DontCreateNativeWidgetSiblings; }
int Qt_AA_MacPluginApplication() { return Qt::AA_MacPluginApplication; }
int Qt_AA_DontUseNativeMenuBar() { return Qt::AA_DontUseNativeMenuBar; }
int Qt_AA_MacDontSwapCtrlAndMeta() { return Qt::AA_MacDontSwapCtrlAndMeta; }
int Qt_AA_Use96Dpi() { return Qt::AA_Use96Dpi; }
int Qt_AA_X11InitThreads() { return Qt::AA_X11InitThreads; }
int Qt_AA_SynthesizeTouchForUnhandledMouseEvents() { return Qt::AA_SynthesizeTouchForUnhandledMouseEvents; }
int Qt_AA_SynthesizeMouseForUnhandledTouchEvents() { return Qt::AA_SynthesizeMouseForUnhandledTouchEvents; }
int Qt_AA_UseHighDpiPixmaps() { return Qt::AA_UseHighDpiPixmaps; }
int Qt_AA_ForceRasterWidgets() { return Qt::AA_ForceRasterWidgets; }
int Qt_AA_UseDesktopOpenGL() { return Qt::AA_UseDesktopOpenGL; }
int Qt_AA_UseOpenGLES() { return Qt::AA_UseOpenGLES; }
int Qt_AA_ImmediateWidgetCreation() { return Qt::AA_ImmediateWidgetCreation; }
int Qt_AA_MSWindowsUseDirect3DByDefault() { return Qt::AA_MSWindowsUseDirect3DByDefault; }

// ApplicationState
int Qt_ApplicationSuspended() { return Qt::ApplicationSuspended; }
int Qt_ApplicationHidden() { return Qt::ApplicationHidden; }
int Qt_ApplicationInactive() { return Qt::ApplicationInactive; }
int Qt_ApplicationActive() { return Qt::ApplicationActive; }

// ArrowType
int Qt_NoArrow() { return Qt::NoArrow; }
int Qt_UpArrow() { return Qt::UpArrow; }
int Qt_DownArrow() { return Qt::DownArrow; }
int Qt_LeftArrow() { return Qt::LeftArrow; }
int Qt_RightArrow() { return Qt::RightArrow; }

// AspectRatioMode
int Qt_IgnoreAspectRatio() { return Qt::IgnoreAspectRatio; }
int Qt_KeepAspectRatio() { return Qt::KeepAspectRatio; }
int Qt_KeepAspectRatioByExpanding() { return Qt::KeepAspectRatioByExpanding; }

// Axis
int Qt_XAxis() { return Qt::XAxis; }
int Qt_YAxis() { return Qt::YAxis; }
int Qt_ZAxis() { return Qt::ZAxis; }

// BGMode
int Qt_TransparentMode() { return Qt::TransparentMode; }
int Qt_OpaqueMode() { return Qt::OpaqueMode; }

// BrushStyle
int Qt_NoBrush() { return Qt::NoBrush; }
int Qt_SolidPattern() { return Qt::SolidPattern; }
int Qt_Dense1Pattern() { return Qt::Dense1Pattern; }
int Qt_Dense2Pattern() { return Qt::Dense2Pattern; }
int Qt_Dense3Pattern() { return Qt::Dense3Pattern; }
int Qt_Dense4Pattern() { return Qt::Dense4Pattern; }
int Qt_Dense5Pattern() { return Qt::Dense5Pattern; }
int Qt_Dense6Pattern() { return Qt::Dense6Pattern; }
int Qt_Dense7Pattern() { return Qt::Dense7Pattern; }
int Qt_HorPattern() { return Qt::HorPattern; }
int Qt_VerPattern() { return Qt::VerPattern; }
int Qt_CrossPattern() { return Qt::CrossPattern; }
int Qt_BDiagPattern() { return Qt::BDiagPattern; }
int Qt_FDiagPattern() { return Qt::FDiagPattern; }
int Qt_DiagCrossPattern() { return Qt::DiagCrossPattern; }
int Qt_LinearGradientPattern() { return Qt::LinearGradientPattern; }
int Qt_ConicalGradientPattern() { return Qt::ConicalGradientPattern; }
int Qt_RadialGradientPattern() { return Qt::RadialGradientPattern; }
int Qt_TexturePattern() { return Qt::TexturePattern; }

// CaseSensitivity
int Qt_CaseInsensitive() { return Qt::CaseInsensitive; }
int Qt_CaseSensitive() { return Qt::CaseSensitive; }

// CheckState
int Qt_Unchecked() { return Qt::Unchecked; }
int Qt_PartiallyChecked() { return Qt::PartiallyChecked; }
int Qt_Checked() { return Qt::Checked; }

// ClipOperation
int Qt_NoClip() { return Qt::NoClip; }
int Qt_ReplaceClip() { return Qt::ReplaceClip; }
int Qt_IntersectClip() { return Qt::IntersectClip; }

// ConnectionType
int Qt_AutoConnection() { return Qt::AutoConnection; }
int Qt_DirectConnection() { return Qt::DirectConnection; }
int Qt_QueuedConnection() { return Qt::QueuedConnection; }
int Qt_BlockingQueuedConnection() { return Qt::BlockingQueuedConnection; }
int Qt_UniqueConnection() { return Qt::UniqueConnection; }

// ContextMenuPolicy
int Qt_NoContextMenu() { return Qt::NoContextMenu; }
int Qt_PreventContextMenu() { return Qt::PreventContextMenu; }
int Qt_DefaultContextMenu() { return Qt::DefaultContextMenu; }
int Qt_ActionsContextMenu() { return Qt::ActionsContextMenu; }
int Qt_CustomContextMenu() { return Qt::CustomContextMenu; }

// CoordinateSystem
int Qt_DeviceCoordinates() { return Qt::DeviceCoordinates; }
int Qt_LogicalCoordinates() { return Qt::LogicalCoordinates; }

// Corner
int Qt_TopLeftCorner() { return Qt::TopLeftCorner; }
int Qt_TopRightCorner() { return Qt::TopRightCorner; }
int Qt_BottomLeftCorner() { return Qt::BottomLeftCorner; }
int Qt_BottomRightCorner() { return Qt::BottomRightCorner; }

// CursorMoveStyle
int Qt_LogicalMoveStyle() { return Qt::LogicalMoveStyle; }
int Qt_VisualMoveStyle() { return Qt::VisualMoveStyle; }

// CursorShape
int Qt_ArrowCursor() { return Qt::ArrowCursor; }
int Qt_UpArrowCursor() { return Qt::UpArrowCursor; }
int Qt_CrossCursor() { return Qt::CrossCursor; }
int Qt_WaitCursor() { return Qt::WaitCursor; }
int Qt_IBeamCursor() { return Qt::IBeamCursor; }
int Qt_SizeVerCursor() { return Qt::SizeVerCursor; }
int Qt_SizeHorCursor() { return Qt::SizeHorCursor; }
int Qt_SizeBDiagCursor() { return Qt::SizeBDiagCursor; }
int Qt_SizeFDiagCursor() { return Qt::SizeFDiagCursor; }
int Qt_SizeAllCursor() { return Qt::SizeAllCursor; }
int Qt_BlankCursor() { return Qt::BlankCursor; }
int Qt_SplitVCursor() { return Qt::SplitVCursor; }
int Qt_SplitHCursor() { return Qt::SplitHCursor; }
int Qt_PointingHandCursor() { return Qt::PointingHandCursor; }
int Qt_ForbiddenCursor() { return Qt::ForbiddenCursor; }
int Qt_OpenHandCursor() { return Qt::OpenHandCursor; }
int Qt_ClosedHandCursor() { return Qt::ClosedHandCursor; }
int Qt_WhatsThisCursor() { return Qt::WhatsThisCursor; }
int Qt_BusyCursor() { return Qt::BusyCursor; }
int Qt_DragMoveCursor() { return Qt::DragMoveCursor; }
int Qt_DragCopyCursor() { return Qt::DragCopyCursor; }
int Qt_DragLinkCursor() { return Qt::DragLinkCursor; }
int Qt_BitmapCursor() { return Qt::BitmapCursor; }

// DateFormat
int Qt_TextDate() { return Qt::TextDate; }
int Qt_ISODate() { return Qt::ISODate; }
int Qt_SystemLocaleShortDate() { return Qt::SystemLocaleShortDate; }
int Qt_SystemLocaleLongDate() { return Qt::SystemLocaleLongDate; }
int Qt_DefaultLocaleShortDate() { return Qt::DefaultLocaleShortDate; }
int Qt_DefaultLocaleLongDate() { return Qt::DefaultLocaleLongDate; }
int Qt_SystemLocaleDate() { return Qt::SystemLocaleDate; }
int Qt_LocaleDate() { return Qt::LocaleDate; }
int Qt_LocalDate() { return Qt::LocalDate; }
int Qt_RFC2822Date() { return Qt::RFC2822Date; }

// DayOfWeek
int Qt_Monday() { return Qt::Monday; }
int Qt_Tuesday() { return Qt::Tuesday; }
int Qt_Wednesday() { return Qt::Wednesday; }
int Qt_Thursday() { return Qt::Thursday; }
int Qt_Friday() { return Qt::Friday; }
int Qt_Saturday() { return Qt::Saturday; }
int Qt_Sunday() { return Qt::Sunday; }

// DockWidgetArea
int Qt_LeftDockWidgetArea() { return Qt::LeftDockWidgetArea; }
int Qt_RightDockWidgetArea() { return Qt::RightDockWidgetArea; }
int Qt_TopDockWidgetArea() { return Qt::TopDockWidgetArea; }
int Qt_BottomDockWidgetArea() { return Qt::BottomDockWidgetArea; }
int Qt_AllDockWidgetAreas() { return Qt::AllDockWidgetAreas; }
int Qt_NoDockWidgetArea() { return Qt::NoDockWidgetArea; }

// DropAction
int Qt_CopyAction() { return Qt::CopyAction; }
int Qt_MoveAction() { return Qt::MoveAction; }
int Qt_LinkAction() { return Qt::LinkAction; }
int Qt_ActionMask() { return Qt::ActionMask; }
int Qt_IgnoreAction() { return Qt::IgnoreAction; }
int Qt_TargetMoveAction() { return Qt::TargetMoveAction; }

// Edge
int Qt_TopEdge() { return Qt::TopEdge; }
int Qt_LeftEdge() { return Qt::LeftEdge; }
int Qt_RightEdge() { return Qt::RightEdge; }
int Qt_BottomEdge() { return Qt::BottomEdge; }

// EventPriority
int Qt_HighEventPriority() { return Qt::HighEventPriority; }
int Qt_NormalEventPriority() { return Qt::NormalEventPriority; }
int Qt_LowEventPriority() { return Qt::LowEventPriority; }

// FillRule
int Qt_OddEvenFill() { return Qt::OddEvenFill; }
int Qt_WindingFill() { return Qt::WindingFill; }

// FindChildOption
int Qt_FindDirectChildrenOnly() { return Qt::FindDirectChildrenOnly; }
int Qt_FindChildrenRecursively() { return Qt::FindChildrenRecursively; }

// FocusPolicy
int Qt_TabFocus() { return Qt::TabFocus; }
int Qt_ClickFocus() { return Qt::ClickFocus; }
int Qt_StrongFocus() { return Qt::StrongFocus; }
int Qt_WheelFocus() { return Qt::WheelFocus; }
int Qt_NoFocus() { return Qt::NoFocus; }

// FocusReason
int Qt_MouseFocusReason() { return Qt::MouseFocusReason; }
int Qt_TabFocusReason() { return Qt::TabFocusReason; }
int Qt_BacktabFocusReason() { return Qt::BacktabFocusReason; }
int Qt_ActiveWindowFocusReason() { return Qt::ActiveWindowFocusReason; }
int Qt_PopupFocusReason() { return Qt::PopupFocusReason; }
int Qt_ShortcutFocusReason() { return Qt::ShortcutFocusReason; }
int Qt_MenuBarFocusReason() { return Qt::MenuBarFocusReason; }
int Qt_OtherFocusReason() { return Qt::OtherFocusReason; }

// GestureFlag
int Qt_DontStartGestureOnChildren() { return Qt::DontStartGestureOnChildren; }
int Qt_ReceivePartialGestures() { return Qt::ReceivePartialGestures; }
int Qt_IgnoredGesturesPropagateToParent() { return Qt::IgnoredGesturesPropagateToParent; }

// GestureState
int Qt_GestureStarted() { return Qt::GestureStarted; }
int Qt_GestureUpdated() { return Qt::GestureUpdated; }
int Qt_GestureFinished() { return Qt::GestureFinished; }
int Qt_GestureCanceled() { return Qt::GestureCanceled; }

// GestureType
int Qt_TapGesture() { return Qt::TapGesture; }
int Qt_TapAndHoldGesture() { return Qt::TapAndHoldGesture; }
int Qt_PanGesture() { return Qt::PanGesture; }
int Qt_PinchGesture() { return Qt::PinchGesture; }
int Qt_SwipeGesture() { return Qt::SwipeGesture; }
int Qt_CustomGesture() { return Qt::CustomGesture; }

// GlobalColor
int Qt_white() { return Qt::white; }
int Qt_black() { return Qt::black; }
int Qt_red() { return Qt::red; }
int Qt_darkRed() { return Qt::darkRed; }
int Qt_green() { return Qt::green; }
int Qt_darkGreen() { return Qt::darkGreen; }
int Qt_blue() { return Qt::blue; }
int Qt_darkBlue() { return Qt::darkBlue; }
int Qt_cyan() { return Qt::cyan; }
int Qt_darkCyan() { return Qt::darkCyan; }
int Qt_magenta() { return Qt::magenta; }
int Qt_darkMagenta() { return Qt::darkMagenta; }
int Qt_yellow() { return Qt::yellow; }
int Qt_darkYellow() { return Qt::darkYellow; }
int Qt_gray() { return Qt::gray; }
int Qt_darkGray() { return Qt::darkGray; }
int Qt_lightGray() { return Qt::lightGray; }
int Qt_transparent() { return Qt::transparent; }
int Qt_color0() { return Qt::color0; }
int Qt_color1() { return Qt::color1; }

// HitTestAccuracy
int Qt_ExactHit() { return Qt::ExactHit; }
int Qt_FuzzyHit() { return Qt::FuzzyHit; }

// ImageConversionFlag
int Qt_AutoColor() { return Qt::AutoColor; }
int Qt_ColorOnly() { return Qt::ColorOnly; }
int Qt_MonoOnly() { return Qt::MonoOnly; }
int Qt_DiffuseDither() { return Qt::DiffuseDither; }
int Qt_OrderedDither() { return Qt::OrderedDither; }
int Qt_ThresholdDither() { return Qt::ThresholdDither; }
int Qt_ThresholdAlphaDither() { return Qt::ThresholdAlphaDither; }
int Qt_OrderedAlphaDither() { return Qt::OrderedAlphaDither; }
int Qt_DiffuseAlphaDither() { return Qt::DiffuseAlphaDither; }
int Qt_PreferDither() { return Qt::PreferDither; }
int Qt_AvoidDither() { return Qt::AvoidDither; }
int Qt_NoOpaqueDetection() { return Qt::NoOpaqueDetection; }
int Qt_NoFormatConversion() { return Qt::NoFormatConversion; }

// InputMethodHint
int Qt_ImhNone() { return Qt::ImhNone; }
int Qt_ImhHiddenText() { return Qt::ImhHiddenText; }
int Qt_ImhSensitiveData() { return Qt::ImhSensitiveData; }
int Qt_ImhNoAutoUppercase() { return Qt::ImhNoAutoUppercase; }
int Qt_ImhPreferNumbers() { return Qt::ImhPreferNumbers; }
int Qt_ImhPreferUppercase() { return Qt::ImhPreferUppercase; }
int Qt_ImhPreferLowercase() { return Qt::ImhPreferLowercase; }
int Qt_ImhNoPredictiveText() { return Qt::ImhNoPredictiveText; }
int Qt_ImhDate() { return Qt::ImhDate; }
int Qt_ImhTime() { return Qt::ImhTime; }
int Qt_ImhPreferLatin() { return Qt::ImhPreferLatin; }
int Qt_ImhMultiLine() { return Qt::ImhMultiLine; }
int Qt_ImhDigitsOnly() { return Qt::ImhDigitsOnly; }
int Qt_ImhFormattedNumbersOnly() { return Qt::ImhFormattedNumbersOnly; }
int Qt_ImhUppercaseOnly() { return Qt::ImhUppercaseOnly; }
int Qt_ImhLowercaseOnly() { return Qt::ImhLowercaseOnly; }
int Qt_ImhDialableCharactersOnly() { return Qt::ImhDialableCharactersOnly; }
int Qt_ImhEmailCharactersOnly() { return Qt::ImhEmailCharactersOnly; }
int Qt_ImhUrlCharactersOnly() { return Qt::ImhUrlCharactersOnly; }
int Qt_ImhLatinOnly() { return Qt::ImhLatinOnly; }
int Qt_ImhExclusiveInputMask() { return Qt::ImhExclusiveInputMask; }

// InputMethodQuery
int Qt_ImEnabled() { return Qt::ImEnabled; }
int Qt_ImMicroFocus() { return Qt::ImMicroFocus; }
int Qt_ImCursorRectangle() { return Qt::ImCursorRectangle; }
int Qt_ImFont() { return Qt::ImFont; }
int Qt_ImCursorPosition() { return Qt::ImCursorPosition; }
int Qt_ImSurroundingText() { return Qt::ImSurroundingText; }
int Qt_ImCurrentSelection() { return Qt::ImCurrentSelection; }
int Qt_ImMaximumTextLength() { return Qt::ImMaximumTextLength; }
int Qt_ImAnchorPosition() { return Qt::ImAnchorPosition; }
int Qt_ImHints() { return Qt::ImHints; }
int Qt_ImPreferredLanguage() { return Qt::ImPreferredLanguage; }
int Qt_ImPlatformData() { return Qt::ImPlatformData; }
int Qt_ImAbsolutePosition() { return Qt::ImAbsolutePosition; }
int Qt_ImTextBeforeCursor() { return Qt::ImTextBeforeCursor; }
int Qt_ImTextAfterCursor() { return Qt::ImTextAfterCursor; }
int Qt_ImQueryInput() { return Qt::ImQueryInput; }
int Qt_ImQueryAll() { return Qt::ImQueryAll; }

// ItemDataRole
int Qt_DisplayRole() { return Qt::DisplayRole; }
int Qt_DecorationRole() { return Qt::DecorationRole; }
int Qt_EditRole() { return Qt::EditRole; }
int Qt_ToolTipRole() { return Qt::ToolTipRole; }
int Qt_StatusTipRole() { return Qt::StatusTipRole; }
int Qt_WhatsThisRole() { return Qt::WhatsThisRole; }
int Qt_SizeHintRole() { return Qt::SizeHintRole; }
int Qt_FontRole() { return Qt::FontRole; }
int Qt_TextAlignmentRole() { return Qt::TextAlignmentRole; }
int Qt_BackgroundRole() { return Qt::BackgroundRole; }
int Qt_BackgroundColorRole() { return Qt::BackgroundColorRole; }
int Qt_ForegroundRole() { return Qt::ForegroundRole; }
int Qt_TextColorRole() { return Qt::TextColorRole; }
int Qt_CheckStateRole() { return Qt::CheckStateRole; }
int Qt_InitialSortOrderRole() { return Qt::InitialSortOrderRole; }
int Qt_AccessibleTextRole() { return Qt::AccessibleTextRole; }
int Qt_AccessibleDescriptionRole() { return Qt::AccessibleDescriptionRole; }
int Qt_UserRole() { return Qt::UserRole; }

// ItemFlag
int Qt_NoItemFlags() { return Qt::NoItemFlags; }
int Qt_ItemIsSelectable() { return Qt::ItemIsSelectable; }
int Qt_ItemIsEditable() { return Qt::ItemIsEditable; }
int Qt_ItemIsDragEnabled() { return Qt::ItemIsDragEnabled; }
int Qt_ItemIsDropEnabled() { return Qt::ItemIsDropEnabled; }
int Qt_ItemIsUserCheckable() { return Qt::ItemIsUserCheckable; }
int Qt_ItemIsEnabled() { return Qt::ItemIsEnabled; }
int Qt_ItemIsTristate() { return Qt::ItemIsTristate; }
int Qt_ItemNeverHasChildren() { return Qt::ItemNeverHasChildren; }

// ItemSelectionMode
int Qt_ContainsItemShape() { return Qt::ContainsItemShape; }
int Qt_IntersectsItemShape() { return Qt::IntersectsItemShape; }
int Qt_ContainsItemBoundingRect() { return Qt::ContainsItemBoundingRect; }
int Qt_IntersectsItemBoundingRect() { return Qt::IntersectsItemBoundingRect; }

// Key
int Qt_Key_Escape() { return Qt::Key_Escape; }
int Qt_Key_Tab() { return Qt::Key_Tab; }
int Qt_Key_Backtab() { return Qt::Key_Backtab; }
int Qt_Key_Backspace() { return Qt::Key_Backspace; }
int Qt_Key_Return() { return Qt::Key_Return; }
int Qt_Key_Enter() { return Qt::Key_Enter; }
int Qt_Key_Insert() { return Qt::Key_Insert; }
int Qt_Key_Delete() { return Qt::Key_Delete; }
int Qt_Key_Pause() { return Qt::Key_Pause; }
int Qt_Key_Print() { return Qt::Key_Print; }
int Qt_Key_SysReq() { return Qt::Key_SysReq; }
int Qt_Key_Clear() { return Qt::Key_Clear; }
int Qt_Key_Home() { return Qt::Key_Home; }
int Qt_Key_End() { return Qt::Key_End; }
int Qt_Key_Left() { return Qt::Key_Left; }
int Qt_Key_Up() { return Qt::Key_Up; }
int Qt_Key_Right() { return Qt::Key_Right; }
int Qt_Key_Down() { return Qt::Key_Down; }
int Qt_Key_PageUp() { return Qt::Key_PageUp; }
int Qt_Key_PageDown() { return Qt::Key_PageDown; }
int Qt_Key_Shift() { return Qt::Key_Shift; }
int Qt_Key_Control() { return Qt::Key_Control; }
int Qt_Key_Meta() { return Qt::Key_Meta; }
int Qt_Key_Alt() { return Qt::Key_Alt; }
int Qt_Key_AltGr() { return Qt::Key_AltGr; }
int Qt_Key_CapsLock() { return Qt::Key_CapsLock; }
int Qt_Key_NumLock() { return Qt::Key_NumLock; }
int Qt_Key_ScrollLock() { return Qt::Key_ScrollLock; }
int Qt_Key_F1() { return Qt::Key_F1; }
int Qt_Key_F2() { return Qt::Key_F2; }
int Qt_Key_F3() { return Qt::Key_F3; }
int Qt_Key_F4() { return Qt::Key_F4; }
int Qt_Key_F5() { return Qt::Key_F5; }
int Qt_Key_F6() { return Qt::Key_F6; }
int Qt_Key_F7() { return Qt::Key_F7; }
int Qt_Key_F8() { return Qt::Key_F8; }
int Qt_Key_F9() { return Qt::Key_F9; }
int Qt_Key_F10() { return Qt::Key_F10; }
int Qt_Key_F11() { return Qt::Key_F11; }
int Qt_Key_F12() { return Qt::Key_F12; }
int Qt_Key_F13() { return Qt::Key_F13; }
int Qt_Key_F14() { return Qt::Key_F14; }
int Qt_Key_F15() { return Qt::Key_F15; }
int Qt_Key_F16() { return Qt::Key_F16; }
int Qt_Key_F17() { return Qt::Key_F17; }
int Qt_Key_F18() { return Qt::Key_F18; }
int Qt_Key_F19() { return Qt::Key_F19; }
int Qt_Key_F20() { return Qt::Key_F20; }
int Qt_Key_F21() { return Qt::Key_F21; }
int Qt_Key_F22() { return Qt::Key_F22; }
int Qt_Key_F23() { return Qt::Key_F23; }
int Qt_Key_F24() { return Qt::Key_F24; }
int Qt_Key_F25() { return Qt::Key_F25; }
int Qt_Key_F26() { return Qt::Key_F26; }
int Qt_Key_F27() { return Qt::Key_F27; }
int Qt_Key_F28() { return Qt::Key_F28; }
int Qt_Key_F29() { return Qt::Key_F29; }
int Qt_Key_F30() { return Qt::Key_F30; }
int Qt_Key_F31() { return Qt::Key_F31; }
int Qt_Key_F32() { return Qt::Key_F32; }
int Qt_Key_F33() { return Qt::Key_F33; }
int Qt_Key_F34() { return Qt::Key_F34; }
int Qt_Key_F35() { return Qt::Key_F35; }
int Qt_Key_Super_L() { return Qt::Key_Super_L; }
int Qt_Key_Super_R() { return Qt::Key_Super_R; }
int Qt_Key_Menu() { return Qt::Key_Menu; }
int Qt_Key_Hyper_L() { return Qt::Key_Hyper_L; }
int Qt_Key_Hyper_R() { return Qt::Key_Hyper_R; }
int Qt_Key_Help() { return Qt::Key_Help; }
int Qt_Key_Direction_L() { return Qt::Key_Direction_L; }
int Qt_Key_Direction_R() { return Qt::Key_Direction_R; }
int Qt_Key_Space() { return Qt::Key_Space; }
int Qt_Key_Any() { return Qt::Key_Any; }
int Qt_Key_Exclam() { return Qt::Key_Exclam; }
int Qt_Key_QuoteDbl() { return Qt::Key_QuoteDbl; }
int Qt_Key_NumberSign() { return Qt::Key_NumberSign; }
int Qt_Key_Dollar() { return Qt::Key_Dollar; }
int Qt_Key_Percent() { return Qt::Key_Percent; }
int Qt_Key_Ampersand() { return Qt::Key_Ampersand; }
int Qt_Key_Apostrophe() { return Qt::Key_Apostrophe; }
int Qt_Key_ParenLeft() { return Qt::Key_ParenLeft; }
int Qt_Key_ParenRight() { return Qt::Key_ParenRight; }
int Qt_Key_Asterisk() { return Qt::Key_Asterisk; }
int Qt_Key_Plus() { return Qt::Key_Plus; }
int Qt_Key_Comma() { return Qt::Key_Comma; }
int Qt_Key_Minus() { return Qt::Key_Minus; }
int Qt_Key_Period() { return Qt::Key_Period; }
int Qt_Key_Slash() { return Qt::Key_Slash; }
int Qt_Key_0() { return Qt::Key_0; }
int Qt_Key_1() { return Qt::Key_1; }
int Qt_Key_2() { return Qt::Key_2; }
int Qt_Key_3() { return Qt::Key_3; }
int Qt_Key_4() { return Qt::Key_4; }
int Qt_Key_5() { return Qt::Key_5; }
int Qt_Key_6() { return Qt::Key_6; }
int Qt_Key_7() { return Qt::Key_7; }
int Qt_Key_8() { return Qt::Key_8; }
int Qt_Key_9() { return Qt::Key_9; }
int Qt_Key_Colon() { return Qt::Key_Colon; }
int Qt_Key_Semicolon() { return Qt::Key_Semicolon; }
int Qt_Key_Less() { return Qt::Key_Less; }
int Qt_Key_Equal() { return Qt::Key_Equal; }
int Qt_Key_Greater() { return Qt::Key_Greater; }
int Qt_Key_Question() { return Qt::Key_Question; }
int Qt_Key_At() { return Qt::Key_At; }
int Qt_Key_A() { return Qt::Key_A; }
int Qt_Key_B() { return Qt::Key_B; }
int Qt_Key_C() { return Qt::Key_C; }
int Qt_Key_D() { return Qt::Key_D; }
int Qt_Key_E() { return Qt::Key_E; }
int Qt_Key_F() { return Qt::Key_F; }
int Qt_Key_G() { return Qt::Key_G; }
int Qt_Key_H() { return Qt::Key_H; }
int Qt_Key_I() { return Qt::Key_I; }
int Qt_Key_J() { return Qt::Key_J; }
int Qt_Key_K() { return Qt::Key_K; }
int Qt_Key_L() { return Qt::Key_L; }
int Qt_Key_M() { return Qt::Key_M; }
int Qt_Key_N() { return Qt::Key_N; }
int Qt_Key_O() { return Qt::Key_O; }
int Qt_Key_P() { return Qt::Key_P; }
int Qt_Key_Q() { return Qt::Key_Q; }
int Qt_Key_R() { return Qt::Key_R; }
int Qt_Key_S() { return Qt::Key_S; }
int Qt_Key_T() { return Qt::Key_T; }
int Qt_Key_U() { return Qt::Key_U; }
int Qt_Key_V() { return Qt::Key_V; }
int Qt_Key_W() { return Qt::Key_W; }
int Qt_Key_X() { return Qt::Key_X; }
int Qt_Key_Y() { return Qt::Key_Y; }
int Qt_Key_Z() { return Qt::Key_Z; }
int Qt_Key_BracketLeft() { return Qt::Key_BracketLeft; }
int Qt_Key_Backslash() { return Qt::Key_Backslash; }
int Qt_Key_BracketRight() { return Qt::Key_BracketRight; }
int Qt_Key_AsciiCircum() { return Qt::Key_AsciiCircum; }
int Qt_Key_Underscore() { return Qt::Key_Underscore; }
int Qt_Key_QuoteLeft() { return Qt::Key_QuoteLeft; }
int Qt_Key_BraceLeft() { return Qt::Key_BraceLeft; }
int Qt_Key_Bar() { return Qt::Key_Bar; }
int Qt_Key_BraceRight() { return Qt::Key_BraceRight; }
int Qt_Key_AsciiTilde() { return Qt::Key_AsciiTilde; }
int Qt_Key_nobreakspace() { return Qt::Key_nobreakspace; }
int Qt_Key_exclamdown() { return Qt::Key_exclamdown; }
int Qt_Key_cent() { return Qt::Key_cent; }
int Qt_Key_sterling() { return Qt::Key_sterling; }
int Qt_Key_currency() { return Qt::Key_currency; }
int Qt_Key_yen() { return Qt::Key_yen; }
int Qt_Key_brokenbar() { return Qt::Key_brokenbar; }
int Qt_Key_section() { return Qt::Key_section; }
int Qt_Key_diaeresis() { return Qt::Key_diaeresis; }
int Qt_Key_copyright() { return Qt::Key_copyright; }
int Qt_Key_ordfeminine() { return Qt::Key_ordfeminine; }
int Qt_Key_guillemotleft() { return Qt::Key_guillemotleft; }
int Qt_Key_notsign() { return Qt::Key_notsign; }
int Qt_Key_hyphen() { return Qt::Key_hyphen; }
int Qt_Key_registered() { return Qt::Key_registered; }
int Qt_Key_macron() { return Qt::Key_macron; }
int Qt_Key_degree() { return Qt::Key_degree; }
int Qt_Key_plusminus() { return Qt::Key_plusminus; }
int Qt_Key_twosuperior() { return Qt::Key_twosuperior; }
int Qt_Key_threesuperior() { return Qt::Key_threesuperior; }
int Qt_Key_acute() { return Qt::Key_acute; }
int Qt_Key_mu() { return Qt::Key_mu; }
int Qt_Key_paragraph() { return Qt::Key_paragraph; }
int Qt_Key_periodcentered() { return Qt::Key_periodcentered; }
int Qt_Key_cedilla() { return Qt::Key_cedilla; }
int Qt_Key_onesuperior() { return Qt::Key_onesuperior; }
int Qt_Key_masculine() { return Qt::Key_masculine; }
int Qt_Key_guillemotright() { return Qt::Key_guillemotright; }
int Qt_Key_onequarter() { return Qt::Key_onequarter; }
int Qt_Key_onehalf() { return Qt::Key_onehalf; }
int Qt_Key_threequarters() { return Qt::Key_threequarters; }
int Qt_Key_questiondown() { return Qt::Key_questiondown; }
int Qt_Key_Agrave() { return Qt::Key_Agrave; }
int Qt_Key_Aacute() { return Qt::Key_Aacute; }
int Qt_Key_Acircumflex() { return Qt::Key_Acircumflex; }
int Qt_Key_Atilde() { return Qt::Key_Atilde; }
int Qt_Key_Adiaeresis() { return Qt::Key_Adiaeresis; }
int Qt_Key_Aring() { return Qt::Key_Aring; }
int Qt_Key_AE() { return Qt::Key_AE; }
int Qt_Key_Ccedilla() { return Qt::Key_Ccedilla; }
int Qt_Key_Egrave() { return Qt::Key_Egrave; }
int Qt_Key_Eacute() { return Qt::Key_Eacute; }
int Qt_Key_Ecircumflex() { return Qt::Key_Ecircumflex; }
int Qt_Key_Ediaeresis() { return Qt::Key_Ediaeresis; }
int Qt_Key_Igrave() { return Qt::Key_Igrave; }
int Qt_Key_Iacute() { return Qt::Key_Iacute; }
int Qt_Key_Icircumflex() { return Qt::Key_Icircumflex; }
int Qt_Key_Idiaeresis() { return Qt::Key_Idiaeresis; }
int Qt_Key_ETH() { return Qt::Key_ETH; }
int Qt_Key_Ntilde() { return Qt::Key_Ntilde; }
int Qt_Key_Ograve() { return Qt::Key_Ograve; }
int Qt_Key_Oacute() { return Qt::Key_Oacute; }
int Qt_Key_Ocircumflex() { return Qt::Key_Ocircumflex; }
int Qt_Key_Otilde() { return Qt::Key_Otilde; }
int Qt_Key_Odiaeresis() { return Qt::Key_Odiaeresis; }
int Qt_Key_multiply() { return Qt::Key_multiply; }
int Qt_Key_Ooblique() { return Qt::Key_Ooblique; }
int Qt_Key_Ugrave() { return Qt::Key_Ugrave; }
int Qt_Key_Uacute() { return Qt::Key_Uacute; }
int Qt_Key_Ucircumflex() { return Qt::Key_Ucircumflex; }
int Qt_Key_Udiaeresis() { return Qt::Key_Udiaeresis; }
int Qt_Key_Yacute() { return Qt::Key_Yacute; }
int Qt_Key_THORN() { return Qt::Key_THORN; }
int Qt_Key_ssharp() { return Qt::Key_ssharp; }
int Qt_Key_division() { return Qt::Key_division; }
int Qt_Key_ydiaeresis() { return Qt::Key_ydiaeresis; }
int Qt_Key_Multi_key() { return Qt::Key_Multi_key; }
int Qt_Key_Codeinput() { return Qt::Key_Codeinput; }
int Qt_Key_SingleCandidate() { return Qt::Key_SingleCandidate; }
int Qt_Key_MultipleCandidate() { return Qt::Key_MultipleCandidate; }
int Qt_Key_PreviousCandidate() { return Qt::Key_PreviousCandidate; }
int Qt_Key_Mode_switch() { return Qt::Key_Mode_switch; }
int Qt_Key_Kanji() { return Qt::Key_Kanji; }
int Qt_Key_Muhenkan() { return Qt::Key_Muhenkan; }
int Qt_Key_Henkan() { return Qt::Key_Henkan; }
int Qt_Key_Romaji() { return Qt::Key_Romaji; }
int Qt_Key_Hiragana() { return Qt::Key_Hiragana; }
int Qt_Key_Katakana() { return Qt::Key_Katakana; }
int Qt_Key_Hiragana_Katakana() { return Qt::Key_Hiragana_Katakana; }
int Qt_Key_Zenkaku() { return Qt::Key_Zenkaku; }
int Qt_Key_Hankaku() { return Qt::Key_Hankaku; }
int Qt_Key_Zenkaku_Hankaku() { return Qt::Key_Zenkaku_Hankaku; }
int Qt_Key_Touroku() { return Qt::Key_Touroku; }
int Qt_Key_Massyo() { return Qt::Key_Massyo; }
int Qt_Key_Kana_Lock() { return Qt::Key_Kana_Lock; }
int Qt_Key_Kana_Shift() { return Qt::Key_Kana_Shift; }
int Qt_Key_Eisu_Shift() { return Qt::Key_Eisu_Shift; }
int Qt_Key_Eisu_toggle() { return Qt::Key_Eisu_toggle; }
int Qt_Key_Hangul() { return Qt::Key_Hangul; }
int Qt_Key_Hangul_Start() { return Qt::Key_Hangul_Start; }
int Qt_Key_Hangul_End() { return Qt::Key_Hangul_End; }
int Qt_Key_Hangul_Hanja() { return Qt::Key_Hangul_Hanja; }
int Qt_Key_Hangul_Jamo() { return Qt::Key_Hangul_Jamo; }
int Qt_Key_Hangul_Romaja() { return Qt::Key_Hangul_Romaja; }
int Qt_Key_Hangul_Jeonja() { return Qt::Key_Hangul_Jeonja; }
int Qt_Key_Hangul_Banja() { return Qt::Key_Hangul_Banja; }
int Qt_Key_Hangul_PreHanja() { return Qt::Key_Hangul_PreHanja; }
int Qt_Key_Hangul_PostHanja() { return Qt::Key_Hangul_PostHanja; }
int Qt_Key_Hangul_Special() { return Qt::Key_Hangul_Special; }
int Qt_Key_Dead_Grave() { return Qt::Key_Dead_Grave; }
int Qt_Key_Dead_Acute() { return Qt::Key_Dead_Acute; }
int Qt_Key_Dead_Circumflex() { return Qt::Key_Dead_Circumflex; }
int Qt_Key_Dead_Tilde() { return Qt::Key_Dead_Tilde; }
int Qt_Key_Dead_Macron() { return Qt::Key_Dead_Macron; }
int Qt_Key_Dead_Breve() { return Qt::Key_Dead_Breve; }
int Qt_Key_Dead_Abovedot() { return Qt::Key_Dead_Abovedot; }
int Qt_Key_Dead_Diaeresis() { return Qt::Key_Dead_Diaeresis; }
int Qt_Key_Dead_Abovering() { return Qt::Key_Dead_Abovering; }
int Qt_Key_Dead_Doubleacute() { return Qt::Key_Dead_Doubleacute; }
int Qt_Key_Dead_Caron() { return Qt::Key_Dead_Caron; }
int Qt_Key_Dead_Cedilla() { return Qt::Key_Dead_Cedilla; }
int Qt_Key_Dead_Ogonek() { return Qt::Key_Dead_Ogonek; }
int Qt_Key_Dead_Iota() { return Qt::Key_Dead_Iota; }
int Qt_Key_Dead_Voiced_Sound() { return Qt::Key_Dead_Voiced_Sound; }
int Qt_Key_Dead_Semivoiced_Sound() { return Qt::Key_Dead_Semivoiced_Sound; }
int Qt_Key_Dead_Belowdot() { return Qt::Key_Dead_Belowdot; }
int Qt_Key_Dead_Hook() { return Qt::Key_Dead_Hook; }
int Qt_Key_Dead_Horn() { return Qt::Key_Dead_Horn; }
int Qt_Key_Back() { return Qt::Key_Back; }
int Qt_Key_Forward() { return Qt::Key_Forward; }
int Qt_Key_Stop() { return Qt::Key_Stop; }
int Qt_Key_Refresh() { return Qt::Key_Refresh; }
int Qt_Key_VolumeDown() { return Qt::Key_VolumeDown; }
int Qt_Key_VolumeMute() { return Qt::Key_VolumeMute; }
int Qt_Key_VolumeUp() { return Qt::Key_VolumeUp; }
int Qt_Key_BassBoost() { return Qt::Key_BassBoost; }
int Qt_Key_BassUp() { return Qt::Key_BassUp; }
int Qt_Key_BassDown() { return Qt::Key_BassDown; }
int Qt_Key_TrebleUp() { return Qt::Key_TrebleUp; }
int Qt_Key_TrebleDown() { return Qt::Key_TrebleDown; }
int Qt_Key_MediaPlay() { return Qt::Key_MediaPlay; }
int Qt_Key_MediaStop() { return Qt::Key_MediaStop; }
int Qt_Key_MediaPrevious() { return Qt::Key_MediaPrevious; }
int Qt_Key_MediaNext() { return Qt::Key_MediaNext; }
int Qt_Key_MediaRecord() { return Qt::Key_MediaRecord; }
int Qt_Key_MediaPause() { return Qt::Key_MediaPause; }
int Qt_Key_MediaTogglePlayPause() { return Qt::Key_MediaTogglePlayPause; }
int Qt_Key_HomePage() { return Qt::Key_HomePage; }
int Qt_Key_Favorites() { return Qt::Key_Favorites; }
int Qt_Key_Search() { return Qt::Key_Search; }
int Qt_Key_Standby() { return Qt::Key_Standby; }
int Qt_Key_OpenUrl() { return Qt::Key_OpenUrl; }
int Qt_Key_LaunchMail() { return Qt::Key_LaunchMail; }
int Qt_Key_LaunchMedia() { return Qt::Key_LaunchMedia; }
int Qt_Key_Launch0() { return Qt::Key_Launch0; }
int Qt_Key_Launch1() { return Qt::Key_Launch1; }
int Qt_Key_Launch2() { return Qt::Key_Launch2; }
int Qt_Key_Launch3() { return Qt::Key_Launch3; }
int Qt_Key_Launch4() { return Qt::Key_Launch4; }
int Qt_Key_Launch5() { return Qt::Key_Launch5; }
int Qt_Key_Launch6() { return Qt::Key_Launch6; }
int Qt_Key_Launch7() { return Qt::Key_Launch7; }
int Qt_Key_Launch8() { return Qt::Key_Launch8; }
int Qt_Key_Launch9() { return Qt::Key_Launch9; }
int Qt_Key_LaunchA() { return Qt::Key_LaunchA; }
int Qt_Key_LaunchB() { return Qt::Key_LaunchB; }
int Qt_Key_LaunchC() { return Qt::Key_LaunchC; }
int Qt_Key_LaunchD() { return Qt::Key_LaunchD; }
int Qt_Key_LaunchE() { return Qt::Key_LaunchE; }
int Qt_Key_LaunchF() { return Qt::Key_LaunchF; }
int Qt_Key_LaunchG() { return Qt::Key_LaunchG; }
int Qt_Key_LaunchH() { return Qt::Key_LaunchH; }
int Qt_Key_MonBrightnessUp() { return Qt::Key_MonBrightnessUp; }
int Qt_Key_MonBrightnessDown() { return Qt::Key_MonBrightnessDown; }
int Qt_Key_KeyboardLightOnOff() { return Qt::Key_KeyboardLightOnOff; }
int Qt_Key_KeyboardBrightnessUp() { return Qt::Key_KeyboardBrightnessUp; }
int Qt_Key_KeyboardBrightnessDown() { return Qt::Key_KeyboardBrightnessDown; }
int Qt_Key_PowerOff() { return Qt::Key_PowerOff; }
int Qt_Key_WakeUp() { return Qt::Key_WakeUp; }
int Qt_Key_Eject() { return Qt::Key_Eject; }
int Qt_Key_ScreenSaver() { return Qt::Key_ScreenSaver; }
int Qt_Key_WWW() { return Qt::Key_WWW; }
int Qt_Key_Memo() { return Qt::Key_Memo; }
int Qt_Key_LightBulb() { return Qt::Key_LightBulb; }
int Qt_Key_Shop() { return Qt::Key_Shop; }
int Qt_Key_History() { return Qt::Key_History; }
int Qt_Key_AddFavorite() { return Qt::Key_AddFavorite; }
int Qt_Key_HotLinks() { return Qt::Key_HotLinks; }
int Qt_Key_BrightnessAdjust() { return Qt::Key_BrightnessAdjust; }
int Qt_Key_Finance() { return Qt::Key_Finance; }
int Qt_Key_Community() { return Qt::Key_Community; }
int Qt_Key_AudioRewind() { return Qt::Key_AudioRewind; }
int Qt_Key_BackForward() { return Qt::Key_BackForward; }
int Qt_Key_ApplicationLeft() { return Qt::Key_ApplicationLeft; }
int Qt_Key_ApplicationRight() { return Qt::Key_ApplicationRight; }
int Qt_Key_Book() { return Qt::Key_Book; }
int Qt_Key_CD() { return Qt::Key_CD; }
int Qt_Key_Calculator() { return Qt::Key_Calculator; }
int Qt_Key_ToDoList() { return Qt::Key_ToDoList; }
int Qt_Key_ClearGrab() { return Qt::Key_ClearGrab; }
int Qt_Key_Close() { return Qt::Key_Close; }
int Qt_Key_Copy() { return Qt::Key_Copy; }
int Qt_Key_Cut() { return Qt::Key_Cut; }
int Qt_Key_Display() { return Qt::Key_Display; }
int Qt_Key_DOS() { return Qt::Key_DOS; }
int Qt_Key_Documents() { return Qt::Key_Documents; }
int Qt_Key_Excel() { return Qt::Key_Excel; }
int Qt_Key_Explorer() { return Qt::Key_Explorer; }
int Qt_Key_Game() { return Qt::Key_Game; }
int Qt_Key_Go() { return Qt::Key_Go; }
int Qt_Key_iTouch() { return Qt::Key_iTouch; }
int Qt_Key_LogOff() { return Qt::Key_LogOff; }
int Qt_Key_Market() { return Qt::Key_Market; }
int Qt_Key_Meeting() { return Qt::Key_Meeting; }
int Qt_Key_MenuKB() { return Qt::Key_MenuKB; }
int Qt_Key_MenuPB() { return Qt::Key_MenuPB; }
int Qt_Key_MySites() { return Qt::Key_MySites; }
int Qt_Key_News() { return Qt::Key_News; }
int Qt_Key_OfficeHome() { return Qt::Key_OfficeHome; }
int Qt_Key_Option() { return Qt::Key_Option; }
int Qt_Key_Paste() { return Qt::Key_Paste; }
int Qt_Key_Phone() { return Qt::Key_Phone; }
int Qt_Key_Calendar() { return Qt::Key_Calendar; }
int Qt_Key_Reply() { return Qt::Key_Reply; }
int Qt_Key_Reload() { return Qt::Key_Reload; }
int Qt_Key_RotateWindows() { return Qt::Key_RotateWindows; }
int Qt_Key_RotationPB() { return Qt::Key_RotationPB; }
int Qt_Key_RotationKB() { return Qt::Key_RotationKB; }
int Qt_Key_Save() { return Qt::Key_Save; }
int Qt_Key_Send() { return Qt::Key_Send; }
int Qt_Key_Spell() { return Qt::Key_Spell; }
int Qt_Key_SplitScreen() { return Qt::Key_SplitScreen; }
int Qt_Key_Support() { return Qt::Key_Support; }
int Qt_Key_TaskPane() { return Qt::Key_TaskPane; }
int Qt_Key_Terminal() { return Qt::Key_Terminal; }
int Qt_Key_Tools() { return Qt::Key_Tools; }
int Qt_Key_Travel() { return Qt::Key_Travel; }
int Qt_Key_Video() { return Qt::Key_Video; }
int Qt_Key_Word() { return Qt::Key_Word; }
int Qt_Key_Xfer() { return Qt::Key_Xfer; }
int Qt_Key_ZoomIn() { return Qt::Key_ZoomIn; }
int Qt_Key_ZoomOut() { return Qt::Key_ZoomOut; }
int Qt_Key_Away() { return Qt::Key_Away; }
int Qt_Key_Messenger() { return Qt::Key_Messenger; }
int Qt_Key_WebCam() { return Qt::Key_WebCam; }
int Qt_Key_MailForward() { return Qt::Key_MailForward; }
int Qt_Key_Pictures() { return Qt::Key_Pictures; }
int Qt_Key_Music() { return Qt::Key_Music; }
int Qt_Key_Battery() { return Qt::Key_Battery; }
int Qt_Key_Bluetooth() { return Qt::Key_Bluetooth; }
int Qt_Key_WLAN() { return Qt::Key_WLAN; }
int Qt_Key_UWB() { return Qt::Key_UWB; }
int Qt_Key_AudioForward() { return Qt::Key_AudioForward; }
int Qt_Key_AudioRepeat() { return Qt::Key_AudioRepeat; }
int Qt_Key_AudioRandomPlay() { return Qt::Key_AudioRandomPlay; }
int Qt_Key_Subtitle() { return Qt::Key_Subtitle; }
int Qt_Key_AudioCycleTrack() { return Qt::Key_AudioCycleTrack; }
int Qt_Key_Time() { return Qt::Key_Time; }
int Qt_Key_Hibernate() { return Qt::Key_Hibernate; }
int Qt_Key_View() { return Qt::Key_View; }
int Qt_Key_TopMenu() { return Qt::Key_TopMenu; }
int Qt_Key_PowerDown() { return Qt::Key_PowerDown; }
int Qt_Key_Suspend() { return Qt::Key_Suspend; }
int Qt_Key_ContrastAdjust() { return Qt::Key_ContrastAdjust; }
int Qt_Key_TouchpadToggle() { return Qt::Key_TouchpadToggle; }
int Qt_Key_TouchpadOn() { return Qt::Key_TouchpadOn; }
int Qt_Key_TouchpadOff() { return Qt::Key_TouchpadOff; }
int Qt_Key_MicMute() { return Qt::Key_MicMute; }
int Qt_Key_Red() { return Qt::Key_Red; }
int Qt_Key_Green() { return Qt::Key_Green; }
int Qt_Key_Yellow() { return Qt::Key_Yellow; }
int Qt_Key_Blue() { return Qt::Key_Blue; }
int Qt_Key_ChannelUp() { return Qt::Key_ChannelUp; }
int Qt_Key_ChannelDown() { return Qt::Key_ChannelDown; }
int Qt_Key_MediaLast() { return Qt::Key_MediaLast; }
int Qt_Key_unknown() { return Qt::Key_unknown; }
int Qt_Key_Call() { return Qt::Key_Call; }
int Qt_Key_Camera() { return Qt::Key_Camera; }
int Qt_Key_CameraFocus() { return Qt::Key_CameraFocus; }
int Qt_Key_Context1() { return Qt::Key_Context1; }
int Qt_Key_Context2() { return Qt::Key_Context2; }
int Qt_Key_Context3() { return Qt::Key_Context3; }
int Qt_Key_Context4() { return Qt::Key_Context4; }
int Qt_Key_Flip() { return Qt::Key_Flip; }
int Qt_Key_Hangup() { return Qt::Key_Hangup; }
int Qt_Key_No() { return Qt::Key_No; }
int Qt_Key_Select() { return Qt::Key_Select; }
int Qt_Key_Yes() { return Qt::Key_Yes; }
int Qt_Key_ToggleCallHangup() { return Qt::Key_ToggleCallHangup; }
int Qt_Key_VoiceDial() { return Qt::Key_VoiceDial; }
int Qt_Key_LastNumberRedial() { return Qt::Key_LastNumberRedial; }
int Qt_Key_Execute() { return Qt::Key_Execute; }
int Qt_Key_Printer() { return Qt::Key_Printer; }
int Qt_Key_Play() { return Qt::Key_Play; }
int Qt_Key_Sleep() { return Qt::Key_Sleep; }
int Qt_Key_Zoom() { return Qt::Key_Zoom; }
int Qt_Key_Cancel() { return Qt::Key_Cancel; }

// KeyboardModifier
int Qt_NoModifier() { return Qt::NoModifier; }
int Qt_ShiftModifier() { return Qt::ShiftModifier; }
int Qt_ControlModifier() { return Qt::ControlModifier; }
int Qt_AltModifier() { return Qt::AltModifier; }
int Qt_MetaModifier() { return Qt::MetaModifier; }
int Qt_KeypadModifier() { return Qt::KeypadModifier; }
int Qt_GroupSwitchModifier() { return Qt::GroupSwitchModifier; }

// LayoutDirection
int Qt_LeftToRight() { return Qt::LeftToRight; }
int Qt_RightToLeft() { return Qt::RightToLeft; }
int Qt_LayoutDirectionAuto() { return Qt::LayoutDirectionAuto; }

// MaskMode
int Qt_MaskInColor() { return Qt::MaskInColor; }
int Qt_MaskOutColor() { return Qt::MaskOutColor; }

// MatchFlag
int Qt_MatchExactly() { return Qt::MatchExactly; }
int Qt_MatchFixedString() { return Qt::MatchFixedString; }
int Qt_MatchContains() { return Qt::MatchContains; }
int Qt_MatchStartsWith() { return Qt::MatchStartsWith; }
int Qt_MatchEndsWith() { return Qt::MatchEndsWith; }
int Qt_MatchCaseSensitive() { return Qt::MatchCaseSensitive; }
int Qt_MatchRegExp() { return Qt::MatchRegExp; }
int Qt_MatchWildcard() { return Qt::MatchWildcard; }
int Qt_MatchWrap() { return Qt::MatchWrap; }
int Qt_MatchRecursive() { return Qt::MatchRecursive; }

// Modifier
int Qt_SHIFT() { return Qt::SHIFT; }
int Qt_META() { return Qt::META; }
int Qt_CTRL() { return Qt::CTRL; }
int Qt_ALT() { return Qt::ALT; }
int Qt_UNICODE_ACCEL() { return Qt::UNICODE_ACCEL; }

// MouseButton
int Qt_NoButton() { return Qt::NoButton; }
int Qt_AllButtons() { return Qt::AllButtons; }
int Qt_LeftButton() { return Qt::LeftButton; }
int Qt_RightButton() { return Qt::RightButton; }
int Qt_MidButton() { return Qt::MidButton; }
int Qt_MiddleButton() { return Qt::MiddleButton; }
int Qt_BackButton() { return Qt::BackButton; }
int Qt_XButton1() { return Qt::XButton1; }
int Qt_ExtraButton1() { return Qt::ExtraButton1; }
int Qt_ForwardButton() { return Qt::ForwardButton; }
int Qt_XButton2() { return Qt::XButton2; }
int Qt_ExtraButton2() { return Qt::ExtraButton2; }
int Qt_TaskButton() { return Qt::TaskButton; }
int Qt_ExtraButton3() { return Qt::ExtraButton3; }
int Qt_ExtraButton4() { return Qt::ExtraButton4; }
int Qt_ExtraButton5() { return Qt::ExtraButton5; }
int Qt_ExtraButton6() { return Qt::ExtraButton6; }
int Qt_ExtraButton7() { return Qt::ExtraButton7; }
int Qt_ExtraButton8() { return Qt::ExtraButton8; }
int Qt_ExtraButton9() { return Qt::ExtraButton9; }
int Qt_ExtraButton10() { return Qt::ExtraButton10; }
int Qt_ExtraButton11() { return Qt::ExtraButton11; }
int Qt_ExtraButton12() { return Qt::ExtraButton12; }
int Qt_ExtraButton13() { return Qt::ExtraButton13; }
int Qt_ExtraButton14() { return Qt::ExtraButton14; }
int Qt_ExtraButton15() { return Qt::ExtraButton15; }
int Qt_ExtraButton16() { return Qt::ExtraButton16; }
int Qt_ExtraButton17() { return Qt::ExtraButton17; }
int Qt_ExtraButton18() { return Qt::ExtraButton18; }
int Qt_ExtraButton19() { return Qt::ExtraButton19; }
int Qt_ExtraButton20() { return Qt::ExtraButton20; }
int Qt_ExtraButton21() { return Qt::ExtraButton21; }
int Qt_ExtraButton22() { return Qt::ExtraButton22; }
int Qt_ExtraButton23() { return Qt::ExtraButton23; }
int Qt_ExtraButton24() { return Qt::ExtraButton24; }

// MouseEventFlag
int Qt_MouseEventCreatedDoubleClick() { return Qt::MouseEventCreatedDoubleClick; }

// MouseEventSource
int Qt_MouseEventNotSynthesized() { return Qt::MouseEventNotSynthesized; }
int Qt_MouseEventSynthesizedBySystem() { return Qt::MouseEventSynthesizedBySystem; }
int Qt_MouseEventSynthesizedByQt() { return Qt::MouseEventSynthesizedByQt; }

// NativeGestureType

// NavigationMode
int Qt_NavigationModeNone() { return Qt::NavigationModeNone; }
int Qt_NavigationModeKeypadTabOrder() { return Qt::NavigationModeKeypadTabOrder; }
int Qt_NavigationModeKeypadDirectional() { return Qt::NavigationModeKeypadDirectional; }
int Qt_NavigationModeCursorAuto() { return Qt::NavigationModeCursorAuto; }
int Qt_NavigationModeCursorForceVisible() { return Qt::NavigationModeCursorForceVisible; }

// Orientation
int Qt_Horizontal() { return Qt::Horizontal; }
int Qt_Vertical() { return Qt::Vertical; }

// PenCapStyle
int Qt_SquareCap() { return Qt::SquareCap; }
int Qt_FlatCap() { return Qt::FlatCap; }
int Qt_RoundCap() { return Qt::RoundCap; }

// PenJoinStyle
int Qt_BevelJoin() { return Qt::BevelJoin; }
int Qt_MiterJoin() { return Qt::MiterJoin; }
int Qt_RoundJoin() { return Qt::RoundJoin; }
int Qt_SvgMiterJoin() { return Qt::SvgMiterJoin; }

// PenStyle
int Qt_SolidLine() { return Qt::SolidLine; }
int Qt_DashDotLine() { return Qt::DashDotLine; }
int Qt_NoPen() { return Qt::NoPen; }
int Qt_DashLine() { return Qt::DashLine; }
int Qt_DotLine() { return Qt::DotLine; }
int Qt_DashDotDotLine() { return Qt::DashDotDotLine; }
int Qt_CustomDashLine() { return Qt::CustomDashLine; }

// ScreenOrientation
int Qt_PrimaryOrientation() { return Qt::PrimaryOrientation; }
int Qt_LandscapeOrientation() { return Qt::LandscapeOrientation; }
int Qt_PortraitOrientation() { return Qt::PortraitOrientation; }
int Qt_InvertedLandscapeOrientation() { return Qt::InvertedLandscapeOrientation; }
int Qt_InvertedPortraitOrientation() { return Qt::InvertedPortraitOrientation; }

// ScrollBarPolicy
int Qt_ScrollBarAsNeeded() { return Qt::ScrollBarAsNeeded; }
int Qt_ScrollBarAlwaysOff() { return Qt::ScrollBarAlwaysOff; }
int Qt_ScrollBarAlwaysOn() { return Qt::ScrollBarAlwaysOn; }

// ScrollPhase
int Qt_ScrollBegin() { return Qt::ScrollBegin; }
int Qt_ScrollUpdate() { return Qt::ScrollUpdate; }
int Qt_ScrollEnd() { return Qt::ScrollEnd; }

// ShortcutContext
int Qt_WidgetShortcut() { return Qt::WidgetShortcut; }
int Qt_WidgetWithChildrenShortcut() { return Qt::WidgetWithChildrenShortcut; }
int Qt_WindowShortcut() { return Qt::WindowShortcut; }
int Qt_ApplicationShortcut() { return Qt::ApplicationShortcut; }

// SizeHint
int Qt_MinimumSize() { return Qt::MinimumSize; }
int Qt_PreferredSize() { return Qt::PreferredSize; }
int Qt_MaximumSize() { return Qt::MaximumSize; }
int Qt_MinimumDescent() { return Qt::MinimumDescent; }

// SizeMode
int Qt_AbsoluteSize() { return Qt::AbsoluteSize; }
int Qt_RelativeSize() { return Qt::RelativeSize; }

// SortOrder
int Qt_AscendingOrder() { return Qt::AscendingOrder; }
int Qt_DescendingOrder() { return Qt::DescendingOrder; }

// TextElideMode
int Qt_ElideLeft() { return Qt::ElideLeft; }
int Qt_ElideRight() { return Qt::ElideRight; }
int Qt_ElideMiddle() { return Qt::ElideMiddle; }
int Qt_ElideNone() { return Qt::ElideNone; }

// TextFlag
int Qt_TextSingleLine() { return Qt::TextSingleLine; }
int Qt_TextDontClip() { return Qt::TextDontClip; }
int Qt_TextExpandTabs() { return Qt::TextExpandTabs; }
int Qt_TextShowMnemonic() { return Qt::TextShowMnemonic; }
int Qt_TextWordWrap() { return Qt::TextWordWrap; }
int Qt_TextWrapAnywhere() { return Qt::TextWrapAnywhere; }
int Qt_TextHideMnemonic() { return Qt::TextHideMnemonic; }
int Qt_TextDontPrint() { return Qt::TextDontPrint; }
int Qt_TextIncludeTrailingSpaces() { return Qt::TextIncludeTrailingSpaces; }
int Qt_TextJustificationForced() { return Qt::TextJustificationForced; }

// TextFormat
int Qt_PlainText() { return Qt::PlainText; }
int Qt_RichText() { return Qt::RichText; }
int Qt_AutoText() { return Qt::AutoText; }

// TextInteractionFlag
int Qt_NoTextInteraction() { return Qt::NoTextInteraction; }
int Qt_TextSelectableByMouse() { return Qt::TextSelectableByMouse; }
int Qt_TextSelectableByKeyboard() { return Qt::TextSelectableByKeyboard; }
int Qt_LinksAccessibleByMouse() { return Qt::LinksAccessibleByMouse; }
int Qt_LinksAccessibleByKeyboard() { return Qt::LinksAccessibleByKeyboard; }
int Qt_TextEditable() { return Qt::TextEditable; }
int Qt_TextEditorInteraction() { return Qt::TextEditorInteraction; }
int Qt_TextBrowserInteraction() { return Qt::TextBrowserInteraction; }

// TileRule
int Qt_StretchTile() { return Qt::StretchTile; }
int Qt_RepeatTile() { return Qt::RepeatTile; }
int Qt_RoundTile() { return Qt::RoundTile; }

// TimeSpec
int Qt_LocalTime() { return Qt::LocalTime; }
int Qt_UTC() { return Qt::UTC; }
int Qt_OffsetFromUTC() { return Qt::OffsetFromUTC; }
int Qt_TimeZone() { return Qt::TimeZone; }

// TimerType
int Qt_PreciseTimer() { return Qt::PreciseTimer; }
int Qt_CoarseTimer() { return Qt::CoarseTimer; }
int Qt_VeryCoarseTimer() { return Qt::VeryCoarseTimer; }

// ToolBarArea
int Qt_LeftToolBarArea() { return Qt::LeftToolBarArea; }
int Qt_RightToolBarArea() { return Qt::RightToolBarArea; }
int Qt_TopToolBarArea() { return Qt::TopToolBarArea; }
int Qt_BottomToolBarArea() { return Qt::BottomToolBarArea; }
int Qt_AllToolBarAreas() { return Qt::AllToolBarAreas; }
int Qt_NoToolBarArea() { return Qt::NoToolBarArea; }

// ToolButtonStyle
int Qt_ToolButtonIconOnly() { return Qt::ToolButtonIconOnly; }
int Qt_ToolButtonTextOnly() { return Qt::ToolButtonTextOnly; }
int Qt_ToolButtonTextBesideIcon() { return Qt::ToolButtonTextBesideIcon; }
int Qt_ToolButtonTextUnderIcon() { return Qt::ToolButtonTextUnderIcon; }
int Qt_ToolButtonFollowStyle() { return Qt::ToolButtonFollowStyle; }

// TouchPointState
int Qt_TouchPointPressed() { return Qt::TouchPointPressed; }
int Qt_TouchPointMoved() { return Qt::TouchPointMoved; }
int Qt_TouchPointStationary() { return Qt::TouchPointStationary; }
int Qt_TouchPointReleased() { return Qt::TouchPointReleased; }

// TransformationMode
int Qt_FastTransformation() { return Qt::FastTransformation; }
int Qt_SmoothTransformation() { return Qt::SmoothTransformation; }

// UIEffect
int Qt_UI_AnimateMenu() { return Qt::UI_AnimateMenu; }
int Qt_UI_FadeMenu() { return Qt::UI_FadeMenu; }
int Qt_UI_AnimateCombo() { return Qt::UI_AnimateCombo; }
int Qt_UI_AnimateTooltip() { return Qt::UI_AnimateTooltip; }
int Qt_UI_FadeTooltip() { return Qt::UI_FadeTooltip; }
int Qt_UI_AnimateToolBox() { return Qt::UI_AnimateToolBox; }

// WhiteSpaceMode
int Qt_WhiteSpaceNormal() { return Qt::WhiteSpaceNormal; }
int Qt_WhiteSpacePre() { return Qt::WhiteSpacePre; }
int Qt_WhiteSpaceNoWrap() { return Qt::WhiteSpaceNoWrap; }

// WidgetAttribute
int Qt_WA_AcceptDrops() { return Qt::WA_AcceptDrops; }
int Qt_WA_AlwaysShowToolTips() { return Qt::WA_AlwaysShowToolTips; }
int Qt_WA_ContentsPropagated() { return Qt::WA_ContentsPropagated; }
int Qt_WA_CustomWhatsThis() { return Qt::WA_CustomWhatsThis; }
int Qt_WA_DeleteOnClose() { return Qt::WA_DeleteOnClose; }
int Qt_WA_Disabled() { return Qt::WA_Disabled; }
int Qt_WA_DontShowOnScreen() { return Qt::WA_DontShowOnScreen; }
int Qt_WA_ForceDisabled() { return Qt::WA_ForceDisabled; }
int Qt_WA_ForceUpdatesDisabled() { return Qt::WA_ForceUpdatesDisabled; }
int Qt_WA_GroupLeader() { return Qt::WA_GroupLeader; }
int Qt_WA_Hover() { return Qt::WA_Hover; }
int Qt_WA_InputMethodEnabled() { return Qt::WA_InputMethodEnabled; }
int Qt_WA_KeyboardFocusChange() { return Qt::WA_KeyboardFocusChange; }
int Qt_WA_KeyCompression() { return Qt::WA_KeyCompression; }
int Qt_WA_LayoutOnEntireRect() { return Qt::WA_LayoutOnEntireRect; }
int Qt_WA_LayoutUsesWidgetRect() { return Qt::WA_LayoutUsesWidgetRect; }
int Qt_WA_MacNoClickThrough() { return Qt::WA_MacNoClickThrough; }
int Qt_WA_MacOpaqueSizeGrip() { return Qt::WA_MacOpaqueSizeGrip; }
int Qt_WA_MacShowFocusRect() { return Qt::WA_MacShowFocusRect; }
int Qt_WA_MacNormalSize() { return Qt::WA_MacNormalSize; }
int Qt_WA_MacSmallSize() { return Qt::WA_MacSmallSize; }
int Qt_WA_MacMiniSize() { return Qt::WA_MacMiniSize; }
int Qt_WA_MacVariableSize() { return Qt::WA_MacVariableSize; }
int Qt_WA_MacBrushedMetal() { return Qt::WA_MacBrushedMetal; }
int Qt_WA_Mapped() { return Qt::WA_Mapped; }
int Qt_WA_MouseNoMask() { return Qt::WA_MouseNoMask; }
int Qt_WA_MouseTracking() { return Qt::WA_MouseTracking; }
int Qt_WA_Moved() { return Qt::WA_Moved; }
int Qt_WA_MSWindowsUseDirect3D() { return Qt::WA_MSWindowsUseDirect3D; }
int Qt_WA_NoBackground() { return Qt::WA_NoBackground; }
int Qt_WA_NoChildEventsForParent() { return Qt::WA_NoChildEventsForParent; }
int Qt_WA_NoChildEventsFromChildren() { return Qt::WA_NoChildEventsFromChildren; }
int Qt_WA_NoMouseReplay() { return Qt::WA_NoMouseReplay; }
int Qt_WA_NoMousePropagation() { return Qt::WA_NoMousePropagation; }
int Qt_WA_TransparentForMouseEvents() { return Qt::WA_TransparentForMouseEvents; }
int Qt_WA_NoSystemBackground() { return Qt::WA_NoSystemBackground; }
int Qt_WA_OpaquePaintEvent() { return Qt::WA_OpaquePaintEvent; }
int Qt_WA_OutsideWSRange() { return Qt::WA_OutsideWSRange; }
int Qt_WA_PaintOnScreen() { return Qt::WA_PaintOnScreen; }
int Qt_WA_PaintUnclipped() { return Qt::WA_PaintUnclipped; }
int Qt_WA_PendingMoveEvent() { return Qt::WA_PendingMoveEvent; }
int Qt_WA_PendingResizeEvent() { return Qt::WA_PendingResizeEvent; }
int Qt_WA_QuitOnClose() { return Qt::WA_QuitOnClose; }
int Qt_WA_Resized() { return Qt::WA_Resized; }
int Qt_WA_RightToLeft() { return Qt::WA_RightToLeft; }
int Qt_WA_SetCursor() { return Qt::WA_SetCursor; }
int Qt_WA_SetFont() { return Qt::WA_SetFont; }
int Qt_WA_SetPalette() { return Qt::WA_SetPalette; }
int Qt_WA_SetStyle() { return Qt::WA_SetStyle; }
int Qt_WA_ShowModal() { return Qt::WA_ShowModal; }
int Qt_WA_StaticContents() { return Qt::WA_StaticContents; }
int Qt_WA_StyleSheet() { return Qt::WA_StyleSheet; }
int Qt_WA_TranslucentBackground() { return Qt::WA_TranslucentBackground; }
int Qt_WA_UnderMouse() { return Qt::WA_UnderMouse; }
int Qt_WA_UpdatesDisabled() { return Qt::WA_UpdatesDisabled; }
int Qt_WA_WindowModified() { return Qt::WA_WindowModified; }
int Qt_WA_WindowPropagation() { return Qt::WA_WindowPropagation; }
int Qt_WA_MacAlwaysShowToolWindow() { return Qt::WA_MacAlwaysShowToolWindow; }
int Qt_WA_SetLocale() { return Qt::WA_SetLocale; }
int Qt_WA_StyledBackground() { return Qt::WA_StyledBackground; }
int Qt_WA_ShowWithoutActivating() { return Qt::WA_ShowWithoutActivating; }
int Qt_WA_NativeWindow() { return Qt::WA_NativeWindow; }
int Qt_WA_DontCreateNativeAncestors() { return Qt::WA_DontCreateNativeAncestors; }
int Qt_WA_X11NetWmWindowTypeDesktop() { return Qt::WA_X11NetWmWindowTypeDesktop; }
int Qt_WA_X11NetWmWindowTypeDock() { return Qt::WA_X11NetWmWindowTypeDock; }
int Qt_WA_X11NetWmWindowTypeToolBar() { return Qt::WA_X11NetWmWindowTypeToolBar; }
int Qt_WA_X11NetWmWindowTypeMenu() { return Qt::WA_X11NetWmWindowTypeMenu; }
int Qt_WA_X11NetWmWindowTypeUtility() { return Qt::WA_X11NetWmWindowTypeUtility; }
int Qt_WA_X11NetWmWindowTypeSplash() { return Qt::WA_X11NetWmWindowTypeSplash; }
int Qt_WA_X11NetWmWindowTypeDialog() { return Qt::WA_X11NetWmWindowTypeDialog; }
int Qt_WA_X11NetWmWindowTypeDropDownMenu() { return Qt::WA_X11NetWmWindowTypeDropDownMenu; }
int Qt_WA_X11NetWmWindowTypePopupMenu() { return Qt::WA_X11NetWmWindowTypePopupMenu; }
int Qt_WA_X11NetWmWindowTypeToolTip() { return Qt::WA_X11NetWmWindowTypeToolTip; }
int Qt_WA_X11NetWmWindowTypeNotification() { return Qt::WA_X11NetWmWindowTypeNotification; }
int Qt_WA_X11NetWmWindowTypeCombo() { return Qt::WA_X11NetWmWindowTypeCombo; }
int Qt_WA_X11NetWmWindowTypeDND() { return Qt::WA_X11NetWmWindowTypeDND; }
int Qt_WA_MacFrameworkScaled() { return Qt::WA_MacFrameworkScaled; }
int Qt_WA_AcceptTouchEvents() { return Qt::WA_AcceptTouchEvents; }
int Qt_WA_TouchPadAcceptSingleTouchEvents() { return Qt::WA_TouchPadAcceptSingleTouchEvents; }
int Qt_WA_X11DoNotAcceptFocus() { return Qt::WA_X11DoNotAcceptFocus; }

// WindowFrameSection
int Qt_NoSection() { return Qt::NoSection; }
int Qt_LeftSection() { return Qt::LeftSection; }
int Qt_TopLeftSection() { return Qt::TopLeftSection; }
int Qt_TopSection() { return Qt::TopSection; }
int Qt_TopRightSection() { return Qt::TopRightSection; }
int Qt_RightSection() { return Qt::RightSection; }
int Qt_BottomRightSection() { return Qt::BottomRightSection; }
int Qt_BottomSection() { return Qt::BottomSection; }
int Qt_BottomLeftSection() { return Qt::BottomLeftSection; }
int Qt_TitleBarArea() { return Qt::TitleBarArea; }

// WindowModality
int Qt_NonModal() { return Qt::NonModal; }
int Qt_WindowModal() { return Qt::WindowModal; }
int Qt_ApplicationModal() { return Qt::ApplicationModal; }

// WindowState
int Qt_WindowNoState() { return Qt::WindowNoState; }
int Qt_WindowMinimized() { return Qt::WindowMinimized; }
int Qt_WindowMaximized() { return Qt::WindowMaximized; }
int Qt_WindowFullScreen() { return Qt::WindowFullScreen; }
int Qt_WindowActive() { return Qt::WindowActive; }

// WindowType
int Qt_Widget() { return Qt::Widget; }
int Qt_Window() { return Qt::Window; }
int Qt_Dialog() { return Qt::Dialog; }
int Qt_Sheet() { return Qt::Sheet; }
int Qt_Drawer() { return Qt::Drawer; }
int Qt_Popup() { return Qt::Popup; }
int Qt_Tool() { return Qt::Tool; }
int Qt_ToolTip() { return Qt::ToolTip; }
int Qt_SplashScreen() { return Qt::SplashScreen; }
int Qt_Desktop() { return Qt::Desktop; }
int Qt_SubWindow() { return Qt::SubWindow; }
int Qt_ForeignWindow() { return Qt::ForeignWindow; }
int Qt_CoverWindow() { return Qt::CoverWindow; }
int Qt_MSWindowsFixedSizeDialogHint() { return Qt::MSWindowsFixedSizeDialogHint; }
int Qt_MSWindowsOwnDC() { return Qt::MSWindowsOwnDC; }
int Qt_BypassWindowManagerHint() { return Qt::BypassWindowManagerHint; }
int Qt_X11BypassWindowManagerHint() { return Qt::X11BypassWindowManagerHint; }
int Qt_FramelessWindowHint() { return Qt::FramelessWindowHint; }
int Qt_NoDropShadowWindowHint() { return Qt::NoDropShadowWindowHint; }
int Qt_CustomizeWindowHint() { return Qt::CustomizeWindowHint; }
int Qt_WindowTitleHint() { return Qt::WindowTitleHint; }
int Qt_WindowSystemMenuHint() { return Qt::WindowSystemMenuHint; }
int Qt_WindowMinimizeButtonHint() { return Qt::WindowMinimizeButtonHint; }
int Qt_WindowMaximizeButtonHint() { return Qt::WindowMaximizeButtonHint; }
int Qt_WindowMinMaxButtonsHint() { return Qt::WindowMinMaxButtonsHint; }
int Qt_WindowCloseButtonHint() { return Qt::WindowCloseButtonHint; }
int Qt_WindowContextHelpButtonHint() { return Qt::WindowContextHelpButtonHint; }
int Qt_MacWindowToolBarButtonHint() { return Qt::MacWindowToolBarButtonHint; }
int Qt_WindowFullscreenButtonHint() { return Qt::WindowFullscreenButtonHint; }
int Qt_BypassGraphicsProxyWidget() { return Qt::BypassGraphicsProxyWidget; }
int Qt_WindowShadeButtonHint() { return Qt::WindowShadeButtonHint; }
int Qt_WindowStaysOnTopHint() { return Qt::WindowStaysOnTopHint; }
int Qt_WindowStaysOnBottomHint() { return Qt::WindowStaysOnBottomHint; }
int Qt_WindowOkButtonHint() { return Qt::WindowOkButtonHint; }
int Qt_WindowCancelButtonHint() { return Qt::WindowCancelButtonHint; }
int Qt_WindowTransparentForInput() { return Qt::WindowTransparentForInput; }
int Qt_WindowOverridesSystemGestures() { return Qt::WindowOverridesSystemGestures; }
int Qt_WindowDoesNotAcceptFocus() { return Qt::WindowDoesNotAcceptFocus; }
int Qt_WindowType_Mask() { return Qt::WindowType_Mask; }
