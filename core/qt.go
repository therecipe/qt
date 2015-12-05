package core

//#include "core.h"
import "C"
import ()

//Qt::AlignmentFlag
type Qt__AlignmentFlag int64

const (
	Qt__AlignLeft            = Qt__AlignmentFlag(0x0001)
	Qt__AlignLeading         = Qt__AlignmentFlag(Qt__AlignLeft)
	Qt__AlignRight           = Qt__AlignmentFlag(0x0002)
	Qt__AlignTrailing        = Qt__AlignmentFlag(Qt__AlignRight)
	Qt__AlignHCenter         = Qt__AlignmentFlag(0x0004)
	Qt__AlignJustify         = Qt__AlignmentFlag(0x0008)
	Qt__AlignAbsolute        = Qt__AlignmentFlag(0x0010)
	Qt__AlignHorizontal_Mask = Qt__AlignmentFlag(Qt__AlignLeft | Qt__AlignRight | Qt__AlignHCenter | Qt__AlignJustify | Qt__AlignAbsolute)
	Qt__AlignTop             = Qt__AlignmentFlag(0x0020)
	Qt__AlignBottom          = Qt__AlignmentFlag(0x0040)
	Qt__AlignVCenter         = Qt__AlignmentFlag(0x0080)
	Qt__AlignBaseline        = Qt__AlignmentFlag(0x0100)
	Qt__AlignVertical_Mask   = Qt__AlignmentFlag(Qt__AlignTop | Qt__AlignBottom | Qt__AlignVCenter | Qt__AlignBaseline)
	Qt__AlignCenter          = Qt__AlignmentFlag(Qt__AlignVCenter | Qt__AlignHCenter)
)

//Qt::AnchorPoint
type Qt__AnchorPoint int64

const (
	Qt__AnchorLeft             = Qt__AnchorPoint(0)
	Qt__AnchorHorizontalCenter = Qt__AnchorPoint(1)
	Qt__AnchorRight            = Qt__AnchorPoint(2)
	Qt__AnchorTop              = Qt__AnchorPoint(3)
	Qt__AnchorVerticalCenter   = Qt__AnchorPoint(4)
	Qt__AnchorBottom           = Qt__AnchorPoint(5)
)

//Qt::ApplicationAttribute
type Qt__ApplicationAttribute int64

const (
	Qt__AA_ImmediateWidgetCreation                = Qt__ApplicationAttribute(0)
	Qt__AA_MSWindowsUseDirect3DByDefault          = Qt__ApplicationAttribute(1)
	Qt__AA_DontShowIconsInMenus                   = Qt__ApplicationAttribute(2)
	Qt__AA_NativeWindows                          = Qt__ApplicationAttribute(3)
	Qt__AA_DontCreateNativeWidgetSiblings         = Qt__ApplicationAttribute(4)
	Qt__AA_MacPluginApplication                   = Qt__ApplicationAttribute(5)
	Qt__AA_DontUseNativeMenuBar                   = Qt__ApplicationAttribute(6)
	Qt__AA_MacDontSwapCtrlAndMeta                 = Qt__ApplicationAttribute(7)
	Qt__AA_Use96Dpi                               = Qt__ApplicationAttribute(8)
	Qt__AA_X11InitThreads                         = Qt__ApplicationAttribute(10)
	Qt__AA_SynthesizeTouchForUnhandledMouseEvents = Qt__ApplicationAttribute(11)
	Qt__AA_SynthesizeMouseForUnhandledTouchEvents = Qt__ApplicationAttribute(12)
	Qt__AA_UseHighDpiPixmaps                      = Qt__ApplicationAttribute(13)
	Qt__AA_ForceRasterWidgets                     = Qt__ApplicationAttribute(14)
	Qt__AA_UseDesktopOpenGL                       = Qt__ApplicationAttribute(15)
	Qt__AA_UseOpenGLES                            = Qt__ApplicationAttribute(16)
	Qt__AA_UseSoftwareOpenGL                      = Qt__ApplicationAttribute(17)
	Qt__AA_ShareOpenGLContexts                    = Qt__ApplicationAttribute(18)
	Qt__AA_SetPalette                             = Qt__ApplicationAttribute(19)
	Qt__AA_AttributeCount                         = Qt__ApplicationAttribute(20)
)

//Qt::ApplicationState
type Qt__ApplicationState int64

const (
	Qt__ApplicationSuspended = Qt__ApplicationState(0x00000000)
	Qt__ApplicationHidden    = Qt__ApplicationState(0x00000001)
	Qt__ApplicationInactive  = Qt__ApplicationState(0x00000002)
	Qt__ApplicationActive    = Qt__ApplicationState(0x00000004)
)

//Qt::ArrowType
type Qt__ArrowType int64

const (
	Qt__NoArrow    = Qt__ArrowType(0)
	Qt__UpArrow    = Qt__ArrowType(1)
	Qt__DownArrow  = Qt__ArrowType(2)
	Qt__LeftArrow  = Qt__ArrowType(3)
	Qt__RightArrow = Qt__ArrowType(4)
)

//Qt::AspectRatioMode
type Qt__AspectRatioMode int64

const (
	Qt__IgnoreAspectRatio          = Qt__AspectRatioMode(0)
	Qt__KeepAspectRatio            = Qt__AspectRatioMode(1)
	Qt__KeepAspectRatioByExpanding = Qt__AspectRatioMode(2)
)

//Qt::Axis
type Qt__Axis int64

const (
	Qt__XAxis = Qt__Axis(0)
	Qt__YAxis = Qt__Axis(1)
	Qt__ZAxis = Qt__Axis(2)
)

//Qt::BGMode
type Qt__BGMode int64

const (
	Qt__TransparentMode = Qt__BGMode(0)
	Qt__OpaqueMode      = Qt__BGMode(1)
)

//Qt::BrushStyle
type Qt__BrushStyle int64

var (
	Qt__NoBrush                = Qt__BrushStyle(0)
	Qt__SolidPattern           = Qt__BrushStyle(1)
	Qt__Dense1Pattern          = Qt__BrushStyle(2)
	Qt__Dense2Pattern          = Qt__BrushStyle(3)
	Qt__Dense3Pattern          = Qt__BrushStyle(4)
	Qt__Dense4Pattern          = Qt__BrushStyle(5)
	Qt__Dense5Pattern          = Qt__BrushStyle(6)
	Qt__Dense6Pattern          = Qt__BrushStyle(7)
	Qt__Dense7Pattern          = Qt__BrushStyle(8)
	Qt__HorPattern             = Qt__BrushStyle(9)
	Qt__VerPattern             = Qt__BrushStyle(10)
	Qt__CrossPattern           = Qt__BrushStyle(11)
	Qt__BDiagPattern           = Qt__BrushStyle(12)
	Qt__FDiagPattern           = Qt__BrushStyle(13)
	Qt__DiagCrossPattern       = Qt__BrushStyle(14)
	Qt__LinearGradientPattern  = Qt__BrushStyle(15)
	Qt__RadialGradientPattern  = Qt__BrushStyle(16)
	Qt__ConicalGradientPattern = Qt__BrushStyle(17)
	Qt__TexturePattern         = Qt__BrushStyle(24)
)

//Qt::CaseSensitivity
type Qt__CaseSensitivity int64

const (
	Qt__CaseInsensitive = Qt__CaseSensitivity(0)
	Qt__CaseSensitive   = Qt__CaseSensitivity(1)
)

//Qt::CheckState
type Qt__CheckState int64

const (
	Qt__Unchecked        = Qt__CheckState(0)
	Qt__PartiallyChecked = Qt__CheckState(1)
	Qt__Checked          = Qt__CheckState(2)
)

//Qt::ClipOperation
type Qt__ClipOperation int64

const (
	Qt__NoClip        = Qt__ClipOperation(0)
	Qt__ReplaceClip   = Qt__ClipOperation(1)
	Qt__IntersectClip = Qt__ClipOperation(2)
)

//Qt::ConnectionType
type Qt__ConnectionType int64

const (
	Qt__AutoConnection           = Qt__ConnectionType(0)
	Qt__DirectConnection         = Qt__ConnectionType(1)
	Qt__QueuedConnection         = Qt__ConnectionType(2)
	Qt__BlockingQueuedConnection = Qt__ConnectionType(3)
	Qt__UniqueConnection         = Qt__ConnectionType(0x80)
)

//Qt::ContextMenuPolicy
type Qt__ContextMenuPolicy int64

const (
	Qt__NoContextMenu      = Qt__ContextMenuPolicy(0)
	Qt__DefaultContextMenu = Qt__ContextMenuPolicy(1)
	Qt__ActionsContextMenu = Qt__ContextMenuPolicy(2)
	Qt__CustomContextMenu  = Qt__ContextMenuPolicy(3)
	Qt__PreventContextMenu = Qt__ContextMenuPolicy(4)
)

//Qt::CoordinateSystem
type Qt__CoordinateSystem int64

const (
	Qt__DeviceCoordinates  = Qt__CoordinateSystem(0)
	Qt__LogicalCoordinates = Qt__CoordinateSystem(1)
)

//Qt::Corner
type Qt__Corner int64

const (
	Qt__TopLeftCorner     = Qt__Corner(0x00000)
	Qt__TopRightCorner    = Qt__Corner(0x00001)
	Qt__BottomLeftCorner  = Qt__Corner(0x00002)
	Qt__BottomRightCorner = Qt__Corner(0x00003)
)

//Qt::CursorMoveStyle
type Qt__CursorMoveStyle int64

var (
	Qt__LogicalMoveStyle = Qt__CursorMoveStyle(0)
	Qt__VisualMoveStyle  = Qt__CursorMoveStyle(1)
)

//Qt::CursorShape
type Qt__CursorShape int64

const (
	Qt__ArrowCursor        = Qt__CursorShape(0)
	Qt__UpArrowCursor      = Qt__CursorShape(1)
	Qt__CrossCursor        = Qt__CursorShape(2)
	Qt__WaitCursor         = Qt__CursorShape(3)
	Qt__IBeamCursor        = Qt__CursorShape(4)
	Qt__SizeVerCursor      = Qt__CursorShape(5)
	Qt__SizeHorCursor      = Qt__CursorShape(6)
	Qt__SizeBDiagCursor    = Qt__CursorShape(7)
	Qt__SizeFDiagCursor    = Qt__CursorShape(8)
	Qt__SizeAllCursor      = Qt__CursorShape(9)
	Qt__BlankCursor        = Qt__CursorShape(10)
	Qt__SplitVCursor       = Qt__CursorShape(11)
	Qt__SplitHCursor       = Qt__CursorShape(12)
	Qt__PointingHandCursor = Qt__CursorShape(13)
	Qt__ForbiddenCursor    = Qt__CursorShape(14)
	Qt__WhatsThisCursor    = Qt__CursorShape(15)
	Qt__BusyCursor         = Qt__CursorShape(16)
	Qt__OpenHandCursor     = Qt__CursorShape(17)
	Qt__ClosedHandCursor   = Qt__CursorShape(18)
	Qt__DragCopyCursor     = Qt__CursorShape(19)
	Qt__DragMoveCursor     = Qt__CursorShape(20)
	Qt__DragLinkCursor     = Qt__CursorShape(21)
	Qt__LastCursor         = Qt__CursorShape(Qt__DragLinkCursor)
	Qt__BitmapCursor       = Qt__CursorShape(24)
	Qt__CustomCursor       = Qt__CursorShape(25)
)

//Qt::DateFormat
type Qt__DateFormat int64

