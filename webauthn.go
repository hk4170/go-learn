package main

import (
    "fmt"
    "net/http"

    "github.com/go-webauthn/webauthn"
    "github.com/go-webauthn/webauthn/protocol"
    "github.com/go-webauthn/webauthn/webauthn"
)

// 定义用户结构体（需实现webauthn.User接口）
type User struct {
    ID          []byte
    Name        string
    DisplayName string
    Credentials []webauthn.Credential // 存储用户的Passkey凭证
}

func (u *User) WebAuthnID() []byte {
    return u.ID
}

func (u *User) WebAuthnName() string {
    return u.Name
}

func (u *User) WebAuthnDisplayName() string {
    return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
    return u.Credentials
}

func (u *User) WebAuthnIcon() string {
    return ""
}

var (
    wa *webauthn.WebAuthn
    // 模拟用户存储
    userStore = make(map[string]*User)
)

func init() {
    // 初始化WebAuthn（Passkey）配置
    config := &webauthn.Config{
        RPID:          "your-domain.com", // 你的服务端域名
        RPName:        "Your App Name",   // 应用名称
        RPOrigin:      "https://your-domain.com", // 服务端源地址（需HTTPS）
        Attestation:   protocol.AttestationNone,
        AuthenticatorSelection: protocol.AuthenticatorSelection{
            ResidentKey: protocol.ResidentKeyRequired, // 要求Passkey（常驻密钥）
            UserVerification: protocol.VerificationRequired,
        },
    }

    var err error
    wa, err = webauthn.New(config)
    if err != nil {
        panic(fmt.Sprintf("初始化webauthn失败: %v", err))
    }

    // 模拟创建一个测试用户
    testUser := &User{
        ID:          []byte("user123"),
        Name:        "test@your-domain.com",
        DisplayName: "Test User",
        Credentials: []webauthn.Credential{},
    }
    userStore["user123"] = testUser
}

// Passkey注册接口
func registerHandler(w http.ResponseWriter, r *http.Request) {
    user := userStore["user123"] // 获取当前用户

    // 1. 生成注册选项
    registerOptions, sessionData, err := wa.BeginRegistration(user)
    if err != nil {
        http.Error(w, fmt.Sprintf("生成注册选项失败: %v", err), http.StatusInternalServerError)
        return
    }

    // 2. 存储sessionData（需关联用户，用于后续验证）
    // （实际项目中需存储在会话或缓存中，如Redis）

    // 3. 返回注册选项给客户端（JSON格式）
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(registerOptions); err != nil {
        http.Error(w, fmt.Sprintf("返回注册选项失败: %v", err), http.StatusInternalServerError)
        return
    }
}

// Passkey认证接口
func authenticateHandler(w http.ResponseWriter, r *http.Request) {
    user := userStore["user123"] // 获取当前用户

    // 1. 生成认证选项
    authOptions, sessionData, err := wa.BeginAuthentication(user)
    if err != nil {
        http.Error(w, fmt.Sprintf("生成认证选项失败: %v", err), http.StatusInternalServerError)
        return
    }

    // 2. 存储sessionData（用于后续验证）
    // （实际项目中需存储在会话或缓存中）

    // 3. 返回认证选项给客户端（JSON格式）
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(authOptions); err != nil {
        http.Error(w, fmt.Sprintf("返回认证选项失败: %v", err), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/passkey/register", registerHandler)
    http.HandleFunc("/passkey/authenticate", authenticateHandler)
    fmt.Println("服务启动在 :8080")
    // 实际生产环境需使用HTTPS（Passkey要求安全上下文）
    if err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil); err != nil {
        panic(fmt.Sprintf("服务启动失败: %v", err)))
    }
}