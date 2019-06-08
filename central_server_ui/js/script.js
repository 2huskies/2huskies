let app = new Vue({
    el: '#app',
    data: {
        menuLinks: [
            { text: 'Моя анкета', href: "/profile" },
            { text: 'Мои заявки', href: "/orders" },
            { text: 'ВУЗы', href: "/institutes" },
            { text: 'Справка', href: "/help" }
        ],
        profile: {
            avatar: '../html/img/photo_s.jpg',
            fio: 'Фамилия Имя Отчество'
        }
    },
    methods: {

    }
});










