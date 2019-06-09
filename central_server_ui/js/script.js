let siteurl = '';
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
            name1: '',
            name2: '',
            name3: '',
            schoolName: 'Средняя общеобразовательная школа № 9 им. А.С. Пушкина, г. Пермь',
            schoolDocID: '07718000264307',
            inn: '123123123123',
            snils: '123-123-123-12',
            passport: {
                passportNumber: '1221 980890',
                gettingDate: '09.09.1999',
                unitCode: '123-123',
                whoGave: 'Отделом УФМС №123 г. Пермь',
                registrationAddress: 'СПб, Невский проспект, 12/2, 12',
                birthday: '',
                birthPlace: 'г. Архангельск',
            },
            contacts: {
                phone: '(123)123-12-12',
                email: 'test@test.test',
            },
            livingPlace: 'Пермский край, г. Пермь, пр. Декабристов, 11/12 - 16',
            parentsData: 'Белова Анна Николаевна ОАО "Авиадвигатель" +7 911 123-12-11, Белов Андрей Витальевич МОАУ СОШ №28  +7 911 123-12-22',

        },
        showSavedAlert: false
    },
    beforeMount: function (){
        axios({
            method: 'get',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            url: siteurl + '/abiturient'
        })
            .then ((response) => {
                this.profile.name1 = response.data.first_name;
                this.profile.name2 = response.data.last_name;
                this.profile.name3 = response.data.middle_name;
                this.profile.passport.birthday = response.data.birth_date.substr(0,10);
                this.profile.passport.birthPlace = response.data.birth_place;
                this.profile.passport.registrationAddress = response.data.address;
                this.profile.contacts.phone = response.data.phone_number;

                console.log("response.data = ");
                console.log(response.data);
            })
            .catch (function (error){
                console.log(error)
            });
    },

    methods: {

    }
});










