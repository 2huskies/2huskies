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
            photo: '../html/img/photo.png',
            name1: 'Василиса',
            name2: 'Белова',
            name3: 'Андреевна',
            schoolName: '',
            schoolID: '',
            inn: '123123123123',
            snils: '123-123-123-12',
            passport:{
                passportNumber: '1221 980890',
                gettingDate: '09.09.1999',
                unitCode: '123-123',
                whoGave: 'Отделом УФМС №123 г. Пермь',
                registrationAddress: '',
                birthday: '01.01.1980',
                birthPlace: '',
            },
            contacts: {
                phone: '(123)123-12-12',
                email: 'test@test.test',
            },
            livingPlace: '',
            parentsData: '',

        }
    },
    methods: {

    }
});










