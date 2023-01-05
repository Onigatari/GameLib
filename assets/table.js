$(document).ready(function () {
    $('#gamesList').DataTable({
        "language": {
            "search": "Поиск:",
            "lengthMenu":     "Кол-во игр на страницы _MENU_",
            "loadingRecords": "Загрузка...",
            "processing":     "",
            "zeroRecords":    "Не найдено совпадающих игр",
            "paginate": {
                "first":      "Первый",
                "last":       "Последний",
                "next":       "Следующий",
                "previous":   "Предыдущий"
            },
        },
        ordering: false,
        info: false,
        lengthMenu: [
            [10, 50, 100, -1],
            [10, 50, 100, 'All'],
        ],
    });
});
