# excelThinker




# 资料
fyne官方文档https://go-circle.cn/fyne-press/v1.0/1-getting-started/hello-world.html

# windows环境配置
1. 安装msys2;
2. 启动UCRT64环境；
   pacman -Syu
   pacman -S git mingw-w64-x86_64-toolchain
   pacman -S mingw-w64-ucrt-x86_64-gcc
   gcc --version
3. D:\software\msys2\ucrt64\bin加到环境变量
4. UCRT64环境变量PATH加入Go的bin目录；


# 打包
fyne package -os windows -icon ttt.jpg



# 包划分
fyne将功能划分到多个子包中：

fyne.io/fyne：提供所有fyne应用程序代码共用的基础定义，包括数据类型和接口；

fyne.io/fyne/app：提供创建应用程序的 API；

fyne.io/fyne/canvas：提供Fyne使用的绘制 API；

fyne.io/fyne/dialog：提供对话框组件；

fyne.io/fyne/layout：提供多种界面布局；

fyne.io/fyne/widget：提供多种组件，fyne所有的窗体控件和交互元素都在这个子包中。


