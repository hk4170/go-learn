package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	Name string
	Age  int
}
// EncryptGobWithPassword 使用密码加密 gob 数据
// 返回值：加密后的数据（含 nonce），以及 salt，都必须保存好
func EncryptGobWithPassword(data any, password string) ([]byte, []byte, error) {
	// --- 1. 用 gob 编码数据 ---
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, nil, fmt.Errorf("gob encode error: %v", err)
	}
	plaintext := buf.Bytes()

	// --- 2. 生成随机 salt（推荐 16 字节）---
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, nil, fmt.Errorf("生成 salt 失败: %v", err)
	}

	// --- 3. 用 PBKDF2 从密码 + salt 派生密钥（32 字节 = AES-256）---
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// --- 4. 用 AES-GCM 加密 ---
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("aes new cipher 失败: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, fmt.Errorf("gcm new 失败: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize()) // 通常为 12 字节
	if _, err := rand.Read(nonce); err != nil {
		return nil, nil, fmt.Errorf("生成 nonce 失败: %v", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// 注意：这里返回的是 密文（含 nonce）和 salt
	// 实际使用时，可以把 salt 和 ciphertext 一起存储（比如拼接或分开存）
	return ciphertext, salt, nil
}
// DecryptGobWithPassword 解密数据并解码为原始数据结构
func DecryptGobWithPassword(encryptedData []byte, salt []byte, password string, out any) error {
	// --- 1. 用相同密码 + salt 派生密钥 ---
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// --- 2. AES-GCM 解密 ---
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("aes new cipher 失败: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("gcm new 失败: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return errors.New("加密数据太短，无法提取 nonce")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("gcm 打开失败: %v", err)
	}

	// --- 3. 用 gob 解码 ---
	buf := bytes.NewBuffer(plaintext)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("gob 解码失败: %v", err)
	}

	return nil
}
func main() {
	// 原始数据
	user := User{Name: "Bob", Age: 30}
	password := "mySecurepass!#"

	// --- 加密 ---
	encryptedData, salt, err := EncryptGobWithPassword(user, password)
	if err != nil {
		panic(err)
	}

	fmt.Printf("加密成功！\nSalt (hex): %s\n密文 (hex): %s\n",
		hex.EncodeToString(salt),
		hex.EncodeToString(encryptedData))

	// --- 解密 ---
	var decryptedUser User
	err = DecryptGobWithPassword(encryptedData, salt, password, &decryptedUser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("解密成功！用户信息: %+v\n", decryptedUser)

	// --- 测试错误密码 ---
	err = DecryptGobWithPassword(encryptedData, salt, "wrongPassword", &decryptedUser)
	if err != nil {
		fmt.Println("使用错误密码解密（预期失败）:", err)
	}
}
