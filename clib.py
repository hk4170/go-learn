import ctypes
import os

def main():
    lib_path = "./clib.so" 
    if not os.path.exists(lib_path):
        print(f"错误：动态库文件 {lib_path} 不存在，请检查编译步骤！")
        return
    
    # 加载动态库
    lib = ctypes.CDLL(lib_path)
    
    # 2. 配置 C 函数的参数类型和返回值类型（关键：避免类型不匹配崩溃）
    # 告诉 ctypes：CalculateSum 接收 1 个 C.char* 类型参数，返回 C.int 类型
    lib.CalculateSum.argtypes = [ctypes.c_char_p]  # c_char_p 对应 C 的 char*
    lib.CalculateSum.restype = ctypes.c_int       # c_int 对应 C 的 int
    
    # 3. 准备输入数据（需转为 C 兼容格式）
    # Python 字符串 → 字节流（C 语言只能识别字节流）
    input_data = "10,20,30,40"  # 要传递给 Go 的数字列表（逗号分隔）
    c_input = input_data.encode("utf-8")  # 转为 bytes，对应 C 的 char*
    
    # 4. 调用 Go 动态库中的函数
    result = lib.CalculateSum(c_input)
    
    # 5. 处理返回结果
    print(f"Python 传入数据：{input_data}")
    print(f"Go 动态库计算结果（总和）：{result}")  # 输出：Go 动态库计算结果（总和）：100

if __name__ == "__main__":
    main()