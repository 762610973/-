# HTML

---------------------------------------------------

- 超文本标记语言

- 超文本指的是超链接，使用超链接

------------------------------------------------------

- 根标签`<html></html>`有两个子标签，分别是`<head></head>`和`<body></body>`,`<head></head>`里面的`<title></title>`是网页的标题

- 标签一般成对出现，但是也存在一些自结束标签`<img />,<input />`

  

  ### 简单概要 

```html
<!doctype html>
<html>
	<head>
	<meta charset="utf-8">
	<title>标签的属性</title>
	</head>
	<body>
	<!--
	注释不能多层嵌套
	属性，在标签（开始标签或自结束标签）中可以设置属性
	在结束标签中不能设置
	属性是一个名值对结构（x(名字) = y（结构））
	属性和标签或其他属性应该使用空格隔开
	属性应该根据文档中的规定来编写，有些属性有属性值，有些没有
	文档声明(doctype)
		--文档声明用来高速浏览器当前网页的版本
		--html5de文档声明
		<!doctype html>
		<!DOCTYPE HTML>
	编码：将字符转换为二进制码的过程成为编码
	解码： 将二进制码转换为字符的过程称为解码
	字符集：编码和解码所采用的规则成为字符集
	16进制较为常用，和二进制配合
	gbk，utf-8，asi
	-->
	<h1>这是我的<font color="red" size='3.1'>第二个</font>网页！</h1>
	</body>
</html>

```

### meta标签

```html
<!doctype html>
<html lang="zh">
<!-- 文档声明，声明当前网页的版本 -->
<html>
<!-- 标签就是元素，元素就是标签-->
	<head>
		<!--meta 标签用来设置网页的元数据，这里meta用来设置网页的字符集，避免乱码问题-->
		<meta charset="utf-8">
		<!--title中的内容会显示在浏览器的标题栏，搜索引擎会主要根据title中的内容来判断网页的主要内容-->
		<title>this is my third html test</title>
		<!--body标签是html的子元素，表示网页的主体，网页中所有的可见内容都应该在body里-->
		<body>
		</body>
	</head>
</html>
```

```html
<!doctype html>
<html lang="zh">
<html>
<head>
     <meta charset='UTF-8'>
    <!--
    	在网页中编写的多个空格默认情况会自动被浏览器解析为一个空格
    	不能书写一些特殊符号
    	比如：多个连续的空格、字母两侧的大于小于号
	-->
    <!--
        meta主要设置网页的一些元数据，元数据不给用户看
        charset 指定网页的字符集
        name指定网页的名称
        content指定的数据的内容
        keywords表示网站的关键字，可以同时指定多个关键字，关键字使用，隔开
        description用于指定网站的描述
        title标签的内容会作为搜索结果的超链接上的文字显示
    -->
    <meta name="keywords" content = "html5,前端，css3" charset=""UTF-8>
    <meta name= "description" content="这是我学习前端的第二个天">
    <meta http-equiv="refresh" content="3;url = https://www.baidu.com">
    <title>meta标签</title>
</head>
<body>
    <h1>一级标题</h1>
     <p>内容</p>
     <p>&gt;&lt;&copy;&ni;</p>
</body>
</html>
```