var (
	Qt__TextDate               = Qt__DateFormat(0)
	Qt__ISODate                = Qt__DateFormat(1)
	Qt__SystemLocaleDate       = Qt__DateFormat(2)
	Qt__LocalDate              = Qt__DateFormat(Qt__SystemLocaleDate)
	Qt__LocaleDate             = Qt__DateFormat(C.Qt_LocaleDate_Type())
	Qt__SystemLocaleShortDate  = Qt__DateFormat(C.Qt_SystemLocaleShortDate_Type())
	Qt__SystemLocaleLongDate   = Qt__DateFormat(C.Qt_SystemLocaleLongDate_Type())
	Qt__DefaultLocaleShortDate = Qt__DateFormat(C.Qt_DefaultLocaleShortDate_Type())
	Qt__DefaultLocaleLongDate  = Qt__DateFormat(C.Qt_DefaultLocaleLongDate_Type())
	Qt__RFC2822Date            = Qt__DateFormat(C.Qt_RFC2822Date_Type())
)

//Qt::DayOfWeek
type Qt__DayOfWeek int64

const (
	Qt__Monday    = Qt__DayOfWeek(1)
	Qt__Tuesday   = Qt__DayOfWeek(2)
	Qt__Wednesday = Qt__DayOfWeek(3)
	Qt__Thursday  = Qt__DayOfWeek(4)
	Qt__Friday    = Qt__DayOfWeek(5)
	Qt__Saturday  = Qt__DayOfWeek(6)
	Qt__Sunday    = Qt__DayOfWeek(7)
)

//Qt::DockWidgetArea
type Qt__DockWidgetArea int64

const (
	Qt__LeftDockWidgetArea   = Qt__DockWidgetArea(0x1)
	Qt__RightDockWidgetArea  = Qt__DockWidgetArea(0x2)
	Qt__TopDockWidgetArea    = Qt__DockWidgetArea(0x4)
	Qt__BottomDockWidgetArea = Qt__DockWidgetArea(0x8)
	Qt__DockWidgetArea_Mask  = Qt__DockWidgetArea(0xf)
	Qt__AllDockWidgetAreas   = Qt__DockWidgetArea(Qt__DockWidgetArea_Mask)
	Qt__NoDockWidgetArea     = Qt__DockWidgetArea(0)
)

//Qt::DropAction
type Qt__DropAction int64

const (
	Qt__CopyAction       = Qt__DropAction(0x1)
	Qt__MoveAction       = Qt__DropAction(0x2)
	Qt__LinkAction       = Qt__DropAction(0x4)
	Qt__ActionMask       = Qt__DropAction(0xff)
	Qt__TargetMoveAction = Qt__DropAction(0x8002)
	Qt__IgnoreAction     = Qt__DropAction(0x0)
)

//Qt::Edge
type Qt__Edge int64

const (
	Qt__TopEdge    = Qt__Edge(0x00001)
	Qt__LeftEdge   = Qt__Edge(0x00002)
	Qt__RightEdge  = Qt__Edge(0x00004)
	Qt__BottomEdge = Qt__Edge(0x00008)
)

//Qt::EventPriority
type Qt__EventPriority int64

const (
	Qt__HighEventPriority   = Qt__EventPriority(1)
	Qt__NormalEventPriority = Qt__EventPriority(0)
	Qt__LowEventPriority    = Qt__EventPriority(-1)
)

//Qt::FillRule
type Qt__FillRule int64

const (
	Qt__OddEvenFill = Qt__FillRule(0)
	Qt__WindingFill = Qt__FillRule(1)
)

//Qt::FindChildOption
type Qt__FindChildOption int64

const (
	Qt__FindDirectChildrenOnly  = Qt__FindChildOption(0x0)
	Qt__FindChildrenRecursively = Qt__FindChildOption(0x1)
)

//Qt::FocusPolicy
type Qt__FocusPolicy int64

const (
	Qt__NoFocus     = Qt__FocusPolicy(0)
	Qt__TabFocus    = Qt__FocusPolicy(0x1)
	Qt__ClickFocus  = Qt__FocusPolicy(0x2)
	Qt__StrongFocus = Qt__FocusPolicy(Qt__TabFocus | Qt__ClickFocus | 0x8)
	Qt__WheelFocus  = Qt__FocusPolicy(Qt__StrongFocus | 0x4)
)

//Qt::FocusReason
type Qt__FocusReason int64

const (
	Qt__MouseFocusReason        = Qt__FocusReason(0)
	Qt__TabFocusReason          = Qt__FocusReason(1)
	Qt__BacktabFocusReason      = Qt__FocusReason(2)
	Qt__ActiveWindowFocusReason = Qt__FocusReason(3)
	Qt__PopupFocusReason        = Qt__FocusReason(4)
	Qt__ShortcutFocusReason     = Qt__FocusReason(5)
	Qt__MenuBarFocusReason      = Qt__FocusReason(6)
	Qt__OtherFocusReason        = Qt__FocusReason(7)
	Qt__NoFocusReason           = Qt__FocusReason(8)
)

//Qt::GestureFlag
type Qt__GestureFlag int64

const (
	Qt__DontStartGestureOnChildren       = Qt__GestureFlag(0x01)
	Qt__ReceivePartialGestures           = Qt__GestureFlag(0x02)
	Qt__IgnoredGesturesPropagateToParent = Qt__GestureFlag(0x04)
)

//Qt::GestureState
type Qt__GestureState int64

const (
	Qt__NoGesture       = Qt__GestureState(0)
	Qt__GestureStarted  = Qt__GestureState(1)
	Qt__GestureUpdated  = Qt__GestureState(2)
	Qt__GestureFinished = Qt__GestureState(3)
	Qt__GestureCanceled = Qt__GestureState(4)
)

//Qt::GestureType
type Qt__GestureType int64

var (
	Qt__TapGesture        = Qt__GestureType(1)
	Qt__TapAndHoldGesture = Qt__GestureType(2)
	Qt__PanGesture        = Qt__GestureType(3)
	Qt__PinchGesture      = Qt__GestureType(4)
	Qt__SwipeGesture      = Qt__GestureType(5)
	Qt__CustomGesture     = Qt__GestureType(0x0100)
	Qt__LastGestureType   = Qt__GestureType(C.Qt_LastGestureType_Type())
)

//Qt::GlobalColor
type Qt__GlobalColor int64

const (
	Qt__color0      = Qt__GlobalColor(0)
	Qt__color1      = Qt__GlobalColor(1)
	Qt__black       = Qt__GlobalColor(2)
	Qt__white       = Qt__GlobalColor(3)
	Qt__darkGray    = Qt__GlobalColor(4)
	Qt__gray        = Qt__GlobalColor(5)
	Qt__lightGray   = Qt__GlobalColor(6)
	Qt__red         = Qt__GlobalColor(7)
	Qt__green       = Qt__GlobalColor(8)
	Qt__blue        = Qt__GlobalColor(9)
	Qt__cyan        = Qt__GlobalColor(10)
	Qt__magenta     = Qt__GlobalColor(11)
	Qt__yellow      = Qt__GlobalColor(12)
	Qt__darkRed     = Qt__GlobalColor(13)
	Qt__darkGreen   = Qt__GlobalColor(14)
	Qt__darkBlue    = Qt__GlobalColor(15)
	Qt__darkCyan    = Qt__GlobalColor(16)
	Qt__darkMagenta = Qt__GlobalColor(17)
	Qt__darkYellow  = Qt__GlobalColor(18)
	Qt__transparent = Qt__GlobalColor(19)
)

//Qt::HitTestAccuracy
type Qt__HitTestAccuracy int64

const (
	Qt__ExactHit = Qt__HitTestAccuracy(0)
	Qt__FuzzyHit = Qt__HitTestAccuracy(1)
)

//Qt::ImageConversionFlag
type Qt__ImageConversionFlag int64

const (
	Qt__ColorMode_Mask       = Qt__ImageConversionFlag(0x00000003)
	Qt__AutoColor            = Qt__ImageConversionFlag(0x00000000)
	Qt__ColorOnly            = Qt__ImageConversionFlag(0x00000003)
	Qt__MonoOnly             = Qt__ImageConversionFlag(0x00000002)
	Qt__AlphaDither_Mask     = Qt__ImageConversionFlag(0x0000000c)
	Qt__ThresholdAlphaDither = Qt__ImageConversionFlag(0x00000000)
	Qt__OrderedAlphaDither   = Qt__ImageConversionFlag(0x00000004)
	Qt__DiffuseAlphaDither   = Qt__ImageConversionFlag(0x00000008)
	Qt__NoAlpha              = Qt__ImageConversionFlag(0x0000000c)
	Qt__Dither_Mask          = Qt__ImageConversionFlag(0x00000030)
	Qt__DiffuseDither        = Qt__ImageConversionFlag(0x00000000)
	Qt__OrderedDither        = Qt__ImageConversionFlag(0x00000010)
	Qt__ThresholdDither      = Qt__ImageConversionFlag(0x00000020)
	Qt__DitherMode_Mask      = Qt__ImageConversionFlag(0x000000c0)
	Qt__AutoDither           = Qt__ImageConversionFlag(0x00000000)
	Qt__PreferDither         = Qt__ImageConversionFlag(0x00000040)
	Qt__AvoidDither          = Qt__ImageConversionFlag(0x00000080)
	Qt__NoOpaqueDetection    = Qt__ImageConversionFlag(0x00000100)
	Qt__NoFormatConversion   = Qt__ImageConversionFlag(0x00000200)
)

//Qt::InputMethodHint
type Qt__InputMethodHint int64

const (
	Qt__ImhNone                   = Qt__InputMethodHint(0x0)
	Qt__ImhHiddenText             = Qt__InputMethodHint(0x1)
	Qt__ImhSensitiveData          = Qt__InputMethodHint(0x2)
	Qt__ImhNoAutoUppercase        = Qt__InputMethodHint(0x4)
	Qt__ImhPreferNumbers          = Qt__InputMethodHint(0x8)
	Qt__ImhPreferUppercase        = Qt__InputMethodHint(0x10)
	Qt__ImhPreferLowercase        = Qt__InputMethodHint(0x20)
	Qt__ImhNoPredictiveText       = Qt__InputMethodHint(0x40)
	Qt__ImhDate                   = Qt__InputMethodHint(0x80)
	Qt__ImhTime                   = Qt__InputMethodHint(0x100)
	Qt__ImhPreferLatin            = Qt__InputMethodHint(0x200)
	Qt__ImhMultiLine              = Qt__InputMethodHint(0x400)
	Qt__ImhDigitsOnly             = Qt__InputMethodHint(0x10000)
	Qt__ImhFormattedNumbersOnly   = Qt__InputMethodHint(0x20000)
	Qt__ImhUppercaseOnly          = Qt__InputMethodHint(0x40000)
	Qt__ImhLowercaseOnly          = Qt__InputMethodHint(0x80000)
	Qt__ImhDialableCharactersOnly = Qt__InputMethodHint(0x100000)
	Qt__ImhEmailCharactersOnly    = Qt__InputMethodHint(0x200000)
	Qt__ImhUrlCharactersOnly      = Qt__InputMethodHint(0x400000)
	Qt__ImhLatinOnly              = Qt__InputMethodHint(0x800000)
	Qt__ImhExclusiveInputMask     = Qt__InputMethodHint(0xffff0000)
)

//Qt::InputMethodQuery
type Qt__InputMethodQuery int64

