#### 盒子模型

我们可以将div看成一个盒子
1. content: 表示内容，内容的宽用width表示，高用height表示
2. padding:表示盒子到边框之间的距离
3. border：表示边框的厚度
4. margin: 表示盒子与盒子之间的距离

看这样一段代码
```html
<!DOCTYPE html>

<html>
<head>
    <style type = "text/css">
    #foo {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
    }
    #foo1 {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    </style>

</head>

<body>
    <div id = "foo">这是盒子的内容</div>
    <div id = "foo1">这是另外一个盒子</div>
</body>
</html>
```

![tWGY4.png](https://s1.328888.xyz/2022/05/28/tWGY4.png)



当我们将这段代码放到浏览器中会发现还有一点距离,这是因为每个浏览器在渲染的时候会给这些margin,padding,border预设一些值，
我们需要将这些值给消除，所以我们可以使用`*`,这种通配选择器来选中所有,但是这种选择器的效率较低
```html
<!DOCTYPE html>

<html>
<head>
    <style type = "text/css">
    *{
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    #foo1 {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    </style>

</head>

<body>
    <div id = "foo">这是盒子的内容</div>
    <div id = "foo1">这是另外一个盒子</div>
</body>
</html>
```
![tWQdk.png](https://s1.328888.xyz/2022/05/28/tWQdk.png)

#### margin的使用
将上面的`foo`中的`margin`改为`margin:20px`得到如下效果
```html
<!DOCTYPE html>

<html>
<head>
    <style type = "text/css">
    *{
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        margin: 20px;
    }
    #foo1 {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    </style>

</head>

<body>
    <div id = "foo">这是盒子的内容</div>
    <div id = "foo1">这是另外一个盒子</div>
</body>
</html>
```
![t5MRv.png](https://s1.328888.xyz/2022/05/28/t5MRv.png)

这是发现上下左右都有一点空隙了，如果我们想左边的空隙和上边的空隙不一样,那么我们如何操作呢？
```html
<!DOCTYPE html>

<html>
<head>
    <style type = "text/css">
    *{
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        margin: 10px 20px 10px 20px;
    }
    #foo1 {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    </style>

</head>

<body>
    <div id = "foo">这是盒子的内容</div>
    <div id = "foo1">这是另外一个盒子</div>
</body>
</html>
```
![tW2KZ.png](https://s1.328888.xyz/2022/05/28/tW2KZ.png)

由此可以看出margin的四个值分别是上,右,下,左

如果想单独设值某一方向上的`margin`值的时候可以使用`margin-top`,`margin-right`,`margin-botom`,`margin-left`

使用`margin-bottom`和`margin-bottom`的时候要注意`margin`的外边距有重叠效果，就是在竖直方向上当有两个元素外边距连在一起的时候，此时两者的距离不是两者之和而是取最大值
这是演示的例子
```html
<!DOCTYPE html>

<html>
<head>
    <style type = "text/css">
    *{
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        margin-bottom: 10px;
    }
    #foo1 {
        margin-top: 20px;
        width: 100px;
        height: 100px;
        background: blue;
        padding: 10px;
        border: solid red 10px;
        
    }
    </style>

</head>

<body>
    <div id = "foo">这是盒子的内容</div>
    <div id = "foo1">这是另外一个盒子</div>
</body>
</html>
```
![tWTFF.png](https://s1.328888.xyz/2022/05/28/tWTFF.png)

#### pading的使用
1. padding的颜色就是盒子的背景色
2. padding也有上，右，下，左四个属性

我们设置了一个父`div`，然后又设置了一个子`div`,让子`div`铺满父`div`,然后使用`padding`扩展父`div`

```html
<!DOCTYPE HTML>

<html>
<head>

    <style type = "text/css">
    * {
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 300px;
        height: 300px;
        background: gold;
        padding: 50px;
    }
    #son {
        width: 100%;
        height: 100%;
        background: black;
    }
    </style>
</head>
<body>
    <div id = "foo">
        <div id = "son"></div>
    </div>
</body>

</html>
```
![t7gMC.png](https://s1.328888.xyz/2022/05/28/t7gMC.png)


上,右,下，左分别设置不同的值
```html
<!DOCTYPE HTML>

<html>
<head>

    <style type = "text/css">
    * {
        margin: 0;
        padding: 0;
        border: 0;
    }
    #foo {
        width: 300px;
        height: 300px;
        background: gold;
        padding: 10px 20px 10px 20px;
    }
    #son {
        width: 100%;
        height: 100%;
        background: black;
    }
    </style>
</head>
<body>
    <div id = "foo">
        <div id = "son"></div>
    </div>
</body>

</html>
```
![tsxyT.png](https://s1.328888.xyz/2022/05/28/tsxyT.png)

此时的背景是黑色，但是我们设置了`padding`
#### border的使用
1. border表示边框
