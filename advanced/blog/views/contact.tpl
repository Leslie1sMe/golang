{{template "header2.tpl"}}

<section class="row page-content-wrap">
    <div class="container">
        <h2 class="page-title text-center">contact</h2>
        <div class="row">
            <div class="col-md-8 page-contents">
                <div class="row page-content">
                    <div class="contents row">
                        <br>
                        <form action="#" method="post" class="contact-form row">
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
                                <input type="url" id="website" class="form-control" placeholder="Your website url">
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
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <div class="col-md-4 sidebar">
                <!--Author Widget-->
                <aside class="row m0 widget-author widget has-pt">
                    <div class="widget-author-inner row">
                        <div class="author-avatar row"><a href="#"><img src="static/images/author.png" alt="" class="img-circle"></a></div>
                        <a href="#"><h2 class="author-name">{{.author}}</h2></a>
                        <h5 class="author-title">{{.professor}}</h5>
                        <p>Some people hear their own inner voices with great clearness. And they live by what they hear. Such people become crazy, or they become legends.</p>
                    </div>
                </aside>
                <!--Twitter Widget-->
                <aside class="row m0 widget widget-twitter">
                    <div class="widget-twitter-inner">
                        <div class="row tweet-texts">
                            <p>Keep Trying To Change The Fucking World</p>
                        </div>
                        <div class="row timepast">{{ d }}</div>
                    </div>
                </aside>
            </div>
        </div>
    </div>
</section>

{{template "footer.tpl"}}

