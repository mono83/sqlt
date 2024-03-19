package hash

import "hash/crc32"

// CRC32 performs CRC32 hashing, producing unsigned
// 32-bit value that can be used in tables to provide
// efficient way to index strings.
func CRC32(s string) uint32 {
	c := crc32.New(crc32.IEEETable)
	_, _ = c.Write([]byte(s))
	return c.Sum32()
}
