{{template "header2.tpl"}}
<section class="row content-wrap single-nosidebar">
    <div class="container">
        <div class="row">
            <div class="col-sm-12 single-post-contents">
                <article class="single-post-content v2 v2p1 row m0 post">
                    <header class="row">

                        <div class="row">
                            <h5 class="taxonomy pull-left"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count pull-right"><img src="static/images/comment-icon-gray.png" alt="">100+</div>
                        </div>
                        <h2 class="post-title">{{.title}}</h2>
                        <h5 class="post-meta">
                            By <a href="#">{{ .author }}</a> | <a href="#" class="date">{{.created}}</a>
                        </h5>
                    </header>
                    <div class="featured-content row m0">
                        <a href="#"><img src="static/images/featured-posts/3.jpg" alt=""></a>
                    </div>
                    <div class="post-content row">
                        <h3>{{.info}}</h3>
                                   <br>
                                   <br>
                                   <br>
                        <div class="row">
                            <div class="col-sm-10">
                                <p>{{ .content }}</p>
                            </div>
                        </div>

                        <br>
                    </div>
                    <div class="row m0 tags">
                        <a href="#" class="tag">{{.tags}}</a>
                    </div>

                    <ul class="pager">
                    {{if .is_first}}
                        <li>
                           <h4>Previous Articles</h4>
                           <h2 class="post-title"><a href="#">没有啦，已经到最开始啦</a></h2>
                       </li>
                        {{else}}
                         <li>
                        <h4>Previous Articles</h4>
                        <h2 class="post-title"><a href="/detail/{{.front.Id}}">{{.front.Title}}</a></h2>
                        </li>
                        {{end}}
                        <li>
                            <h4>Next Articles</h4>
                            <h2 class="post-title"><a href="/detail/{{.backend.Id}}">{{.backend.Title}}</a></h2>
                        </li>
                    </ul>
                    </div>
                </article>
            </div>
        </div>
    </div>
</section>

<!--Footer-->
{{template "footer.tpl"}}

