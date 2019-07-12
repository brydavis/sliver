// +build windows
// +build 386 !amd64
package win

import (
	"unsafe"
)

type INPUTCONTEXT struct {
	storage [320]byte
}

func (this *INPUTCONTEXT) HWnd() *HWND { // 4
	return (*HWND)(unsafe.Pointer(&this.storage[0]))
}
func (this *INPUTCONTEXT) FOpen() *BOOL { // 4
	return (*BOOL)(unsafe.Pointer(&this.storage[4]))
}
func (this *INPUTCONTEXT) PtStatusWndPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[8]))
}
func (this *INPUTCONTEXT) PtSoftKbdPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[16]))
}
func (this *INPUTCONTEXT) FdwConversion() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *INPUTCONTEXT) FdwSentence() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *INPUTCONTEXT) LfFont() *LOGFONT { // 92
	return (*LOGFONT)(unsafe.Pointer(&this.storage[32]))
}
func (this *INPUTCONTEXT) CfCompForm() *COMPOSITIONFORM { // 28
	return (*COMPOSITIONFORM)(unsafe.Pointer(&this.storage[124]))
}
func (this *INPUTCONTEXT) CfCandForm() *[4]CANDIDATEFORM { // 128
	return (*[4]CANDIDATEFORM)(unsafe.Pointer(&this.storage[152]))
}
func (this *INPUTCONTEXT) HCompStr() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[280]))
}
func (this *INPUTCONTEXT) HCandInfo() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[284]))
}
func (this *INPUTCONTEXT) HGuideLine() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[288]))
}
func (this *INPUTCONTEXT) HPrivate() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[292]))
}
func (this *INPUTCONTEXT) DwNumMsgBuf() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[296]))
}
func (this *INPUTCONTEXT) HMsgBuf() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[300]))
}
func (this *INPUTCONTEXT) FdwInit() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[304]))
}
func (this *INPUTCONTEXT) DwReserve() *[3]DWORD { // 12
	return (*[3]DWORD)(unsafe.Pointer(&this.storage[308]))
}

