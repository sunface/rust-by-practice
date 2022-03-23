(function() {
    var path = window.location.pathname;
    if (path.endsWith("/print.html")) {
        return;
    }

    var images = document.querySelectorAll("main img")
    Array.prototype.forEach.call(images, function(img) {
        img.addEventListener("click", function() {
            BigPicture({
                el: img,
            });
        });
    });

    // Un-active everything when you click it
    Array.prototype.forEach.call(document.getElementsByClassName("pagetoc")[0].children, function(el) {
        el.addEventHandler("click", function() {
            Array.prototype.forEach.call(document.getElementsByClassName("pagetoc")[0].children, function(el) {
                el.classList.remove("active");
            });
            el.classList.add("active");
        });
    });

    var updateFunction = function() {

        var id;
        var elements = document.getElementsByClassName("header");
        Array.prototype.forEach.call(elements, function(el) {
            if (window.pageYOffset >= el.offsetTop) {
                id = el;
            }
        });

        Array.prototype.forEach.call(document.getElementsByClassName("pagetoc")[0].children, function(el) {
            el.classList.remove("active");
        });

        Array.prototype.forEach.call(document.getElementsByClassName("pagetoc")[0].children, function(el) {
            if (id.href.localeCompare(el.href) == 0) {
                el.classList.add("active");
            }
        });
    };

    // Populate sidebar on load
    window.addEventListener('load', function() {
        var pagetoc = document.getElementsByClassName("pagetoc")[0];
        var elements = document.getElementsByClassName("header");
        Array.prototype.forEach.call(elements, function(el) {
            var link = document.createElement("a");

            // Indent shows hierarchy
            var indent = "";
            switch (el.parentElement.tagName) {
                case "H1":
                    return;
                // case "H2":
                //     indent = "20px";
                //     break;
                case "H3":
                    indent = "20px";
                    break;
                case "H4":
                    indent = "40px";
                    break;
                default:
                    break;
            }

            link.appendChild(document.createTextNode(el.text));
            link.style.paddingLeft = indent;
            link.href = el.href;
            pagetoc.appendChild(link);
        });
        updateFunction.call();
    });

    // Handle active elements on scroll
    window.addEventListener("scroll", updateFunction);

    var p = path.replace("index.html", "");
    p = p.replace(".html", "");
    var strs = p.split("/");
    if (strs[strs.length-1] == ""){
        strs.pop()
    } 
    var str = strs[strs.length-1];
    var title = document.querySelector("main>h1,h2>a").textContent
    var gitalk = new Gitalk({
        clientID: '8e4b2cf9529ebb3dcad6',
        clientSecret: '6f6e8c23575a780bdb1faba3c17be08d76dc35f8',
        repo: 'rust-by-practice-comments',
        owner: 'sunface',
        admin: ["sunface"], 
        labels: ['comments'],
        title: title,
        createIssueManually: false,
        id: str,
        distractionFreeMode: true
    });
    gitalk.render('gitalk-container');
})();