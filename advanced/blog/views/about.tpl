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

<section class="row page-content-wrap">
    <div class="container">
        <h2 class="page-title text-center">about me</h2>
        <div class="row">
            <div class="col-md-8 page-contents">
                <div class="row page-content">
                    <img src="static/images/author.jpg" alt="" class="img-responsive">

                    <div class="contents row">
                        <h4>Mark Samder, lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</h4>

                        <p>It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                        <br>
                        <div class="row">
                            <div class="col-sm-6">
                                <p>It has survived not only five centuries, remaining essentially unchanged. <a href="www.mywebsite.com">www.mywebsite.com</a> it was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                            </div>
                            <div class="col-sm-6">
                                <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.</p>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
            <div class="col-md-4 sidebar">
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
                <aside class="row m0 widget widget-facebook">
                    <div class="widget-inner widget-facebook-inner">
                        <a href="#" class="facebook-link"><i class="fa fa-facebook-square"></i>find us on facebook</a>
                    </div>
                    <ul class="nav social-nav">
                        <li><a href="#"><i class="fa fa-twitter"></i></a></li>
                        <li><a href="#"><i class="fa fa-facebook-official"></i></a></li>
                        <li><a href="#"><i class="fa fa-google-plus"></i></a></li>
                        <li><a href="#"><i class="fa fa-instagram"></i></a></li>
                        <li><a href="#"><i class="fa fa-envelope"></i></a></li>
                    </ul>
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
<script src="static/js/theme.js"></script>
</body>
</html>
