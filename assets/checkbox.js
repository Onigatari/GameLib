$(':checkbox').change(function () {
    $.post("/updateDone", {updateDoneGame: $(this).val()}).done(function () {
        window.location.reload();
    });
});
