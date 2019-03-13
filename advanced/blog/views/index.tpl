{{template "header2.tpl"}}
<section class="row featured-post-carousel">
    <div class="item post">
        <img src="static/images/featured-posts/1.jpg" alt="" class="img-responsive main-bg">
        <div class="post-content">
            <div class="container">
                <h5 class="post-meta"><i>in</i> <a href="#">Author Profile</a> | <a href="#">feb 17, 2016</a></h5>
                <h2 class="post-title"><a href="single.html">测试1</a></h2>
                <a href="single.html" class="btn btn-primary"><span>read more</span></a>
            </div>
        </div>
    </div>
    <div class="item post">
        <img src="static/images/featured-posts/1-1.jpg" alt="" class="img-responsive main-bg">
        <div class="post-content">
            <div class="container">
                <h5 class="post-meta"><i>in</i> <a href="#">Author Profile</a> | <a href="#">feb 17, 2016</a></h5>
                <h2 class="post-title"><a href="single.html">测试2</a></h2>
                <a href="single.html" class="btn btn-primary"><span>read more</span></a>
            </div>
        </div>
    </div>
    <div class="item post">
        <img src="static/images/featured-posts/1-2.jpg" alt="" class="img-responsive main-bg">
        <div class="post-content">
            <div class="container">
                <h5 class="post-meta"><i>in</i> <a href="#">Author Profile</a> | <a href="#">feb 17, 2016</a></h5>
                <h2 class="post-title"><a href="single.html">测试3</a></h2>
                <a href="single.html" class="btn btn-primary"><span>read more</span></a>
            </div>
        </div>
    </div>
</section>
<section class="row content-wrap">
    <div class="container">
        <div class="row" id="post-masonry">
            {{range $val:=.articles}}
            <article class="col-sm-4 post post-masonry post-format-image">
                <div class="post-wrapper row">
                    <div class="featured-content row">
                        <a href="/detail/{{$val.Id}}"><img src="static/images/posts/masonry/2.jpg" alt="" class="img-responsive"></a>
                    </div>
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">{{$val.Created}}</a>
                            <span class="post-author"><i>by</i><a href="#">{{$val.Author}}</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">{{$val.Title}}</a></h3>
                        <p>{{$val.Info}}</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
        {{ end }}
        </div>
    </div>
</section>

{{template "footer.tpl"}}
