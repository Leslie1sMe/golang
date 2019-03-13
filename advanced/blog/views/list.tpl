{{template "header2.tpl"}}
<section class="row content-wrap">
    <div class="container">
        <div class="row">
            <div class="col-md-8 single-post-contents">
                <article class="single-post-content row m0 post">
                   {{range $val:= .posts}}
                   <header class="row">
                    <h5 class="post-meta">
                        <a href="/list" class="date">{{$val.Created}}</a>
                        <span class="post-author"><i>by</i><a href="/list">{{$val.Author}}</a></span>
                    </h5>
                    <h2 class="post-title">{{$val.Title}}</h2>
                    <div class="row">
                        <h5 class="taxonomy pull-left"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                        <div class="response-count pull-right"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                    </div>
                   </header>
                <div class="featured-content row m0">
                    <a href="/detail/{{$val.Id}}"><img src="static/images/posts/6.jpg" alt=""></a>
                </div>
                   <div class="row m0 tags">
                    <a href="#" class="tag">{{$val.Tags}}</a>
                   </div>
                   {{ end }}

                    <div class="row m0 comments">
                        <!--Comments-->
                        <div class="media comment">
                            <div class="media-left">
                            </div>
                            <div class="media-body">
                                <!--Comments-->
                                <div class="media comment reply">
                                    <div class="media-left">
                                    </div>
                                    <div class="media-body">
                                        <!--Comments-->
                                        <div class="media comment reply">
                                            <div class="media-left">
                                            </div>
                                            <div class="media-body">
                                                <!--Comments-->
                                                <div class="media comment reply">
                                                    <div class="media-left">
                                                    </div>
                                                    <div class="media-body">

                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>

                        <!--Comments-->
                        <div class="media comment">
                            <div class="media-left">
                            </div>
                            <div class="media-body">
                            </div>
                        </div>
                    </div>
                    </div>
                </article>
            </div>
            <div class="col-md-4 sidebar">
                <!--Author Widget-->
                <aside class="row m0 widget-author widget">
                    <div class="widget-author-inner row">
                        <div class="author-avatar row"><a href="#"><img src="static/images/author.png" alt="" class="img-circle"></a></div>
                        <a href="#"><h2 class="author-name">Leslie</h2></a>
                        <h5 class="author-title">Backend Engineer</h5>
                        <p>The word nature is derived from the Latin word natura, or "essential qualities, innate disposition", and in ancient times, literally meant "birth".</p>
                    </div>
                </aside>
                <aside class="row m0 widget widget-twitter">
                    <div class="widget-twitter-inner">
                        <div class="row tweet-texts">
                            <p>最新文章请关注我的微信公众号JosephCircle <a href="https://leslie.net.cn"></a> <a href="https://leslie.net.cn">https://leslie.net.cn</a></p>
                        </div>
                        <div class="row timepast">1 day ago</div>
                    </div>
                </aside>
                <aside class="row m0 widget widget-instagram">
                    <div class="widget-twitter-inner">
                     <div class="row tweet-texts">
                        <p>Github地址: <a href="https://github.com/Leslie1sMe"></a> <a href="https://github.com/Leslie1sMe">https://github.com/Leslie1sMe</a></p>
                    </div>
                        <div id="instafeed"></div>
                    </div>
                </aside>
            </div>
        </div>
    </div>
</section>

{{template "footer.tpl"}}
