<p align="center">
<img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/zhengpeiqiang.png" style="border-radius:50%;margin: 0 auto;" width="20%" />
</p>

<h1 align="center">
快速编码系统
</h1>

### organization

<a href="https://github.com/zhengpeiqiang"><img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/LCLCF_circle.png" style="width:60px;height:60px;margin: 0 auto;" width="8%" /></a>
**LCLCF**

#### target

```
要达到效果，产品给予到产品经理，产品经理可以通过语言描述去实时新增功能，例如【页面做到支付界面，产品经理说一句，需要支持微信支付，那么就会新增微信支付的支持】
```

#### 预估需要涉及的功能步骤如下

```
1.声音转文字（这部分比较麻烦，直接略过，用文字输入到接口替代）
    1.1.声音转成语音文件（使用window的录音机，指令调用window的录音功能【开始录音】【结束录音】【查询录音文件地址】【录音文件提取】）
    1.2.语音文件转成文字（可以用市面的api做这个功能）
2.文字转代码
    2.1.文字语义理解
    2.2.语义匹配系统预设内容，新增相应功能
    2.3.相关功能curd要正确
3.测试
4.生成一份swagger文档
5.集成发布
```
