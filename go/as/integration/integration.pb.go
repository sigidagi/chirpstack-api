// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/integration/integration.proto

package integration

import (
	fmt "fmt"
	common "github.com/brocaar/chirpstack-api/go/v3/common"
	gw "github.com/brocaar/chirpstack-api/go/v3/gw"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorType int32

const (
	// Unknown type.
	ErrorType_UNKNOWN ErrorType = 0
	// Error related to the downlink payload size.
	// Usually seen when the payload exceeded the maximum allowed payload size.
	ErrorType_DOWNLINK_PAYLOAD_SIZE ErrorType = 1
	// Error related to the downlink frame-counter.
	// Usually seen when the frame-counter has already been used.
	ErrorType_DOWNLINK_FCNT ErrorType = 2
	// Uplink codec error.
	ErrorType_UPLINK_CODEC ErrorType = 3
	// Downlink codec error.
	ErrorType_DOWNLINK_CODEC ErrorType = 4
	// OTAA error.
	ErrorType_OTAA ErrorType = 5
	// Uplink frame-counter was reset.
	ErrorType_UPLINK_FCNT_RESET ErrorType = 6
	// Uplink MIC error.
	ErrorType_UPLINK_MIC ErrorType = 7
	// Uplink frame-counter retransmission.
	ErrorType_UPLINK_FCNT_RETRANSMISSION ErrorType = 8
	// Downlink gateway error.
	ErrorType_DOWNLINK_GATEWAY ErrorType = 9
)

var ErrorType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DOWNLINK_PAYLOAD_SIZE",
	2: "DOWNLINK_FCNT",
	3: "UPLINK_CODEC",
	4: "DOWNLINK_CODEC",
	5: "OTAA",
	6: "UPLINK_FCNT_RESET",
	7: "UPLINK_MIC",
	8: "UPLINK_FCNT_RETRANSMISSION",
	9: "DOWNLINK_GATEWAY",
}

var ErrorType_value = map[string]int32{
	"UNKNOWN":                    0,
	"DOWNLINK_PAYLOAD_SIZE":      1,
	"DOWNLINK_FCNT":              2,
	"UPLINK_CODEC":               3,
	"DOWNLINK_CODEC":             4,
	"OTAA":                       5,
	"UPLINK_FCNT_RESET":          6,
	"UPLINK_MIC":                 7,
	"UPLINK_FCNT_RETRANSMISSION": 8,
	"DOWNLINK_GATEWAY":           9,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}

func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{0}
}

// UplinkEvent is the message sent when an uplink payload has been received.
type UplinkEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Receiving gateway RX info.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,5,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// TX info.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,6,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// Device has ADR bit set.
	Adr bool `protobuf:"varint,7,opt,name=adr,proto3" json:"adr,omitempty"`
	// Data-rate.
	Dr uint32 `protobuf:"varint,8,opt,name=dr,proto3" json:"dr,omitempty"`
	// Frame counter.
	FCnt uint32 `protobuf:"varint,9,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Frame port.
	FPort uint32 `protobuf:"varint,10,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// FRMPayload data.
	Data []byte `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
	// JSON string containing the decoded object.
	// Note that this is only set when a codec is configured in the Device Profile.
	ObjectJson string `protobuf:"bytes,12,opt,name=object_json,json=objectJSON,proto3" json:"object_json,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UplinkEvent) Reset()         { *m = UplinkEvent{} }
func (m *UplinkEvent) String() string { return proto.CompactTextString(m) }
func (*UplinkEvent) ProtoMessage()    {}
func (*UplinkEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{0}
}

func (m *UplinkEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkEvent.Unmarshal(m, b)
}
func (m *UplinkEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkEvent.Marshal(b, m, deterministic)
}
func (m *UplinkEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkEvent.Merge(m, src)
}
func (m *UplinkEvent) XXX_Size() int {
	return xxx_messageInfo_UplinkEvent.Size(m)
}
func (m *UplinkEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkEvent proto.InternalMessageInfo

func (m *UplinkEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *UplinkEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *UplinkEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *UplinkEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *UplinkEvent) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *UplinkEvent) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkEvent) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *UplinkEvent) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *UplinkEvent) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *UplinkEvent) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *UplinkEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UplinkEvent) GetObjectJson() string {
	if m != nil {
		return m.ObjectJson
	}
	return ""
}

