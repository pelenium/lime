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
        createSiteCards(data)
    })
    .catch(error => {
        console.error('Произошла ошибка:', error);
    });


function createSiteCards(jsonData) {
    const mainPanel = document.getElementById("main-panel");

    jsonData.forEach(item => {
        const card = document.createElement("div");
        card.classList.add("card");

        const titleLink = document.createElement("a");
        titleLink.textContent = item.title;
        titleLink.href = item.url;
        titleLink.id = "site-title"

        const title = document.createElement("h2");
        title.appendChild(titleLink);

        const domain = new URL(item.url).hostname;

        const link = document.createElement("a");
        link.href = domain;
        link.textContent = domain;

        card.appendChild(title);
        card.appendChild(link);

        mainPanel.appendChild(card);
    });
}