type IP_ADAPTER_ADDRESSES_LH struct {
	union1                 ULONGLONG
	Next                   *IP_ADAPTER_ADDRESSES_LH
	AdapterName            PCHAR
	FirstUnicastAddress    PIP_ADAPTER_UNICAST_ADDRESS_LH
	FirstAnycastAddress    PIP_ADAPTER_ANYCAST_ADDRESS_XP
	FirstMulticastAddress  PIP_ADAPTER_MULTICAST_ADDRESS_XP
	FirstDnsServerAddress  PIP_ADAPTER_DNS_SERVER_ADDRESS_XP
	DnsSuffix              PWCHAR
	Description            PWCHAR
	FriendlyName           PWCHAR
	PhysicalAddress        [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	PhysicalAddressLength  ULONG
	union2                 ULONG
	Mtu                    ULONG
	IfType                 IFTYPE
	OperStatus             IF_OPER_STATUS
	Ipv6IfIndex            IF_INDEX
	ZoneIndices            [16]ULONG
	FirstPrefix            PIP_ADAPTER_PREFIX_XP
	TransmitLinkSpeed      ULONG64
	ReceiveLinkSpeed       ULONG64
	FirstWinsServerAddress PIP_ADAPTER_WINS_SERVER_ADDRESS_LH
	FirstGatewayAddress    PIP_ADAPTER_GATEWAY_ADDRESS_LH
	Ipv4Metric             ULONG
	Ipv6Metric             ULONG
	Luid                   IF_LUID
	Dhcpv4Server           SOCKET_ADDRESS
	CompartmentId          NET_IF_COMPARTMENT_ID
	NetworkGuid            NET_IF_NETWORK_GUID
	ConnectionType         NET_IF_CONNECTION_TYPE
	TunnelType             TUNNEL_TYPE
	Dhcpv6Server           SOCKET_ADDRESS
	Dhcpv6ClientDuid       [MAX_DHCPV6_DUID_LENGTH]BYTE
	Dhcpv6ClientDuidLength ULONG
	Dhcpv6Iaid             ULONG
	FirstDnsSuffix         PIP_ADAPTER_DNS_SUFFIX
	padding1               [4]byte
}
type IP_ADAPTER_ANYCAST_ADDRESS_XP struct {
	union1   ULONGLONG
	Next     *IP_ADAPTER_ANYCAST_ADDRESS_XP
	Address  SOCKET_ADDRESS
	padding1 [4]byte
}
type IP_ADAPTER_DNS_SERVER_ADDRESS_XP struct {
	union1   ULONGLONG
	Next     *IP_ADAPTER_DNS_SERVER_ADDRESS_XP
	Address  SOCKET_ADDRESS
	padding1 [4]byte
}
type IP_ADAPTER_GATEWAY_ADDRESS_LH struct {
	union1   ULONGLONG
	Next     *IP_ADAPTER_GATEWAY_ADDRESS_LH
	Address  SOCKET_ADDRESS
	padding1 [4]byte
}
type IP_ADAPTER_MULTICAST_ADDRESS_XP struct {
	union1   ULONGLONG
	Next     *IP_ADAPTER_MULTICAST_ADDRESS_XP
	Address  SOCKET_ADDRESS
	padding1 [4]byte
}
type IP_ADAPTER_WINS_SERVER_ADDRESS_LH struct {
	union1   ULONGLONG
	Next     *IP_ADAPTER_WINS_SERVER_ADDRESS_LH
	Address  SOCKET_ADDRESS
	padding1 [4]byte
}
type MIDIHDR struct {
	storage [64]byte
}

func (this *MIDIHDR) LpData() *LPSTR { // 4
	return (*LPSTR)(unsafe.Pointer(&this.storage[0]))
}
func (this *MIDIHDR) DwBufferLength() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[4]))
}
func (this *MIDIHDR) DwBytesRecorded() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[8]))
}
func (this *MIDIHDR) DwUser() *DWORD_PTR { // 4
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[12]))
}
func (this *MIDIHDR) DwFlags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[16]))
}
func (this *MIDIHDR) LpNext() **MIDIHDR { // 4
	return (**MIDIHDR)(unsafe.Pointer(&this.storage[20]))
}
func (this *MIDIHDR) Reserved() *DWORD_PTR { // 4
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[24]))
}
func (this *MIDIHDR) DwOffset() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *MIDIHDR) DwReserved() *[8]DWORD_PTR { // 32
	return (*[8]DWORD_PTR)(unsafe.Pointer(&this.storage[32]))
}

type MIXERLINECONTROLS struct {
	CbStruct  DWORD
	DwLineID  DWORD
	union1    DWORD
	CControls DWORD
	Cbmxctrl  DWORD
	storage1  [4]byte
}
type PRINTDLG struct {
	storage [66]byte
}

func (this *PRINTDLG) LStructSize() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PRINTDLG) HwndOwner() *HWND { // 4
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *PRINTDLG) HDevMode() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[8]))
}
func (this *PRINTDLG) HDevNames() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[12]))
}
func (this *PRINTDLG) HDC() *HDC { // 4
	return (*HDC)(unsafe.Pointer(&this.storage[16]))
}
func (this *PRINTDLG) Flags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[20]))
}
func (this *PRINTDLG) NFromPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *PRINTDLG) NToPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[26]))
}
func (this *PRINTDLG) NMinPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *PRINTDLG) NMaxPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[30]))
}
func (this *PRINTDLG) NCopies() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[32]))
}
func (this *PRINTDLG) HInstance() *HINSTANCE { // 4
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[34]))
}
func (this *PRINTDLG) LCustData() *LPARAM { // 4
	return (*LPARAM)(unsafe.Pointer(&this.storage[38]))
}
func (this *PRINTDLG) LpfnPrintHook() *uintptr { // 4
	return (*uintptr)(unsafe.Pointer(&this.storage[42]))
}
func (this *PRINTDLG) LpfnSetupHook() *uintptr { // 4
	return (*uintptr)(unsafe.Pointer(&this.storage[46]))
}
func (this *PRINTDLG) LpPrintTemplateName() *LPCWSTR { // 4
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[50]))
}
func (this *PRINTDLG) LpSetupTemplateName() *LPCWSTR { // 4
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[54]))
}
func (this *PRINTDLG) HPrintTemplate() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[58]))
}
func (this *PRINTDLG) HSetupTemplate() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[62]))
}