func (m *UplinkEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// JoinEvent is the message sent when a device joined the network.
// Note that this is only sent after the first received uplink after the
// device (re)activation.
type JoinEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Device address.
	DevAddr []byte `protobuf:"bytes,5,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
	// Receiving gateway RX info.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,6,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// TX info.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,7,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// Data-rate.
	Dr uint32 `protobuf:"varint,8,opt,name=dr,proto3" json:"dr,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JoinEvent) Reset()         { *m = JoinEvent{} }
func (m *JoinEvent) String() string { return proto.CompactTextString(m) }
func (*JoinEvent) ProtoMessage()    {}
func (*JoinEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{1}
}

func (m *JoinEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinEvent.Unmarshal(m, b)
}
func (m *JoinEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinEvent.Marshal(b, m, deterministic)
}
func (m *JoinEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinEvent.Merge(m, src)
}
func (m *JoinEvent) XXX_Size() int {
	return xxx_messageInfo_JoinEvent.Size(m)
}
func (m *JoinEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinEvent.DiscardUnknown(m)
}

var xxx_messageInfo_JoinEvent proto.InternalMessageInfo

func (m *JoinEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *JoinEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *JoinEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *JoinEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *JoinEvent) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *JoinEvent) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *JoinEvent) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *JoinEvent) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *JoinEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// AckEvent is the message sent when a confirmation on a confirmed downlink
// has been received -or- when the downlink timed out.
type AckEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Frame was acknowledged.
	Acknowledged bool `protobuf:"varint,5,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	// Downlink frame counter to which the acknowledgement relates.
	FCnt uint32 `protobuf:"varint,6,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AckEvent) Reset()         { *m = AckEvent{} }
func (m *AckEvent) String() string { return proto.CompactTextString(m) }
func (*AckEvent) ProtoMessage()    {}
func (*AckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{2}
}

func (m *AckEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckEvent.Unmarshal(m, b)
}
func (m *AckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckEvent.Marshal(b, m, deterministic)
}
func (m *AckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckEvent.Merge(m, src)
}
func (m *AckEvent) XXX_Size() int {
	return xxx_messageInfo_AckEvent.Size(m)
}
func (m *AckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AckEvent proto.InternalMessageInfo

func (m *AckEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *AckEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *AckEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *AckEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *AckEvent) GetAcknowledged() bool {
	if m != nil {
		return m.Acknowledged
	}
	return false
}

func (m *AckEvent) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *AckEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// TxAckEvent is the message sent when a downlink was acknowledged by the gateway
// for transmission. As a downlink can be scheduled in the future, this event
// does not confirm that the message has already been transmitted.
type TxAckEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Downlink frame-counter.
	FCnt uint32 `protobuf:"varint,5,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TxAckEvent) Reset()         { *m = TxAckEvent{} }
func (m *TxAckEvent) String() string { return proto.CompactTextString(m) }
func (*TxAckEvent) ProtoMessage()    {}
func (*TxAckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{3}
}

func (m *TxAckEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxAckEvent.Unmarshal(m, b)
}
func (m *TxAckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxAckEvent.Marshal(b, m, deterministic)
}
func (m *TxAckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxAckEvent.Merge(m, src)
}
func (m *TxAckEvent) XXX_Size() int {
	return xxx_messageInfo_TxAckEvent.Size(m)
}
func (m *TxAckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TxAckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TxAckEvent proto.InternalMessageInfo

func (m *TxAckEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *TxAckEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *TxAckEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *TxAckEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *TxAckEvent) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *TxAckEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// ErrorEvent is the message sent when an error occurred.
type ErrorEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Error type.
	Type ErrorType `protobuf:"varint,5,opt,name=type,proto3,enum=integration.ErrorType" json:"type,omitempty"`
	// Error message.
	Error string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// Downlink frame-counter (in case the downlink is related to a scheduled downlink).
	FCnt uint32 `protobuf:"varint,7,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ErrorEvent) Reset()         { *m = ErrorEvent{} }
func (m *ErrorEvent) String() string { return proto.CompactTextString(m) }
func (*ErrorEvent) ProtoMessage()    {}
func (*ErrorEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{4}
}

func (m *ErrorEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorEvent.Unmarshal(m, b)
}
func (m *ErrorEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorEvent.Marshal(b, m, deterministic)
}
func (m *ErrorEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorEvent.Merge(m, src)
}
func (m *ErrorEvent) XXX_Size() int {
	return xxx_messageInfo_ErrorEvent.Size(m)
}
func (m *ErrorEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorEvent proto.InternalMessageInfo

func (m *ErrorEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ErrorEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ErrorEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *ErrorEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *ErrorEvent) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_UNKNOWN
}

func (m *ErrorEvent) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorEvent) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *ErrorEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// StatusEvent is the message sent when a device-status mac-command was sent
// by the device.
type StatusEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// The demodulation signal-to-noise ratio in dB for the last successfully
	// received device-status request by the Network Server.
	Margin uint32 `protobuf:"varint,5,opt,name=margin,proto3" json:"margin,omitempty"`
	// Device is connected to an external power source.
	ExternalPowerSource bool `protobuf:"varint,6,opt,name=external_power_source,json=externalPowerSource,proto3" json:"external_power_source,omitempty"`
	// Battery level is not available.
	BatteryLevelUnavailable bool `protobuf:"varint,7,opt,name=battery_level_unavailable,json=batteryLevelUnavailable,proto3" json:"battery_level_unavailable,omitempty"`
	// Battery level.
	BatteryLevel float32 `protobuf:"fixed32,8,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"`
	// User-defined device tags.
	Tags                 map[string]string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusEvent) Reset()         { *m = StatusEvent{} }
