// Code generated by protoc-gen-go.
// source: coprocess_session_state.proto
// DO NOT EDIT!

package coprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AccessSpec struct {
	Url     string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Methods []string `protobuf:"bytes,2,rep,name=methods" json:"methods,omitempty"`
}

func (m *AccessSpec) Reset()                    { *m = AccessSpec{} }
func (m *AccessSpec) String() string            { return proto.CompactTextString(m) }
func (*AccessSpec) ProtoMessage()               {}
func (*AccessSpec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *AccessSpec) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AccessSpec) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

type AccessDefinition struct {
	ApiName     string        `protobuf:"bytes,1,opt,name=api_name,json=apiName" json:"api_name,omitempty"`
	ApiId       string        `protobuf:"bytes,2,opt,name=api_id,json=apiId" json:"api_id,omitempty"`
	Versions    []string      `protobuf:"bytes,3,rep,name=versions" json:"versions,omitempty"`
	AllowedUrls []*AccessSpec `protobuf:"bytes,4,rep,name=allowed_urls,json=allowedUrls" json:"allowed_urls,omitempty"`
}

func (m *AccessDefinition) Reset()                    { *m = AccessDefinition{} }
func (m *AccessDefinition) String() string            { return proto.CompactTextString(m) }
func (*AccessDefinition) ProtoMessage()               {}
func (*AccessDefinition) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *AccessDefinition) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *AccessDefinition) GetApiId() string {
	if m != nil {
		return m.ApiId
	}
	return ""
}

func (m *AccessDefinition) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *AccessDefinition) GetAllowedUrls() []*AccessSpec {
	if m != nil {
		return m.AllowedUrls
	}
	return nil
}

type BasicAuthData struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *BasicAuthData) Reset()                    { *m = BasicAuthData{} }
func (m *BasicAuthData) String() string            { return proto.CompactTextString(m) }
func (*BasicAuthData) ProtoMessage()               {}
func (*BasicAuthData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *BasicAuthData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BasicAuthData) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type JWTData struct {
	Secret string `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
}

func (m *JWTData) Reset()                    { *m = JWTData{} }
func (m *JWTData) String() string            { return proto.CompactTextString(m) }
func (*JWTData) ProtoMessage()               {}
func (*JWTData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *JWTData) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type Monitor struct {
	TriggerLimits []float64 `protobuf:"fixed64,1,rep,packed,name=trigger_limits,json=triggerLimits" json:"trigger_limits,omitempty"`
}

func (m *Monitor) Reset()                    { *m = Monitor{} }
func (m *Monitor) String() string            { return proto.CompactTextString(m) }
func (*Monitor) ProtoMessage()               {}
func (*Monitor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Monitor) GetTriggerLimits() []float64 {
	if m != nil {
		return m.TriggerLimits
	}
	return nil
}

type SessionState struct {
	LastCheck               int64                        `protobuf:"varint,1,opt,name=last_check,json=lastCheck" json:"last_check,omitempty"`
	Allowance               float64                      `protobuf:"fixed64,2,opt,name=allowance" json:"allowance,omitempty"`
	Rate                    float64                      `protobuf:"fixed64,3,opt,name=rate" json:"rate,omitempty"`
	Per                     float64                      `protobuf:"fixed64,4,opt,name=per" json:"per,omitempty"`
	Expires                 int64                        `protobuf:"varint,5,opt,name=expires" json:"expires,omitempty"`
	QuotaMax                int64                        `protobuf:"varint,6,opt,name=quota_max,json=quotaMax" json:"quota_max,omitempty"`
	QuotaRenews             int64                        `protobuf:"varint,7,opt,name=quota_renews,json=quotaRenews" json:"quota_renews,omitempty"`
	QuotaRemaining          int64                        `protobuf:"varint,8,opt,name=quota_remaining,json=quotaRemaining" json:"quota_remaining,omitempty"`
	QuotaRenewalRate        int64                        `protobuf:"varint,9,opt,name=quota_renewal_rate,json=quotaRenewalRate" json:"quota_renewal_rate,omitempty"`
	AccessRights            map[string]*AccessDefinition `protobuf:"bytes,10,rep,name=access_rights,json=accessRights" json:"access_rights,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OrgId                   string                       `protobuf:"bytes,11,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	OauthClientId           string                       `protobuf:"bytes,12,opt,name=oauth_client_id,json=oauthClientId" json:"oauth_client_id,omitempty"`
	OauthKeys               map[string]string            `protobuf:"bytes,13,rep,name=oauth_keys,json=oauthKeys" json:"oauth_keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BasicAuthData           *BasicAuthData               `protobuf:"bytes,14,opt,name=basic_auth_data,json=basicAuthData" json:"basic_auth_data,omitempty"`
	JwtData                 *JWTData                     `protobuf:"bytes,15,opt,name=jwt_data,json=jwtData" json:"jwt_data,omitempty"`
	HmacEnabled             bool                         `protobuf:"varint,16,opt,name=hmac_enabled,json=hmacEnabled" json:"hmac_enabled,omitempty"`
	HmacSecret              string                       `protobuf:"bytes,17,opt,name=hmac_secret,json=hmacSecret" json:"hmac_secret,omitempty"`
	IsInactive              bool                         `protobuf:"varint,18,opt,name=is_inactive,json=isInactive" json:"is_inactive,omitempty"`
	ApplyPolicyId           string                       `protobuf:"bytes,19,opt,name=apply_policy_id,json=applyPolicyId" json:"apply_policy_id,omitempty"`
	DataExpires             int64                        `protobuf:"varint,20,opt,name=data_expires,json=dataExpires" json:"data_expires,omitempty"`
	Monitor                 *Monitor                     `protobuf:"bytes,21,opt,name=monitor" json:"monitor,omitempty"`
	EnableDetailedRecording bool                         `protobuf:"varint,22,opt,name=enable_detailed_recording,json=enableDetailedRecording" json:"enable_detailed_recording,omitempty"`
	Metadata                string                       `protobuf:"bytes,23,opt,name=metadata" json:"metadata,omitempty"`
	Tags                    []string                     `protobuf:"bytes,24,rep,name=tags" json:"tags,omitempty"`
	Alias                   string                       `protobuf:"bytes,25,opt,name=alias" json:"alias,omitempty"`
	LastUpdated             string                       `protobuf:"bytes,26,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
	IdExtractorDeadline     int64                        `protobuf:"varint,27,opt,name=id_extractor_deadline,json=idExtractorDeadline" json:"id_extractor_deadline,omitempty"`
	SessionLifetime         int64                        `protobuf:"varint,28,opt,name=session_lifetime,json=sessionLifetime" json:"session_lifetime,omitempty"`
}

