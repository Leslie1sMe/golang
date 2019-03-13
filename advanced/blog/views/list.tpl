<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<!--========== The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags ==========-->
<title>Chivalric</title>

<!--==========Dependency============-->
<link rel="stylesheet" href="static/css/bootstrap.min.css">
<link rel="stylesheet" href="static/css/font-awesome.min.css">
<link rel="stylesheet" href="static/vendors/owl-carousel/assets/owl.carousel.css">
<link rel="stylesheet" href="static/vendors/magnific-popup/magnific-popup.css">
<link rel="stylesheet" href="static/vendors/flexslider/flexslider.css">

<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Kanit:500">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Josefin+Sans:600,700italic">
<link href='https://fonts.googleapis.com/css?family=Dosis:400,200,300,500,600,800,700' rel='stylesheet' type='text/css'>
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Muli:400,300,300italic,400italic">
<link href='https://fonts.googleapis.com/css?family=Roboto:400,100,100italic,300,300italic,400italic,500italic,500,700italic,700,900,900italic' rel='stylesheet' type='text/css'>
<link href='https://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>
<link href='https://fonts.googleapis.com/css?family=Fredoka+One' rel='stylesheet' type='text/css'>

<!--==========Theme Styles==========-->
<link href="static/css/style.css" rel="stylesheet">
<link href="static/css/theme/green.css" rel="stylesheet">

<!--========== HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries ==========-->
<!--========== WARNING: Respond.js doesn't work if you view the page via file:// ==========-->
<!--==========[if lt IE 9]>
    <script src="js/html5shiv.min.js"></script>
    <script src="js/respond.min.js"></script>
