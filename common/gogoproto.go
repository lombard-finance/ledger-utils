package common

// GogoprotoCustomType is an interface that can be implemented by types that need custom serialization and
// deserialization logic for use with gogoproto. Types implementing this interface must provide methods for
// marshaling, unmarshaling, and determining the size of the serialized data.
// [gogoproto custom types](https://github.com/cosmos/gogoproto/blob/70f82eb45331b1eb0db349b02a73a9d8e914305f/proto/custom_gogo.go#L33-L37)
type GogoprotoCustomType interface {
	// Marshal serializes the instance into a byte array
	Marshal() ([]byte, error)

	// MarshalTo serializes the instance into the provided byte slice returning the number of bytes written
	// this is not present in the referenced interface but it is required when building the chain
	MarshalTo(data []byte) (int, error)

	// Unmarshal deserializes the instance from a byte array
	Unmarshal(data []byte) error

	// Size returns the size of the instance in bytes
	Size() int
}
