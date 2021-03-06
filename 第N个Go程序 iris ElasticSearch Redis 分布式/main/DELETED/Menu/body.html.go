
// Code generated by hero.
// source: C:\Users\hduoct\Desktop\GoWeb\Menu\body.html
// DO NOT EDIT!
package Menu
import (
	"github.com/shiyanhui/hero"
	"io"
	"../../Entity"
	"strconv"
	"github.com/kataras/iris"
)
func MenuWriter(id int64,m map[string]int64,entity Entity.Entity,page int,alen int,ctx iris.Context,w io.Writer) (int, error){
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>index</title>
    <!-- Bootstrap -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style type="text/css">
        body{ font-family: Microsoft YaHei,'宋体' , Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;}
    </style>
    <script src="https://cdn.bootcss.com/markdown.js/0.6.0-beta1/markdown.min.js"></script>
</head>
<body>
`)
	_buffer.WriteString(`
<nav id="navbar-example" class="navbar navbar-default navbar-static" role="navigation">
    <div class="container-fluid">
        <div class="navbar-header">
            <button class="navbar-toggle" type="button" data-toggle="collapse"
                    data-target=".bs-js-navbar-scrollspy">
                <span class="sr-only">切换导航</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/">Web开发学习笔记</a>
        </div>
        <div class="collapse navbar-collapse bs-js-navbar-scrollspy">
            <ul class="nav navbar-nav">`)
	for i,m := range entity.MenuList {
		_buffer.WriteString(`<li><a href="/menu/`+strconv.Itoa(i+1)+`/1">`+m.Name+`</a></li>`)
	}
	_buffer.WriteString(`
            </ul>
            <form class="navbar-form navbar-left">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Search" value="golang是最好的语言">
                </div>
                <button type="submit" class="btn btn-default">Submit</button>
            </form>
            <ul class="nav navbar-nav navbar-right">
				<li><a href="/adminlogin">后台</a></li>`)
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		_buffer.WriteString(`<li><a href="/register">注册</a></li>`+`<li><a href="/login">登录</a></li>`)
	}else {
		_buffer.WriteString(`<li><a href="/user">欢迎你&nbsp;:&nbsp;`+Entity.Sess.Start(ctx).GetString("Username")+`&nbsp;&nbsp;&nbsp;更多</a></li>`)
	}
	_buffer.WriteString(`
            </ul>
        </div>
    </div>
</nav>
<div class="cow">
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
    <div class="col-md-5 col-lg-5 col-sm-10 col-xs-10">
        <div>
			<ol class="breadcrumb">
			  <li><a href="/">主页</a></li>
			  <li class="active">`)
	_buffer.WriteString(entity.Menu.Name)
	_buffer.WriteString(`</li>
			</ol>
            <h3>分类:`)
	hero.EscapeHTML(entity.Menu.Name,_buffer)
	_buffer.WriteString(`(共 `)
	_buffer.WriteString(strconv.Itoa(len(entity.ArticleList)))
	_buffer.WriteString(` 篇文章)</h3><br/><div class="list-group">`)
	for _, menu := range entity.ArticleList {
		_buffer.WriteString(`<a href="/article/`)
		hero.EscapeHTML(strconv.FormatInt(menu.Id,10),_buffer)
		_buffer.WriteString(`" class="list-group-item">
`)
		hero.EscapeHTML(menu.Title, _buffer)
		_buffer.WriteString(`
</a>`)
	}
	_buffer.WriteString(`
            </div>
        </div>`)
	var pages [5]string
	for k := range pages {
		if alen >5 {
			pages[k] = strconv.Itoa(page+k)
		} else {
			pages[k] = strconv.Itoa(k+1)
		}
	}
	var class [5]string
	for k := range class {
		if page == k+1 || page >5 {
			class[k] = "active"
		}
		if (k + 1) > alen {
			class[k] = "disabled"
		}
	}
	_buffer.WriteString(`
			<nav aria-label="...">
			  <ul class="pagination">
				<li class=`)
	if page == 1 {_buffer.WriteString("disabled")}
	_buffer.WriteString(`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+strconv.Itoa(page-1)+`" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>
				<li class=`+class[0]+`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+pages[0]+`">1 <span class="sr-only">(current)</span></a></li>
				<li class=`+class[1]+`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+pages[1]+`">2 <span class="sr-only">(current)</span></a></li>
				<li class=`+class[2]+`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+pages[2]+`">3 <span class="sr-only">(current)</span></a></li>
				<li class=`+class[3]+`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+pages[3]+`">4 <span class="sr-only">(current)</span></a></li>
				<li class=`+class[4]+`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+pages[4]+`">5 <span class="sr-only">(current)</span></a></li>
				<li class=`)
	if page == alen {_buffer.WriteString("disabled")}
	_buffer.WriteString(`><a href="/menu/`+strconv.FormatInt(id,10)+`/`+strconv.Itoa(page+1)+`" aria-label="Next"><span aria-hidden="true">&raquo;</span></a></li>
			  </ul>
			</nav>
    </div>
	<div class="col-md-3 col-lg-3 hidden-sm hidden-xs">
		<h3>分类</h3>
		<ul class="list-group">`)
		for k,v := range m {
		  _buffer.WriteString(`<li class="list-group-item">
			<span class="badge">`+strconv.FormatInt(v,10)+`
			</span>
			<a href="/Classify/`+k+`/1">`+k+`</a>
		  </li>`)
		}
		_buffer.WriteString(`</ul>
	</div>
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
</div>
`)

	_buffer.WriteString(`
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}