const (
	Qt__ImEnabled           = Qt__InputMethodQuery(0x1)
	Qt__ImCursorRectangle   = Qt__InputMethodQuery(0x2)
	Qt__ImMicroFocus        = Qt__InputMethodQuery(0x2)
	Qt__ImFont              = Qt__InputMethodQuery(0x4)
	Qt__ImCursorPosition    = Qt__InputMethodQuery(0x8)
	Qt__ImSurroundingText   = Qt__InputMethodQuery(0x10)
	Qt__ImCurrentSelection  = Qt__InputMethodQuery(0x20)
	Qt__ImMaximumTextLength = Qt__InputMethodQuery(0x40)
	Qt__ImAnchorPosition    = Qt__InputMethodQuery(0x80)
	Qt__ImHints             = Qt__InputMethodQuery(0x100)
	Qt__ImPreferredLanguage = Qt__InputMethodQuery(0x200)
	Qt__ImAbsolutePosition  = Qt__InputMethodQuery(0x400)
	Qt__ImTextBeforeCursor  = Qt__InputMethodQuery(0x800)
	Qt__ImTextAfterCursor   = Qt__InputMethodQuery(0x1000)
	Qt__ImPlatformData      = Qt__InputMethodQuery(0x80000000)
	Qt__ImQueryInput        = Qt__InputMethodQuery(Qt__ImCursorRectangle | Qt__ImCursorPosition | Qt__ImSurroundingText | Qt__ImCurrentSelection | Qt__ImAnchorPosition)
	Qt__ImQueryAll          = Qt__InputMethodQuery(0xffffffff)
)

//Qt::ItemDataRole
type Qt__ItemDataRole int64

const (
	Qt__DisplayRole               = Qt__ItemDataRole(0)
	Qt__DecorationRole            = Qt__ItemDataRole(1)
	Qt__EditRole                  = Qt__ItemDataRole(2)
	Qt__ToolTipRole               = Qt__ItemDataRole(3)
	Qt__StatusTipRole             = Qt__ItemDataRole(4)
	Qt__WhatsThisRole             = Qt__ItemDataRole(5)
	Qt__FontRole                  = Qt__ItemDataRole(6)
	Qt__TextAlignmentRole         = Qt__ItemDataRole(7)
	Qt__BackgroundColorRole       = Qt__ItemDataRole(8)
	Qt__BackgroundRole            = Qt__ItemDataRole(8)
	Qt__TextColorRole             = Qt__ItemDataRole(9)
	Qt__ForegroundRole            = Qt__ItemDataRole(9)
	Qt__CheckStateRole            = Qt__ItemDataRole(10)
	Qt__AccessibleTextRole        = Qt__ItemDataRole(11)
	Qt__AccessibleDescriptionRole = Qt__ItemDataRole(12)
	Qt__SizeHintRole              = Qt__ItemDataRole(13)
	Qt__InitialSortOrderRole      = Qt__ItemDataRole(14)
	Qt__DisplayPropertyRole       = Qt__ItemDataRole(27)
	Qt__DecorationPropertyRole    = Qt__ItemDataRole(28)
	Qt__ToolTipPropertyRole       = Qt__ItemDataRole(29)
	Qt__StatusTipPropertyRole     = Qt__ItemDataRole(30)
	Qt__WhatsThisPropertyRole     = Qt__ItemDataRole(31)
	Qt__UserRole                  = Qt__ItemDataRole(0x0100)
)

//Qt::ItemFlag
type Qt__ItemFlag int64

const (
	Qt__NoItemFlags          = Qt__ItemFlag(0)
	Qt__ItemIsSelectable     = Qt__ItemFlag(1)
	Qt__ItemIsEditable       = Qt__ItemFlag(2)
	Qt__ItemIsDragEnabled    = Qt__ItemFlag(4)
	Qt__ItemIsDropEnabled    = Qt__ItemFlag(8)
	Qt__ItemIsUserCheckable  = Qt__ItemFlag(16)
	Qt__ItemIsEnabled        = Qt__ItemFlag(32)
	Qt__ItemIsTristate       = Qt__ItemFlag(64)
	Qt__ItemNeverHasChildren = Qt__ItemFlag(128)
	Qt__ItemIsUserTristate   = Qt__ItemFlag(256)
)

//Qt::ItemSelectionMode
type Qt__ItemSelectionMode int64

const (
	Qt__ContainsItemShape          = Qt__ItemSelectionMode(0x0)
	Qt__IntersectsItemShape        = Qt__ItemSelectionMode(0x1)
	Qt__ContainsItemBoundingRect   = Qt__ItemSelectionMode(0x2)
	Qt__IntersectsItemBoundingRect = Qt__ItemSelectionMode(0x3)
)

//Qt::ItemSelectionOperation
type Qt__ItemSelectionOperation int64

const (
	Qt__ReplaceSelection = Qt__ItemSelectionOperation(0)
	Qt__AddToSelection   = Qt__ItemSelectionOperation(1)
)

//Qt::Key
type Qt__Key int64