<![endif]==========-->
</head>
<body class="home">
{{template "header2.tpl"}}
<section class="row content-wrap">
    <div class="container">
        <div class="row">
            <div class="col-md-8 single-post-contents">
                <article class="single-post-content row m0 post">
                   {{range $val:= .posts}}
                   <header class="row">
                    <h5 class="post-meta">
                        <a href="#" class="date">{{$val.Created}}</a>
                        <span class="post-author"><i>by</i><a href="#">{{$val.Author}}</a></span>
                    </h5>
                    <h2 class="post-title">{{$val.Title}}</h2>
                    <div class="row">
                        <h5 class="taxonomy pull-left"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                        <div class="response-count pull-right"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                    </div>
                   </header>
                <div class="featured-content row m0">
                    <a href="#"><img src="static/images/posts/6.jpg" alt=""></a>
                </div>
                   <div class="row m0 tags">
                    <a href="#" class="tag">{{$val.Tags}}</a>
                   </div>
                   {{ end }}
                    <ul class="pager">
                        <li>
                            <h4>Previous Articles</h4>
                            <h2 class="post-title"><a href="single2.html">Nature, in the broadest sense, is the natural...</a></h2>
                            <h5 class="taxonomy pull-left"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                        </li>
                        <li>
                            <h4>Next Articles</h4>
                            <h2 class="post-title"><a href="single2.html">Nature, in the broadest sense, is the natural...</a></h2>
                            <h5 class="taxonomy pull-left"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                        </li>
                    </ul>

                    <div class="row m0 comments">
                        <h5 class="response-count">5 comments<a href="#comment-form" class="btn btn-primary pull-right"><span>add comment</span></a></h5>
                        <!--Comments-->
                        <div class="media comment">
                            <div class="media-left">
                                <a href="#"><img src="static/images/posts/c1.jpg" alt="" class="img-circle"></a>
                            </div>
                            <div class="media-body">
                                <h4><a href="#">Mark Sanders</a></h4>
                                <h5><a href="#" class="date">feb 17, 2016 at 3:30pm</a> | <a href="#" class="reply-link">reply</a></h5>
                                <p>It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.</p>
                                <!--Comments-->
                                <div class="media comment reply">
                                    <div class="media-left">
                                        <a href="#"><img src="static/images/posts/c2.jpg" alt="" class="img-circle"></a>
                                    </div>
                                    <div class="media-body">
                                        <h4><a href="#">Mark Sanders</a></h4>
                                        <h5><a href="#" class="date">feb 17, 2016 at 3:30pm</a> | <a href="#" class="reply-link">reply</a></h5>
                                        <p>It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.</p>
                                        <!--Comments-->
                                        <div class="media comment reply">
                                            <div class="media-left">
                                                <a href="#"><img src="static/images/posts/c1.jpg" alt="" class="img-circle"></a>
                                            </div>
                                            <div class="media-body">
                                                <h4><a href="#">Mark Sanders</a></h4>
                                                <h5><a href="#" class="date">feb 17, 2016 at 3:30pm</a> | <a href="#" class="reply-link">reply</a></h5>
                                                <p>Remaining essentially unchanged.</p>
                                                <!--Comments-->
                                                <div class="media comment reply">
                                                    <div class="media-left">
                                                        <a href="#"><img src="static/images/posts/c2.jpg" alt="" class="img-circle"></a>
                                                    </div>
                                                    <div class="media-body">
                                                        <h4><a href="#">Mark Sanders</a></h4>
                                                        <h5><a href="#" class="date">feb 17, 2016 at 3:30pm</a> | <a href="#" class="reply-link">reply</a></h5>
                                                        <p>Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit</p>
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
                                <a href="#"><img src="static/images/posts/c1.jpg" alt="" class="img-circle"></a>
                            </div>
                            <div class="media-body">
                                <h4><a href="#">Mark Sanders</a></h4>
                                <h5><a href="#" class="date">feb 17, 2016 at 3:30pm</a> | <a href="#" class="reply-link">reply</a></h5>
                                <p>It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.</p>
                            </div>
                        </div>
                    </div>
                    </div>
                    <form action="#" method="post" class="comment-form row" id="comment-form">
                        <h5 class="form-title">leave a reply</h5>
                        <div class="form-group col-sm-6">
                            <label for="name">full name*</label>
                            <input type="text" id="name" class="form-control" placeholder="Your name" required>
                        </div>
                        <div class="form-group col-sm-6">
                            <label for="email">full name*</label>
                            <input type="email" id="email" class="form-control" placeholder="Your email address here" required>
                        </div>
                        <div class="form-group col-sm-6">
                            <label for="website">website</label>
                            <input type="url" id="website" class="form-control" placeholder="Your website url" >
                        </div>
                        <div class="form-group col-sm-6">
                            <label for="subject">subject</label>
                            <input type="text" id="subject" class="form-control" placeholder="Write subject here" required>
                        </div>
                        <div class="form-group col-sm-12">
                            <label for="message">message</label>
                            <textarea name="message" id="message" class="form-control" placeholder="Write message here"></textarea>
                        </div>
                        <div class="col-sm-12">
                            <button type="submit" class="btn-primary"><span>send</span></button>
                            <h5 class="pull-right">*required fields</h5>
                        </div>
                    </form>
                </article>
            </div>
            <div class="col-md-4 sidebar">
                <!--Author Widget-->
                <aside class="row m0 widget-author widget">
                    <div class="widget-author-inner row">
                        <div class="author-avatar row"><a href="#"><img src="static/images/author.png" alt="" class="img-circle"></a></div>
                        <a href="#"><h2 class="author-name">Mark Sanders</h2></a>
                        <h5 class="author-title">small title</h5>
                        <p>The word nature is derived from the Latin word natura, or "essential qualities, innate disposition", and in ancient times, literally meant "birth".</p>
                        <a href="#" class="know-more">know more</a>
                    </div>
                    <ul class="nav social-nav">
                        <li><a href="#"><i class="fa fa-twitter"></i></a></li>
                        <li><a href="#"><i class="fa fa-facebook-official"></i></a></li>
                        <li><a href="#"><i class="fa fa-google-plus"></i></a></li>
                        <li><a href="#"><i class="fa fa-instagram"></i></a></li>
                        <li><a href="#"><i class="fa fa-envelope"></i></a></li>
                    </ul>
                </aside>
                <!--Twitter Widget-->
                <aside class="row m0 widget widget-twitter">
                    <div class="widget-twitter-inner">
                        <h5 class="widget-meta"><i class="fa fa-twitter"></i>Twitter feed <a href="http://twitter.com/chivalricblog">@chivalricblog</a></h5>
                        <div class="row tweet-texts">
                            <p>Check out new post on my blog <a href="http://twitter.com/#natureshot">#natureshot</a> <a href="http://bit.ly/blog">http://bit.ly/blog</a></p>
                        </div>
                        <div class="row timepast">1 day ago</div>
                    </div>
                </aside>
                <!--Instagram Widget-->
                <aside class="row m0 widget widget-instagram">
                    <div class="widget-instagram-inner">
                        <h5 class="widget-meta"><i class="fa fa-twitter"></i>instagram feed <a href="http://twitter.com/chivalricblog">@chivalricblog</a></h5>
                        <div id="instafeed"></div>
                    </div>
                </aside>
            </div>
        </div>
    </div>
</section>

{{template "footer.tpl"}}
<!--========== jQuery (necessary for Bootstrap's JavaScript plugins) ==========-->
<script src="static/js/jquery-2.2.0.min.js"></script>
<script src="static/js/bootstrap.min.js"></script>
<script src="static/vendors/owl-carousel/owl.carousel.min.js"></script>
<script src="static/vendors/magnific-popup/jquery.magnific-popup.min.js"></script>
<script src="static/vendors/instafeed/instafeed.min.js"></script>
<script src="static/vendors/imagesLoaded/imagesloaded.pkgd.min.js"></script>
<script src="static/vendors/isotope/isotope.pkgd.min.js"></script>
<script src="static/vendors/flexslider/jquery.flexslider-min.js"></script>
<script src="static/js/theme.js"></script>
</body>
</html>
