
<header class="row transparent black header1" data-spy="affix" data-offset-top="0" id="header">
    <div class="container">
        <div class="row top-header">
            <div class="col-sm-4 search-form-col">
                <form action="#" method="get" class="search-form">
                    <div class="input-group">
                        <span class="input-group-addon"><img src="static/images/search-icon-white.png" alt=""></span>
                        <input type="search" class="form-control" placeholder="search">
                    </div>
                </form>
            </div>
            <div class="col-sm-4 logo-col text-center">
            </div>
            <div class="col-sm-4 menu-trigger-col">
                <button class="menu-trigger pull-right">
                    <span class="active-page">home</span>
                    <img src="static/images/menu-align-white.png" alt="" class="icon-burger">
                    <img src="static/images/menu-close-white.png" alt="" class="icon-cross">
                </button>
            </div>
        </div>
    </div>
    <div class="row menu-section">
        <div class="container">
            <div class="row">
                <div class="col-sm-8 menu-col">
                    <div class="row">
                        <ul class="nav column-menu black-bg">
                            <li><a href="{{urlfor  "MainController.Get" }}">HOME</a></li>
                            <li><a href="{{urlfor  "MainController.About" }}">ABOUT</a></li>
                            <li><a href="{{urlfor "ArticleController.List" }}">LIST</a></li>

                        </ul>
                        <ul class="nav column-menu black-bg">
                            <li><a href="{{urlfor  "MainController.Contact" }}">Contact</a></li>
                            {{ range $val:= .m }}
                            <li><a href="{{urlfor  "MainController.Contact" }}">{{.m.Name}}</a></li>
                            {{ end }}
                        </ul>
                    </div>
                </div>
                <div class="col-sm-4 subscribe-col">
                    <h5 class="widget-title">subscribe to our newsletter.</h5>
                    <form action="#" method="post" class="form-inline subscribe-form">
                        <div class="form-group">
                            <input type="email" class="form-control" placeholder="Email">
                        </div>
                        <button type="submit" class="btn btn-primary btn-sm"><span>send</span></button>
                    </form>
                    <ul class="nav social-nav dark">
                        <li><a href="#"><i class="fa fa-twitter"></i></a></li>
                        <li><a href="#"><i class="fa fa-facebook-official"></i></a></li>
                        <li><a href="#"><i class="fa fa-google-plus"></i></a></li>
                        <li><a href="#"><i class="fa fa-instagram"></i></a></li>
                        <li><a href="#"><i class="fa fa-envelope"></i></a></li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</header>