const (
	Qt__Key_Escape                 = Qt__Key(0x01000000)
	Qt__Key_Tab                    = Qt__Key(0x01000001)
	Qt__Key_Backtab                = Qt__Key(0x01000002)
	Qt__Key_Backspace              = Qt__Key(0x01000003)
	Qt__Key_Return                 = Qt__Key(0x01000004)
	Qt__Key_Enter                  = Qt__Key(0x01000005)
	Qt__Key_Insert                 = Qt__Key(0x01000006)
	Qt__Key_Delete                 = Qt__Key(0x01000007)
	Qt__Key_Pause                  = Qt__Key(0x01000008)
	Qt__Key_Print                  = Qt__Key(0x01000009)
	Qt__Key_SysReq                 = Qt__Key(0x0100000a)
	Qt__Key_Clear                  = Qt__Key(0x0100000b)
	Qt__Key_Home                   = Qt__Key(0x01000010)
	Qt__Key_End                    = Qt__Key(0x01000011)
	Qt__Key_Left                   = Qt__Key(0x01000012)
	Qt__Key_Up                     = Qt__Key(0x01000013)
	Qt__Key_Right                  = Qt__Key(0x01000014)
	Qt__Key_Down                   = Qt__Key(0x01000015)
	Qt__Key_PageUp                 = Qt__Key(0x01000016)
	Qt__Key_PageDown               = Qt__Key(0x01000017)
	Qt__Key_Shift                  = Qt__Key(0x01000020)
	Qt__Key_Control                = Qt__Key(0x01000021)
	Qt__Key_Meta                   = Qt__Key(0x01000022)
	Qt__Key_Alt                    = Qt__Key(0x01000023)
	Qt__Key_CapsLock               = Qt__Key(0x01000024)
	Qt__Key_NumLock                = Qt__Key(0x01000025)
	Qt__Key_ScrollLock             = Qt__Key(0x01000026)
	Qt__Key_F1                     = Qt__Key(0x01000030)
	Qt__Key_F2                     = Qt__Key(0x01000031)
	Qt__Key_F3                     = Qt__Key(0x01000032)
	Qt__Key_F4                     = Qt__Key(0x01000033)
	Qt__Key_F5                     = Qt__Key(0x01000034)
	Qt__Key_F6                     = Qt__Key(0x01000035)
	Qt__Key_F7                     = Qt__Key(0x01000036)
	Qt__Key_F8                     = Qt__Key(0x01000037)
	Qt__Key_F9                     = Qt__Key(0x01000038)
	Qt__Key_F10                    = Qt__Key(0x01000039)
	Qt__Key_F11                    = Qt__Key(0x0100003a)
	Qt__Key_F12                    = Qt__Key(0x0100003b)
	Qt__Key_F13                    = Qt__Key(0x0100003c)
	Qt__Key_F14                    = Qt__Key(0x0100003d)
	Qt__Key_F15                    = Qt__Key(0x0100003e)
	Qt__Key_F16                    = Qt__Key(0x0100003f)
	Qt__Key_F17                    = Qt__Key(0x01000040)
	Qt__Key_F18                    = Qt__Key(0x01000041)
	Qt__Key_F19                    = Qt__Key(0x01000042)
	Qt__Key_F20                    = Qt__Key(0x01000043)
	Qt__Key_F21                    = Qt__Key(0x01000044)
	Qt__Key_F22                    = Qt__Key(0x01000045)
	Qt__Key_F23                    = Qt__Key(0x01000046)
	Qt__Key_F24                    = Qt__Key(0x01000047)
	Qt__Key_F25                    = Qt__Key(0x01000048)
	Qt__Key_F26                    = Qt__Key(0x01000049)
	Qt__Key_F27                    = Qt__Key(0x0100004a)
	Qt__Key_F28                    = Qt__Key(0x0100004b)
	Qt__Key_F29                    = Qt__Key(0x0100004c)
	Qt__Key_F30                    = Qt__Key(0x0100004d)
	Qt__Key_F31                    = Qt__Key(0x0100004e)
	Qt__Key_F32                    = Qt__Key(0x0100004f)
	Qt__Key_F33                    = Qt__Key(0x01000050)
	Qt__Key_F34                    = Qt__Key(0x01000051)
	Qt__Key_F35                    = Qt__Key(0x01000052)
	Qt__Key_Super_L                = Qt__Key(0x01000053)
	Qt__Key_Super_R                = Qt__Key(0x01000054)
	Qt__Key_Menu                   = Qt__Key(0x01000055)
	Qt__Key_Hyper_L                = Qt__Key(0x01000056)
	Qt__Key_Hyper_R                = Qt__Key(0x01000057)
	Qt__Key_Help                   = Qt__Key(0x01000058)
	Qt__Key_Direction_L            = Qt__Key(0x01000059)
	Qt__Key_Direction_R            = Qt__Key(0x01000060)
	Qt__Key_Space                  = Qt__Key(0x20)
	Qt__Key_Any                    = Qt__Key(Qt__Key_Space)
	Qt__Key_Exclam                 = Qt__Key(0x21)
	Qt__Key_QuoteDbl               = Qt__Key(0x22)
	Qt__Key_NumberSign             = Qt__Key(0x23)
	Qt__Key_Dollar                 = Qt__Key(0x24)
	Qt__Key_Percent                = Qt__Key(0x25)
	Qt__Key_Ampersand              = Qt__Key(0x26)
	Qt__Key_Apostrophe             = Qt__Key(0x27)
	Qt__Key_ParenLeft              = Qt__Key(0x28)
	Qt__Key_ParenRight             = Qt__Key(0x29)
	Qt__Key_Asterisk               = Qt__Key(0x2a)
	Qt__Key_Plus                   = Qt__Key(0x2b)
	Qt__Key_Comma                  = Qt__Key(0x2c)
	Qt__Key_Minus                  = Qt__Key(0x2d)
	Qt__Key_Period                 = Qt__Key(0x2e)
	Qt__Key_Slash                  = Qt__Key(0x2f)
	Qt__Key_0                      = Qt__Key(0x30)
	Qt__Key_1                      = Qt__Key(0x31)
	Qt__Key_2                      = Qt__Key(0x32)
	Qt__Key_3                      = Qt__Key(0x33)
	Qt__Key_4                      = Qt__Key(0x34)
	Qt__Key_5                      = Qt__Key(0x35)
	Qt__Key_6                      = Qt__Key(0x36)
	Qt__Key_7                      = Qt__Key(0x37)
	Qt__Key_8                      = Qt__Key(0x38)
	Qt__Key_9                      = Qt__Key(0x39)
	Qt__Key_Colon                  = Qt__Key(0x3a)
	Qt__Key_Semicolon              = Qt__Key(0x3b)
	Qt__Key_Less                   = Qt__Key(0x3c)
	Qt__Key_Equal                  = Qt__Key(0x3d)
	Qt__Key_Greater                = Qt__Key(0x3e)
	Qt__Key_Question               = Qt__Key(0x3f)
	Qt__Key_At                     = Qt__Key(0x40)
	Qt__Key_A                      = Qt__Key(0x41)
	Qt__Key_B                      = Qt__Key(0x42)
	Qt__Key_C                      = Qt__Key(0x43)
	Qt__Key_D                      = Qt__Key(0x44)
	Qt__Key_E                      = Qt__Key(0x45)
	Qt__Key_F                      = Qt__Key(0x46)
	Qt__Key_G                      = Qt__Key(0x47)
	Qt__Key_H                      = Qt__Key(0x48)
	Qt__Key_I                      = Qt__Key(0x49)
	Qt__Key_J                      = Qt__Key(0x4a)
	Qt__Key_K                      = Qt__Key(0x4b)
	Qt__Key_L                      = Qt__Key(0x4c)
	Qt__Key_M                      = Qt__Key(0x4d)
	Qt__Key_N                      = Qt__Key(0x4e)
	Qt__Key_O                      = Qt__Key(0x4f)
	Qt__Key_P                      = Qt__Key(0x50)
	Qt__Key_Q                      = Qt__Key(0x51)
	Qt__Key_R                      = Qt__Key(0x52)
	Qt__Key_S                      = Qt__Key(0x53)
	Qt__Key_T                      = Qt__Key(0x54)
	Qt__Key_U                      = Qt__Key(0x55)
	Qt__Key_V                      = Qt__Key(0x56)
	Qt__Key_W                      = Qt__Key(0x57)
	Qt__Key_X                      = Qt__Key(0x58)
	Qt__Key_Y                      = Qt__Key(0x59)
	Qt__Key_Z                      = Qt__Key(0x5a)
	Qt__Key_BracketLeft            = Qt__Key(0x5b)
	Qt__Key_Backslash              = Qt__Key(0x5c)
	Qt__Key_BracketRight           = Qt__Key(0x5d)
	Qt__Key_AsciiCircum            = Qt__Key(0x5e)
	Qt__Key_Underscore             = Qt__Key(0x5f)
	Qt__Key_QuoteLeft              = Qt__Key(0x60)
	Qt__Key_BraceLeft              = Qt__Key(0x7b)
	Qt__Key_Bar                    = Qt__Key(0x7c)
	Qt__Key_BraceRight             = Qt__Key(0x7d)
	Qt__Key_AsciiTilde             = Qt__Key(0x7e)
	Qt__Key_nobreakspace           = Qt__Key(0x0a0)
	Qt__Key_exclamdown             = Qt__Key(0x0a1)
	Qt__Key_cent                   = Qt__Key(0x0a2)
	Qt__Key_sterling               = Qt__Key(0x0a3)
	Qt__Key_currency               = Qt__Key(0x0a4)
	Qt__Key_yen                    = Qt__Key(0x0a5)
	Qt__Key_brokenbar              = Qt__Key(0x0a6)
	Qt__Key_section                = Qt__Key(0x0a7)
	Qt__Key_diaeresis              = Qt__Key(0x0a8)
	Qt__Key_copyright              = Qt__Key(0x0a9)
	Qt__Key_ordfeminine            = Qt__Key(0x0aa)
	Qt__Key_guillemotleft          = Qt__Key(0x0ab)
	Qt__Key_notsign                = Qt__Key(0x0ac)
	Qt__Key_hyphen                 = Qt__Key(0x0ad)
	Qt__Key_registered             = Qt__Key(0x0ae)
	Qt__Key_macron                 = Qt__Key(0x0af)
	Qt__Key_degree                 = Qt__Key(0x0b0)
	Qt__Key_plusminus              = Qt__Key(0x0b1)
	Qt__Key_twosuperior            = Qt__Key(0x0b2)
	Qt__Key_threesuperior          = Qt__Key(0x0b3)
	Qt__Key_acute                  = Qt__Key(0x0b4)
	Qt__Key_mu                     = Qt__Key(0x0b5)
	Qt__Key_paragraph              = Qt__Key(0x0b6)
	Qt__Key_periodcentered         = Qt__Key(0x0b7)
	Qt__Key_cedilla                = Qt__Key(0x0b8)
	Qt__Key_onesuperior            = Qt__Key(0x0b9)
	Qt__Key_masculine              = Qt__Key(0x0ba)
	Qt__Key_guillemotright         = Qt__Key(0x0bb)
	Qt__Key_onequarter             = Qt__Key(0x0bc)
	Qt__Key_onehalf                = Qt__Key(0x0bd)
	Qt__Key_threequarters          = Qt__Key(0x0be)
	Qt__Key_questiondown           = Qt__Key(0x0bf)
	Qt__Key_Agrave                 = Qt__Key(0x0c0)
	Qt__Key_Aacute                 = Qt__Key(0x0c1)
	Qt__Key_Acircumflex            = Qt__Key(0x0c2)
	Qt__Key_Atilde                 = Qt__Key(0x0c3)
	Qt__Key_Adiaeresis             = Qt__Key(0x0c4)
	Qt__Key_Aring                  = Qt__Key(0x0c5)
	Qt__Key_AE                     = Qt__Key(0x0c6)
	Qt__Key_Ccedilla               = Qt__Key(0x0c7)
	Qt__Key_Egrave                 = Qt__Key(0x0c8)
	Qt__Key_Eacute                 = Qt__Key(0x0c9)
	Qt__Key_Ecircumflex            = Qt__Key(0x0ca)
	Qt__Key_Ediaeresis             = Qt__Key(0x0cb)
	Qt__Key_Igrave                 = Qt__Key(0x0cc)
	Qt__Key_Iacute                 = Qt__Key(0x0cd)
	Qt__Key_Icircumflex            = Qt__Key(0x0ce)
	Qt__Key_Idiaeresis             = Qt__Key(0x0cf)
	Qt__Key_ETH                    = Qt__Key(0x0d0)
	Qt__Key_Ntilde                 = Qt__Key(0x0d1)
	Qt__Key_Ograve                 = Qt__Key(0x0d2)
	Qt__Key_Oacute                 = Qt__Key(0x0d3)
	Qt__Key_Ocircumflex            = Qt__Key(0x0d4)
	Qt__Key_Otilde                 = Qt__Key(0x0d5)
	Qt__Key_Odiaeresis             = Qt__Key(0x0d6)
	Qt__Key_multiply               = Qt__Key(0x0d7)
	Qt__Key_Ooblique               = Qt__Key(0x0d8)
	Qt__Key_Ugrave                 = Qt__Key(0x0d9)
	Qt__Key_Uacute                 = Qt__Key(0x0da)
	Qt__Key_Ucircumflex            = Qt__Key(0x0db)
	Qt__Key_Udiaeresis             = Qt__Key(0x0dc)
	Qt__Key_Yacute                 = Qt__Key(0x0dd)
	Qt__Key_THORN                  = Qt__Key(0x0de)
	Qt__Key_ssharp                 = Qt__Key(0x0df)
	Qt__Key_division               = Qt__Key(0x0f7)
	Qt__Key_ydiaeresis             = Qt__Key(0x0ff)
	Qt__Key_AltGr                  = Qt__Key(0x01001103)
	Qt__Key_Multi_key              = Qt__Key(0x01001120)
	Qt__Key_Codeinput              = Qt__Key(0x01001137)
	Qt__Key_SingleCandidate        = Qt__Key(0x0100113c)
	Qt__Key_MultipleCandidate      = Qt__Key(0x0100113d)
	Qt__Key_PreviousCandidate      = Qt__Key(0x0100113e)
	Qt__Key_Mode_switch            = Qt__Key(0x0100117e)
	Qt__Key_Kanji                  = Qt__Key(0x01001121)
	Qt__Key_Muhenkan               = Qt__Key(0x01001122)
	Qt__Key_Henkan                 = Qt__Key(0x01001123)
	Qt__Key_Romaji                 = Qt__Key(0x01001124)
	Qt__Key_Hiragana               = Qt__Key(0x01001125)
	Qt__Key_Katakana               = Qt__Key(0x01001126)
	Qt__Key_Hiragana_Katakana      = Qt__Key(0x01001127)
	Qt__Key_Zenkaku                = Qt__Key(0x01001128)
	Qt__Key_Hankaku                = Qt__Key(0x01001129)
	Qt__Key_Zenkaku_Hankaku        = Qt__Key(0x0100112a)
	Qt__Key_Touroku                = Qt__Key(0x0100112b)
	Qt__Key_Massyo                 = Qt__Key(0x0100112c)
	Qt__Key_Kana_Lock              = Qt__Key(0x0100112d)
	Qt__Key_Kana_Shift             = Qt__Key(0x0100112e)
	Qt__Key_Eisu_Shift             = Qt__Key(0x0100112f)
	Qt__Key_Eisu_toggle            = Qt__Key(0x01001130)
	Qt__Key_Hangul                 = Qt__Key(0x01001131)
	Qt__Key_Hangul_Start           = Qt__Key(0x01001132)
	Qt__Key_Hangul_End             = Qt__Key(0x01001133)
	Qt__Key_Hangul_Hanja           = Qt__Key(0x01001134)
	Qt__Key_Hangul_Jamo            = Qt__Key(0x01001135)
	Qt__Key_Hangul_Romaja          = Qt__Key(0x01001136)
	Qt__Key_Hangul_Jeonja          = Qt__Key(0x01001138)
	Qt__Key_Hangul_Banja           = Qt__Key(0x01001139)
	Qt__Key_Hangul_PreHanja        = Qt__Key(0x0100113a)
	Qt__Key_Hangul_PostHanja       = Qt__Key(0x0100113b)
	Qt__Key_Hangul_Special         = Qt__Key(0x0100113f)
	Qt__Key_Dead_Grave             = Qt__Key(0x01001250)
	Qt__Key_Dead_Acute             = Qt__Key(0x01001251)
	Qt__Key_Dead_Circumflex        = Qt__Key(0x01001252)
	Qt__Key_Dead_Tilde             = Qt__Key(0x01001253)
	Qt__Key_Dead_Macron            = Qt__Key(0x01001254)
	Qt__Key_Dead_Breve             = Qt__Key(0x01001255)
	Qt__Key_Dead_Abovedot          = Qt__Key(0x01001256)
	Qt__Key_Dead_Diaeresis         = Qt__Key(0x01001257)
	Qt__Key_Dead_Abovering         = Qt__Key(0x01001258)
	Qt__Key_Dead_Doubleacute       = Qt__Key(0x01001259)
	Qt__Key_Dead_Caron             = Qt__Key(0x0100125a)
	Qt__Key_Dead_Cedilla           = Qt__Key(0x0100125b)
	Qt__Key_Dead_Ogonek            = Qt__Key(0x0100125c)
	Qt__Key_Dead_Iota              = Qt__Key(0x0100125d)
	Qt__Key_Dead_Voiced_Sound      = Qt__Key(0x0100125e)
	Qt__Key_Dead_Semivoiced_Sound  = Qt__Key(0x0100125f)
	Qt__Key_Dead_Belowdot          = Qt__Key(0x01001260)
	Qt__Key_Dead_Hook              = Qt__Key(0x01001261)
	Qt__Key_Dead_Horn              = Qt__Key(0x01001262)
	Qt__Key_Back                   = Qt__Key(0x01000061)
	Qt__Key_Forward                = Qt__Key(0x01000062)
	Qt__Key_Stop                   = Qt__Key(0x01000063)
	Qt__Key_Refresh                = Qt__Key(0x01000064)
	Qt__Key_VolumeDown             = Qt__Key(0x01000070)
	Qt__Key_VolumeMute             = Qt__Key(0x01000071)
	Qt__Key_VolumeUp               = Qt__Key(0x01000072)
	Qt__Key_BassBoost              = Qt__Key(0x01000073)
	Qt__Key_BassUp                 = Qt__Key(0x01000074)
	Qt__Key_BassDown               = Qt__Key(0x01000075)
	Qt__Key_TrebleUp               = Qt__Key(0x01000076)
	Qt__Key_TrebleDown             = Qt__Key(0x01000077)
	Qt__Key_MediaPlay              = Qt__Key(0x01000080)
	Qt__Key_MediaStop              = Qt__Key(0x01000081)
	Qt__Key_MediaPrevious          = Qt__Key(0x01000082)
	Qt__Key_MediaNext              = Qt__Key(0x01000083)
	Qt__Key_MediaRecord            = Qt__Key(0x01000084)
	Qt__Key_MediaPause             = Qt__Key(0x1000085)
	Qt__Key_MediaTogglePlayPause   = Qt__Key(0x1000086)
	Qt__Key_HomePage               = Qt__Key(0x01000090)
	Qt__Key_Favorites              = Qt__Key(0x01000091)
	Qt__Key_Search                 = Qt__Key(0x01000092)
	Qt__Key_Standby                = Qt__Key(0x01000093)
	Qt__Key_OpenUrl                = Qt__Key(0x01000094)
	Qt__Key_LaunchMail             = Qt__Key(0x010000a0)
	Qt__Key_LaunchMedia            = Qt__Key(0x010000a1)
	Qt__Key_Launch0                = Qt__Key(0x010000a2)
	Qt__Key_Launch1                = Qt__Key(0x010000a3)
	Qt__Key_Launch2                = Qt__Key(0x010000a4)
	Qt__Key_Launch3                = Qt__Key(0x010000a5)
	Qt__Key_Launch4                = Qt__Key(0x010000a6)
	Qt__Key_Launch5                = Qt__Key(0x010000a7)
	Qt__Key_Launch6                = Qt__Key(0x010000a8)
	Qt__Key_Launch7                = Qt__Key(0x010000a9)
	Qt__Key_Launch8                = Qt__Key(0x010000aa)
	Qt__Key_Launch9                = Qt__Key(0x010000ab)
	Qt__Key_LaunchA                = Qt__Key(0x010000ac)
	Qt__Key_LaunchB                = Qt__Key(0x010000ad)
	Qt__Key_LaunchC                = Qt__Key(0x010000ae)
	Qt__Key_LaunchD                = Qt__Key(0x010000af)
	Qt__Key_LaunchE                = Qt__Key(0x010000b0)
	Qt__Key_LaunchF                = Qt__Key(0x010000b1)
	Qt__Key_MonBrightnessUp        = Qt__Key(0x010000b2)
	Qt__Key_MonBrightnessDown      = Qt__Key(0x010000b3)
	Qt__Key_KeyboardLightOnOff     = Qt__Key(0x010000b4)
	Qt__Key_KeyboardBrightnessUp   = Qt__Key(0x010000b5)
	Qt__Key_KeyboardBrightnessDown = Qt__Key(0x010000b6)
	Qt__Key_PowerOff               = Qt__Key(0x010000b7)
	Qt__Key_WakeUp                 = Qt__Key(0x010000b8)
	Qt__Key_Eject                  = Qt__Key(0x010000b9)
	Qt__Key_ScreenSaver            = Qt__Key(0x010000ba)
	Qt__Key_WWW                    = Qt__Key(0x010000bb)
	Qt__Key_Memo                   = Qt__Key(0x010000bc)
	Qt__Key_LightBulb              = Qt__Key(0x010000bd)
	Qt__Key_Shop                   = Qt__Key(0x010000be)
	Qt__Key_History                = Qt__Key(0x010000bf)
	Qt__Key_AddFavorite            = Qt__Key(0x010000c0)
	Qt__Key_HotLinks               = Qt__Key(0x010000c1)
	Qt__Key_BrightnessAdjust       = Qt__Key(0x010000c2)
	Qt__Key_Finance                = Qt__Key(0x010000c3)
	Qt__Key_Community              = Qt__Key(0x010000c4)
	Qt__Key_AudioRewind            = Qt__Key(0x010000c5)
	Qt__Key_BackForward            = Qt__Key(0x010000c6)
	Qt__Key_ApplicationLeft        = Qt__Key(0x010000c7)
	Qt__Key_ApplicationRight       = Qt__Key(0x010000c8)
	Qt__Key_Book                   = Qt__Key(0x010000c9)
	Qt__Key_CD                     = Qt__Key(0x010000ca)
	Qt__Key_Calculator             = Qt__Key(0x010000cb)
	Qt__Key_ToDoList               = Qt__Key(0x010000cc)
	Qt__Key_ClearGrab              = Qt__Key(0x010000cd)
	Qt__Key_Close                  = Qt__Key(0x010000ce)
	Qt__Key_Copy                   = Qt__Key(0x010000cf)
	Qt__Key_Cut                    = Qt__Key(0x010000d0)
	Qt__Key_Display                = Qt__Key(0x010000d1)
	Qt__Key_DOS                    = Qt__Key(0x010000d2)
	Qt__Key_Documents              = Qt__Key(0x010000d3)
	Qt__Key_Excel                  = Qt__Key(0x010000d4)
	Qt__Key_Explorer               = Qt__Key(0x010000d5)
	Qt__Key_Game                   = Qt__Key(0x010000d6)
	Qt__Key_Go                     = Qt__Key(0x010000d7)
	Qt__Key_iTouch                 = Qt__Key(0x010000d8)
	Qt__Key_LogOff                 = Qt__Key(0x010000d9)
	Qt__Key_Market                 = Qt__Key(0x010000da)
	Qt__Key_Meeting                = Qt__Key(0x010000db)
	Qt__Key_MenuKB                 = Qt__Key(0x010000dc)
	Qt__Key_MenuPB                 = Qt__Key(0x010000dd)
	Qt__Key_MySites                = Qt__Key(0x010000de)
	Qt__Key_News                   = Qt__Key(0x010000df)
	Qt__Key_OfficeHome             = Qt__Key(0x010000e0)
	Qt__Key_Option                 = Qt__Key(0x010000e1)
	Qt__Key_Paste                  = Qt__Key(0x010000e2)
	Qt__Key_Phone                  = Qt__Key(0x010000e3)
	Qt__Key_Calendar               = Qt__Key(0x010000e4)
	Qt__Key_Reply                  = Qt__Key(0x010000e5)
	Qt__Key_Reload                 = Qt__Key(0x010000e6)
	Qt__Key_RotateWindows          = Qt__Key(0x010000e7)
	Qt__Key_RotationPB             = Qt__Key(0x010000e8)
	Qt__Key_RotationKB             = Qt__Key(0x010000e9)
	Qt__Key_Save                   = Qt__Key(0x010000ea)
	Qt__Key_Send                   = Qt__Key(0x010000eb)
	Qt__Key_Spell                  = Qt__Key(0x010000ec)
	Qt__Key_SplitScreen            = Qt__Key(0x010000ed)
	Qt__Key_Support                = Qt__Key(0x010000ee)
	Qt__Key_TaskPane               = Qt__Key(0x010000ef)
	Qt__Key_Terminal               = Qt__Key(0x010000f0)
	Qt__Key_Tools                  = Qt__Key(0x010000f1)
	Qt__Key_Travel                 = Qt__Key(0x010000f2)
	Qt__Key_Video                  = Qt__Key(0x010000f3)
	Qt__Key_Word                   = Qt__Key(0x010000f4)
	Qt__Key_Xfer                   = Qt__Key(0x010000f5)
	Qt__Key_ZoomIn                 = Qt__Key(0x010000f6)
	Qt__Key_ZoomOut                = Qt__Key(0x010000f7)
	Qt__Key_Away                   = Qt__Key(0x010000f8)
	Qt__Key_Messenger              = Qt__Key(0x010000f9)
	Qt__Key_WebCam                 = Qt__Key(0x010000fa)
	Qt__Key_MailForward            = Qt__Key(0x010000fb)
	Qt__Key_Pictures               = Qt__Key(0x010000fc)
	Qt__Key_Music                  = Qt__Key(0x010000fd)
	Qt__Key_Battery                = Qt__Key(0x010000fe)
	Qt__Key_Bluetooth              = Qt__Key(0x010000ff)
	Qt__Key_WLAN                   = Qt__Key(0x01000100)
	Qt__Key_UWB                    = Qt__Key(0x01000101)
	Qt__Key_AudioForward           = Qt__Key(0x01000102)
	Qt__Key_AudioRepeat            = Qt__Key(0x01000103)
	Qt__Key_AudioRandomPlay        = Qt__Key(0x01000104)
	Qt__Key_Subtitle               = Qt__Key(0x01000105)
	Qt__Key_AudioCycleTrack        = Qt__Key(0x01000106)
	Qt__Key_Time                   = Qt__Key(0x01000107)
	Qt__Key_Hibernate              = Qt__Key(0x01000108)
	Qt__Key_View                   = Qt__Key(0x01000109)
	Qt__Key_TopMenu                = Qt__Key(0x0100010a)
	Qt__Key_PowerDown              = Qt__Key(0x0100010b)
	Qt__Key_Suspend                = Qt__Key(0x0100010c)
	Qt__Key_ContrastAdjust         = Qt__Key(0x0100010d)
	Qt__Key_LaunchG                = Qt__Key(0x0100010e)
	Qt__Key_LaunchH                = Qt__Key(0x0100010f)
	Qt__Key_TouchpadToggle         = Qt__Key(0x01000110)
	Qt__Key_TouchpadOn             = Qt__Key(0x01000111)
	Qt__Key_TouchpadOff            = Qt__Key(0x01000112)
	Qt__Key_MicMute                = Qt__Key(0x01000113)
	Qt__Key_Red                    = Qt__Key(0x01000114)
	Qt__Key_Green                  = Qt__Key(0x01000115)
	Qt__Key_Yellow                 = Qt__Key(0x01000116)
	Qt__Key_Blue                   = Qt__Key(0x01000117)
	Qt__Key_ChannelUp              = Qt__Key(0x01000118)
	Qt__Key_ChannelDown            = Qt__Key(0x01000119)
	Qt__Key_Guide                  = Qt__Key(0x0100011a)
	Qt__Key_Info                   = Qt__Key(0x0100011b)
	Qt__Key_Settings               = Qt__Key(0x0100011c)
	Qt__Key_MicVolumeUp            = Qt__Key(0x0100011d)
	Qt__Key_MicVolumeDown          = Qt__Key(0x0100011e)
	Qt__Key_New                    = Qt__Key(0x01000120)
	Qt__Key_Open                   = Qt__Key(0x01000121)
	Qt__Key_Find                   = Qt__Key(0x01000122)
	Qt__Key_Undo                   = Qt__Key(0x01000123)
	Qt__Key_Redo                   = Qt__Key(0x01000124)
	Qt__Key_MediaLast              = Qt__Key(0x0100ffff)
	Qt__Key_Select                 = Qt__Key(0x01010000)
	Qt__Key_Yes                    = Qt__Key(0x01010001)
	Qt__Key_No                     = Qt__Key(0x01010002)
	Qt__Key_Cancel                 = Qt__Key(0x01020001)
	Qt__Key_Printer                = Qt__Key(0x01020002)
	Qt__Key_Execute                = Qt__Key(0x01020003)
	Qt__Key_Sleep                  = Qt__Key(0x01020004)
	Qt__Key_Play                   = Qt__Key(0x01020005)
	Qt__Key_Zoom                   = Qt__Key(0x01020006)
	Qt__Key_Exit                   = Qt__Key(0x0102000a)
	Qt__Key_Context1               = Qt__Key(0x01100000)
	Qt__Key_Context2               = Qt__Key(0x01100001)
	Qt__Key_Context3               = Qt__Key(0x01100002)
	Qt__Key_Context4               = Qt__Key(0x01100003)
	Qt__Key_Call                   = Qt__Key(0x01100004)
	Qt__Key_Hangup                 = Qt__Key(0x01100005)
	Qt__Key_Flip                   = Qt__Key(0x01100006)
	Qt__Key_ToggleCallHangup       = Qt__Key(0x01100007)
	Qt__Key_VoiceDial              = Qt__Key(0x01100008)
	Qt__Key_LastNumberRedial       = Qt__Key(0x01100009)
	Qt__Key_Camera                 = Qt__Key(0x01100020)
	Qt__Key_CameraFocus            = Qt__Key(0x01100021)
	Qt__Key_unknown                = Qt__Key(0x01ffffff)
)