func (m *SessionState) Reset()                    { *m = SessionState{} }
func (m *SessionState) String() string            { return proto.CompactTextString(m) }
func (*SessionState) ProtoMessage()               {}
func (*SessionState) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *SessionState) GetLastCheck() int64 {
	if m != nil {
		return m.LastCheck
	}
	return 0
}

func (m *SessionState) GetAllowance() float64 {
	if m != nil {
		return m.Allowance
	}
	return 0
}

func (m *SessionState) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *SessionState) GetPer() float64 {
	if m != nil {
		return m.Per
	}
	return 0
}

func (m *SessionState) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *SessionState) GetQuotaMax() int64 {
	if m != nil {
		return m.QuotaMax
	}
	return 0
}

func (m *SessionState) GetQuotaRenews() int64 {
	if m != nil {
		return m.QuotaRenews
	}
	return 0
}

func (m *SessionState) GetQuotaRemaining() int64 {
	if m != nil {
		return m.QuotaRemaining
	}
	return 0
}

func (m *SessionState) GetQuotaRenewalRate() int64 {
	if m != nil {
		return m.QuotaRenewalRate
	}
	return 0
}

func (m *SessionState) GetAccessRights() map[string]*AccessDefinition {
	if m != nil {
		return m.AccessRights
	}
	return nil
}

func (m *SessionState) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *SessionState) GetOauthClientId() string {
	if m != nil {
		return m.OauthClientId
	}
	return ""
}

func (m *SessionState) GetOauthKeys() map[string]string {
	if m != nil {
		return m.OauthKeys
	}
	return nil
}

func (m *SessionState) GetBasicAuthData() *BasicAuthData {
	if m != nil {
		return m.BasicAuthData
	}
	return nil
}

func (m *SessionState) GetJwtData() *JWTData {
	if m != nil {
		return m.JwtData
	}
	return nil
}

func (m *SessionState) GetHmacEnabled() bool {
	if m != nil {
		return m.HmacEnabled
	}
	return false
}

func (m *SessionState) GetHmacSecret() string {
	if m != nil {
		return m.HmacSecret
	}
	return ""
}

func (m *SessionState) GetIsInactive() bool {
	if m != nil {
		return m.IsInactive
	}
	return false
}

