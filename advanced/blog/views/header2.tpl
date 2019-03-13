<header class="row transparent white" data-spy="affix" data-offset-top="300" id="header">
    <div class="container">
        <div class="row top-header">
            <div class="col-sm-4 search-form-col">
                <form action="#" method="get" class="search-form">
                    <div class="input-group">
                        <span class="input-group-addon"><img src="static/images/search-icon-dark.png" alt=""></span>
                        <input type="search" class="form-control" placeholder="search">
                    </div>
                </form>
            </div>
            <div class="col-sm-4 logo-col text-center">
            </div>
            <div class="col-sm-4 menu-trigger-col">
                <button class="menu-trigger black pull-right">
                    <span class="active-page">navigation</span>
                    <img src="static/images/menu-align-dark.png" alt="" class="icon-burger">
                    <img src="static/images/menu-close-dark.png" alt="" class="icon-cross">
                </button>
            </div>
        </div>
    </div>
    <div class="row menu-section">
        <div class="container">
            <div class="row">
                <div class="col-sm-8 menu-col">
                    <div class="row">
                        <ul class="nav column-menu white-bg">
                            <li><a href={{urlfor "MainController.Get"}}>HOME</a></li>
                            <li><a href={{urlfor "MainController.About"}}>About</a></li>
                            <li><a href={{urlfor "ArticleController.List"}}>ARTICLEs</a></li>
                        </ul>
                        <ul class="nav column-menu white-bg">
                            <li><a href={{urlfor "MainController.Contact"}}>contact</a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</header>
