$(document).ready(function () {
    $('#gamesList').DataTable({
        ordering: false,
        info: false,
        lengthMenu: [
            [10, 50, 100, -1],
            [10, 50, 100, 'All'],
        ],
    });
});