```html
<!doctype html>
<html lang="zh">
<html>
<head>
    <meta charset="UTF-8">
    <!--
        meta主要设置网页的一些元数据，元数据不给用户看
        charset 指定网页的字符集
        name指定网页的名称
        content指定的数据的内容
        keywords表示网站的关键字，可以同时指定多个关键字，关键字使用，隔开
        description用于指定网站的描述
        title标签的内容会作为搜索结果的超链接上的文字显示
    -->
    <meta name="keywords" content = "html5,前端，css3" charset=""UTF-8>
    <meta name= "description" content="这是我学习前端的第二个天">
    <!--<meta http-equiv="refresh" content="3;url = https://www.baidu.com">-->
    <title>meta标签</title>
</head>
<body>
    <!--在网页中HTML专门用来负责网页的结构，所以在使用html标签时，应该关注的是标签的语义，而不是他的样式-->
    <!--
        标题标签：一共有六级标签，重要性递减，h1最重要，h1在网页中的重要性仅次于title标签，一般情况下一个页面只会有一个h1
        一般情况下标题标签只会使用到h1~h3，h4~h6很少用
        标题标签是块元素，p标签也是一个块元素
        在页面中独占一行的元素称为块元素（block element）
        hgroup标签用来为标题分组，可以将一组相关的标题同时放入到hgroup
        在页面中不会独占一行的元素称为行内元素
    -->
    <hgroup>
        <h1>一级标题</h1>
        <h2>二级标题</h2>
    </hgroup>
    <h2>三级标题</h2>
    <p>在p标签中的内容就表示一个段落</p>
    <p>em标签用于表示语音音调的一个<em>加重</em></p>
    <p><strong>striong表示强调重要的内容</strong></p>
    鲁迅说：
    <blockquote>blockquote 表示一个长引用</blockquote>
    子曰<q> < q > 表示一个短引用</q>
    <br>
    <p>br标签表示换行</p>
</body>
</html>
```

### 语义化标签

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>布局标签</title>
</head>
<body>
    <!--
        header表示网页的头部
        footer表示网页的底部
        nav表示网页中的导航
        main表示网页的主题部分（一个网页中只有一个main）
        aside表示和主题相关的其他内容（侧边栏）
        article表示独立的文章
        section表示独立的区块，上边的标签都不能表示时使用section
        div没有语义。用来表示一个区块儿，目前是主要的布局元素
        span行内元素，没有任何的语义，一般用于在网页中选中文字
-->
    <header>aa</header>
    <main>bb</main>
    <nav>cc</nav>
    <aside>dd</aside>
    <article>ee</article>
    <footer>ff</footer>
    <section>gg</section>
    <div>hh</div>
    <span>ii</span>
</body>
</html>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
<!--
    块元素（block element）：在网页中一般通过块元素来对页面进行布局
    行内元素：行内元素主要用来包裹文字
    一般情况下会在块元素内放行内元素，而不会在行内元素中放块元素
    块元素基本上什么都能放
    p元素不能放任何的块元素
    浏览器在解析网页时，会自动对网页中不符合规范的内容进行修正
-->
</body>
</html>
```

### 超链接

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>超链接</title>
    <meta charset="UTF-8">
</head>
<body>
<!-- <br>换行
    使用a标签来定义超链接。超链接也是一个行内元素。
    在超链接中可以嵌套除它自身之外的任何元素
    属性：href指定跳转的目标路径
        值可以是一个外部网站的地址
        也可以是一个内部页面的地址
 -->
    <a href="https://kamier.top">超链接</a>
    <br>
    <a href="实体.html">超链接</a>
</body>
</html>
```

### 列表

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>列表</title>
</head>
<body>
<!--
    列表（list）
        1、铅笔
        2、尺子
    在HTML中也可以创建列表，html列表分为三种
        1、有序列表
        2、无序列表
        3、定义列表
    有序列表，使用ol标签来创建有序列表，使用li表示列表项
    无序列表，使用ul标签创建无序列表，使用li表示列表项
    定义列表，使用dl标签来创建一个定义列表，使用dt来表示定义的内容，使用dd来对内容进行解释说明
-->
<ul>
    <li>结构</li>
    <li>表现</li>
    <li>行为</li>
</ul>
<ol>
    <li>apple</li>
    <li>peach</li>
    <li>pear</li>
</ol>
<dl>
    <dt>结构</dt>
    <dd>结构表示网页网页的结构，结构用来规定网页中哪里是标题，哪里是段落</dd>
</dl>
<ul>
    <li>
        嵌套
        <ul>
            <li>嵌套</li>
        </ul>
    </li>
