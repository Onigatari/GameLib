$(':checkbox').change(function () {
    $.post("/api/updateDone", {updateDoneGame: $(this).val()}).done(function () {
        window.location.reload();
    });
});
