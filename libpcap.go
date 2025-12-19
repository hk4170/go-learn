package main

/*
#cgo LDFLAGS: -lpcap  
// 链接 libpcap 库
#include <pcap.h>
#include <stdio.h>
#include <stdlib.h>

// 自定义 C 回调函数，处理抓到的包
void packet_handler(u_char *user, const struct pcap_pkthdr *h, const u_char *bytes) {
    printf("抓到包，长度: %d\n", h->len);
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    var errbuf *C.char = (*C.char)(C.calloc(1, C.PCAP_ERRBUF_SIZE))
    defer C.free(unsafe.Pointer(errbuf))

    // 指定网卡（Mac 常用 en0，Linux 常用 eth0）
    device := C.CString("eth0")
    defer C.free(unsafe.Pointer(device))

    // 打开网卡，进入抓包模式
    handle := C.pcap_open_live(device, C.int(65535), C.int(1), C.int(1000), errbuf)
    if handle == nil {
        fmt.Printf("打开网卡失败: %s\n", C.GoString(errbuf))
        return
    }
    defer C.pcap_close(handle)

    // 循环抓包，调用回调函数
    fmt.Println("开始抓包...")
    C.pcap_loop(handle, C.int(-1), C.pcap_handler(C.packet_handler), nil)
}