# todo zpq
# todo python 基础函数
# todo iterable 指的大概是数组之类的
# a = [0, 1, 2, 3, 4]
# for k in a:
#     print("key ==> " + str(a.index(k)) + "; value ==> " + str(k))

# abs() 返回一个数的绝对值。实参可以是整数或浮点数。如果实参是一个复数，返回它的模
absValue = 3 / 10
print(abs(absValue))
# 0.3

# all(iterable) 如果 iterable 的所有元素为真（或迭代器为空），返回 True
# def all(iterable):
#     for element in iterable:
#         if not element:
#             return False
#     return True
allValue = [5, 4, 3, 2, 1]
print(all(allValue))
# True

# any(iterable) 如果 iterable 的任一元素为真则返回 True。 如果迭代器为空，返回 False
# def any(iterable):
#     for element in iterable:
#         if element:
#             return True
#     return False
anyValue = [5, 4, 3, 2, 1]
print(any(anyValue))
# True

# ascii(object)
# 就像函数 repr()，返回一个对象可打印的字符串，但是 repr() 返回的字符串中非 ASCII 编码的字符，会使用 \x、\u 和 \U 来转义。
# 生成的字符串和 Python 2 的 repr() 返回的结果相似
asciiValue = "中国"
print(ascii(asciiValue))
# '\u4e2d\u56fd'

# bool 判断是否为真
boolValue = "ashdj"
print(bool(boolValue))

# bytearray
bytearrayValue = "一二"
print(bytearray(bytearrayValue, 'utf-8'))
