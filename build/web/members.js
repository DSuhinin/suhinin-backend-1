$(document).ready(function() {
    const token = localStorage.getItem("token");
    if(!token) {
        window.location.replace("index.html");
        return
    }

    $.ajax({
        type: "GET",
        url: 'http://localhost:8080/members',
        async: false,
        contentType: 'application/json',
        headers: {
            "Authorization": `Bearer ${token}`
        },
        statusCode: {
            200: function(data) {
                $('#text').text(data.text)
            },
            401: function() {
                localStorage.removeItem("token");
                window.location.replace("index.html");
            }
        },
    })
});
