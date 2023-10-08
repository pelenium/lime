const urlParams = new URLSearchParams(window.location.href.split('?')[1]);
const reqValue = urlParams.get('req');

fetch(`/api/${reqValue}`)
    .then(response => {
        if (!response.ok) {
            throw new Error(`Ошибка HTTP: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('Произошла ошибка:', error);
    });

