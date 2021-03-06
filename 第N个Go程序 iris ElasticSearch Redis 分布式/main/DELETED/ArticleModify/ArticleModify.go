package articleModify

import (
	"github.com/shiyanhui/hero"
	"io"
	"../../Entity"
	"strconv"
)
func ArticleToWriter(menulist []Entity.Menu,article Entity.Article,w io.Writer) (int, error){
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
	<div class="row" style="margin-top:5%;">
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
		<div class="col-md-3 col-lg-3 col-sm-4 col-xs-4">
				<ul class="nav nav-pills nav-stacked">
				  <li role="presentation"><a href="/backend/1">修改用户</a></li>
				  <li role="presentation" class="active"><a href="/articlemodifylist/1">修改文章</a></li>
				  <li role="presentation"><a href="/articleinsert">增加文章</a></li>
				  <li role="presentation"><a href="/upload">上传文件</a></li>
				</ul>
		</div>
		<div class="col-md-5 col-lg-5 col-sm-6 col-xs-6">
			<form method="post" name="form" id="form" action="/articleModify">
        		<div class="form-group">
            		<label for="Title">标题</label>
            		<input type="text" class="form-control" id="Title" name="Title" value="`)
	_buffer.WriteString(article.Title)
	_buffer.WriteString(`"></input>
					<label for="Classify">分类</label>
            		<input type="text" class="form-control" id="Classify" name="Classify" value="`)
	_buffer.WriteString(article.Classify)
	_buffer.WriteString(`"></input>
        		</div>
        		<select name="Menu" id="Menu" class="form-control">`)
	menu,_ := strconv.Atoi(article.Menu)
	for i,v := range menulist {
		_buffer.WriteString(`<option value="`+strconv.Itoa(i+1)+`"`)
		if menu == i+1 {_buffer.WriteString(` selected = "selected"`)}
		_buffer.WriteString(`>`+v.Name+`</option>`)
	}
	_buffer.WriteString(`</select>
        		<textarea class="form-control" rows="20" name="Content" id="Content">`)
	_buffer.WriteString(string(article.Content))
	_buffer.WriteString(`
        		</textarea>
				<input type="text" id ="Id" name="Id" style="display:none" value="`)
	_buffer.WriteString(strconv.FormatInt(article.Id,10))
	_buffer.WriteString(`"></input>
			<button type="submit" class="btn btn-default" name="btn" id="btn" onclick="return validate()">确认</button>
    		</form>
		</div>
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
	</div>
	<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
	<script src="../js/Modify.js"></script>
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}