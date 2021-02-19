$(".tab a").on("click", function (e) {
    e.preventDefault();

    $(this).parent().addClass("active");
    $(this).parent().siblings().removeClass("active");

    target = $(this).attr("href");

    $(".tab-content > div").not(target).hide();

    $(target).fadeIn(600);
    $("input").val("")
});

$('#login-btn').click(function (e) {
    e.preventDefault();
    $.ajax({
        type: "POST",
        url: 'http://localhost:8080/auth/signin',
        contentType: 'application/json',
        async: false,
        data: JSON.stringify({
            email: $('#signin input[name="email"]').val(),
            password: $('#signin input[name="password"]').val(),
        }),
        statusCode: {
            200: function(resp) {
                localStorage.setItem("token", resp.token);
                window.location.replace("members.html");
            },
            400: function () {
                alert("Incorrect credentials");
            },
            404: function () {
                alert("Incorrect credentials");
            },
            401: function() {
                alert("Incorrect credentials");
            }
        }
    })
});

$('#register-btn').click(function (e) {
    e.preventDefault();
    $.ajax({
        type: "POST",
        url: 'http://localhost:8080/auth/signup',
        contentType: 'application/json',
        async: false,
        data: JSON.stringify({
            email: $('#signup input[name="email"]').val(),
            password: $('#signup input[name="password"]').val(),
            confirm_password: $('#signup input[name="confirm-password"]').val(),
        }),
        statusCode: {
            200: function() {
                window.location.replace("index.html");
            },
            400: function(data) {
                alert(data.responseJSON.message);
            }
        }
    })
});

$('#logout').click(function (e) {
    e.preventDefault();
    $.ajax({
        type: "GET",
        url: 'http://localhost:8080/auth/signout',
        async: false,
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        statusCode: {
            200: function() {
                localStorage.removeItem("token");
                window.location.replace("index.html");
            },
            401: function() {
                localStorage.removeItem("token");
                window.location.replace("index.html");
            }
        },
    })
});



