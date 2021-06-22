<p align="center">
<img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/zhengpeiqiang.png" style="border-radius:50%;margin: 0 auto;" width="20%" />
</p>

<h1 align="center">
window10 + python3.6.5 使用 face_recognition
</h1>

### organization

<a href="https://github.com/zhengpeiqiang"><img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/LCLCF_circle.png" style="width:60px;height:60px;margin: 0 auto;" width="8%" /></a>
**LCLCF**

### <font color='red'> ** 亲测 python 3.9 版本的几乎没办法引入 dlib 包；并且有说要引入vs 的也是很麻烦的，引入vs之后去 install dlib 非常的吃cpu，cpu瞬间100% **</font>

步骤1
```
pip install cmake
```
步骤2
```
python -m pip install --upgrade pip
```
步骤3
```
pip install wheel
```
步骤4
```
pip install boost
```
步骤5

- 下载 dlib-19.7.0-cp36-cp36m-win_amd64.whl

- [dlib地址](https://pypi.org/simple/dlib/)

- 找到对应版本

- 任意放到一个文件夹下，在该文件夹下命令行执行下面的命令
```
pip install dlib-19.7.0-cp36-cp36m-win_amd64.whl
```
步骤6 大功告成！！！
```
pip install face_recognition
```