//Qt::KeyboardModifier
type Qt__KeyboardModifier int64

const (
	Qt__NoModifier           = Qt__KeyboardModifier(0x00000000)
	Qt__ShiftModifier        = Qt__KeyboardModifier(0x02000000)
	Qt__ControlModifier      = Qt__KeyboardModifier(0x04000000)
	Qt__AltModifier          = Qt__KeyboardModifier(0x08000000)
	Qt__MetaModifier         = Qt__KeyboardModifier(0x10000000)
	Qt__KeypadModifier       = Qt__KeyboardModifier(0x20000000)
	Qt__GroupSwitchModifier  = Qt__KeyboardModifier(0x40000000)
	Qt__KeyboardModifierMask = Qt__KeyboardModifier(0xfe000000)
)

//Qt::LayoutDirection
type Qt__LayoutDirection int64

const (
	Qt__LeftToRight         = Qt__LayoutDirection(0)
	Qt__RightToLeft         = Qt__LayoutDirection(1)
	Qt__LayoutDirectionAuto = Qt__LayoutDirection(2)
)

//Qt::MaskMode
type Qt__MaskMode int64

const (
	Qt__MaskInColor  = Qt__MaskMode(0)
	Qt__MaskOutColor = Qt__MaskMode(1)
)

//Qt::MatchFlag
type Qt__MatchFlag int64

