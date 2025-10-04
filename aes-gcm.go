package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
)
// 示例数据结构
type User struct {
	Name string
	Age  int
}

// 加密函数：将任意数据（比如结构体）先 gob 编码，再用 AES-GCM 加密
func EncryptGobWithAES(data any, key []byte) ([]byte, error) {
	// --- 第一步：用 gob 编码数据 ---
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, fmt.Errorf("gob encode error: %v", err)
	}
	plaintext := buf.Bytes()

	// --- 第二步：使用 AES-GCM 加密 ---

	// 检查密钥长度：AES-128（16字节）、AES-192（24字节）、AES-256（32字节）
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid AES key size (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes new cipher error: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("gcm new error: %v", err)
	}

	// 创建随机 nonce（推荐长度是 gcm.NonceSize()，通常为 12 字节）
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	// 加密：nonce + 密文
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

// 解密函数：先解密 AES-GCM 数据，再 gob 解码为原始数据
func DecryptGobWithAES(encryptedData []byte, key []byte, out any) error {
	// --- 第一步：使用 AES-GCM 解密 ---

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("aes new cipher error: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("gcm new error: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return errors.New("encrypted data too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("gcm open error: %v", err)
	}

	// --- 第二步：用 gob 解码 ---
	buf := bytes.NewBuffer(plaintext)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("gob decode error: %v", err)
	}

	return nil
}

func main() {
	// 原始数据
	user := User{Name: "Alice", Age: 25}

	// AES 密钥（必须是 16, 24 或 32 字节，对应 AES-128, AES-192, AES-256）
	// 注意：在实际项目中，密钥应安全生成/存储，不要硬编码！
	key := []byte("this_is_a_32byte_key_for_aes_256!!") // 32 字节 => AES-256

	// 加密
	encrypted, err := EncryptGobWithAES(user, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后的数据（十六进制）: %x\n", encrypted)

	// 解密
	var decrypted User
	err = DecryptGobWithAES(encrypted, key, &decrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后的数据: %+v\n", decrypted)
}