type SIZE_T uint32
type STRRET struct {
	UType UINT
	cStr  [260]byte
}
type TASKDIALOGCONFIG struct {
	storage [96]byte
}

func (this *TASKDIALOGCONFIG) CbSize() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[0]))
}
func (this *TASKDIALOGCONFIG) HwndParent() *HWND {
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *TASKDIALOGCONFIG) HInstance() *HINSTANCE {
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[8]))
}
func (this *TASKDIALOGCONFIG) DwFlags() *TASKDIALOG_FLAGS {
	return (*TASKDIALOG_FLAGS)(unsafe.Pointer(&this.storage[12]))
}
func (this *TASKDIALOGCONFIG) DwCommonButtons() *TASKDIALOG_COMMON_BUTTON_FLAGS {
	return (*TASKDIALOG_COMMON_BUTTON_FLAGS)(unsafe.Pointer(&this.storage[16]))
}
func (this *TASKDIALOGCONFIG) PszWindowTitle() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[20]))
}
func (this *TASKDIALOGCONFIG) HMainIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[24]))
}
func (this *TASKDIALOGCONFIG) PszMainIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[24]))
}
func (this *TASKDIALOGCONFIG) PszMainInstruction() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[28]))
}
func (this *TASKDIALOGCONFIG) PszContent() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[32]))
}
func (this *TASKDIALOGCONFIG) CButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[36]))
}
func (this *TASKDIALOGCONFIG) PButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[40]))
}
func (this *TASKDIALOGCONFIG) NDefaultButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[44]))
}
func (this *TASKDIALOGCONFIG) CRadioButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[48]))
}
func (this *TASKDIALOGCONFIG) PRadioButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[52]))
}
func (this *TASKDIALOGCONFIG) NDefaultRadioButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[56]))
}
func (this *TASKDIALOGCONFIG) PszVerificationText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[60]))
}
func (this *TASKDIALOGCONFIG) PszExpandedInformation() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[64]))
}
func (this *TASKDIALOGCONFIG) PszExpandedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[68]))
}
func (this *TASKDIALOGCONFIG) PszCollapsedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[72]))
}
func (this *TASKDIALOGCONFIG) HFooterIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[76]))
}
func (this *TASKDIALOGCONFIG) PszFooterIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[76]))
}
func (this *TASKDIALOGCONFIG) PszFooter() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[80]))
}
func (this *TASKDIALOGCONFIG) PfCallback() *uintptr {
	return (*uintptr)(unsafe.Pointer(&this.storage[84]))
}
func (this *TASKDIALOGCONFIG) LpCallbackData() *LONG_PTR {
	return (*LONG_PTR)(unsafe.Pointer(&this.storage[88]))
}
func (this *TASKDIALOGCONFIG) CxWidth() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[92]))
}

type VARIANT struct {
	union1 [16]byte
}

func (this *VARIANT) PRecInfo() *IRecordInfo {
	return (*IRecordInfo)(unsafe.Pointer(&this.union1[12]))
}

type WSADATA struct {
	WVersion       uint16
	WHighVersion   uint16
	SzDescription  [WSADESCRIPTION_LEN + 1]byte
	SzSystemStatus [WSASYS_STATUS_LEN + 1]byte
	IMaxSockets    uint16
	IMaxUdpDg      uint16
	LpVendorInfo   *byte
}