const (
	Qt__MatchExactly       = Qt__MatchFlag(0)
	Qt__MatchContains      = Qt__MatchFlag(1)
	Qt__MatchStartsWith    = Qt__MatchFlag(2)
	Qt__MatchEndsWith      = Qt__MatchFlag(3)
	Qt__MatchRegExp        = Qt__MatchFlag(4)
	Qt__MatchWildcard      = Qt__MatchFlag(5)
	Qt__MatchFixedString   = Qt__MatchFlag(8)
	Qt__MatchCaseSensitive = Qt__MatchFlag(16)
	Qt__MatchWrap          = Qt__MatchFlag(32)
	Qt__MatchRecursive     = Qt__MatchFlag(64)
)

//Qt::Modifier
type Qt__Modifier int64

const (
	Qt__META          = Qt__Modifier(Qt__MetaModifier)
	Qt__SHIFT         = Qt__Modifier(Qt__ShiftModifier)
	Qt__CTRL          = Qt__Modifier(Qt__ControlModifier)
	Qt__ALT           = Qt__Modifier(Qt__AltModifier)
	Qt__MODIFIER_MASK = Qt__Modifier(Qt__KeyboardModifierMask)
	Qt__UNICODE_ACCEL = Qt__Modifier(0x00000000)
)

//Qt::MouseButton
type Qt__MouseButton int64

const (
	Qt__NoButton        = Qt__MouseButton(0x00000000)
	Qt__LeftButton      = Qt__MouseButton(0x00000001)
	Qt__RightButton     = Qt__MouseButton(0x00000002)
	Qt__MidButton       = Qt__MouseButton(0x00000004)
	Qt__MiddleButton    = Qt__MouseButton(Qt__MidButton)
	Qt__BackButton      = Qt__MouseButton(0x00000008)
	Qt__XButton1        = Qt__MouseButton(Qt__BackButton)
	Qt__ExtraButton1    = Qt__MouseButton(Qt__XButton1)
	Qt__ForwardButton   = Qt__MouseButton(0x00000010)
	Qt__XButton2        = Qt__MouseButton(Qt__ForwardButton)
	Qt__ExtraButton2    = Qt__MouseButton(Qt__ForwardButton)
	Qt__TaskButton      = Qt__MouseButton(0x00000020)
	Qt__ExtraButton3    = Qt__MouseButton(Qt__TaskButton)
	Qt__ExtraButton4    = Qt__MouseButton(0x00000040)
	Qt__ExtraButton5    = Qt__MouseButton(0x00000080)
	Qt__ExtraButton6    = Qt__MouseButton(0x00000100)
	Qt__ExtraButton7    = Qt__MouseButton(0x00000200)
	Qt__ExtraButton8    = Qt__MouseButton(0x00000400)
	Qt__ExtraButton9    = Qt__MouseButton(0x00000800)
	Qt__ExtraButton10   = Qt__MouseButton(0x00001000)
	Qt__ExtraButton11   = Qt__MouseButton(0x00002000)
	Qt__ExtraButton12   = Qt__MouseButton(0x00004000)
	Qt__ExtraButton13   = Qt__MouseButton(0x00008000)
	Qt__ExtraButton14   = Qt__MouseButton(0x00010000)
	Qt__ExtraButton15   = Qt__MouseButton(0x00020000)
	Qt__ExtraButton16   = Qt__MouseButton(0x00040000)
	Qt__ExtraButton17   = Qt__MouseButton(0x00080000)
	Qt__ExtraButton18   = Qt__MouseButton(0x00100000)
	Qt__ExtraButton19   = Qt__MouseButton(0x00200000)
	Qt__ExtraButton20   = Qt__MouseButton(0x00400000)
	Qt__ExtraButton21   = Qt__MouseButton(0x00800000)
	Qt__ExtraButton22   = Qt__MouseButton(0x01000000)
	Qt__ExtraButton23   = Qt__MouseButton(0x02000000)
	Qt__ExtraButton24   = Qt__MouseButton(0x04000000)
	Qt__AllButtons      = Qt__MouseButton(0x07ffffff)
	Qt__MaxMouseButton  = Qt__MouseButton(Qt__ExtraButton24)
	Qt__MouseButtonMask = Qt__MouseButton(0xffffffff)
)

//Qt::MouseEventFlag
type Qt__MouseEventFlag int64

const (
	Qt__MouseEventCreatedDoubleClick = Qt__MouseEventFlag(0x01)
	Qt__MouseEventFlagMask           = Qt__MouseEventFlag(0xFF)
)

//Qt::MouseEventSource
type Qt__MouseEventSource int64

const (
	Qt__MouseEventNotSynthesized      = Qt__MouseEventSource(0)
	Qt__MouseEventSynthesizedBySystem = Qt__MouseEventSource(1)
	Qt__MouseEventSynthesizedByQt     = Qt__MouseEventSource(2)
)

//Qt::NativeGestureType
type Qt__NativeGestureType int64

const (
	Qt__BeginNativeGesture     = Qt__NativeGestureType(0)
	Qt__EndNativeGesture       = Qt__NativeGestureType(1)
	Qt__PanNativeGesture       = Qt__NativeGestureType(2)
	Qt__ZoomNativeGesture      = Qt__NativeGestureType(3)
	Qt__SmartZoomNativeGesture = Qt__NativeGestureType(4)
	Qt__RotateNativeGesture    = Qt__NativeGestureType(5)
	Qt__SwipeNativeGesture     = Qt__NativeGestureType(6)
)

//Qt::NavigationMode
type Qt__NavigationMode int64

const (
	Qt__NavigationModeNone               = Qt__NavigationMode(0)
	Qt__NavigationModeKeypadTabOrder     = Qt__NavigationMode(1)
	Qt__NavigationModeKeypadDirectional  = Qt__NavigationMode(2)
	Qt__NavigationModeCursorAuto         = Qt__NavigationMode(3)
	Qt__NavigationModeCursorForceVisible = Qt__NavigationMode(4)
)

//Qt::Orientation
type Qt__Orientation int64

const (
	Qt__Horizontal = Qt__Orientation(0x1)
	Qt__Vertical   = Qt__Orientation(0x2)
)

//Qt::PenCapStyle
type Qt__PenCapStyle int64

var (
	Qt__FlatCap      = Qt__PenCapStyle(0x00)
	Qt__SquareCap    = Qt__PenCapStyle(0x10)
	Qt__RoundCap     = Qt__PenCapStyle(0x20)
	Qt__MPenCapStyle = Qt__PenCapStyle(0x30)
)

//Qt::PenJoinStyle
type Qt__PenJoinStyle int64

var (
	Qt__MiterJoin     = Qt__PenJoinStyle(0x00)
	Qt__BevelJoin     = Qt__PenJoinStyle(0x40)
	Qt__RoundJoin     = Qt__PenJoinStyle(0x80)
	Qt__SvgMiterJoin  = Qt__PenJoinStyle(0x100)
	Qt__MPenJoinStyle = Qt__PenJoinStyle(0x1c0)
)

//Qt::PenStyle
type Qt__PenStyle int64

var (
	Qt__NoPen          = Qt__PenStyle(0)
	Qt__SolidLine      = Qt__PenStyle(1)
	Qt__DashLine       = Qt__PenStyle(2)
	Qt__DotLine        = Qt__PenStyle(3)
	Qt__DashDotLine    = Qt__PenStyle(4)
	Qt__DashDotDotLine = Qt__PenStyle(5)
	Qt__CustomDashLine = Qt__PenStyle(6)
	Qt__MPenStyle      = Qt__PenStyle(0x0f)
)

//Qt::ScreenOrientation
type Qt__ScreenOrientation int64

const (
	Qt__PrimaryOrientation           = Qt__ScreenOrientation(0x00000000)
	Qt__PortraitOrientation          = Qt__ScreenOrientation(0x00000001)
	Qt__LandscapeOrientation         = Qt__ScreenOrientation(0x00000002)
	Qt__InvertedPortraitOrientation  = Qt__ScreenOrientation(0x00000004)
	Qt__InvertedLandscapeOrientation = Qt__ScreenOrientation(0x00000008)
)

//Qt::ScrollBarPolicy
type Qt__ScrollBarPolicy int64

const (
	Qt__ScrollBarAsNeeded  = Qt__ScrollBarPolicy(0)
	Qt__ScrollBarAlwaysOff = Qt__ScrollBarPolicy(1)
	Qt__ScrollBarAlwaysOn  = Qt__ScrollBarPolicy(2)
)

//Qt::ScrollPhase
type Qt__ScrollPhase int64

const (
	Qt__ScrollBegin  = Qt__ScrollPhase(1)
	Qt__ScrollUpdate = Qt__ScrollPhase(2)
	Qt__ScrollEnd    = Qt__ScrollPhase(3)
)

//Qt::ShortcutContext
type Qt__ShortcutContext int64

