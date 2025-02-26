// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package migration

import (
	_ "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "github.com/pokt-network/poktroll/api/poktroll/shared"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_EventImportMorseClaimableAccounts                          protoreflect.MessageDescriptor
	fd_EventImportMorseClaimableAccounts_created_at_height        protoreflect.FieldDescriptor
	fd_EventImportMorseClaimableAccounts_morse_account_state_hash protoreflect.FieldDescriptor
	fd_EventImportMorseClaimableAccounts_num_accounts             protoreflect.FieldDescriptor
)

func init() {
	file_poktroll_migration_event_proto_init()
	md_EventImportMorseClaimableAccounts = File_poktroll_migration_event_proto.Messages().ByName("EventImportMorseClaimableAccounts")
	fd_EventImportMorseClaimableAccounts_created_at_height = md_EventImportMorseClaimableAccounts.Fields().ByName("created_at_height")
	fd_EventImportMorseClaimableAccounts_morse_account_state_hash = md_EventImportMorseClaimableAccounts.Fields().ByName("morse_account_state_hash")
	fd_EventImportMorseClaimableAccounts_num_accounts = md_EventImportMorseClaimableAccounts.Fields().ByName("num_accounts")
}

var _ protoreflect.Message = (*fastReflection_EventImportMorseClaimableAccounts)(nil)

type fastReflection_EventImportMorseClaimableAccounts EventImportMorseClaimableAccounts

func (x *EventImportMorseClaimableAccounts) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventImportMorseClaimableAccounts)(x)
}

func (x *EventImportMorseClaimableAccounts) slowProtoReflect() protoreflect.Message {
	mi := &file_poktroll_migration_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventImportMorseClaimableAccounts_messageType fastReflection_EventImportMorseClaimableAccounts_messageType
var _ protoreflect.MessageType = fastReflection_EventImportMorseClaimableAccounts_messageType{}

type fastReflection_EventImportMorseClaimableAccounts_messageType struct{}

func (x fastReflection_EventImportMorseClaimableAccounts_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventImportMorseClaimableAccounts)(nil)
}
func (x fastReflection_EventImportMorseClaimableAccounts_messageType) New() protoreflect.Message {
	return new(fastReflection_EventImportMorseClaimableAccounts)
}
func (x fastReflection_EventImportMorseClaimableAccounts_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventImportMorseClaimableAccounts
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventImportMorseClaimableAccounts) Descriptor() protoreflect.MessageDescriptor {
	return md_EventImportMorseClaimableAccounts
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventImportMorseClaimableAccounts) Type() protoreflect.MessageType {
	return _fastReflection_EventImportMorseClaimableAccounts_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventImportMorseClaimableAccounts) New() protoreflect.Message {
	return new(fastReflection_EventImportMorseClaimableAccounts)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventImportMorseClaimableAccounts) Interface() protoreflect.ProtoMessage {
	return (*EventImportMorseClaimableAccounts)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventImportMorseClaimableAccounts) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CreatedAtHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.CreatedAtHeight)
		if !f(fd_EventImportMorseClaimableAccounts_created_at_height, value) {
			return
		}
	}
	if len(x.MorseAccountStateHash) != 0 {
		value := protoreflect.ValueOfBytes(x.MorseAccountStateHash)
		if !f(fd_EventImportMorseClaimableAccounts_morse_account_state_hash, value) {
			return
		}
	}
	if x.NumAccounts != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NumAccounts)
		if !f(fd_EventImportMorseClaimableAccounts_num_accounts, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventImportMorseClaimableAccounts) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		return x.CreatedAtHeight != int64(0)
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		return len(x.MorseAccountStateHash) != 0
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		return x.NumAccounts != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventImportMorseClaimableAccounts) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		x.CreatedAtHeight = int64(0)
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		x.MorseAccountStateHash = nil
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		x.NumAccounts = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventImportMorseClaimableAccounts) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		value := x.CreatedAtHeight
		return protoreflect.ValueOfInt64(value)
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		value := x.MorseAccountStateHash
		return protoreflect.ValueOfBytes(value)
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		value := x.NumAccounts
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventImportMorseClaimableAccounts) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		x.CreatedAtHeight = value.Int()
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		x.MorseAccountStateHash = value.Bytes()
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		x.NumAccounts = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventImportMorseClaimableAccounts) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		panic(fmt.Errorf("field created_at_height of message poktroll.migration.EventImportMorseClaimableAccounts is not mutable"))
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		panic(fmt.Errorf("field morse_account_state_hash of message poktroll.migration.EventImportMorseClaimableAccounts is not mutable"))
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		panic(fmt.Errorf("field num_accounts of message poktroll.migration.EventImportMorseClaimableAccounts is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventImportMorseClaimableAccounts) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.migration.EventImportMorseClaimableAccounts.created_at_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "poktroll.migration.EventImportMorseClaimableAccounts.morse_account_state_hash":
		return protoreflect.ValueOfBytes(nil)
	case "poktroll.migration.EventImportMorseClaimableAccounts.num_accounts":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.migration.EventImportMorseClaimableAccounts"))
		}
		panic(fmt.Errorf("message poktroll.migration.EventImportMorseClaimableAccounts does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventImportMorseClaimableAccounts) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in poktroll.migration.EventImportMorseClaimableAccounts", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventImportMorseClaimableAccounts) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventImportMorseClaimableAccounts) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventImportMorseClaimableAccounts) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventImportMorseClaimableAccounts) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventImportMorseClaimableAccounts)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.CreatedAtHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.CreatedAtHeight))
		}
		l = len(x.MorseAccountStateHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NumAccounts != 0 {
			n += 1 + runtime.Sov(uint64(x.NumAccounts))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventImportMorseClaimableAccounts)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.NumAccounts != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NumAccounts))
			i--
			dAtA[i] = 0x18
		}
		if len(x.MorseAccountStateHash) > 0 {
			i -= len(x.MorseAccountStateHash)
			copy(dAtA[i:], x.MorseAccountStateHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MorseAccountStateHash)))
			i--
			dAtA[i] = 0x12
		}
		if x.CreatedAtHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CreatedAtHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventImportMorseClaimableAccounts)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventImportMorseClaimableAccounts: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventImportMorseClaimableAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreatedAtHeight", wireType)
				}
				x.CreatedAtHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CreatedAtHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MorseAccountStateHash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MorseAccountStateHash = append(x.MorseAccountStateHash[:0], dAtA[iNdEx:postIndex]...)
				if x.MorseAccountStateHash == nil {
					x.MorseAccountStateHash = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NumAccounts", wireType)
				}
				x.NumAccounts = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NumAccounts |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: poktroll/migration/event.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventImportMorseClaimableAccounts is emitted when the MorseClaimableAccounts are created on-chain.
type EventImportMorseClaimableAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The height (on Shannon) at which the MorseAccountState was created on-chain.
	CreatedAtHeight int64 `protobuf:"varint,1,opt,name=created_at_height,json=createdAtHeight,proto3" json:"created_at_height,omitempty"`
	// The onchain computed sha256 hash of the entire MorseAccountState containing the MorseClaimableAccounts which were imported.
	MorseAccountStateHash []byte `protobuf:"bytes,2,opt,name=morse_account_state_hash,json=morseAccountStateHash,proto3" json:"morse_account_state_hash,omitempty"`
	// Number of claimable accounts (EOAs) collected from Morse state export
	// NOTE: Account balances include consolidated application and supplier actor stakes
	NumAccounts uint64 `protobuf:"varint,3,opt,name=num_accounts,json=numAccounts,proto3" json:"num_accounts,omitempty"`
}

func (x *EventImportMorseClaimableAccounts) Reset() {
	*x = EventImportMorseClaimableAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poktroll_migration_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventImportMorseClaimableAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventImportMorseClaimableAccounts) ProtoMessage() {}

// Deprecated: Use EventImportMorseClaimableAccounts.ProtoReflect.Descriptor instead.
func (*EventImportMorseClaimableAccounts) Descriptor() ([]byte, []int) {
	return file_poktroll_migration_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventImportMorseClaimableAccounts) GetCreatedAtHeight() int64 {
	if x != nil {
		return x.CreatedAtHeight
	}
	return 0
}

func (x *EventImportMorseClaimableAccounts) GetMorseAccountStateHash() []byte {
	if x != nil {
		return x.MorseAccountStateHash
	}
	return nil
}

func (x *EventImportMorseClaimableAccounts) GetNumAccounts() uint64 {
	if x != nil {
		return x.NumAccounts
	}
	return 0
}

var File_poktroll_migration_event_proto protoreflect.FileDescriptor

var file_poktroll_migration_event_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x72, 0x73, 0x65, 0x5f, 0x6f,
	0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a,
	0x21, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x72, 0x73,
	0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0xea,
	0xde, 0x1f, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x55, 0x0a, 0x18, 0x6d, 0x6f, 0x72, 0x73, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x1c, 0xea, 0xde, 0x1f, 0x18, 0x6d, 0x6f, 0x72,
	0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x52, 0x15, 0x6d, 0x6f, 0x72, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x33, 0x0a, 0x0c,
	0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x10, 0xea, 0xde, 0x1f, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x42, 0xb6, 0x01, 0xd8, 0xe2, 0x1e, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x4d, 0x58, 0xaa, 0x02, 0x12, 0x50, 0x6f, 0x6b, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02,
	0x12, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1e, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x3a,
	0x3a, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_poktroll_migration_event_proto_rawDescOnce sync.Once
	file_poktroll_migration_event_proto_rawDescData = file_poktroll_migration_event_proto_rawDesc
)

func file_poktroll_migration_event_proto_rawDescGZIP() []byte {
	file_poktroll_migration_event_proto_rawDescOnce.Do(func() {
		file_poktroll_migration_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_poktroll_migration_event_proto_rawDescData)
	})
	return file_poktroll_migration_event_proto_rawDescData
}

var file_poktroll_migration_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_poktroll_migration_event_proto_goTypes = []interface{}{
	(*EventImportMorseClaimableAccounts)(nil), // 0: poktroll.migration.EventImportMorseClaimableAccounts
}
var file_poktroll_migration_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_poktroll_migration_event_proto_init() }
func file_poktroll_migration_event_proto_init() {
	if File_poktroll_migration_event_proto != nil {
		return
	}
	file_poktroll_migration_morse_onchain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_poktroll_migration_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventImportMorseClaimableAccounts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poktroll_migration_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poktroll_migration_event_proto_goTypes,
		DependencyIndexes: file_poktroll_migration_event_proto_depIdxs,
		MessageInfos:      file_poktroll_migration_event_proto_msgTypes,
	}.Build()
	File_poktroll_migration_event_proto = out.File
	file_poktroll_migration_event_proto_rawDesc = nil
	file_poktroll_migration_event_proto_goTypes = nil
	file_poktroll_migration_event_proto_depIdxs = nil
}
