function createCards() {
    fetch('/api')
        .then(function (response) {
            return response.json();
        })
        .then(function (jsonArray) {
            jsonArray.forEach(function (i) {
                const ip = i.ip;
                const generation = i.generation;
                const height = i.height;
                const version = i.version;
                const workTime = i.work_time;
                const minedEver = i.mined_ever;
                const minedToday = i.mined_today;
                const nodeStatus = i.node_status;

                blocksToday += minedToday;

                const existingCard = document.querySelector(`.node-card[data-ip="${ip}"]`);
                if (existingCard) {
                    updateCard(existingCard, height, version, generation, workTime, minedEver, nodeStatus);
                } else {
                    createCard(ip, height, version, generation, workTime, minedEver, nodeStatus);
                }
            });
        })
        .catch(function (error) {
            console.log('Ошибка:', error);
        });
}