const (
	Qt__WidgetShortcut             = Qt__ShortcutContext(0)
	Qt__WindowShortcut             = Qt__ShortcutContext(1)
	Qt__ApplicationShortcut        = Qt__ShortcutContext(2)
	Qt__WidgetWithChildrenShortcut = Qt__ShortcutContext(3)
)

//Qt::SizeHint
type Qt__SizeHint int64

const (
	Qt__MinimumSize    = Qt__SizeHint(0)
	Qt__PreferredSize  = Qt__SizeHint(1)
	Qt__MaximumSize    = Qt__SizeHint(2)
	Qt__MinimumDescent = Qt__SizeHint(3)
	Qt__NSizeHints     = Qt__SizeHint(4)
)

//Qt::SizeMode
type Qt__SizeMode int64

const (
	Qt__AbsoluteSize = Qt__SizeMode(0)
	Qt__RelativeSize = Qt__SizeMode(1)
)

//Qt::SortOrder
type Qt__SortOrder int64

const (
	Qt__AscendingOrder  = Qt__SortOrder(0)
	Qt__DescendingOrder = Qt__SortOrder(1)
)

//Qt::TabFocusBehavior
type Qt__TabFocusBehavior int64

const (
	Qt__NoTabFocus           = Qt__TabFocusBehavior(0x00)
	Qt__TabFocusTextControls = Qt__TabFocusBehavior(0x01)
	Qt__TabFocusListControls = Qt__TabFocusBehavior(0x02)
	Qt__TabFocusAllControls  = Qt__TabFocusBehavior(0xff)
)

//Qt::TextElideMode
type Qt__TextElideMode int64

const (
	Qt__ElideLeft   = Qt__TextElideMode(0)
	Qt__ElideRight  = Qt__TextElideMode(1)
	Qt__ElideMiddle = Qt__TextElideMode(2)
	Qt__ElideNone   = Qt__TextElideMode(3)
)

//Qt::TextFlag
type Qt__TextFlag int64

const (
	Qt__TextSingleLine            = Qt__TextFlag(0x0100)
	Qt__TextDontClip              = Qt__TextFlag(0x0200)
	Qt__TextExpandTabs            = Qt__TextFlag(0x0400)
	Qt__TextShowMnemonic          = Qt__TextFlag(0x0800)
	Qt__TextWordWrap              = Qt__TextFlag(0x1000)
	Qt__TextWrapAnywhere          = Qt__TextFlag(0x2000)
	Qt__TextDontPrint             = Qt__TextFlag(0x4000)
	Qt__TextIncludeTrailingSpaces = Qt__TextFlag(0x08000000)
	Qt__TextHideMnemonic          = Qt__TextFlag(0x8000)
	Qt__TextJustificationForced   = Qt__TextFlag(0x10000)
	Qt__TextForceLeftToRight      = Qt__TextFlag(0x20000)
	Qt__TextForceRightToLeft      = Qt__TextFlag(0x40000)
	Qt__TextLongestVariant        = Qt__TextFlag(0x80000)
	Qt__TextBypassShaping         = Qt__TextFlag(0x100000)
)

//Qt::TextFormat
type Qt__TextFormat int64

const (
	Qt__PlainText = Qt__TextFormat(0)
	Qt__RichText  = Qt__TextFormat(1)
	Qt__AutoText  = Qt__TextFormat(2)
)

//Qt::TextInteractionFlag
type Qt__TextInteractionFlag int64

const (
	Qt__NoTextInteraction         = Qt__TextInteractionFlag(0)
	Qt__TextSelectableByMouse     = Qt__TextInteractionFlag(1)
	Qt__TextSelectableByKeyboard  = Qt__TextInteractionFlag(2)
	Qt__LinksAccessibleByMouse    = Qt__TextInteractionFlag(4)
	Qt__LinksAccessibleByKeyboard = Qt__TextInteractionFlag(8)
	Qt__TextEditable              = Qt__TextInteractionFlag(16)
	Qt__TextEditorInteraction     = Qt__TextInteractionFlag(Qt__TextSelectableByMouse | Qt__TextSelectableByKeyboard | Qt__TextEditable)
	Qt__TextBrowserInteraction    = Qt__TextInteractionFlag(Qt__TextSelectableByMouse | Qt__LinksAccessibleByMouse | Qt__LinksAccessibleByKeyboard)
)

//Qt::TileRule
type Qt__TileRule int64

const (
	Qt__StretchTile = Qt__TileRule(0)
	Qt__RepeatTile  = Qt__TileRule(1)
	Qt__RoundTile   = Qt__TileRule(2)
)

//Qt::TimeSpec
type Qt__TimeSpec int64

const (
	Qt__LocalTime     = Qt__TimeSpec(0)
	Qt__UTC           = Qt__TimeSpec(1)
	Qt__OffsetFromUTC = Qt__TimeSpec(2)
	Qt__TimeZone      = Qt__TimeSpec(3)
)

//Qt::TimerType
type Qt__TimerType int64

const (
	Qt__PreciseTimer    = Qt__TimerType(0)
	Qt__CoarseTimer     = Qt__TimerType(1)
	Qt__VeryCoarseTimer = Qt__TimerType(2)
)

//Qt::ToolBarArea
type Qt__ToolBarArea int64

const (
	Qt__LeftToolBarArea   = Qt__ToolBarArea(0x1)
	Qt__RightToolBarArea  = Qt__ToolBarArea(0x2)
	Qt__TopToolBarArea    = Qt__ToolBarArea(0x4)
	Qt__BottomToolBarArea = Qt__ToolBarArea(0x8)
	Qt__ToolBarArea_Mask  = Qt__ToolBarArea(0xf)
	Qt__AllToolBarAreas   = Qt__ToolBarArea(Qt__ToolBarArea_Mask)
	Qt__NoToolBarArea     = Qt__ToolBarArea(0)
)

//Qt::ToolButtonStyle
type Qt__ToolButtonStyle int64

var (
	Qt__ToolButtonIconOnly       = Qt__ToolButtonStyle(0)
	Qt__ToolButtonTextOnly       = Qt__ToolButtonStyle(1)
	Qt__ToolButtonTextBesideIcon = Qt__ToolButtonStyle(2)
	Qt__ToolButtonTextUnderIcon  = Qt__ToolButtonStyle(3)
	Qt__ToolButtonFollowStyle    = Qt__ToolButtonStyle(4)
)

//Qt::TouchPointState
type Qt__TouchPointState int64

const (
	Qt__TouchPointPressed    = Qt__TouchPointState(0x01)
	Qt__TouchPointMoved      = Qt__TouchPointState(0x02)
	Qt__TouchPointStationary = Qt__TouchPointState(0x04)
	Qt__TouchPointReleased   = Qt__TouchPointState(0x08)
)

//Qt::TransformationMode
type Qt__TransformationMode int64

const (
	Qt__FastTransformation   = Qt__TransformationMode(0)
	Qt__SmoothTransformation = Qt__TransformationMode(1)
)

//Qt::UIEffect
type Qt__UIEffect int64

const (
	Qt__UI_General        = Qt__UIEffect(0)
	Qt__UI_AnimateMenu    = Qt__UIEffect(1)
	Qt__UI_FadeMenu       = Qt__UIEffect(2)
	Qt__UI_AnimateCombo   = Qt__UIEffect(3)
	Qt__UI_AnimateTooltip = Qt__UIEffect(4)
	Qt__UI_FadeTooltip    = Qt__UIEffect(5)
	Qt__UI_AnimateToolBox = Qt__UIEffect(6)
)

//Qt::WhiteSpaceMode
type Qt__WhiteSpaceMode int64

const (
	Qt__WhiteSpaceNormal        = Qt__WhiteSpaceMode(0)
	Qt__WhiteSpacePre           = Qt__WhiteSpaceMode(1)
	Qt__WhiteSpaceNoWrap        = Qt__WhiteSpaceMode(2)
	Qt__WhiteSpaceModeUndefined = Qt__WhiteSpaceMode(-1)
)

//Qt::WidgetAttribute
type Qt__WidgetAttribute int64