</ul>
</body>
</html>
```

### 相对路径

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="witth=device-width,initial-scal">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    </head>
<body>
    <!--
     当需要跳转一个服务器内部的页面时，一般都会使用相对路径
    相对路径都会使用.或../开头
    ./表示当前文件所在的目录
    ../表示当前文件所在目录的上一级目录
    target 属性：可用于指定超链接打开的位置
        可选值：
        _self 默认值 在当前页面中打开超链接
        _blank 在一个新的页面中打开超链接
    -->
    <a href="https://kamier.top" target="_self">kamier.top</a>
    <a href="#bottom">去底部</a>
    <a href="#">这是一个占位符</a>
    <a href="javascript:;">这是一个新的超链接，作为href的属性，此时点击什么也不会发生</a>
    <!--
        在开发中#作为超链接的路径的占位符使用
        可以在超链接的href属性设置为#，这样点击超链接以后页面不会发生跳转，而是转到当前页面的顶部的位置
        id属性（唯一不重复），每一个标签都可以添加一个id属性
        id属性就是元素的唯一标识，同一个页面中不能出现重复的id属性
        字母开头，不要数字开头
        可以跳转到任意位置
    -->
    <a id="bottom" href="#">回到顶部</a>
</body>
</html>
```

### 图片标签

```html
<!DOCTYPE html>
<html lang="en">
    <head><meta charset="UTF-8">
        <meta name="viewport" content="with=device-width,initial-">
    </head>
    <body>
        <img src="https://img1.baidu.com/it/u=3631594132,3972221223&fm=26&fmt=auto">
    </body>
    <!--
        图片标签用于从当前页面中引入一个外部图片
        使用img标签引入外部图片，img标签是一个自结束标签
        img输入替换元素
        属性：src 是图片的路径
              alt 是图片的描述，这个描述默认情况下不会显示，有些路径无法加载时显示
              搜索引擎会根据alt中的内容来识别图片，如果不写alt属性则图片不会被搜索引擎所搜索
              width：图片的宽度（单位是像素）
              height 图片的高度，宽高如果只修改了一个，则另一个会等比例缩放
            注意：在pc下，一般不建议修改图片的大小
            （浪费内存，失真），需要多大的图片就裁多大
            移动端，经常需要对图片进行缩放
        图片的格式：
        jpeg（jpg）支持的颜色丰富，不支持透明和动图，一般用来显示照片
        gif 支持的颜色较少，支持简单透明，支持动图，颜色单一，动图
        png 支持的颜色丰富，支持复杂透明，不支持动图，颜色丰富，复杂透明的图片（专为网页而生）
        wepg：谷歌推出的专门用来表示网页中的图片的一种格式，具备其他图片格式的所有优点，文件小
              缺点：兼容性不好    
        base64 将图片使用base64进行编码，这样可以直接将图片转换为字符，通过字符的形式来引入图片
        一般是一些需要和网页一起加载的图片才会使用base64
        效果一样用小的，效果不一样，用效果好的
    -->
</html>
```

### 内联框架 || 音视频嵌入

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="witth=device-width,initial-scal">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
   <!-- 内联框架，用于向当前页面中引入一个其他页面
    frameborder 指定内联框架的边框
-->
   <!--<iframe src="https://kamier.top" width="800" aria-hidden="600" frameborder='0'></iframe>-->
    <!-- audio标签用来向页面中引入一个外部的音频文件的
    音频文件引入时，默认情况下不允许用户自己控制播放停止
    属性：controls是否允许用户控制播放
    autoplay 音频文件是否自动播放，如果设置了autoplay，则音乐在打开页面时会自动播放
    但是目前来讲大部分浏览器都不会自动对音乐进行播放
    loop循环播放 ，这三个属性没有值，用到就写，不用就不写
-->
    <audio controls>
        <!--对不起您的浏览器不支持-->
        <source src="">
        <source src="">
        <embed src="" type="audio/map3">
    </audio>
    <audio src="" controls autoplay></audio>
    <!--  <video src="C:\Users\76261\Desktop\HTML\p13-013_尚硅谷_Linux实操篇_远程登录XShell5[1080P].flv" controls autoplay></audio>-->
    <video controls>
        <source src="">
            <embed src="" type="video/mp4 ">
    </video>
    <embed src="" type="audio/map3">
    <iframe frameborder="0" src="https://v.qq.com/txp/iframe/player.html?vid=f0041qllzll" allowFullScreen="true" width="600" height="600"></iframe>
</body>
</html>
```

