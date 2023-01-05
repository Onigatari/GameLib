$(document).on('click', '#buttonRandom', function () {
    $.ajax({
        url: '/api/random',
        method: 'get',
        dataType: 'json',
        success: function (data) {
            const element = $("#labelRandomGame");
            const newText = data.name;
            element.text(newText);
            $("#randomGame").show();
        }
    });
});