const (
	Qt__WA_Disabled                        = Qt__WidgetAttribute(0)
	Qt__WA_UnderMouse                      = Qt__WidgetAttribute(1)
	Qt__WA_MouseTracking                   = Qt__WidgetAttribute(2)
	Qt__WA_ContentsPropagated              = Qt__WidgetAttribute(3)
	Qt__WA_OpaquePaintEvent                = Qt__WidgetAttribute(4)
	Qt__WA_NoBackground                    = Qt__WidgetAttribute(Qt__WA_OpaquePaintEvent)
	Qt__WA_StaticContents                  = Qt__WidgetAttribute(5)
	Qt__WA_LaidOut                         = Qt__WidgetAttribute(7)
	Qt__WA_PaintOnScreen                   = Qt__WidgetAttribute(8)
	Qt__WA_NoSystemBackground              = Qt__WidgetAttribute(9)
	Qt__WA_UpdatesDisabled                 = Qt__WidgetAttribute(10)
	Qt__WA_Mapped                          = Qt__WidgetAttribute(11)
	Qt__WA_MacNoClickThrough               = Qt__WidgetAttribute(12)
	Qt__WA_InputMethodEnabled              = Qt__WidgetAttribute(14)
	Qt__WA_WState_Visible                  = Qt__WidgetAttribute(15)
	Qt__WA_WState_Hidden                   = Qt__WidgetAttribute(16)
	Qt__WA_ForceDisabled                   = Qt__WidgetAttribute(32)
	Qt__WA_KeyCompression                  = Qt__WidgetAttribute(33)
	Qt__WA_PendingMoveEvent                = Qt__WidgetAttribute(34)
	Qt__WA_PendingResizeEvent              = Qt__WidgetAttribute(35)
	Qt__WA_SetPalette                      = Qt__WidgetAttribute(36)
	Qt__WA_SetFont                         = Qt__WidgetAttribute(37)
	Qt__WA_SetCursor                       = Qt__WidgetAttribute(38)
	Qt__WA_NoChildEventsFromChildren       = Qt__WidgetAttribute(39)
	Qt__WA_WindowModified                  = Qt__WidgetAttribute(41)
	Qt__WA_Resized                         = Qt__WidgetAttribute(42)
	Qt__WA_Moved                           = Qt__WidgetAttribute(43)
	Qt__WA_PendingUpdate                   = Qt__WidgetAttribute(44)
	Qt__WA_InvalidSize                     = Qt__WidgetAttribute(45)
	Qt__WA_MacBrushedMetal                 = Qt__WidgetAttribute(46)
	Qt__WA_MacMetalStyle                   = Qt__WidgetAttribute(Qt__WA_MacBrushedMetal)
	Qt__WA_CustomWhatsThis                 = Qt__WidgetAttribute(47)
	Qt__WA_LayoutOnEntireRect              = Qt__WidgetAttribute(48)
	Qt__WA_OutsideWSRange                  = Qt__WidgetAttribute(49)
	Qt__WA_GrabbedShortcut                 = Qt__WidgetAttribute(50)
	Qt__WA_TransparentForMouseEvents       = Qt__WidgetAttribute(51)
	Qt__WA_PaintUnclipped                  = Qt__WidgetAttribute(52)
	Qt__WA_SetWindowIcon                   = Qt__WidgetAttribute(53)
	Qt__WA_NoMouseReplay                   = Qt__WidgetAttribute(54)
	Qt__WA_DeleteOnClose                   = Qt__WidgetAttribute(55)
	Qt__WA_RightToLeft                     = Qt__WidgetAttribute(56)
	Qt__WA_SetLayoutDirection              = Qt__WidgetAttribute(57)
	Qt__WA_NoChildEventsForParent          = Qt__WidgetAttribute(58)
	Qt__WA_ForceUpdatesDisabled            = Qt__WidgetAttribute(59)
	Qt__WA_WState_Created                  = Qt__WidgetAttribute(60)
	Qt__WA_WState_CompressKeys             = Qt__WidgetAttribute(61)
	Qt__WA_WState_InPaintEvent             = Qt__WidgetAttribute(62)
	Qt__WA_WState_Reparented               = Qt__WidgetAttribute(63)
	Qt__WA_WState_ConfigPending            = Qt__WidgetAttribute(64)
	Qt__WA_WState_Polished                 = Qt__WidgetAttribute(66)
	Qt__WA_WState_DND                      = Qt__WidgetAttribute(67)
	Qt__WA_WState_OwnSizePolicy            = Qt__WidgetAttribute(68)
	Qt__WA_WState_ExplicitShowHide         = Qt__WidgetAttribute(69)
	Qt__WA_ShowModal                       = Qt__WidgetAttribute(70)
	Qt__WA_MouseNoMask                     = Qt__WidgetAttribute(71)
	Qt__WA_GroupLeader                     = Qt__WidgetAttribute(72)
	Qt__WA_NoMousePropagation              = Qt__WidgetAttribute(73)
	Qt__WA_Hover                           = Qt__WidgetAttribute(74)
	Qt__WA_InputMethodTransparent          = Qt__WidgetAttribute(75)
	Qt__WA_QuitOnClose                     = Qt__WidgetAttribute(76)
	Qt__WA_KeyboardFocusChange             = Qt__WidgetAttribute(77)
	Qt__WA_AcceptDrops                     = Qt__WidgetAttribute(78)
	Qt__WA_DropSiteRegistered              = Qt__WidgetAttribute(79)
	Qt__WA_ForceAcceptDrops                = Qt__WidgetAttribute(Qt__WA_DropSiteRegistered)
	Qt__WA_WindowPropagation               = Qt__WidgetAttribute(80)
	Qt__WA_NoX11EventCompression           = Qt__WidgetAttribute(81)
	Qt__WA_TintedBackground                = Qt__WidgetAttribute(82)
	Qt__WA_X11OpenGLOverlay                = Qt__WidgetAttribute(83)
	Qt__WA_AlwaysShowToolTips              = Qt__WidgetAttribute(84)
	Qt__WA_MacOpaqueSizeGrip               = Qt__WidgetAttribute(85)
	Qt__WA_SetStyle                        = Qt__WidgetAttribute(86)
	Qt__WA_SetLocale                       = Qt__WidgetAttribute(87)
	Qt__WA_MacShowFocusRect                = Qt__WidgetAttribute(88)
	Qt__WA_MacNormalSize                   = Qt__WidgetAttribute(89)
	Qt__WA_MacSmallSize                    = Qt__WidgetAttribute(90)
	Qt__WA_MacMiniSize                     = Qt__WidgetAttribute(91)
	Qt__WA_LayoutUsesWidgetRect            = Qt__WidgetAttribute(92)
	Qt__WA_StyledBackground                = Qt__WidgetAttribute(93)
	Qt__WA_MSWindowsUseDirect3D            = Qt__WidgetAttribute(94)
	Qt__WA_CanHostQMdiSubWindowTitleBar    = Qt__WidgetAttribute(95)
	Qt__WA_MacAlwaysShowToolWindow         = Qt__WidgetAttribute(96)
	Qt__WA_StyleSheet                      = Qt__WidgetAttribute(97)
	Qt__WA_ShowWithoutActivating           = Qt__WidgetAttribute(98)
	Qt__WA_X11BypassTransientForHint       = Qt__WidgetAttribute(99)
	Qt__WA_NativeWindow                    = Qt__WidgetAttribute(100)
	Qt__WA_DontCreateNativeAncestors       = Qt__WidgetAttribute(101)
	Qt__WA_MacVariableSize                 = Qt__WidgetAttribute(102)
	Qt__WA_DontShowOnScreen                = Qt__WidgetAttribute(103)
	Qt__WA_X11NetWmWindowTypeDesktop       = Qt__WidgetAttribute(104)
	Qt__WA_X11NetWmWindowTypeDock          = Qt__WidgetAttribute(105)
	Qt__WA_X11NetWmWindowTypeToolBar       = Qt__WidgetAttribute(106)
	Qt__WA_X11NetWmWindowTypeMenu          = Qt__WidgetAttribute(107)
	Qt__WA_X11NetWmWindowTypeUtility       = Qt__WidgetAttribute(108)
	Qt__WA_X11NetWmWindowTypeSplash        = Qt__WidgetAttribute(109)
	Qt__WA_X11NetWmWindowTypeDialog        = Qt__WidgetAttribute(110)
	Qt__WA_X11NetWmWindowTypeDropDownMenu  = Qt__WidgetAttribute(111)
	Qt__WA_X11NetWmWindowTypePopupMenu     = Qt__WidgetAttribute(112)
	Qt__WA_X11NetWmWindowTypeToolTip       = Qt__WidgetAttribute(113)
	Qt__WA_X11NetWmWindowTypeNotification  = Qt__WidgetAttribute(114)
	Qt__WA_X11NetWmWindowTypeCombo         = Qt__WidgetAttribute(115)
	Qt__WA_X11NetWmWindowTypeDND           = Qt__WidgetAttribute(116)
	Qt__WA_MacFrameworkScaled              = Qt__WidgetAttribute(117)
	Qt__WA_SetWindowModality               = Qt__WidgetAttribute(118)
	Qt__WA_WState_WindowOpacitySet         = Qt__WidgetAttribute(119)
	Qt__WA_TranslucentBackground           = Qt__WidgetAttribute(120)
	Qt__WA_AcceptTouchEvents               = Qt__WidgetAttribute(121)
	Qt__WA_WState_AcceptedTouchBeginEvent  = Qt__WidgetAttribute(122)
	Qt__WA_TouchPadAcceptSingleTouchEvents = Qt__WidgetAttribute(123)
	Qt__WA_X11DoNotAcceptFocus             = Qt__WidgetAttribute(126)
	Qt__WA_MacNoShadow                     = Qt__WidgetAttribute(127)
	Qt__WA_AlwaysStackOnTop                = Qt__WidgetAttribute(128)
	Qt__WA_AttributeCount                  = Qt__WidgetAttribute(129)
)

//Qt::WindowFrameSection
type Qt__WindowFrameSection int64

const (
	Qt__NoSection          = Qt__WindowFrameSection(0)
	Qt__LeftSection        = Qt__WindowFrameSection(1)
	Qt__TopLeftSection     = Qt__WindowFrameSection(2)
	Qt__TopSection         = Qt__WindowFrameSection(3)
	Qt__TopRightSection    = Qt__WindowFrameSection(4)
	Qt__RightSection       = Qt__WindowFrameSection(5)
	Qt__BottomRightSection = Qt__WindowFrameSection(6)
	Qt__BottomSection      = Qt__WindowFrameSection(7)
	Qt__BottomLeftSection  = Qt__WindowFrameSection(8)
	Qt__TitleBarArea       = Qt__WindowFrameSection(9)
)

//Qt::WindowModality
type Qt__WindowModality int64

const (
	Qt__NonModal         = Qt__WindowModality(0)
	Qt__WindowModal      = Qt__WindowModality(1)
	Qt__ApplicationModal = Qt__WindowModality(2)
)

//Qt::WindowState
type Qt__WindowState int64

const (
	Qt__WindowNoState    = Qt__WindowState(0x00000000)
	Qt__WindowMinimized  = Qt__WindowState(0x00000001)
	Qt__WindowMaximized  = Qt__WindowState(0x00000002)
	Qt__WindowFullScreen = Qt__WindowState(0x00000004)
	Qt__WindowActive     = Qt__WindowState(0x00000008)
)

//Qt::WindowType
type Qt__WindowType int64

const (
	Qt__Widget                              = Qt__WindowType(0x00000000)
	Qt__Window                              = Qt__WindowType(0x00000001)
	Qt__Dialog                              = Qt__WindowType(0x00000002 | Qt__Window)
	Qt__Sheet                               = Qt__WindowType(0x00000004 | Qt__Window)
	Qt__Drawer                              = Qt__WindowType(Qt__Sheet | Qt__Dialog)
	Qt__Popup                               = Qt__WindowType(0x00000008 | Qt__Window)
	Qt__Tool                                = Qt__WindowType(Qt__Popup | Qt__Dialog)
	Qt__ToolTip                             = Qt__WindowType(Qt__Popup | Qt__Sheet)
	Qt__SplashScreen                        = Qt__WindowType(Qt__ToolTip | Qt__Dialog)
	Qt__Desktop                             = Qt__WindowType(0x00000010 | Qt__Window)
	Qt__SubWindow                           = Qt__WindowType(0x00000012)
	Qt__ForeignWindow                       = Qt__WindowType(0x00000020 | Qt__Window)
	Qt__CoverWindow                         = Qt__WindowType(0x00000040 | Qt__Window)
	Qt__WindowType_Mask                     = Qt__WindowType(0x000000ff)
	Qt__MSWindowsFixedSizeDialogHint        = Qt__WindowType(0x00000100)
	Qt__MSWindowsOwnDC                      = Qt__WindowType(0x00000200)
	Qt__BypassWindowManagerHint             = Qt__WindowType(0x00000400)
	Qt__X11BypassWindowManagerHint          = Qt__WindowType(Qt__BypassWindowManagerHint)
	Qt__FramelessWindowHint                 = Qt__WindowType(0x00000800)
	Qt__WindowTitleHint                     = Qt__WindowType(0x00001000)
	Qt__WindowSystemMenuHint                = Qt__WindowType(0x00002000)
	Qt__WindowMinimizeButtonHint            = Qt__WindowType(0x00004000)
	Qt__WindowMaximizeButtonHint            = Qt__WindowType(0x00008000)
	Qt__WindowMinMaxButtonsHint             = Qt__WindowType(Qt__WindowMinimizeButtonHint | Qt__WindowMaximizeButtonHint)
	Qt__WindowContextHelpButtonHint         = Qt__WindowType(0x00010000)
	Qt__WindowShadeButtonHint               = Qt__WindowType(0x00020000)
	Qt__WindowStaysOnTopHint                = Qt__WindowType(0x00040000)
	Qt__WindowTransparentForInput           = Qt__WindowType(0x00080000)
	Qt__WindowOverridesSystemGestures       = Qt__WindowType(0x00100000)
	Qt__WindowDoesNotAcceptFocus            = Qt__WindowType(0x00200000)
	Qt__MaximizeUsingFullscreenGeometryHint = Qt__WindowType(0x00400000)
	Qt__CustomizeWindowHint                 = Qt__WindowType(0x02000000)
	Qt__WindowStaysOnBottomHint             = Qt__WindowType(0x04000000)
	Qt__WindowCloseButtonHint               = Qt__WindowType(0x08000000)
	Qt__MacWindowToolBarButtonHint          = Qt__WindowType(0x10000000)
	Qt__BypassGraphicsProxyWidget           = Qt__WindowType(0x20000000)
	Qt__NoDropShadowWindowHint              = Qt__WindowType(0x40000000)
	Qt__WindowFullscreenButtonHint          = Qt__WindowType(0x80000000)
	Qt__WindowOkButtonHint                  = Qt__WindowType(0x00080000)
	Qt__WindowCancelButtonHint              = Qt__WindowType(0x00100000)
)