func (m *StatusEvent) String() string { return proto.CompactTextString(m) }
func (*StatusEvent) ProtoMessage()    {}
func (*StatusEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{5}
}

func (m *StatusEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusEvent.Unmarshal(m, b)
}
func (m *StatusEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusEvent.Marshal(b, m, deterministic)
}
func (m *StatusEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusEvent.Merge(m, src)
}
func (m *StatusEvent) XXX_Size() int {
	return xxx_messageInfo_StatusEvent.Size(m)
}
func (m *StatusEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StatusEvent proto.InternalMessageInfo

func (m *StatusEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *StatusEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *StatusEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *StatusEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *StatusEvent) GetMargin() uint32 {
	if m != nil {
		return m.Margin
	}
	return 0
}

func (m *StatusEvent) GetExternalPowerSource() bool {
	if m != nil {
		return m.ExternalPowerSource
	}
	return false
}

func (m *StatusEvent) GetBatteryLevelUnavailable() bool {
	if m != nil {
		return m.BatteryLevelUnavailable
	}
	return false
}

func (m *StatusEvent) GetBatteryLevel() float32 {
	if m != nil {
		return m.BatteryLevel
	}
	return 0
}

func (m *StatusEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// LocationEvent is the message sent when a geolocation resolve was returned.
type LocationEvent struct {
	// Application ID.
	ApplicationId uint64 `protobuf:"varint,1,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,4,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Location.
	Location *common.Location `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	// User-defined device tags.
	Tags map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Uplink IDs used for geolocation.
	UplinkIds            [][]byte `protobuf:"bytes,7,rep,name=uplink_ids,json=uplinkIDs,proto3" json:"uplink_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationEvent) Reset()         { *m = LocationEvent{} }
func (m *LocationEvent) String() string { return proto.CompactTextString(m) }
func (*LocationEvent) ProtoMessage()    {}
func (*LocationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ba82681be587591, []int{6}
}

func (m *LocationEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationEvent.Unmarshal(m, b)
}
func (m *LocationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationEvent.Marshal(b, m, deterministic)
}
func (m *LocationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationEvent.Merge(m, src)
}
func (m *LocationEvent) XXX_Size() int {
	return xxx_messageInfo_LocationEvent.Size(m)
}
func (m *LocationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LocationEvent proto.InternalMessageInfo

func (m *LocationEvent) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *LocationEvent) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *LocationEvent) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *LocationEvent) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *LocationEvent) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *LocationEvent) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LocationEvent) GetUplinkIds() [][]byte {
	if m != nil {
		return m.UplinkIds
	}
	return nil
}

func init() {
	proto.RegisterEnum("integration.ErrorType", ErrorType_name, ErrorType_value)
	proto.RegisterType((*UplinkEvent)(nil), "integration.UplinkEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.UplinkEvent.TagsEntry")
	proto.RegisterType((*JoinEvent)(nil), "integration.JoinEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.JoinEvent.TagsEntry")
	proto.RegisterType((*AckEvent)(nil), "integration.AckEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.AckEvent.TagsEntry")
	proto.RegisterType((*TxAckEvent)(nil), "integration.TxAckEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.TxAckEvent.TagsEntry")
	proto.RegisterType((*ErrorEvent)(nil), "integration.ErrorEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.ErrorEvent.TagsEntry")
	proto.RegisterType((*StatusEvent)(nil), "integration.StatusEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.StatusEvent.TagsEntry")
	proto.RegisterType((*LocationEvent)(nil), "integration.LocationEvent")
	proto.RegisterMapType((map[string]string)(nil), "integration.LocationEvent.TagsEntry")
}

func init() {
	proto.RegisterFile("as/integration/integration.proto", fileDescriptor_5ba82681be587591)
}

var fileDescriptor_5ba82681be587591 = []byte{
	// 925 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x4f, 0x8f, 0xda, 0x46,
	0x18, 0xc6, 0x6b, 0xf3, 0xcf, 0x7e, 0x81, 0xad, 0x33, 0x9b, 0x4d, 0xbc, 0x2b, 0xb5, 0xa1, 0xdb,
	0x56, 0x22, 0x51, 0x0b, 0xd2, 0x6e, 0x9b, 0xae, 0x72, 0xa3, 0x40, 0x2b, 0x92, 0x8d, 0x59, 0x19,
	0xd0, 0x36, 0xb9, 0x58, 0x83, 0x3d, 0x38, 0xce, 0x1a, 0x8f, 0x35, 0x0c, 0xb0, 0x7c, 0x82, 0x7e,
	0x8a, 0x7e, 0x8e, 0x9e, 0xfa, 0x29, 0x7a, 0xec, 0xa1, 0x87, 0x7e, 0x91, 0xca, 0x63, 0xaf, 0xb1,
	0xa3, 0x95, 0x5a, 0x45, 0x1c, 0x38, 0x31, 0xf3, 0xbc, 0x0f, 0xf6, 0xbc, 0xbf, 0xc7, 0xe3, 0x31,
	0x34, 0xf0, 0xa2, 0xed, 0x05, 0x9c, 0xb8, 0x0c, 0x73, 0x8f, 0x06, 0xd9, 0x71, 0x2b, 0x64, 0x94,
	0x53, 0x54, 0xcd, 0x48, 0x27, 0x87, 0x36, 0x9d, 0xcf, 0x69, 0xd0, 0x8e, 0x7f, 0x62, 0xc7, 0x49,
	0xd5, 0x5d, 0xb7, 0xdd, 0x75, 0x3c, 0x39, 0xfd, 0xa7, 0x00, 0xd5, 0x49, 0xe8, 0x7b, 0xc1, 0x4d,
	0x7f, 0x45, 0x02, 0x8e, 0xbe, 0x86, 0x03, 0x1c, 0x86, 0xbe, 0x67, 0x8b, 0x0b, 0x58, 0x9e, 0xa3,
	0x4b, 0x0d, 0xa9, 0x59, 0x34, 0xeb, 0x19, 0x75, 0xd0, 0x43, 0x4f, 0x41, 0xcb, 0xda, 0x02, 0x3c,
	0x27, 0xba, 0xdc, 0x90, 0x9a, 0xaa, 0xf9, 0x69, 0x46, 0x37, 0xf0, 0x9c, 0xa0, 0x27, 0x50, 0x75,
	0xc8, 0xca, 0xb3, 0x49, 0xec, 0x2a, 0x08, 0x17, 0xc4, 0x92, 0x30, 0x3c, 0x86, 0x8a, 0x43, 0x56,
	0x16, 0x59, 0x7a, 0x7a, 0xb1, 0x21, 0x35, 0x6b, 0x66, 0xd9, 0x21, 0xab, 0xfe, 0x64, 0x80, 0x9e,
	0x42, 0x85, 0xdd, 0x5a, 0x5e, 0x30, 0xa3, 0x7a, 0xa9, 0x51, 0x68, 0x56, 0xcf, 0xb4, 0x96, 0xbb,
	0x6e, 0xc5, 0xab, 0x35, 0x7f, 0x19, 0x04, 0x33, 0x6a, 0x96, 0xd9, 0x6d, 0xf4, 0x1b, 0x59, 0x79,
	0x62, 0x2d, 0x37, 0xa4, 0xbc, 0x75, 0x9c, 0x58, 0x79, 0x6c, 0xd5, 0xa0, 0x80, 0x1d, 0xa6, 0x57,
	0x1a, 0x52, 0x53, 0x31, 0xa3, 0x21, 0x3a, 0x00, 0xd9, 0x61, 0xba, 0xd2, 0x90, 0x9a, 0x75, 0x53,
	0x76, 0x18, 0x3a, 0x84, 0xd2, 0xcc, 0xb2, 0x03, 0xae, 0xab, 0x42, 0x2a, 0xce, 0xba, 0x01, 0x47,
	0x47, 0x50, 0x9e, 0x59, 0x21, 0x65, 0x5c, 0x07, 0xa1, 0x96, 0x66, 0x57, 0x94, 0x71, 0x84, 0xa0,
	0xe8, 0x60, 0x8e, 0xf5, 0xaa, 0x58, 0xb9, 0x18, 0x47, 0x1d, 0xd3, 0xe9, 0x7b, 0x62, 0x73, 0xeb,
	0xfd, 0x82, 0x06, 0x7a, 0x2d, 0xee, 0x38, 0x96, 0x5e, 0x8e, 0x86, 0x06, 0x7a, 0x0e, 0x45, 0x8e,
	0xdd, 0x85, 0x5e, 0x17, 0x5d, 0x9d, 0xb6, 0xb2, 0x29, 0x66, 0xc2, 0x68, 0x8d, 0xb1, 0xbb, 0xe8,
	0x07, 0x9c, 0x6d, 0x4c, 0xe1, 0x3f, 0xf9, 0x01, 0xd4, 0x54, 0x8a, 0xfa, 0xb8, 0x21, 0x1b, 0x11,
	0x8f, 0x6a, 0x46, 0x43, 0xf4, 0x10, 0x4a, 0x2b, 0xec, 0x2f, 0xef, 0x92, 0x88, 0x27, 0x2f, 0xe4,
	0x0b, 0xe9, 0xf4, 0xd7, 0x02, 0xa8, 0x2f, 0xa9, 0x17, 0xec, 0x5f, 0xc6, 0xc7, 0xa0, 0x44, 0x05,
	0xec, 0x38, 0x4c, 0x2f, 0x89, 0x4a, 0x64, 0xec, 0x38, 0x0e, 0xcb, 0xc6, 0x5f, 0xfe, 0xff, 0xf1,
	0x57, 0xfe, 0x23, 0xfe, 0x0f, 0xc3, 0xfe, 0x2e, 0xc9, 0x42, 0x15, 0xb7, 0x68, 0xe4, 0xb2, 0x48,
	0x91, 0xed, 0x2e, 0x89, 0x3f, 0x64, 0x50, 0x3a, 0xf6, 0x1e, 0x6e, 0xb6, 0x53, 0xa8, 0x61, 0xfb,
	0x26, 0xa0, 0x6b, 0x9f, 0x38, 0x2e, 0x71, 0x44, 0x18, 0x8a, 0x99, 0xd3, 0xb6, 0x1b, 0xa3, 0x9c,
	0xd9, 0x18, 0xe7, 0x09, 0xc0, 0x8a, 0x00, 0xf8, 0x24, 0x07, 0xf0, 0xae, 0xd3, 0xdd, 0xf1, 0xfb,
	0x4d, 0x06, 0x18, 0xdf, 0xee, 0x25, 0xc1, 0x94, 0x4e, 0x29, 0x43, 0xe7, 0xfb, 0x84, 0x4e, 0xfc,
	0x04, 0x7f, 0x91, 0xa3, 0xb3, 0xed, 0x63, 0x77, 0x7c, 0xfe, 0x96, 0x01, 0xfa, 0x8c, 0x51, 0xb6,
	0x7f, 0x7c, 0x9e, 0x41, 0x91, 0x6f, 0x42, 0x22, 0xf0, 0x1c, 0x9c, 0x3d, 0xca, 0xa1, 0x10, 0x4b,
	0x1e, 0x6f, 0x42, 0x62, 0x0a, 0x4f, 0xd4, 0x20, 0x89, 0x24, 0xf1, 0xa4, 0xa9, 0x66, 0x3c, 0xd9,
	0x12, 0xae, 0xdc, 0x43, 0x58, 0xb9, 0x87, 0xf0, 0x96, 0xc4, 0xee, 0x08, 0xff, 0x5e, 0x80, 0xea,
	0x88, 0x63, 0xbe, 0x5c, 0xec, 0x1f, 0xe2, 0x47, 0x50, 0x9e, 0x63, 0xe6, 0x7a, 0x41, 0xf2, 0x0c,
	0x26, 0x33, 0x74, 0x06, 0x47, 0xe4, 0x96, 0x13, 0x16, 0x60, 0xdf, 0x0a, 0xe9, 0x9a, 0x30, 0x6b,
	0x41, 0x97, 0xcc, 0x26, 0x02, 0xaf, 0x62, 0x1e, 0xde, 0x15, 0xaf, 0xa2, 0xda, 0x48, 0x94, 0xd0,
	0x0b, 0x38, 0x9e, 0x62, 0xce, 0x09, 0xdb, 0x58, 0x3e, 0x59, 0x11, 0xdf, 0x5a, 0x06, 0x78, 0x85,
	0x3d, 0x1f, 0x4f, 0x7d, 0x92, 0x9c, 0x9e, 0x8f, 0x13, 0xc3, 0x65, 0x54, 0x9f, 0x6c, 0xcb, 0xe8,
	0x4b, 0xa8, 0xe7, 0xfe, 0x2b, 0xde, 0xb7, 0xb2, 0x59, 0xcb, 0xfa, 0xd3, 0x53, 0x50, 0xbd, 0xe7,
	0x14, 0xcc, 0x00, 0xde, 0x5d, 0x72, 0x7f, 0xc9, 0x50, 0xbf, 0xa4, 0x31, 0xe8, 0xfd, 0xcb, 0xee,
	0x1b, 0x50, 0xfc, 0x64, 0x71, 0x22, 0xbd, 0xe8, 0x10, 0x4b, 0xbe, 0xdb, 0xee, 0x16, 0x6d, 0xa6,
	0x0e, 0x74, 0x91, 0x7b, 0xaf, 0x7c, 0x95, 0x83, 0x97, 0xeb, 0xf1, 0x43, 0x7c, 0xe8, 0x33, 0x80,
	0xa5, 0x38, 0x18, 0x2d, 0xcf, 0x89, 0xdf, 0xda, 0x35, 0x53, 0x8d, 0x95, 0x41, 0xef, 0xe3, 0xe9,
	0x3e, 0xfb, 0x53, 0x02, 0x35, 0xdd, 0xc6, 0xa8, 0x0a, 0x95, 0x89, 0xf1, 0xca, 0x18, 0x5e, 0x1b,
	0xda, 0x27, 0xe8, 0x18, 0x8e, 0x7a, 0xc3, 0x6b, 0xe3, 0x72, 0x60, 0xbc, 0xb2, 0xae, 0x3a, 0x6f,
	0x2e, 0x87, 0x9d, 0x9e, 0x35, 0x1a, 0xbc, 0xed, 0x6b, 0x12, 0x7a, 0x00, 0xf5, 0xb4, 0xf4, 0x53,
	0xd7, 0x18, 0x6b, 0x32, 0xd2, 0xa0, 0x36, 0xb9, 0x12, 0x42, 0x77, 0xd8, 0xeb, 0x77, 0xb5, 0x02,
	0x42, 0x70, 0x90, 0x9a, 0x62, 0xad, 0x88, 0x14, 0x28, 0x0e, 0xc7, 0x9d, 0x8e, 0x56, 0x42, 0x47,
	0xf0, 0x20, 0xf1, 0x47, 0x17, 0xb0, 0xcc, 0xfe, 0xa8, 0x3f, 0xd6, 0xca, 0xe8, 0x00, 0x20, 0x91,
	0x5f, 0x0f, 0xba, 0x5a, 0x05, 0x7d, 0x0e, 0x27, 0x79, 0xdb, 0xd8, 0xec, 0x18, 0xa3, 0xd7, 0x83,
	0xd1, 0x68, 0x30, 0x34, 0x34, 0x05, 0x3d, 0x04, 0x2d, 0xbd, 0xc9, 0xcf, 0x9d, 0x71, 0xff, 0xba,
	0xf3, 0x46, 0x53, 0x7f, 0xbc, 0x78, 0xfb, 0xdc, 0xf5, 0xf8, 0xbb, 0xe5, 0x34, 0xca, 0xa2, 0x3d,
	0x65, 0xd4, 0xc6, 0x98, 0xb5, 0xed, 0x77, 0x1e, 0x0b, 0x17, 0x1c, 0xdb, 0x37, 0xdf, 0xe2, 0xd0,
	0x6b, 0xbb, 0xb4, 0xbd, 0x3a, 0x6f, 0xe7, 0x3f, 0xcd, 0xa7, 0x65, 0xf1, 0x81, 0x7d, 0xfe, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x6a, 0x5d, 0x53, 0xb3, 0x0b, 0x00, 0x00,
}