func (m *SessionState) GetApplyPolicyId() string {
	if m != nil {
		return m.ApplyPolicyId
	}
	return ""
}

func (m *SessionState) GetDataExpires() int64 {
	if m != nil {
		return m.DataExpires
	}
	return 0
}

func (m *SessionState) GetMonitor() *Monitor {
	if m != nil {
		return m.Monitor
	}
	return nil
}

func (m *SessionState) GetEnableDetailedRecording() bool {
	if m != nil {
		return m.EnableDetailedRecording
	}
	return false
}

func (m *SessionState) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *SessionState) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SessionState) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SessionState) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

func (m *SessionState) GetIdExtractorDeadline() int64 {
	if m != nil {
		return m.IdExtractorDeadline
	}
	return 0
}

func (m *SessionState) GetSessionLifetime() int64 {
	if m != nil {
		return m.SessionLifetime
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessSpec)(nil), "coprocess.AccessSpec")
	proto.RegisterType((*AccessDefinition)(nil), "coprocess.AccessDefinition")
	proto.RegisterType((*BasicAuthData)(nil), "coprocess.BasicAuthData")
	proto.RegisterType((*JWTData)(nil), "coprocess.JWTData")
	proto.RegisterType((*Monitor)(nil), "coprocess.Monitor")
	proto.RegisterType((*SessionState)(nil), "coprocess.SessionState")
}

