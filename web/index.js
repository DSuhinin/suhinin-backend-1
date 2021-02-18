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
        success: function () {
            window.location.replace("members.html");
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
        success: function () {
            alert("Success Signup!");
        }
    })
});

$('#logout').click(function (e) {
    e.preventDefault();
    alert(localStorage.getItem("jwt-token"));
});



