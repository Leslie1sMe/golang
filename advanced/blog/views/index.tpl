<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<!--========== The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags ==========-->
<title>Leslie`s At Home</title>

<!--==========Dependency============-->
<link rel="stylesheet" href="static/css/bootstrap.min.css">
<link rel="stylesheet" href="static/css/font-awesome.min.css">
<link rel="stylesheet" href="static/vendors/owl-carousel/assets/owl.carousel.css">
<link rel="stylesheet" href="static/vendors/magnific-popup/magnific-popup.css">

<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Kanit:500">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Josefin+Sans:600,700italic">
<link href='https://fonts.googleapis.com/css?family=Dosis:400,200,300,500,600,800,700' rel='stylesheet' type='text/css'>
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Muli:400,300,300italic,400italic">
<link href='https://fonts.googleapis.com/css?family=Roboto:400,100,100italic,300,300italic,400italic,500italic,500,700italic,700,900,900italic' rel='stylesheet' type='text/css'>

<!--==========Theme Styles==========-->
<link href="static/css/style.css" rel="stylesheet">
<link href="static/css/theme/green.css" rel="stylesheet">

<!--========== HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries ==========-->
<!--========== WARNING: Respond.js doesn't work if you view the page via file:// ==========-->
<!--==========[if lt IE 9]>
    <script src="static/js/html5shiv.min.js"></script>
    <script src="static/js/respond.min.js"></script>
<![endif]==========-->
</head>
<body class="home">
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
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-image">
                <div class="post-wrapper row">
                    <div class="featured-content row">
                        <a href="single.html"><img src="static/images/posts/masonry/1.jpg" alt="" class="img-responsive"></a>
                    </div>
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">feb 17, 2016</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">测试4</a></h3>
                        <p>概括</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-image">
                <div class="post-wrapper row">
                    <div class="featured-content row">
                        <a href="single.html"><img src="static/images/posts/masonry/2.jpg" alt="" class="img-responsive"></a>
                    </div>
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">feb 17, 2016</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">测试5</a></h3>
                        <p>概括</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
            <!--Author Widget-->
            <aside class="col-sm-4 widget-author widget widget-with-posts post">
                <div class="widget-author-inner row">
                    <div class="author-avatar row"><a href="#"><img src="static/images/author.png" alt="" class="img-circle"></a></div>
                    <a href="#"><h2 class="author-name">Leslie</h2></a>
                    <h5 class="author-title">Backend Egineer</h5>
                    <p>Some people hear their own inner voices with great clearness. And they live by what they hear. Such people become crazy, or they become legends</p>
                </div>
            </aside>
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-gallery">
                <div class="post-wrapper row">
                    <div class="featured-content row">
                        <div class="gallery-of-post">
                            <div class="item"><img src="static/images/posts/masonry/3.jpg" alt=""></div>
                            <div class="item"><img src="static/images/posts/masonry/3.jpg" alt=""></div>
                            <div class="item"><img src="static/images/posts/masonry/3.jpg" alt=""></div>
                            <div class="item"><img src="static/images/posts/masonry/3.jpg" alt=""></div>
                            <div class="item"><img src="static/images/posts/masonry/3.jpg" alt=""></div>
                        </div>
                    </div>
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">feb 17, 2016</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">测试6.</a></h3>
                        <p>概括</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-text">
                <div class="post-wrapper row">
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">feb 17, 2016</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">测试.</a></h3>
                        <p>概括</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-video">
                <div class="post-wrapper row">
                    <div class="featured-content row">
                        <a href="single.html">
                            <img src="static/images/posts/masonry/4.jpg" alt="" class="img-responsive">
                            <img src="static/images/play-btn.png" alt="" class="video-mark">
                        </a>
                    </div>
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">feb 17, 2016</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title"><a href="single.html">测试.</a></h3>
                        <p>概括</p>
                        <footer class="row">
                            <h5 class="taxonomy"><i>in</i> <a href="#">image</a>, <a href="#">entertainment</a></h5>
                            <div class="response-count"><img src="static/images/comment-icon-gray.png" alt="">5</div>
                        </footer>
                    </div>
                </div>
            </article>
            <!--Twitter Widget-->
            <!--Blog Post-->
            <article class="col-sm-4 post post-masonry post-format-quote">
                <div class="post-wrapper row">
                    <div class="post-excerpt row">
                        <h5 class="post-meta">
                            <a href="#" class="date">thr 13, 2019</a>
                            <span class="post-author"><i>by</i><a href="#">Leslie</a></span>
                        </h5>
                        <h3 class="post-title">If everybody learns this simple art of loving his work, whatever it is, enjoying it without asking for any recognition, we would have more beautiful and celebrating world.</h3>
                        <footer class="row">

                        </footer>
                    </div>
                </div>
            </article>
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
<script src="static/js/theme.js"></script>
</body>
</html>