func init() { proto.RegisterFile("coprocess_session_state.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x55, 0x5d, 0x4f, 0x1b, 0x47,
	0x14, 0x95, 0x31, 0x60, 0xef, 0xb5, 0x0d, 0x64, 0x02, 0xc9, 0x00, 0x89, 0x6a, 0x2c, 0x35, 0x75,
	0xa4, 0x14, 0xb5, 0xf4, 0x05, 0x45, 0x95, 0xda, 0x34, 0xf0, 0xe0, 0x36, 0x49, 0xab, 0xa5, 0x51,
	0x5f, 0x2a, 0x8d, 0x2e, 0xbb, 0x37, 0xf6, 0x84, 0xfd, 0xea, 0xcc, 0x18, 0xe3, 0xbf, 0xd2, 0x3f,
	0xd8, 0xbf, 0x51, 0xcd, 0xdd, 0x59, 0x30, 0xa2, 0x7d, 0x9b, 0x7b, 0xce, 0x99, 0xeb, 0x33, 0x33,
	0x67, 0xaf, 0xe1, 0x79, 0x52, 0x56, 0xa6, 0x4c, 0xc8, 0x5a, 0x65, 0xc9, 0x5a, 0x5d, 0x16, 0xca,
	0x3a, 0x74, 0x74, 0x5c, 0x99, 0xd2, 0x95, 0x22, 0xba, 0xa5, 0x47, 0xa7, 0x00, 0x6f, 0x12, 0xbf,
	0xba, 0xa8, 0x28, 0x11, 0x3b, 0xd0, 0x9e, 0x9b, 0x4c, 0xb6, 0x86, 0xad, 0x71, 0x14, 0xfb, 0xa5,
	0x90, 0xd0, 0xc9, 0xc9, 0xcd, 0xca, 0xd4, 0xca, 0xb5, 0x61, 0x7b, 0x1c, 0xc5, 0x4d, 0x39, 0xfa,
	0xbb, 0x05, 0x3b, 0xf5, 0xd6, 0x33, 0xfa, 0xa4, 0x0b, 0xed, 0x74, 0x59, 0x88, 0x7d, 0xe8, 0x62,
	0xa5, 0x55, 0x81, 0x39, 0x85, 0x2e, 0x1d, 0xac, 0xf4, 0x07, 0xcc, 0x49, 0xec, 0xc1, 0xa6, 0xa7,
	0x74, 0x2a, 0xd7, 0x98, 0xd8, 0xc0, 0x4a, 0x4f, 0x52, 0x71, 0x00, 0xdd, 0x6b, 0x32, 0xde, 0xa2,
	0x95, 0x6d, 0xfe, 0x85, 0xdb, 0x5a, 0x9c, 0x42, 0x1f, 0xb3, 0xac, 0x5c, 0x50, 0xaa, 0xe6, 0x26,
	0xb3, 0x72, 0x7d, 0xd8, 0x1e, 0xf7, 0x4e, 0xf6, 0x8e, 0x6f, 0xed, 0x1f, 0xdf, 0x79, 0x8f, 0x7b,
	0x41, 0xfa, 0xd1, 0x64, 0x76, 0xf4, 0x03, 0x0c, 0x7e, 0x42, 0xab, 0x93, 0x37, 0x73, 0x37, 0x3b,
	0x43, 0x87, 0xfe, 0x67, 0x2a, 0xb4, 0x76, 0x51, 0x9a, 0x34, 0x18, 0xbb, 0xad, 0x85, 0x80, 0xf5,
	0x19, 0xda, 0x59, 0xf0, 0xc5, 0xeb, 0xd1, 0x11, 0x74, 0x7e, 0xfe, 0xe3, 0x77, 0xde, 0xfa, 0x04,
	0x36, 0x2d, 0x25, 0x86, 0x5c, 0xd8, 0x18, 0xaa, 0xd1, 0x37, 0xd0, 0x79, 0x5f, 0x16, 0xda, 0x95,
	0x46, 0x7c, 0x09, 0x5b, 0xce, 0xe8, 0xe9, 0x94, 0x8c, 0xca, 0x74, 0xae, 0x9d, 0x95, 0xad, 0x61,
	0x7b, 0xdc, 0x8a, 0x07, 0x01, 0x7d, 0xc7, 0xe0, 0xe8, 0x9f, 0x08, 0xfa, 0x17, 0xf5, 0x7b, 0x5c,
	0xf8, 0xe7, 0x10, 0xcf, 0x01, 0x32, 0xb4, 0x4e, 0x25, 0x33, 0x4a, 0xae, 0xb8, 0x7d, 0x3b, 0x8e,
	0x3c, 0xf2, 0xd6, 0x03, 0xe2, 0x19, 0x44, 0x7c, 0x28, 0x2c, 0x12, 0x62, 0x77, 0xad, 0xf8, 0x0e,
	0xf0, 0xb6, 0x0d, 0x3a, 0x92, 0x6d, 0x26, 0x78, 0xed, 0x1f, 0xb0, 0x22, 0x23, 0xd7, 0x19, 0xf2,
	0x4b, 0xff, 0x80, 0x74, 0x53, 0x69, 0x43, 0x56, 0x6e, 0x70, 0xff, 0xa6, 0x14, 0x87, 0x10, 0xfd,
	0x35, 0x2f, 0x1d, 0xaa, 0x1c, 0x6f, 0xe4, 0x26, 0x73, 0x5d, 0x06, 0xde, 0xe3, 0x8d, 0x38, 0x82,
	0x7e, 0x4d, 0x1a, 0x2a, 0x68, 0x61, 0x65, 0x87, 0xf9, 0x1e, 0x63, 0x31, 0x43, 0xe2, 0x2b, 0xd8,
	0x6e, 0x24, 0x39, 0xea, 0x42, 0x17, 0x53, 0xd9, 0x65, 0xd5, 0x56, 0x50, 0x05, 0x54, 0xbc, 0x02,
	0xb1, 0xd2, 0x0b, 0x33, 0xc5, 0xb6, 0x23, 0xd6, 0xee, 0xdc, 0x75, 0xc4, 0x2c, 0xf6, 0x47, 0xf8,
	0x00, 0x03, 0xe4, 0x57, 0x55, 0x46, 0x4f, 0x67, 0xce, 0x4a, 0xe0, 0x57, 0x7f, 0xb9, 0xf2, 0xea,
	0xab, 0x77, 0x18, 0x22, 0x10, 0xb3, 0xf6, 0xbc, 0x70, 0x66, 0x19, 0xf7, 0x71, 0x05, 0xf2, 0xb9,
	0x2b, 0xcd, 0xd4, 0xe7, 0xae, 0x57, 0xe7, 0xae, 0x34, 0xd3, 0x49, 0x2a, 0x5e, 0xc0, 0x76, 0x89,
	0x73, 0x37, 0x53, 0x49, 0xa6, 0xa9, 0x70, 0x9e, 0xef, 0x33, 0x3f, 0x60, 0xf8, 0x2d, 0xa3, 0x93,
	0x54, 0x9c, 0x03, 0xd4, 0xba, 0x2b, 0x5a, 0x5a, 0x39, 0x60, 0x2f, 0x2f, 0xfe, 0xcf, 0xcb, 0xaf,
	0x5e, 0xf9, 0x0b, 0x2d, 0x83, 0x91, 0xa8, 0x6c, 0x6a, 0xf1, 0x23, 0x6c, 0x5f, 0xfa, 0x40, 0x2a,
	0xee, 0x95, 0xa2, 0x43, 0xb9, 0x35, 0x6c, 0x8d, 0x7b, 0x27, 0x72, 0xa5, 0xd7, 0xbd, 0xc8, 0xc6,
	0x83, 0xcb, 0x7b, 0x09, 0xfe, 0x1a, 0xba, 0x9f, 0x17, 0xae, 0xde, 0xba, 0xcd, 0x5b, 0xc5, 0xca,
	0xd6, 0x10, 0xd6, 0xb8, 0xf3, 0x79, 0xe1, 0x58, 0x7e, 0x04, 0xfd, 0x59, 0x8e, 0x89, 0xa2, 0x02,
	0x2f, 0x33, 0x4a, 0xe5, 0xce, 0xb0, 0x35, 0xee, 0xc6, 0x3d, 0x8f, 0x9d, 0xd7, 0x90, 0xf8, 0x02,
	0xb8, 0x54, 0x21, 0xdd, 0x8f, 0xf8, 0xf8, 0xe0, 0xa1, 0x0b, 0x46, 0xbc, 0x40, 0x5b, 0xa5, 0x0b,
	0x4c, 0x9c, 0xbe, 0x26, 0x29, 0xb8, 0x05, 0x68, 0x3b, 0x09, 0x88, 0xbf, 0x44, 0xac, 0xaa, 0x6c,
	0xa9, 0xaa, 0x32, 0xd3, 0xc9, 0xd2, 0x5f, 0xe2, 0xe3, 0xfa, 0x12, 0x19, 0xfe, 0x8d, 0xd1, 0x49,
	0xea, 0xcd, 0x78, 0xdf, 0xaa, 0x49, 0xe2, 0x6e, 0x9d, 0x26, 0x8f, 0x9d, 0x87, 0x34, 0xbe, 0x82,
	0x4e, 0x5e, 0x7f, 0x4d, 0x72, 0xef, 0xc1, 0xe9, 0xc2, 0x77, 0x16, 0x37, 0x12, 0xf1, 0x1a, 0xf6,
	0xeb, 0x83, 0xa9, 0x94, 0x1c, 0xea, 0x8c, 0x52, 0x65, 0x28, 0x29, 0x4d, 0xea, 0x53, 0xf8, 0x84,
	0x7d, 0x3e, 0xad, 0x05, 0x67, 0x81, 0x8f, 0x1b, 0xda, 0x8f, 0x82, 0x9c, 0x1c, 0xf2, 0x45, 0x3e,
	0xad, 0x47, 0x41, 0x53, 0xfb, 0x6f, 0xca, 0xe1, 0xd4, 0x4a, 0xc9, 0x93, 0x88, 0xd7, 0x62, 0x17,
	0x36, 0x30, 0xd3, 0x68, 0xe5, 0x7e, 0x98, 0x5b, 0xbe, 0xf0, 0x47, 0xe2, 0x4f, 0x77, 0x5e, 0xa5,
	0xe8, 0x28, 0x95, 0x07, 0x4c, 0xf6, 0x3c, 0xf6, 0xb1, 0x86, 0xc4, 0x09, 0xec, 0xe9, 0x54, 0xd1,
	0x8d, 0x33, 0x98, 0xb8, 0xd2, 0xa8, 0x94, 0x30, 0xcd, 0x74, 0x41, 0xf2, 0x90, 0x8f, 0xff, 0x58,
	0xa7, 0xe7, 0x0d, 0x77, 0x16, 0x28, 0xf1, 0x12, 0x76, 0x9a, 0x89, 0x9d, 0xe9, 0x4f, 0xe4, 0x74,
	0x4e, 0xf2, 0x19, 0xcb, 0xb7, 0x03, 0xfe, 0x2e, 0xc0, 0x07, 0x7f, 0xc2, 0xa3, 0x07, 0xd9, 0xf7,
	0x03, 0xe0, 0x8a, 0x96, 0xcd, 0x04, 0xbf, 0xa2, 0xa5, 0xf8, 0x16, 0x36, 0xae, 0x31, 0x9b, 0xd7,
	0x03, 0xa4, 0x77, 0x72, 0xf8, 0x60, 0x7a, 0xde, 0x8d, 0xef, 0xb8, 0x56, 0xbe, 0x5e, 0x3b, 0x6d,
	0x1d, 0x7c, 0x0f, 0x5b, 0xf7, 0xd3, 0xfc, 0x1f, 0xad, 0x77, 0x57, 0x5b, 0x47, 0x2b, 0xbb, 0x2f,
	0x37, 0xf9, 0x8f, 0xe6, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x72, 0xdf, 0x5f, 0x89,
	0x06, 0x00, 0x00,
}