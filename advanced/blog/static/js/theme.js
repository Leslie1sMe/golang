(function ($) {
    "use strict";  
    
    /*Show Menu*/
    $('.menu-trigger').on('click', function(){
        if ( $(this).hasClass('active') ){
            $(this).removeClass('active');
            $('#header').removeClass('menu-active')
        }else{
            $(this).addClass('active');
            $('#header').addClass('menu-active')
        }
    });
        
    /*Header Sticky*/
    $(window).scroll(function() {
        var scroll = $(window).scrollTop();

        if ( scroll >= 100 ) {
            $("#header").addClass("add-bg")
        } else {
            $("#header").removeClass("add-bg")
        }
    }); 
    
    /*==========================================================================================*/
    /*==============================Function Declaration==============================*/
    /*Featured Post Carousel*/
    function featuredPostCarosel(){
        if ( $('.featured-post-carousel').length ){
            $('.featured-post-carousel').owlCarousel({
                items: 1,
                margin: 0,
                loop: false,
                nav: true,
                navContainer: '.featured-post-carousel',
                navText: ['<i class="fa fa-angle-left"></i>','<i class="fa fa-angle-right"></i>']
            })
        }
    }
    
    /*Gallery Post*/    
    function galleryOfPost () {
        if( $('.gallery-of-post').length ){
            $('.gallery-of-post').each(function(){
                $('.gallery-of-post').owlCarousel({
                    items: 1,
                    margin: 0,
                    loop: false,
                    nav: false,
                    dots: true
                })
            })
        }
    }
    
    /*Gallery Post*/    
    function galleryOfPost2 () {
        if( $('.gallery-of-post2').length ){
            $('.gallery-of-post2').each(function(){
                $('.gallery-of-post2').owlCarousel({
                    items: 2,
                    margin: 0,
                    loop: false,
                    nav: true,
                    navText: ['<i class="fa fa-angle-left"></i>','<i class="fa fa-angle-right"></i>'],
                    dots: true,
                    responsiveClass:true,
                    responsive:{
                        0:{
                            items:1,
                            dots: false
                        },
                        992:{
                            items:2
                        }
                    }
                })
            })
        }
    }
    
    /*Instafeed Widget*/
    function instafeedWidget(){
        if( $('#instafeed').length ){
            var userFeed = new Instafeed({
                get: 'user',
                tagName: 'awesome',
                userId: '1101771199',
                resolution: 'standard_resolution',
                accessToken: '1101771199.1677ed0.64ed520ad0c94f358dd92dddc082c33f',
                limit: 2,
                template: '<div class="item"><a href="{{image}}" data-source="{{image}}" title="{{caption}}"><img src="{{image}}" alt="{{caption}}></a></div>',
                after: function () {
                    for (var i = 0; i < 1; i++) {
                        var $newdiv = $('<div class="item"></div>').html('<img src="http://placehold.it/360x300">');
                        $('#instafeed').append($newdiv)
                    }
                    var owl = $('#instafeed');
                    owl.owlCarousel({
                        items: 1,
                        loop: true,
                        margin: 0,
                        nav: false,
                        autoplay: true,
                        animateOut: 'fadeOut',
                        animateIn: 'fadeIn'
                    })
                }
            });
            userFeed.run()
        }
    }
    
    function zoomGallery () {
        if ( $('#instafeed').length ){
            $('#instafeed').magnificPopup({
                delegate: 'a',
                type: 'image',
                closeOnContentClick: false,
                closeBtnInside: false,
                mainClass: 'mfp-with-zoom mfp-img-mobile',
                image: {
                    verticalFit: true,
                    titleSrc: function(item) {
                        return item.el.attr('title') + ' &middot; <a class="image-source-link" href="'+item.el.attr('data-source')+'" target="_blank">image source</a>'
                    }
                },
                gallery: {
                    enabled: true
                },
                zoom: {
                    enabled: true,
                    duration: 300, // don't foget to change the duration also in CSS
                    opener: function(element) {
                        return element.find('img')
                    }
                }

            })
        }
    }
    
    /*Post Masonry*/
    function postMasonry(){
        if ( $('#post-masonry').length ){
            var $container = $('#post-masonry');
            
            $container.imagesLoaded( function() {
                $container.isotope({
                  itemSelector: '.post',
                  layoutMode: 'masonry',
                })
            })            
        }
    }
    
    /*Post Share*/
    function postShare () {
        if( $('.post-share').length ){
            $('.post-share button').on('click', function(){
                if ( $(this).parent().hasClass('active') ){
                    $(this).parent().removeClass('active')
                }else{
                    $(this).parent().addClass('active')
                }
            })
        }
    }
    
    function thumbSilder(){
          // The slider being synced must be initialized first
        if( $('.thumbCarousel').length ){
            $('#carousel').flexslider({
                animation: "slide",
                controlNav: false,
                animationLoop: false,
                prevText: '<i class="fa fa-caret-left"></i>',
                nextText: '<i class="fa fa-caret-right"></i>',    
                slideshow: false,
                itemWidth: 98,
                itemMargin: 10,
                asNavFor: '#slider'
            });

            $('#slider').flexslider({
                animation: "slide",
                controlNav: true,
                directionNav: false,
                animationLoop: false,
                slideshow: false,
                sync: "#carousel"
            })
        }
    }
    
    /*==========================================================================================*/
    /*Function Call*/
    featuredPostCarosel();
    instafeedWidget();
    galleryOfPost ();
    galleryOfPost2 ();
    postMasonry();
    zoomGallery();
    postShare ();
    thumbSilder()
    
    
})(jQuery)